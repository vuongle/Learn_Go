package business

import (
	"context"
	"todo/common"
	"todo/modules/todo/entity"
)

// Define an interface for use case "get todo item by id (or by another column)"
type GetTodoStorage interface {
	GetTodo(ctx context.Context, cond map[string]interface{}) (*entity.TodoItem, error)
}

// Define a struct for use case "get todo item by id (or by another column)"
type getTodoBusiness struct {
	store GetTodoStorage
}

// Method of "getTodoBusiness" struct. Get a todo item from db, via the interface.
func (biz *getTodoBusiness) GetTodoById(ctx context.Context, id int) (*entity.TodoItem, error) {
	data, err := biz.store.GetTodo(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(entity.EntityName, err)
	}

	return data, nil
}

func NewGetTodoBiz(store GetTodoStorage) *getTodoBusiness {
	return &getTodoBusiness{store: store}
}
