// +build static,system_libgit2

package main

import (
	"context"
	"encoding/gob"
	"errors"
	"flag"
	"fmt"
	"io"

	git "github.com/libgit2/git2go/v31"
	"gitlab.com/gitlab-org/gitaly/internal/git2go"
)

type rebaseSubcommand struct{}

func (cmd *rebaseSubcommand) Flags() *flag.FlagSet {
	return flag.NewFlagSet("rebase", flag.ExitOnError)
}

func (cmd *rebaseSubcommand) Run(ctx context.Context, r io.Reader, w io.Writer) error {
	var request git2go.RebaseCommand
	if err := gob.NewDecoder(r).Decode(&request); err != nil {
		return err
	}

	commitID, err := cmd.rebase(ctx, &request)
	return gob.NewEncoder(w).Encode(git2go.Result{
		CommitID: commitID,
		Error:    git2go.SerializableError(err),
	})
}

func (cmd *rebaseSubcommand) verify(ctx context.Context, r *git2go.RebaseCommand) error {
	if r.Repository == "" {
		return errors.New("missing repository")
	}
	if r.Committer.Name == "" {
		return errors.New("missing committer name")
	}
	if r.Committer.Email == "" {
		return errors.New("missing committer email")
	}
	if r.BranchName == "" {
		return errors.New("missing branch name")
	}
	if r.UpstreamBranch == "" {
		return errors.New("missing upstream branch")
	}
	return nil
}

func (cmd *rebaseSubcommand) rebase(ctx context.Context, request *git2go.RebaseCommand) (string, error) {
	if err := cmd.verify(ctx, request); err != nil {
		return "", err
	}

	repo, err := git.OpenRepository(request.Repository)
	if err != nil {
		return "", fmt.Errorf("open repository: %w", err)
	}

	opts, err := git.DefaultRebaseOptions()
	if err != nil {
		return "", fmt.Errorf("get rebase options: %w", err)
	}
	opts.InMemory = 1

	branch, err := repo.AnnotatedCommitFromRevspec(fmt.Sprintf("refs/heads/%s", request.BranchName))
	if err != nil {
		return "", fmt.Errorf("look up branch %q: %w", request.BranchName, err)
	}

	onto, err := repo.AnnotatedCommitFromRevspec(fmt.Sprintf("refs/heads/%s", request.UpstreamBranch))
	if err != nil {
		return "", fmt.Errorf("look up upstream branch %q: %w", request.UpstreamBranch, err)
	}

	rebase, err := repo.InitRebase(branch, onto, onto, &opts)
	if err != nil {
		return "", fmt.Errorf("initiate rebase: %w", err)
	}

	committer := git.Signature(request.Committer)
	var oid *git.Oid
	for {
		op, err := rebase.Next()
		if git.IsErrorCode(err, git.ErrIterOver) {
			break
		} else if err != nil {
			return "", fmt.Errorf("rebase iterate: %w", err)
		}

		commit, err := repo.LookupCommit(op.Id)
		if err != nil {
			return "", fmt.Errorf("lookup commit: %w", err)
		}

		oid = op.Id.Copy()
		err = rebase.Commit(oid, nil, &committer, commit.Message())
		if err != nil {
			return "", fmt.Errorf("commit %q: %w", op.Id.String(), err)
		}
	}

	if oid == nil {
		return branch.Id().String(), nil
	}

	if err = rebase.Finish(); err != nil {
		return "", fmt.Errorf("finish rebase: %w", err)
	}

	return oid.String(), nil
}
