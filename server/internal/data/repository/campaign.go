package repository

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/dto/request"
	infra "debian-ecommerce/internal/infra/transaction"
	"debian-ecommerce/pkg/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CampaignRepository interface {
	GetAll(ctx context.Context, f request.CampaignListQuery) ([]entity.Campaign, int64, error)
	Create(ctx context.Context, campaign *entity.Campaign) (*entity.Campaign, error)
	GetWithProduct(ctx context.Context, id uint) (*entity.Campaign, error)
	Update(ctx context.Context, id uint, campaign *entity.Campaign) error
	UpdateActive(ctx context.Context, id uint, data map[string]interface{}) error
	CreateCampaignProduct(ctx context.Context, campaign *entity.CampaignProduct) error
	ReplaceCampaignProducts(ctx context.Context, campaignID uint, productIDs []uint) error
}

type campaignRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewCampaignRepo(db *gorm.DB, log *zap.Logger) CampaignRepository {
	return &campaignRepository{
		db:  db,
		log: log,
	}
}

func (r *campaignRepository) Create(ctx context.Context, campaign *entity.Campaign) (*entity.Campaign, error) {
	db := infra.GetDB(ctx, r.db)
	r.log.Info("Creating campaign",
		zap.String("name", campaign.Name),
	)

	err := db.Create(campaign).Error
	if err != nil {
		r.log.Error("Failed to create campaign",
			zap.String("name", campaign.Name),
			zap.Error(err))
		return nil, err
	}

	r.log.Info("Campaign created successfully",
		zap.Uint("id", campaign.ID),
		zap.String("name", campaign.Name))

	return campaign, nil
}

func (r *campaignRepository) CreateCampaignProduct(ctx context.Context, campaign *entity.CampaignProduct) error {
	db := infra.GetDB(ctx, r.db)
	r.log.Info("Creating campaign",
		zap.Uint("campaign_id", campaign.CampaignID),
		zap.Uint("product_id", campaign.ProductID),
	)

	err := db.Create(campaign).Error
	if err != nil {
		r.log.Error("Failed to create campaign",
			zap.Uint("campaign_id", campaign.CampaignID),
			zap.Uint("product_id", campaign.ProductID),
			zap.Error(err))
		return err
	}

	r.log.Info("Campaign created successfully",
		zap.Uint("id", campaign.ID),
		zap.Uint("campaign_id", campaign.CampaignID),
		zap.Uint("product_id", campaign.ProductID),
	)

	return nil
}

func (r *campaignRepository) GetAll(ctx context.Context, f request.CampaignListQuery) ([]entity.Campaign, int64, error) {
	db := infra.GetDB(ctx, r.db)
	var campaigns []entity.Campaign
	var total int64
	query := db.Model(&entity.Campaign{})

	// Get total campaign
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Limit(f.Limit).Offset(f.Offset).Find(&campaigns).Error
	if err != nil {
		r.log.Error("Error query get campaign list", zap.Error(err))
		return nil, 0, err
	}

	return campaigns, total, nil
}

func (r *campaignRepository) GetWithProduct(ctx context.Context, id uint) (*entity.Campaign, error) {
	db := infra.GetDB(ctx, r.db)
	r.log.Info("Get campaign by id",
		zap.Uint("id", id),
	)

	var campaign entity.Campaign
	query := db.Model(&campaign).
		Preload("CampaignProducts").
		Preload("CampaignProducts.Product")

	err := query.First(&campaign, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			r.log.Warn("campaign not found", zap.Uint("id", id))
			return nil, utils.ErrCampaignNotFound
		} else {
			r.log.Error("Failed to get campaign",
				zap.Uint("id", id),
				zap.Error(err))
			return nil, err
		}
	}

	return &campaign, nil
}

func (r *campaignRepository) Update(ctx context.Context, id uint, campaign *entity.Campaign) error {
	db := infra.GetDB(ctx, r.db)

	result := db.Model(&entity.Campaign{}).
		Where("id = ?", id).
		Updates(campaign)

	if result.Error != nil {
		r.log.Error("Error query update campaign", zap.Error(result.Error))
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrCampaignNotFound
	}

	return nil
}

func (r *campaignRepository) UpdateActive(ctx context.Context, id uint, data map[string]interface{}) error {
	db := infra.GetDB(ctx, r.db)

	result := db.Model(&entity.Campaign{}).
		Where("id = ?", id).
		Updates(data)

	if result.Error != nil {
		r.log.Error("Error query update campaign", zap.Error(result.Error))
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrCampaignNotFound
	}

	return nil
}

func (r *campaignRepository) Delete(ctx context.Context, id uint) error {
	db := infra.GetDB(ctx, r.db)
	err := db.Delete(&entity.Campaign{}, id).Error
	if err != nil {
		r.log.Error("Error query delete campaign", zap.Error(err))
		return err
	}
	return nil
}

func (r *campaignRepository) ReplaceCampaignProducts(
	ctx context.Context,
	campaignID uint,
	productIDs []uint,
) error {
	db := infra.GetDB(ctx, r.db)

	// 1. Delete old relations
	if err := db.
		Where("campaign_id = ?", campaignID).
		Delete(&entity.CampaignProduct{}).
		Error; err != nil {
		return err
	}

	// 2. Insert new relations
	var rows []entity.CampaignProduct
	for _, pid := range productIDs {
		rows = append(rows, entity.CampaignProduct{
			CampaignID: campaignID,
			ProductID:  pid,
		})
	}

	if len(rows) == 0 {
		return nil
	}

	err := db.WithContext(ctx).Create(&rows).Error
	if err != nil {
		r.log.Error("Failed to create campaign products", zap.Error(err))
		return err
	}
	return nil
}
