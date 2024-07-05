package business_test

import (
	"context"
	"errors"
	"testing"
	"todo/common"
	"todo/modules/todo/business"
	"todo/modules/todo/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockListTodosStorage struct {
	mock.Mock
}

func (m *MockListTodosStorage) ListTodos(
	ctx context.Context,
	filter *entity.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]entity.TodoItem, error) {
	args := m.Called(ctx, filter, paging)

	// Use 2 following if to fix the error "panic: interface conversion: interface {} is nil, not []entity.TodoItem"
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]entity.TodoItem), args.Error(1)
}

func TestListTodos(t *testing.T) {

	mockStorage := new(MockListTodosStorage)

	mockStorage.On("ListTodos", mock.Anything, mock.Anything, mock.Anything).Return([]entity.TodoItem{
		{
			Title:       "test title",
			Description: "test description",
		},
		{
			Title:       "test title 2",
			Description: "test description 2",
		},
	}, nil)

	biz := business.NewListTodosBiz(mockStorage)

	data, err := biz.ListTodos(context.Background(),
		&entity.Filter{
			Status: "All",
		},
		&common.Paging{
			Page:  2,
			Limit: 10,
		})

	assert.Nil(t, err)
	assert.NotEmpty(t, data)
	assert.Equal(t, 2, len(data))
}

func TestListTodosWithErr(t *testing.T) {

	mockStorage := new(MockListTodosStorage)

	mockStorage.On("ListTodos", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("some error"))

	biz := business.NewListTodosBiz(mockStorage)

	data, err := biz.ListTodos(context.Background(),
		&entity.Filter{
			Status: "All",
		},
		&common.Paging{
			Page:  2,
			Limit: 10,
		})

	assert.Nil(t, data)
	assert.Error(t, err)
}
