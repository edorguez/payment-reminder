package models

import "gorm.io/gorm"

func AutoMigrateModels(DB *gorm.DB) {
	DB.AutoMigrate(AlertTemplate{})
	DB.AutoMigrate(Alert{})
}
