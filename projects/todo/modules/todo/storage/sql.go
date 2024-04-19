package storage

import "gorm.io/gorm"

// Contain an instance of underlying mysql. Because this struct is not access from external, must define NewSQLStore() function
// to return this struct itself.
type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
