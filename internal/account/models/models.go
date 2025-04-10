package models

import "gorm.io/gorm"

func AutoMigrateModels(DB *gorm.DB) {
	DB.AutoMigrate(User{})
	DB.AutoMigrate(UserPlan{})
}
