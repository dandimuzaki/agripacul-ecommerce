package repository

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	infra "debian-ecommerce/internal/infra/transaction"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type WishlistRepository interface {
	Create(ctx context.Context, wishlist *entity.Wishlist) error
	FindByCustomerID(ctx context.Context, customerID uint) ([]entity.Wishlist, error)
	Delete(ctx context.Context, customerID uint, productID uint) error
	FindOne(ctx context.Context, customerID, productID uint) (*entity.Wishlist, error)
}

type wishlistRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewWishlistRepository(db *gorm.DB, log *zap.Logger) WishlistRepository {
	return &wishlistRepository{
		DB:  db,
		Log: log,
	}
}

func (r *wishlistRepository) Create(ctx context.Context, wishlist *entity.Wishlist) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Create(wishlist).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}

func (r *wishlistRepository) FindByCustomerID(ctx context.Context, customerID uint) ([]entity.Wishlist, error) {
	db := infra.GetDB(ctx, r.DB)
	var wishlists []entity.Wishlist
	if err := db.WithContext(ctx).Preload("Product").Where("customer_id = ?", customerID).Find(&wishlists).Error; err != nil {
		r.Log.Error(err.Error())
		return nil, err
	}
	return wishlists, nil
}

func (r *wishlistRepository) FindOne(ctx context.Context, customerID, productID uint) (*entity.Wishlist, error) {
	db := infra.GetDB(ctx, r.DB)
	var wishlist entity.Wishlist
	if err := db.WithContext(ctx).Where("customer_id = ? AND product_id = ?", customerID, productID).First(&wishlist).Error; err != nil {
		return nil, err
	}
	return &wishlist, nil
}

func (r *wishlistRepository) Delete(ctx context.Context, customerID uint, productID uint) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Where("customer_id = ? AND product_id = ?", customerID, productID).Delete(&entity.Wishlist{}).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}
