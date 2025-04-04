package models

import "time"

type User struct {
	ID              int64     `gorm:"primaryKey"`
	IDUserPlan      int64     `gorm:"type:bigserial;not null"`
	UserPlan        UserPlan  `gorm:"foreignKey:IDUserPlan;references:ID"`
	Email           string    `gorm:"type:varchar(100);not null"`
	LastPaymentDate time.Time `gorm:"type:timestamptz;not null"`
	CreatedAt       time.Time `gorm:"type:timestamptz;not null;default:now()"`
	ModifiedAt      time.Time `gorm:"type:timestamptz;not null;default:now()"`
}
