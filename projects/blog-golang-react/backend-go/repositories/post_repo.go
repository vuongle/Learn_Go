package repositories

import (
	"blog-go-api/database"
	"blog-go-api/models"
	"errors"
)

var (
	ErrCantCreatePost = errors.New("can not create post")
	ErrCantUpdatePost = errors.New("can not create post")
	ErrCantDeletePost = errors.New("can not remove post")
)

func CreatePost(post *models.Blog) error {
	if err := database.DB.Create(post).Error; err != nil {
		return ErrCantCreatePost
	}

	return nil
}

func UpdatePost(post *models.Blog) error {
	if err := database.DB.Model(post).Updates(*post).Error; err != nil {
		return ErrCantUpdatePost
	}

	return nil
}

func DeletePost(post *models.Blog) error {
	if err := database.DB.Delete(&post).Error; err != nil {
		return ErrCantDeletePost
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

func GetPostById(id *int) models.Blog {
	var blog models.Blog
	database.DB.Where("id=?", id).Preload("User").First(&blog)

	return blog
}

func GetPostsByUserId(userId *string) []models.Blog {
	var blogs []models.Blog
	database.DB.Model(&blogs).Where("user_id=?", userId).Preload("User").Find(&blogs)

	return blogs
}
