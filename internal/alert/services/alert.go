package services

import (
	"context"
	models "github.com/edorguez/payment-reminder/internal/alert/models"
	"github.com/edorguez/payment-reminder/internal/alert/repository"
)

type AlertService interface {
	Create(ctx context.Context, alert *models.Alert) error
	FindByID(ctx context.Context, id int64) (*models.Alert, error)
	Update(ctx context.Context, id int64, newUser *models.Alert) error
	Delete(ctx context.Context, id int64) error
}

type alertService struct {
	repo              repository.AlertRepository
	alertTemplateRepo repository.AlertTemplateRepository
	userCacheRepo     repository.UserCacheRepository
}

func NewAlertService(repo repository.AlertRepository, alertTemplateRepo repository.AlertTemplateRepository, userCacheRepo repository.UserCacheRepository) AlertService {
	return &alertService{
		repo:              repo,
		alertTemplateRepo: alertTemplateRepo,
		userCacheRepo:     userCacheRepo,
	}
}

func (s *alertService) Create(ctx context.Context, alert *models.Alert) error {
	_, err := s.userCacheRepo.FindByID(ctx, alert.UserID)
	if err != nil {
		return err
	}

	_, err = s.alertTemplateRepo.FindByID(ctx, alert.AlertTemplateID)
	if err != nil {
		return err
	}

	_, err = s.repo.Create(ctx, alert)

	return err
}

func (s *alertService) FindByID(ctx context.Context, id int64) (*models.Alert, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *alertService) Update(ctx context.Context, id int64, newAlert *models.Alert) error {
	_, err := s.userCacheRepo.FindByID(ctx, newAlert.UserID)
	if err != nil {
		return err
	}

	_, err = s.alertTemplateRepo.FindByID(ctx, newAlert.AlertTemplateID)
	if err != nil {
		return err
	}

	return s.repo.Update(ctx, id, newAlert)
}

func (s *alertService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
