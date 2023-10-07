package database

import (
	"CleanArchitecture/features/users/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
