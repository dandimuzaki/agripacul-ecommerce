package usecase

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/data/repository"
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/dto/response"
	"debian-ecommerce/pkg/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CartUsecase interface {
	GetCart(ctx context.Context) (*response.CartResponse, error)
	AddItem(ctx context.Context, req request.AddCartItemRequest) error
	UpdateItem(ctx context.Context, itemID uint, req request.UpdateCartItemRequest) error
	RemoveItem(ctx context.Context, itemID uint) error
	ClearCart(ctx context.Context) error
	BatchSelectItem(ctx context.Context, req request.BatchSelectCartItemRequest) error
}

type cartUsecase struct {
	repo *repository.Repository
	log  *zap.Logger
}

func NewCartUsecase(repo *repository.Repository, log *zap.Logger) CartUsecase {
	return &cartUsecase{
		repo: repo,
		log:  log,
	}
}

func (u *cartUsecase) GetCart(ctx context.Context) (*response.CartResponse, error) {
	// Get user id from context
	userIDVal := ctx.Value("user_id")
	userID, ok := userIDVal.(uint)
	if !ok {
		return nil, utils.ErrUserNotFound
	}

	// Get customer id
	customer, err := u.repo.CustomerRepo.FindCustomerByUserID(ctx, userID)
	if err != nil {
		u.log.Error("Error find customer by user id", zap.Error(err), zap.Uint("user_id", userID))
		return nil, err
	}

	cart, err := u.getOrCreateCart(ctx, customer.ID)
	if err != nil {
		return nil, err
	}

	res := &response.CartResponse{
		ID:         cart.ID,
		CustomerID: cart.CustomerID,
		Items:      []response.CartItemResponse{},
	}

	var totalPrice float64
	var totalItems int
	var totalSelectedPrice float64
	var totalSelectedItems int

	for _, item := range cart.Items {
		product := response.ProductSnapshot{
			ID: item.SKU.ProductID,
			Name: item.SKU.Product.Name,
			MainImageURL: item.SKU.Product.MainImageURL,
			Slug: item.SKU.Product.Slug,
		}

		var discount float64
		if item.SKU.SalePrice != nil {
			discount = (item.SKU.Price - *item.SKU.SalePrice)/100
		}
		price := response.PriceSnapshot{
			UnitPrice: item.SKU.Price,
			SalePrice: item.SKU.SalePrice,
			DiscountPercentage: discount,
		}

		comb, err := u.repo.VariantValueRepo.GetVariantCombination(ctx, item.SKUID)
		if err != nil {
			return nil, err
		}

		subTotal := item.SKU.Price * float64(item.Quantity)
		cartItem := response.CartItemResponse{
			ID:          item.ID,
			SKUID:       item.SKUID,
			SKUCode:     item.SKU.SKUCode,
			Product: product,
			Variants: comb,
			Price: price,
			Quantity:    item.Quantity,
			Stock: item.SKU.Stock,
			SubTotal:    subTotal,
			IsSelected: item.IsSelected,
			IsAvailable: item.SKU.Stock >= item.Quantity,
		}

		res.Items = append(res.Items, cartItem)
		totalPrice += subTotal
		totalItems += item.Quantity

		if item.IsSelected {
			totalSelectedPrice += subTotal
			totalSelectedItems += item.Quantity
		}
	}

	totalSummary := response.TotalSnapshot{
		TotalItems: totalItems,
		TotalPrice: totalPrice,
		TotalSelectedItems: totalSelectedItems,
		TotalSelectedPrice: totalSelectedPrice,
	}

	res.Summary = totalSummary

	return res, nil
}

