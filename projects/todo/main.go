package main

import (
	"fmt"
	"log"
	"os"
	"todo/middleware"
	transport "todo/modules/todo/transport/gin"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Entry point of go app
func main() {

	// 1. Connect to mysql
	dsn := os.Getenv("MYSQL_CON")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected to MySQL", db)

	// 2. Start a http server
	r := gin.Default()

	// Add middleware to all apis : 1st way
	r.Use(middleware.Recovery())

	// 3. Create REST APIS
	//// Naming conventions for apis
	//// POST /v1/todos (for create action)
	//// GET /v1/todos (for read action)
	//// (PUT | PATCH) /v1/todos (for update action)
	//// DELETE /v1/todos/:id (for delete action)

	// create a group apis for v1
	v1 := r.Group("/v1" /*,middleware.Recovery()*/) // Add middleware to one group : 2nd way
	{
		// create a group apis for "todos"
		todos := v1.Group("/todos")
		{
			todos.POST("" /*middleware.Recovery(), */, transport.CreateTodoItem(db)) // Add middleware to one api : 3rd way
			todos.GET("", transport.ListTodoItems(db))
			todos.GET("/:id", transport.GetTodoItem(db))
			todos.PATCH("/:id", transport.UpdateTodoItemById(db))
			todos.DELETE("/:id", transport.DeleteTodoItemById(db))
		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
