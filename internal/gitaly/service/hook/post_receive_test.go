package hook

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/gitaly/internal/gitaly/config"
	"gitlab.com/gitlab-org/gitaly/internal/helper/text"
	"gitlab.com/gitlab-org/gitaly/internal/metadata/featureflag"
	"gitlab.com/gitlab-org/gitaly/internal/praefect/metadata"
	"gitlab.com/gitlab-org/gitaly/internal/testhelper"
	"gitlab.com/gitlab-org/gitaly/proto/go/gitalypb"
	"gitlab.com/gitlab-org/gitaly/streamio"
	"google.golang.org/grpc/codes"
)

func TestPostReceiveInvalidArgument(t *testing.T) {
	serverSocketPath, stop := runHooksServer(t, config.Config.Hooks)
	defer stop()

	client, conn := newHooksClient(t, serverSocketPath)
	defer conn.Close()

	ctx, cancel := testhelper.Context()
	defer cancel()

	stream, err := client.PostReceiveHook(ctx)
	require.NoError(t, err)
	require.NoError(t, stream.Send(&gitalypb.PostReceiveHookRequest{}), "empty repository should result in an error")
	_, err = stream.Recv()

	testhelper.RequireGrpcError(t, err, codes.InvalidArgument)
}

func transactionEnv(t *testing.T, primary bool) string {
	t.Helper()

	transaction := metadata.Transaction{
		ID:      1234,
		Node:    "node-1",
		Primary: primary,
	}

	env, err := transaction.Env()
	require.NoError(t, err)

	return env
}

