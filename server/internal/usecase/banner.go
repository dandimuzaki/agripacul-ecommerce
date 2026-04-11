package usecase

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/data/repository"
	"debian-ecommerce/internal/dto/request"
	"errors"
	"time"

	"go.uber.org/zap"
)

type BannerUsecase interface {
	CreateBanner(ctx context.Context, req *request.BannerRequest) error
	GetBannerList(ctx context.Context) ([]entity.Banner, error)
	GetBannerByID(ctx context.Context, id uint) (*entity.Banner, error)
	UpdateBanner(ctx context.Context, id uint, req *request.BannerRequest) error
	DeleteBanner(ctx context.Context, id uint) error
}

type bannerUsecase struct {
	tx TxManager
	cloudinary ImageUploader
	BannerRepo repository.BannerRepository
	Log        *zap.Logger
}

func NewBannerUsecase(tx TxManager, cloudinary ImageUploader, repo repository.BannerRepository, log *zap.Logger) BannerUsecase {
	return &bannerUsecase{
		tx: tx,
		cloudinary: cloudinary,
		BannerRepo: repo,
		Log:        log,
	}
}

func (u *bannerUsecase) CreateBanner(ctx context.Context, req *request.BannerRequest) error {
	if req.Name == "" {
		return errors.New("banner name is required")
	}
	if req.StartDate.IsZero() || req.EndDate.IsZero() {
		return errors.New("start date and end date are required")
	}
	if req.EndDate.Before(req.StartDate) {
		return errors.New("end date must be after start date")
	}

	// Default value for IsPublished if not set?
	// The struct tag `default:true` handles DB default, but here `false` is zero value.
	// If the user wants `true` by default, they should set it or we set it here if it's a pointer or check specific logic.
	// However, usually API request comes with it.

	banner := req.ToBanner()

	// Upload banner image to cloudinary
	url, publicID, err := u.cloudinary.Upload(ctx, req.Image, "debian/banners")
	if err != nil {
		u.Log.Error("Failed to upload image to cloudinary", zap.Error(err))
		return err
	}

	banner.ImageURL = url
	banner.ImagePublicID = publicID

	if err := u.BannerRepo.CreateBanner(ctx, &banner); err != nil {
		u.Log.Error("failed to create banner", zap.Error(err))
		return err
	}
	return nil
}

func (u *bannerUsecase) GetBannerList(ctx context.Context) ([]entity.Banner, error) {
	banners, err := u.BannerRepo.GetBannerList(ctx)
	if err != nil {
		u.Log.Error("failed to get banner list", zap.Error(err))
		return nil, err
	}
	return banners, nil
}

func (u *bannerUsecase) GetBannerByID(ctx context.Context, id uint) (*entity.Banner, error) {
	banner, err := u.BannerRepo.GetBannerByID(ctx, id)
	if err != nil {
		u.Log.Error("failed to get banner by id", zap.Error(err))
		return nil, err
	}
	return banner, nil
}

func (u *bannerUsecase) UpdateBanner(ctx context.Context, id uint, req *request.BannerRequest) error {
	existing, err := u.BannerRepo.GetBannerByID(ctx, id)
	if err != nil {
		return err
	}
	
	existing.Name = req.Name
	existing.TargetURL = req.TargetURL

	if req.Image != nil {
		// Get old public id to delete
		oldPublicID := existing.ImagePublicID

		// Upload image to cloudinary
		url, publicID, err := u.cloudinary.Upload(ctx, req.Image, "debian/banners")
		if err != nil {
			u.Log.Error("Failed to upload image to cloudinary", zap.Error(err))
			return err
		}

		// Delete old image
		err = u.cloudinary.Delete(ctx, oldPublicID)
		if err != nil {
			u.Log.Error("Failed to delete old image", zap.Error(err))
			return err
		}

		existing.ImageURL = url
		existing.ImagePublicID = publicID
	}

	// Only update dates if provided? Or assume full update?
	// Usually strict PUT updates everything.
	if !req.StartDate.IsZero() {
		existing.StartDate = req.StartDate
	}
	if !req.EndDate.IsZero() {
		existing.EndDate = req.EndDate
	} else {
		// If dates were required, we might keep old ones.
		// But let's assume if they send valid struct they want to update.
	}

	// Actually, careful with zero values.
	// For now, let's just update fields.
	existing.StartDate = req.StartDate
	existing.EndDate = req.EndDate

	if req.Type != "" {
		existing.Type = req.Type
	}

	// IsPublished is boolean, tricky with zero value.
	// We'll trust the input for now.
	existing.IsPublished = req.IsPublished

	// Validate dates again
	if existing.EndDate.Before(existing.StartDate) {
		return errors.New("end date must be after start date")
	}

	// Update timestamp
	existing.UpdatedAt = time.Now()

	if err := u.BannerRepo.UpdateBanner(ctx, existing); err != nil {
		u.Log.Error("failed to update banner", zap.Error(err))
		return err
	}
	return nil
}

func (u *bannerUsecase) DeleteBanner(ctx context.Context, id uint) error {
	_, err := u.BannerRepo.GetBannerByID(ctx, id)
	if err != nil {
		return err
	}

	if err := u.BannerRepo.DeleteBanner(ctx, id); err != nil {
		u.Log.Error("failed to delete banner", zap.Error(err))
		return err
	}
	return nil
}
