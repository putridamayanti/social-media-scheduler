package db

import (
	"gorm.io/gorm"
	"social-media-scheduler/internal/models"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Session{})
}
