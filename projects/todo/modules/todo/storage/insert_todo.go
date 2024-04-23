package storage

import (
	"context"
	"fmt"
	"todo/common"
	"todo/modules/todo/entity"
)

// Define a method belongs to sqlStore.
// With this definition, the sqlStore struct implements CreateTodoStorage interface
// Can put this method in sql.go file.
func (s *sqlStore) Create(ctx context.Context, data *entity.TodoCreationBody) error {
	fmt.Println("--------------- STORAGE")
	if dbErr := s.db.Create(&data).Error; dbErr != nil {
		// wrap the db error into pre-defined error
		return common.ErrDB(dbErr)
	}

	return nil
}
