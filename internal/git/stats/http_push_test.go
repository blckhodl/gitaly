//go:build !gitaly_test_sha256

package stats

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git/gittest"
	"gitlab.com/gitlab-org/gitaly/v15/internal/gitaly/config"
	"gitlab.com/gitlab-org/gitaly/v15/internal/helper/text"
	"gitlab.com/gitlab-org/gitaly/v15/internal/testhelper"
	"gitlab.com/gitlab-org/gitaly/v15/internal/testhelper/testcfg"
)

func TestPerformHTTPPush(t *testing.T) {
	cfg, _, targetRepoPath := testcfg.BuildWithRepo(t)
	gitCmdFactory := gittest.NewCommandFactory(t, cfg)
	ctx := testhelper.Context(t)

	serverPort, stopGitServer := gittest.HTTPServer(ctx, t, gitCmdFactory, targetRepoPath, nil)
	defer func() {
		require.NoError(t, stopGitServer())
	}()
	url := fmt.Sprintf("http://localhost:%d/%s", serverPort, filepath.Base(targetRepoPath))

	for _, tc := range []struct {
		desc            string
		preparePush     func(t *testing.T, cfg config.Cfg) ([]PushCommand, io.Reader)
		expectedErr     error
		expectedTimings []string
		expectedStats   HTTPSendPack
	}{
		{
			desc: "single revision",
			preparePush: func(t *testing.T, cfg config.Cfg) ([]PushCommand, io.Reader) {
				_, repoPath := gittest.CreateRepository(ctx, t, cfg, gittest.CreateRepositoryConfig{
					SkipCreationViaService: true,
					Seed:                   gittest.SeedGitLabTest,
				})

				commit := gittest.WriteCommit(t, cfg, repoPath)
				revisions := strings.NewReader(commit.String())
				pack := gittest.ExecOpts(t, cfg, gittest.ExecConfig{Stdin: revisions},
					"-C", repoPath, "pack-objects", "--stdout", "--revs", "--thin", "--delta-base-offset", "-q",
				)

				return []PushCommand{
					{OldOID: git.ObjectHashSHA1.ZeroOID, NewOID: commit, Reference: "refs/heads/foobar"},
				}, bytes.NewReader(pack)
			},
			expectedTimings: []string{
				"start", "header", "pack-sideband", "unpack-ok", "response-body", "end",
			},
			expectedStats: HTTPSendPack{
				stats: SendPack{
					updatedRefs:       1,
					packets:           2,
					largestPacketSize: 44,
					multiband: map[string]*bandInfo{
						"pack": {
							packets: 1,
							size:    44,
						},
						"progress": {},
						"error":    {},
					},
				},
			},
		},
		{
			desc: "many revisions",
			preparePush: func(t *testing.T, cfg config.Cfg) ([]PushCommand, io.Reader) {
				_, repoPath := gittest.CreateRepository(ctx, t, cfg, gittest.CreateRepositoryConfig{
					SkipCreationViaService: true,
					Seed:                   gittest.SeedGitLabTest,
				})

				commands := make([]PushCommand, 1000)
				commits := make([]string, len(commands))
				for i := 0; i < len(commands); i++ {
					commit := gittest.WriteCommit(t, cfg, repoPath)
					commits[i] = commit.String()
					commands[i] = PushCommand{
						OldOID:    git.ObjectHashSHA1.ZeroOID,
						NewOID:    commit,
						Reference: git.ReferenceName(fmt.Sprintf("refs/heads/branch-%d", i)),
					}
				}

				revisions := strings.NewReader(strings.Join(commits, "\n"))
				pack := gittest.ExecOpts(t, cfg, gittest.ExecConfig{Stdin: revisions},
					"-C", repoPath, "pack-objects", "--stdout", "--revs", "--thin", "--delta-base-offset", "-q",
				)

				return commands, bytes.NewReader(pack)
			},
			expectedTimings: []string{
				"start", "header", "pack-sideband", "unpack-ok", "response-body", "end",
			},
			expectedStats: HTTPSendPack{
				stats: SendPack{
					updatedRefs:       1000,
					packets:           2,
					largestPacketSize: 28909,
					multiband: map[string]*bandInfo{
						"pack": {
							packets: 1,
							size:    28909,
						},
						"progress": {},
						"error":    {},
					},
				},
			},
		},
		{
			desc: "branch deletion",
			preparePush: func(t *testing.T, cfg config.Cfg) ([]PushCommand, io.Reader) {
				commit := gittest.Exec(t, cfg, "-C", targetRepoPath, "rev-parse", "refs/heads/feature")
				oldOID := git.ObjectID(text.ChompBytes(commit))

				return []PushCommand{
					{OldOID: oldOID, NewOID: git.ObjectHashSHA1.ZeroOID, Reference: "refs/heads/feature"},
				}, nil
			},
			expectedTimings: []string{
				"start", "header", "pack-sideband", "unpack-ok", "response-body", "end",
			},
			expectedStats: HTTPSendPack{
				stats: SendPack{
					updatedRefs:       1,
					packets:           2,
					largestPacketSize: 45,
					multiband: map[string]*bandInfo{
						"pack": {
							packets: 1,
							size:    45,
						},
						"progress": {},
						"error":    {},
					},
				},
			},
		},
		{
			desc: "failing delete",
			preparePush: func(t *testing.T, cfg config.Cfg) ([]PushCommand, io.Reader) {
				oldOID := git.ObjectID(strings.Repeat("1", 40))

				return []PushCommand{
					{OldOID: oldOID, NewOID: git.ObjectHashSHA1.ZeroOID, Reference: "refs/heads/master"},
				}, nil
			},
			expectedErr: fmt.Errorf("parsing packfile response: %w",
				errors.New("reference update failed: \"ng refs/heads/master deletion of the current branch prohibited\\n\"")),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			commands, packfile := tc.preparePush(t, cfg)

			start := time.Now()

			stats, err := PerformHTTPPush(ctx, url, "", "", commands, packfile, false)
			require.Equal(t, tc.expectedErr, err)
			if err != nil {
				return
			}

			end := time.Now()

			timings := map[string]*time.Time{
				"start":         &stats.SendPack.start,
				"header":        &stats.SendPack.header,
				"pack-sideband": &stats.SendPack.stats.multiband["pack"].firstPacket,
				"unpack-ok":     &stats.SendPack.stats.unpackOK,
				"response-body": &stats.SendPack.stats.responseBody,
				"end":           &end,
			}

			previousTime := start
			for _, expectedTiming := range tc.expectedTimings {
				timing := timings[expectedTiming]
				require.True(t, timing.After(previousTime),
					"expected to receive %q packet before before %q, but received at %q",
					expectedTiming, previousTime, timing)
				previousTime = *timing
				*timing = time.Time{}
			}

			stats.SendPack.stats.ReportProgress = nil
			require.Equal(t, tc.expectedStats, stats.SendPack)
		})
	}
}
