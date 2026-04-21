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

	"go.uber.org/zap"
)

type OrderService interface {
	Create(ctx context.Context, req request.CreateOrderRequest) (*entity.Order, error)
	GetAll(ctx context.Context, req request.OrderQueryParams) (*response.PaginatedResponse[response.OrderSummary], error)
	GetOrderHistory(ctx context.Context, req request.OrderQueryParams) (*response.PaginatedResponse[response.OrderSummary], error)
	GetDetails(ctx context.Context, id uint) (*response.OrderDetails, error)
	Pay(ctx context.Context, id uint) error
	Confirm(ctx context.Context, id uint) error
	Ship(ctx context.Context) error
	Deliver(ctx context.Context) error
	Complete(ctx context.Context) error
	ManualComplete(ctx context.Context, id uint) error
	Cancel(ctx context.Context, id uint, req request.CancelOrderRequest) error
}

type orderService struct {
	tx   TxManager
	repo *repository.Repository
	log  *zap.Logger
}

func NewOrderService(tx TxManager, repo *repository.Repository, log *zap.Logger) OrderService {
	return &orderService{
		tx:   tx,
		repo: repo,
		log:  log,
	}
}

func (s *orderService) Create(ctx context.Context, req request.CreateOrderRequest) (*entity.Order, error) {
	txStart := time.Now()
	defer func() {
		s.log.Info("CreateOrder duration", zap.Duration("duration", time.Since(txStart)))
	}()

	var order entity.Order
	err := s.tx.WithinTx(ctx, func(ctx context.Context) error {
		// Get customer
		customer, err := s.getCustomerFromContext(ctx)
		if err != nil {
			return err
		}

		// Get cart items
		items, err := s.repo.CartRepo.GetSelectedCartItems(ctx, customer.ID)
		if err != nil {
			s.log.Error("Failed to get cart items", zap.Error(err))
			return err
		}

		// Validate items and calculate subtotal (batch lock SKUs)
		skus, subtotal, err := s.validateAndLockItems(ctx, items)
		if err != nil {
			return err
		}

		// Create order
		order = s.buildOrder(customer.ID, subtotal, req)
		
		// Apply promotion/voucher
		if err := s.applyPromotion(ctx, &order, req, subtotal); err != nil {
			return err
		}
		
		order.Total = order.Subtotal - order.DiscountAmount + req.Shipping.Cost

		// Save order
		orderID, err := s.repo.OrderRepo.CreateOrder(ctx, &order)
		if err != nil {
			s.log.Error("Failed to create order", zap.Error(err))
			return err
		}
		order.ID = *orderID

		// Process order items (batch operations)
		if err := s.processOrderItems(ctx, items, skus, &order, customer.ID); err != nil {
			return err
		}

		// Create payment
		if err := s.createOrderPayment(ctx, &order, req.PaymentMethodID); err != nil {
			return err
		}

		// Create shipping
		if err := s.createOrderShipping(ctx, &order, req); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		s.log.Error("Transaction failed", zap.Error(err))
		return nil, err
	}
	return &order, nil
}

func (s *orderService) getCustomerFromContext(ctx context.Context) (*entity.Customer, error) {
	userID, ok := ctx.Value("user_id").(uint)
	if !ok {
		return nil, utils.ErrInvalidUserID
	}

	customer, err := s.repo.CustomerRepo.FindCustomerByUserID(ctx, userID)
	if err != nil {
		s.log.Error("Failed to find customer", zap.Error(err), zap.Uint("user_id", userID))
		return nil, err
	}
	return customer, nil
}

func (s *orderService) validateAndLockItems(ctx context.Context, items []entity.CartItem) (map[uint]*entity.SKU, float64, error) {
	if len(items) == 0 {
		return nil, 0, utils.ErrCartEmpty
	}

	// Extract SKU IDs
	skuIDs := make([]uint, len(items))
	skuQuantity := make(map[uint]int)
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, 0, utils.ErrInvalidQuantity
		}
		skuIDs[i] = item.SKUID
		skuQuantity[item.SKUID] = item.Quantity
	}

	// Batch lock SKUs
	skus, err := s.repo.SKURepo.LockSKUsByIDs(ctx, skuIDs)
	if err != nil {
		s.log.Error("Failed to lock SKUs", zap.Error(err))
		return nil, 0, err
	}

	// Map SKUs by ID for easy access
	skuMap := make(map[uint]*entity.SKU)
	for _, sku := range skus {
		skuMap[sku.ID] = &sku
	}

	// Validate and calculate subtotal
	var subtotal float64
	for _, item := range items {
		sku := skuMap[item.SKUID]
		
		if sku.Status != entity.SKUStatusActive {
			s.log.Error("SKU inactive", zap.Uint("sku_id", sku.ID))
			return nil, 0, utils.ErrSKUInactive
		}

		if item.Quantity > sku.Stock {
			s.log.Error("Insufficient stock", 
				zap.Uint("sku_id", sku.ID),
				zap.Int("requested", item.Quantity),
				zap.Int("available", sku.Stock))
			return nil, 0, utils.ErrInsufficientStock
		}

		subtotal += sku.Price * float64(item.Quantity)
	}

	return skuMap, subtotal, nil
}

func (s *orderService) buildOrder(customerID uint, subtotal float64, req request.CreateOrderRequest) entity.Order {
	return entity.Order{
		CustomerID: customerID,
		Status:     entity.OrderStatusCreated,
		Notes:      req.Notes,
		Subtotal:   subtotal,
	}
}

