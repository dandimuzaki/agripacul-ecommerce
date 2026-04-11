package seeder

import (
	"encoding/json"
	"time"

	"debian-ecommerce/internal/data/entity"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type OrderSeeder struct {
	DB *gorm.DB
}

func NewOrderSeeder(db *gorm.DB) *OrderSeeder {
	return &OrderSeeder{DB: db}
}

func (s *OrderSeeder) Run() error {
	var orders []entity.Order
	if err := s.DB.Find(&orders).Error; err != nil {
		return err
	}

	if len(orders) > 0 {
		return nil // Sudah ada order, skip seeding
	}

	if err := SeedOrdersAndReviews(s.DB); err != nil {
		return err
	}

	// // Ambil data yang diperlukan
	// var customers []entity.Customer
	// if err := s.DB.Preload("User", "is_active = ?", true).
	// 	Limit(8).
	// 	Find(&customers).Error; err != nil {
	// 	return err
	// }

	// var paymentMethods []entity.PaymentMethod
	// if err := s.DB.Where("is_active = ?", true).
	// 	Limit(5).
	// 	Find(&paymentMethods).Error; err != nil {
	// 	return err
	// }

	// var promotions []entity.Promotion
	// if err := s.DB.Where("is_published = ? AND end_date > ?",
	// 	true, time.Now()).
	// 	Limit(3).
	// 	Find(&promotions).Error; err != nil {
	// 	return err
	// }

	// var skus []entity.SKU
	// if err := s.DB.Where("is_active = ? AND stock > ?", true, 0).
	// 	Preload("Product").
	// 	Preload("Product.Category").
	// 	Limit(20).
	// 	Find(&skus).Error; err != nil {
	// 	return err
	// }

	// if len(customers) == 0 || len(paymentMethods) == 0 || len(skus) == 0 {
	// 	return nil
	// }

	// // Hapus data lama
	// s.DB.Exec("DELETE FROM order_items")
	// s.DB.Exec("DELETE FROM orders")

	// now := time.Now()

	// // Data orders dengan berbagai status
	// orders = []entity.Order{
	// 	// Order baru (pending payment)
	// 	{
	// 		CustomerID:      customers[0].ID,
	// 		Status:          entity.OrderStatusCreated,
	// 		PromotionID:     getPromotionID(promotions, 0),
	// 		Subtotal:        0, // Akan dihitung
	// 		DiscountAmount:  0,
	// 		Total:           0,
	// 		Notes:           "Tolong dikirim sebelum jam 5 sore",
	// 	},
	// 	// Order dalam proses (sudah bayar)
	// 	{
	// 		CustomerID:      customers[0].ID,
	// 		Status:          entity.OrderStatusProcess,
	// 		PromotionID:     getPromotionID(promotions, 1),
	// 		Subtotal:        0,
	// 		DiscountAmount:  0,
	// 		Total:           0,
	// 		Notes:           "",
	// 	},
	// 	// Order dikirim
	// 	{
	// 		CustomerID:      customers[1].ID,
	// 		Status:          entity.OrderStatusProcess,
	// 		PromotionID:     nil,
	// 		Subtotal:        0,
	// 		DiscountAmount:  0,
	// 		Total:           0,
	// 		Notes:           "Kirim ke alamat kantor",
	// 	},
	// 	// Order delivered
	// 	{
	// 		CustomerID:      customers[1].ID,
	// 		Status:          entity.OrderStatusCompleted,
	// 		PromotionID:     getPromotionID(promotions, 2),
	// 		Subtotal:        0,
	// 		DiscountAmount:  0,
	// 		Total:           0,
	// 		Notes:           "",
	// 	},
	// }

	// // Create orders
	// for i := range orders {
	// 	if err := s.DB.Create(&orders[i]).Error; err != nil {
	// 		return err
	// 	}

	// 	// Buat order items untuk setiap order
	// 	if err := s.createOrderItems(&orders[i], skus); err != nil {
	// 		return err
	// 	}

	// 	// Update total order berdasarkan items
	// 	if err := s.updateOrderTotal(&orders[i]); err != nil {
	// 		return err
	// 	}

	// 	// Buat promo snapshot jika ada promo
	// 	if orders[i].PromotionID != nil {
	// 		if err := s.createPromotionSnapshot(&orders[i]); err != nil {
	// 			return err
	// 		}
	// 	}

	// 	// Update created_at untuk order yang sudah lama
	// 	s.updateOrderTimestamps(&orders[i], i, now)
	// }

	// // Update promo usages untuk order yang menggunakan promo
	// s.updatePromoUsages(orders, customers)

	// for _, order := range orders {
	// 	s.createOrderDetails(&order)
	// }

	return nil
}

func (s *OrderSeeder) createOrderItems(order *entity.Order, skus []entity.SKU) error {
	var orderItems []entity.OrderItem

	// Tentukan berapa banyak items berdasarkan jenis order
	itemCount := 1
	switch order.Status {
	case entity.OrderStatusProcess:
		itemCount = 2
	case entity.OrderStatusCompleted:
		itemCount = 2
	case entity.OrderStatusCancelled:
		itemCount = 1
	}

	// Buat order items
	for i := 0; i < itemCount && i < len(skus); i++ {
		sku := skus[i]
		quantity := 1

		// Quantity lebih banyak untuk item murah
		if sku.Price < 100000 {
			quantity = 2
		}
		if sku.Price < 50000 {
			quantity = 3
		}

		// Buat snapshots
		productSnapshot, err := json.Marshal(map[string]interface{}{
			"id":          sku.Product.ID,
			"name":        sku.Product.Name,
			"category":    sku.Product.Category.Name,
			"description": sku.Product.Description,
		})
		if err != nil {
			return err
		}

		skuSnapshot, err := json.Marshal(map[string]interface{}{
			"id":       sku.ID,
			"sku_code": sku.SKUCode,
			"price":    sku.Price,
			"variant":  "Sample variant info",
		})
		if err != nil {
			return err
		}

		orderItem := entity.OrderItem{
			OrderID:         order.ID,
			SKUID:           sku.ID,
			Quantity:        quantity,
			UnitPrice:       sku.Price,
			TotalPrice:      sku.Price * float64(quantity),
			ProductSnapshot: datatypes.JSON(productSnapshot),
			SKUSnapshot:     datatypes.JSON(skuSnapshot),
		}

		orderItems = append(orderItems, orderItem)
	}

	// Save order items
	for i := range orderItems {
		if err := s.DB.Create(&orderItems[i]).Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *OrderSeeder) updateOrderTotal(order *entity.Order) error {
	// Hitung subtotal dari order items
	var subtotal float64
	var orderItems []entity.OrderItem

	if err := s.DB.Where("order_id = ?", order.ID).Find(&orderItems).Error; err != nil {
		return err
	}

	for _, item := range orderItems {
		subtotal += item.TotalPrice
	}

	// Hitung discount jika ada promo
	discountAmount := 0.0
	if order.PromotionID != nil {
		var promotion entity.Promotion
		if err := s.DB.First(&promotion, order.PromotionID).Error; err == nil {
			switch promotion.DiscountType {
			case entity.DiscountType(entity.DiscountPercentage):
				discountAmount = subtotal * promotion.DiscountValue / 100
				if promotion.MaximumDiscount > 0 && discountAmount > promotion.MaximumDiscount {
					discountAmount = promotion.MaximumDiscount
				}
			case entity.DiscountType(entity.DiscountAmount):
				discountAmount = promotion.DiscountValue
			}
		}
	}

	// Update order totals
	order.Subtotal = subtotal
	order.DiscountAmount = discountAmount
	order.Total = subtotal - discountAmount

	return s.DB.Save(order).Error
}

func (s *OrderSeeder) createPromotionSnapshot(order *entity.Order) error {
	if order.PromotionID == nil {
		return nil
	}

	var promotion entity.Promotion
	if err := s.DB.First(&promotion, order.PromotionID).Error; err != nil {
		return err
	}

	promoSnapshot, err := json.Marshal(map[string]interface{}{
		"id":             promotion.ID,
		"name":           promotion.Name,
		"type":           promotion.Type,
		"discount_type":  promotion.DiscountType,
		"discount_value": promotion.DiscountValue,
		"voucher_code":   promotion.VoucherCode,
	})
	if err != nil {
		return err
	}

	order.PromotionSnapshot = datatypes.JSON(promoSnapshot)
	return s.DB.Save(order).Error
}

func (s *OrderSeeder) updateOrderTimestamps(order *entity.Order, index int, now time.Time) {
	// Set created_at yang berbeda untuk simulasi timeline
	switch index {
	case 0: // Order baru - hari ini
		order.CreatedAt = now
	case 1: // Process - 1 jam yang lalu
		order.CreatedAt = now.Add(-1 * time.Hour)
	case 2: // Shipped - 1 hari yang lalu
		order.CreatedAt = now.Add(-24 * time.Hour)
	case 3: // Delivered - 2 hari yang lalu
		order.CreatedAt = now.Add(-48 * time.Hour)
	case 4: // Completed - 5 hari yang lalu
		order.CreatedAt = now.Add(-120 * time.Hour)
	case 5: // Cancelled - 3 hari yang lalu
		order.CreatedAt = now.Add(-72 * time.Hour)
	}

	s.DB.Save(order)
}

func (s *OrderSeeder) updatePromoUsages(orders []entity.Order, customers []entity.Customer) {
	// Buat promo usage untuk order yang menggunakan promo
	for i, order := range orders {
		if order.PromotionID != nil && i < len(customers) {
			promoUsage := entity.PromoUsage{
				PromotionID: *order.PromotionID,
				CustomerID:  customers[i].ID,
				OrderID:     order.ID,
				UsedAt:      order.CreatedAt,
			}
			s.DB.Create(&promoUsage)
		}
	}
}

// Helper functions
func getPromotionID(promotions []entity.Promotion, index int) *uint {
	if index < len(promotions) {
		return &promotions[index].ID
	}
	return nil
}

func findPaymentMethodID(methods []entity.PaymentMethod, name string) uint {
	for _, method := range methods {
		if method.Name == name {
			return method.ID
		}
	}
	return methods[0].ID // Fallback ke method pertama
}

func (s *OrderSeeder) createOrderDetails(order *entity.Order) {
	shipping := entity.OrderShipping{
		OrderID: order.ID,
		RecipientName: order.Customer.FullName,
		Label: "Home",
		PhoneNumber: order.Customer.PhoneNumber,
		Province: "Jawa Barat",
		Regency: "Kota Cimahi",
		District: "Kecamatan Cimahi Utara",
		Subdistrict: "Kelurahan Pasirkaliki",
		CourierName: "JNE",
		CourierCode: "jne",
		Cost: 20000,
		ETD: 2,
		Status: entity.ShippingStatusPending,
	}

	payment := entity.OrderPayment{
		OrderID: order.ID,
		PaymentMethodID: 1,
		Amount: order.Total,
		Status: entity.PaymentStatusPaid,
	}

	s.DB.Create(&shipping)
	s.DB.Create(&payment)
}