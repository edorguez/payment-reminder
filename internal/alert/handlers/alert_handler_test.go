package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/dvln/testify/suite"
	"github.com/edorguez/payment-reminder/internal/alert/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAlertService implements AlertService interface for testing
type MockAlertService struct {
	mock.Mock
}

func (m *MockAlertService) Create(ctx context.Context, alert *models.Alert) error {
	args := m.Called(ctx, alert)
	return args.Error(0)
}

func (m *MockAlertService) FindByID(ctx context.Context, id uint) (*models.Alert, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Alert), args.Error(1)
}

func (m *MockAlertService) Update(ctx context.Context, id uint, newAlert *models.Alert) error {
	args := m.Called(ctx, id, newAlert)
	return args.Error(0)
}

func (m *MockAlertService) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type AlertHandlerTestSuite struct {
	suite.Suite
	service *MockAlertService
	handler *AlertHandler
	router  *gin.Engine
}

func (suite *AlertHandlerTestSuite) SetupTest() {
	suite.service = new(MockAlertService)
	suite.handler = NewAlertHandler(suite.service)

	gin.SetMode(gin.TestMode)
	suite.router = gin.New()

	// Setup routes
	suite.router.POST("/alerts", suite.handler.Create)
	suite.router.GET("/alerts/:id", suite.handler.FindById)
	suite.router.PUT("/alerts/:id", suite.handler.Update)
	suite.router.DELETE("/alerts/:id", suite.handler.Delete)
}

func TestAlertHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(AlertHandlerTestSuite))
}

func (suite *AlertHandlerTestSuite) TestCreateAlert_Success() {
	alert := models.Alert{
		UserID:          1,
		AlertTemplateID: 1,
		Name:            "Test Alert",
		Description:     "Test Description",
		PhoneNumber:     "+1234567890",
		HourConcurrence: 12,
		StartAt:         time.Now(),
		IsActive:        true,
	}

	suite.service.On("Create", mock.Anything, &alert).Return(nil)

	body, _ := json.Marshal(alert)
	req, _ := http.NewRequest("POST", "/alerts", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusNoContent, w.Code)
	suite.service.AssertExpectations(suite.T())
}

func (suite *AlertHandlerTestSuite) TestCreateAlert_BadRequest() {
	invalidAlert := map[string]interface{}{
		"user_id": "invalid", // Should be integer
	}

	body, _ := json.Marshal(invalidAlert)
	req, _ := http.NewRequest("POST", "/alerts", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *AlertHandlerTestSuite) TestFindByID_Success() {
	testID := uint(1)
	expectedAlert := &models.Alert{
		ID:   int64(testID),
		Name: "Test Alert",
	}

	suite.service.On("FindByID", mock.Anything, testID).Return(expectedAlert, nil)

	req, _ := http.NewRequest("GET", "/alerts/"+strconv.FormatUint(uint64(testID), 10), nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)

	var response models.Alert
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(suite.T(), expectedAlert.ID, response.ID)
}

func (suite *AlertHandlerTestSuite) TestFindByID_NotFound() {
	testID := uint(999)
	suite.service.On("FindByID", mock.Anything, testID).Return(&models.Alert{}, errors.New("not found"))

	req, _ := http.NewRequest("GET", "/alerts/999", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusNotFound, w.Code)
}

func (suite *AlertHandlerTestSuite) TestUpdateAlert_Success() {
	testID := uint(1)
	alert := models.Alert{
		AlertTemplateID: 2,
		Name:            "Updated Alert",
	}

	suite.service.On("Update", mock.Anything, testID, &alert).Return(nil)

	body, _ := json.Marshal(alert)
	req, _ := http.NewRequest("PUT", "/alerts/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusNoContent, w.Code)
}

func (suite *AlertHandlerTestSuite) TestUpdateAlert_InvalidID() {
	req, _ := http.NewRequest("PUT", "/alerts/invalid", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *AlertHandlerTestSuite) TestDeleteAlert_Success() {
	testID := uint(1)
	suite.service.On("Delete", mock.Anything, testID).Return(nil)

	req, _ := http.NewRequest("DELETE", "/alerts/1", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusNoContent, w.Code)
}

func (suite *AlertHandlerTestSuite) TestDeleteAlert_ServiceError() {
	testID := uint(999)
	suite.service.On("Delete", mock.Anything, testID).Return(errors.New("database error"))

	req, _ := http.NewRequest("DELETE", "/alerts/999", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}
