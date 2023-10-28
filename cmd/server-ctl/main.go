package main

import (
	"context"
	"log"
	"os/signal"
	"sync"
	"syscall"

	apps "github.com/kumin/AndPadDating/apps/server-ctl"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGSYS)
	defer done()
	server, err := apps.BuildServer()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		go server.Start(ctx)
		select {
		case <-ctx.Done():
			cancel()
		}
		defer func() {
			wg.Done()
		}()
	}(ctx)
	wg.Wait()
}
