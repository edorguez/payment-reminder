package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type UserCache struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
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
		return nil, fmt.Errorf("failed to marshal user: %w", err)
	}

	if err := r.redis.Set(ctx, strconv.FormatInt(user.ID, 10), data, 0).Err(); err != nil {
		return nil, fmt.Errorf("failed to store user: %w", err)
	}

	return &user.ID, nil
}

func (r *userCacheRepository) FindByID(ctx context.Context, id int64) (*UserCache, error) {
	data, err := r.redis.Get(ctx, strconv.FormatInt(id, 10)).Bytes()
	if err == redis.Nil {
		return nil, fmt.Errorf("User not found")
	}

	if err != nil {
		return nil, fmt.Errorf("Failed to get user: %w", err)
	}

	var user UserCache
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal user: %w", err)
	}

	return &user, nil
}

func (r *userCacheRepository) Update(ctx context.Context, user UserCache) error {
	count, err := r.redis.Exists(ctx, strconv.FormatInt(user.ID, 10)).Result()
	if err != nil {
		return fmt.Errorf("Failed to check user existence: %w", err)
	}
	if count <= 0 {
		return fmt.Errorf("User not found")
	}

	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("Failed to marshal user: %w", err)
	}

	if err := r.redis.Set(ctx, strconv.FormatInt(user.ID, 10), data, 0).Err(); err != nil {
		return fmt.Errorf("Failed to update user: %w", err)
	}

	return nil
}

func (r *userCacheRepository) Delete(ctx context.Context, id int64) error {
	if err := r.redis.Del(ctx, strconv.FormatInt(id, 10)).Err(); err != nil {
		return fmt.Errorf("Failed to delete user: %w", err)
	}
	return nil
}
