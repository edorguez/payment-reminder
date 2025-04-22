package models

import "time"

type UserPlan struct {
	ID          int64     `gorm:"primaryKey"`
	Name        string    `gorm:"type:varchar(50);not null"`
	Description string    `gorm:"type:varchar(200);not null"`
	CreatedAt   time.Time `gorm:"type:timestamptz;not null;default:now()"`
	ModifiedAt  time.Time `gorm:"type:timestamptz;not null;default:now()"`
}
