package repository

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	infra "debian-ecommerce/internal/infra/transaction"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BannerRepository interface {
	CreateBanner(ctx context.Context, banner *entity.Banner) error
	GetBannerList(ctx context.Context) ([]entity.Banner, error)
	GetBannerByID(ctx context.Context, id uint) (*entity.Banner, error)
	UpdateBanner(ctx context.Context, banner *entity.Banner) error
	DeleteBanner(ctx context.Context, id uint) error
}

type bannerRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewBannerRepository(db *gorm.DB, log *zap.Logger) BannerRepository {
	return &bannerRepository{
		DB:  db,
		Log: log,
	}
}

func (r *bannerRepository) CreateBanner(ctx context.Context, banner *entity.Banner) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Create(banner).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}

func (r *bannerRepository) GetBannerList(ctx context.Context) ([]entity.Banner, error) {
	db := infra.GetDB(ctx, r.DB)
	var banners []entity.Banner
	if err := db.WithContext(ctx).Find(&banners).Error; err != nil {
		r.Log.Error(err.Error())
		return nil, err
	}
	return banners, nil
}

func (r *bannerRepository) GetBannerByID(ctx context.Context, id uint) (*entity.Banner, error) {
	db := infra.GetDB(ctx, r.DB)
	var banner entity.Banner
	if err := db.WithContext(ctx).First(&banner, id).Error; err != nil {
		r.Log.Error(err.Error())
		return nil, err
	}
	return &banner, nil
}

func (r *bannerRepository) UpdateBanner(ctx context.Context, banner *entity.Banner) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Save(banner).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}

func (r *bannerRepository) DeleteBanner(ctx context.Context, id uint) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Delete(&entity.Banner{}, id).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}
