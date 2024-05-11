package business

import (
	"context"
	"todo/common"
	"todo/modules/todo/entity"
)

// Define an interface for use case "update todo item by id (or by another column)"
type UpdateTodoStorage interface {
	GetTodo(ctx context.Context, cond map[string]interface{}) (*entity.TodoItem, error)
	UpdateTodo(ctx context.Context, cond map[string]interface{}, data *entity.TodoUpdateBody) error
}

// Define a struct for use case "update todo item by id (or by another column)"
type updateTodoBusiness struct {
	store UpdateTodoStorage
}

// Initialize an instance of "updateTodoBusiness" struct
func NewUpdateTodoBiz(store UpdateTodoStorage) *updateTodoBusiness {
	return &updateTodoBusiness{store: store}
}

// Method of "getTodoBusiness" struct. Get a todo item from db, via the interface.
func (biz *updateTodoBusiness) UpdateTodoById(ctx context.Context, id int, dataUpdate *entity.TodoUpdateBody) error {

	// First, get data from db and check(assumption this is a logic)
	data, err := biz.store.GetTodo(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrCannotGetEntity(entity.EntityName, err)
		}

		return common.ErrCannotUpdateEntity(entity.EntityName, err)
	}

	// because "data" is a pointer. Yo get value from the pointer, use "*"
	if data.Status != nil && *data.Status == entity.StatusDeleted {
		return common.ErrEntityDeleted(entity.EntityName, entity.ErrItemDeleted)
	}

	// Second, start updating
	if err := biz.store.UpdateTodo(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(entity.EntityName, err)
	}

	return nil
}
