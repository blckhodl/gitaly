package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git/gittest"
	"gitlab.com/gitlab-org/gitaly/v15/internal/gitaly/config"
	"gitlab.com/gitlab-org/gitaly/v15/internal/gitaly/service/setup"
	"gitlab.com/gitlab-org/gitaly/v15/internal/testhelper"
	"gitlab.com/gitlab-org/gitaly/v15/internal/testhelper/testcfg"
	"gitlab.com/gitlab-org/gitaly/v15/internal/testhelper/testserver"
	"gitlab.com/gitlab-org/gitaly/v15/proto/go/gitalypb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestConfigLocator_GetRepoPath(t *testing.T) {
	t.Parallel()
	const storageName = "exists"
	ctx := testhelper.Context(t)
	cfg := testcfg.Build(t, testcfg.WithStorages(storageName, "removed"))
	cfg.SocketPath = testserver.RunGitalyServer(t, cfg, nil, setup.RegisterAll)
	repo, repoPath := gittest.CreateRepository(ctx, t, cfg)
	locator := config.NewLocator(cfg)

	t.Run("proper repository path", func(t *testing.T) {
		path, err := locator.GetRepoPath(repo)
		require.NoError(t, err)
		require.Equal(t, repoPath, path)
	})

	// The storage name still present in the storages list, but not on the disk.
	require.NoError(t, os.RemoveAll(cfg.Storages[1].Path))

	// The repository path exists on the disk, but it is not a git repository.
	const notRepositoryFolder = "not-a-git-repo"
	require.NoError(t, os.MkdirAll(filepath.Join(cfg.Storages[0].Path, notRepositoryFolder), 0o755))

	for _, tc := range []struct {
		desc   string
		repo   *gitalypb.Repository
		expErr error
	}{
		{
			desc:   "storage is empty",
			repo:   &gitalypb.Repository{RelativePath: repo.RelativePath},
			expErr: status.Error(codes.InvalidArgument, `GetStorageByName: no such storage: ""`),
		},
		{
			desc:   "unknown storage",
			repo:   &gitalypb.Repository{StorageName: "invalid", RelativePath: repo.RelativePath},
			expErr: status.Error(codes.InvalidArgument, `GetStorageByName: no such storage: "invalid"`),
		},
		{
			desc:   "storage doesn't exist on disk",
			repo:   &gitalypb.Repository{StorageName: cfg.Storages[1].Name, RelativePath: repo.RelativePath},
			expErr: status.Errorf(codes.NotFound, `GetPath: does not exist: stat %s: no such file or directory`, cfg.Storages[1].Path),
		},
		{
			desc:   "relative path is empty",
			repo:   &gitalypb.Repository{StorageName: storageName, RelativePath: ""},
			expErr: status.Error(codes.InvalidArgument, `GetPath: relative path missing from storage_name:"exists"`),
		},
		{
			desc:   "unknown relative path",
			repo:   &gitalypb.Repository{StorageName: storageName, RelativePath: "invalid"},
			expErr: status.Errorf(codes.NotFound, `GetRepoPath: not a git repository: %q`, filepath.Join(cfg.Storages[0].Path, "invalid")),
		},
		{
			desc:   "path exists but not a git repository",
			repo:   &gitalypb.Repository{StorageName: storageName, RelativePath: notRepositoryFolder},
			expErr: status.Errorf(codes.NotFound, `GetRepoPath: not a git repository: %q`, filepath.Join(cfg.Storages[0].Path, notRepositoryFolder)),
		},
		{
			desc:   "relative path escapes parent folder",
			repo:   &gitalypb.Repository{StorageName: storageName, RelativePath: "../.."},
			expErr: status.Error(codes.InvalidArgument, `GetRepoPath: relative path escapes root directory`),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			locator := config.NewLocator(cfg)
			_, err := locator.GetRepoPath(tc.repo)
			require.Equal(t, tc.expErr, err)
		})
	}
}

