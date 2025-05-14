package utils

import (
	"errors"
	customerrors "github.com/edorguez/payment-reminder/pkg/core/errors"
	"net/http"
)

func MapCodeToHTTPStatus(err error) int {
	switch {
	case errors.Is(err, customerrors.ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, customerrors.ErrInvalidInput):
		return http.StatusBadRequest
	case errors.Is(err, customerrors.ErrConflict):
		return http.StatusConflict
	case errors.Is(err, customerrors.ErrPublishEvent):
		return http.StatusInternalServerError
	case errors.Is(err, customerrors.ErrConsumeEvent):
		return http.StatusInternalServerError
	case errors.Is(err, customerrors.ErrFirebase):
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
