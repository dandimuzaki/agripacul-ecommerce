package repository

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	infra "debian-ecommerce/internal/infra/transaction"
	"debian-ecommerce/pkg/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SKURepository interface{
	GetIDsByProductID(ctx context.Context, productID uint) ([]uint, error)
	GetByID(ctx context.Context, skuID uint) (*entity.SKU, error)
	Create(ctx context.Context, sku *entity.SKU) (*entity.SKU, error)
	Update(ctx context.Context, id uint, product *entity.SKU) error
	Delete(ctx context.Context, id uint) error
	CreateSKUVariantValue(ctx context.Context, SKU *entity.SKUVariantValue) (*entity.SKUVariantValue, error)
	DeleteSKUVariantValueBySKUIDs(ctx context.Context, SKUIDs []uint) error
	DeleteByVariantValue(ctx context.Context, variantValueIDs []uint) error
	GetVariantValueIDs(ctx context.Context, skuID uint) ([]uint, error)
	FindSKUIDsByVariantValues(ctx context.Context, valueIDs []uint) ([]uint, error)
	ArchiveBatch(ctx context.Context, skuIDs []uint) error
	LockSKUsByIDs(ctx context.Context, skuIDs []uint) ([]entity.SKU, error)
	GetByProductID(ctx context.Context, productID uint) ([]entity.SKU, error)
}

type skuRepository struct {
	db *gorm.DB
	log *zap.Logger
}

func NewSKURepo(db *gorm.DB, log *zap.Logger) SKURepository {
	return &skuRepository{
		db: db,
		log: log,
	}
}

func (r *skuRepository) GetIDsByProductID(ctx context.Context, productID uint) ([]uint, error) {
	db := infra.GetDB(ctx, r.db)
	var ids []uint
	err := db.
		Model(&entity.SKU{}).
		Where("product_id = ?", productID).
		Pluck("id", &ids).
		Error

	if err != nil {
		r.log.Error("Error query get sku ids", zap.Error(err))
		return nil, err
	}

	return ids, nil
}

func (r *skuRepository) GetByID(ctx context.Context, skuID uint) (*entity.SKU, error) {
	db := infra.GetDB(ctx, r.db)
	var sku entity.SKU
	err := db.
		Model(&sku).Preload("Product").
		First(&sku, skuID).
		Error

	if err != nil {
		r.log.Error("Error query get sku", zap.Error(err))
		return nil, err
	}

	return &sku, nil
}

func (r *skuRepository) Create(ctx context.Context, SKU *entity.SKU) (*entity.SKU, error) {
	db := infra.GetDB(ctx, r.db)
	r.log.Info("Creating SKU",
		zap.Uint("product_id", SKU.ProductID),
		zap.String("sku_code", SKU.SKUCode),
	)

	err := db.Create(SKU).Error
	if err != nil {
		r.log.Error("Failed to create SKU",
			zap.String("sku_code", SKU.SKUCode),
			zap.Error(err))
		return nil, err
	}

	r.log.Info("SKU created successfully",
		zap.Uint("id", SKU.ID),
		zap.String("sku_code", SKU.SKUCode))

	return SKU, nil
}

func (r *skuRepository) Update(ctx context.Context, id uint, SKU *entity.SKU) error {
	db := infra.GetDB(ctx, r.db)

	result := db.Model(&entity.SKU{}).
		Where("id = ?", id).
		Updates(SKU)

	if result.Error != nil {
		r.log.Error("Error query update SKU", zap.Error(result.Error))
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrSKUNotFound
	}

	return nil
}

func (r *skuRepository) Delete(ctx context.Context, id uint) error {
	db := infra.GetDB(ctx, r.db)
	result := db.Delete(&entity.SKU{}, id)
	if result.Error != nil {
		r.log.Error("Error query delete SKU", zap.Error(result.Error))
		return result.Error
	}
	if result.RowsAffected == 0 {
		return utils.ErrSKUNotFound
	}
	return nil
}

func (r *skuRepository) DeleteByVariantValue(ctx context.Context, variantValueIDs []uint) error {
	db := infra.GetDB(ctx, r.db)

	subQuery := db.
		Model(&entity.SKUVariantValue{}).
		Select("sku_id").
		Where("variant_value_id IN ?", variantValueIDs)

	result := db.
		Where("id IN (?)", subQuery).
		Delete(&entity.SKU{})

	if result.Error != nil {
		r.log.Error("Error deleting SKU by variant value", zap.Error(result.Error))
		return result.Error
	}

	return nil
}

