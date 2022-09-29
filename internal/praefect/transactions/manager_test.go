package transactions

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/gitaly/v15/internal/praefect/config"
	"gitlab.com/gitlab-org/gitaly/v15/internal/testhelper"
)

func TestManager_FailTransactionNode_Error(t *testing.T) {
	voters := []Voter{
		{Name: "1", Votes: 1, vote: nil},
		{Name: "2", Votes: 1, vote: nil},
		{Name: "3", Votes: 1, vote: nil},
	}

	manager := NewManager(config.Config{})

	err := manager.FailTransactionNode(0, voters[0].Name)
	require.ErrorIs(t, err, ErrNotFound)
}

func TestManager_FailTransactionNode_Success(t *testing.T) {
	ctx := testhelper.Context(t)

	voters := []Voter{
		{Name: "1", Votes: 1, vote: nil},
		{Name: "2", Votes: 1, vote: nil},
		{Name: "3", Votes: 1, vote: nil},
	}

	manager := NewManager(config.Config{})

	transaction, c, err := manager.RegisterTransaction(ctx, voters, 2)
	require.NoError(t, err)
	defer func() {
		err := c()
		require.NoError(t, err)
	}()

	err = manager.FailTransactionNode(transaction.ID(), voters[0].Name)
	require.NoError(t, err)
}

func TestManager_IsTransactionQuorumPossible_Error(t *testing.T) {
	manager := NewManager(config.Config{})

	possible, err := manager.IsTransactionQuorumPossible(0)
	require.ErrorIs(t, err, ErrNotFound)
	require.False(t, possible)
}

func TestManager_IsTransactionQuorumPossible_Success(t *testing.T) {
	ctx := testhelper.Context(t)

	voters := []Voter{
		{Name: "1", Votes: 1, vote: nil},
		{Name: "2", Votes: 1, vote: nil},
		{Name: "3", Votes: 1, vote: nil},
	}

	manager := NewManager(config.Config{})

	transaction, c, err := manager.RegisterTransaction(ctx, voters, 2)
	require.NoError(t, err)
	defer func() {
		err := c()
		require.NoError(t, err)
	}()

	tr := manager.transactions[transaction.ID()]
	_, err = tr.getOrCreateSubtransaction(voters[0].Name)
	require.NoError(t, err)

	possible, err := manager.IsTransactionQuorumPossible(transaction.ID())
	require.NoError(t, err)
	require.True(t, possible)
}