func TestConfigLocator_GetPath(t *testing.T) {
	t.Parallel()
	const storageName = "exists"
	ctx := testhelper.Context(t)
	cfg := testcfg.Build(t, testcfg.WithStorages(storageName, "removed"))
	cfg.SocketPath = testserver.RunGitalyServer(t, cfg, nil, setup.RegisterAll)
	repo, repoPath := gittest.CreateRepository(ctx, t, cfg)

	// The storage name still present in the storages list, but not on the disk.
	require.NoError(t, os.RemoveAll(cfg.Storages[1].Path))

	// The repository path exists on the disk, but it is not a git repository.
	const notRepositoryFolder = "not-a-git-repo"
	require.NoError(t, os.MkdirAll(filepath.Join(cfg.Storages[0].Path, notRepositoryFolder), 0o755))

	for _, tc := range []struct {
		desc string
		repo *gitalypb.Repository
		path string
	}{
		{
			desc: "proper repository path",
			repo: repo,
			path: repoPath,
		},
		{
			desc: "path doesn't exist",
			repo: &gitalypb.Repository{StorageName: storageName, RelativePath: "doesnt/exist"},
			path: filepath.Join(cfg.Storages[0].Path, "doesnt/exist"),
		},
		{
			desc: "path exists but not a git repository",
			repo: &gitalypb.Repository{StorageName: storageName, RelativePath: notRepositoryFolder},
			path: filepath.Join(cfg.Storages[0].Path, notRepositoryFolder),
		},
		{
			desc: "relative path includes travels",
			repo: &gitalypb.Repository{StorageName: storageName, RelativePath: "some/../other"},
			path: filepath.Join(cfg.Storages[0].Path, "other"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			locator := config.NewLocator(cfg)
			path, err := locator.GetPath(tc.repo)
			require.NoError(t, err)
			require.Equal(t, tc.path, path)
		})
	}

	for _, tc := range []struct {
		desc   string
		repo   *gitalypb.Repository
		expErr error
	}{
		{
			desc:   "storage is empty",
			repo:   &gitalypb.Repository{RelativePath: repo.RelativePath},
			expErr: status.Error(codes.InvalidArgument, `GetStorageByName: no such storage: ""`),
		},
		{
			desc:   "unknown storage",
			repo:   &gitalypb.Repository{StorageName: "invalid", RelativePath: repo.RelativePath},
			expErr: status.Error(codes.InvalidArgument, `GetStorageByName: no such storage: "invalid"`),
		},
		{
			desc:   "storage doesn't exist on disk",
			repo:   &gitalypb.Repository{StorageName: cfg.Storages[1].Name, RelativePath: repo.RelativePath},
			expErr: status.Errorf(codes.NotFound, `GetPath: does not exist: stat %s: no such file or directory`, cfg.Storages[1].Path),
		},
		{
			desc:   "relative path is empty",
			repo:   &gitalypb.Repository{StorageName: storageName, RelativePath: ""},
			expErr: status.Error(codes.InvalidArgument, `GetPath: relative path missing from storage_name:"exists"`),
		},
		{
			desc:   "relative path escapes parent folder",
			repo:   &gitalypb.Repository{StorageName: storageName, RelativePath: "../.."},
			expErr: status.Error(codes.InvalidArgument, `GetRepoPath: relative path escapes root directory`),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			locator := config.NewLocator(cfg)
			_, err := locator.GetPath(tc.repo)
			require.Equal(t, tc.expErr, err)
		})
	}
}

func TestConfigLocator_CacheDir(t *testing.T) {
	t.Parallel()
	const storageName = "exists"
	cfg := testcfg.Build(t, testcfg.WithStorages(storageName, "removed"))
	locator := config.NewLocator(cfg)

	t.Run("storage exists", func(t *testing.T) {
		path, err := locator.CacheDir(storageName)
		require.NoError(t, err)
		require.Equal(t, path, filepath.Join(cfg.Storages[0].Path, "+gitaly/cache"))
	})

	t.Run("storage doesn't exist on disk", func(t *testing.T) {
		require.NoError(t, os.RemoveAll(cfg.Storages[1].Path))
		path, err := locator.CacheDir(cfg.Storages[1].Name)
		require.NoError(t, err)
		require.Equal(t, filepath.Join(cfg.Storages[1].Path, "+gitaly/cache"), path)
	})

	t.Run("unknown storage", func(t *testing.T) {
		_, err := locator.CacheDir("unknown")
		require.Equal(t, status.Error(codes.InvalidArgument, `cache dir: no such storage: "unknown"`), err)
	})
}

func TestConfigLocator_StateDir(t *testing.T) {
	t.Parallel()
	const storageName = "exists"
	cfg := testcfg.Build(t, testcfg.WithStorages(storageName, "removed"))
	locator := config.NewLocator(cfg)

	t.Run("storage exists", func(t *testing.T) {
		path, err := locator.StateDir(storageName)
		require.NoError(t, err)
		require.Equal(t, path, filepath.Join(cfg.Storages[0].Path, "+gitaly/state"))
	})

	t.Run("storage doesn't exist on disk", func(t *testing.T) {
		require.NoError(t, os.RemoveAll(cfg.Storages[1].Path))
		path, err := locator.StateDir(cfg.Storages[1].Name)
		require.NoError(t, err)
		require.Equal(t, filepath.Join(cfg.Storages[1].Path, "+gitaly/state"), path)
	})

	t.Run("unknown storage", func(t *testing.T) {
		_, err := locator.StateDir("unknown")
		require.Equal(t, status.Error(codes.InvalidArgument, `state dir: no such storage: "unknown"`), err)
	})
}

func TestConfigLocator_TempDir(t *testing.T) {
	t.Parallel()
	const storageName = "exists"
	cfg := testcfg.Build(t, testcfg.WithStorages(storageName, "removed"))
	locator := config.NewLocator(cfg)

	t.Run("storage exists", func(t *testing.T) {
		path, err := locator.TempDir(storageName)
		require.NoError(t, err)
		require.Equal(t, path, filepath.Join(cfg.Storages[0].Path, "+gitaly/tmp"))
	})

	t.Run("storage doesn't exist on disk", func(t *testing.T) {
		require.NoError(t, os.RemoveAll(cfg.Storages[1].Path))
		path, err := locator.TempDir(cfg.Storages[1].Name)
		require.NoError(t, err)
		require.Equal(t, filepath.Join(cfg.Storages[1].Path, "+gitaly/tmp"), path)
	})

	t.Run("unknown storage", func(t *testing.T) {
		_, err := locator.TempDir("unknown")
		require.Equal(t, status.Error(codes.InvalidArgument, `tmp dir: no such storage: "unknown"`), err)
	})
}
