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

type MockCreateTodoStorage struct {
	mock.Mock
}

func (m *MockCreateTodoStorage) Create(ctx context.Context, data *entity.TodoCreationBody) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func TestCreateTodoNoErr(t *testing.T) {

	// create the mock storage object
	mockStorage := new(MockCreateTodoStorage)

	// set up expectations when call Create() func
	mockStorage.On("Create", context.Background(), mock.Anything).Return(nil)

	// create biz object and call the code we are testing (CreateTodo())
	biz := business.NewCreateTodoBiz(mockStorage)
	err := biz.CreateTodo(context.Background(), &entity.TodoCreationBody{

		Title:       "test",
		Description: "description",
	})

	assert.Nil(t, err)
}

func TestCreateTodoWithErr(t *testing.T) {

	// create the mock storage object
	mockStorage := new(MockCreateTodoStorage)

	// set up expectations when call Create() func
	mockStorage.On("Create", context.Background(), mock.Anything).Return(errors.New("some error"))

	// create biz object and call the code we are testing (CreateTodo())
	biz := business.NewCreateTodoBiz(mockStorage)
	err := biz.CreateTodo(context.Background(), &entity.TodoCreationBody{

		Title:       "test",
		Description: "description",
	})

	assert.NotNil(t, err)
}

func TestCreateTodoWithEmptyTitle(t *testing.T) {

	// create the mock storage object
	mockStorage := new(MockCreateTodoStorage)

	// set up expectations when call Create() func
	mockStorage.On("Create", context.Background(), mock.Anything).Return(entity.ErrTitleBlank)

	// create biz object and call the code we are testing (CreateTodo())
	biz := business.NewCreateTodoBiz(mockStorage)
	err := biz.CreateTodo(context.Background(), &entity.TodoCreationBody{

		Title:       "",
		Description: "description",
	})

	assert.Equal(t, entity.ErrTitleBlank, err)
}
