package usecase

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/data/repository"
	"debian-ecommerce/pkg/utils"
)

func GetUserFromContext(ctx context.Context, repo *repository.Repository) (*entity.User, error) {
	userID, ok := ctx.Value("user_id").(uint)
	if !ok {
		return nil, utils.ErrInvalidUserID
	}

	user, err := repo.UserRepo.FindUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetCustomerFromContext(ctx context.Context, repo *repository.Repository) (*entity.Customer, error) {
	userID, ok := ctx.Value("user_id").(uint)
	if !ok {
		return nil, utils.ErrInvalidUserID
	}

	customer, err := repo.CustomerRepo.FindCustomerByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return customer, nil
}