//go:build !gitaly_test_sha256

package main

import (
	"fmt"
	"net"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/gitaly/v15/internal/testhelper"
	"gitlab.com/gitlab-org/gitaly/v15/proto/go/gitalypb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// svcRegistrar is a function that registers a gRPC service with a server
// instance
type svcRegistrar func(*grpc.Server)

func registerHealthService(srv *grpc.Server) {
	healthSrvr := health.NewServer()
	grpc_health_v1.RegisterHealthServer(srv, healthSrvr)
	healthSrvr.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
}

func registerServerService(impl gitalypb.ServerServiceServer) svcRegistrar {
	return func(srv *grpc.Server) {
		gitalypb.RegisterServerServiceServer(srv, impl)
	}
}

func listenAndServe(tb testing.TB, svcs []svcRegistrar) (net.Listener, testhelper.Cleanup) {
	tb.Helper()

	tmp := testhelper.TempDir(tb)

	ln, err := net.Listen("unix", filepath.Join(tmp, "gitaly"))
	require.NoError(tb, err)

	srv := grpc.NewServer()

	for _, s := range svcs {
		s(srv)
	}

	go func() { require.NoError(tb, srv.Serve(ln)) }()
	ctx := testhelper.Context(tb)

	// verify the service is up
	addr := fmt.Sprintf("%s://%s", ln.Addr().Network(), ln.Addr())
	cc, err := grpc.DialContext(ctx, addr, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(tb, err)
	require.NoError(tb, cc.Close())

	return ln, func() {
		srv.Stop()
	}
}
