package usecase

import (
	"context"
	"debian-ecommerce/internal/data/repository"
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/dto/response"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CustomerUsecase interface {
	UpdateProfile(ctx context.Context, userID uint, req request.UpdateProfileRequest) error
	GetCustomerByUserID(ctx context.Context, userID uint) (*response.CustomerProfile, error)
}

type customerUsecase struct {
	repo repository.CustomerRepository
	cloudinary ImageUploader
	log  *zap.Logger
}

func NewCustomerUsecase(repo repository.CustomerRepository, cloudinary ImageUploader, log *zap.Logger) CustomerUsecase {
	return &customerUsecase{
		repo: repo,
		cloudinary: cloudinary,
		log:  log,
	}
}

func (u *customerUsecase) UpdateProfile(ctx context.Context, userID uint, req request.UpdateProfileRequest) error {
	customer, err := u.repo.FindCustomerByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("customer profile not found")
		}
		return err
	}

	if req.ProfileImage != nil {
		// Get old public id to delete
		oldPublicID := customer.ProfileImagePublicID

		// Upload image to cloudinary
		url, publicID, err := u.cloudinary.Upload(ctx, req.ProfileImage, "debian/categories")
		if err != nil {
			u.log.Error("Failed to upload image to cloudinary", zap.Error(err))
			return err
		}

		// Delete old image
		err = u.cloudinary.Delete(ctx, oldPublicID)
		if err != nil {
			u.log.Error("Failed to delete old image", zap.Error(err))
			return err
		}

		customer.ProfileImageURL = url
		customer.ProfileImagePublicID = publicID
	}

	if req.FullName != "" {
		customer.FullName = req.FullName
	}
	if req.PhoneNumber != "" {
		customer.PhoneNumber = req.PhoneNumber
	}
	if req.DateOfBirth != "" {
		dateOfBirth, err := time.Parse("2006-01-02", req.DateOfBirth)
		if err != nil {
			return err
		}
		customer.DateOfBirth = &dateOfBirth
	}

	err = u.repo.UpdateCustomer(ctx, customer)
	if err != nil {
		u.log.Error("Failed to update customer", zap.Error(err))
		return err
	}

	return nil
}

func (u *customerUsecase) GetCustomerByUserID(ctx context.Context, userID uint) (*response.CustomerProfile, error) {
	cust, err := u.repo.FindCustomerByUserID(ctx, userID)
	if err != nil {
		u.log.Error("Failed to get customer profile", zap.Error(err))
	}
	return response.ToCustomerProfile(cust), nil
}