func TestPostReceive(t *testing.T) {
	rubyDir := config.Config.Ruby.Dir
	defer func(rubyDir string) {
		config.Config.Ruby.Dir = rubyDir
	}(rubyDir)

	cwd, err := os.Getwd()
	require.NoError(t, err)
	config.Config.Ruby.Dir = filepath.Join(cwd, "testdata")

	serverSocketPath, stop := runHooksServer(t, config.Config.Hooks)
	defer stop()

	testRepo, _, cleanupFn := testhelper.NewTestRepo(t)
	defer cleanupFn()

	client, conn := newHooksClient(t, serverSocketPath)
	defer conn.Close()

	testCases := []struct {
		desc   string
		stdin  io.Reader
		req    gitalypb.PostReceiveHookRequest
		status int32
		stdout string
		stderr string
	}{
		{
			desc:  "everything OK",
			stdin: bytes.NewBufferString("a\nb\nc\nd\ne\nf\ng"),
			req: gitalypb.PostReceiveHookRequest{
				Repository:           testRepo,
				EnvironmentVariables: []string{"GL_ID=key_id", "GL_USERNAME=username", "GL_PROTOCOL=protocol", "GL_REPOSITORY=repository"},
				GitPushOptions:       []string{"option0", "option1"}},
			status: 0,
			stdout: "OK",
			stderr: "",
		},
		{
			desc:  "missing stdin",
			stdin: bytes.NewBuffer(nil),
			req: gitalypb.PostReceiveHookRequest{
				Repository:           testRepo,
				EnvironmentVariables: []string{"GL_ID=key_id", "GL_USERNAME=username", "GL_PROTOCOL=protocol", "GL_REPOSITORY=repository"},
				GitPushOptions:       []string{"option0"},
			},
			status: 1,
			stdout: "",
			stderr: "FAIL",
		},
		{
			desc:  "missing gl_id",
			stdin: bytes.NewBufferString("a\nb\nc\nd\ne\nf\ng"),
			req: gitalypb.PostReceiveHookRequest{
				Repository:           testRepo,
				EnvironmentVariables: []string{"GL_ID=", "GL_USERNAME=username", "GL_PROTOCOL=protocol", "GL_REPOSITORY=repository"},
				GitPushOptions:       []string{"option0"},
			},
			status: 1,
			stdout: "",
			stderr: "FAIL",
		},
		{
			desc:  "missing gl_username",
			stdin: bytes.NewBufferString("a\nb\nc\nd\ne\nf\ng"),
			req: gitalypb.PostReceiveHookRequest{
				Repository:           testRepo,
				EnvironmentVariables: []string{"GL_ID=key-123", "GL_USERNAME=", "GL_PROTOCOL=protocol", "GL_REPOSITORY=repository"},
				GitPushOptions:       []string{"option0"},
			},
			status: 1,
			stdout: "",
			stderr: "FAIL",
		},
		{
			desc:  "missing gl_protocol",
			stdin: bytes.NewBufferString("a\nb\nc\nd\ne\nf\ng"),
			req: gitalypb.PostReceiveHookRequest{
				Repository:           testRepo,
				EnvironmentVariables: []string{"GL_ID=key-123", "GL_USERNAME=username", "GL_PROTOCOL=", "GL_REPOSITORY=repository"},
				GitPushOptions:       []string{"option0"},
			},
			status: 1,
			stdout: "",
			stderr: "FAIL",
		},
		{
			desc:  "missing gl_repository value",
			stdin: bytes.NewBufferString("a\nb\nc\nd\ne\nf\ng"),
			req: gitalypb.PostReceiveHookRequest{
				Repository:           testRepo,
				EnvironmentVariables: []string{"GL_ID=key-123", "GL_USERNAME=username", "GL_PROTOCOL=protocol", "GL_REPOSITORY="},
				GitPushOptions:       []string{"option0"},
			},
			status: 1,
			stdout: "",
			stderr: "FAIL",
		},
		{
			desc:  "missing git push option",
			stdin: bytes.NewBufferString("a\nb\nc\nd\ne\nf\ng"),
			req: gitalypb.PostReceiveHookRequest{
				Repository:           testRepo,
				EnvironmentVariables: []string{"GL_ID=key-123", "GL_USERNAME=username", "GL_PROTOCOL=protocol", "GL_REPOSITORY=repository"},
			},
			status: 1,
			stdout: "",
			stderr: "FAIL",
		},
		{
			desc:  "primary fails with missing stdin because hook gets executed",
			stdin: bytes.NewBuffer(nil),
			req: gitalypb.PostReceiveHookRequest{
				Repository: testRepo,
				EnvironmentVariables: []string{
					"GL_ID=key_id",
					"GL_USERNAME=username",
					"GL_PROTOCOL=protocol",
					"GL_REPOSITORY=repository",
					transactionEnv(t, true),
				},
				GitPushOptions: []string{"option0"},
			},
			status: 1,
			stdout: "",
			stderr: "FAIL",
		},
		{
			desc:  "secondary succeeds with missing stdin because hook does not get executed",
			stdin: bytes.NewBuffer(nil),
			req: gitalypb.PostReceiveHookRequest{
				Repository: testRepo,
				EnvironmentVariables: []string{
					"GL_ID=key_id",
					"GL_USERNAME=username",
					"GL_PROTOCOL=protocol",
					"GL_REPOSITORY=repository",
					transactionEnv(t, false),
				},
				GitPushOptions: []string{"option0"},
			},
			status: 0,
			stdout: "",
			stderr: "",
		},
		{
			desc:  "Go hook correctly honors the primary flag",
			stdin: bytes.NewBuffer(nil),
			req: gitalypb.PostReceiveHookRequest{
				Repository: testRepo,
				EnvironmentVariables: []string{
					"GL_ID=key_id",
					"GL_USERNAME=username",
					"GL_PROTOCOL=protocol",
					"GL_REPOSITORY=repository",
					transactionEnv(t, false),
					fmt.Sprintf("%s=true", featureflag.GoPostReceiveHookEnvVar),
				},
				GitPushOptions: []string{"option0"},
			},
			status: 0,
			stdout: "",
			stderr: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ctx, cancel := testhelper.Context()
			defer cancel()

			stream, err := client.PostReceiveHook(ctx)
			require.NoError(t, err)
			require.NoError(t, stream.Send(&tc.req))

			go func() {
				writer := streamio.NewWriter(func(p []byte) error {
					return stream.Send(&gitalypb.PostReceiveHookRequest{Stdin: p})
				})
				_, err := io.Copy(writer, tc.stdin)
				require.NoError(t, err)
				require.NoError(t, stream.CloseSend(), "close send")
			}()

			var status int32
			var stdout, stderr bytes.Buffer
			for {
				resp, err := stream.Recv()
				if err == io.EOF {
					break
				}

				_, err = stdout.Write(resp.GetStdout())
				require.NoError(t, err)
				stderr.Write(resp.GetStderr())
				status = resp.GetExitStatus().GetValue()
			}

			assert.Equal(t, tc.status, status)
			assert.Equal(t, tc.stderr, text.ChompBytes(stderr.Bytes()), "hook stderr")
			assert.Equal(t, tc.stdout, text.ChompBytes(stdout.Bytes()), "hook stdout")
		})
	}
}

