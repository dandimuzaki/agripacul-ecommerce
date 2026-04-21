package repository

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	infra "debian-ecommerce/internal/infra/transaction"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PaymentMethodRepository interface {
	Create(ctx context.Context, paymentMethod *entity.PaymentMethod) error
	Update(ctx context.Context, paymentMethod *entity.PaymentMethod) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*entity.PaymentMethod, error)
	FindAll(ctx context.Context, page, limit int, search string, isActive *bool) ([]entity.PaymentType, int64, error)
	CheckNameExists(ctx context.Context, name string, excludeID uint) (bool, error)
	GetAll(ctx context.Context) ([]entity.PaymentType, error)
}

type paymentMethodRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewPaymentMethodRepository(db *gorm.DB, log *zap.Logger) PaymentMethodRepository {
	return &paymentMethodRepository{
		DB:  db,
		Log: log,
	}
}

// Create inserts a new payment method
func (r *paymentMethodRepository) Create(ctx context.Context, paymentMethod *entity.PaymentMethod) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Create(paymentMethod).Error; err != nil {
		r.Log.Error("failed to create payment method", zap.Error(err))
		return err
	}
	return nil
}

// Update updates an existing payment method
func (r *paymentMethodRepository) Update(ctx context.Context, paymentMethod *entity.PaymentMethod) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Save(paymentMethod).Error; err != nil {
		r.Log.Error("failed to update payment method",
			zap.Error(err),
			zap.Uint("id", paymentMethod.ID),
		)
		return err
	}
	return nil
}

// Delete soft deletes a payment method
func (r *paymentMethodRepository) Delete(ctx context.Context, id uint) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Delete(&entity.PaymentMethod{}, id).Error; err != nil {
		r.Log.Error("failed to delete payment method",
			zap.Error(err),
			zap.Uint("id", id),
		)
		return err
	}
	return nil
}

// FindByID finds payment method by ID
func (r *paymentMethodRepository) FindByID(ctx context.Context, id uint) (*entity.PaymentMethod, error) {
	db := infra.GetDB(ctx, r.DB)
	var paymentMethod entity.PaymentMethod
	if err := db.WithContext(ctx).First(&paymentMethod, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Return nil, nil for not found (sesuai pola EmployeeRepo)
		}
		r.Log.Error("failed to find payment method by ID",
			zap.Error(err),
			zap.Uint("id", id),
		)
		return nil, err
	}
	return &paymentMethod, nil
}

// FindAll retrieves all payment methods with pagination and filters
func (r *paymentMethodRepository) FindAll(ctx context.Context, page, limit int, search string, isActive *bool) ([]entity.PaymentType, int64, error) {
	var paymentTypes []entity.PaymentType
	var total int64

	db := infra.GetDB(ctx, r.DB).Model(&entity.PaymentType{}).Joins("JOIN payment_methods pm ON payment_types.id = pm.payment_type_id")

	// Apply search filter
	if search != "" {
		searchTerm := "%" + search + "%"
		db = db.Where("LOWER(pm.name) LIKE LOWER(?)", searchTerm)
	}

	// Apply is_active filter
	if isActive != nil {
		db = db.Where("pm.is_active = ?", *isActive)
	}

	// Get total count
	if err := db.WithContext(ctx).Count(&total).Error; err != nil {
		r.Log.Error("failed to count payment methods", zap.Error(err))
		return nil, 0, err
	}

	// Apply pagination and order
	offset := (page - 1) * limit
	if err := db.WithContext(ctx).
		Model(&entity.PaymentType{}).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&paymentTypes).
		Preload("Methods").Error; err != nil {
		r.Log.Error("failed to find payment methods", zap.Error(err))
		return nil, 0, err
	}

	return paymentTypes, total, nil
}

// GetAll retrieves all payment methods
func (r *paymentMethodRepository) GetAll(ctx context.Context) ([]entity.PaymentType, error) {
	var paymentTypes []entity.PaymentType
	db := infra.GetDB(ctx, r.DB).Model(&entity.PaymentType{})

	if err := db.
		Preload("Methods").
		Order("name ASC").
		Find(&paymentTypes).Error; err != nil {
		r.Log.Error("failed to find payment methods", zap.Error(err))
		return nil, err
	}

	return paymentTypes, nil
}

// CheckNameExists checks if payment method name already exists
func (r *paymentMethodRepository) CheckNameExists(ctx context.Context, name string, excludeID uint) (bool, error) {
	db := infra.GetDB(ctx, r.DB).Model(&entity.PaymentMethod{}).Where("name = ?", name)

	if excludeID > 0 {
		db = db.Where("id != ?", excludeID)
	}

	var count int64
	if err := db.WithContext(ctx).Count(&count).Error; err != nil {
		r.Log.Error("failed to check payment method name existence",
			zap.Error(err),
			zap.String("name", name),
		)
		return false, err
	}

	return count > 0, nil
}
