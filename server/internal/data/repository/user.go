package repository

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	infra "debian-ecommerce/internal/infra/transaction"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
	FindUserByID(ctx context.Context, id uint) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id uint) error
}

type userRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewUserRepository(db *gorm.DB, log *zap.Logger) UserRepository {
	return &userRepository{
		DB:  db,
		Log: log,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Create(user).Error; err != nil {
		r.Log.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	db := infra.GetDB(ctx, r.DB)
	var user entity.User
	if err := db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindUserByID(ctx context.Context, id uint) (*entity.User, error) {
	db := infra.GetDB(ctx, r.DB)
	var user entity.User
	if err := db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Save(user).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id uint) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Delete(&entity.User{}, id).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}