func (s *orderService) applyPromotion(ctx context.Context, order *entity.Order, req request.CreateOrderRequest, subtotal float64) error {
	// Can't use both
	if req.PromotionID != nil && req.VoucherCode != nil {
		s.log.Error("Cannot use both promotion and voucher")
		return utils.ErrPromotionNotApplicable
	}

	if req.PromotionID == nil && req.VoucherCode == nil {
		return nil
	}

	var promo *entity.Promotion
	var err error

	if req.PromotionID != nil {
		promo, err = s.repo.PromotionRepo.LockPromoByID(ctx, *req.PromotionID)
	} else {
		promo, err = s.repo.PromotionRepo.LockPromoByCode(ctx, *req.VoucherCode)
	}

	if err != nil {
		s.log.Error("Failed to get promotion", zap.Error(err))
		return err
	}

	// Validate minimum order
	if subtotal < promo.MinimumOrderValue {
		s.log.Error("Insufficient minimum order", zap.Uint("promo_id", promo.ID))
		return utils.ErrPromotionNotApplicable
	}

	// Apply to order
	order.PromotionID = &promo.ID
	
	snapshot, err := utils.ToJSONB(entity.NewPromotionSnapshot(*promo))
	if err != nil {
		s.log.Error("Failed to create promotion snapshot", zap.Error(err))
		return err
	}
	order.PromotionSnapshot = snapshot
	
	order.DiscountAmount = s.calculateDiscount(subtotal, promo)

	// Update promotion usage count
	promo.UsedCount++
	promo.UpdatedAt = time.Now()
	return s.repo.PromotionRepo.Update(ctx, promo)
}

func (s *orderService) calculateDiscount(subtotal float64, promo *entity.Promotion) float64 {
	switch promo.DiscountType {
	case entity.DiscountAmount:
		return promo.DiscountValue
	case entity.DiscountPercentage:
		discount := subtotal * promo.DiscountValue / 100
		if discount > promo.MaximumDiscount {
			return promo.MaximumDiscount
		}
		return discount
	default:
		return 0
	}
}

func (s *orderService) processOrderItems(ctx context.Context, items []entity.CartItem, skuMap map[uint]*entity.SKU, order *entity.Order, customerID uint) error {
	// Pre-fetch all variant combinations in one query
	skuIDs := make([]uint, 0, len(items))
	for _, item := range items {
		skuIDs = append(skuIDs, item.SKUID)
	}
	
	variantsMap, err := s.repo.VariantValueRepo.GetVariantCombinations(ctx, skuIDs)
	if err != nil {
		return err
	}

	// Prepare batch data
	orderItems := make([]entity.OrderItem, 0, len(items))
	inventoryLogs := make([]entity.InventoryLog, 0, len(items))
	productUpdates := make(map[uint]int)
	cartItemIDs := make([]uint, 0, len(items))

	for _, item := range items {
		sku := skuMap[item.SKUID]
		
		// Update stock
		sku.Stock -= item.Quantity
		if err := s.repo.SKURepo.Update(ctx, sku.ID, sku); err != nil {
			s.log.Error("Failed to update SKU", zap.Error(err), zap.Uint("sku_id", sku.ID))
			return err
		}

		// Prepare inventory log
		inventoryLogs = append(inventoryLogs, entity.InventoryLog{
			SKUID:             sku.ID,
			Type:              entity.InventoryLogTypeOut,
			QuantityChange:    item.Quantity,
			CurrentStockAfter: sku.Stock,
			ReferenceID:       &order.ID,
			ReferenceType:     "order",
		})

		// Track product sold count
		productUpdates[sku.ProductID] += item.Quantity

		// Create order item
		orderItem, err := s.buildOrderItem(order.ID, sku, item, variantsMap[sku.ID])
		if err != nil {
			return err
		}
		orderItems = append(orderItems, *orderItem)

		cartItemIDs = append(cartItemIDs, item.ID)
	}

	// Batch insert inventory logs
	if err := s.repo.InventoryRepo.BatchCreateInventoryLogs(ctx, inventoryLogs); err != nil {
		s.log.Error("Failed to create inventory logs", zap.Error(err))
		return err
	}

	// Batch update product sold counts
	if err := s.batchUpdateProductSoldCounts(ctx, productUpdates); err != nil {
		return err
	}

	// Batch insert order items
	if err := s.repo.OrderRepo.BatchCreateOrderItems(ctx, orderItems); err != nil {
		s.log.Error("Failed to create order items", zap.Error(err))
		return err
	}

	// Batch remove cart items
	if err := s.repo.CartRepo.BatchRemoveItems(ctx, cartItemIDs); err != nil {
		s.log.Error("Failed to remove cart items", zap.Error(err))
		return err
	}

	// Record promotion usage
	if order.PromotionID != nil {
		if err := s.recordPromotionUsage(ctx, customerID, order.PromotionID, order.ID); err != nil {
			return err
		}
	}

	return nil
}

