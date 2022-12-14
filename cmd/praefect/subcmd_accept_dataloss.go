package main

import (
	"context"
	"flag"
	"fmt"

	"gitlab.com/gitlab-org/gitaly/v15/internal/praefect/config"
	"gitlab.com/gitlab-org/gitaly/v15/proto/go/gitalypb"
)

const (
	acceptDatalossCmdName = "accept-dataloss"
)

type acceptDatalossSubcommand struct {
	virtualStorage       string
	relativePath         string
	authoritativeStorage string
}

func (cmd *acceptDatalossSubcommand) FlagSet() *flag.FlagSet {
	fs := flag.NewFlagSet(acceptDatalossCmdName, flag.ContinueOnError)
	fs.StringVar(&cmd.virtualStorage, paramVirtualStorage, "", "name of the repository's virtual storage")
	fs.StringVar(&cmd.relativePath, paramRelativePath, "", "repository to accept data loss for")
	fs.StringVar(&cmd.authoritativeStorage, paramAuthoritativeStorage, "", "storage with the repository to consider as authoritative")
	return fs
}

func (cmd *acceptDatalossSubcommand) Exec(flags *flag.FlagSet, cfg config.Config) error {
	if flags.NArg() > 0 {
		return unexpectedPositionalArgsError{Command: flags.Name()}
	} else if cmd.virtualStorage == "" {
		return requiredParameterError(paramVirtualStorage)
	} else if cmd.relativePath == "" {
		return requiredParameterError(paramRelativePath)
	} else if cmd.authoritativeStorage == "" {
		return requiredParameterError(paramAuthoritativeStorage)
	}

	nodeAddr, err := getNodeAddress(cfg)
	if err != nil {
		return err
	}
	ctx := context.TODO()

	conn, err := subCmdDial(ctx, nodeAddr, cfg.Auth.Token, defaultDialTimeout)
	if err != nil {
		return fmt.Errorf("error dialing: %w", err)
	}
	defer conn.Close()

	client := gitalypb.NewPraefectInfoServiceClient(conn)
	if _, err := client.SetAuthoritativeStorage(ctx, &gitalypb.SetAuthoritativeStorageRequest{
		VirtualStorage:       cmd.virtualStorage,
		RelativePath:         cmd.relativePath,
		AuthoritativeStorage: cmd.authoritativeStorage,
	}); err != nil {
		return fmt.Errorf("set authoritative storage: %w", err)
	}

	return nil
}
