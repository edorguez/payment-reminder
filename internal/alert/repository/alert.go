package repository

import (
	"time"

	models "github.com/edorguez/payment-reminder/internal/alert/models"
	"github.com/edorguez/payment-reminder/pkg/core/errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AlertRepository interface {
	Create(ctx *gin.Context, alert *models.Alert) (*models.Alert, error)
	FindByID(ctx *gin.Context, id int64) (*models.Alert, error)
	Update(ctx *gin.Context, id int64, newAlert *models.Alert) error
	Delete(ctx *gin.Context, id int64) error
}

type alertRepository struct {
	DB *gorm.DB
}

func NewAlertRepository(DB *gorm.DB) AlertRepository {
	return &alertRepository{
		DB: DB,
	}
}

func (r *alertRepository) Create(ctx *gin.Context, alert *models.Alert) (*models.Alert, error) {
	createdAlert := r.DB.Create(alert)
	if createdAlert.Error != nil {
		return nil, &errors.Error{Err: errors.ErrGeneral, Message: createdAlert.Error.Error()}
	}

	return alert, nil
}

func (r *alertRepository) FindByID(ctx *gin.Context, id int64) (*models.Alert, error) {
	var alert models.Alert

	r.DB.First(&alert, id)
	if alert.ID == 0 {
		return nil, &errors.Error{Err: errors.ErrNotFound, Message: "Alert not found"}
	}

	return &alert, nil
}

func (r *alertRepository) Update(ctx *gin.Context, id int64, newAlert *models.Alert) error {
	var oldAlert models.Alert

	r.DB.First(&oldAlert, id)

	if oldAlert.ID == 0 {
		return &errors.Error{Err: errors.ErrNotFound, Message: "Alert not found"}
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

func (r *alertRepository) Delete(ctx *gin.Context, id int64) error {
	var alert models.Alert

	r.DB.First(&alert, id)

	if alert.ID == 0 {
		return &errors.Error{Err: errors.ErrNotFound, Message: "Alert not found"}
	}

	r.DB.Unscoped().Delete(&alert)

	return nil
}
