package models

import "time"

type Alert struct {
	ID              int64         `gorm:"primaryKey"`
	UserID          int64         `gorm:"type:bigint;not null"`
	AlertTemplateID int64         `gorm:"type:bigint;not null"`
	Name            string        `gorm:"type:varchar(50);not null"`
	Description     string        `gorm:"type:varchar(200)"`
	PhoneNumber     string        `gorm:"type:varchar(20);not null"`
	HourConcurrence uint16        `gorm:"type:smallint"`
	StartAt         time.Time     `gorm:"type:timestamptz;not null"`
	IsActive        bool          `gorm:"type:boolean;not null"`
	CreatedAt       time.Time     `gorm:"type:timestamptz;not null;default:now()"`
	ModifiedAt      time.Time     `gorm:"type:timestamptz;not null;default:now()"`
	AlertTemplate   AlertTemplate `gorm:"foreignKey:AlertTemplateID;references:ID"`
}
