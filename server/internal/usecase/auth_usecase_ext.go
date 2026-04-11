package usecase

import (
	"context"
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/pkg/utils"
	template "debian-ecommerce/pkg/utils/email_template"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (u *authUsecase) IsEmailAvailable(ctx context.Context, email string) (bool, error) {
	_, err := u.UserRepo.FindUserByEmail(ctx, email)
	if err == nil {
		return false, nil // Email exists
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true, nil // Email does not exist
	}
	return false, err
}

func (u *authUsecase) RequestResetPassword(
	ctx context.Context,
	req request.ForgotPasswordRequest,
) error {
	user, err := u.UserRepo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		// Prevent email enumeration
		return nil
	}

	// Generate OTP
	otp, err := utils.GenerateOTP(6)
	if err != nil {
		u.Log.Error("Failed to generate OTP", zap.Error(err))
		return err
	}

	// Hash OTP
	hashedOTP := utils.HashPassword(otp)

	// Save to Redis (TTL 10 min)
	if err := u.ResetPasswordRepo.SaveOTP(
		ctx,
		user.ID,
		string(hashedOTP),
		10*time.Minute,
	); err != nil {
		return err
	}

	// Send Email
	subject := "Password Reset Request for Your Account"
	body := template.RequestResetPassword(user.Email, otp, "Debian E-commerce")
	if err := u.EmailService.SendEmail(user.Email, subject, body); err != nil {
		u.Log.Error("failed to send welcome email", zap.Error(err), zap.String("email", user.Email))
		return nil
	}

	return nil
}


func (u *authUsecase) ResetPassword(
	ctx context.Context,
	req request.ResetPasswordRequest,
) error {

	// Get user by email
	user, err := u.UserRepo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		u.Log.Error("Failed to get user by email", zap.Error(err))
		return err
	}

	// Check attempts
	attempts, _ := u.ResetPasswordRepo.IncrementAttempt(ctx, user.ID)
	if attempts > 5 {
		u.Log.Warn("Too many attempts", zap.String("email", req.Email))
		return fmt.Errorf("too many attempts")
	}

	// Validate OTP
	if err := u.ResetPasswordRepo.ValidateOTP(ctx, user.ID, req.OTP); err != nil {
		u.Log.Error("Failed to generate OTP", zap.Error(err))
		return err
	}

	// Hash new password
	hashedPwd := utils.HashPassword(req.NewPassword)

	// Update password in DB
	user.PasswordHash = string(hashedPwd)
	user.PasswordChangedAt = time.Now()
	if err := u.UserRepo.UpdateUser(ctx, user); err != nil {
		return err
	}

	// Cleanup
	return u.ResetPasswordRepo.DeleteOTP(ctx, user.ID)
}
