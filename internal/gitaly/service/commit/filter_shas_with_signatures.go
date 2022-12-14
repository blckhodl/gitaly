package commit

import (
	"context"
	"errors"
	"io"

	"gitlab.com/gitlab-org/gitaly/v15/internal/git"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git/catfile"
	"gitlab.com/gitlab-org/gitaly/v15/internal/helper"
	"gitlab.com/gitlab-org/gitaly/v15/proto/go/gitalypb"
)

func (s *server) FilterShasWithSignatures(bidi gitalypb.CommitService_FilterShasWithSignaturesServer) error {
	firstRequest, err := bidi.Recv()
	if err != nil {
		return err
	}

	if err = validateFirstFilterShasWithSignaturesRequest(firstRequest); err != nil {
		return helper.ErrInvalidArgument(err)
	}

	if err := s.filterShasWithSignatures(bidi, firstRequest); err != nil {
		return helper.ErrInternal(err)
	}
	return nil
}

func validateFirstFilterShasWithSignaturesRequest(in *gitalypb.FilterShasWithSignaturesRequest) error {
	if in.Repository == nil {
		return errors.New("no repository given")
	}
	return nil
}

func (s *server) filterShasWithSignatures(bidi gitalypb.CommitService_FilterShasWithSignaturesServer, firstRequest *gitalypb.FilterShasWithSignaturesRequest) error {
	ctx := bidi.Context()
	repo := s.localrepo(firstRequest.GetRepository())

	objectReader, cancel, err := s.catfileCache.ObjectReader(ctx, repo)
	if err != nil {
		return err
	}
	defer cancel()

	request := firstRequest
	for {
		shas, err := filterCommitShasWithSignatures(ctx, objectReader, request.GetShas())
		if err != nil {
			return err
		}

		if err := bidi.Send(&gitalypb.FilterShasWithSignaturesResponse{Shas: shas}); err != nil {
			return err
		}

		request, err = bidi.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}
	}
}

func filterCommitShasWithSignatures(ctx context.Context, objectReader catfile.ObjectReader, shas [][]byte) ([][]byte, error) {
	var foundShas [][]byte
	for _, sha := range shas {
		commit, err := catfile.GetCommit(ctx, objectReader, git.Revision(sha))
		if catfile.IsNotFound(err) {
			continue
		}

		if err != nil {
			return nil, err
		}

		if commit.SignatureType == gitalypb.SignatureType_NONE {
			continue
		}

		foundShas = append(foundShas, sha)
	}

	return foundShas, nil
}
