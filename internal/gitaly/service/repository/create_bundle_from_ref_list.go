package repository

import (
	"bytes"
	"io"

	"gitlab.com/gitlab-org/gitaly/v15/internal/command"
	gitalyerrors "gitlab.com/gitlab-org/gitaly/v15/internal/errors"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git"
	"gitlab.com/gitlab-org/gitaly/v15/internal/helper"
	"gitlab.com/gitlab-org/gitaly/v15/proto/go/gitalypb"
	"gitlab.com/gitlab-org/gitaly/v15/streamio"
)

func (s *server) CreateBundleFromRefList(stream gitalypb.RepositoryService_CreateBundleFromRefListServer) error {
	firstRequest, err := stream.Recv()
	if err != nil {
		return err
	}

	if firstRequest.GetRepository() == nil {
		return helper.ErrInvalidArgument(gitalyerrors.ErrEmptyRepository)
	}

	ctx := stream.Context()

	if _, err := s.Cleanup(ctx, &gitalypb.CleanupRequest{Repository: firstRequest.GetRepository()}); err != nil {
		return err
	}

	firstRead := true
	reader := streamio.NewReader(func() ([]byte, error) {
		var request *gitalypb.CreateBundleFromRefListRequest
		if firstRead {
			firstRead = false
			request = firstRequest
		} else {
			var err error
			request, err = stream.Recv()
			if err != nil {
				return nil, err
			}
		}
		return append(bytes.Join(request.GetPatterns(), []byte("\n")), '\n'), nil
	})

	var stderr bytes.Buffer

	repo := s.localrepo(firstRequest.GetRepository())
	cmd, err := repo.Exec(ctx,
		git.SubSubCmd{
			Name:   "bundle",
			Action: "create",
			Flags: []git.Option{
				git.OutputToStdout,
				git.Flag{Name: "--ignore-missing"},
				git.Flag{Name: "--stdin"},
			},
		},
		git.WithStdin(reader),
		git.WithStderr(&stderr),
	)
	if err != nil {
		return helper.ErrInternalf("cmd start failed: %w", err)
	}

	writer := streamio.NewWriter(func(p []byte) error {
		return stream.Send(&gitalypb.CreateBundleFromRefListResponse{Data: p})
	})

	_, err = io.Copy(writer, cmd)
	if err != nil {
		return helper.ErrInternalf("stream writer failed: %w", err)
	}

	err = cmd.Wait()
	if isExitWithCode(err, 128) && bytes.HasPrefix(stderr.Bytes(), []byte("fatal: Refusing to create empty bundle.")) {
		return helper.ErrFailedPreconditionf("cmd wait failed: refusing to create empty bundle")
	} else if err != nil {
		return helper.ErrInternalf("cmd wait failed: %w, stderr: %q", err, stderr.String())
	}

	return nil
}

func isExitWithCode(err error, code int) bool {
	actual, ok := command.ExitStatus(err)
	if !ok {
		return false
	}

	return code == actual
}
