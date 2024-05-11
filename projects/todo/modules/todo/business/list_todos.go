package business

import (
	"context"
	"todo/common"
	"todo/modules/todo/entity"
)

// Define an interface for use case "list todo items with filters"
type ListTodosStorage interface {
	// ...string: means that there can be more params(fourth, fifth, ...) and params are optional
	ListTodos(
		ctx context.Context,
		filter *entity.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]entity.TodoItem, error)
}

// Define a struct for use case "list todo items with filters"
type listTodosBusiness struct {
	store ListTodosStorage
}

// Method of "listTodosBusiness" struct. Get a todo item from db, via the interface.
func (biz *listTodosBusiness) ListTodos(
	ctx context.Context,
	filter *entity.Filter,
	paging *common.Paging,
) ([]entity.TodoItem, error) {
	data, err := biz.store.ListTodos(ctx, filter, paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewListTodosBiz(store ListTodosStorage) *listTodosBusiness {
	return &listTodosBusiness{store: store}
}
