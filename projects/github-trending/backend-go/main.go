package main

import (
	"context"
	"fmt"
	"github-trending-api/db"
	"github-trending-api/handlers"
	"github-trending-api/helper"
	"github-trending-api/logger"
	repo_impl "github-trending-api/repositories/impl"
	"github-trending-api/routers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

func init() {
	// Init logger
	logger.InitLogger(false)
}

func main() {
	// Connect to mysql
	sql := &db.Sql{
		Host: "127.0.0.1", // use this line if accessing mysql in local pc
		//Host:     "host.docker.internal", // use this line if accessing mysql in a docker
		Port:     3306,
		Username: "root",
		Password: "root",
		DbName:   "github_trending",
	}
	sql.Connect()
	defer sql.Close()

	// Init api server
	e := echo.New()

	// Init a custom validator and assign it to echo validator
	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate()
	e.Validator = structValidator

	// Init handlers, repositories
	userHandler := handlers.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}
	repoHandler := handlers.RepoHandler{
		GithubRepo: repo_impl.NewGithubRepo(sql),
	}

	// Defines apis(routers)
	api := routers.API{
		Echo:        e,
		UserHandler: userHandler,
		RepoHandler: repoHandler,
	}
	api.SetupRouter()

	// start a new routine to crawl data
	go scheduleUpdateTrending(3*time.Minute, repoHandler)

	// Listen the Interrupt or shurdown from OS
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	// Start server (in docker, running 3000)
	go func() {
		if err := e.Start(":3001"); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("shutting down the server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logger.Fatalf("gracefully shutdown the server: %v", err)
	}
}

func scheduleUpdateTrending(timeSchedule time.Duration, handler handlers.RepoHandler) {

	// Create a new ticker
	// after a duration defined in "timeSchedule" -> this ticker sends a value to a channel("ticker.C") associated inside it
	ticker := time.NewTicker(timeSchedule)
	go func() {

		// for: used to always wait the "ticker.C" channel after each runs
		for {
			// select - case syntax: wait channels to ready
			// which channel is ready first -> that case runs and exits the select
			select {
			case <-ticker.C:
				fmt.Println("Crawling github trending repositories...")
				helper.CrawlTrendingRepos(handler.GithubRepo)
			}
		}
	}()
}
