package business_test

import (
	"context"
	"errors"
	"testing"
	"todo/modules/todo/business"
	"todo/modules/todo/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDeleteTodoStorage struct {
	mock.Mock
}

func (m *MockDeleteTodoStorage) GetTodo(ctx context.Context, cond map[string]interface{}) (*entity.TodoItem, error) {
	args := m.Called(ctx, cond)

	// Use 2 following if to fix the error "panic: interface conversion: interface {} is nil, not []entity.TodoItem"
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.TodoItem), args.Error(1)
}

func (m *MockDeleteTodoStorage) DeleteTodo(ctx context.Context, cond map[string]interface{}) error {
	args := m.Called(ctx, cond)
	return args.Error(0)
}

func TestDeleteTodoById(t *testing.T) {
	// create the mock storage object
	mockStorage := new(MockDeleteTodoStorage)

	// set up expectations when call DeleteTodo() func
	mockStorage.On("GetTodo", context.Background(), mock.Anything).Return(&entity.TodoItem{
		Title:       "test title",
		Description: "test description",
	}, nil)
	mockStorage.On("DeleteTodo", context.Background(), mock.Anything).Return(nil)

	// create biz object and call the code we are testing (DeleteTodoById())
	biz := business.NewDeleteTodoBiz(mockStorage)
	err := biz.DeleteTodoById(context.Background(), 1)

	assert.Nil(t, err)
}

func TestDeleteTodoByIdNotExist(t *testing.T) {
	// create the mock storage object
	mockStorage := new(MockDeleteTodoStorage)

	// set up expectations when call DeleteTodo() func
	mockStorage.On("GetTodo", context.Background(), mock.Anything).Return(nil, errors.New("some error"))
	mockStorage.On("DeleteTodo", context.Background(), mock.Anything).Return(nil)

	// create biz object and call the code we are testing (DeleteTodoById())
	biz := business.NewDeleteTodoBiz(mockStorage)
	err := biz.DeleteTodoById(context.Background(), 1)

	assert.Error(t, err)
}

func TestDeleteTodoFail(t *testing.T) {
	// create the mock storage object
	mockStorage := new(MockDeleteTodoStorage)

	// set up expectations when call DeleteTodo() func
	mockStorage.On("GetTodo", context.Background(), mock.Anything).Return(&entity.TodoItem{
		Title:       "test title",
		Description: "test description",
	}, nil)
	mockStorage.On("DeleteTodo", context.Background(), mock.Anything).Return(errors.New("some error"))

	// create biz object and call the code we are testing (DeleteTodoById())
	biz := business.NewDeleteTodoBiz(mockStorage)
	err := biz.DeleteTodoById(context.Background(), 1)

	assert.Error(t, err)
}
