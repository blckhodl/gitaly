package operations

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	gitalyerrors "gitlab.com/gitlab-org/gitaly/v15/internal/errors"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git/catfile"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git/localrepo"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git/updateref"
	"gitlab.com/gitlab-org/gitaly/v15/internal/gitaly/hook"
	"gitlab.com/gitlab-org/gitaly/v15/internal/helper"
	"gitlab.com/gitlab-org/gitaly/v15/proto/go/gitalypb"
)

func validateUserDeleteTagRequest(in *gitalypb.UserDeleteTagRequest) error {
	if in.GetRepository() == nil {
		return gitalyerrors.ErrEmptyRepository
	}
	if len(in.GetTagName()) == 0 {
		return errors.New("empty tag name")
	}
	if in.GetUser() == nil {
		return errors.New("empty user")
	}
	return nil
}

//nolint: stylecheck // This is unintentionally missing documentation.
func (s *Server) UserDeleteTag(ctx context.Context, req *gitalypb.UserDeleteTagRequest) (*gitalypb.UserDeleteTagResponse, error) {
	if err := validateUserDeleteTagRequest(req); err != nil {
		return nil, helper.ErrInvalidArgument(err)
	}
	referenceName := git.ReferenceName(fmt.Sprintf("refs/tags/%s", req.TagName))
	revision, err := s.localrepo(req.GetRepository()).ResolveRevision(ctx, referenceName.Revision())
	if err != nil {
		if errors.Is(err, git.ErrReferenceNotFound) {
			return nil, helper.ErrFailedPreconditionf("tag not found: %s", req.TagName)
		}
		return nil, helper.ErrInternalf("resolve revision %q: %w", referenceName, err)
	}

	if err := s.updateReferenceWithHooks(ctx, req.Repository, req.User, nil, referenceName, git.ObjectHashSHA1.ZeroOID, revision); err != nil {
		var customHookErr updateref.CustomHookError
		if errors.As(err, &customHookErr) {
			return &gitalypb.UserDeleteTagResponse{
				PreReceiveError: customHookErr.Error(),
			}, nil
		}

		var updateRefError updateref.Error
		if errors.As(err, &updateRefError) {
			return nil, helper.ErrFailedPrecondition(err)
		}

		return nil, helper.ErrInternal(err)
	}

	return &gitalypb.UserDeleteTagResponse{}, nil
}

func validateUserCreateTag(req *gitalypb.UserCreateTagRequest) error {
	if len(req.TagName) == 0 {
		return errors.New("empty tag name")
	}

	if err := git.ValidateRevision(req.TagName); err != nil {
		return fmt.Errorf("invalid tag name: %w", err)
	}

	if req.User == nil {
		return errors.New("empty user")
	}

	if len(req.TargetRevision) == 0 {
		return errors.New("empty target revision")
	}

	if bytes.Contains(req.Message, []byte("\000")) {
		return errors.New("tag message contains NUL byte")
	}

	if req.GetRepository() == nil {
		return gitalyerrors.ErrEmptyRepository
	}

	return nil
}

