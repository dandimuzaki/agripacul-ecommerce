package repository

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/dto/request"
	infra "debian-ecommerce/internal/infra/transaction"
	"debian-ecommerce/pkg/utils"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepository interface{
	GetAll(ctx context.Context, f request.OrderQueryParams) ([]entity.Order, int64, error)
	CreateOrder(ctx context.Context, order *entity.Order) (*uint, error)
	BatchCreateOrderItems(ctx context.Context, orders []entity.OrderItem) error
	CreateOrderPayment(ctx context.Context, order *entity.OrderPayment) error
	CreateOrderShipping(ctx context.Context, order *entity.OrderShipping) error
	GetByID(ctx context.Context, id uint) (*entity.Order, error)
	UpdateOrder(ctx context.Context, id uint, order *entity.Order) error
	UpdateOrderPayment(ctx context.Context, id uint, order *entity.OrderPayment) error
	UpdateOrderShipping(ctx context.Context, id uint, order *entity.OrderShipping) error
	MarkOrdersAsShipped(ctx context.Context) error
	MarkOrdersAsDelivered(ctx context.Context) error
	MarkOrdersAsCompleted(ctx context.Context) error
	GetOrderSKUsMap(ctx context.Context, orderID uint) (map[uint]uint, error)
}

type orderRepository struct {
	db *gorm.DB
	log *zap.Logger
}

func NewOrderRepo(db *gorm.DB, log *zap.Logger) OrderRepository {
	return &orderRepository{
		db: db,
		log: log,
	}
}

func (r *orderRepository) CreateOrder(ctx context.Context, order *entity.Order) (*uint, error) {
	db := infra.GetDB(ctx, r.db)
	r.log.Info("Creating order",
		zap.Uint("customer_id", order.CustomerID),
	)

	err := db.Create(order).Error
	if err != nil {
		r.log.Error("Failed to create order",
			zap.Uint("customer_id", order.CustomerID),
			zap.Error(err))
		return nil, err
	}

	r.log.Info("order created successfully",
		zap.Uint("id", order.ID),
		zap.Uint("customer_id", order.CustomerID))

	return &order.ID, nil
}

func (r *orderRepository) BatchCreateOrderItems(ctx context.Context, orders []entity.OrderItem) error {
	db := infra.GetDB(ctx, r.db)
	r.log.Info("Creating order items",
		zap.Any("order_ids", orders),
	)

	err := db.Create(&orders).Error
	if err != nil {
		r.log.Error("Failed to create order items",
			zap.Error(err))
		return err
	}

	r.log.Info("Order item created successfully",
	)

	return nil
}

func (r *orderRepository) CreateOrderPayment(ctx context.Context, order *entity.OrderPayment) error {
	db := infra.GetDB(ctx, r.db)
	r.log.Info("Creating order payment",
		zap.Uint("order_id", order.OrderID),
		zap.Uint("payment_method_id", order.PaymentMethodID),
	)

	err := db.Create(order).Error
	if err != nil {
		r.log.Error("Failed to create order payment",
			zap.Uint("order_id", order.OrderID),
			zap.Uint("payment_method_id", order.PaymentMethodID),
			zap.Error(err))
		return err
	}

	r.log.Info("Order payment created successfully",
		zap.Uint("id", order.ID),
		zap.Uint("order_id", order.OrderID),
		zap.Uint("payment_method_id", order.PaymentMethodID),
	)

	return nil
}

func (r *orderRepository) CreateOrderShipping(ctx context.Context, order *entity.OrderShipping) error {
	db := infra.GetDB(ctx, r.db)
	r.log.Info("Creating order shipping",
		zap.Uint("order_id", order.OrderID),
		zap.String("courier_name", order.CourierName),
	)

	err := db.Create(order).Error
	if err != nil {
		r.log.Error("Failed to create order shipping",
			zap.Uint("order_id", order.OrderID),
			zap.String("courier_name", order.CourierName),
			zap.Error(err))
		return err
	}

	r.log.Info("Order shipping created successfully",
		zap.Uint("id", order.ID),
		zap.Uint("order_id", order.OrderID),
		zap.String("courier_name", order.CourierName),
	)

	return nil
}

