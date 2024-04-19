package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"todo/common"
	"todo/modules/todo/entity"
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

	// 3. Create REST APIS
	//// Naming conventions for apis
	//// POST /v1/todos (for create action)
	//// GET /v1/todos (for read action)
	//// (PUT | PATCH) /v1/todos (for update action)
	//// DELETE /v1/todos/:id (for delete action)
	// create a group apis for v1
	v1 := r.Group("/v1")
	{
		// create a group apis for "todos"
		todos := v1.Group("/todos")
		{
			todos.POST("", transport.CreateTodoItem(db))
			todos.GET("", getTodoItems(db))
			todos.GET("/:id", transport.GetTodoItem(db))
			todos.PATCH("/:id", updateTodoItemById(db))
			todos.DELETE("/:id", deleteTodoItemById(db))
		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func updateTodoItemById(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

		// Get param id from request and validate
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		// When the param id is ok -> parse data from request
		var data entity.TodoUpdateBody
		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		// update data to db
		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		// Response to client
		// ctx.JSON(http.StatusOK, gin.H{
		// 	"data": true,
		// })
		ctx.JSON(http.StatusOK, common.SingleSuccessResponse(true))
	}
}

func deleteTodoItemById(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

		// Get param id from request and validate
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		// When the param id is ok -> delete data
		updateBody := map[string]interface{}{
			"status": "Deleted",
		}
		if err := db.Table(entity.TodoItem{}.TableName()).Where("id = ?", id).Updates(updateBody).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})

			return
		}

		// Response to client
		// ctx.JSON(http.StatusOK, gin.H{
		// 	"data": true,
		// })
		ctx.JSON(http.StatusOK, common.SingleSuccessResponse(true))
	}
}

func getTodoItems(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.Paging

		err := ctx.ShouldBind(&paging) // pass pointer of data (not pass data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		paging.Process()
		var result []entity.TodoItem
		offset := (paging.Page - 1) * paging.Limit

		dbErr := db.Where("status <> ?", "Deleted").
			Order("id desc").
			Offset(offset).
			Limit(paging.Limit).
			Find(&result).Error
		if dbErr != nil {
			ctx.JSON(http.StatusCreated, gin.H{
				"error": dbErr.Error(),
			})

			return
		}

		// ctx.JSON(http.StatusOK, gin.H{
		// 	"data": result,
		// })
		ctx.JSON(http.StatusOK, common.SuccessResponse(result, paging, nil))
	}
}
