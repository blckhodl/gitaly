//go:build !gitaly_test_sha256

package transactions

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/gitaly/v15/internal/testhelper"
	"gitlab.com/gitlab-org/gitaly/v15/internal/transaction/voting"
)

func TestTransactionCancellationWithEmptyTransaction(t *testing.T) {
	ctx := testhelper.Context(t)

	tx, err := newTransaction(1, []Voter{
		{Name: "voter", Votes: 1},
	}, 1)
	require.NoError(t, err)

	tx.cancel()

	// When canceling a transaction, no more votes may happen.
	err = tx.vote(ctx, "voter", voting.VoteFromData([]byte{}))
	require.Error(t, err)
	require.Equal(t, err, ErrTransactionCanceled)
}

func TestTransaction_DidVote(t *testing.T) {
	ctx := testhelper.Context(t)

	tx, err := newTransaction(1, []Voter{
		{Name: "v1", Votes: 1},
		{Name: "v2", Votes: 0},
	}, 1)
	require.NoError(t, err)

	// An unregistered voter did not vote.
	require.False(t, tx.DidVote("unregistered"))
	// And neither of the registered ones did cast a vote yet.
	require.False(t, tx.DidVote("v1"))
	require.False(t, tx.DidVote("v2"))

	// One of both nodes does cast a vote.
	require.NoError(t, tx.vote(ctx, "v1", voting.VoteFromData([]byte{})))
	require.True(t, tx.DidVote("v1"))
	require.False(t, tx.DidVote("v2"))

	// And now the second node does cast a vote, too.
	require.NoError(t, tx.vote(ctx, "v2", voting.VoteFromData([]byte{})))
	require.True(t, tx.DidVote("v1"))
	require.True(t, tx.DidVote("v2"))
}

func TestTransaction_addFailedNode(t *testing.T) {
	node := "1"

	voters := []Voter{
		{Name: "1", Votes: 1, vote: nil},
		{Name: "2", Votes: 1, vote: nil},
		{Name: "3", Votes: 1, vote: nil},
	}

	transaction, err := newTransaction(1, voters, 2)
	require.NoError(t, err)

	transaction.addFailedNode(node)

	_, ok := transaction.failedNodes[node]

	require.True(t, ok)
}

func TestTransaction_isQuorumPossible(t *testing.T) {
	voters := []Voter{
		{Name: "1", Votes: 1, vote: nil},
		{Name: "2", Votes: 1, vote: nil},
		{Name: "3", Votes: 1, vote: nil},
	}

	threshold := uint(2)
	sub1, err := newSubtransaction(voters, threshold)
	require.NoError(t, err)
	sub2, err := newSubtransaction(voters, threshold)
	require.NoError(t, err)

	for _, tc := range []struct {
		desc           string
		subs           []*subtransaction
		threshold      uint
		failedNodes    map[string]struct{}
		quorumPossible bool
		quorumError    error
	}{
		{
			desc:           "No subtransactions",
			subs:           nil,
			failedNodes:    map[string]struct{}{},
			threshold:      threshold,
			quorumPossible: false,
			quorumError:    errors.New("transaction has no subtransactions"),
		},
		{
			desc:           "Single subtransaction",
			subs:           []*subtransaction{sub1},
			failedNodes:    map[string]struct{}{},
			threshold:      threshold,
			quorumPossible: true,
			quorumError:    nil,
		},
		{
			desc:           "Two subtransactions and one failed node",
			subs:           []*subtransaction{sub1, sub2},
			failedNodes:    map[string]struct{}{"1": {}},
			threshold:      threshold,
			quorumPossible: true,
			quorumError:    nil,
		},
		{
			desc:           "Two subtransactions and two failed nodes",
			subs:           []*subtransaction{sub1, sub2},
			failedNodes:    map[string]struct{}{"1": {}, "2": {}},
			threshold:      threshold,
			quorumPossible: false,
			quorumError:    nil,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			transaction, err := newTransaction(1, voters, tc.threshold)
			require.NoError(t, err)

			transaction.subtransactions = tc.subs
			transaction.failedNodes = tc.failedNodes

			possible, err := transaction.isQuorumPossible()
			require.Equal(t, tc.quorumPossible, possible)
			require.Equal(t, tc.quorumError, err)
		})
	}
}
