package repository

import (
	"context"
	"time"

	models "github.com/edorguez/payment-reminder/internal/account/models"
	"github.com/edorguez/payment-reminder/pkg/core/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	FindByID(ctx context.Context, id int64) (*models.User, error)
	FindByFirebaseID(ctx context.Context, id string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) *models.User
	Update(ctx context.Context, id int64, newUser *models.User) error
	Delete(ctx context.Context, id int64) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	createdUser := r.DB.Create(user)
	if createdUser.Error != nil {
		return nil, &errors.Error{Err: errors.ErrGeneral, Message: createdUser.Error.Error()}
	}

	return user, nil
}

func (r *userRepository) FindByID(ctx context.Context, id int64) (*models.User, error) {
	var user models.User

	r.DB.First(&user, id)
	if user.ID == 0 {
		return nil, &errors.Error{Err: errors.ErrNotFound, Message: "User not found"}
	}

	return &user, nil
}

func (r *userRepository) FindByFirebaseID(ctx context.Context, id string) (*models.User, error) {
	var users []models.User

	r.DB.Where("firebaseUid = ?", id).Find(&users)
	if len(users) == 0 {
		return nil, &errors.Error{Err: errors.ErrNotFound, Message: "User not found"}
	}

	return &users[0], nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) *models.User {
	var users []models.User

	r.DB.Where("email LIKE ?", "%"+email+"%").Find(&users)
	if len(users) == 0 {
		return nil
	}

	return &users[0]
}

func (r *userRepository) Update(ctx context.Context, id int64, newUser *models.User) error {
	var oldUser models.User

	r.DB.First(&oldUser, id)

	if oldUser.ID == 0 {
		return &errors.Error{Err: errors.ErrNotFound, Message: "User not found"}
	}

	oldUser.UserPlanID = newUser.UserPlanID
	oldUser.Email = newUser.Email
	oldUser.LastPaymentDate = newUser.LastPaymentDate
	oldUser.ModifiedAt = time.Now()

	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	var user models.User

	r.DB.First(&user, id)

	if user.ID == 0 {
		return &errors.Error{Err: errors.ErrNotFound, Message: "User not found"}
	}

	r.DB.Unscoped().Delete(&user)

	return nil
}
