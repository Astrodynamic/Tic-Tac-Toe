package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"tictactoe/internal/di"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(di.Module())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := app.Start(ctx); err != nil {
		panic(err)
	}

	<-ctx.Done()

	if err := app.Stop(context.Background()); err != nil {
		panic(err)
	}
}