func (r *orderRepository) GetAll(
	ctx context.Context,
	f request.OrderQueryParams,
) ([]entity.Order, int64, error) {

	db := infra.GetDB(ctx, r.db)

	// ===============================
	// COUNT QUERY
	// ===============================

	countDB := db.Model(&entity.Order{})

	if f.CustomerID != nil {
		countDB = countDB.Where("customer_id = ?", *f.CustomerID)
	}

	// Date filters
	if f.Period != "" {
		now := time.Now()
		start := f.StartDate

		switch f.Period {
		case request.Today:
			start = now.AddDate(0, 0, -1)
		case request.Week:
			start = now.AddDate(0, 0, -7)
		case request.Month:
			start = now.AddDate(0, -1, 0)
		}

		countDB = countDB.Where("created_at BETWEEN ? AND ?", start, now)

	} else if !f.StartDate.IsZero() {

		countDB = countDB.Where("created_at >= ?", f.StartDate)

		if f.EndDate != nil {
			countDB = countDB.Where("created_at <= ?", *f.EndDate)
		}
	}

	// Order status filter
	switch f.Status {

	case
		string(entity.OrderStatusProcess),
		string(entity.OrderStatusCompleted),
		string(entity.OrderStatusCancelled):

		countDB = countDB.Where("status = ?", f.Status)
	
	case "confirmed":
		countDB = countDB.Where("status = ?", entity.OrderStatusProcess)

	case string(entity.PaymentStatusPending),
		string(entity.PaymentStatusPaid),
		string(entity.PaymentStatusExpired):

		countDB = countDB.Where(`
			EXISTS (
				SELECT 1
				FROM order_payments
				WHERE order_payments.order_id = orders.id
				AND order_payments.status = ?
			)`, f.Status)

	case
		string(entity.ShippingStatusShipped),
		string(entity.ShippingStatusDelivered):

		countDB = countDB.Where(`
			EXISTS (
				SELECT 1
				FROM order_shippings
				WHERE order_shippings.order_id = orders.id
				AND order_shippings.status = ?
			)`, f.Status)

	}

	// Shipping method filter
	if f.ShippingMethod != "" {

		countDB = countDB.Where(`
			EXISTS (
				SELECT 1
				FROM order_shippings
				WHERE order_shippings.order_id = orders.id
				AND order_shippings.courier_name = ?
			)`, f.ShippingMethod)

	}

	var total int64

	if err := countDB.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if total == 0 {
		return []entity.Order{}, 0, nil
	}

	// ===============================
	// STEP 1: GET ORDER IDS
	// ===============================

	idQuery := countDB.Select("orders.id")

	// Sorting
	sortDesc := f.SortOrder == request.SortDesc
	switch f.SortBy {
	case request.OrderSortByCreatedAt:
		idQuery = idQuery.Order(clause.OrderByColumn{
			Column: clause.Column{Name: "created_at"},
			Desc:   sortDesc,
		})
	case request.OrderSortByUpdatedAt:
		idQuery = idQuery.Order(clause.OrderByColumn{
			Column: clause.Column{Name: "updated_at"},
			Desc:   sortDesc,
		})
	case request.OrderSortByTotal:
		idQuery = idQuery.Order(clause.OrderByColumn{
			Column: clause.Column{Name: "total"},
			Desc:   sortDesc,
		})
	default:
		idQuery = idQuery.Order("created_at DESC")
	}

	var orderIDs []uint

	err := idQuery.
		Limit(f.Limit).
		Offset(f.Offset).
		Pluck("orders.id", &orderIDs).
		Error

	if err != nil {
		r.log.Error("Error query order ids", zap.Error(err))
		return nil, 0, err
	}

	if len(orderIDs) == 0 {
		return []entity.Order{}, total, nil
	}

	// ===============================
	// STEP 2: FETCH ORDERS WITH PRELOAD
	// ===============================

	var orders []entity.Order

	// build array string: ARRAY[5,2,9,1]
	idStrs := make([]string, len(orderIDs))
	for i, id := range orderIDs {
		idStrs[i] = fmt.Sprintf("%d", id)
	}

	orderExpr := fmt.Sprintf(
		"array_position(ARRAY[%s]::int[], orders.id)",
		strings.Join(idStrs, ","),
	)

	err = db.
		Preload("Items").
		Preload("Customer.User").
		Preload("Shipping").
		Preload("Payment").Preload("Payment.PaymentMethod").
		Where("orders.id IN ?", orderIDs).
		Order(orderExpr).
		Find(&orders).Error

	if err != nil {
		r.log.Error("Error query get order list", zap.Error(err))
		return nil, 0, err
	}

	return orders, total, nil
}

func (r *orderRepository) MarkOrdersAsShipped(ctx context.Context) error {
	db := infra.GetDB(ctx, r.db)

	r.log.Info("Mark orders as shipped...")

	query := `
		UPDATE order_shippings
		SET
				status = ?,
				shipped_at = NOW(),
				updated_at = NOW(),
				tracking_number = ?
		FROM orders o
		WHERE
				order_shippings.order_id = o.id
				AND o.status = ?
				AND order_shippings.status = ?
				AND order_shippings.shipped_at IS NULL
				AND o.confirmed_at IS NOT NULL;
	`

	err := db.Exec(
		query,
		entity.ShippingStatusShipped,
		utils.GenerateTrackingNumber(),
		entity.OrderStatusProcess,
		entity.ShippingStatusPending,
	).Error

	if err != nil {
		r.log.Error("Failed to mark orders as shipped")
		return err
	}

	r.log.Info("Successfully mark orders as shipped")

	return nil
}

