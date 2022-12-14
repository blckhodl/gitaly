//go:build !gitaly_test_sha256

package backup

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/gitaly/v15/internal/testhelper"
	_ "gocloud.dev/blob/memblob"
)

func TestStorageServiceSink(t *testing.T) {
	t.Parallel()
	ctx := testhelper.Context(t)

	sss, err := NewStorageServiceSink(ctx, "mem://test_bucket")
	require.NoError(t, err)
	defer func() { require.NoError(t, sss.Close()) }()

	t.Run("write and retrieve", func(t *testing.T) {
		const relativePath = "path/to/data"

		data := []byte("test")

		require.NoError(t, sss.Write(ctx, relativePath, bytes.NewReader(data)))

		reader, err := sss.GetReader(ctx, relativePath)
		require.NoError(t, err)
		defer func() { require.NoError(t, reader.Close()) }()

		retrieved, err := io.ReadAll(reader)
		require.NoError(t, err)
		require.Equal(t, data, retrieved)
	})

	t.Run("not existing path", func(t *testing.T) {
		reader, err := sss.GetReader(ctx, "not-existing")
		require.Equal(t, fmt.Errorf(`storage service sink: new reader for "not-existing": %w`, ErrDoesntExist), err)
		require.Nil(t, reader)
	})
}
