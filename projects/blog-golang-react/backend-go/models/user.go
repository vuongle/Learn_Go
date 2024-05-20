package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  []byte `json:"password"`
	Phone     string `json:"phone"`
}

func (u *User) HashPassword(password string) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	u.Password = hashed
}

func (u *User) CompareWithHashedPassword(password string) error {
	return bcrypt.CompareHashAndPassword(u.Password, []byte(password))
}
