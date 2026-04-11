package usecase

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/data/repository"
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/dto/response"

	"go.uber.org/zap"
)

type CampaignService interface {
	GetAll(ctx context.Context, req request.CampaignListQuery) (*response.PaginatedResponse[entity.Campaign], error)
	GetWithProduct(ctx context.Context, campaignID uint) (*response.CampaignResponse, error)
	Create(ctx context.Context, req request.CampaignRequest) error
	Update(ctx context.Context, id uint, req request.CampaignRequest) error
}

type campaignService struct {
	tx   TxManager
	repo *repository.Repository
	log  *zap.Logger
}

func NewCampaignService(tx TxManager, repo *repository.Repository, log *zap.Logger) CampaignService {
	return &campaignService{
		tx:   tx,
		repo: repo,
		log:  log,
	}
}

func (s *campaignService) GetAll(ctx context.Context, req request.CampaignListQuery) (*response.PaginatedResponse[entity.Campaign], error) {
	campaigns, total, err := s.repo.CampaignRepo.GetAll(ctx, req)
	if err != nil {
		s.log.Error("Failed to get campaigns", zap.Error(err))
		return nil, err
	}

	return response.NewPaginatedResponse(
		campaigns,
		req.Page,
		req.Limit,
		total,
	), nil
}

func (s *campaignService) GetWithProduct(ctx context.Context, campaignID uint) (*response.CampaignResponse, error) {
	campaign, err := s.repo.CampaignRepo.GetWithProduct(ctx, campaignID)
	if err != nil {
		s.log.Error("Failed to get campaign", zap.Error(err))
		return nil, err
	}

	var products []response.ProductSummary
	for _, p := range campaign.CampaignProducts {
		products = append(products, *response.ToProductSummary(&p.Product))
	}

	res := response.CampaignResponse{
		ID: campaign.ID,
		Name: campaign.Name,
		Description: campaign.Description,
		Type: campaign.Type,
		StartDate: campaign.StartDate,
		EndDate: campaign.EndDate,
		Products: products,
	}

	return &res, nil
}

func (s *campaignService) Create(ctx context.Context, req request.CampaignRequest) error {
	campaign := entity.Campaign{
		Name: req.Name,
		Description: req.Description,
		Type: req.Type,
		StartDate: req.StartDate,
		EndDate: req.EndDate,
	}

	err := s.tx.WithinTx(ctx, func(ctx context.Context) error {
		createdCampaign, err := s.repo.CampaignRepo.Create(ctx, &campaign)
		if err != nil {
			s.log.Error("Failed to create campaign", zap.Error(err))
			return err
		}

		for _, pid := range req.ProductIDs {
			campProduct := entity.CampaignProduct{
				CampaignID: createdCampaign.ID,
				ProductID: pid,
			}
			err := s.repo.CampaignRepo.CreateCampaignProduct(ctx, &campProduct)
			if err != nil {
				s.log.Error("Failed to create campaign", zap.Error(err))
				return err
			}
		}
		return nil
	})

	if err != nil {
		s.log.Error("Failed to create campaign", zap.Error(err))
		return err
	}

	return nil
}

func (s *campaignService) Update(ctx context.Context, id uint, req request.CampaignRequest) error {
	campaign := entity.Campaign{
		Name: req.Name,
		Description: req.Description,
		Type: req.Type,
		StartDate: req.StartDate,
		EndDate: req.EndDate,
	}

	err := s.tx.WithinTx(ctx, func(ctx context.Context) error {
		err := s.repo.CampaignRepo.Update(ctx, id, &campaign)
		if err != nil {
			s.log.Error("Failed to update campaign", zap.Error(err))
			return err
		}

		err = s.repo.CampaignRepo.ReplaceCampaignProducts(ctx, id, req.ProductIDs)
		if err != nil {
			s.log.Error("Failed to update campaign", zap.Error(err))
			return err
		}
		return nil
	})

	if err != nil {
		s.log.Error("Failed to update campaign", zap.Error(err))
		return err
	}

	return nil
}