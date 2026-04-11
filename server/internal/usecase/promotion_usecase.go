package usecase

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/data/repository"
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/dto/response"
	"debian-ecommerce/pkg/utils"
	"fmt"
	"time"
)

type PromotionUsecase interface {
	CreatePromotion(ctx context.Context, req request.PromotionRequest) (*response.PromotionResponse, error)
	GetPromotionByID(ctx context.Context, id uint) (*response.PromotionResponse, error)
	GetPromotionList(ctx context.Context, filter request.PromotionFilterQuery) (*response.PaginatedResponse[response.PromotionResponse], error)
	UpdatePromotion(ctx context.Context, id uint, req request.PromotionUpdateRequest) (*response.PromotionResponse, error)
	DeletePromotion(ctx context.Context, id uint) error
	ValidatePromotion(ctx context.Context, promotionID uint, customerID uint, totalAmount float64) (*entity.Promotion, error)
}

type promotionUsecase struct {
	promotionRepo repository.PromotionRepository
	txManager     TxManager
}

func NewPromotionUsecase(promotionRepo repository.PromotionRepository, txManager TxManager) PromotionUsecase {
	return &promotionUsecase{
		promotionRepo: promotionRepo,
		txManager:     txManager,
	}
}

func (uc *promotionUsecase) CreatePromotion(ctx context.Context, req request.PromotionRequest) (*response.PromotionResponse, error) {
	// Validasi tanggal
	if req.StartDate.After(req.EndDate) {
		return nil, fmt.Errorf("start date cannot be after end date")
	}

	var promotion *entity.Promotion
	var err error

	// Gunakan transaction jika tersedia
	if uc.txManager != nil {
		err = uc.txManager.WithinTx(ctx, func(txCtx context.Context) error {
			return uc.createPromotionWithTransaction(txCtx, req, &promotion)
		})
	} else {
		err = uc.createPromotionWithTransaction(ctx, req, &promotion)
	}

	if err != nil {
		return nil, err
	}

	return uc.mapToPromotionResponse(promotion), nil
}

func (uc *promotionUsecase) createPromotionWithTransaction(ctx context.Context, req request.PromotionRequest, promotion **entity.Promotion) error {
	// Validate voucher code uniqueness if provided
	if req.Type == entity.VoucherCode && req.VoucherCode != nil {
		exists, err := uc.promotionRepo.CheckVoucherCodeExists(ctx, *req.VoucherCode)
		if err != nil {
			return fmt.Errorf("%w: failed to check voucher code", utils.ErrDBConnection)
		}
		if exists {
			return fmt.Errorf("%w: voucher code already exists", utils.ErrDBDuplicate)
		}
	}

	// Create promotion
	newPromotion := &entity.Promotion{
		Name:              req.Name,
		StartDate:         req.StartDate,
		EndDate:           req.EndDate,
		Type:              req.Type,
		Description:       req.Description,
		IsPublished:       req.IsPublished,
		DiscountType:      req.DiscountType,
		DiscountValue:     req.DiscountValue,
		MinimumOrderValue: req.MinimumOrderValue,
		MaximumDiscount:   req.MaximumDiscount,
		UsageLimit:        req.UsageLimit,
		VoucherCode:       req.VoucherCode,
		IsPublic:          req.IsPublic,
	}

	// Create promo products if provided
	if len(req.ProductIDs) > 0 {
		for _, productID := range req.ProductIDs {
			newPromotion.PromoProducts = append(newPromotion.PromoProducts, entity.PromoProduct{
				ProductID: productID,
			})
		}
	}

	if err := uc.promotionRepo.Create(ctx, newPromotion); err != nil {
		return fmt.Errorf("%w: failed to create promotion", utils.ErrDBTransaction)
	}

	*promotion = newPromotion
	return nil
}

func (uc *promotionUsecase) GetPromotionByID(ctx context.Context, id uint) (*response.PromotionResponse, error) {
	promotion, err := uc.promotionRepo.GetByID(ctx, id)
	if err != nil {
		return nil, utils.ErrPromotionNotFound
	}

	return uc.mapToPromotionResponse(promotion), nil
}

func (uc *promotionUsecase) GetPromotionList(ctx context.Context, filter request.PromotionFilterQuery) (*response.PaginatedResponse[response.PromotionResponse], error) {
	promotions, total, err := uc.promotionRepo.GetList(ctx, filter)
	if err != nil {
		return nil, utils.ErrInternalServer
	}

	var promotionResponses []response.PromotionResponse
	for _, promotion := range promotions {
		promotionResponses = append(promotionResponses, *uc.mapToPromotionResponse(&promotion))
	}

	return response.NewPaginatedResponse(
		promotionResponses,
		filter.Page,
		filter.Limit,
		total,
	), nil
}

func (uc *promotionUsecase) UpdatePromotion(ctx context.Context, id uint, req request.PromotionUpdateRequest) (*response.PromotionResponse, error) {
	var updatedPromotion *entity.Promotion
	var err error

	if uc.txManager != nil {
		err = uc.txManager.WithinTx(ctx, func(txCtx context.Context) error {
			return uc.updatePromotionWithTransaction(txCtx, id, req, &updatedPromotion)
		})
	} else {
		err = uc.updatePromotionWithTransaction(ctx, id, req, &updatedPromotion)
	}

	if err != nil {
		return nil, err
	}

	return uc.mapToPromotionResponse(updatedPromotion), nil
}

