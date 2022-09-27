package repository

//nolint:depguard
import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	log "github.com/sirupsen/logrus"
	gitalyerrors "gitlab.com/gitlab-org/gitaly/v15/internal/errors"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git/catfile"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git/housekeeping"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git/localrepo"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git/stats"
	"gitlab.com/gitlab-org/gitaly/v15/internal/helper"
	"gitlab.com/gitlab-org/gitaly/v15/proto/go/gitalypb"
)

func (s *server) GarbageCollect(ctx context.Context, in *gitalypb.GarbageCollectRequest) (*gitalypb.GarbageCollectResponse, error) {
	ctxlogger := ctxlogrus.Extract(ctx)
	ctxlogger.WithFields(log.Fields{
		"WriteBitmaps": in.GetCreateBitmap(),
	}).Debug("GarbageCollect")

	if in.GetRepository() == nil {
		return nil, helper.ErrInvalidArgument(gitalyerrors.ErrEmptyRepository)
	}

	repo := s.localrepo(in.GetRepository())

	if err := housekeeping.CleanupWorktrees(ctx, repo); err != nil {
		return nil, err
	}

	if err := s.cleanupKeepArounds(ctx, repo); err != nil {
		return nil, helper.ErrInternalf("cleanup keep-arounds: %w", err)
	}

	// Perform housekeeping to cleanup stale lockfiles that may block GC
	if err := s.housekeepingManager.CleanStaleData(ctx, repo); err != nil {
		ctxlogger.WithError(err).Warn("Pre gc housekeeping failed")
	}

	if err := s.gc(ctx, in); err != nil {
		return nil, helper.ErrInternalf("garbage collect: %w", err)
	}

	if err := housekeeping.WriteCommitGraph(ctx, repo, housekeeping.WriteCommitGraphConfig{
		ReplaceChain: true,
	}); err != nil {
		return nil, err
	}

	stats.LogObjectsInfo(ctx, repo)

	return &gitalypb.GarbageCollectResponse{}, nil
}

func (s *server) gc(ctx context.Context, in *gitalypb.GarbageCollectRequest) error {
	config := append(housekeeping.GetRepackGitConfig(ctx, in.GetRepository(), in.CreateBitmap), git.ConfigPair{Key: "gc.writeCommitGraph", Value: "false"})

	var flags []git.Option
	if in.Prune {
		flags = append(flags, git.Flag{Name: "--prune=30.minutes.ago"})
	}

	cmd, err := s.gitCmdFactory.New(ctx, in.GetRepository(),
		git.SubCmd{Name: "gc", Flags: flags},
		git.WithConfig(config...),
	)
	if err != nil {
		if git.IsInvalidArgErr(err) {
			return helper.ErrInvalidArgumentf("gitCommand: %w", err)
		}

		return helper.ErrInternal(fmt.Errorf("gitCommand: %w", err))
	}

	if err := cmd.Wait(); err != nil {
		return helper.ErrInternal(fmt.Errorf("cmd wait: %w", err))
	}

	return nil
}

func (s *server) cleanupKeepArounds(ctx context.Context, repo *localrepo.Repo) error {
	repoPath, err := repo.Path()
	if err != nil {
		return nil
	}

	objectInfoReader, cancel, err := s.catfileCache.ObjectInfoReader(ctx, repo)
	if err != nil {
		return nil
	}
	defer cancel()

	keepAroundsPrefix := "refs/keep-around"
	keepAroundsPath := filepath.Join(repoPath, keepAroundsPrefix)

	refInfos, err := ioutil.ReadDir(keepAroundsPath)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("list refs/keep-around: %w", err)
	}

	for _, info := range refInfos {
		if info.IsDir() {
			continue
		}

		refName := fmt.Sprintf("%s/%s", keepAroundsPrefix, info.Name())
		path := filepath.Join(repoPath, keepAroundsPrefix, info.Name())

		if err = checkRef(ctx, objectInfoReader, refName, info); err == nil {
			continue
		}

		if err := s.fixRef(ctx, repo, objectInfoReader, path, refName, info.Name()); err != nil {
			return fmt.Errorf("fix ref: %w", err)
		}
	}

	return nil
}

func checkRef(ctx context.Context, objectInfoReader catfile.ObjectInfoReader, refName string, info os.FileInfo) error {
	if info.Size() == 0 {
		return errors.New("checkRef: Ref file is empty")
	}

	_, err := objectInfoReader.Info(ctx, git.Revision(refName))
	return err
}

func (s *server) fixRef(ctx context.Context, repo *localrepo.Repo, objectInfoReader catfile.ObjectInfoReader, refPath string, name string, sha string) error {
	// So the ref is broken, let's get rid of it
	if err := os.RemoveAll(refPath); err != nil {
		return err
	}

	// If the sha is not in the the repository, we can't fix it
	if _, err := objectInfoReader.Info(ctx, git.Revision(sha)); err != nil {
		return nil
	}

	// The name is a valid sha, recreate the ref
	return repo.ExecAndWait(ctx, git.SubCmd{
		Name: "update-ref",
		Args: []string{name, sha},
	}, git.WithRefTxHook(repo))
}
