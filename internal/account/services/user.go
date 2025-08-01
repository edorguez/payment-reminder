package services

import (
	"time"

	models "github.com/edorguez/payment-reminder/internal/account/models"
	"github.com/edorguez/payment-reminder/internal/account/repository"
	"github.com/edorguez/payment-reminder/pkg/constants"
	"github.com/edorguez/payment-reminder/pkg/core/errors"
	"github.com/edorguez/payment-reminder/pkg/kafka"
	"github.com/edorguez/payment-reminder/pkg/kafka/events"
	"github.com/edorguez/payment-reminder/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	Create(ctx *gin.Context) error
	FindByID(ctx *gin.Context, id int64) (*models.User, error)
	FindByFirebaseID(ctx *gin.Context, id string) (*models.User, error)
	FindByEmail(ctx *gin.Context, email string) (*models.User, error)
	Update(ctx *gin.Context, id int64, newUser *models.User) error
	Delete(ctx *gin.Context, id int64) error
	// VerifyToken(ctx context.Context, token string) (string, error)
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

func (s *userService) Create(ctx *gin.Context) error {
	claims, ok := middleware.ExtractClaims(ctx)
	if !ok {
		return &errors.Error{Err: errors.ErrGeneral, Message: "Claims not found"}
	}

	_, err := s.FindByEmail(ctx, claims.Email)
	if err == nil {
		return &errors.Error{Err: errors.ErrInvalidInput, Message: "Error creating user, email is already in use"}
	}

	_, err = s.FindByFirebaseID(ctx, claims.FirbaseUID)
	if err == nil {
		return &errors.Error{Err: errors.ErrInvalidInput, Message: "Error creating user, firebase UID already exists"}
	}

	u := &models.User{
		FirebaseUID:     claims.FirbaseUID,
		UserPlanID:      constants.UserPlanBasic,
		Name:            claims.Name,
		Email:           claims.Email,
		LastPaymentDate: time.Now().UTC(),
		LastLoginDate:   time.Now().UTC(),
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

func (s *userService) FindByID(ctx *gin.Context, id int64) (*models.User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *userService) FindByFirebaseID(ctx *gin.Context, id string) (*models.User, error) {
	return s.repo.FindByFirebaseID(ctx, id)
}

func (s *userService) FindByEmail(ctx *gin.Context, email string) (*models.User, error) {
	return s.repo.FindByEmail(ctx, email)
}

func (s *userService) Update(ctx *gin.Context, id int64, newUser *models.User) error {
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

func (s *userService) Delete(ctx *gin.Context, id int64) error {
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

// func (s *userService) VerifyToken(ctx context.Context, token string) (string, error) {
// 	t, err := s.firebase.VerifyIDToken(ctx, token)
// 	if err != nil {
// 		return "", &errors.Error{Err: errors.ErrFirebase, Message: err.Error()}
// 	}
// 	return t.UID, nil
// }
