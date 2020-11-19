package hook

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/gitaly/internal/gitaly/config"
	"gitlab.com/gitlab-org/gitaly/internal/testhelper"
)

func TestUpdate_customHooks(t *testing.T) {
	repo, repoPath, cleanup := testhelper.NewTestRepo(t)
	defer cleanup()

	hookManager := NewManager(GitlabAPIStub, config.Config)

	standardEnv := []string{
		fmt.Sprintf("GITALY_SOCKET=%s", config.Config.GitalyInternalSocketPath()),
		"GITALY_TOKEN=secret",
		"GL_ID=1234",
		fmt.Sprintf("GL_PROJECT_PATH=%s", repo.GetGlProjectPath()),
		"GL_PROTOCOL=web",
		fmt.Sprintf("GL_REPO=%s", repo),
		fmt.Sprintf("GL_REPOSITORY=%s", repo.GetGlRepository()),
		"GL_USERNAME=user",
	}

	hash1 := strings.Repeat("1", 40)
	hash2 := strings.Repeat("2", 40)

	testCases := []struct {
		desc           string
		env            []string
		hook           string
		reference      string
		oldHash        string
		newHash        string
		expectedErr    string
		expectedStdout string
		expectedStderr string
	}{
		{
			desc:           "hook receives environment variables",
			env:            standardEnv,
			hook:           "#!/bin/sh\nenv | grep -e '^GL_' -e '^GITALY_' | sort\n",
			expectedStdout: strings.Join(standardEnv, "\n") + "\n",
		},
		{
			desc:           "hook receives arguments",
			env:            standardEnv,
			reference:      "refs/heads/master",
			oldHash:        hash1,
			newHash:        hash2,
			hook:           "#!/bin/sh\nprintf '%s\\n' \"$@\"\n",
			expectedStdout: fmt.Sprintf("refs/heads/master\n%s\n%s\n", hash1, hash2),
		},
		{
			desc:           "stdout and stderr are passed through",
			env:            standardEnv,
			reference:      "refs/heads/master",
			oldHash:        hash1,
			newHash:        hash2,
			hook:           "#!/bin/sh\necho foo >&1\necho bar >&2\n",
			expectedStdout: "foo\n",
			expectedStderr: "bar\n",
		},
		{
			desc:      "standard input is empty",
			env:       standardEnv,
			reference: "refs/heads/master",
			oldHash:   hash1,
			newHash:   hash2,
			hook:      "#!/bin/sh\ncat\n",
		},
		{
			desc:        "invalid script causes failure",
			env:         standardEnv,
			reference:   "refs/heads/master",
			oldHash:     hash1,
			newHash:     hash2,
			hook:        "",
			expectedErr: "exec format error",
		},
		{
			desc:        "errors are passed through",
			env:         standardEnv,
			reference:   "refs/heads/master",
			oldHash:     hash1,
			newHash:     hash2,
			hook:        "#!/bin/sh\nexit 123\n",
			expectedErr: "exit status 123",
		},
		{
			desc:           "errors are passed through with stderr and stdout",
			env:            standardEnv,
			reference:      "refs/heads/master",
			oldHash:        hash1,
			newHash:        hash2,
			hook:           "#!/bin/sh\necho foo >&1\necho bar >&2\nexit 123\n",
			expectedStdout: "foo\n",
			expectedStderr: "bar\n",
			expectedErr:    "exit status 123",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ctx, cleanup := testhelper.Context()
			defer cleanup()

			cleanup, err := testhelper.WriteCustomHook(repoPath, "update", []byte(tc.hook))
			require.NoError(t, err)
			defer cleanup()

			var stdout, stderr bytes.Buffer
			err = hookManager.UpdateHook(ctx, repo, tc.reference, tc.oldHash, tc.newHash, tc.env, &stdout, &stderr)

			if tc.expectedErr != "" {
				require.Contains(t, err.Error(), tc.expectedErr)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tc.expectedStdout, stdout.String())
			require.Equal(t, tc.expectedStderr, stderr.String())
		})
	}
}