func (r *skuRepository) CreateSKUVariantValue(ctx context.Context, SKU *entity.SKUVariantValue) (*entity.SKUVariantValue, error) {
	db := infra.GetDB(ctx, r.db)
	r.log.Info("Creating SKU",
		zap.Uint("sku_id", SKU.SKUID),
		zap.Uint("variant_value_id", SKU.VariantValueID),
	)

	err := db.Create(SKU).Error
	if err != nil {
		r.log.Error("Failed to create SKU",
			zap.Uint("sku_id", SKU.SKUID),
			zap.Error(err))
		return nil, err
	}

	r.log.Info("SKU created successfully",
		zap.Uint("id", SKU.ID),
		zap.Uint("sku_id", SKU.SKUID))

	return SKU, nil
}

func (r *skuRepository) DeleteSKUVariantValueBySKUIDs(ctx context.Context, SKUIDs []uint) error {
	db := infra.GetDB(ctx, r.db)
	result := db.Model(&entity.SKUVariantValue{}).Where("sku_id IN ?", SKUIDs).Delete(&entity.SKUVariantValue{})
	if result.Error != nil {
		r.log.Error("Error query delete SKU", zap.Error(result.Error))
		return result.Error
	}
	if result.RowsAffected == 0 {
		return utils.ErrSKUNotFound
	}
	return nil
}

func (r *skuRepository) GetVariantValueIDs(ctx context.Context, skuID uint) ([]uint, error) {
	db := infra.GetDB(ctx, r.db)
	var ids []uint

	err := db.Table("sku_variant_values").
		Where("sku_id = ?", skuID).
		Pluck("variant_value_id", &ids).
		Error

	if err != nil {
		r.log.Error("Error get variant value ids", zap.Error(err))
		return nil, err
	}

	return ids, nil
}

func (r *skuRepository) FindSKUIDsByVariantValues(ctx context.Context, valueIDs []uint) ([]uint, error) {
	if len(valueIDs) == 0 {
		return []uint{}, nil
	}

	db := infra.GetDB(ctx, r.db)

	var skuIDs []uint
	err := db.
		Model(&entity.SKUVariantValue{}).
		Distinct("sku_id").
		Where("variant_value_id IN ?", valueIDs).
		Pluck("sku_id", &skuIDs).
		Error

	if err != nil {
		r.log.Error("Error find sku ids", zap.Error(err))
		return nil, err
	}

	return skuIDs, nil
}

func (r *skuRepository) ArchiveBatch(ctx context.Context, skuIDs []uint) error {
	if len(skuIDs) == 0 {
		return nil
	}

	db := infra.GetDB(ctx, r.db)

	err := db.Model(&entity.SKU{}).
		Where("id IN ?", skuIDs).
		Where("status != ?", entity.SKUStatusArchived).
		Updates(map[string]interface{}{
			"status": "archived",
		}).
		Error
	if err != nil {
		r.log.Error("Error batch archive SKUs", zap.Error(err))
		return err
	}
	return nil
}

func (r *skuRepository) LockSKUsByIDs(ctx context.Context, skuIDs []uint) ([]entity.SKU, error) {
	db := infra.GetDB(ctx, r.db)
	var skus []entity.SKU
	err := db.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Preload("Product").
		Preload("Product.Category").
		Where("id IN ?", skuIDs).
		Where("status = ?", entity.SKUStatusActive).
		Find(&skus).Error
	if err != nil {
		r.log.Error("Error lock sku", zap.Error(err), zap.Any("sku_id", skuIDs))
		return nil, err
	}
	return skus, nil
}

func (r *skuRepository) GetByProductID(ctx context.Context, productID uint) ([]entity.SKU, error) {
	db := infra.GetDB(ctx, r.db)
	var skus []entity.SKU
	err := db.
		Model(&entity.SKU{}).
		Preload("Images").
		Where("product_id = ?", productID).
		Find(&skus).
		Error

	if err != nil {
		r.log.Error("Error query get SKUs", zap.Error(err))
		return nil, err
	}

	return skus, nil
}