func (s *orderService) buildOrderItem(orderID uint, sku *entity.SKU, item entity.CartItem, variants []entity.VariantCombination) (*entity.OrderItem, error) {
	// Create product snapshot
	productSnapshot, err := utils.ToJSONB(entity.NewProductSnapshot(sku.Product))
	if err != nil {
		s.log.Error("Failed to create product snapshot", zap.Error(err))
		return nil, err
	}

	// Create SKU snapshot with variants
	skuSnapshot := entity.NewSKUSnapshot(*sku)
	skuSnapshot.Variants = variants
	
	skuJSON, err := utils.ToJSONB(skuSnapshot)
	if err != nil {
		s.log.Error("Failed to create SKU snapshot", zap.Error(err))
		return nil, err
	}

	// Determine unit price
	unitPrice := sku.Price
	if sku.SalePrice != nil {
		unitPrice = *sku.SalePrice
	}

	return &entity.OrderItem{
		OrderID:         orderID,
		SKUID:           sku.ID,
		Quantity:        item.Quantity,
		UnitPrice:       unitPrice,
		TotalPrice:      unitPrice * float64(item.Quantity),
		ProductSnapshot: productSnapshot,
		SKUSnapshot:     skuJSON,
	}, nil
}

func (s *orderService) batchUpdateProductSoldCounts(ctx context.Context, updates map[uint]int) error {
	if len(updates) == 0 {
		return nil
	}

	// Get all product IDs
	productIDs := make([]uint, 0, len(updates))
	for id := range updates {
		productIDs = append(productIDs, id)
	}

	// Batch fetch products
	products, err := s.repo.ProductRepo.BatchGetByIDs(ctx, productIDs)
	if err != nil {
		s.log.Error("Failed to get products", zap.Error(err))
		return err
	}

	// Update sold counts
	for _, product := range products {
		product.SoldCount += updates[product.ID]
	}

	// Batch update
	return s.repo.ProductRepo.BatchUpdate(ctx, products)
}

func (s *orderService) recordPromotionUsage(ctx context.Context, customerID uint, promotionID *uint, orderID uint) error {
	promoUsage := entity.PromoUsage{
		CustomerID:  customerID,
		PromotionID: *promotionID,
		OrderID:     orderID,
		UsedAt:      time.Now(),
	}
	
	if err := s.repo.PromotionRepo.CreatePromoUsage(ctx, &promoUsage); err != nil {
		s.log.Error("Failed to create promo usage", zap.Error(err))
		return err
	}
	return nil
}

func (s *orderService) createOrderPayment(ctx context.Context, order *entity.Order, paymentMethodID uint) error {
	orderPayment := entity.OrderPayment{
		OrderID:         order.ID,
		PaymentMethodID: paymentMethodID,
		Amount:          order.Total,
		Status:          entity.PaymentStatusPaid,
		ExpiredAt:       time.Now().Add(5 * time.Minute),
	}
	
	if err := s.repo.OrderRepo.CreateOrderPayment(ctx, &orderPayment); err != nil {
		s.log.Error("Failed to create order payment", zap.Error(err))
		return err
	}
	return nil
}

func (s *orderService) createOrderShipping(ctx context.Context, order *entity.Order, req request.CreateOrderRequest) error {
	address, err := s.repo.AddressRepo.FindByID(req.ShippingAddressID)
	if err != nil {
		s.log.Error("Failed to get address", zap.Error(err))
		return err
	}

	orderShipping := entity.OrderShipping{
		OrderID:         order.ID,
		RecipientName:   address.RecipientName,
		Label:           address.Label,
		PhoneNumber:     address.PhoneNumber,
		DetailAddress:   address.DetailAddress,
		Province:        address.Province.Name,
		Regency:         address.Regency.Name,
		District:        address.District.Name,
		Subdistrict:     address.Subdistrict.Name,
		PostalCode:      address.PostalCode,
		CourierName:     req.Shipping.Name,
		CourierCode:     req.Shipping.Code,
		CourierService:  req.Shipping.Service,
		Cost:            req.Shipping.Cost,
		ETD:             utils.ParseETDToDays(req.Shipping.ETD),
		Status:          entity.ShippingStatusPending,
	}
	
	if err := s.repo.OrderRepo.CreateOrderShipping(ctx, &orderShipping); err != nil {
		s.log.Error("Failed to create order shipping", zap.Error(err))
		return err
	}
	return nil
}

// func (s *orderService) Create(ctx context.Context, req request.CreateOrderRequest) error {
// 	txStart := time.Now()
// 	err := s.tx.WithinTx(ctx, func(ctx context.Context) error {
// 		// Get user id from context
// 		userID := ctx.Value("user_id").(uint)

// 		// Get customer id
// 		customer, err := s.repo.CustomerRepo.FindCustomerByUserID(ctx, userID)
// 		if err != nil {
// 			s.log.Error("Error find customer by user id", zap.Error(err), zap.Uint("user_id", userID))
// 			return err
// 		}

// 		// Get ordered items (items selected in the cart)
// 		items, err := s.repo.CartRepo.GetSelectedCartItems(ctx, customer.ID)

// 		var subtotal float64
// 		for _, item := range items {
// 			// Lock sku row
// 			sku, err := s.repo.SKURepo.LockSKURow(ctx, item.SKUID)
// 			if err != nil {
// 				s.log.Error("Error get sku", zap.Error(err))
// 				return err
// 			}

// 			// Validate item
// 			if sku.Status != entity.SKUStatusActive {
// 				s.log.Error("SKU is not active", zap.Error(err), zap.Uint("sku_id", sku.ID))
// 				return utils.ErrSKUInactive
// 			}

// 			if item.Quantity == 0 {
// 				s.log.Error("Quantity must be greater than 0", zap.Error(err), zap.Int("quantity", item.Quantity))
// 				return utils.ErrInvalidQuantity
// 			}

