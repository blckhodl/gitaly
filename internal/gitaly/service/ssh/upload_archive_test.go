//go:build !gitaly_test_sha256

package ssh

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git/gittest"
	"gitlab.com/gitlab-org/gitaly/v15/internal/helper"
	"gitlab.com/gitlab-org/gitaly/v15/internal/testhelper"
	"gitlab.com/gitlab-org/gitaly/v15/internal/testhelper/testcfg"
	"gitlab.com/gitlab-org/gitaly/v15/proto/go/gitalypb"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestFailedUploadArchiveRequestDueToTimeout(t *testing.T) {
	t.Parallel()

	cfg := testcfg.Build(t)

	cfg.SocketPath = runSSHServerWithOptions(t, cfg, []ServerOpt{WithArchiveRequestTimeout(100 * time.Microsecond)})

	ctx := testhelper.Context(t)
	repo, _ := gittest.CreateRepository(ctx, t, cfg, gittest.CreateRepositoryConfig{
		Seed: gittest.SeedGitLabTest,
	})

	client := newSSHClient(t, cfg.SocketPath)

	stream, err := client.SSHUploadArchive(ctx)
	require.NoError(t, err)

	// The first request is not limited by timeout, but also not under attacker control
	require.NoError(t, stream.Send(&gitalypb.SSHUploadArchiveRequest{Repository: repo}))

	// Because the client says nothing, the server would block. Because of
	// the timeout, it won't block forever, and return with a non-zero exit
	// code instead.
	requireFailedSSHStream(t, helper.ErrDeadlineExceededf("waiting for packfile negotiation: context canceled"), func() (int32, error) {
		resp, err := stream.Recv()
		if err != nil {
			return 0, err
		}

		var code int32
		if status := resp.GetExitStatus(); status != nil {
			code = status.Value
		}

		return code, nil
	})
}

func TestFailedUploadArchiveRequestDueToValidationError(t *testing.T) {
	t.Parallel()

	cfg := testcfg.Build(t)

	serverSocketPath := runSSHServer(t, cfg)

	client := newSSHClient(t, serverSocketPath)

	tests := []struct {
		Desc string
		Req  *gitalypb.SSHUploadArchiveRequest
		Code codes.Code
	}{
		{
			Desc: "Repository.RelativePath is empty",
			Req:  &gitalypb.SSHUploadArchiveRequest{Repository: &gitalypb.Repository{StorageName: cfg.Storages[0].Name, RelativePath: ""}},
			Code: codes.InvalidArgument,
		},
		{
			Desc: "Repository is nil",
			Req:  &gitalypb.SSHUploadArchiveRequest{Repository: nil},
			Code: codes.InvalidArgument,
		},
		{
			Desc: "Data exists on first request",
			Req:  &gitalypb.SSHUploadArchiveRequest{Repository: &gitalypb.Repository{StorageName: cfg.Storages[0].Name, RelativePath: "path/to/repo"}, Stdin: []byte("Fail")},
			Code: func() codes.Code {
				if testhelper.IsPraefectEnabled() {
					return codes.NotFound
				}

				return codes.InvalidArgument
			}(),
		},
	}

	for _, test := range tests {
		t.Run(test.Desc, func(t *testing.T) {
			ctx := testhelper.Context(t)
			stream, err := client.SSHUploadArchive(ctx)
			if err != nil {
				t.Fatal(err)
			}

			if err = stream.Send(test.Req); err != nil {
				t.Fatal(err)
			}
			require.NoError(t, stream.CloseSend())

			err = testUploadArchiveFailedResponse(t, stream)
			testhelper.RequireGrpcCode(t, err, test.Code)
		})
	}
}

func TestUploadArchiveSuccess(t *testing.T) {
	t.Parallel()

	cfg := testcfg.Build(t)

	testcfg.BuildGitalySSH(t, cfg)

	cfg.SocketPath = runSSHServer(t, cfg)

	ctx := testhelper.Context(t)
	repo, _ := gittest.CreateRepository(ctx, t, cfg, gittest.CreateRepositoryConfig{
		Seed: gittest.SeedGitLabTest,
	})

	payload, err := protojson.Marshal(&gitalypb.SSHUploadArchiveRequest{
		Repository: repo,
	})
	require.NoError(t, err)

	gittest.ExecOpts(t, cfg, gittest.ExecConfig{
		Env: []string{
			fmt.Sprintf("GITALY_ADDRESS=%s", cfg.SocketPath),
			fmt.Sprintf("GITALY_PAYLOAD=%s", payload),
			fmt.Sprintf("PATH=%s", ".:"+os.Getenv("PATH")),
			fmt.Sprintf(`GIT_SSH_COMMAND=%s upload-archive`, cfg.BinaryPath("gitaly-ssh")),
		},
	}, "archive", "master", "--remote=git@localhost:test/test.git")
}

func testUploadArchiveFailedResponse(t *testing.T, stream gitalypb.SSHService_SSHUploadArchiveClient) error {
	var err error
	var res *gitalypb.SSHUploadArchiveResponse

	for err == nil {
		res, err = stream.Recv()
		require.Nil(t, res.GetStdout())
	}

	return err
}
