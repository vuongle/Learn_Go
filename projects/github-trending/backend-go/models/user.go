package models

import "time"

type User struct {
	UserId    string    `json:"user_id" db:"user_id,omitempty"`
	FullName  string    `json:"full_name,omitempty" db:"full_name,omitempty"`
	Email     string    `json:"email,omitempty" db:"email,omitempty"`
	Password  string    `json:"password,omitempty" db:"password,omitempty"`
	Role      string    `json:"role,omitempty" db:"role,omitempty"`
	Token     string    `json:"-"`
	CreatedAt time.Time `json:"-" db:"created_at,omitempty"`
	UpdatedAt time.Time `json:"-" db:"updated_at,omitempty"`
}
