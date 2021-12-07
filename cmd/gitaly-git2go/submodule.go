//go:build static && system_libgit2
// +build static,system_libgit2

package main

import (
	"context"
	"encoding/gob"
	"flag"
	"fmt"
	"io"
	"time"

	git "github.com/libgit2/git2go/v32"
	"gitlab.com/gitlab-org/gitaly/v14/cmd/gitaly-git2go/git2goutil"
	"gitlab.com/gitlab-org/gitaly/v14/internal/git2go"
)

type submoduleSubcommand struct {
	request string
}

func (cmd *submoduleSubcommand) Flags() *flag.FlagSet {
	flags := flag.NewFlagSet("submodule", flag.ExitOnError)
	flags.StringVar(&cmd.request, "request", "", "git2go.SubmoduleCommand")
	return flags
}

func (cmd *submoduleSubcommand) Run(_ context.Context, r io.Reader, w io.Writer) error {
	var request git2go.SubmoduleCommand
	if cmd.request == "" {
		if err := gob.NewDecoder(r).Decode(&request); err != nil {
			return fmt.Errorf("deserializing submodule command request: %w", err)
		}
	} else {
		var err error
		request, err = git2go.SubmoduleCommandFromSerialized(cmd.request)
		if err != nil {
			return fmt.Errorf("deserializing submodule command request: %w", err)
		}
	}

	res, err := cmd.run(request)
	if err != nil {
		return err
	}

	if cmd.request == "" {
		return gob.NewEncoder(w).Encode(res)
	}

	return res.SerializeTo(w)
}

func (cmd *submoduleSubcommand) run(request git2go.SubmoduleCommand) (*git2go.SubmoduleResult, error) {
	if request.AuthorDate.IsZero() {
		request.AuthorDate = time.Now()
	}

	smCommitOID, err := git.NewOid(request.CommitSHA)
	if err != nil {
		return nil, fmt.Errorf("converting %s to OID: %w", request.CommitSHA, err)
	}

	repo, err := git2goutil.OpenRepository(request.Repository)
	if err != nil {
		return nil, fmt.Errorf("open repository: %w", err)
	}

	fullBranchRefName := "refs/heads/" + request.Branch
	o, err := repo.RevparseSingle(fullBranchRefName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", git2go.LegacyErrPrefixInvalidBranch, err)
	}

	startCommit, err := o.AsCommit()
	if err != nil {
		return nil, fmt.Errorf("peeling %s as a commit: %w", o.Id(), err)
	}

	rootTree, err := startCommit.Tree()
	if err != nil {
		return nil, fmt.Errorf("root tree from starting commit: %w", err)
	}

	index, err := git.NewIndex()
	if err != nil {
		return nil, fmt.Errorf("creating new index: %w", err)
	}

	if err := index.ReadTree(rootTree); err != nil {
		return nil, fmt.Errorf("reading root tree into index: %w", err)
	}

	smEntry, err := index.EntryByPath(request.Submodule, 0)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: %w",
			git2go.LegacyErrPrefixInvalidSubmodulePath, err,
		)
	}

	if smEntry.Id.Cmp(smCommitOID) == 0 {
		//nolint
		return nil, fmt.Errorf(
			"The submodule %s is already at %s",
			request.Submodule, request.CommitSHA,
		)
	}

	if smEntry.Mode != git.FilemodeCommit {
		return nil, fmt.Errorf(
			"%s: %w",
			git2go.LegacyErrPrefixInvalidSubmodulePath, err,
		)
	}

	newEntry := *smEntry      // copy by value
	newEntry.Id = smCommitOID // assign new commit SHA
	if err := index.Add(&newEntry); err != nil {
		return nil, fmt.Errorf("add new submodule entry to index: %w", err)
	}

	newRootTreeOID, err := index.WriteTreeTo(repo)
	if err != nil {
		return nil, fmt.Errorf("write index to repo: %w", err)
	}

	newTree, err := repo.LookupTree(newRootTreeOID)
	if err != nil {
		return nil, fmt.Errorf("looking up new submodule entry root tree: %w", err)
	}

	committer := git.Signature(
		git2go.NewSignature(
			request.AuthorName,
			request.AuthorMail,
			request.AuthorDate,
		),
	)
	newCommitOID, err := repo.CreateCommit(
		"", // caller should update branch with hooks
		&committer,
		&committer,
		request.Message,
		newTree,
		startCommit,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: %w",
			git2go.LegacyErrPrefixFailedCommit, err,
		)
	}

	return &git2go.SubmoduleResult{
		CommitID: newCommitOID.String(),
	}, nil
}
