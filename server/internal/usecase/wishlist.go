package usecase

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/data/repository"
	"errors"

	"go.uber.org/zap"
)

type WishlistUsecase interface {
	AddWishlist(ctx context.Context, userID uint, productID uint) error
	GetWishlist(ctx context.Context, userID uint) ([]entity.Wishlist, error)
	RemoveWishlist(ctx context.Context, userID uint, productID uint) error
}

type wishlistUsecase struct {
	wishlistRepo repository.WishlistRepository
	customerRepo repository.CustomerRepository
	productRepo  repository.ProductRepository
	log          *zap.Logger
}

func NewWishlistUsecase(wishlistRepo repository.WishlistRepository, customerRepo repository.CustomerRepository, productRepo repository.ProductRepository, log *zap.Logger) WishlistUsecase {
	return &wishlistUsecase{
		wishlistRepo: wishlistRepo,
		customerRepo: customerRepo,
		productRepo:  productRepo,
		log:          log,
	}
}

func (u *wishlistUsecase) AddWishlist(ctx context.Context, userID uint, productID uint) error {
	// Find customer by user ID
	customer, err := u.customerRepo.FindCustomerByUserID(ctx, userID)
	if err != nil {
		u.log.Error("Failed to find customer", zap.Error(err))
		return err
	}

	// Check if product exists
	_, err = u.productRepo.GetByID(ctx, productID)
	if err != nil {
		u.log.Error("Failed to find product", zap.Error(err))
		return err
	}

	// Check if already in wishlist
	existing, err := u.wishlistRepo.FindOne(ctx, customer.ID, productID)
	if err == nil && existing != nil {
		return errors.New("product already in wishlist")
	}

	wishlist := &entity.Wishlist{
		CustomerID: customer.ID,
		ProductID:  productID,
	}

	if err := u.wishlistRepo.Create(ctx, wishlist); err != nil {
		u.log.Error("Failed to create wishlist", zap.Error(err))
		return err
	}

	return nil
}

func (u *wishlistUsecase) GetWishlist(ctx context.Context, userID uint) ([]entity.Wishlist, error) {
	customer, err := u.customerRepo.FindCustomerByUserID(ctx, userID)
	if err != nil {
		u.log.Error("Failed to find customer", zap.Error(err))
		return nil, err
	}

	wishlists, err := u.wishlistRepo.FindByCustomerID(ctx, customer.ID)
	if err != nil {
		u.log.Error("Failed to get wishlists", zap.Error(err))
		return nil, err
	}

	return wishlists, nil
}

func (u *wishlistUsecase) RemoveWishlist(ctx context.Context, userID uint, productID uint) error {
	customer, err := u.customerRepo.FindCustomerByUserID(ctx, userID)
	if err != nil {
		u.log.Error("Failed to find customer", zap.Error(err))
		return err
	}

	if err := u.wishlistRepo.Delete(ctx, customer.ID, productID); err != nil {
		u.log.Error("Failed to delete wishlist", zap.Error(err))
		return err
	}

	return nil
}
