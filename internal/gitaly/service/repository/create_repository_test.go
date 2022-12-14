//go:build !gitaly_test_sha256

package repository

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git/gittest"
	"gitlab.com/gitlab-org/gitaly/v15/internal/gitaly/config"
	"gitlab.com/gitlab-org/gitaly/v15/internal/gitaly/config/auth"
	"gitlab.com/gitlab-org/gitaly/v15/internal/gitaly/transaction"
	"gitlab.com/gitlab-org/gitaly/v15/internal/helper/text"
	"gitlab.com/gitlab-org/gitaly/v15/internal/metadata"
	"gitlab.com/gitlab-org/gitaly/v15/internal/praefect/praefectutil"
	"gitlab.com/gitlab-org/gitaly/v15/internal/testhelper"
	"gitlab.com/gitlab-org/gitaly/v15/internal/testhelper/testcfg"
	"gitlab.com/gitlab-org/gitaly/v15/internal/testhelper/testserver"
	"gitlab.com/gitlab-org/gitaly/v15/internal/transaction/txinfo"
	"gitlab.com/gitlab-org/gitaly/v15/proto/go/gitalypb"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateRepository_missingAuth(t *testing.T) {
	t.Parallel()
	ctx := testhelper.Context(t)

	cfg, repo, _ := testcfg.BuildWithRepo(t, testcfg.WithBase(config.Cfg{Auth: auth.Config{Token: "some"}}))

	_, serverSocketPath := runRepositoryService(t, cfg, nil)
	client := newRepositoryClient(t, config.Cfg{Auth: auth.Config{Token: ""}}, serverSocketPath)

	_, err := client.CreateRepository(ctx, &gitalypb.CreateRepositoryRequest{Repository: repo})

	testhelper.RequireGrpcCode(t, err, codes.Unauthenticated)
}

func TestCreateRepository_successful(t *testing.T) {
	t.Parallel()
	ctx := testhelper.Context(t)

	cfg, client := setupRepositoryServiceWithoutRepo(t)

	relativePath := "create-repository-test.git"

	repo := &gitalypb.Repository{StorageName: cfg.Storages[0].Name, RelativePath: relativePath}
	req := &gitalypb.CreateRepositoryRequest{Repository: repo}
	_, err := client.CreateRepository(ctx, req)
	require.NoError(t, err)

	repoDir := filepath.Join(cfg.Storages[0].Path, gittest.GetReplicaPath(ctx, t, cfg, repo))

	require.NoError(t, unix.Access(repoDir, unix.R_OK))
	require.NoError(t, unix.Access(repoDir, unix.W_OK))
	require.NoError(t, unix.Access(repoDir, unix.X_OK))

	for _, dir := range []string{repoDir, filepath.Join(repoDir, "refs")} {
		fi, err := os.Stat(dir)
		require.NoError(t, err)
		require.True(t, fi.IsDir(), "%q must be a directory", fi.Name())

		require.NoError(t, unix.Access(dir, unix.R_OK))
		require.NoError(t, unix.Access(dir, unix.W_OK))
		require.NoError(t, unix.Access(dir, unix.X_OK))
	}

	symRef := testhelper.MustReadFile(t, path.Join(repoDir, "HEAD"))
	require.Equal(t, symRef, []byte(fmt.Sprintf("ref: %s\n", git.DefaultRef)))
}

func TestCreateRepository_withDefaultBranch(t *testing.T) {
	cfg, client := setupRepositoryServiceWithoutRepo(t)

	ctx := testhelper.Context(t)

	testCases := []struct {
		name              string
		defaultBranch     string
		expected          string
		expectedErrString string
	}{
		{
			name:          "valid default branch",
			defaultBranch: "develop",
			expected:      "refs/heads/develop",
		},
		{
			name:          "empty branch name",
			defaultBranch: "",
			expected:      "refs/heads/main",
		},
		{
			name:              "invalid branch name",
			defaultBranch:     "./.lock",
			expected:          "refs/heads/main",
			expectedErrString: `creating repository: exit status 128, stderr: "fatal: invalid initial branch name: './.lock'\n"`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := &gitalypb.Repository{StorageName: cfg.Storages[0].Name, RelativePath: gittest.NewRepositoryName(t, true)}

			req := &gitalypb.CreateRepositoryRequest{Repository: repo, DefaultBranch: []byte(tc.defaultBranch)}
			_, err := client.CreateRepository(ctx, req)
			if tc.expectedErrString != "" {
				require.Contains(t, err.Error(), tc.expectedErrString)
			} else {
				require.NoError(t, err)
				repoPath := filepath.Join(cfg.Storages[0].Path, gittest.GetReplicaPath(ctx, t, cfg, repo))
				symRef := text.ChompBytes(gittest.Exec(
					t,
					cfg,
					"-C", repoPath,
					"symbolic-ref", "HEAD"))
				require.Equal(t, tc.expected, symRef)
			}
		})
	}
}

