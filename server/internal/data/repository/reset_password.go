package repository

import (
	"context"
	"debian-ecommerce/pkg/utils"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type ResetPasswordRepository interface {
	SaveOTP(ctx context.Context, userID uint, hashedOTP string, exp time.Duration) error
	ValidateOTP(ctx context.Context, userID uint, inputOTP string) error
	DeleteOTP(ctx context.Context, userID uint) error
	IncrementAttempt(ctx context.Context, userID uint) (int64, error)
}

type resetPasswordRepository struct {
	Redis *redis.Client
	log *zap.Logger
}

func NewResetPasswordRepository(redis *redis.Client, log *zap.Logger) ResetPasswordRepository {
	return &resetPasswordRepository{Redis: redis}
}

func (r *resetPasswordRepository) SaveOTP(
	ctx context.Context,
	userID uint,
	hashedOTP string,
	exp time.Duration,
) error {
	key := fmt.Sprintf("reset_pwd:otp:%d", userID)
	return r.Redis.Set(ctx, key, hashedOTP, exp).Err()
}

func (r *resetPasswordRepository) ValidateOTP(
	ctx context.Context,
	userID uint,
	inputOTP string,
) error {
	key := fmt.Sprintf("reset_pwd:otp:%d", userID)

	hashedOTP, err := r.Redis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return fmt.Errorf("otp expired or invalid")
		}
		return err
	}

	if !utils.CheckPassword(inputOTP, hashedOTP) {
		return fmt.Errorf("invalid otp")
	}

	return nil
}

func (r *resetPasswordRepository) IncrementAttempt(
	ctx context.Context,
	userID uint,
) (int64, error) {
	key := fmt.Sprintf("reset_pwd:attempt:%d", userID)

	attempts, err := r.Redis.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	// Ensure TTL exists (same as OTP TTL)
	r.Redis.Expire(ctx, key, 10*time.Minute)

	return attempts, nil
}

func (r *resetPasswordRepository) DeleteOTP(
	ctx context.Context,
	userID uint,
) error {
	otpKey := fmt.Sprintf("reset_pwd:otp:%d", userID)
	attemptKey := fmt.Sprintf("reset_pwd:attempt:%d", userID)

	return r.Redis.Del(ctx, otpKey, attemptKey).Err()
}
