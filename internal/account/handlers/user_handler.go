package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/edorguez/payment-reminder/internal/account/models"
	"github.com/edorguez/payment-reminder/internal/account/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) FindById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 32)

	user, err := h.service.FindByID(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) FindByEmail(ctx *gin.Context) {
	email := ctx.Query("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email query param is required"})
		return
	}

	user := h.service.FindByEmail(ctx, email)
	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) Update(ctx *gin.Context) {
	type updateUserRequest struct {
		UserPlanID      int64     `json:"user_plan_id" binding:"required"`
		Email           string    `json:"email" binding:"required"`
		LastPaymentDate time.Time `json:"last_payment_date" binding:"required"`
	}
	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idParam := ctx.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 32)

	user := models.User{
		UserPlanID:      req.UserPlanID,
		Email:           req.Email,
		LastPaymentDate: req.LastPaymentDate,
	}

	err := h.service.Update(ctx, uint(id), &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *UserHandler) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 32)

	err := h.service.Delete(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
