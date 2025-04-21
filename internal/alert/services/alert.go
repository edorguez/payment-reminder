package services

import (
	"context"
	models "github.com/edorguez/payment-reminder/internal/alert/models"
	"github.com/edorguez/payment-reminder/internal/alert/repository"
)

type AlertService interface {
	Create(ctx context.Context, alert *models.Alert) error
	FindByID(ctx context.Context, id uint) (*models.Alert, error)
	Update(ctx context.Context, id uint, newUser *models.Alert) error
	Delete(ctx context.Context, id uint) error
}

type alertService struct {
	repo          repository.AlertRepository
	userCacheRepo repository.UserCacheRepository
}

func NewAlertService(repo repository.AlertRepository, userCacheRepo repository.UserCacheRepository) AlertService {
	return &alertService{
		repo:          repo,
		userCacheRepo: userCacheRepo,
	}
}

func (s *alertService) Create(ctx context.Context, alert *models.Alert) error {
	_, err := s.repo.Create(ctx, alert)
	return err
}

func (s *alertService) FindByID(ctx context.Context, id uint) (*models.Alert, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *alertService) Update(ctx context.Context, id uint, newAlert *models.Alert) error {
	return s.repo.Update(ctx, id, newAlert)
}

func (s *alertService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