func TestCreateRepository_failure(t *testing.T) {
	t.Parallel()
	ctx := testhelper.Context(t)

	cfg, client := setupRepositoryServiceWithoutRepo(t)

	// This tests what happens if a file already exists at the path of the repository. In order to have
	// the test working with Praefect, we use the relative path to ensure the directory conflicts with
	// the repository as Praefect would create it.
	replicaPath := praefectutil.DeriveReplicaPath(1)
	storagePath := cfg.Storages[0].Path
	fullPath := filepath.Join(storagePath, replicaPath)

	require.NoError(t, os.MkdirAll(filepath.Dir(fullPath), os.ModePerm))
	_, err := os.Create(fullPath)
	require.NoError(t, err)

	_, err = client.CreateRepository(ctx, &gitalypb.CreateRepositoryRequest{
		Repository: &gitalypb.Repository{StorageName: cfg.Storages[0].Name, RelativePath: replicaPath},
	})
	testhelper.RequireGrpcCode(t, err, codes.AlreadyExists)
}

func TestCreateRepository_invalidArguments(t *testing.T) {
	t.Parallel()
	ctx := testhelper.Context(t)

	_, client := setupRepositoryServiceWithoutRepo(t)

	testCases := []struct {
		repo *gitalypb.Repository
		code codes.Code
	}{
		{
			repo: &gitalypb.Repository{StorageName: "does not exist", RelativePath: "foobar.git"},
			code: codes.InvalidArgument,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc.repo), func(t *testing.T) {
			_, err := client.CreateRepository(ctx, &gitalypb.CreateRepositoryRequest{Repository: tc.repo})

			require.Error(t, err)
			testhelper.RequireGrpcCode(t, err, tc.code)
		})
	}
}

func TestCreateRepository_transactional(t *testing.T) {
	t.Parallel()
	ctx := testhelper.Context(t)

	txManager := transaction.NewTrackingManager()

	cfg, client := setupRepositoryServiceWithoutRepo(t, testserver.WithTransactionManager(txManager))

	ctx, err := txinfo.InjectTransaction(ctx, 1, "node", true)
	require.NoError(t, err)
	ctx = metadata.IncomingToOutgoing(ctx)

	t.Run("initial creation without refs", func(t *testing.T) {
		txManager.Reset()

		repo := &gitalypb.Repository{
			StorageName:  cfg.Storages[0].Name,
			RelativePath: "repo.git",
		}
		_, err = client.CreateRepository(ctx, &gitalypb.CreateRepositoryRequest{Repository: repo})
		require.NoError(t, err)

		require.DirExists(t, filepath.Join(cfg.Storages[0].Path, gittest.GetReplicaPath(ctx, t, cfg, repo)))
		require.Equal(t, 2, len(txManager.Votes()), "expected transactional vote")
	})

	t.Run("idempotent creation with preexisting refs", func(t *testing.T) {
		txManager.Reset()

		// The above test creates the second repository on the server. As this test can run with Praefect in front of it,
		// we'll use the next replica path Praefect will assign in order to ensure this repository creation conflicts even
		// with Praefect in front of it.
		repo, _ := gittest.CreateRepository(ctx, t, cfg, gittest.CreateRepositoryConfig{
			RelativePath: praefectutil.DeriveReplicaPath(2),
		})

		_, err = client.CreateRepository(ctx, &gitalypb.CreateRepositoryRequest{
			Repository: repo,
		})
		if testhelper.IsPraefectEnabled() {
			testhelper.ProtoEqual(t, status.Error(codes.AlreadyExists, "route repository creation: reserve repository id: repository already exists"), err)
		} else {
			testhelper.ProtoEqual(t, status.Error(codes.AlreadyExists, "creating repository: repository exists already"), err)
		}
	})
}

func TestCreateRepository_idempotent(t *testing.T) {
	t.Parallel()
	ctx := testhelper.Context(t)

	cfg, client := setupRepositoryServiceWithoutRepo(t)

	repo := &gitalypb.Repository{
		StorageName: cfg.Storages[0].Name,
		// This creates the first repository on the server. As this test can run with Praefect in front of it,
		// we'll use the next replica path Praefect will assign in order to ensure this repository creation
		// conflicts even with Praefect in front of it.
		RelativePath: praefectutil.DeriveReplicaPath(1),
	}
	gittest.CreateRepository(ctx, t, cfg, gittest.CreateRepositoryConfig{
		RelativePath: repo.RelativePath,
	})

	req := &gitalypb.CreateRepositoryRequest{Repository: repo}
	_, err := client.CreateRepository(ctx, req)
	if testhelper.IsPraefectEnabled() {
		testhelper.ProtoEqual(t, status.Error(codes.AlreadyExists, "route repository creation: reserve repository id: repository already exists"), err)
	} else {
		testhelper.ProtoEqual(t, status.Error(codes.AlreadyExists, "creating repository: repository exists already"), err)
	}
}
