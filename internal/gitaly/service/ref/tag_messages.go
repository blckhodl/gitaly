package ref

import (
	"bytes"
	"fmt"
	"io"

	gitalyerrors "gitlab.com/gitlab-org/gitaly/v15/internal/errors"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git/catfile"
	"gitlab.com/gitlab-org/gitaly/v15/internal/helper"
	"gitlab.com/gitlab-org/gitaly/v15/proto/go/gitalypb"
	"gitlab.com/gitlab-org/gitaly/v15/streamio"
)

func (s *server) GetTagMessages(request *gitalypb.GetTagMessagesRequest, stream gitalypb.RefService_GetTagMessagesServer) error {
	if err := validateGetTagMessagesRequest(request); err != nil {
		return helper.ErrInvalidArgument(err)
	}
	if err := s.getAndStreamTagMessages(request, stream); err != nil {
		return helper.ErrInternal(err)
	}

	return nil
}

func validateGetTagMessagesRequest(request *gitalypb.GetTagMessagesRequest) error {
	if request.GetRepository() == nil {
		return gitalyerrors.ErrEmptyRepository
	}

	return nil
}

func (s *server) getAndStreamTagMessages(request *gitalypb.GetTagMessagesRequest, stream gitalypb.RefService_GetTagMessagesServer) error {
	ctx := stream.Context()
	repo := s.localrepo(request.GetRepository())

	objectReader, cancel, err := s.catfileCache.ObjectReader(ctx, repo)
	if err != nil {
		return fmt.Errorf("creating object reader: %w", err)
	}
	defer cancel()

	for _, tagID := range request.GetTagIds() {
		tag, err := catfile.GetTag(ctx, objectReader, git.Revision(tagID), "")
		if err != nil {
			return fmt.Errorf("failed to get tag: %w", err)
		}

		if err := stream.Send(&gitalypb.GetTagMessagesResponse{TagId: tagID}); err != nil {
			return err
		}
		sw := streamio.NewWriter(func(p []byte) error {
			return stream.Send(&gitalypb.GetTagMessagesResponse{Message: p})
		})

		msgReader := bytes.NewReader(tag.Message)

		if _, err = io.Copy(sw, msgReader); err != nil {
			return fmt.Errorf("failed to send response: %w", err)
		}
	}
	return nil
}