func (u *cartUsecase) AddItem(ctx context.Context, req request.AddCartItemRequest) error {
	if req.Quantity <= 0 {
		return utils.ErrInsufficientQuantity
	}
	
	// Get user id from context
	userIDVal := ctx.Value("user_id")
	userID, ok := userIDVal.(uint)
	if !ok {
		return utils.ErrUserNotFound
	}

	// Get customer id
	customer, err := u.repo.CustomerRepo.FindCustomerByUserID(ctx, userID)
	if err != nil {
		u.log.Error("Error find customer by user id", zap.Error(err), zap.Uint("user_id", userID))
		return err
	}
	
	cart, err := u.getOrCreateCart(ctx, customer.ID)
	if err != nil {
		return err
	}

	existingItem, err := u.repo.CartRepo.GetCartItemBySKU(ctx, cart.ID, req.SKUID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	data := make(map[string]interface{})
	if existingItem != nil {
		newQuantity := existingItem.Quantity + req.Quantity
		data["quantity"] = newQuantity
		return u.repo.CartRepo.UpdateItem(ctx, existingItem.ID, customer.ID, data)
	}

	newItem := &entity.CartItem{
		CartID:   cart.ID,
		SKUID:    req.SKUID,
		Quantity: req.Quantity,
	}

	return u.repo.CartRepo.AddItem(ctx, newItem)
}

func (u *cartUsecase) UpdateItem(ctx context.Context, itemID uint, req request.UpdateCartItemRequest) error {
	// Get user id from context
	userIDVal := ctx.Value("user_id")
	userID, ok := userIDVal.(uint)
	if !ok {
		return utils.ErrUserNotFound
	}

	// Get customer id
	customer, err := u.repo.CustomerRepo.FindCustomerByUserID(ctx, userID)
	if err != nil {
		u.log.Error("Error find customer by user id", zap.Error(err), zap.Uint("user_id", userID))
		return err
	}

	// Get item by id
	_, err = u.repo.CartRepo.GetItemByID(ctx, itemID, customer.ID)
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	if req.Quantity != nil {
		data["quantity"] = req.Quantity
	}

	if req.IsSelected != nil {
		data["is_selected"] = req.IsSelected
	}
	return u.repo.CartRepo.UpdateItem(ctx, itemID, customer.ID, data)
}

func (u *cartUsecase) RemoveItem(ctx context.Context, itemID uint) error {
	// Get user id from context
	userIDVal := ctx.Value("user_id")
	userID, ok := userIDVal.(uint)
	if !ok {
		return utils.ErrUserNotFound
	}

	// Get customer id
	customer, err := u.repo.CustomerRepo.FindCustomerByUserID(ctx, userID)
	if err != nil {
		u.log.Error("Error find customer by user id", zap.Error(err), zap.Uint("user_id", userID))
		return err
	}

	cart, err := u.getOrCreateCart(ctx, customer.ID)
	if err != nil {
		return err
	}

	// Verify ownership
	found := false
	for _, item := range cart.Items {
		if item.ID == itemID {
			found = true
			break
		}
	}

	if !found {
		return utils.ErrCartItemNotFound
	}

	return u.repo.CartRepo.BatchRemoveItems(ctx, []uint{itemID})
}

func (u *cartUsecase) ClearCart(ctx context.Context) error {
	// Get user id from context
	userIDVal := ctx.Value("user_id")
	userID, ok := userIDVal.(uint)
	if !ok {
		return utils.ErrUserNotFound
	}

	// Get customer id
	customer, err := u.repo.CustomerRepo.FindCustomerByUserID(ctx, userID)
	if err != nil {
		u.log.Error("Error find customer by user id", zap.Error(err), zap.Uint("user_id", userID))
		return err
	}

	cart, err := u.getOrCreateCart(ctx, customer.ID)
	if err != nil {
		return err
	}
	return u.repo.CartRepo.ClearCart(ctx, cart.ID)
}

// Helper to get or create cart
func (u *cartUsecase) getOrCreateCart(ctx context.Context, customerID uint) (*entity.Cart, error) {
	cart, err := u.repo.CartRepo.GetCartByCustomerID(ctx, customerID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			newCart := &entity.Cart{CustomerID: customerID}
			if err := u.repo.CartRepo.CreateCart(ctx, newCart); err != nil {
				return nil, err
			}
			return newCart, nil
		}
		return nil, err
	}
	return cart, nil
}

func (u *cartUsecase) BatchSelectItem(ctx context.Context, req request.BatchSelectCartItemRequest) error {
	// Get user id from context
	userIDVal := ctx.Value("user_id")
	userID, ok := userIDVal.(uint)
	if !ok {
		return utils.ErrUserNotFound
	}

	// Get customer id
	customer, err := u.repo.CustomerRepo.FindCustomerByUserID(ctx, userID)
	if err != nil {
		u.log.Error("Error find customer by user id", zap.Error(err), zap.Uint("user_id", userID))
		return err
	}
	
	_, err = u.getOrCreateCart(ctx, customer.ID)
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	data["is_selected"] = req.IsSelected
	err = u.repo.CartRepo.BatchSelectItem(ctx, customer.ID, data)
	if err != nil {
		u.log.Error("Error batch select cart items", zap.Error(err), zap.Uint("user_id", userID))
		return err
	}

	return nil
}