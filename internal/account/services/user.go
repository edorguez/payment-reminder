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
	FindByID(ctx context.Context, id int64) (*models.User, error)
	FindByEmail(ctx context.Context, email string) *models.User
	Update(ctx context.Context, id int64, newUser *models.User) error
	Delete(ctx context.Context, id int64) error
}

type userService struct {
	repo     repository.UserRepository
	producer *kafka.Producer
}

func NewUserService(repo repository.UserRepository, producer *kafka.Producer) UserService {
	return &userService{
		repo:     repo,
		producer: producer,
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

func (s *userService) FindByID(ctx context.Context, id int64) (*models.User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *userService) FindByEmail(ctx context.Context, email string) *models.User {
	return s.repo.FindByEmail(ctx, email)
}

func (s *userService) Update(ctx context.Context, id int64, newUser *models.User) error {
	err := s.repo.Update(ctx, id, newUser)
	if err != nil {
		return err
	}

	event := events.UserEvent{
		EventType: constants.UserUpdatedEvent,
		UserID:    id,
		Email:     newUser.Email,
	}

	return s.producer.SendEvent(event)
}

func (s *userService) Delete(ctx context.Context, id int64) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	event := events.UserEvent{
		EventType: constants.UserDeletedEvent,
		UserID:    id,
	}

	return s.producer.SendEvent(event)
}
