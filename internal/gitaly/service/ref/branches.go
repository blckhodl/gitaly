package ref

import (
	"context"
	"errors"

	gitalyerrors "gitlab.com/gitlab-org/gitaly/v15/internal/errors"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git"
	"gitlab.com/gitlab-org/gitaly/v15/internal/helper"
	"gitlab.com/gitlab-org/gitaly/v15/proto/go/gitalypb"
)

func (s *server) FindBranch(ctx context.Context, req *gitalypb.FindBranchRequest) (*gitalypb.FindBranchResponse, error) {
	if req.GetRepository() == nil {
		return nil, helper.ErrInvalidArgument(gitalyerrors.ErrEmptyRepository)
	}
	if len(req.GetName()) == 0 {
		return nil, helper.ErrInvalidArgumentf("Branch name cannot be empty")
	}

	repo := s.localrepo(req.GetRepository())

	branchName := git.NewReferenceNameFromBranchName(string(req.GetName()))
	branchRef, err := repo.GetReference(ctx, branchName)
	if err != nil {
		if errors.Is(err, git.ErrReferenceNotFound) {
			return &gitalypb.FindBranchResponse{}, nil
		}
		return nil, err
	}
	commit, err := repo.ReadCommit(ctx, git.Revision(branchRef.Target))
	if err != nil {
		return nil, err
	}

	branch, ok := branchName.Branch()
	if !ok {
		return nil, helper.ErrInvalidArgumentf("reference is not a branch")
	}

	return &gitalypb.FindBranchResponse{
		Branch: &gitalypb.Branch{
			Name:         []byte(branch),
			TargetCommit: commit,
		},
	}, nil
}
