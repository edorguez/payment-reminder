package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/edorguez/payment-reminder/internal/account/models"
	"github.com/edorguez/payment-reminder/internal/account/services"
	customerrors "github.com/edorguez/payment-reminder/pkg/core/errors"
	"github.com/edorguez/payment-reminder/pkg/utils"
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

func (h *UserHandler) Create(ctx *gin.Context) {
	err := h.service.Create(ctx)
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

	ctx.Status(http.StatusCreated)
}

func (h *UserHandler) FindById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	user, err := h.service.FindByID(ctx, id)
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

	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) ListOrFind(ctx *gin.Context) {
	email := ctx.Query("email")
	firebaseId := ctx.Query("firebaseId")

	switch {
	case email != "" && firebaseId != "":
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "only one filter allowed"})
		return
	case email != "":
		user, err := h.service.FindByEmail(ctx, email)
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
		ctx.JSON(http.StatusOK, user)
	case firebaseId != "":
		user, err := h.service.FindByFirebaseID(ctx, firebaseId)
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

		ctx.JSON(http.StatusOK, user)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "either email or firebaseId query param is required"})
		return
	}
}

func (h *UserHandler) FindByFirebaseId(ctx *gin.Context) {
	id := ctx.Query("firebaseId")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Firebase ID query param is required"})
		return
	}

}

func (h *UserHandler) Update(ctx *gin.Context) {
	type updateUserRequest struct {
		UserPlanID int64  `json:"user_plan_id" binding:"required"`
		Email      string `json:"email" binding:"required"`
	}
	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idParam := ctx.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	user := models.User{
		UserPlanID: req.UserPlanID,
		Email:      req.Email,
	}

	err := h.service.Update(ctx, id, &user)
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

func (h *UserHandler) Delete(ctx *gin.Context) {
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
