package migrations

import (
	"gorm.io/gorm"
	"task-golang/models"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
