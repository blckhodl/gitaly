package objectpool

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"gitlab.com/gitlab-org/gitaly/v15/internal/git"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git/localrepo"
)

// clone a repository to a pool, without setting the alternates, is not the
// responsibility of this function.
func (o *ObjectPool) clone(ctx context.Context, repo *localrepo.Repo) error {
	repoPath, err := repo.Path()
	if err != nil {
		return err
	}

	var stderr bytes.Buffer
	cmd, err := o.gitCmdFactory.NewWithoutRepo(ctx,
		git.SubCmd{
			Name: "clone",
			Flags: []git.Option{
				git.Flag{Name: "--quiet"},
				git.Flag{Name: "--bare"},
				git.Flag{Name: "--local"},
			},
			Args: []string{repoPath, o.FullPath()},
		},
		git.WithRefTxHook(repo),
		git.WithStderr(&stderr),
	)
	if err != nil {
		return fmt.Errorf("spawning clone: %w", err)
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("cloning to pool: %w, stderr: %q", err, stderr.String())
	}

	return nil
}

func (o *ObjectPool) removeHooksDir() error {
	return os.RemoveAll(filepath.Join(o.FullPath(), "hooks"))
}
