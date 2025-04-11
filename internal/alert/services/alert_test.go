package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/dvln/testify/assert"
	"github.com/edorguez/payment-reminder/internal/alert/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// MockAlertRepository implements AlertRepository interface for testing
type MockAlertRepository struct {
	mock.Mock
}

func (m *MockAlertRepository) Create(ctx context.Context, alert *models.Alert) error {
	args := m.Called(ctx, alert)
	return args.Error(0)
}

func (m *MockAlertRepository) FindByID(ctx context.Context, id uint) (*models.Alert, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Alert), args.Error(1)
}

func (m *MockAlertRepository) Update(ctx context.Context, id uint, newAlert *models.Alert) error {
	args := m.Called(ctx, id, newAlert)
	return args.Error(0)
}

func (m *MockAlertRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type AlertServiceTestSuite struct {
	suite.Suite
	repo    *MockAlertRepository
	service AlertService
	ctx     context.Context
}

func (suite *AlertServiceTestSuite) SetupTest() {
	suite.repo = new(MockAlertRepository)
	suite.service = NewAlertService(suite.repo)
	suite.ctx = context.Background()
}

func (suite *AlertServiceTestSuite) TestCreateAlert_Success() {
	alert := &models.Alert{
		Name:        "Test Alert",
		Description: "Test Description",
		StartAt:     time.Now(),
	}

	suite.repo.On("Create", suite.ctx, alert).Return(nil)

	err := suite.service.Create(suite.ctx, alert)
	assert.NoError(suite.T(), err)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *AlertServiceTestSuite) TestCreateAlert_Error() {
	alert := &models.Alert{Name: "Test Alert"}
	expectedErr := errors.New("repository error")

	suite.repo.On("Create", suite.ctx, alert).Return(expectedErr)

	err := suite.service.Create(suite.ctx, alert)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedErr, err)
}

func (suite *AlertServiceTestSuite) TestFindByID_Success() {
	testID := uint(1)
	expectedAlert := &models.Alert{
		ID:   int64(testID),
		Name: "Existing Alert",
	}

	suite.repo.On("FindByID", suite.ctx, testID).Return(expectedAlert, nil)

	result, err := suite.service.FindByID(suite.ctx, testID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedAlert, result)
}

func (suite *AlertServiceTestSuite) TestFindByID_NotFound() {
	testID := uint(999)
	expectedErr := errors.New("not found")

	suite.repo.On("FindByID", suite.ctx, testID).Return(&models.Alert{}, expectedErr)

	_, err := suite.service.FindByID(suite.ctx, testID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedErr, err)
}

func (suite *AlertServiceTestSuite) TestUpdateAlert_Success() {
	testID := uint(1)
	newAlert := &models.Alert{
		Name:        "Updated Alert",
		Description: "Updated Description",
	}

	suite.repo.On("Update", suite.ctx, testID, newAlert).Return(nil)

	err := suite.service.Update(suite.ctx, testID, newAlert)
	assert.NoError(suite.T(), err)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *AlertServiceTestSuite) TestUpdateAlert_Error() {
	testID := uint(1)
	newAlert := &models.Alert{Name: "Updated Alert"}
	expectedErr := errors.New("update failed")

	suite.repo.On("Update", suite.ctx, testID, newAlert).Return(expectedErr)

	err := suite.service.Update(suite.ctx, testID, newAlert)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedErr, err)
}

func (suite *AlertServiceTestSuite) TestDeleteAlert_Success() {
	testID := uint(1)

	suite.repo.On("Delete", suite.ctx, testID).Return(nil)

	err := suite.service.Delete(suite.ctx, testID)
	assert.NoError(suite.T(), err)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *AlertServiceTestSuite) TestDeleteAlert_Error() {
	testID := uint(999)
	expectedErr := errors.New("delete failed")

	suite.repo.On("Delete", suite.ctx, testID).Return(expectedErr)

	err := suite.service.Delete(suite.ctx, testID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedErr, err)
}

func TestAlertServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AlertServiceTestSuite))
}
