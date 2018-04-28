package repository

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	pb "gitlab.com/gitlab-org/gitaly-proto/go"
	"gitlab.com/gitlab-org/gitaly/internal/helper"
	"gitlab.com/gitlab-org/gitaly/internal/helper/housekeeping"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	worktreePrefix       = "gitlab-worktree"
	rebaseWorktreePrefix = "rebase"
	freshTimeout         = 15 * time.Minute
)

func (s *server) IsRebaseInProgress(ctx context.Context, req *pb.IsRebaseInProgressRequest) (*pb.IsRebaseInProgressResponse, error) {
	if err := validateIsRebaseInProgressRequest(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "IsRebaseInProgress: %v", err)
	}

	repoPath, err := helper.GetRepoPath(req.GetRepository())
	if err != nil {
		return nil, err
	}

	inProg, err := freshWorktree(repoPath, rebaseWorktreePrefix, req.GetRebaseId())
	if err != nil {
		return nil, err
	}
	return &pb.IsRebaseInProgressResponse{InProgress: inProg}, nil
}

func freshWorktree(repoPath, prefix, id string) (bool, error) {
	worktreePath := path.Join(repoPath, worktreePrefix, fmt.Sprintf("%s-%s", prefix, id))

	fs, err := os.Stat(worktreePath)
	if err != nil {
		return false, nil
	}

	if time.Since(fs.ModTime()) > freshTimeout {
		if err = os.RemoveAll(worktreePath); err != nil {
			if err = housekeeping.FixDirectoryPermissions(worktreePath); err != nil {
				return false, err
			}
			err = os.RemoveAll(worktreePath)
		}
		return false, err
	}

	return true, nil
}

func validateIsRebaseInProgressRequest(req *pb.IsRebaseInProgressRequest) error {
	if req.GetRepository() == nil {
		return fmt.Errorf("empty Repository")
	}

	if req.GetRebaseId() == "" {
		return fmt.Errorf("empty RebaseId")
	}

	if strings.Contains(req.GetRebaseId(), "/") {
		return fmt.Errorf("RebaseId contains '/'")
	}

	return nil
}
