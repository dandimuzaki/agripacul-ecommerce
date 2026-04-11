package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type TokenRepository interface {
	SaveToken(ctx context.Context, userID uint, token string, exp time.Duration) error
	ValidateToken(ctx context.Context, userID uint, token string) error
	DeleteToken(ctx context.Context, userID uint) error
}

type tokenRepository struct {
	Redis *redis.Client
}

func NewTokenRepository(redis *redis.Client) TokenRepository {
	return &tokenRepository{
		Redis: redis,
	}
}

func (r *tokenRepository) SaveToken(ctx context.Context, userID uint, token string, exp time.Duration) error {
	key := fmt.Sprintf("auth:%d", userID)
	return r.Redis.Set(ctx, key, token, exp).Err()
}

func (r *tokenRepository) ValidateToken(ctx context.Context, userID uint, token string) error {
	key := fmt.Sprintf("auth:%d", userID)
	val, err := r.Redis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return fmt.Errorf("session expired or invalid")
		}
		return err
	}
	if val != token {
		return fmt.Errorf("invalid token")
	}
	return nil
}

func (r *tokenRepository) DeleteToken(ctx context.Context, userID uint) error {
	key := fmt.Sprintf("auth:%d", userID)
	return r.Redis.Del(ctx, key).Err()
}
