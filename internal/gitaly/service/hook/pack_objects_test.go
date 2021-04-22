package hook

import (
	"bytes"
	"context"
	"io"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/gitaly/internal/gitaly/config"
	"gitlab.com/gitlab-org/gitaly/internal/streamcache"
	"gitlab.com/gitlab-org/gitaly/internal/testhelper"
	"gitlab.com/gitlab-org/gitaly/internal/testhelper/testcfg"
	"gitlab.com/gitlab-org/gitaly/internal/testhelper/testserver"
	"gitlab.com/gitlab-org/gitaly/proto/go/gitalypb"
	"google.golang.org/grpc/codes"
)

func TestServer_PackObjectsHook_invalidArgument(t *testing.T) {
	_, repo, _, client := setupHookService(t)

	ctx, cancel := testhelper.Context()
	defer cancel()

	testCases := []struct {
		desc string
		req  *gitalypb.PackObjectsHookRequest
	}{
		{desc: "empty", req: &gitalypb.PackObjectsHookRequest{}},
		{desc: "repo, no args", req: &gitalypb.PackObjectsHookRequest{Repository: repo}},
		{desc: "repo, bad args", req: &gitalypb.PackObjectsHookRequest{Repository: repo, Args: []string{"rm", "-rf"}}},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			stream, err := client.PackObjectsHook(ctx)
			require.NoError(t, err)
			require.NoError(t, stream.Send(tc.req))

			_, err = stream.Recv()
			testhelper.RequireGrpcError(t, err, codes.InvalidArgument)
		})
	}
}

func cfgWithCache(t *testing.T) (config.Cfg, *gitalypb.Repository, string) {
	cfg, repo, repoPath := testcfg.BuildWithRepo(t)
	cfg.PackObjectsCache.Enabled = true
	cfg.PackObjectsCache.Dir = testhelper.TempDir(t)
	return cfg, repo, repoPath
}

func TestServer_PackObjectsHook(t *testing.T) {
	ctx, cancel := testhelper.Context()
	defer cancel()

	cfg, repo, repoPath := cfgWithCache(t)

	testCases := []struct {
		desc  string
		stdin string
		args  []string
	}{
		{
			desc:  "clone 1 branch",
			stdin: "3dd08961455abf80ef9115f4afdc1c6f968b503c\n--not\n\n",
			args:  []string{"pack-objects", "--revs", "--thin", "--stdout", "--progress", "--delta-base-offset"},
		},
		{
			desc:  "shallow clone 1 branch",
			stdin: "--shallow 1e292f8fedd741b75372e19097c76d327140c312\n1e292f8fedd741b75372e19097c76d327140c312\n--not\n\n",
			args:  []string{"--shallow-file", "", "pack-objects", "--revs", "--thin", "--stdout", "--shallow", "--progress", "--delta-base-offset", "--include-tag"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			logger, hook := test.NewNullLogger()

			serverSocketPath := runHooksServer(t, cfg, nil, testserver.WithLogger(logger))
			client, conn := newHooksClient(t, serverSocketPath)
			defer conn.Close()

			stream, err := client.PackObjectsHook(ctx)
			require.NoError(t, err)

			require.NoError(t, stream.Send(&gitalypb.PackObjectsHookRequest{
				Repository: repo,
				Args:       tc.args,
			}))

			require.NoError(t, stream.Send(&gitalypb.PackObjectsHookRequest{
				Stdin: []byte(tc.stdin),
			}))
			require.NoError(t, stream.CloseSend())

			var stdout []byte
			for err == nil {
				var resp *gitalypb.PackObjectsHookResponse
				resp, err = stream.Recv()
				stdout = append(stdout, resp.GetStdout()...)
				if stderr := resp.GetStderr(); len(stderr) > 0 {
					t.Log(string(stderr))
				}
			}
			require.Equal(t, io.EOF, err)

			testhelper.MustRunCommand(
				t,
				bytes.NewReader(stdout),
				"git", "-C", repoPath, "index-pack", "--stdin", "--fix-thin",
			)

			for _, msg := range []string{"served bytes", "generated bytes"} {
				t.Run(msg, func(t *testing.T) {
					var entry *logrus.Entry
					for _, e := range hook.AllEntries() {
						if e.Message == msg {
							entry = e
						}
					}

					require.NotNil(t, entry)
					require.NotEmpty(t, entry.Data["cache_key"])
					require.Greater(t, entry.Data["bytes"], int64(0))
				})
			}
		})
	}
}