// 			if item.Quantity > sku.Stock {
// 				s.log.Error("Insufficient stock", zap.Error(err), zap.Int("quantity", item.Quantity))
// 				return utils.ErrInsufficientStock
// 			}

// 			// Calculate subtotal
// 			subtotal += sku.Price * float64(item.Quantity)
// 		}

// 		// Initialize order
// 		order := entity.Order{
// 			CustomerID: customer.ID,
// 			Status: entity.OrderStatusCreated,
// 			Notes: req.Notes,
// 			Subtotal: subtotal,
// 		}

// 		// Direct discount and voucher must not be used at the same time
// 		if req.PromotionID != nil && req.VoucherCode != nil {
// 			s.log.Error("Direct discount and voucher code must not be used at the same time", zap.Error(err))
// 			return utils.ErrPromotionNotApplicable
// 		}

// 		// Get promo
// 		if req.PromotionID != nil {
// 			promo, err := s.repo.PromotionRepo.LockPromoByID(ctx, *req.PromotionID)
// 			if err != nil {
// 				s.log.Error("Failed to get promo", zap.Error(err), zap.Uint("promo_id", *req.PromotionID))
// 				return err
// 			}

// 			if subtotal < promo.MinimumOrderValue {
// 				s.log.Error("Insufficient minimum order", zap.Error(err), zap.Uint("promo_id", *req.PromotionID))
// 				return utils.ErrPromotionNotApplicable
// 			}

// 			order.PromotionID = &promo.ID
// 			snapshot, err := utils.ToJSONB(entity.NewPromotionSnapshot(*promo))
// 			if err != nil {
// 				s.log.Error("Error generate promotion snapshot", zap.Error(err), zap.Uint("promo_id", *req.PromotionID))
// 				return err
// 			}
// 			order.PromotionSnapshot = snapshot

// 			switch promo.DiscountType {
// 			case entity.DiscountAmount:
// 				order.DiscountAmount = promo.DiscountValue
// 			case entity.DiscountPercentage:
// 				discount := subtotal * promo.DiscountValue / 100
// 				if discount < promo.MaximumDiscount {
// 					order.DiscountAmount = discount
// 				} else {
// 					order.DiscountAmount = promo.MaximumDiscount
// 				}
// 			}

// 			// Update promotion
// 			promo.UsedCount += 1
// 			promo.UpdatedAt = time.Now()
// 			err = s.repo.PromotionRepo.Update(ctx, promo)
// 			if err != nil {
// 				s.log.Error("Failed to update promo", zap.Error(err), zap.Uint("promo_id", *req.PromotionID))
// 				return err
// 			}
// 		}

// 		// Get voucher code
// 		if req.VoucherCode != nil {
// 			voucher, err := s.repo.PromotionRepo.LockPromoByCode(ctx, *req.VoucherCode)
// 			if err != nil {
// 				s.log.Error("Failed to get voucher", zap.Error(err), zap.Uint("promo_id", *req.PromotionID))
// 				return err
// 			}
// 			if subtotal < voucher.MinimumOrderValue {
// 				s.log.Error("Insufficient minimum order", zap.Error(err), zap.Uint("promo_id", *req.PromotionID))
// 				return utils.ErrPromotionNotApplicable
// 			}
// 			order.PromotionID = &voucher.ID
// 			snapshot, err := utils.ToJSONB(entity.NewPromotionSnapshot(*voucher))
// 			if err != nil {
// 				s.log.Error("Error generate promotion snapshot", zap.Error(err), zap.Uint("promo_id", *req.PromotionID))
// 				return err
// 			}
// 			order.PromotionSnapshot = snapshot
// 			switch voucher.DiscountType {
// 			case entity.DiscountAmount:
// 				order.DiscountAmount = voucher.DiscountValue
// 			case entity.DiscountPercentage:
// 				discount := subtotal * voucher.DiscountValue / 100
// 				if discount < voucher.MaximumDiscount {
// 					order.DiscountAmount = discount
// 				} else {
// 					order.DiscountAmount = voucher.MaximumDiscount
// 				}
// 			}

// 			// Update promotion
// 			voucher.UsedCount += 1
// 			voucher.UpdatedAt = time.Now()
// 			err = s.repo.PromotionRepo.Update(ctx, voucher)
// 			if err != nil {
// 				s.log.Error("Failed to update promo", zap.Error(err), zap.Uint("promo_id", voucher.ID))
// 				return err
// 			}
// 		}

// 		// Create order
// 		order.Total = order.Subtotal - order.DiscountAmount + req.Shipping.Cost
// 		orderID, err := s.repo.OrderRepo.CreateOrder(ctx, &order)
// 		if err != nil {
// 			s.log.Error("Error create order", zap.Error(err))
// 			return err
// 		}

// 		// Record promotion usage
// 		if order.PromotionID != nil {
// 			promoUsage := entity.PromoUsage{
// 				CustomerID: customer.ID,
// 				PromotionID: *order.PromotionID,
// 				OrderID: *orderID,
// 				UsedAt: time.Now(),
// 			}
// 			err = s.repo.PromotionRepo.CreatePromoUsage(ctx, &promoUsage)
// 			if err != nil {
// 				s.log.Error("Error create promo usage", zap.Error(err), zap.Uint("promoID", promoUsage.PromotionID))
// 				return err
// 			}
// 		}

