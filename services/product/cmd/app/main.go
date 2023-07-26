package main

import (
	"context"
	"os"
	"syscall"
	"time"

	"github.com/pwnedgod/tanshogyo/pkg/common/util/signal"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app"
)

const DurationShutdownTimeout = 60 * time.Second

func main() {
	application, cleanup, err := app.InjectApplication()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	listener := func(os.Signal) {
		ctx, cancel := context.WithTimeout(context.Background(), DurationShutdownTimeout)
		application.Stop(ctx)
		cancel()
	}

	signal.On(syscall.SIGTERM, listener)
	signal.On(syscall.SIGINT, listener)

	signal.Listen()
	defer signal.Close()

	if err := application.Run(context.Background()); err != nil {
		panic(err)
	}
}
