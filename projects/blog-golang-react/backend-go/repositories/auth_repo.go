package repositories

import (
	"blog-go-api/database"
	"blog-go-api/models"
	"errors"
)

var (
	ErrCantRegisterUser = errors.New("can not create user")
)

func IsEmailExisted(email string) bool {
	var userData models.User
	database.DB.Where("email=?", email).First(&userData)

	return userData.Id != 0
}

func FindUserByEmail(email string) (*models.User, error) {
	var userData models.User
	if err := database.DB.Where("email=?", email).First(&userData).Error; err != nil {
		return nil, err
	}

	return &userData, nil
}

func RegisterUser(user *models.User) error {

	if err := database.DB.Create(user).Error; err != nil {
		return ErrCantRegisterUser
	}

	return nil
}
