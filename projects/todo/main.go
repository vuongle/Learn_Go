package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Create enum for status. Data type of enum is int therefore, use iota. iota starts by 0, ...
// Therefore,  TodoStatus has value 0 | 1 | 2
type TodoStatus int

const (
	StatusDoing TodoStatus = iota
	StatusDone
	StatusDeleted
)

var allStatuses = [3]string{"Doing", "Done", "Deleted"}

// Define a method "String" with value receiver(item TodoStatus)
// Logic: Convery int to string
func (item TodoStatus) String() string {
	return allStatuses[item]
}

// Define a function to parse a string to int
// This function returns 2 value: int(TodoStatus) and error
func parseStringToTodoStatus(s string) (TodoStatus, error) {
	for i := range allStatuses {
		if allStatuses[i] == s {
			return TodoStatus(i), nil
		}
	}

	return TodoStatus(0), errors.New("invalid status string")
}

// Define a method "Scan" with pointer receiver(item *TodoStatus)
// Logic: scan value from db -> convert to bytes -> convert to string
// This method is automatically called by GORM
func (item *TodoStatus) Scan(value interface{}) error {
	// convert value from db to byte array
	b, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
	}

	// then convert byte array to string
	v, err := parseStringToTodoStatus(string(b))
	if err != nil {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
	}

	// change value of the pointer. use syntax "* + pointer variable"
	// first: the pointer points to int
	// now: points to string
	*item = v

	// no error -> return nil. If there is any errors -> already returned above
	return nil
}

// This method is automatically called by GORM
// Logic: convert int to json string(this json string includes " + value + ") in byte form
func (item *TodoStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

/*
 * Create a struct that maps to a table in db.
 */
type TodoItem struct {
	Id          int         `json:"id" gorm:"column:id"`
	Title       string      `json:"title" gorm:"column:title"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *TodoStatus `json:"status" gorm:"column:status"`
	CreatedAt   *time.Time  `json:"created_at" gorm:"column:created_at"` // use pointer so that if db is null, this field still has data
	UpdatedAt   *time.Time  `json:"updated_at" gorm:"column:updated_at"`
}

func (TodoItem) TableName() string { return "todo_items" }

/*
 * Create a struct that maps to the request body of POST.
 */
type TodoCreationBody struct {
	Id          int    `json:"-" gorm:"column:id"` //`json:"-": not set in post body, gorm:"column:id" -> need to get from db
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
}

func (TodoCreationBody) TableName() string { return TodoItem{}.TableName() }

type TodoUpdateBody struct {
	// Reason to use pointer here (not use normal data type like aboth structs), because:
	// GORM only update fields having actual value and correct data type
	// if pass title as "" or "12" -> GORM ignores
	// therefore; if want GORM to update, use pointer. Because,
	// when set title="" -> its pointer is not nil ====> GORM updates
	Title       *string `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Status      *string `json:"status" gorm:"column:status"`
}

func (TodoUpdateBody) TableName() string { return TodoItem{}.TableName() }

// Define a struct and its methods for paging
type Paging struct {
	// use tag "form" because page/limit are passed as query string ?page=1&limit=3
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"` // not pass -> set "-"
}

func (p *Paging) process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 1
	}
}

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
			todos.POST("", createTodoItem(db))
			todos.GET("", getTodoItems(db))
			todos.GET("/:id", getTodoItemById(db))
			todos.PATCH("/:id", updateTodoItemById(db))
			todos.DELETE("/:id", deleteTodoItemById(db))
		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// Create a new todo item and return a func having type "func(*gin.Context)"
func createTodoItem(db *gorm.DB) func(*gin.Context) {
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

func getTodoItemById(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

		// Get param id from request and validate
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		// When the param id is ok -> query data
		var data TodoItem
		data.Id = id
		if err := db.First(&data).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})

			return
		}

		// Response to client
		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
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
		var data TodoUpdateBody
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
		ctx.JSON(http.StatusOK, gin.H{
			"data": true,
		})
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
		if err := db.Table(TodoItem{}.TableName()).Where("id = ?", id).Updates(updateBody).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})

			return
		}

		// Response to client
		ctx.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}

func getTodoItems(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var paging Paging

		err := ctx.ShouldBind(&paging) // pass pointer of data (not pass data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		paging.process()
		var result []TodoItem
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

		ctx.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}
