package repository

import (
	"os"
	"path/filepath"

	"gitlab.com/gitlab-org/gitaly/v15/internal/command"
	"gitlab.com/gitlab-org/gitaly/v15/proto/go/gitalypb"
	"gitlab.com/gitlab-org/gitaly/v15/streamio"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const customHooksDir = "custom_hooks"

func (s *server) BackupCustomHooks(in *gitalypb.BackupCustomHooksRequest, stream gitalypb.RepositoryService_BackupCustomHooksServer) error {
	repoPath, err := s.locator.GetPath(in.Repository)
	if err != nil {
		return status.Errorf(codes.Internal, "BackupCustomHooks: getting repo path failed %v", err)
	}

	writer := streamio.NewWriter(func(p []byte) error {
		return stream.Send(&gitalypb.BackupCustomHooksResponse{Data: p})
	})

	if _, err := os.Lstat(filepath.Join(repoPath, customHooksDir)); os.IsNotExist(err) {
		return nil
	}

	ctx := stream.Context()
	tar := []string{"tar", "-c", "-f", "-", "-C", repoPath, customHooksDir}
	cmd, err := command.New(ctx, tar, command.WithStdout(writer))
	if err != nil {
		return status.Errorf(codes.Internal, "%v", err)
	}

	if err := cmd.Wait(); err != nil {
		return status.Errorf(codes.Internal, "%v", err)
	}

	return nil
}
