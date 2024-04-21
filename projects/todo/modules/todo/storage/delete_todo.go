package storage

import (
	"context"
	"fmt"
	"todo/modules/todo/entity"
)

func (s *sqlStore) DeleteTodo(ctx context.Context, cond map[string]interface{}) error {
	fmt.Println("--------------- STORAGE")

	// Delete logic here means soft update(change status to "Deleted")
	updateBody := map[string]interface{}{
		"status": "Deleted",
	}
	if err := s.db.Table(entity.TodoItem{}.TableName()).
		Where(cond).
		Updates(updateBody).Error; err != nil {

		return err
	}

	return nil
}
