package repository

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/dto/request"
	infra "debian-ecommerce/internal/infra/transaction"
	"debian-ecommerce/pkg/utils"
)

type PromotionRepository interface {
	Create(ctx context.Context, promotion *entity.Promotion) error
	GetByID(ctx context.Context, id uint) (*entity.Promotion, error)
	GetList(ctx context.Context, filter request.PromotionFilterQuery) ([]entity.Promotion, int64, error) // UPDATED
	Update(ctx context.Context, promotion *entity.Promotion) error
	Delete(ctx context.Context, id uint) error
	CheckVoucherCodeExists(ctx context.Context, code string, excludeID ...uint) (bool, error)
	GetActivePromotionsByProduct(ctx context.Context, productID uint) ([]entity.Promotion, error)
	LockPromoByID(ctx context.Context, promoID uint) (*entity.Promotion, error)
	LockPromoByCode(ctx context.Context, code string) (*entity.Promotion, error)
	CreatePromoUsage(ctx context.Context, promo *entity.PromoUsage) error
}

type promotionRepository struct {
	db *gorm.DB
	log *zap.Logger
}

func NewPromotionRepository(db *gorm.DB, log *zap.Logger) PromotionRepository {
	return &promotionRepository{db: db, log: log}
}

func (r *promotionRepository) Create(ctx context.Context, promotion *entity.Promotion) error {
	db := infra.GetDB(ctx, r.db)
	return db.Create(promotion).Error
}

func (r *promotionRepository) GetByID(ctx context.Context, id uint) (*entity.Promotion, error) {
	db := infra.GetDB(ctx, r.db)
	var promotion entity.Promotion
	err := db.Preload("PromoProducts.Product").First(&promotion, id).Error
	if err != nil {
		return nil, err
	}
	return &promotion, nil
}

func (r *promotionRepository) GetList(ctx context.Context, filter request.PromotionFilterQuery) ([]entity.Promotion, int64, error) {
	db := infra.GetDB(ctx, r.db)
	var promotions []entity.Promotion
	var total int64

	query := db.Model(&entity.Promotion{})

	// Apply filters
	if filter.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name+"%")
	}
	if filter.Type != "" {
		query = query.Where("type = ?", filter.Type)
	}
	if filter.IsPublished {
		query = query.Where("is_published = ?", filter.IsPublished)
	}
	if !filter.StartDateFrom.IsZero() {
		query = query.Where("start_date >= ?", filter.StartDateFrom)
	}
	if !filter.StartDateTo.IsZero() {
		query = query.Where("start_date <= ?", filter.StartDateTo)
	}
	if !filter.EndDateFrom.IsZero() {
		query = query.Where("end_date >= ?", filter.EndDateFrom)
	}
	if !filter.EndDateTo.IsZero() {
		query = query.Where("end_date <= ?", filter.EndDateTo)
	}
	if filter.MinimumOrderValue != 0 {
		query = query.Where("minimum_order_value < ?", filter.MinimumOrderValue)
	}
	if filter.Available {
		query = query.Where("used_count < usage_limit")
	}
	if filter.VoucherCode != "" {
		query = query.Where("voucher_code ILIKE ?", "%"+filter.VoucherCode+"%")
	}

	// Count total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	query = query.Limit(filter.Limit).Offset(filter.Offset)

	// Apply sorting (tambahkan default sort jika tidak ada)
	sortBy := filter.SortBy
	sortOrder := filter.SortOrder
	if sortBy == "" {
		sortBy = "created_at"
	}
	if sortOrder == "" {
		sortOrder = "desc"
	}
	query = query.Order(string(sortBy) + " " + string(sortOrder))

	// Execute query
	if err := query.Find(&promotions).Error; err != nil {
		return nil, 0, err
	}

	return promotions, total, nil
}

func (r *promotionRepository) Update(ctx context.Context, promotion *entity.Promotion) error {
	db := infra.GetDB(ctx, r.db)

	result := db.Model(&entity.Promotion{}).
		Where("id = ?", promotion.ID).
		Updates(promotion)

	if result.Error != nil {
		r.log.Error("Error query update promotion", zap.Error(result.Error))
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrPromotionNotFound
	}

	return nil
}

func (r *promotionRepository) Delete(ctx context.Context, id uint) error {
	db := infra.GetDB(ctx, r.db)
	return db.Delete(&entity.Promotion{}, id).Error
}

func (r *promotionRepository) CheckVoucherCodeExists(ctx context.Context, code string, excludeID ...uint) (bool, error) {
	db := infra.GetDB(ctx, r.db)
	var count int64
	query := db.Model(&entity.Promotion{}).Where("voucher_code = ?", code)

	if len(excludeID) > 0 && excludeID[0] > 0 {
		query = query.Where("id != ?", excludeID[0])
	}

	err := query.Count(&count).Error
	return count > 0, err
}

func (r *promotionRepository) GetActivePromotionsByProduct(ctx context.Context, productID uint) ([]entity.Promotion, error) {
	db := infra.GetDB(ctx, r.db)
	var promotions []entity.Promotion

	now := time.Now()
	err := db.Joins("JOIN promo_products ON promo_products.promotion_id = promotions.id").
		Where("promo_products.product_id = ?", productID).
		Where("is_published = ?", true).
		Where("start_date <= ?", now).
		Where("end_date >= ?", now).
		Where("used_count < usage_limit").
		Preload("PromoProducts").
		Find(&promotions).Error

	return promotions, err
}

func (r *promotionRepository) LockPromoByID(ctx context.Context, promoID uint) (*entity.Promotion, error) {
	db := infra.GetDB(ctx, r.db)
	var promo entity.Promotion
	err := db.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", promoID).
		Where("used_count < usage_limit").
		First(&promo).Error
	if err != nil {
		r.log.Error("Error lock promo row", zap.Error(err), zap.Uint("promoID", promoID))
		return nil, err
	}
	return &promo, nil
}

func (r *promotionRepository) LockPromoByCode(ctx context.Context, code string) (*entity.Promotion, error) {
	db := infra.GetDB(ctx, r.db)
	var promo entity.Promotion
	err := db.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("voucher_code = ?", code).
		Where("used_count < usage_limit").
		Where("is_published = ?", true).
		Where("start_date <= ?", time.Now()).
		Where("end_date >= ?", time.Now()).
		First(&promo).Error
	if err != nil {
		r.log.Error("Error lock promo row", zap.Error(err), zap.String("voucher_code", code))
		return nil, err
	}
	return &promo, nil
}

func (r *promotionRepository) CreatePromoUsage(ctx context.Context, promo *entity.PromoUsage) error {
	db := infra.GetDB(ctx, r.db)
	r.log.Info("Creating promo usage",
		zap.Uint("customer_id", promo.CustomerID),
		zap.Uint("order_id", promo.OrderID),
		zap.Uint("promotion_id", promo.PromotionID),
	)

	err := db.Create(promo).Error
	if err != nil {
		r.log.Error("Failed to create promo usage",
			zap.Uint("customer_id", promo.CustomerID),
			zap.Uint("order_id", promo.OrderID),
			zap.Uint("promotion_id", promo.PromotionID),
			zap.Error(err))
		return err
	}

	r.log.Info("Promo usage created successfully",
		zap.Uint("id", promo.ID),
		zap.Uint("customer_id", promo.CustomerID),
		zap.Uint("order_id", promo.OrderID),
		zap.Uint("promotion_id", promo.PromotionID),
	)

	return nil
}