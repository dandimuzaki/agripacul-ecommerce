package repository

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	infra "debian-ecommerce/internal/infra/transaction"
	"debian-ecommerce/pkg/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type VariantValueRepository interface{
	GetIDsByTypeID(ctx context.Context, typeID uint) ([]uint, error)
	GetValuesByValueIDs(ctx context.Context, valueIDs []uint) ([]string, error)
	GetVariantCombination(ctx context.Context, skuID uint) ([]entity.VariantCombination, error)
	Create(ctx context.Context, vValue *entity.VariantValue) (*entity.VariantValue, error)
	Update(ctx context.Context, id uint, vValue *entity.VariantValue) error
	Delete(ctx context.Context, id uint) error
	GetVariantCombinations(ctx context.Context, skuIDs []uint) (map[uint][]entity.VariantCombination, error)
}

type variantValueRepository struct {
	db *gorm.DB
	log *zap.Logger
}

func NewVariantValueRepo(db *gorm.DB, log *zap.Logger) VariantValueRepository {
	return &variantValueRepository{
		db: db,
		log: log,
	}
}

func (r *variantValueRepository) GetIDsByTypeID(ctx context.Context, typeID uint) ([]uint, error) {
	db := infra.GetDB(ctx, r.db)
	var ids []uint
	err := db.
		Model(&entity.VariantValue{}).
		Where("variant_type_id = ?", typeID).
		Pluck("id", &ids).
		Error

	if err != nil {
		r.log.Error("Error query get variant value ids", zap.Error(err))
		return nil, err
	}

	return ids, nil
}

func (r *variantValueRepository) GetValuesByValueIDs(ctx context.Context, valueIDs []uint) ([]string, error) {
	db := infra.GetDB(ctx, r.db)
	var values []string
	err := db.
		Model(&entity.VariantValue{}).
		Where("id IN ?", valueIDs).
		Pluck("value", &values).
		Error

	if err != nil {
		r.log.Error("Error query get variant values", zap.Error(err))
		return nil, err
	}

	return values, nil
}



func (r *variantValueRepository) GetVariantCombination(ctx context.Context, skuID uint) ([]entity.VariantCombination, error) {
	db := infra.GetDB(ctx, r.db)
	var variant []entity.VariantCombination
	err := db.
		Model(&entity.VariantValue{}).
		Joins("JOIN sku_variant_values svv ON svv.variant_value_id = variant_values.id").
		Joins("JOIN variant_types vt ON vt.id = variant_values.variant_type_id").
		Where("svv.sku_id = ?", skuID).
		Pluck("vt.name, variant_values.value", &variant).
		Error

	if err != nil {
		r.log.Error("Error query get variant combination", zap.Error(err))
		return nil, err
	}

	return variant, nil
}

func (r *variantValueRepository) BatchGetVariantCombinations(ctx context.Context, skuIDs []uint) ([]entity.VariantCombination, error) {
	db := infra.GetDB(ctx, r.db)
	var variant []entity.VariantCombination
	err := db.
		Model(&entity.VariantValue{}).
		Joins("JOIN sku_variant_values svv ON svv.variant_value_id = variant_values.id").
		Joins("JOIN variant_types vt ON vt.id = variant_values.variant_type_id").
		Where("svv.sku_id IN ?", skuIDs).
		Pluck("vt.name, variant_values.value", &variant).
		Error

	if err != nil {
		r.log.Error("Error query get variant combination", zap.Error(err))
		return nil, err
	}

	return variant, nil
}

func (r *variantValueRepository) Create(ctx context.Context, vValue *entity.VariantValue) (*entity.VariantValue, error) {
	db := infra.GetDB(ctx, r.db)
	r.log.Info("Creating variant value",
		zap.String("value", vValue.Value),
		zap.Uint("variant_type_id", vValue.VariantTypeID),
	)

	err := db.Create(vValue).Error
	if err != nil {
		r.log.Error("Failed to create variant value",
			zap.String("value", vValue.Value),
			zap.Error(err))
		return nil, err
	}

	r.log.Info("Variant value created successfully",
		zap.Uint("id", vValue.ID),
		zap.String("value", vValue.Value))

	return vValue, nil
}

func (r *variantValueRepository) Update(ctx context.Context, id uint, vValue *entity.VariantValue) error {
	db := infra.GetDB(ctx, r.db)

	result := db.Model(&entity.VariantValue{}).
		Where("id = ?", id).
		Updates(vValue)

	if result.Error != nil {
		r.log.Error("Error query update variant value", zap.Error(result.Error))
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrVariantValueNotFound
	}

	return nil
}

func (r *variantValueRepository) Delete(ctx context.Context, id uint) error {
	db := infra.GetDB(ctx, r.db)
	result := db.Delete(&entity.VariantValue{}, id)
	if result.Error != nil {
		r.log.Error("Error query delete variant value", zap.Error(result.Error))
		return result.Error
	}
	if result.RowsAffected == 0 {
		return utils.ErrVariantValueNotFound
	}
	return nil
}

// Fixed variant combination query
func (r *variantValueRepository) GetVariantCombinations(ctx context.Context, skuIDs []uint) (map[uint][]entity.VariantCombination, error) {
	var results []entity.VariantCombination
	db := infra.GetDB(ctx, r.db)
	
	err := db.Table("sku_variant_values").
		Select("sku_variant_values.sku_id, variant_types.name, variant_values.value").
		Joins("JOIN variant_values ON variant_values.id = sku_variant_values.variant_value_id").
		Joins("JOIN variant_types ON variant_types.id = variant_values.variant_type_id").
		Where("sku_variant_values.sku_id IN ?", skuIDs).
		Scan(&results).Error

	if err != nil {
		r.log.Error("Failed to get variant combinations", zap.Error(err))
		return nil, err
	}

	// Map results by SKU ID
	variantsMap := make(map[uint][]entity.VariantCombination)
	for _, r := range results {
		variantsMap[r.SKUID] = append(variantsMap[r.SKUID], entity.VariantCombination{
			Name:  r.Name,
			Value: r.Value,
		})
	}

	return variantsMap, nil
}
