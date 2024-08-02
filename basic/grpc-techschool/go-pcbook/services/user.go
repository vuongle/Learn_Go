package services

import (
	"fmt"

	"github.com/vuongle/grpc/entities"
	"github.com/vuongle/grpc/storages"
	"golang.org/x/crypto/bcrypt"
)

func NewUser(username string, password string, role string) (*entities.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}

	user := &entities.User{
		Username:     username,
		HashPassword: string(hashPassword),
		Role:         role,
	}

	return user, nil
}

func SeedUsers(userStore storages.UserStore) error {
	err := CreateUser(userStore, "admin1", "secret", "admin")
	if err != nil {
		return err
	}

	return CreateUser(userStore, "user1", "secret", "user")
}

func CreateUser(
	userStore storages.UserStore,
	username string,
	password string,
	role string) error {
	user, err := NewUser(username, password, role)
	if err != nil {
		return err
	}

	return userStore.Save(user)
}