// 		// Create order_items
// 		for _, item := range items {
// 			sku, err := s.repo.SKURepo.LockSKURow(ctx, item.SKUID)
// 			if err != nil {
// 				s.log.Error("Error get sku", zap.Error(err), zap.Uint("skuID", item.SKUID))
// 				return err
// 			}

// 			// Reduce SKU stock
// 			sku.Stock -= item.Quantity
// 			err = s.repo.SKURepo.Update(ctx, sku.ID, sku)
// 			if err != nil {
// 				s.log.Error("Failed to update sku", zap.Error(err), zap.Uint("sku_id", sku.ID))
// 				return err
// 			}

// 			// Record inventory log
// 			inventoryLog := entity.InventoryLog{
// 				SKUID: sku.ID,
// 				Type: entity.InventoryLogTypeOut,
// 				QuantityChange: item.Quantity,
// 				CurrentStockAfter: sku.Stock - item.Quantity,
// 				ReferenceID: orderID,
// 				ReferenceType: "order",
// 			}
// 			_, err = s.repo.InventoryRepo.CreateInventoryLog(ctx, &inventoryLog)
// 			if err != nil {
// 				s.log.Error("Failed to create inventory log", zap.Error(err), zap.Uint("sku_id", sku.ID))
// 				return err
// 			}

// 			// Update sold count
// 			product, err := s.repo.ProductRepo.GetByID(ctx, sku.ProductID)
// 			if err != nil {
// 				s.log.Error("Failed to get product", zap.Error(err), zap.Uint("product_id", sku.ProductID))
// 				return err
// 			}
// 			product.SoldCount += item.Quantity
// 			err = s.repo.ProductRepo.Update(ctx, product.ID, product)
// 			if err != nil {
// 				s.log.Error("Failed to update product", zap.Error(err), zap.Uint("sku_id", sku.ID))
// 				return err
// 			}

// 			// Initialize order_item
// 			orderItem := entity.OrderItem{
// 				OrderID: *orderID,
// 				SKUID: sku.ID,
// 				Quantity: item.Quantity,
// 			}

// 			// Create product snapshot
// 			productSnapshot, err := utils.ToJSONB(entity.NewProductSnapshot(sku.Product))
// 			if err != nil {
// 				s.log.Error("Error create product snapshot", zap.Error(err), zap.Uint("productID", sku.ProductID))
// 				return err
// 			}

// 			// Create sku snapshot
// 			skuSnapshot := entity.NewSKUSnapshot(*sku)
// 			variants, err := s.repo.VariantValueRepo.GetVariantCombination(ctx, sku.ID)
// 			if err != nil {
// 				s.log.Error("Error get variants", zap.Error(err), zap.Uint("skuID", sku.ID))
// 				return err
// 			}
// 			skuSnapshot.Variants = variants
// 			skuJSON, err := utils.ToJSONB(skuSnapshot)
// 			if err != nil {
// 				s.log.Error("Error create sku snapshot", zap.Error(err), zap.Uint("skuID", sku.ID))
// 				return err
// 			}

// 			// Assign snapshot to order_item
// 			orderItem.ProductSnapshot = productSnapshot
// 			orderItem.SKUSnapshot = skuJSON

// 			// Assign price to order_item
// 			if sku.SalePrice != nil {
// 				orderItem.UnitPrice = *sku.SalePrice
// 			} else {
// 				orderItem.UnitPrice = sku.Price
// 			}
// 			orderItem.TotalPrice = orderItem.UnitPrice * float64(orderItem.Quantity)

// 			// Create order_item
// 			err = s.repo.OrderRepo.CreateOrderItem(ctx, &orderItem)
// 			if err != nil {
// 				s.log.Error("Error create order item", zap.Error(err), zap.Uint("orderID", *orderID))
// 				return err
// 			}

// 			// Delete item in cart
// 			err = s.repo.CartRepo.RemoveItem(ctx, item.ID)
// 			if err != nil {
// 				s.log.Error("Error remove item from cart", zap.Error(err), zap.Uint("orderID", *orderID))
// 				return err
// 			}
// 		}

// 		// Create order_payment
// 		orderPayment := entity.OrderPayment{
// 			OrderID: *orderID,
// 			PaymentMethodID: req.PaymentMethodID,
// 			Amount: order.Total,
// 			Status: entity.PaymentStatusPending,
// 			ExpiredAt: time.Now().Add(5 * time.Minute),
// 		}
// 		err = s.repo.OrderRepo.CreateOrderPayment(ctx, &orderPayment)
// 		if err != nil {
// 			s.log.Error("Error create order payment", zap.Error(err), zap.Uint("orderID", *orderID))
// 			return err
// 		}

// 		// Get customer address
// 		address, err := s.repo.AddressRepo.FindByID(req.ShippingAddressID)
// 		if err != nil {
// 			s.log.Error("Error get address", zap.Error(err), zap.Uint("customerID", customer.ID))
// 			return err
// 		}

// 		// Create order_shipping
// 		etd := utils.ParseETDToDays(req.Shipping.ETD)
// 		orderShipping := entity.OrderShipping{
// 			OrderID: *orderID,
// 			RecipientName: address.RecipientName,
// 			Label: address.Label,
// 			PhoneNumber: customer.Phone,
// 			DetailAddress: address.DetailAddress,
// 			Province: address.Province.Name,
// 			Regency: address.Regency.Name,
// 			District: address.District.Name,
// 			Subdistrict: address.Subdistrict.Name,
// 			PostalCode: address.PostalCode,
// 			CourierName: req.Shipping.Name,
// 			CourierCode: req.Shipping.Code,
// 			CourierService: req.Shipping.Service,
// 			Cost: req.Shipping.Cost,
// 			ETD: etd,
// 			Status: entity.ShippingStatusPending,
// 		}
// 		err = s.repo.OrderRepo.CreateOrderShipping(ctx, &orderShipping)
// 		if err != nil {
// 			s.log.Error("Error create order shipping", zap.Error(err), zap.Uint("orderID", *orderID))
// 			return err
// 		}