func TestParsePackObjectsArgs(t *testing.T) {
	testCases := []struct {
		desc string
		args []string
		out  *packObjectsArgs
		err  error
	}{
		{desc: "no args", args: []string{"pack-objects", "--stdout"}, out: &packObjectsArgs{}},
		{desc: "no args shallow", args: []string{"--shallow-file", "", "pack-objects", "--stdout"}, out: &packObjectsArgs{shallowFile: true}},
		{desc: "with args", args: []string{"pack-objects", "--foo", "-x", "--stdout"}, out: &packObjectsArgs{flags: []string{"--foo", "-x"}}},
		{desc: "with args shallow", args: []string{"--shallow-file", "", "pack-objects", "--foo", "--stdout", "-x"}, out: &packObjectsArgs{shallowFile: true, flags: []string{"--foo", "-x"}}},
		{desc: "missing stdout", args: []string{"pack-objects"}, err: errNoStdout},
		{desc: "no pack objects", args: []string{"zpack-objects"}, err: errNoPackObjects},
		{desc: "non empty shallow", args: []string{"--shallow-file", "z", "pack-objects"}, err: errNoPackObjects},
		{desc: "bad global", args: []string{"-c", "foo=bar", "pack-objects"}, err: errNoPackObjects},
		{desc: "non flag arg", args: []string{"pack-objects", "--foo", "x"}, err: errNonFlagArg},
		{desc: "non flag arg shallow", args: []string{"--shallow-file", "", "pack-objects", "--foo", "x"}, err: errNonFlagArg},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			args, err := parsePackObjectsArgs(tc.args)
			require.Equal(t, tc.out, args)
			require.Equal(t, tc.err, err)
		})
	}
}

func TestServer_PackObjectsHook_separateContext(t *testing.T) {
	cfg, repo, repoPath := cfgWithCache(t)

	startRequest := func(ctx context.Context, stream gitalypb.HookService_PackObjectsHookClient) {
		require.NoError(t, stream.Send(&gitalypb.PackObjectsHookRequest{
			Repository: repo,
			Args:       []string{"pack-objects", "--revs", "--thin", "--stdout", "--progress", "--delta-base-offset"},
		}))

		require.NoError(t, stream.Send(&gitalypb.PackObjectsHookRequest{
			Stdin: []byte("3dd08961455abf80ef9115f4afdc1c6f968b503c\n--not\n\n"),
		}))

		require.NoError(t, stream.CloseSend())
	}

	serverSocketPath := runHooksServer(t, cfg, nil)

	client1, conn1 := newHooksClient(t, serverSocketPath)
	defer conn1.Close()

	// Use a timeout to make the first call fail: just canceling the context
	// does not propagate reliably. This timeout must be long enough for the
	// second RPC call to send over its request details, so that it shares
	// the cache entry with the first request.
	const timeout = 100 * time.Millisecond
	ctx1, cancel1 := context.WithTimeout(context.Background(), timeout)
	defer cancel1()

	stream1, err := client1.PackObjectsHook(ctx1)
	require.NoError(t, err)
	startRequest(ctx1, stream1)

	client2, conn2 := newHooksClient(t, serverSocketPath)
	defer conn2.Close()

	ctx2, cancel2 := testhelper.Context()
	defer cancel2()

	stream2, err := client2.PackObjectsHook(ctx2)
	require.NoError(t, err)
	startRequest(ctx2, stream2)

	// If we correctly decoupled the cache from stream1, then cancelation of
	// stream1 should not distrupt stream2.
	time.Sleep(2 * timeout)

	var stdout []byte
	for err == nil {
		var resp *gitalypb.PackObjectsHookResponse
		resp, err = stream2.Recv()
		stdout = append(stdout, resp.GetStdout()...)
	}
	require.Equal(t, io.EOF, err)

	testhelper.MustRunCommand(
		t,
		bytes.NewReader(stdout),
		"git", "-C", repoPath, "index-pack", "--stdin", "--fix-thin",
	)
}

func TestServer_PackObjectsHook_usesCache(t *testing.T) {
	cfg, repo, repoPath := cfgWithCache(t)

	tlc := &streamcache.TestLoggingCache{}
	serverSocketPath := runHooksServer(t, cfg, []serverOption{func(s *server) {
		tlc.Cache = s.packObjectsCache
		s.packObjectsCache = tlc
	}})

	doRequest := func() {
		ctx, cancel := testhelper.Context()
		defer cancel()

		client, conn := newHooksClient(t, serverSocketPath)
		defer conn.Close()

		stream, err := client.PackObjectsHook(ctx)
		require.NoError(t, err)

		require.NoError(t, stream.Send(&gitalypb.PackObjectsHookRequest{
			Repository: repo,
			Args:       []string{"pack-objects", "--revs", "--thin", "--stdout", "--progress", "--delta-base-offset"},
		}))

		require.NoError(t, stream.Send(&gitalypb.PackObjectsHookRequest{
			Stdin: []byte("3dd08961455abf80ef9115f4afdc1c6f968b503c\n--not\n\n"),
		}))
		require.NoError(t, stream.CloseSend())

		var stdout []byte
		for err == nil {
			var resp *gitalypb.PackObjectsHookResponse
			resp, err = stream.Recv()
			stdout = append(stdout, resp.GetStdout()...)
		}
		require.Equal(t, io.EOF, err)

		testhelper.MustRunCommand(
			t,
			bytes.NewReader(stdout),
			"git", "-C", repoPath, "index-pack", "--stdin", "--fix-thin",
		)
	}

	const N = 5
	for i := 0; i < N; i++ {
		doRequest()
	}

	entries := tlc.Entries()
	require.Len(t, entries, N)
	first := entries[0]
	require.NotEmpty(t, first.Key)
	require.True(t, first.Created)
	require.NoError(t, first.Err)

	for i := 1; i < N; i++ {
		require.Equal(t, first.Key, entries[i].Key, "all requests had the same cache key")
		require.False(t, entries[i].Created, "all requests except the first were cache hits")
		require.NoError(t, entries[i].Err)
	}
}
