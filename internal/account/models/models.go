package models

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AutoMigrateModels(DB *gorm.DB) {
	DB.AutoMigrate(User{})
	DB.AutoMigrate(UserPlan{})

	insertDefaultsUserPlan(DB)
}

func insertDefaultsUserPlan(DB *gorm.DB) {
	defaultPeople := []UserPlan{
		{ID: 1, Name: "Pro", Description: "Pro plan"},
		{ID: 2, Name: "Basic", Description: "Basic plan"},
	}

	err := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}}, // Check conflict on "id"
		DoNothing: true,                          // Do nothing on conflict
	}).Create(&defaultPeople).Error

	if err != nil {
		log.Fatalf("Failed to insert User Plan defaults: %v", err)
	}
}