// 		return nil
// 	})
// 	if err != nil {
// 		s.log.Error("Transaction create order failed", zap.Error(err))
// 		return err
// 	}
// 	defer func() {
// 			s.log.Info(
// 					"CreateOrder duration",
// 					zap.Duration("duration", time.Since(txStart)),
// 			)
// 	}()
// 	return nil
// }

func (s *orderService) Pay(ctx context.Context, id uint) error {
	err  := s.tx.WithinTx(ctx, func(ctx context.Context) error {
		order, err := s.repo.OrderRepo.GetByID(ctx, id)
		if err != nil {
			s.log.Error("Failed to get order", zap.Error(err), zap.Uint("orderID", id))
			return err
		}

		if order.Payment.ExpiredAt.Before(time.Now()) {
			s.log.Error("Payment expired", zap.Error(err))
			return utils.ErrPaymentExpired
		}
		
		switch order.Payment.Status {
		case entity.PaymentStatusPaid:
			s.log.Error("Payment already paid", zap.Error(err))
			return utils.ErrOrderAlreadyPaid
		case entity.PaymentStatusExpired:
			s.log.Error("Payment expired", zap.Error(err))
			return utils.ErrPaymentExpired
		case entity.PaymentStatusPending:
			order.Payment.Status = entity.PaymentStatusPaid
			now := time.Now()
			order.Payment.PaidAt = &now
			order.Payment.UpdatedAt = now
			err = s.repo.OrderRepo.UpdateOrderPayment(ctx, order.ID, &order.Payment)
			if err != nil {
				s.log.Error("Failed to update order payment", zap.Error(err), zap.Uint("orderID", id))
				return err
			}
		}

		return nil
	})

	if err != nil {
		s.log.Error("Transaction pay order failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *orderService) Confirm(ctx context.Context, id uint) error {
	err  := s.tx.WithinTx(ctx, func(ctx context.Context) error {
		order, err := s.repo.OrderRepo.GetByID(ctx, id)
		if err != nil {
			s.log.Error("Failed to get order", zap.Error(err), zap.Uint("orderID", id))
			return err
		}

		if order.Status != entity.OrderStatusCreated {
			s.log.Error("Cannot confirm processed order", zap.Error(err), zap.Uint("orderID", id))
			return utils.ErrInvalidOrderStatusTransition
		}

		if order.Payment.Status != entity.PaymentStatusPaid {
			s.log.Error("Cannot confirm unpaid order", zap.Error(err), zap.Uint("orderID", id))
			return utils.ErrOrderUnpaid
		}

		order.Status = entity.OrderStatusProcess
		now := time.Now()
		order.ConfirmedAt = &now
		order.UpdatedAt = now
		err = s.repo.OrderRepo.UpdateOrder(ctx, order.ID, order)
		if err != nil {
			s.log.Error("Failed to update order", zap.Error(err), zap.Uint("orderID", id))
			return err
		}

		return nil
	})

	if err != nil {
		s.log.Error("Transaction confirm order failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *orderService) Ship(ctx context.Context) error {
	err := s.repo.OrderRepo.MarkOrdersAsShipped(ctx)
	if err != nil {
		s.log.Error("Failed to update order", zap.Error(err))
		return err
	}
	return nil
}

func (s *orderService) Deliver(ctx context.Context) error {
	err := s.repo.OrderRepo.MarkOrdersAsDelivered(ctx)
	if err != nil {
		s.log.Error("Failed to update order", zap.Error(err))
		return err
	}
	return nil
}

func (s *orderService) Complete(ctx context.Context) error {
	err := s.repo.OrderRepo.MarkOrdersAsCompleted(ctx)
	if err != nil {
		s.log.Error("Failed to update order", zap.Error(err))
		return err
	}
	return nil
}

func (s *orderService) ManualComplete(ctx context.Context, id uint) error {
	err  := s.tx.WithinTx(ctx, func(ctx context.Context) error {
		order, err := s.repo.OrderRepo.GetByID(ctx, id)
		if err != nil {
			s.log.Error("Failed to get order", zap.Error(err), zap.Uint("orderID", id))
			return err
		}

		if order.Status != entity.OrderStatusProcess {
			s.log.Error("Cannot complete unprocessed order or completed order", zap.Error(err), zap.Uint("orderID", id))
			return utils.ErrInvalidOrderStatusTransition
		}

		if order.Shipping.Status != entity.ShippingStatusDelivered {
			s.log.Error("Cannot complete undelivered order", zap.Error(err), zap.Uint("orderID", id))
			return utils.ErrInvalidOrderStatusTransition
		}

		if order.Payment.Status != entity.PaymentStatusPaid {
			s.log.Error("Cannot complete unpaid order", zap.Error(err), zap.Uint("orderID", id))
			return utils.ErrOrderUnpaid
		}

		order.Status = entity.OrderStatusCompleted
		now := time.Now()
		order.CompletedAt = &now
		order.UpdatedAt = now
		err = s.repo.OrderRepo.UpdateOrder(ctx, order.ID, order)
		if err != nil {
			s.log.Error("Failed to update order", zap.Error(err), zap.Uint("orderID", id))
			return err
		}

		return nil
	})

	if err != nil {
		s.log.Error("Transaction complete order failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *orderService) Cancel(ctx context.Context, id uint, req request.CancelOrderRequest) error {
	role := ctx.Value("role").(string)
	isAdmin := role == string(entity.RoleAdmin)
	err  := s.tx.WithinTx(ctx, func(ctx context.Context) error {
		order, err := s.repo.OrderRepo.GetByID(ctx, id)
		if err != nil {
			s.log.Error("Failed to get order", zap.Error(err), zap.Uint("orderID", id))
			return err
		}

		if !CanCancel(*order, isAdmin) {
			s.log.Error("Cancel order is forbidden", zap.Error(err), zap.Uint("orderID", id))
			return utils.ErrForbidden
		}

		order.Status = entity.OrderStatusCancelled
		order.CancelReason = req.CancelReason
		now := time.Now()
		order.CancelledAt = &now
		order.CancelledBy = role
		order.UpdatedAt = now
		err = s.repo.OrderRepo.UpdateOrder(ctx, order.ID, order)
		if err != nil {
			s.log.Error("Failed to update order", zap.Error(err), zap.Uint("orderID", id))
			return err
		}

		for _, item := range order.Items {
			sku, err := s.repo.SKURepo.GetByID(ctx, item.SKUID)
			if err != nil {
				s.log.Error("Failed to get SKU", zap.Error(err), zap.Uint("SKUID", item.SKUID))
				return err
			}
			sku.Stock += item.Quantity
			err = s.repo.SKURepo.Update(ctx, sku.ID, sku)
			if err != nil {
				s.log.Error("Failed to update SKU stock", zap.Error(err), zap.Uint("SKUID", item.SKUID))
				return err
			}
			product, err := s.repo.ProductRepo.GetByID(ctx, sku.ProductID)
			if err != nil {
				s.log.Error("Failed to get product", zap.Error(err), zap.Uint("productID", sku.ProductID))
				return err
			}
			product.SoldCount -= item.Quantity
			err = s.repo.ProductRepo.Update(ctx, sku.ProductID, product)
			if err != nil {
				s.log.Error("Failed to update SKU stock", zap.Error(err), zap.Uint("SKUID", item.SKUID))
				return err
			}
			inventoryLog := entity.InventoryLog{
				SKUID: sku.ID,
				Type: entity.InventoryLogTypeIn,
				QuantityChange: item.Quantity,
				CurrentStockAfter: sku.Stock + item.Quantity,
				ReferenceID: &order.ID,
				ReferenceType: "order cancelled",
				Notes: req.CancelReason,
			}
			_, err = s.repo.InventoryRepo.CreateInventoryLog(ctx, &inventoryLog)
			if err != nil {
				s.log.Error("Failed to create inventory log", zap.Error(err), zap.Uint("SKUID", item.SKUID))
				return err
			}
		}

		return nil
	})

	if err != nil {
		s.log.Error("Transaction complete order failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *orderService) GetAll(ctx context.Context, req request.OrderQueryParams) (*response.PaginatedResponse[response.OrderSummary], error) {
	role := ctx.Value("role")
	isAdmin := role == entity.RoleAdmin
	
	orders, total, err := s.repo.OrderRepo.GetAll(ctx, req)
	if err != nil {
		s.log.Error("Failed to get orders", zap.Error(err))
		return nil, err
	}

	var res []response.OrderSummary
	for _, o := range orders {
		summary, err := response.ToSummary(o)
		if err != nil {
			s.log.Error("Failed to get order summary", zap.Error(err))
			return nil, err
		}
		summary.DisplayStatus = BuildDisplayStatus(o)
		summary.Cancellation = CanCancel(o, isAdmin)
		summary.CanReview = s.CanReview(ctx, &o)
		res = append(res, *summary)
	}
	return response.NewPaginatedResponse(
		res,
		req.Page,
		req.Limit,
		total,
	), nil
}

func (s *orderService) GetOrderHistory(ctx context.Context, req request.OrderQueryParams) (*response.PaginatedResponse[response.OrderSummary], error) {
	// Get user id from context
	userID := ctx.Value("user_id").(uint)
	role := ctx.Value("role")
	isAdmin := role == entity.RoleAdmin

	// Get customer id
	customer, err := s.repo.CustomerRepo.FindCustomerByUserID(ctx, userID)
	if err != nil {
		s.log.Error("Error find customer by user id", zap.Error(err), zap.Uint("user_id", userID))
		return nil, err
	}

	req.CustomerID = &customer.ID
	
	orders, total, err := s.repo.OrderRepo.GetAll(ctx, req)
	if err != nil {
		s.log.Error("Failed to get orders", zap.Error(err))
		return nil, err
	}

	var res []response.OrderSummary
	for _, o := range orders {
		summary, err := response.ToSummary(o)
		if err != nil {
			s.log.Error("Failed to get order summary", zap.Error(err))
			return nil, err
		}
		summary.DisplayStatus = BuildDisplayStatus(o)
		summary.Cancellation = CanCancel(o, isAdmin)
		summary.CanReview = s.CanReview(ctx, &o)
		res = append(res, *summary)
	}
	return response.NewPaginatedResponse(
		res,
		req.Page,
		req.Limit,
		total,
	), nil
}

func (s *orderService) GetDetails(ctx context.Context, id uint) (*response.OrderDetails, error) {	
	// Get user role
	role := ctx.Value("role")
	isAdmin := role == entity.RoleAdmin
	
	order, err := s.repo.OrderRepo.GetByID(ctx, id)
	if err != nil {
		s.log.Error("Failed to get order details", zap.Error(err), zap.Uint("order id", id))
		return nil, err
	}

	details, err := response.ToDetails(*order)
	if err != nil {
		s.log.Error("Failed to get order details", zap.Error(err), zap.Uint("order id", id))
		return nil, err
	}

	details.DisplayStatus = BuildDisplayStatus(*order)
	details.Steps = BuildSteps(*order)
	details.Cancellation = CanCancel(*order, isAdmin)
	details.CanReview = s.CanReview(ctx, order)

	return details, nil
}

type OrderStep struct {
	Key   string     `json:"key"`
	Label string     `json:"label"`
	Done  bool       `json:"done"`
	At    *time.Time `json:"at,omitempty"`
	Descripttion string `json:"description"`
}

func BuildSteps(order entity.Order) []OrderStep {
	steps := []OrderStep{}

	// Step 1: Order placed
	steps = append(steps, OrderStep{
		Key:   "created",
		Label: "Order Placed",
		Done:  true,
		At:    &order.CreatedAt,
		Descripttion: "Your order is successfully created and please proceed to payment",
	})

	// Step 2: Payment
	steps = append(steps, OrderStep{
		Key:   "paid",
		Label: "Payment Confirmed",
		Done:  order.Payment.PaidAt != nil,
		At:    order.Payment.PaidAt,
		Descripttion: "Your order payment is confirmed and it’s now being reviewed",
	})

	// Step 3: Confirmed
	steps = append(steps, OrderStep{
		Key:   "confirmed",
		Label: "Order Confirmed",
		Done:  order.ConfirmedAt != nil,
		At:    order.ConfirmedAt,
		Descripttion: "Your order has been confirmed and is being prepared.",
	})

	// Step 4: Shipped

	var etd string = "soon"
	if order.Shipping.ETD != 0 {
		etd = fmt.Sprintf("%d", order.Shipping.ETD)
	}

	steps = append(steps, OrderStep{
		Key:   "shipped",
		Label: "Order Shipped",
		Done:  order.Shipping.ShippedAt != nil,
		At:    order.Shipping.ShippedAt,
		Descripttion: fmt.Sprintf("Your order is on its way. Estimated delivery: %v.", etd),
	})

	// Step 5: Delivered
	steps = append(steps, OrderStep{
		Key:   "delivered",
		Label: "Order Delivered",
		Done:  order.Shipping.DeliveredAt != nil,
		At:    order.Shipping.DeliveredAt,
		Descripttion: "Your order has arrived! Please check it at your convenience.",
	})

	// Step 6: Completed
	steps = append(steps, OrderStep{
		Key:   "completed",
		Label: "Order Completed",
		Done:  order.CompletedAt != nil,
		At:    order.CompletedAt,
		Descripttion: "Your order has been completed. Thank you for shopping with us!",
	})

	// Step 7: Cancelled (if cancelled)
	if order.Status == entity.OrderStatusCancelled {
		cancelledAt := order.CancelledAt
		steps = append(steps, OrderStep{
			Key:   "cancelled",
			Label: "Order Cancelled",
			Done:  true,
			At:    cancelledAt,
			Descripttion: "Your order has been cancelled successfully. We’re sorry for the inconvenience.",
		})

			// Set all future steps after cancellation to done=false
		for i, step := range steps {
			if step.At != nil && step.At.After(*cancelledAt) {
					steps[i].Done = false
					steps[i].At = nil
			}
		}
	}

	return steps
}

func BuildDisplayStatus(order entity.Order) string {
	if order.Payment.Status == entity.PaymentStatusPending {
		return "waiting for payment"
	}

	if order.Status == entity.OrderStatusCancelled {
		return "cancelled"
	}
	if order.Status == entity.OrderStatusCompleted {
		return "completed"
	}
	if order.Shipping.Status == entity.ShippingStatusDelivered {
		return "delivered"
	}
	if order.Shipping.Status == entity.ShippingStatusShipped {
		return "shipped"
	}
	if order.Status == entity.OrderStatusProcess {
		return "processing"
	}
	if order.Payment.Status == entity.PaymentStatusPaid {
		return "paid"
	}

	return "pending"
}

func CanCancel(order entity.Order, isAdmin bool) bool {
	if order.Status == entity.OrderStatusCancelled {
		return false // already cancelled
	}

	if isAdmin {
		// Admin can cancel before delivered
		return order.Shipping.DeliveredAt == nil
	}

	// Customer can cancel before shipped
	if order.Shipping.ShippedAt == nil {
		return true
	}

	return false
}

func (s *orderService) CanReview(ctx context.Context, order *entity.Order) (bool) {
	var skuIDs []uint
	for _, item := range order.Items {
		skuIDs = append(skuIDs, item.SKUID)
	}

	// Check for existing review
	reviewedMap, err := s.repo.ReviewRepo.GetReviewedSKUs(ctx, order.Customer.ID, order.ID, skuIDs)
	if err != nil {
    return false
	}

	for _, rev := range reviewedMap {
		if rev {
			return false
		}
	}

	return true
}