func (r *orderRepository) MarkOrdersAsDelivered(ctx context.Context) error {
	db := infra.GetDB(ctx, r.db)

	r.log.Info("Mark orders as delivered...")

	query := `
		UPDATE order_shippings
		SET
			status = ?,
			delivered_at = NOW(),
			updated_at = NOW()
		FROM orders o
		WHERE
			order_shippings.order_id = o.id
			AND o.status = ?
			AND order_shippings.status = ?
			AND order_shippings.shipped_at IS NOT NULL
			AND NOW() >= order_shippings.shipped_at + (order_shippings.etd || ' days')::interval
	`

	err := db.Exec(
		query,
		entity.ShippingStatusDelivered,
		entity.OrderStatusProcess,
		entity.ShippingStatusShipped,
	).Error

	if err != nil {
		r.log.Error("Failed to mark orders as delivered")
		return err
	}

	r.log.Info("Successfully mark orders as delivered")

	return nil
}

func (r *orderRepository) MarkOrdersAsCompleted(ctx context.Context) error {
	db := infra.GetDB(ctx, r.db)

	r.log.Info("Mark orders as completed...")

	query := `
		UPDATE orders
		SET
			status = ?,
			completed_at = NOW(),
			updated_at = NOW()
		FROM order_shippings os
		WHERE
			os.order_id = o.id
			AND orders.status = ?
			AND os.status = ?
			AND os.delivered_at IS NOT NULL
			AND NOW() >= os.delivered_at + '5 days' interval
	`

	err := db.Exec(
		query,
		entity.OrderStatusCompleted,
		entity.OrderStatusProcess,
		entity.ShippingStatusDelivered,
	).Error

	if err != nil {
		r.log.Error("Failed to mark orders as completed")
		return err
	}

	r.log.Info("Successfully mark orders as completed")

	return nil
}

func (r *orderRepository) GetByID(ctx context.Context, id uint) (*entity.Order, error) {
	db := infra.GetDB(ctx, r.db)
	r.log.Info("Get order by id",
		zap.Uint("id", id),
	)

	var order entity.Order
	query := db.Model(&order).Preload("Payment").Preload("Shipping").Preload("Items").Preload("Customer").Preload("Customer.User").Preload("Payment.PaymentMethod")

	err := query.First(&order, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			r.log.Warn("order not found", zap.Uint("id", id))
			return nil, utils.ErrOrderNotFound
		} else {
			r.log.Error("Failed to get order",
				zap.Uint("id", id),
				zap.Error(err))
			return nil, err
		}
	}

	return &order, nil
}

func (r *orderRepository) UpdateOrder(ctx context.Context, id uint, order *entity.Order) error {
	db := infra.GetDB(ctx, r.db)

	result := db.Model(&entity.Order{}).
		Where("id = ?", id).
		Updates(order)

	if result.Error != nil {
		r.log.Error("Error query update order", zap.Error(result.Error))
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrOrderNotFound
	}

	return nil
}

func (r *orderRepository) UpdateOrderPayment(ctx context.Context, id uint, order *entity.OrderPayment) error {
	db := infra.GetDB(ctx, r.db)

	result := db.Model(&entity.OrderPayment{}).
		Where("order_id = ?", id).
		Updates(order)

	if result.Error != nil {
		r.log.Error("Error query update order payment", zap.Error(result.Error))
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrOrderNotFound
	}

	return nil
}

func (r *orderRepository) UpdateOrderShipping(ctx context.Context, id uint, order *entity.OrderShipping) error {
	db := infra.GetDB(ctx, r.db)

	result := db.Model(&entity.OrderShipping{}).
		Where("order_id = ?", id).
		Updates(order)

	if result.Error != nil {
		r.log.Error("Error query update order shipping", zap.Error(result.Error))
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrOrderNotFound
	}

	return nil
}

func (r *orderRepository) GetOrderSKUsMap(
    ctx context.Context,
    orderID uint,
) (map[uint]uint, error) {

	db := infra.GetDB(ctx, r.db)

	type Result struct {
		SKUID     uint `gorm:"column:sku_id"`
		ProductID uint
	}

	var results []Result

	err := db.
		Model(&entity.OrderItem{}).
		Select("order_items.sku_id, s.product_id").
		Joins("JOIN skus s ON s.id = order_items.sku_id").
		Where("order_items.order_id = ?", orderID).
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	skuMap := make(map[uint]uint)
	for _, r := range results {
		skuMap[r.SKUID] = r.ProductID
	}

	r.log.Info("results", zap.Any("results", results))
	r.log.Info("skuMap", zap.Any("skuMap", skuMap))
	return skuMap, nil
}
