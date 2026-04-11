package usecase

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/data/repository"
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/dto/response"
	"debian-ecommerce/pkg/utils"
	"time"

	"go.uber.org/zap"
)

type CheckoutService interface {
	GetPreviewCheckout(ctx context.Context, req request.PreviewCheckout) (*response.PreviewCheckout, error)
	FetchShippingOptions(ctx context.Context, req request.ShippingOptionsRequest) ([]response.ShippingOption, error)
	GetValidPromotions(ctx context.Context, req request.PromotionFilterRequest) ([]response.PromotionSummary, error)
}

type checkoutService struct {
	tx   TxManager
	repo *repository.Repository
	log  *zap.Logger
	config utils.Configuration
}

func NewCheckoutService(tx TxManager, repo *repository.Repository, log *zap.Logger, config utils.Configuration) CheckoutService {
	return &checkoutService{
		tx:   tx,
		repo: repo,
		log:  log,
		config: config,
	}
}

func (s *checkoutService) GetSubtotal(ctx context.Context) ([]response.PreviewItem, *float64, error) {
	// Get user id from context
	userID := ctx.Value("user_id").(uint)

	// Get customer id
	customer, err := s.repo.CustomerRepo.FindCustomerByUserID(ctx, userID)
	if err != nil {
		s.log.Error("Error find customer by user id", zap.Error(err), zap.Uint("user_id", userID))
		return nil, nil, err
	}

	// Get ordered items (items selected in the cart)
	var previewItems []response.PreviewItem
	items, err := s.repo.CartRepo.GetSelectedCartItems(ctx, customer.ID)
	for _, i := range items {
		product := response.ProductSnapshot{
			ID: i.SKU.ProductID,
			Name: i.SKU.Product.Name,
			MainImageURL: i.SKU.Product.MainImageURL,
			Slug: i.SKU.Product.Slug,
		}

		var discount float64
		if i.SKU.SalePrice != nil {
			discount = (i.SKU.Price - *i.SKU.SalePrice)/100
		}
		price := response.PriceSnapshot{
			UnitPrice: i.SKU.Price,
			SalePrice: i.SKU.SalePrice,
			DiscountPercentage: discount,
		}

		comb, err := s.repo.VariantValueRepo.GetVariantCombination(ctx, i.SKUID)
		if err != nil {
			return nil, nil, err
		}

		subTotal := i.SKU.Price * float64(i.Quantity)
		item := response.PreviewItem{
			ID:          i.ID,
			SKUID:       i.SKUID,
			SKUCode:     i.SKU.SKUCode,
			Product: product,
			Variants: comb,
			Price: price,
			Quantity:    i.Quantity,
			Stock: i.SKU.Stock,
			SubTotal:    subTotal,
		}

		previewItems = append(previewItems, item)
	}

	// Get subtotal
	var subtotal float64
	for _, item := range items {
		sku, err := s.repo.SKURepo.GetByID(ctx, item.SKUID)
		if err != nil {
			s.log.Error("Error get sku", zap.Error(err))
			return nil, nil, err
		}
		if sku.SalePrice != nil {
			subtotal += *sku.SalePrice * float64(item.Quantity)
		} else {
			subtotal += sku.Price * float64(item.Quantity)
		}
	}

	return previewItems, &subtotal, nil
}

func (s *checkoutService) GetPreviewCheckout(ctx context.Context, req request.PreviewCheckout) (*response.PreviewCheckout, error) {
	previewItems, subtotal, err := s.GetSubtotal(ctx)
	if err != nil {
		s.log.Error("Failed to get items", zap.Error(err))
	}
	
	// Get shipping cost based on selected shipping
	var cost float64
	if req.SelectedShippingOption != nil {
		cost = req.SelectedShippingOption.Cost
	}

	// Get discount amount based on selected promo
	var discount float64
	if req.SelectedPromotionID != nil {
		selectedPromo, err := s.repo.PromotionRepo.GetByID(ctx, *req.SelectedPromotionID)
		if err != nil {
			s.log.Error("Failed to get promo", zap.Error(err), zap.Uint("promo_id", *req.SelectedPromotionID))
			return nil, err
		}

		if *subtotal < selectedPromo.MinimumOrderValue {
			s.log.Error("Insufficient minimum order", zap.Error(err), zap.Uint("promo_id", *req.SelectedPromotionID))
			return nil, utils.ErrPromotionNotApplicable
		}

		switch selectedPromo.DiscountType {
			case entity.DiscountAmount: 
			discount = selectedPromo.DiscountValue
		 case entity.DiscountPercentage: 
			draftDiscount := *subtotal * selectedPromo.DiscountValue / 100
			if draftDiscount < selectedPromo.MaximumDiscount {
				discount = draftDiscount
			} else {
				discount = selectedPromo.MaximumDiscount
			}
		}
	}

	// Construct totals
	totals := response.Totals{
		Subtotal: *subtotal,
		DiscountAmount: &discount,
		ShippingCost: cost,
		GrandTotal: *subtotal + cost - discount,
	}

	// Construct response
	res := response.PreviewCheckout{
		Items: previewItems,
		Totals: totals,
	}

	return &res, nil
}

func (s *checkoutService) GetValidPromotions(ctx context.Context, req request.PromotionFilterRequest) ([]response.PromotionSummary, error) {
	_, subtotal, err := s.GetSubtotal(ctx)
	if err != nil {
		s.log.Error("Failed to get items", zap.Error(err))
	}
	
	// Fetch promotion list
	promoFilter := request.PromotionFilterQuery {
		StartDateTo: time.Now(),
		EndDateFrom: time.Now(),
		MinimumOrderValue: *subtotal,
		IsActive: true,
		IsPublished: true,
		Available: true,
		Page: req.Page,
		Limit: req.Limit,
		Offset: req.GetOffset(),
	}
	promos, _, err := s.repo.PromotionRepo.GetList(ctx, promoFilter)
	if err != nil {
		s.log.Error("Error fetch promotions", zap.Error(err))
		return nil, err
	}

	var promoList []response.PromotionSummary
	for _, p := range promos {
		promo := response.ToPromoSummary(p)
		promoList = append(promoList, promo)
	}

	return promoList, nil
}