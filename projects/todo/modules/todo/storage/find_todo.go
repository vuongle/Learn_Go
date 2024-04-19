package storage

import (
	"context"
	"fmt"
	"todo/modules/todo/entity"
)

func (s *sqlStore) GetTodo(ctx context.Context, cond map[string]interface{}) (*entity.TodoItem, error) {
	fmt.Println("--------------- STORAGE")
	var data entity.TodoItem
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
