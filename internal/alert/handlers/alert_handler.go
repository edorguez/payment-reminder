package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/edorguez/payment-reminder/internal/alert/models"
	"github.com/edorguez/payment-reminder/internal/alert/services"
	customerrors "github.com/edorguez/payment-reminder/pkg/core/errors"
	"github.com/edorguez/payment-reminder/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AlertHandler struct {
	service services.AlertService
}

func NewAlertHandler(service services.AlertService) *AlertHandler {
	return &AlertHandler{
		service: service,
	}
}

func (h *AlertHandler) Create(ctx *gin.Context) {
	type createAlertRequest struct {
		UserID          int64     `json:"user_id" binding:"required"`
		AlertTemplateID int64     `json:"alert_template_id" binding:"required"`
		Name            string    `json:"name" binding:"required"`
		Description     string    `json:"description" binding:"required"`
		PhoneNumber     string    `json:"phone_number" binding:"required"`
		HourConcurrence uint16    `json:"hour_concurrence" binding:"required"`
		StartAt         time.Time `json:"start_at" biding:"required"`
		IsActive        bool      `json:"is_active" binding:"required"`
	}

	var req createAlertRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	alert := models.Alert{
		UserID:          req.UserID,
		AlertTemplateID: req.AlertTemplateID,
		Name:            req.Name,
		Description:     req.Description,
		PhoneNumber:     req.PhoneNumber,
		HourConcurrence: req.HourConcurrence,
		StartAt:         req.StartAt,
		IsActive:        req.IsActive,
	}

	err := h.service.Create(ctx, &alert)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *AlertHandler) FindById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	alert, err := h.service.FindByID(ctx, id)
	if err != nil {
		var customErr *customerrors.Error
		if errors.As(err, &customErr) {
			status := utils.MapCodeToHTTPStatus(customErr.Err)
			ctx.JSON(status, gin.H{"error": customErr.Message})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, alert)
}

func (h *AlertHandler) Update(ctx *gin.Context) {
	type updateAlertRequest struct {
		AlertTemplateID int64     `json:"alert_template_id" binding:"required"`
		Name            string    `json:"name" binding:"required"`
		Description     string    `json:"description" binding:"required"`
		PhoneNumber     string    `json:"phone_number" binding:"required"`
		HourConcurrence uint16    `json:"hour_concurrence" binding:"required"`
		StartAt         time.Time `json:"start_at" biding:"required"`
		IsActive        bool      `json:"is_active" binding:"required"`
	}

	var req updateAlertRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idParam := ctx.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	alert := models.Alert{
		AlertTemplateID: req.AlertTemplateID,
		Name:            req.Name,
		Description:     req.Description,
		PhoneNumber:     req.PhoneNumber,
		HourConcurrence: req.HourConcurrence,
		StartAt:         req.StartAt,
		IsActive:        req.IsActive,
	}

	err := h.service.Update(ctx, id, &alert)
	if err != nil {
		var customErr *customerrors.Error
		if errors.As(err, &customErr) {
			status := utils.MapCodeToHTTPStatus(customErr.Err)
			ctx.JSON(status, gin.H{"error": customErr.Message})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *AlertHandler) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	err := h.service.Delete(ctx, id)
	if err != nil {
		var customErr *customerrors.Error
		if errors.As(err, &customErr) {
			status := utils.MapCodeToHTTPStatus(customErr.Err)
			ctx.JSON(status, gin.H{"error": customErr.Message})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
