package entity

import (
	"errors"
	"todo/common"
)

const (
	EntityName = "todo_item"
)

var (
	ErrTitleBlank  = errors.New("title can not be blank")
	ErrItemDeleted = errors.New("item deleted")
)

/*
 * Create a struct that maps to a table in db.
 */
type TodoItem struct {
	common.SQLModel             // embed a struct to a struct
	Title           string      `json:"title" gorm:"column:title"`
	Description     string      `json:"description" gorm:"column:description"`
	Status          *TodoStatus `json:"status" gorm:"column:status"`
}

func (TodoItem) TableName() string { return "todo_items" }

/*
 * Create a struct that maps to the request body of POST.
 */
type TodoCreationBody struct {
	Id          int         `json:"-" gorm:"column:id"` //`json:"-": not set in post body, gorm:"column:id" -> need to get from db
	Title       string      `json:"title" gorm:"column:titless"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *TodoStatus `json:"status" gorm:"column:status"`
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
