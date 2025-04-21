package services

import (
	"context"

	models "github.com/edorguez/payment-reminder/internal/account/models"
	"github.com/edorguez/payment-reminder/internal/account/repository"
	"github.com/edorguez/payment-reminder/pkg/constants"
	"github.com/edorguez/payment-reminder/pkg/kafka"
	"github.com/edorguez/payment-reminder/pkg/kafka/events"
)

type UserService interface {
	Create(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id uint) (*models.User, error)
	FindByEmail(ctx context.Context, email string) *models.User
	Update(ctx context.Context, id uint, newUser *models.User) error
	Delete(ctx context.Context, id uint) error
}

type userService struct {
	repo     repository.UserRepository
	producer *kafka.Producer
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Create(ctx context.Context, user *models.User) error {
	user, err := s.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	event := events.UserEvent{
		EventType: constants.UserCreatedEvent,
		UserID:    user.ID,
		Email:     user.Email,
	}

	return s.producer.SendEvent(event)
}

func (s *userService) FindByID(ctx context.Context, id uint) (*models.User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *userService) FindByEmail(ctx context.Context, email string) *models.User {
	return s.repo.FindByEmail(ctx, email)
}

func (s *userService) Update(ctx context.Context, id uint, newUser *models.User) error {
	return s.repo.Update(ctx, id, newUser)
}

func (s *userService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
