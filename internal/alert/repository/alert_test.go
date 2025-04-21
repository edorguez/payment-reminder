package repository

import (
	"context"
	"testing"
	"time"

	"github.com/edorguez/payment-reminder/internal/alert/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AlertRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	repo AlertRepository
}

func (suite *AlertRepositoryTestSuite) SetupSuite() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		suite.T().Fatal("Failed to connect to database:", err)
	}

	// AutoMigrate the Alert model
	if err := db.AutoMigrate(&models.Alert{}); err != nil {
		suite.T().Fatal("Failed to migrate database:", err)
	}

	suite.db = db
}

func (suite *AlertRepositoryTestSuite) SetupTest() {
	// Start a new transaction for each test
	tx := suite.db.Begin()
	suite.repo = NewAlertRepository(tx)
}

func (suite *AlertRepositoryTestSuite) TearDownTest() {
	// Rollback the transaction after each test
	suite.db.Rollback()
}

func (suite *AlertRepositoryTestSuite) TestCreateAlert() {
	ctx := context.Background()
	alert := &models.Alert{
		AlertTemplateID: 1,
		Name:            "Test Alert",
		Description:     "Test Description",
		PhoneNumber:     "+1234567890",
		HourConcurrence: 12,
		StartAt:         time.Now(),
		IsActive:        true,
	}

	_, err := suite.repo.Create(ctx, alert)
	assert.NoError(suite.T(), err)
	assert.NotZero(suite.T(), alert.ID, "Alert ID should not be zero after creation")

	// Verify the alert can be retrieved
	foundAlert, err := suite.repo.FindByID(ctx, uint(alert.ID))
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), alert.Name, foundAlert.Name)
}

func (suite *AlertRepositoryTestSuite) TestFindByID_Exists() {
	ctx := context.Background()
	alert := &models.Alert{Name: "Existing Alert"}
	_, err := suite.repo.Create(ctx, alert)
	assert.NoError(suite.T(), err)

	foundAlert, err := suite.repo.FindByID(ctx, uint(alert.ID))
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), alert.ID, foundAlert.ID)
}

func (suite *AlertRepositoryTestSuite) TestFindByID_NotExists() {
	ctx := context.Background()
	_, err := suite.repo.FindByID(ctx, 999)
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "Alert not found")
}

func (suite *AlertRepositoryTestSuite) TestUpdateAlert_Success() {
	ctx := context.Background()
	originalAlert := &models.Alert{
		Name:        "Original Name",
		Description: "Original Description",
		IsActive:    false,
		StartAt:     time.Now(),
	}
	_, err := suite.repo.Create(ctx, originalAlert)
	assert.NoError(suite.T(), err)

	newAlertData := &models.Alert{
		Name:        "Updated Name",
		Description: "Updated Description",
		IsActive:    true,
		StartAt:     originalAlert.StartAt.Add(1 * time.Hour),
	}

	err = suite.repo.Update(ctx, uint(originalAlert.ID), newAlertData)
	assert.NoError(suite.T(), err)

	updatedAlert, err := suite.repo.FindByID(ctx, uint(originalAlert.ID))
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), newAlertData.Name, updatedAlert.Name)
	assert.Equal(suite.T(), newAlertData.Description, updatedAlert.Description)
	assert.Equal(suite.T(), newAlertData.IsActive, updatedAlert.IsActive)
	assert.True(suite.T(), updatedAlert.StartAt.Equal(newAlertData.StartAt))
	assert.True(suite.T(), updatedAlert.ModifiedAt.After(originalAlert.ModifiedAt))
}

func (suite *AlertRepositoryTestSuite) TestUpdateAlert_NotFound() {
	ctx := context.Background()
	newAlert := &models.Alert{Name: "Non-existent Alert"}
	err := suite.repo.Update(ctx, 999, newAlert)
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "Alert not found")
}

func (suite *AlertRepositoryTestSuite) TestDeleteAlert_Success() {
	ctx := context.Background()
	alert := &models.Alert{Name: "Alert to Delete"}
	_, err := suite.repo.Create(ctx, alert)
	assert.NoError(suite.T(), err)

	err = suite.repo.Delete(ctx, uint(alert.ID))
	assert.NoError(suite.T(), err)

	_, err = suite.repo.FindByID(ctx, uint(alert.ID))
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "Alert not found")
}

func (suite *AlertRepositoryTestSuite) TestDeleteAlert_NotFound() {
	ctx := context.Background()
	err := suite.repo.Delete(ctx, 999)
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "Alert not found")
}

func TestAlertRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AlertRepositoryTestSuite))
}