func TestPostReceiveMessages(t *testing.T) {
	testCases := []struct {
		desc                         string
		basicMessages, alertMessages []string
		expectedStdout               string
	}{
		{
			desc:          "basic MR message",
			basicMessages: []string{"To create a merge request for okay, visit:\n  http://localhost/project/-/merge_requests/new?merge_request"},
			expectedStdout: `
To create a merge request for okay, visit:
  http://localhost/project/-/merge_requests/new?merge_request
`,
		},
		{
			desc:          "alert",
			alertMessages: []string{"something went very wrong"},
			expectedStdout: `
========================================================================

                       something went very wrong

========================================================================
`,
		},
	}

	testRepo, testRepoPath, cleanupFn := testhelper.NewTestRepo(t)
	defer cleanupFn()

	secretToken := "secret token"
	user, password := "user", "password"

	tempDir, cleanup := testhelper.CreateTemporaryGitlabShellDir(t)
	defer cleanup()
	testhelper.WriteShellSecretFile(t, tempDir, secretToken)

	for _, tc := range testCases {
		for _, useGoPreReceive := range []bool{true, false} {
			t.Run(fmt.Sprintf("%s:use_go_pre_receive:%v", tc.desc, useGoPreReceive), func(t *testing.T) {
				c := testhelper.GitlabTestServerOptions{
					User:                        user,
					Password:                    password,
					SecretToken:                 secretToken,
					GLID:                        "key_id",
					GLRepository:                "repository",
					Changes:                     "changes",
					PostReceiveCounterDecreased: true,
					PostReceiveMessages:         tc.basicMessages,
					PostReceiveAlerts:           tc.alertMessages,
					Protocol:                    "protocol",
					RepoPath:                    testRepoPath,
				}

				serverURL, cleanup := testhelper.NewGitlabTestServer(t, c)
				defer cleanup()

				gitlabConfig := config.Gitlab{
					SecretFile: filepath.Join(tempDir, ".gitlab_shell_secret"),
					URL:        serverURL,
					HTTPSettings: config.HTTPSettings{
						User:     user,
						Password: password,
					},
				}

				defer func(cfg config.Cfg) {
					config.Config = cfg
				}(config.Config)

				config.Config.Gitlab = gitlabConfig

				api, err := NewGitlabAPI(gitlabConfig)
				require.NoError(t, err)

				serverSocketPath, stop := runHooksServerWithAPI(t, api, config.Config.Hooks)
				defer stop()

				client, conn := newHooksClient(t, serverSocketPath)
				defer conn.Close()

				ctx, cancel := testhelper.Context()
				defer cancel()

				stream, err := client.PostReceiveHook(ctx)
				require.NoError(t, err)

				envVars := []string{
					"GL_ID=key_id",
					"GL_USERNAME=username",
					"GL_PROTOCOL=protocol",
					"GL_REPOSITORY=repository"}
				if useGoPreReceive {
					envVars = append(envVars, "GITALY_GO_PRERECEIVE=true")
				}

				require.NoError(t, stream.Send(&gitalypb.PostReceiveHookRequest{
					Repository:           testRepo,
					EnvironmentVariables: envVars}))

				go func() {
					writer := streamio.NewWriter(func(p []byte) error {
						return stream.Send(&gitalypb.PostReceiveHookRequest{Stdin: p})
					})
					_, err := writer.Write([]byte("changes"))
					require.NoError(t, err)
					require.NoError(t, stream.CloseSend(), "close send")
				}()

				var status int32
				var stdout, stderr bytes.Buffer
				for {
					resp, err := stream.Recv()
					if err == io.EOF {
						break
					}

					_, err = stdout.Write(resp.GetStdout())
					require.NoError(t, err)
					stderr.Write(resp.GetStderr())
					status = resp.GetExitStatus().GetValue()
				}

				assert.Equal(t, int32(0), status)
				assert.Equal(t, "", text.ChompBytes(stderr.Bytes()), "hook stderr")
				assert.Equal(t, tc.expectedStdout, text.ChompBytes(stdout.Bytes()), "hook stdout")
			})
		}
	}
}

func TestPrintAlert(t *testing.T) {
	testCases := []struct {
		message  string
		expected string
	}{
		{
			message: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur nec mi lectus. Fusce eu ligula in odio hendrerit posuere. Ut semper neque vitae maximus accumsan. In malesuada justo nec leo congue egestas. Vivamus interdum nec libero ac convallis. Praesent euismod et nunc vitae vulputate. Mauris tincidunt ligula urna, bibendum vestibulum sapien luctus eu. Donec sed justo in erat dictum semper. Ut porttitor augue in felis gravida scelerisque. Morbi dolor justo, accumsan et nulla vitae, luctus consectetur est. Donec aliquet erat pellentesque suscipit elementum. Cras posuere eros ipsum, a tincidunt tortor laoreet quis. Mauris varius nulla vitae placerat imperdiet. Vivamus ut ligula odio. Cras nec euismod ligula.",
			expected: `========================================================================

   Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur
  nec mi lectus. Fusce eu ligula in odio hendrerit posuere. Ut semper
    neque vitae maximus accumsan. In malesuada justo nec leo congue
  egestas. Vivamus interdum nec libero ac convallis. Praesent euismod
    et nunc vitae vulputate. Mauris tincidunt ligula urna, bibendum
  vestibulum sapien luctus eu. Donec sed justo in erat dictum semper.
  Ut porttitor augue in felis gravida scelerisque. Morbi dolor justo,
  accumsan et nulla vitae, luctus consectetur est. Donec aliquet erat
      pellentesque suscipit elementum. Cras posuere eros ipsum, a
   tincidunt tortor laoreet quis. Mauris varius nulla vitae placerat
      imperdiet. Vivamus ut ligula odio. Cras nec euismod ligula.

========================================================================`,
		},
		{
			message: "Lorem ipsum dolor sit amet, consectetur",
			expected: `========================================================================

                Lorem ipsum dolor sit amet, consectetur

========================================================================`,
		},
	}

	for _, tc := range testCases {
		var result bytes.Buffer

		require.NoError(t, printAlert(PostReceiveMessage{Message: tc.message}, &result))
		assert.Equal(t, tc.expected, result.String())
	}
}