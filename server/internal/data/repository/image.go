package repository

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	infra "debian-ecommerce/internal/infra/transaction"
	"debian-ecommerce/pkg/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ImageRepository interface{
	GetByID(ctx context.Context, id uint) (*entity.Image, error)
	Create(ctx context.Context, product *entity.Image) error
	Update(ctx context.Context, id uint, product *entity.Image) error
	Delete(ctx context.Context, id uint) error
}

type imageRepository struct {
	db *gorm.DB
	log *zap.Logger
}

func NewImageRepo(db *gorm.DB, log *zap.Logger) ImageRepository {
	return &imageRepository{
		db: db,
		log: log,
	}
}

func (r *imageRepository) GetByID(ctx context.Context, id uint) (*entity.Image, error) {
	db := infra.GetDB(ctx, r.db)
	var img entity.Image
	err := db.Model(&img).First(&img, id).Error
	if err != nil {
		r.log.Error("Failed to get image",
			zap.Uint("image_id", id),
			zap.Error(err))
		return nil, err
	}
	return &img, nil
}

func (r *imageRepository) Create(ctx context.Context, image *entity.Image) error {
	db := infra.GetDB(ctx, r.db)
	r.log.Info("Creating image",
		zap.Uint("product_id", image.ProductID),
		zap.String("image_url", image.ImageURL),
	)

	err := db.Create(image).Error
	if err != nil {
		r.log.Error("Failed to create image",
			zap.String("image_url", image.ImageURL),
			zap.Error(err))
		return err
	}

	r.log.Info("image created successfully",
		zap.Uint("id", image.ID),
		zap.String("image_url", image.ImageURL))

	return nil
}

func (r *imageRepository) Update(ctx context.Context, id uint, image *entity.Image) error {
	db := infra.GetDB(ctx, r.db)

	result := db.Model(&entity.Image{}).
		Where("id = ?", id).
		Updates(image)

	if result.Error != nil {
		r.log.Error("Error query update image", zap.Error(result.Error))
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrImageNotFound
	}

	return nil
}

func (r *imageRepository) Delete(ctx context.Context, id uint) error {
	db := infra.GetDB(ctx, r.db)
	result := db.Delete(&entity.Image{}, id)
	if result.Error != nil {
		r.log.Error("Error query delete image", zap.Error(result.Error))
		return result.Error
	}
	if result.RowsAffected == 0 {
		return utils.ErrImageNotFound
	}
	return nil
}