func (uc *promotionUsecase) updatePromotionWithTransaction(ctx context.Context, id uint, req request.PromotionUpdateRequest, updatedPromotion **entity.Promotion) error {
	// 1. Cari promotion yang ada
	promotion, err := uc.promotionRepo.GetByID(ctx, id)
	if err != nil {
		return utils.ErrPromotionNotFound
	}

	// Validate voucher code uniqueness if provided
	if req.VoucherCode != nil && *req.VoucherCode != "" {
		exists, err := uc.promotionRepo.CheckVoucherCodeExists(ctx, *req.VoucherCode, id)
		if err != nil {
			return fmt.Errorf("%w: failed to check voucher code", utils.ErrDBConnection)
		}
		if exists {
			return fmt.Errorf("%w: voucher code already exists", utils.ErrDBDuplicate)
		}
		promotion.VoucherCode = req.VoucherCode
	}

	// Update fields if provided
	if req.Name != "" {
		promotion.Name = req.Name
	}
	if req.StartDate != nil {
		promotion.StartDate = *req.StartDate
	}
	if req.EndDate != nil {
		promotion.EndDate = *req.EndDate
	}
	if req.Type != "" {
		promotion.Type = req.Type
	}
	if req.Description != nil {
		promotion.Description = *req.Description
	}
	if req.IsPublished != nil {
		promotion.IsPublished = *req.IsPublished
	}
	if req.DiscountType != "" {
		promotion.DiscountType = req.DiscountType
	}
	if req.DiscountValue != nil {
		promotion.DiscountValue = *req.DiscountValue
	}
	if req.MinimumOrderValue != nil {
		promotion.MinimumOrderValue = *req.MinimumOrderValue
	}
	if req.MaximumDiscount != nil {
		promotion.MaximumDiscount = *req.MaximumDiscount
	}
	if req.UsageLimit != nil {
		promotion.UsageLimit = *req.UsageLimit
	}
	if req.IsPublic != nil {
		promotion.IsPublic = *req.IsPublic
	}

	// Update promo products if provided
	if req.ProductIDs != nil {
		// Clear existing promo products
		promotion.PromoProducts = nil

		// Add new promo products
		for _, productID := range req.ProductIDs {
			promotion.PromoProducts = append(promotion.PromoProducts, entity.PromoProduct{
				PromotionID: promotion.ID,
				ProductID:   productID,
			})
		}
	}

	if err := uc.promotionRepo.Update(ctx, promotion); err != nil {
		return fmt.Errorf("%w: failed to update promotion", utils.ErrDBTransaction)
	}

	*updatedPromotion = promotion
	return nil
}

func (uc *promotionUsecase) DeletePromotion(ctx context.Context, id uint) error {
	_, err := uc.promotionRepo.GetByID(ctx, id)
	if err != nil {
		return utils.ErrPromotionNotFound
	}

	if err := uc.promotionRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("%w: failed to delete promotion", utils.ErrDBTransaction)
	}

	return nil
}

func (uc *promotionUsecase) ValidatePromotion(ctx context.Context, promotionID uint, customerID uint, totalAmount float64) (*entity.Promotion, error) {
	promotion, err := uc.promotionRepo.GetByID(ctx, promotionID)
	if err != nil {
		return nil, utils.ErrPromotionNotFound
	}

	now := time.Now()
	if !promotion.IsPublished {
		return nil, utils.ErrPromotionInactive
	}
	if now.Before(promotion.StartDate) || now.After(promotion.EndDate) {
		return nil, utils.ErrPromotionExpired
	}
	if promotion.UsageLimit > 0 {
		// Check usage count (implement this method in repository)
		// usageCount, err := uc.promotionRepo.GetUsageCount(promotionID)
		// if err != nil {
		//     return nil, err
		// }
		// if usageCount >= promotion.UsageLimit {
		//     return nil, utils.ErrPromotionUsageLimit
		// }
	}
	if totalAmount < promotion.MinimumOrderValue {
		return nil, utils.ErrMinimumOrderNotMet
	}

	return promotion, nil
}

func (uc *promotionUsecase) mapToPromotionResponse(promotion *entity.Promotion) *response.PromotionResponse {
	resp := &response.PromotionResponse{
		ID:                promotion.ID,
		Name:              promotion.Name,
		StartDate:         promotion.StartDate,
		EndDate:           promotion.EndDate,
		Type:              promotion.Type,
		Description:       promotion.Description,
		IsPublished:       promotion.IsPublished,
		DiscountType:      promotion.DiscountType,
		DiscountValue:     promotion.DiscountValue,
		MinimumOrderValue: promotion.MinimumOrderValue,
		MaximumDiscount:   promotion.MaximumDiscount,
		UsageLimit:        promotion.UsageLimit,
		VoucherCode:       promotion.VoucherCode,
		IsPublic:          promotion.IsPublic,
		CreatedAt:         promotion.CreatedAt,
		UpdatedAt:         promotion.UpdatedAt,
	}

	// Map promo products
	for _, promoProduct := range promotion.PromoProducts {
		productResp := response.PromoProductResponse{
			ID:        promoProduct.ID,
			ProductID: promoProduct.ProductID,
		}
		if promoProduct.Product.ID != 0 {
			productResp.ProductName = promoProduct.Product.Name

			// AMBIL DATA DARI PRODUCT
			// Nama produk sudah ada
			productResp.ProductName = promoProduct.Product.Name

			// Untuk SKU, ambil dari SKUs jika ada, atau gunakan field lain
			if len(promoProduct.Product.SKUs) > 0 {
				// Ambil SKU pertama
				productResp.ProductSKU = promoProduct.Product.SKUs[0].SKUCode
				productResp.Price = promoProduct.Product.SKUs[0].Price
			} else {
				// Fallback jika tidak ada SKU
				productResp.ProductSKU = promoProduct.Product.Slug // atau string kosong
				productResp.Price = promoProduct.Product.MinPrice  // atau 0
			}
		}
		resp.PromoProducts = append(resp.PromoProducts, productResp)
	}

	return resp
}
