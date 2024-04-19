package common

import "time"

type SQLModel struct {
	Id        int        `json:"id" gorm:"column:id"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"` // use pointer so that if db is null, this field still has data
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}
