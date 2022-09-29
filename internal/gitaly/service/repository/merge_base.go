package repository

import (
	"context"
	"io"

	gitalyerrors "gitlab.com/gitlab-org/gitaly/v15/internal/errors"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git"
	"gitlab.com/gitlab-org/gitaly/v15/internal/helper"
	"gitlab.com/gitlab-org/gitaly/v15/internal/helper/text"
	"gitlab.com/gitlab-org/gitaly/v15/proto/go/gitalypb"
)

func (s *server) FindMergeBase(ctx context.Context, req *gitalypb.FindMergeBaseRequest) (*gitalypb.FindMergeBaseResponse, error) {
	if req.GetRepository() == nil {
		return nil, helper.ErrInvalidArgument(gitalyerrors.ErrEmptyRepository)
	}
	var revisions []string
	for _, rev := range req.GetRevisions() {
		revisions = append(revisions, string(rev))
	}

	if len(revisions) < 2 {
		return nil, helper.ErrInvalidArgumentf("at least 2 revisions are required")
	}

	cmd, err := s.gitCmdFactory.New(ctx, req.GetRepository(),
		git.SubCmd{
			Name: "merge-base",
			Args: revisions,
		},
	)
	if err != nil {
		return nil, helper.ErrInternalf("cmd: %w", err)
	}

	mergeBase, err := io.ReadAll(cmd)
	if err != nil {
		return nil, helper.ErrInternalf("read output: %w", err)
	}

	mergeBaseStr := text.ChompBytes(mergeBase)

	if err := cmd.Wait(); err != nil {
		// On error just return an empty merge base
		return &gitalypb.FindMergeBaseResponse{Base: ""}, nil
	}

	return &gitalypb.FindMergeBaseResponse{Base: mergeBaseStr}, nil
}