//nolint: stylecheck // This is unintentionally missing documentation.
func (s *Server) UserCreateTag(ctx context.Context, req *gitalypb.UserCreateTagRequest) (*gitalypb.UserCreateTagResponse, error) {
	if err := validateUserCreateTag(req); err != nil {
		return nil, helper.ErrInvalidArgumentf("validating request: %w", err)
	}

	targetRevision := git.Revision(req.TargetRevision)
	referenceName := git.ReferenceName(fmt.Sprintf("refs/tags/%s", req.TagName))

	committerTime, err := dateFromProto(req)
	if err != nil {
		return nil, helper.ErrInvalidArgument(err)
	}

	quarantineDir, quarantineRepo, err := s.quarantinedRepo(ctx, req.GetRepository())
	if err != nil {
		return nil, err
	}

	tag, tagID, err := s.createTag(ctx, quarantineRepo, targetRevision, req.TagName, req.Message, req.User, committerTime)
	if err != nil {
		return nil, err
	}

	// We check whether the reference exists up-front so that we can return a proper
	// detailed error to the caller in case the tag cannot be created. While this is
	// racy given that the tag can be created afterwards by a concurrent caller, we'd
	// detect that raciness in `updateReferenceWithHooks()` anyway.
	if oid, err := quarantineRepo.ResolveRevision(ctx, referenceName.Revision()); err != nil {
		// If the reference wasn't found we're on the happy path, otherwise
		// something has gone wrong.
		if !errors.Is(err, git.ErrReferenceNotFound) {
			return nil, helper.ErrInternalf("resolving target reference: %w", err)
		}
	} else {
		// When resolving the revision succeeds the tag reference exists already.
		// Because we don't want to overwrite preexisting references in this RPC we
		// thus return an error and alert the caller of this condition.
		detailedErr, err := helper.ErrWithDetails(
			helper.ErrAlreadyExistsf("tag reference exists already"),
			&gitalypb.UserCreateTagError{
				Error: &gitalypb.UserCreateTagError_ReferenceExists{
					ReferenceExists: &gitalypb.ReferenceExistsError{
						ReferenceName: []byte(referenceName.String()),
						Oid:           oid.String(),
					},
				},
			},
		)
		if err != nil {
			return nil, helper.ErrInternalf("creating detailed error: %w", err)
		}

		return nil, detailedErr
	}

	if err := s.updateReferenceWithHooks(ctx, req.Repository, req.User, quarantineDir, referenceName, tagID, git.ObjectHashSHA1.ZeroOID); err != nil {
		var notAllowedError hook.NotAllowedError
		var customHookErr updateref.CustomHookError
		var updateRefError updateref.Error

		if errors.As(err, &notAllowedError) {
			detailedErr, err := helper.ErrWithDetails(
				helper.ErrPermissionDeniedf("reference update denied by access checks: %w", err),
				&gitalypb.UserCreateTagError{
					Error: &gitalypb.UserCreateTagError_AccessCheck{
						AccessCheck: &gitalypb.AccessCheckError{
							ErrorMessage: notAllowedError.Message,
							UserId:       notAllowedError.UserID,
							Protocol:     notAllowedError.Protocol,
							Changes:      notAllowedError.Changes,
						},
					},
				},
			)
			if err != nil {
				return nil, helper.ErrInternalf("error details: %w", err)
			}

			return nil, detailedErr
		} else if errors.As(err, &customHookErr) {
			detailedErr, err := helper.ErrWithDetails(
				// We explicitly don't include the custom hook error itself
				// in the returned error because that would also contain the
				// standard output or standard error in the error message.
				// It's thus needlessly verbose and duplicates information
				// we have available in the structured error anyway.
				helper.ErrPermissionDeniedf("reference update denied by custom hooks"),
				&gitalypb.UserCreateTagError{
					Error: &gitalypb.UserCreateTagError_CustomHook{
						CustomHook: customHookErr.Proto(),
					},
				},
			)
			if err != nil {
				return nil, helper.ErrInternalf("error details: %w", err)
			}

			return nil, detailedErr
		} else if errors.As(err, &updateRefError) {
			detailedErr, err := helper.ErrWithDetails(
				helper.ErrFailedPreconditionf("reference update failed: %w", err),
				&gitalypb.UserCreateTagError{
					Error: &gitalypb.UserCreateTagError_ReferenceUpdate{
						ReferenceUpdate: &gitalypb.ReferenceUpdateError{
							ReferenceName: []byte(updateRefError.Reference.String()),
							OldOid:        updateRefError.OldOID.String(),
							NewOid:        updateRefError.NewOID.String(),
						},
					},
				},
			)
			if err != nil {
				return nil, helper.ErrInternalf("error details: %w", err)
			}

			return nil, detailedErr
		}

		return nil, helper.ErrInternalf("updating reference: %w", err)
	}

	return &gitalypb.UserCreateTagResponse{
		Tag: tag,
	}, nil
}

