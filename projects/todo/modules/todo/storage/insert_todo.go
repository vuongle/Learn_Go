package storage

import (
	"context"
	"fmt"
	"todo/modules/todo/entity"
)

// Define a method belongs to sqlStore.
// With this definition, the sqlStore struct implements CreateTodoStorage interface
// Can put this method in sql.go file.
func (s *sqlStore) Create(ctx context.Context, data *entity.TodoCreationBody) error {
	fmt.Println("--------------- STORAGE")
	dbErr := s.db.Create(&data).Error
	return dbErr
}
