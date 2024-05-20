package repositories

import (
	"blog-go-api/database"
	"blog-go-api/models"
	"errors"
)

var (
	ErrCantCreatePost = errors.New("can not create user")
)

func CreatePost(post *models.Blog) error {
	if err := database.DB.Create(post).Error; err != nil {
		return ErrCantCreatePost
	}

	return nil
}

func GetAllPosts(page *int, limit *int, offset *int) (int64, []models.Blog) {
	var total int64
	var blogs []models.Blog
	database.DB.Preload("User").Offset(*offset).Limit(*limit).Find(&blogs)
	database.DB.Model(&models.Blog{}).Count(&total)
	return total, blogs
}
