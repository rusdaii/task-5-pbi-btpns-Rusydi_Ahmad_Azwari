package database

import (
	"go-project/models"
)

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Photo{})
}
