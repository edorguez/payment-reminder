package models

import "time"

type AlertTemplate struct {
	ID          int64     `gorm:"primaryKey;type:bigserial"`
	Name        string    `gorm:"type:varchar(50);not null"`
	Description string    `gorm:"type:varchar(200);not null"`
	ContentSID  string    `gorm:"type:varchar(34);not null"`
	IsActive    bool      `gorm:"type:boolean;not null"`
	CreatedAt   time.Time `gorm:"type:timestamptz;not null;default:now()"`
	ModifiedAt  time.Time `gorm:"type:timestamptz;not null;default:now()"`
}
