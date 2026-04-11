package repository

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	infra "debian-ecommerce/internal/infra/transaction"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, customer *entity.Customer) error
	FindCustomerByUserID(ctx context.Context, userID uint) (*entity.Customer, error)
	UpdateCustomer(ctx context.Context, customer *entity.Customer) error
}

type customerRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewCustomerRepository(db *gorm.DB, log *zap.Logger) CustomerRepository {
	return &customerRepository{
		DB:  db,
		Log: log,
	}
}

func (r *customerRepository) CreateCustomer(ctx context.Context, customer *entity.Customer) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Create(customer).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}

func (r *customerRepository) FindCustomerByUserID(ctx context.Context, userID uint) (*entity.Customer, error) {
	db := infra.GetDB(ctx, r.DB)
	var customer entity.Customer
	if err := db.WithContext(ctx).Where("user_id = ?", userID).First(&customer).Preload("User").Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) UpdateCustomer(ctx context.Context, customer *entity.Customer) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Save(customer).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}
