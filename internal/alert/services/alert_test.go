package services

import (
	"context"
	"testing"
	"time"

	"github.com/edorguez/payment-reminder/internal/alert/models"
	mocks "github.com/edorguez/payment-reminder/mocks/alert/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateAlert_Failure(t *testing.T) {
	// Arrange
	alertRepo := new(mocks.MockAlertRepository)
	alertTemplateRepo := new(mocks.MockAlertTemplateRepository)
	userCacheRepo := new(mocks.MockUserCacheRepository)

	service := NewAlertService(alertRepo, alertTemplateRepo, userCacheRepo)

	alert := &models.Alert{
		Name:            "Test Alert",
		Description:     "Test Description",
		UserID:          1,
		AlertTemplateID: 1,
		PhoneNumber:     "1234567",
		HourConcurrence: nil,
		StartAt:         time.Now(),
		IsActive:        true,
	}

	// Act
	service.On("FetchData", 789).
		Return("raw_data", nil).
		Once()

	mockDB.On("StoreData", 789, "PROCESSED_raw_data").
		Return(assert.AnError).
		Once()

	result, err := processor.Process(789)

	mockDB.AssertExpectations(t)
	assert.Error(t, err)
	assert.Empty(t, result)
	err := service.Create(context.Background(), alert)

	// Assert
	assert.Nil(t, err)
}

// func (suite *AlertServiceTestSuite) TestCreateAlert_Error(t *testing.T) {
// 	alert := &models.Alert{Name: "Test Alert"}
// 	expectedErr := errors.New("repository error")
//
// 	suite.repo.On("Create", suite.ctx, alert).Return(expectedErr)
//
// 	err := suite.service.Create(suite.ctx, alert)
// 	assert.Error(suite.T(), err)
// 	assert.Equal(suite.T(), expectedErr, err)
// }
//
// func (suite *AlertServiceTestSuite) TestFindByID_Success() {
// 	testID := int64(1)
// 	expectedAlert := &models.Alert{
// 		ID:   int64(testID),
// 		Name: "Existing Alert",
// 	}
//
// 	suite.repo.On("FindByID", suite.ctx, testID).Return(expectedAlert, nil)
//
// 	result, err := suite.service.FindByID(suite.ctx, testID)
// 	assert.NoError(suite.T(), err)
// 	assert.Equal(suite.T(), expectedAlert, result)
// }
//
// func (suite *AlertServiceTestSuite) TestFindByID_NotFound() {
// 	testID := int64(999)
// 	expectedErr := errors.New("not found")
//
// 	suite.repo.On("FindByID", suite.ctx, testID).Return(&models.Alert{}, expectedErr)
//
// 	_, err := suite.service.FindByID(suite.ctx, testID)
// 	assert.Error(suite.T(), err)
// 	assert.Equal(suite.T(), expectedErr, err)
// }
//
// func (suite *AlertServiceTestSuite) TestUpdateAlert_Success() {
// 	testID := int64(1)
// 	newAlert := &models.Alert{
// 		Name:        "Updated Alert",
// 		Description: "Updated Description",
// 	}
//
// 	suite.repo.On("Update", suite.ctx, testID, newAlert).Return(nil)
//
// 	err := suite.service.Update(suite.ctx, testID, newAlert)
// 	assert.NoError(suite.T(), err)
// 	suite.repo.AssertExpectations(suite.T())
// }
//
// func (suite *AlertServiceTestSuite) TestUpdateAlert_Error() {
// 	testID := int64(1)
// 	newAlert := &models.Alert{Name: "Updated Alert"}
// 	expectedErr := errors.New("update failed")
//
// 	suite.repo.On("Update", suite.ctx, testID, newAlert).Return(expectedErr)
//
// 	err := suite.service.Update(suite.ctx, testID, newAlert)
// 	assert.Error(suite.T(), err)
// 	assert.Equal(suite.T(), expectedErr, err)
// }
//
// func (suite *AlertServiceTestSuite) TestDeleteAlert_Success() {
// 	testID := uint(1)
//
// 	suite.repo.On("Delete", suite.ctx, testID).Return(nil)
//
// 	err := suite.service.Delete(suite.ctx, testID)
// 	assert.NoError(suite.T(), err)
// 	suite.repo.AssertExpectations(suite.T())
// }
//
// func (suite *AlertServiceTestSuite) TestDeleteAlert_Error() {
// 	testID := int64(999)
// 	expectedErr := errors.New("delete failed")
//
// 	suite.repo.On("Delete", suite.ctx, testID).Return(expectedErr)
//
// 	err := suite.service.Delete(suite.ctx, testID)
// 	assert.Error(suite.T(), err)
// 	assert.Equal(suite.T(), expectedErr, err)
// }
//
// func TestAlertServiceTestSuite(t *testing.T) {
// 	suite.Run(t, new(AlertServiceTestSuite))
// }
