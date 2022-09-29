package hook

import (
	"errors"

	gitalyerrors "gitlab.com/gitlab-org/gitaly/v15/internal/errors"
	"gitlab.com/gitlab-org/gitaly/v15/internal/gitaly/hook"
	"gitlab.com/gitlab-org/gitaly/v15/internal/gitaly/transaction"
	"gitlab.com/gitlab-org/gitaly/v15/internal/helper"
	"gitlab.com/gitlab-org/gitaly/v15/proto/go/gitalypb"
	"gitlab.com/gitlab-org/gitaly/v15/streamio"
)

func validateReferenceTransactionHookRequest(in *gitalypb.ReferenceTransactionHookRequest) error {
	if in.GetRepository() == nil {
		return gitalyerrors.ErrEmptyRepository
	}

	return nil
}

func (s *server) ReferenceTransactionHook(stream gitalypb.HookService_ReferenceTransactionHookServer) error {
	request, err := stream.Recv()
	if err != nil {
		return helper.ErrInternalf("receiving first request: %w", err)
	}

	if err := validateReferenceTransactionHookRequest(request); err != nil {
		return helper.ErrInvalidArgument(err)
	}

	var state hook.ReferenceTransactionState
	switch request.State {
	case gitalypb.ReferenceTransactionHookRequest_PREPARED:
		state = hook.ReferenceTransactionPrepared
	case gitalypb.ReferenceTransactionHookRequest_COMMITTED:
		state = hook.ReferenceTransactionCommitted
	case gitalypb.ReferenceTransactionHookRequest_ABORTED:
		state = hook.ReferenceTransactionAborted
	default:
		return helper.ErrInvalidArgument(errors.New("invalid hook state"))
	}

	stdin := streamio.NewReader(func() ([]byte, error) {
		req, err := stream.Recv()
		return req.GetStdin(), err
	})

	if err := s.manager.ReferenceTransactionHook(
		stream.Context(),
		state,
		request.GetEnvironmentVariables(),
		stdin,
	); err != nil {
		switch {
		case errors.Is(err, transaction.ErrTransactionAborted):
			return helper.ErrAbortedf("reference-transaction hook: %w", err)
		case errors.Is(err, transaction.ErrTransactionStopped):
			return helper.ErrFailedPreconditionf("reference-transaction hook: %w", err)
		default:
			return helper.ErrInternalf("reference-transaction hook: %w", err)
		}
	}

	if err := stream.Send(&gitalypb.ReferenceTransactionHookResponse{
		ExitStatus: &gitalypb.ExitStatus{Value: 0},
	}); err != nil {
		return helper.ErrInternalf("sending response: %w", err)
	}

	return nil
}
