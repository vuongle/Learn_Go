package main

import (
	"context"
	"log"
	"orders-api/application"
	"os"
	"os/signal"
)

func main() {
	app := application.NewApp()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := app.Start(ctx); err != nil {
		log.Fatal("Server not started")
	}
}
