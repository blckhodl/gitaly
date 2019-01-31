package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/gitlab-org/gitaly/internal/praefect"
	"gitlab.com/gitlab-org/gitaly/internal/praefect/config"
)

var (
	flagConfig = flag.String("config", "", "Location for the config.toml")
	logger     = log.New()
)

func main() {
	flag.Parse()

	logger.Formatter = &log.JSONFormatter{}

	logger.WithField("config", *flagConfig).Info("reading config file")
	conf, err := config.FromFile(*flagConfig)
	if err != nil {
		logger.Fatalf("%s", err)
	}

	if err := conf.Validate(); err != nil {
		logger.Fatalf("%s", err)
	}

	l, err := net.Listen("tcp", conf.ListenAddr)
	if err != nil {
		logger.Fatalf("%s", err)
	}

	logger.WithField("address", conf.ListenAddr).Info("listening at tcp address")

	logger.Fatalf("%v", run(l, conf))
}

func run(l net.Listener, conf *config.Config) error {
	srv := praefect.NewServer(nil, log.New())

	signals := []os.Signal{syscall.SIGTERM, syscall.SIGINT}
	termCh := make(chan os.Signal, len(signals))
	signal.Notify(termCh, signals...)

	serverErrors := make(chan error, 1)
	go func() { serverErrors <- srv.Start(l) }()

	for _, gitaly := range conf.GitalyServers {
		srv.RegisterNode(gitaly.Name, gitaly.ListenAddr)

		logger.WithField("gitaly listen addr", gitaly.ListenAddr).Info("registered gitaly node")
	}

	var err error
	select {
	case s := <-termCh:
		logger.WithField("signal", s).Warn("received signal, shutting down gracefully")

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		if shutdownErr := srv.Shutdown(ctx); shutdownErr != nil {
			logger.Warn("error received during shutting down: %v", shutdownErr)
		}
		err = fmt.Errorf("received signal: %v", s)
	case err = <-serverErrors:
	}

	return err
}
