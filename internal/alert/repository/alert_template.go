package repository

import (
	"github.com/edorguez/payment-reminder/internal/alert/models"
	"github.com/edorguez/payment-reminder/pkg/core/errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AlertTemplateRepository interface {
	FindByID(ctx *gin.Context, id int64) (*models.AlertTemplate, error)
}

type alertTemplateRepository struct {
	DB *gorm.DB
}

func NewAlertTemplateRepository(DB *gorm.DB) AlertTemplateRepository {
	return &alertTemplateRepository{
		DB: DB,
	}
}

func (r *alertTemplateRepository) FindByID(ctx *gin.Context, id int64) (*models.AlertTemplate, error) {
	var template models.AlertTemplate

	r.DB.First(&template, id)
	if template.ID == 0 {
		return nil, &errors.Error{Err: errors.ErrNotFound, Message: "Alert template not found"}
	}

	return &template, nil
}
