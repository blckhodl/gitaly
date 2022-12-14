package objectpool

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gitlab.com/gitlab-org/gitaly/v15/internal/git"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git/catfile"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git/housekeeping"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git/localrepo"
	"gitlab.com/gitlab-org/gitaly/v15/internal/gitaly/storage"
	"gitlab.com/gitlab-org/gitaly/v15/internal/gitaly/transaction"
)

type errString string

func (err errString) Error() string { return string(err) }

// ErrInvalidPoolDir is returned when the object pool relative path is malformed.
const ErrInvalidPoolDir errString = "invalid object pool directory"

// ObjectPool are a way to de-dupe objects between repositories, where the objects
// live in a pool in a distinct repository which is used as an alternate object
// store for other repositories.
type ObjectPool struct {
	Repo *localrepo.Repo

	gitCmdFactory       git.CommandFactory
	txManager           transaction.Manager
	housekeepingManager housekeeping.Manager

	storageName  string
	storagePath  string
	relativePath string
}

// NewObjectPool will initialize the object with the required data on the storage
// shard. Relative path is validated to match the expected naming and directory
// structure. If the shard cannot be found, this function returns an error.
func NewObjectPool(
	locator storage.Locator,
	gitCmdFactory git.CommandFactory,
	catfileCache catfile.Cache,
	txManager transaction.Manager,
	housekeepingManager housekeeping.Manager,
	storageName,
	relativePath string,
) (*ObjectPool, error) {
	storagePath, err := locator.GetStorageByName(storageName)
	if err != nil {
		return nil, err
	}

	pool := &ObjectPool{
		gitCmdFactory:       gitCmdFactory,
		txManager:           txManager,
		housekeepingManager: housekeepingManager,
		storageName:         storageName,
		storagePath:         storagePath,
		relativePath:        relativePath,
	}
	pool.Repo = localrepo.New(locator, gitCmdFactory, catfileCache, pool)

	if !housekeeping.IsPoolRepository(pool) {
		return nil, ErrInvalidPoolDir
	}

	return pool, nil
}

// GetGitAlternateObjectDirectories for object pools are empty, given pools are
// never a member of another pool, nor do they share Alternate objects with other
// repositories which the pool doesn't contain itself
func (o *ObjectPool) GetGitAlternateObjectDirectories() []string {
	return []string{}
}

// GetGitObjectDirectory satisfies the repository.GitRepo interface, but is not
// used for ObjectPools
func (o *ObjectPool) GetGitObjectDirectory() string {
	return ""
}

// Exists will return true if the pool path exists and is a directory
func (o *ObjectPool) Exists() bool {
	fi, err := os.Stat(o.FullPath())
	if os.IsNotExist(err) || err != nil {
		return false
	}

	return fi.IsDir()
}

// IsValid checks if a repository exists, and if its valid.
func (o *ObjectPool) IsValid() bool {
	if !o.Exists() {
		return false
	}

	return storage.IsGitDirectory(o.FullPath())
}

// Create will create a pool for a repository and pull the required data to this
// pool. `repo` that is passed also joins the repository.
func (o *ObjectPool) Create(ctx context.Context, repo *localrepo.Repo) (err error) {
	if err := o.clone(ctx, repo); err != nil {
		return fmt.Errorf("clone: %v", err)
	}

	if err := o.removeHooksDir(); err != nil {
		return fmt.Errorf("remove hooks: %v", err)
	}

	return nil
}

// Remove will remove the pool, and all its contents without preparing and/or
// updating the repositories depending on this object pool
// Subdirectories will remain to exist, and will never be cleaned up, even when
// these are empty.
func (o *ObjectPool) Remove(ctx context.Context) (err error) {
	return os.RemoveAll(o.FullPath())
}

// Init will initialize an empty pool repository
// if one already exists, it will do nothing
func (o *ObjectPool) Init(ctx context.Context) (err error) {
	targetDir := o.FullPath()

	if storage.IsGitDirectory(targetDir) {
		return nil
	}

	cmd, err := o.gitCmdFactory.NewWithoutRepo(ctx,
		git.SubCmd{
			Name: "init",
			Flags: []git.Option{
				git.Flag{Name: "--bare"},
			},
			Args: []string{targetDir},
		},
	)
	if err != nil {
		return err
	}

	return cmd.Wait()
}

// FromRepo returns an instance of ObjectPool that the repository points to
func FromRepo(
	locator storage.Locator,
	gitCmdFactory git.CommandFactory,
	catfileCache catfile.Cache,
	txManager transaction.Manager,
	housekeepingManager housekeeping.Manager,
	repo *localrepo.Repo,
) (*ObjectPool, error) {
	dir, err := getAlternateObjectDir(repo)
	if err != nil {
		return nil, err
	}

	if dir == "" {
		return nil, nil
	}

	repoPath, err := repo.Path()
	if err != nil {
		return nil, err
	}

	altPathRelativeToStorage, err := objectPathRelativeToStorage(locator, repo.GetStorageName(), dir, repoPath)
	if err != nil {
		return nil, err
	}

	return NewObjectPool(locator, gitCmdFactory, catfileCache, txManager, housekeepingManager, repo.GetStorageName(), filepath.Dir(altPathRelativeToStorage))
}

var (
	// ErrInvalidPoolRepository indicates the directory the alternates file points to is not a valid git repository
	ErrInvalidPoolRepository = errors.New("object pool is not a valid git repository")

	// ErrAlternateObjectDirNotExist indicates a repository does not have an alternates file
	ErrAlternateObjectDirNotExist = errors.New("no alternates directory exists")
)

// getAlternateObjectDir returns the entry in the objects/info/attributes file if it exists
// it will only return the first line of the file if there are multiple lines.
func getAlternateObjectDir(repo *localrepo.Repo) (string, error) {
	altPath, err := repo.InfoAlternatesPath()
	if err != nil {
		return "", err
	}

	if _, err = os.Stat(altPath); err != nil {
		if os.IsNotExist(err) {
			return "", ErrAlternateObjectDirNotExist
		}
		return "", err
	}

	altFile, err := os.Open(altPath)
	if err != nil {
		return "", err
	}
	defer altFile.Close()

	r := bufio.NewReader(altFile)
	b, err := r.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("reading alternates file: %v", err)
	}

	if err == nil {
		b = b[:len(b)-1]
	}

	if bytes.HasPrefix(b, []byte("#")) {
		return "", ErrAlternateObjectDirNotExist
	}

	return string(b), nil
}

// objectPathRelativeToStorage takes an object path that's relative to a repository's object directory
// and returns the path relative to the storage path of the repository.
func objectPathRelativeToStorage(locator storage.Locator, storageName, path, repoPath string) (string, error) {
	storagePath, err := locator.GetStorageByName(storageName)
	if err != nil {
		return "", err
	}
	objectDirPath := filepath.Join(repoPath, "objects")

	poolObjectDirFullPath := filepath.Join(objectDirPath, path)

	if !storage.IsGitDirectory(filepath.Dir(poolObjectDirFullPath)) {
		return "", ErrInvalidPoolRepository
	}

	poolRelPath, err := filepath.Rel(storagePath, poolObjectDirFullPath)
	if err != nil {
		return "", err
	}

	return poolRelPath, nil
}
