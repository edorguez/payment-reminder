package services

import (
	"context"
	"time"

	"firebase.google.com/go/v4/auth"
	models "github.com/edorguez/payment-reminder/internal/account/models"
	"github.com/edorguez/payment-reminder/internal/account/repository"
	"github.com/edorguez/payment-reminder/pkg/constants"
	"github.com/edorguez/payment-reminder/pkg/core/errors"
	"github.com/edorguez/payment-reminder/pkg/kafka"
	"github.com/edorguez/payment-reminder/pkg/kafka/events"
)

type UserService interface {
	Create(ctx context.Context, email string, password string, userPlanId int64) error
	FindByID(ctx context.Context, id int64) (*models.User, error)
	FindByEmail(ctx context.Context, email string) *models.User
	Update(ctx context.Context, id int64, newUser *models.User) error
	Delete(ctx context.Context, id int64) error
}

type userService struct {
	repo     repository.UserRepository
	firebase *auth.Client
	producer *kafka.Producer
}

func NewUserService(repo repository.UserRepository, firebase *auth.Client, producer *kafka.Producer) UserService {
	return &userService{
		repo:     repo,
		firebase: firebase,
		producer: producer,
	}
}

func (s *userService) Create(ctx context.Context, email string, password string, userPlanId int64) error {
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password).
		EmailVerified(false).
		Disabled(false)

	firebaseRes, err := s.firebase.CreateUser(context.Background(), params)
	if err != nil {
		return &errors.Error{Err: errors.ErrFirebase, Message: err.Error()}
	}

	u := &models.User{
		FirebaseUID:     firebaseRes.UID,
		UserPlanID:      userPlanId,
		Email:           email,
		LastPaymentDate: time.Now().UTC(),
	}

	user, err := s.repo.Create(ctx, u)
	if err != nil {
		return err
	}

	event := events.UserEvent{
		EventType: constants.UserCreatedEvent,
		UserID:    user.ID,
		Email:     user.Email,
	}

	if err = s.producer.SendEvent(event); err != nil {
		return &errors.Error{Err: errors.ErrPublishEvent, Message: err.Error()}
	}

	return nil
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

	if err = s.producer.SendEvent(event); err != nil {
		return &errors.Error{Err: errors.ErrPublishEvent, Message: err.Error()}
	}

	return nil
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

	if err = s.producer.SendEvent(event); err != nil {
		return &errors.Error{Err: errors.ErrPublishEvent, Message: err.Error()}
	}

	return nil
}
