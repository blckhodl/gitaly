//go:build !gitaly_test_sha256

package chunk

import (
	"io"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	test "gitlab.com/gitlab-org/gitaly/v15/internal/helper/chunk/testdata"
	"gitlab.com/gitlab-org/gitaly/v15/internal/testhelper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestMain(m *testing.M) {
	testhelper.Run(m)
}

type testSender struct {
	stream test.Test_StreamOutputServer
	output [][]byte
}

func (ts *testSender) Reset() { ts.output = nil }
func (ts *testSender) Append(m proto.Message) {
	ts.output = append(ts.output, m.(*wrapperspb.BytesValue).Value)
}

func (ts *testSender) Send() error {
	return ts.stream.Send(&test.StreamOutputResponse{
		Msg: ts.output,
	})
}

func TestChunker(t *testing.T) {
	s := &server{}
	srv, serverSocketPath := runServer(t, s)
	defer srv.Stop()

	client, conn := newClient(t, serverSocketPath)
	defer conn.Close()
	ctx := testhelper.Context(t)

	stream, err := client.StreamOutput(ctx, &test.StreamOutputRequest{BytesToReturn: 3.5 * maxMessageSize})
	require.NoError(t, err)

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		require.Less(t, proto.Size(resp), maxMessageSize)
	}
}

type server struct {
	test.UnimplementedTestServer
}

func (s *server) StreamOutput(req *test.StreamOutputRequest, srv test.Test_StreamOutputServer) error {
	const kilobyte = 1024

	c := New(&testSender{stream: srv})
	for numBytes := 0; numBytes < int(req.GetBytesToReturn()); numBytes += kilobyte {
		if err := c.Send(&wrapperspb.BytesValue{Value: make([]byte, kilobyte)}); err != nil {
			return err
		}
	}

	if err := c.Flush(); err != nil {
		return err
	}
	return nil
}

func runServer(t *testing.T, s *server, opt ...grpc.ServerOption) (*grpc.Server, string) {
	serverSocketPath := testhelper.GetTemporaryGitalySocketFileName(t)
	grpcServer := grpc.NewServer(opt...)
	test.RegisterTestServer(grpcServer, s)

	lis, err := net.Listen("unix", serverSocketPath)
	require.NoError(t, err)

	go testhelper.MustServe(t, grpcServer, lis)

	return grpcServer, "unix://" + serverSocketPath
}

func newClient(t *testing.T, serverSocketPath string) (test.TestClient, *grpc.ClientConn) {
	connOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(serverSocketPath, connOpts...)
	if err != nil {
		t.Fatal(err)
	}

	return test.NewTestClient(conn), conn
}
