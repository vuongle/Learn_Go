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

type MockGetTodoByIdStorage struct {
	mock.Mock
}

func (m *MockGetTodoByIdStorage) GetTodo(ctx context.Context, cond map[string]interface{}) (*entity.TodoItem, error) {
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

func TestGetTodoNotFound(t *testing.T) {

	// create the mock storage object
	mockStorage := new(MockGetTodoByIdStorage)

	// set up expectations when call GetTodo() func
	mockStorage.On("GetTodo", context.Background(), mock.Anything).Return(nil, nil)

	// create biz object and call the code we are testing (GetTodoById())
	biz := business.NewGetTodoBiz(mockStorage)
	data, err := biz.GetTodoById(context.Background(), 0)

	assert.Nil(t, data)
	assert.Nil(t, err)
}

func TestGetTodoFound(t *testing.T) {

	// create the mock storage object
	mockStorage := new(MockGetTodoByIdStorage)

	// set up expectations when call GetTodo() func
	mockStorage.On("GetTodo", context.Background(), mock.Anything).Return(&entity.TodoItem{
		Title:       "test title",
		Description: "test description",
	}, nil)

	// create biz object and call the code we are testing (GetTodoById())
	biz := business.NewGetTodoBiz(mockStorage)
	data, err := biz.GetTodoById(context.Background(), 1)

	assert.Nil(t, err)
	assert.Equal(t, "test title", data.Title)
	assert.Equal(t, "test description", data.Description)
}

func TestGetTodoError(t *testing.T) {
	// create the mock storage object
	mockStorage := new(MockGetTodoByIdStorage)

	// set up expectations when call GetTodo() func
	mockStorage.On("GetTodo", context.Background(), mock.Anything).Return(nil, errors.New("some error"))

	// create biz object and call the code we are testing (GetTodoById())
	biz := business.NewGetTodoBiz(mockStorage)
	_, err := biz.GetTodoById(context.Background(), 1)

	assert.Error(t, err)
}
