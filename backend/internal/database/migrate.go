package database

import "github.com/Geovannisouza23/go-user-api/internal/models"

func Migrate() {
	DB.AutoMigrate(&models.User{})
}
