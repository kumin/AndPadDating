package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	apps "github.com/kumin/BityDating/apps/server-ctl"
	"github.com/kumin/BityDating/monitor/instrumentation"
	zerolog "github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGSYS)
	defer stop()
	server, err := apps.BuildServer()
	if err != nil {
		log.Fatal(err)
	}
	metricServer, err := apps.BuildMetricServer()
	if err != nil {
		log.Fatal(err)
	}
	shutdown, err := instrumentation.SetupInstrument(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := shutdown(ctx)
		if err != nil {
			zerolog.Error().Msgf(err.Error())
		}
	}()

	ge, ctx := errgroup.WithContext(ctx)
	ge.Go(func() error {
		return server.Start(ctx)
	})
	ge.Go(func() error {
		return metricServer.Start(ctx)
	})
	ge.Go(func() error {
		for range ctx.Done() {
			stop()
			break
		}
		return nil
	})
	if err := ge.Wait(); err != nil {
		panic(err)
	}
}
