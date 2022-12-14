package ssh

import (
	"context"
	"fmt"
	"io"

	"gitlab.com/gitlab-org/gitaly/v15/internal/command"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git"
	"gitlab.com/gitlab-org/gitaly/v15/internal/git/pktline"
)

func monitorStdinCommand(ctx context.Context, gitCmdFactory git.CommandFactory, stdin io.Reader, stdout, stderr io.Writer, sc git.SubCmd, opts ...git.CmdOpt) (*command.Command, *pktline.ReadMonitor, error) {
	stdinPipe, monitor, cleanup, err := pktline.NewReadMonitor(ctx, stdin)
	if err != nil {
		return nil, nil, fmt.Errorf("create monitor: %v", err)
	}

	cmd, err := gitCmdFactory.NewWithoutRepo(ctx, sc, append([]git.CmdOpt{
		git.WithStdin(stdinPipe),
		git.WithStdout(stdout),
		git.WithStderr(stderr),
		git.WithFinalizer(func(*command.Command) { cleanup() }),
	}, opts...)...)
	stdinPipe.Close() // this now belongs to cmd
	if err != nil {
		cleanup()
		return nil, nil, fmt.Errorf("start cmd: %v", err)
	}

	return cmd, monitor, err
}
