package repository

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	infra "debian-ecommerce/internal/infra/transaction"
	"debian-ecommerce/pkg/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type VariantTypeRepository interface{
	GetVariantTypesByProductID(ctx context.Context, productID uint) ([]uint, error)
	Create(ctx context.Context, product *entity.VariantType) (*entity.VariantType, error)
	Update(ctx context.Context, id uint, product *entity.VariantType) error
	Delete(ctx context.Context, id uint) error
}

type variantTypeRepository struct {
	db *gorm.DB
	log *zap.Logger
}

func NewVariantTypeRepo(db *gorm.DB, log *zap.Logger) VariantTypeRepository {
	return &variantTypeRepository{
		db: db,
		log: log,
	}
}

func (r *variantTypeRepository) GetVariantTypesByProductID(ctx context.Context, productID uint) ([]uint, error) {
	db := infra.GetDB(ctx, r.db)
	var typeIDs []uint
	err := db.Model(&entity.VariantType{}).
		Where("variant_types.product_id = ?", productID).
		Pluck("id", &typeIDs).Error
	if err != nil {
		r.log.Error("Error get variant types", zap.Error(err))
		return nil, err
	}
	return typeIDs, nil
}

func (r *variantTypeRepository) Create(ctx context.Context, vType *entity.VariantType) (*entity.VariantType, error) {
	db := infra.GetDB(ctx, r.db)
	r.log.Info("Creating variant type",
		zap.String("name", vType.Name),
		zap.Uint("product_id", vType.ProductID),
	)

	err := db.Create(vType).Error
	if err != nil {
		r.log.Error("Failed to create variant value",
			zap.String("name", vType.Name),
			zap.Error(err))
		return nil, err
	}

	r.log.Info("Variant type created successfully",
		zap.Uint("id", vType.ID),
		zap.String("name", vType.Name))

	return vType, nil
}

func (r *variantTypeRepository) Update(ctx context.Context, id uint, vType *entity.VariantType) error {
	db := infra.GetDB(ctx, r.db)

	result := db.Model(&entity.VariantType{}).
		Where("id = ?", id).
		Updates(vType)

	if result.Error != nil {
		r.log.Error("Error query update variant type", zap.Error(result.Error))
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrVariantTypeNotFound
	}

	return nil
}

func (r *variantTypeRepository) Delete(ctx context.Context, id uint) error {
	db := infra.GetDB(ctx, r.db)
	result := db.Delete(&entity.VariantType{}, id)
	if result.Error != nil {
		r.log.Error("Error query delete variant type", zap.Error(result.Error))
		return result.Error
	}
	if result.RowsAffected == 0 {
		return utils.ErrVariantTypeNotFound
	}
	return nil
}