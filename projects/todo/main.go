package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
 * Create a sturct that maps to a table in db.
 */
type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"` // use pointer so that if db is null, this field still has data
	UpdatedAt   *time.Time `json:"updated_at"`
}

/*
 * Create a sturct that maps to a table in db.
 */
type TodoCreationBody struct {
	Id          int    `json:"-" gorm:"column:id"` //`json:"-": not set in post body, gorm:"column:id" -> need to get from db
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Status      string `json:"status" gorm:"column:status"`
}

func (TodoCreationBody) TableName() string { return "todo_items" }

func main() {
	//dsn := "root:root@tcp(127.0.0.1:3306)/todo_list?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("MYSQL_CON")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected to MySQL", db)

	// now := time.Now().UTC()

	// item := TodoIten{
	// 	Id:          1,
	// 	Title:       "New task",
	// 	Description: "description",
	// 	CreatedAt:   &now,
	// }

	// jsonData, err := json.Marshal(item)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }

	// fmt.Println(string(jsonData))

	// Create and start a http server
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Naming conventions for apis
	// POST /v1/todos (for create action)
	// GET /v1/todos (for read action)
	// (PUT | PATCH) /v1/todos (for update action)
	// DELETE /v1/todos/:id (for delete action)

	// create a group apis for v1
	v1 := r.Group("/v1")
	{
		// create a group apis for "todos"
		todos := v1.Group("/todos")
		{
			todos.POST("", createTodoIem(db))
			todos.GET("")
			todos.GET("/:id")
			todos.PATCH("/:id")
			todos.DELETE("/:id")
		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// Create a new todo item and return a func having type "func(*gin.Context)"
func createTodoIem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data TodoCreationBody

		err := ctx.ShouldBind(&data) // pass pointer of data (not pass data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		dbErr := db.Create(&data).Error
		if dbErr != nil {
			ctx.JSON(http.StatusCreated, gin.H{
				"error": dbErr.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": data.Id,
		})
	}
}
