package storage

import (
	"context"
	"fmt"
	"todo/common"
	"todo/modules/todo/entity"

	"gorm.io/gorm"
)

func (s *sqlStore) GetTodo(ctx context.Context, cond map[string]interface{}) (*entity.TodoItem, error) {
	fmt.Println("--------------- STORAGE")
	var data entity.TodoItem
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}
