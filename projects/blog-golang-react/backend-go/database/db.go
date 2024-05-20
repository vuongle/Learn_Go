package database

import (
	"blog-go-api/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	// Connect to db
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("can not connect to db")
	}

	DB = db

	DB.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}
