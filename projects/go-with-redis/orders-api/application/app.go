package application

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func NewApp() *App {
	app := &App{
		rdb: redis.NewClient(&redis.Options{}),
	}

	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	if err := a.rdb.Ping(ctx).Err(); err != nil {
		return err
	}

	// Implement a feature called "graceful shutdown" in go by:
	//	1. send an error to a channel
	//	2. shutdown the server

	// Send an error to a channel
	ch := make(chan error, 1)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			ch <- errors.New(err.Error())
		}
		close(ch)
	}()

	// Listten an error in the channel
	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		server.Shutdown(timeout)
	}

	return nil
}
