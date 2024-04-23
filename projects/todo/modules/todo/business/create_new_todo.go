package business

import (
	"context"
	"fmt"
	"strings"
	"todo/common"
	"todo/modules/todo/entity"
)

// Create an interface. The interface is intermediate layer between business layer and storage layer.
// business layer does not known the implementation of the storage layer. It access the storage layer via the interface
// when the storage is changed(example from mysql to mongdb) -> the storage's implementation
// is changed but the interface is not changed -> so the business layer is not changed.
// This interface is implemented by "sqlStore" struct because the "sqlStore" struct has the Create() methods having same signatures
// In Golang, the interface will be defined at the places where it is used. Therefore, the interface is defined here, not in storage layer.
type CreateTodoStorage interface {
	Create(ctx context.Context, data *entity.TodoCreationBody) error
}

// Define a struct and its methods for business layer. Because this struct is not access from external, must define NewCreateTodoBiz() function
// to return this struct itself.
type createTodoBusiness struct {
	store CreateTodoStorage
}

// Method of "createTodoBusiness" struct. Insert a todo item into db, via the interface.
func (biz *createTodoBusiness) CreateTodo(ctx context.Context, data *entity.TodoCreationBody) error {
	fmt.Println("--------------- BUSINESS")
	// assume that the following is logic of business layer
	title := strings.TrimSpace(data.Title)
	if title == "" {
		return entity.ErrTitleBlank
	}

	if err := biz.store.Create(ctx, data); err != nil {
		// wrap the err to pre-defined custom error that matches with this logic
		return common.ErrCannotCreateEntity(entity.EntityName, err)
	}

	return nil
}

func NewCreateTodoBiz(store CreateTodoStorage) *createTodoBusiness {
	return &createTodoBusiness{store: store}
}
