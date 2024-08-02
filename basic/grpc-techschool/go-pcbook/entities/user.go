package entities

import "golang.org/x/crypto/bcrypt"

type User struct {
	Username     string
	HashPassword string
	Role         string
}

func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(password))

	return err == nil
}

func (user *User) Clone() *User {
	return &User{
		Username:     user.Username,
		HashPassword: user.HashPassword,
		Role:         user.Role,
	}
}
