//go:build !gitaly_test_sha256

package middleware

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/gitaly/v15/internal/helper"
	"gitlab.com/gitlab-org/gitaly/v15/internal/praefect/grpc-proxy/proxy"
	"gitlab.com/gitlab-org/gitaly/v15/internal/praefect/mock"
	"gitlab.com/gitlab-org/gitaly/v15/internal/praefect/nodes/tracker"
	"gitlab.com/gitlab-org/gitaly/v15/internal/praefect/protoregistry"
	"gitlab.com/gitlab-org/gitaly/v15/internal/testhelper"
	"gitlab.com/gitlab-org/gitaly/v15/internal/testhelper/testcfg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type simpleService struct {
	mock.UnimplementedSimpleServiceServer
}

func (s *simpleService) RepoAccessorUnary(ctx context.Context, in *mock.RepoRequest) (*emptypb.Empty, error) {
	if in.GetRepo() == nil {
		return nil, helper.ErrInternalf("error")
	}

	return &emptypb.Empty{}, nil
}

func (s *simpleService) RepoMutatorUnary(ctx context.Context, in *mock.RepoRequest) (*emptypb.Empty, error) {
	if in.GetRepo() == nil {
		return nil, helper.ErrInternalf("error")
	}

	return &emptypb.Empty{}, nil
}

func TestStreamInterceptor(t *testing.T) {
	ctx := testhelper.Context(t)

	isInErrorWindow := true
	threshold := 5
	errTracker, err := tracker.NewErrors(ctx, func(_, _ time.Time) bool {
		return isInErrorWindow
	}, uint32(threshold), uint32(threshold))
	require.NoError(t, err)
	nodeName := "node-1"

	internalSrv := grpc.NewServer()

	internalServerSocketPath := testhelper.GetTemporaryGitalySocketFileName(t)
	lis, err := net.Listen("unix", internalServerSocketPath)
	require.NoError(t, err)

	registry, err := protoregistry.NewFromPaths("praefect/mock/mock.proto")
	require.NoError(t, err)

	mock.RegisterSimpleServiceServer(internalSrv, &simpleService{})

	go testhelper.MustServe(t, internalSrv, lis)
	defer internalSrv.Stop()

	srvOptions := []grpc.ServerOption{
		grpc.ForceServerCodec(proxy.NewCodec()),
		grpc.UnknownServiceHandler(proxy.TransparentHandler(func(ctx context.Context,
			fullMethodName string,
			peeker proxy.StreamPeeker,
		) (*proxy.StreamParameters, error) {
			cc, err := grpc.Dial("unix://"+internalServerSocketPath,
				grpc.WithDefaultCallOptions(grpc.ForceCodec(proxy.NewCodec())),
				grpc.WithTransportCredentials(insecure.NewCredentials()),
				grpc.WithStreamInterceptor(StreamErrorHandler(registry, errTracker, nodeName)),
			)
			require.NoError(t, err)
			t.Cleanup(func() { testhelper.MustClose(t, cc) })
			f, err := peeker.Peek()
			require.NoError(t, err)
			return proxy.NewStreamParameters(proxy.Destination{Conn: cc, Ctx: ctx, Msg: f}, nil, func() error { return nil }, nil), nil
		})),
	}

	praefectSocket := testhelper.GetTemporaryGitalySocketFileName(t)
	praefectLis, err := net.Listen("unix", praefectSocket)
	require.NoError(t, err)

	praefectSrv := grpc.NewServer(srvOptions...)
	defer praefectSrv.Stop()
	go testhelper.MustServe(t, praefectSrv, praefectLis)

	praefectCC, err := grpc.Dial("unix://"+praefectSocket, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer testhelper.MustClose(t, praefectCC)
	require.NoError(t, err)

	simpleClient := mock.NewSimpleServiceClient(praefectCC)

	_, repo, _ := testcfg.BuildWithRepo(t)

	for i := 0; i < threshold; i++ {
		_, err = simpleClient.RepoAccessorUnary(ctx, &mock.RepoRequest{
			Repo: repo,
		})
		require.NoError(t, err)
		_, err = simpleClient.RepoMutatorUnary(ctx, &mock.RepoRequest{
			Repo: repo,
		})
		require.NoError(t, err)
	}

	assert.False(t, errTracker.WriteThresholdReached(nodeName))
	assert.False(t, errTracker.ReadThresholdReached(nodeName))

	for i := 0; i < threshold; i++ {
		_, err = simpleClient.RepoAccessorUnary(ctx, &mock.RepoRequest{
			Repo: nil,
		})
		require.Error(t, err)
		_, err = simpleClient.RepoMutatorUnary(ctx, &mock.RepoRequest{
			Repo: nil,
		})
		require.Error(t, err)
	}

	assert.True(t, errTracker.WriteThresholdReached(nodeName))
	assert.True(t, errTracker.ReadThresholdReached(nodeName))

	isInErrorWindow = false

	for i := 0; i < threshold; i++ {
		_, err = simpleClient.RepoAccessorUnary(ctx, &mock.RepoRequest{
			Repo: repo,
		})
		require.NoError(t, err)
		_, err = simpleClient.RepoMutatorUnary(ctx, &mock.RepoRequest{
			Repo: repo,
		})
		require.NoError(t, err)
	}

	assert.False(t, errTracker.WriteThresholdReached(nodeName))
	assert.False(t, errTracker.ReadThresholdReached(nodeName))
}