func (s *Server) createTag(
	ctx context.Context,
	repo *localrepo.Repo,
	targetRevision git.Revision,
	tagName []byte,
	message []byte,
	committer *gitalypb.User,
	committerTime time.Time,
) (*gitalypb.Tag, git.ObjectID, error) {
	objectReader, cancel, err := s.catfileCache.ObjectReader(ctx, repo)
	if err != nil {
		return nil, "", helper.ErrInternalf("creating object reader: %w", err)
	}
	defer cancel()

	objectInfoReader, cancel, err := s.catfileCache.ObjectInfoReader(ctx, repo)
	if err != nil {
		return nil, "", helper.ErrInternalf("creating object info reader: %w", err)
	}
	defer cancel()

	// We allow all ways to name a revision that cat-file
	// supports, not just OID. Resolve it.
	targetInfo, err := objectInfoReader.Info(ctx, targetRevision)
	if err != nil {
		return nil, "", helper.ErrFailedPreconditionf("revspec '%s' not found", targetRevision)
	}
	targetObjectID, targetObjectType := targetInfo.Oid, targetInfo.Type

	// Whether we have a "message" parameter tells us if we're
	// making a lightweight tag or an annotated tag. Maybe we can
	// use strings.TrimSpace() eventually, but as the tests show
	// the Ruby idea of Unicode whitespace is different than its
	// idea.
	makingTag := regexp.MustCompile(`\S`).Match(message)

	// If we're creating a tag to another "tag" we'll need to peel
	// (possibly recursively) all the way down to the
	// non-tag. That'll be our returned TargetCommit if it's a
	// commit (nil if tree or blob).
	//
	// If we're not pointing to a tag we pretend our "peeled" is
	// the commit/tree/blob object itself in subsequent logic.
	peeledTargetObjectID, peeledTargetObjectType := targetObjectID, targetObjectType
	if targetObjectType == "tag" {
		peeledTargetObjectInfo, err := objectInfoReader.Info(ctx, targetRevision+"^{}")
		if err != nil {
			return nil, "", helper.ErrInternalf("read object info: %w", err)
		}
		peeledTargetObjectID, peeledTargetObjectType = peeledTargetObjectInfo.Oid, peeledTargetObjectInfo.Type

		// If we're not making an annotated tag ourselves and
		// the user pointed to a "tag" we'll ignore what they
		// asked for and point to what we found when peeling
		// that tag.
		if !makingTag {
			targetObjectID = peeledTargetObjectID
		}
	}

	// At this point we'll either be pointing to an object we were
	// provided with, or creating a new tag object and pointing to
	// that.
	refObjectID := targetObjectID
	var tagObject *gitalypb.Tag
	if makingTag {
		tagObjectID, err := repo.WriteTag(ctx, targetObjectID, targetObjectType, tagName, message, committer, committerTime)
		if err != nil {
			var FormatTagError localrepo.FormatTagError
			if errors.As(err, &FormatTagError) {
				return nil, "", helper.ErrUnknownf("Rugged::InvalidError: failed to parse signature - expected prefix doesn't match actual")
			}

			var MktagError localrepo.MktagError
			if errors.As(err, &MktagError) {
				return nil, "", helper.ErrNotFoundf("Gitlab::Git::CommitError: %s", err.Error())
			}
			return nil, "", helper.ErrInternalf("write tag: %w", err)
		}

		createdTag, err := catfile.GetTag(ctx, objectReader, tagObjectID.Revision(), string(tagName))
		if err != nil {
			return nil, "", helper.ErrInternalf("get tag: %w", err)
		}

		tagObject = &gitalypb.Tag{
			Name:        tagName,
			Id:          tagObjectID.String(),
			Message:     createdTag.Message,
			MessageSize: createdTag.MessageSize,
			// TargetCommit: is filled in below if needed
		}

		refObjectID = tagObjectID
	} else {
		tagObject = &gitalypb.Tag{
			Name: tagName,
			Id:   peeledTargetObjectID.String(),
			// TargetCommit: is filled in below if needed
		}
	}

	// Save ourselves looking this up earlier in case update-ref
	// died
	if peeledTargetObjectType == "commit" {
		peeledTargetCommit, err := catfile.GetCommit(ctx, objectReader, peeledTargetObjectID.Revision())
		if err != nil {
			return nil, "", helper.ErrInternalf("get commit: %w", err)
		}
		tagObject.TargetCommit = peeledTargetCommit
	}

	return tagObject, refObjectID, nil
}
