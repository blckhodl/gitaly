//go:build !gitaly_test_sha256

package commit

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/gitaly/v15/internal/testhelper"
	"gitlab.com/gitlab-org/gitaly/v15/proto/go/gitalypb"
	"gitlab.com/gitlab-org/gitaly/v15/streamio"
	"google.golang.org/grpc/codes"
)

func TestSuccessfulRawBlameRequest(t *testing.T) {
	t.Parallel()

	ctx := testhelper.Context(t)
	_, repo, _, client := setupCommitServiceWithRepo(ctx, t)

	testCases := []struct {
		revision, path, data, blameRange []byte
	}{
		{
			revision:   []byte("e63f41fe459e62e1228fcef60d7189127aeba95a"),
			path:       []byte("files/ruby/popen.rb"),
			data:       testhelper.MustReadFile(t, "testdata/files-ruby-popen-e63f41f-blame.txt"),
			blameRange: []byte{},
		},
		{
			revision:   []byte("e63f41fe459e62e1228fcef60d7189127aeba95a"),
			path:       []byte("files/ruby/popen.rb"),
			data:       testhelper.MustReadFile(t, "testdata/files-ruby-popen-e63f41f-blame.txt")[0:956],
			blameRange: []byte("1,5"),
		},
		{
			revision:   []byte("e63f41fe459e62e1228fcef60d7189127aeba95a"),
			path:       []byte("files/ruby/../ruby/popen.rb"),
			data:       testhelper.MustReadFile(t, "testdata/files-ruby-popen-e63f41f-blame.txt"),
			blameRange: []byte{},
		},
		{
			revision:   []byte("93dcf076a236c837dd47d61f86d95a6b3d71b586"),
			path:       []byte("gitaly/empty-file"),
			data:       []byte{},
			blameRange: []byte{},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("test case: revision=%q path=%q", testCase.revision, testCase.path), func(t *testing.T) {
			request := &gitalypb.RawBlameRequest{
				Repository: repo,
				Revision:   testCase.revision,
				Path:       testCase.path,
				Range:      testCase.blameRange,
			}
			c, err := client.RawBlame(ctx, request)
			require.NoError(t, err)

			sr := streamio.NewReader(func() ([]byte, error) {
				response, err := c.Recv()
				return response.GetData(), err
			})

			blame, err := io.ReadAll(sr)
			require.NoError(t, err)

			require.Equal(t, testCase.data, blame, "blame data mismatched")
		})
	}
}

func TestFailedRawBlameRequest(t *testing.T) {
	t.Parallel()

	ctx := testhelper.Context(t)
	_, repo, _, client := setupCommitServiceWithRepo(ctx, t)

	invalidRepo := &gitalypb.Repository{StorageName: "fake", RelativePath: "path"}

	testCases := []struct {
		description    string
		repo           *gitalypb.Repository
		revision, path []byte
		blameRange     []byte
		code           codes.Code
	}{
		{
			description: "Invalid repo",
			repo:        invalidRepo,
			revision:    []byte("master"),
			path:        []byte("a/b/c"),
			blameRange:  []byte{},
			code:        codes.InvalidArgument,
		},
		{
			description: "Empty revision",
			repo:        repo,
			revision:    []byte(""),
			path:        []byte("a/b/c"),
			blameRange:  []byte{},
			code:        codes.InvalidArgument,
		},
		{
			description: "Empty path",
			repo:        repo,
			revision:    []byte("abcdef"),
			path:        []byte(""),
			blameRange:  []byte{},
			code:        codes.InvalidArgument,
		},
		{
			description: "Invalid revision",
			repo:        repo,
			revision:    []byte("--output=/meow"),
			path:        []byte("a/b/c"),
			blameRange:  []byte{},
			code:        codes.InvalidArgument,
		},
		{
			description: "Invalid range",
			repo:        repo,
			revision:    []byte("abcdef"),
			path:        []byte("a/b/c"),
			blameRange:  []byte("foo"),
			code:        codes.InvalidArgument,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			request := gitalypb.RawBlameRequest{
				Repository: testCase.repo,
				Revision:   testCase.revision,
				Path:       testCase.path,
				Range:      testCase.blameRange,
			}
			c, err := client.RawBlame(ctx, &request)
			require.NoError(t, err)

			testhelper.RequireGrpcCode(t, drainRawBlameResponse(c), testCase.code)
		})
	}
}

func drainRawBlameResponse(c gitalypb.CommitService_RawBlameClient) error {
	var err error
	for err == nil {
		_, err = c.Recv()
	}
	return err
}
