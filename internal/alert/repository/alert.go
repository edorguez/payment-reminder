package repository

import (
	"context"
	"fmt"
	"time"

	models "github.com/edorguez/payment-reminder/internal/alert/models"
	"gorm.io/gorm"
)

type AlertRepository interface {
	Create(ctx context.Context, alert *models.Alert) error
	FindByID(ctx context.Context, id uint) (*models.Alert, error)
	Update(ctx context.Context, id uint, newAlert *models.Alert) error
	Delete(ctx context.Context, id uint) error
}

type alertRepository struct {
	DB *gorm.DB
}

func NewAlertRepository(DB *gorm.DB) AlertRepository {
	return &alertRepository{
		DB: DB,
	}
}

func (r *alertRepository) Create(ctx context.Context, alert *models.Alert) error {
	createdAlert := r.DB.Create(alert)
	return createdAlert.Error
}

func (r *alertRepository) FindByID(ctx context.Context, id uint) (*models.Alert, error) {
	var alert models.Alert

	r.DB.First(&alert, id)
	if alert.ID == 0 {
		return nil, fmt.Errorf("Alert not found")
	}

	return &alert, nil
}

func (r *alertRepository) Update(ctx context.Context, id uint, newAlert *models.Alert) error {
	var oldAlert models.Alert

	r.DB.First(&oldAlert, id)

	if oldAlert.ID == 0 {
		return fmt.Errorf("Alert not found")
	}

	oldAlert.AlertTemplateID = newAlert.AlertTemplateID
	oldAlert.Name = newAlert.Name
	oldAlert.Description = newAlert.Description
	oldAlert.PhoneNumber = newAlert.PhoneNumber
	oldAlert.HourConcurrence = newAlert.HourConcurrence
	oldAlert.StartAt = newAlert.StartAt
	oldAlert.IsActive = newAlert.IsActive
	oldAlert.ModifiedAt = time.Now()

	return nil
}

func (r *alertRepository) Delete(ctx context.Context, id uint) error {
	var alert models.Alert

	r.DB.First(&alert, id)

	if alert.ID == 0 {
		return fmt.Errorf("Alert not found")
	}

	r.DB.Unscoped().Delete(&alert)

	return nil
}
