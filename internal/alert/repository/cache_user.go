package repository

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/edorguez/payment-reminder/pkg/core/errors"
	"github.com/redis/go-redis/v9"
)

type UserCache struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

type UserCacheRepository interface {
	Create(ctx context.Context, user UserCache) (*int64, error)
	FindByID(ctx context.Context, id int64) (*UserCache, error)
	Update(ctx context.Context, user UserCache) error
	Delete(ctx context.Context, id int64) error
}

type userCacheRepository struct {
	redis *redis.Client
}

func NewUserCacheRepository(redis *redis.Client) UserCacheRepository {
	return &userCacheRepository{
		redis: redis,
	}
}

func (r *userCacheRepository) Create(ctx context.Context, user UserCache) (*int64, error) {
	data, err := json.Marshal(user)
	if err != nil {
		return nil, &errors.Error{Err: errors.ErrGeneral, Message: "Failed to marshal user"}
	}

	if err := r.redis.Set(ctx, strconv.FormatInt(user.ID, 10), data, 0).Err(); err != nil {
		return nil, &errors.Error{Err: errors.ErrGeneral, Message: "Failed to store user"}
	}

	return &user.ID, nil
}

func (r *userCacheRepository) FindByID(ctx context.Context, id int64) (*UserCache, error) {
	data, err := r.redis.Get(ctx, strconv.FormatInt(id, 10)).Bytes()
	if err == redis.Nil {
		return nil, &errors.Error{Err: errors.ErrNotFound, Message: "User not found"}
	}

	if err != nil {
		return nil, &errors.Error{Err: errors.ErrGeneral, Message: "Failed to get user"}
	}

	var user UserCache
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, &errors.Error{Err: errors.ErrGeneral, Message: "Failed to unmarshal user"}
	}

	return &user, nil
}

func (r *userCacheRepository) Update(ctx context.Context, user UserCache) error {
	count, err := r.redis.Exists(ctx, strconv.FormatInt(user.ID, 10)).Result()
	if err != nil {
		return &errors.Error{Err: errors.ErrGeneral, Message: "Failed to check user existence"}
	}
	if count <= 0 {
		return &errors.Error{Err: errors.ErrNotFound, Message: "User not found"}
	}

	data, err := json.Marshal(user)
	if err != nil {
		return &errors.Error{Err: errors.ErrGeneral, Message: "Failed to marshal user"}
	}

	if err := r.redis.Set(ctx, strconv.FormatInt(user.ID, 10), data, 0).Err(); err != nil {
		return &errors.Error{Err: errors.ErrGeneral, Message: "Failed to update user"}
	}

	return nil
}

func (r *userCacheRepository) Delete(ctx context.Context, id int64) error {
	if err := r.redis.Del(ctx, strconv.FormatInt(id, 10)).Err(); err != nil {
		return &errors.Error{Err: errors.ErrGeneral, Message: "Failed to delete user"}
	}
	return nil
}
