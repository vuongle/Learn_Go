package storage

import (
	"context"
	"fmt"
	"todo/modules/todo/entity"
)

func (s *sqlStore) UpdateTodo(ctx context.Context, cond map[string]interface{}, dataUpdate *entity.TodoUpdateBody) error {
	fmt.Println("--------------- STORAGE")
	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}
