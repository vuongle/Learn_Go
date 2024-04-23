package storage

import (
	"context"
	"fmt"
	"todo/common"
	"todo/modules/todo/entity"
)

func (s *sqlStore) ListTodos(
	ctx context.Context,
	filter *entity.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]entity.TodoItem, error) {
	fmt.Println("--------------- STORAGE")
	var result []entity.TodoItem
	offset := (paging.Page - 1) * paging.Limit
	db := s.db.Where("status <> ?", "Deleted")

	// After getting todos that are not "Deleted" -> cnt check filter and add filter
	// to Where
	if f := filter; f != nil {
		if s := f.Status; s != "" {
			db.Where("status = ?", s)
		}
	}

	dbErr := db.Table(entity.TodoItem{}.TableName()).
		Order("id desc").
		Offset(offset).
		Limit(paging.Limit).
		Find(&result).Error
	if dbErr != nil {
		return nil, dbErr
	}

	return result, nil
}
