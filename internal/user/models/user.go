package models

import "time"

type User struct {
	ID              int64     `gorm:"primaryKey"`
	UserPlanID      int64     `gorm:"type:bigserial;not null"`
	Email           string    `gorm:"type:varchar(100);not null;unique_index"`
	LastPaymentDate time.Time `gorm:"type:timestamptz;not null"`
	CreatedAt       time.Time `gorm:"type:timestamptz;not null;default:now()"`
	ModifiedAt      time.Time `gorm:"type:timestamptz;not null;default:now()"`
	UserPlan        UserPlan  `gorm:"foreignKey:UserPlanID;references:ID"`
}
