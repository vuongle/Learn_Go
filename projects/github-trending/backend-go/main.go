package main

import (
	"github-trending-api/db"
	"github-trending-api/handlers"
	repo_impl "github-trending-api/repositories/impl"
	"github-trending-api/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	// Connect to db
	sql := &db.Sql{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "root",
		DbName:   "github_trending",
	}
	sql.Connect()
	defer sql.Close()

	// Init api server
	e := echo.New()

	// Init handlers, repositories
	userHandler := handlers.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	// Defines apis(routers)
	api := routers.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupRouter()

	// Start the server
	e.Logger.Fatal(e.Start(":3000"))
}
