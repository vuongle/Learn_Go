package business

import (
	"context"
	"todo/modules/todo/entity"
)

// Define an interface for use case "update todo item by id (or by another column)"
type DeleteTodoStorage interface {
	GetTodo(ctx context.Context, cond map[string]interface{}) (*entity.TodoItem, error)
	DeleteTodo(ctx context.Context, cond map[string]interface{}) error
}

// Define a struct for use case "update todo item by id (or by another column)"
type deleteTodoBusiness struct {
	store DeleteTodoStorage
}

// Initialize an instance of "deleteTodoBusiness" struct
func NewDeleteTodoBiz(store DeleteTodoStorage) *deleteTodoBusiness {
	return &deleteTodoBusiness{store: store}
}

// Method of "getTodoBusiness" struct. Get a todo item from db, via the interface.
func (biz *deleteTodoBusiness) DeleteTodoById(ctx context.Context, id int) error {

	// First, get data from db and check(assumption this is a logic)
	data, err := biz.store.GetTodo(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	// because "data" is a pointer. Yo get value from the pointer, use "*"
	if data.Status != nil && *data.Status == entity.StatusDeleted {
		return entity.ErrItemDeleted
	}

	// Second, start deleting
	if err := biz.store.DeleteTodo(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
