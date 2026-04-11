package seeder

import (
	"time"

	"debian-ecommerce/internal/data/entity"

	"gorm.io/gorm"
)

type PromotionSeeder struct {
	DB *gorm.DB
}

func NewPromotionSeeder(db *gorm.DB) *PromotionSeeder {
	return &PromotionSeeder{DB: db}
}

func (s *PromotionSeeder) Run() error {
	var promos []entity.Promotion
	if err := s.DB.Find(&promos).Error; err != nil {
		return err
	}

	if len(promos) > 0 {
		return nil // Sudah ada promotion, skip seeding
	}

	// Ambil beberapa produk untuk promo
	var products []entity.Product
	if err := s.DB.Where("is_published = ?", true).Limit(10).Find(&products).Error; err != nil {
		return err
	}

	// Ambil beberapa customer untuk promo usage
	var customers []entity.Customer
	if err := s.DB.Preload("User", "is_active = ?", true).Limit(5).Find(&customers).Error; err != nil {
		return err
	}

	now := time.Now()

	// Data promosi
	promotions := []entity.Promotion{
		// Promo aktif - diskon sitewide
		{
			Name:              "Grand Opening Sale 50%",
			StartDate:         now.Add(-7 * 24 * time.Hour),
			EndDate:           now.Add(7 * 24 * time.Hour),
			Type:              entity.DirectDiscount,
			Description:       "Diskon 50% untuk semua produk selama periode grand opening",
			IsPublished:       true,
			DiscountType:      entity.DiscountType(entity.DiscountAmount),
			DiscountValue:     50000,
			MinimumOrderValue: 100000,
			MaximumDiscount:   500000,
			UsageLimit:        1000,
			VoucherCode:       nil, // Direct discount tidak perlu voucher
			IsPublic:           true,
		},
		// Promo aktif - voucher code
		{
			Name:              "Welcome Voucher",
			StartDate:         now.Add(-30 * 24 * time.Hour),
			EndDate:           now.Add(60 * 24 * time.Hour),
			Type:              entity.VoucherCode,
			Description:       "Voucher welcome untuk customer baru, potongan Rp 50.000",
			IsPublished:       true,
			DiscountType:      entity.DiscountType(entity.DiscountAmount),
			DiscountValue:     50000,
			MinimumOrderValue: 200000,
			MaximumDiscount:   50000,
			UsageLimit:        500,
			VoucherCode:       stringPtr("WELCOME50"),
			IsPublic:           true,
		},
		// Promo untuk produk tertentu
		{
			Name:              "Flash Sale Personal Care",
			StartDate:         now,
			EndDate:           now.Add(2 * 24 * time.Hour),
			Type:              entity.DirectDiscount,
			Description:       "Flash sale khusus produk personal care",
			IsPublished:       true,
			DiscountType:      entity.DiscountType(entity.DiscountAmount),
			DiscountValue:     30,
			MinimumOrderValue: 0,
			MaximumDiscount:   300000,
			UsageLimit:        100,
			VoucherCode:       nil,
			IsPublic:           true,
		},
		// Promo akan datang
		{
			Name:              "Black Friday Sale",
			StartDate:         now.Add(14 * 24 * time.Hour),
			EndDate:           now.Add(21 * 24 * time.Hour),
			Type:              entity.VoucherCode,
			Description:       "Black Friday special discount up to 70%",
			IsPublished:       true,
			DiscountType:      entity.DiscountType(entity.DiscountAmount),
			DiscountValue:     70,
			MinimumOrderValue: 500000,
			MaximumDiscount:   1000000,
			UsageLimit:        200,
			VoucherCode:       stringPtr("BLACKFRIDAY70"),
			IsPublic:           true,
		},
		// Promo sudah berakhir
		{
			Name:              "New Year Sale 2023",
			StartDate:         now.Add(-60 * 24 * time.Hour),
			EndDate:           now.Add(-30 * 24 * time.Hour),
			Type:              entity.DirectDiscount,
			Description:       "New Year celebration discount",
			IsPublished:       false,
			DiscountType:      entity.DiscountType(entity.DiscountAmount),
			DiscountValue:     25000,
			MinimumOrderValue: 150000,
			MaximumDiscount:   250000,
			UsageLimit:        300,
			VoucherCode:       nil,
			IsPublic:           false,
		},
		// Promo free shipping
		{
			Name:              "Free Shipping Weekend",
			StartDate:         now.Add(-1 * 24 * time.Hour),
			EndDate:           now.Add(2 * 24 * time.Hour),
			Type:              entity.DirectDiscount,
			Description:       "Gratis ongkir untuk semua order di akhir pekan",
			IsPublished:       true,
			DiscountType:      entity.DiscountType(entity.DiscountAmount),
			DiscountValue:     20000, // Asumsi ongkir rata-rata
			MinimumOrderValue: 100000,
			MaximumDiscount:   20000,
			UsageLimit:        0, // Unlimited
			VoucherCode:       nil,
			IsPublic:           true,
		},
		// Promo khusus member
		{
			Name:              "Member Exclusive 20%",
			StartDate:         now.Add(-15 * 24 * time.Hour),
			EndDate:           now.Add(45 * 24 * time.Hour),
			Type:              entity.VoucherCode,
			Description:       "Diskon eksklusif untuk member terdaftar",
			IsPublished:       true,
			DiscountType:      entity.DiscountType(entity.DiscountPercentage),
			DiscountValue:     20,
			MinimumOrderValue: 100000,
			MaximumDiscount:   200000,
			UsageLimit:        1000,
			VoucherCode:       stringPtr("MEMBER20"),
			IsPublic:           true,
		},
		// Promo bundle
		{
			Name:              "Buy 2 Get 10% Off",
			StartDate:         now.Add(-10 * 24 * time.Hour),
			EndDate:           now.Add(20 * 24 * time.Hour),
			Type:              entity.DirectDiscount,
			Description:       "Diskon khusus untuk pembelian 2 item atau lebih",
			IsPublished:       true,
			DiscountType:      entity.DiscountType(entity.DiscountPercentage),
			DiscountValue:     10,
			MinimumOrderValue: 0,
			MaximumDiscount:   150000,
			UsageLimit:        0,
			VoucherCode:       nil,
			IsPublic:           true,
		},
		// Promo tidak aktif
		{
			Name:              "Test Promo (Inactive)",
			StartDate:         now.Add(-5 * 24 * time.Hour),
			EndDate:           now.Add(5 * 24 * time.Hour),
			Type:              entity.VoucherCode,
			Description:       "Promo test yang tidak aktif",
			IsPublished:       false,
			DiscountType:      entity.DiscountType(entity.DiscountPercentage),
			DiscountValue:     15,
			MinimumOrderValue: 0,
			MaximumDiscount:   100000,
			UsageLimit:        50,
			VoucherCode:       stringPtr("TEST15"),
			IsPublic:           false,
		},
		// Promo besar untuk produk tertentu
		{
			Name:              "Clearance Sale - Up to 60%",
			StartDate:         now.Add(-3 * 24 * time.Hour),
			EndDate:           now.Add(10 * 24 * time.Hour),
			Type:              entity.DirectDiscount,
			Description:       "Clearance sale untuk produk tertentu",
			IsPublished:       true,
			DiscountType:      entity.DiscountType(entity.DiscountPercentage),
			DiscountValue:     60,
			MinimumOrderValue: 0,
			MaximumDiscount:   500000,
			UsageLimit:        200,
			VoucherCode:       nil,
			IsPublic:           true,
		},
	}

	// Hapus data lama
	s.DB.Exec("DELETE FROM promo_usages")
	s.DB.Exec("DELETE FROM promo_products")
	s.DB.Exec("DELETE FROM promotions")

	// Create promotions
	for i := range promotions {
		if err := s.DB.Create(&promotions[i]).Error; err != nil {
			return err
		}
	}

	// Tambahkan promo products untuk promo tertentu
	if len(products) > 0 {
		// Promo Flash Sale Electronics (index 2)
		if len(products) >= 3 {
			// Ambil 3 produk pertama (asumsi elektronik)
			for j := 0; j < 3 && j < len(products); j++ {
				promoProduct := entity.PromoProduct{
					PromotionID: promotions[2].ID,
					ProductID:   products[j].ID,
				}
				if err := s.DB.Create(&promoProduct).Error; err != nil {
					return err
				}
			}
		}

		// Promo Clearance Sale (index 9)
		if len(products) >= 5 {
			// Ambil 5 produk terakhir untuk clearance sale
			start := len(products) - 5
			if start < 0 {
				start = 0
			}
			for j := start; j < len(products); j++ {
				promoProduct := entity.PromoProduct{
					PromotionID: promotions[9].ID,
					ProductID:   products[j].ID,
				}
				if err := s.DB.Create(&promoProduct).Error; err != nil {
					return err
				}
			}
		}
	}

	// Buat beberapa promo usages untuk demo
	// if len(customers) > 0 && len(promotions) > 0 {
	// 	// Customer 1 menggunakan Welcome Voucher
	// 	promoUsage1 := entity.PromoUsage{
	// 		PromotionID: promotions[1].ID, // Welcome Voucher
	// 		CustomerID:  customers[0].ID,
	// 		OrderID:     1, // Asumsi order ID
	// 		UsedAt:      now.Add(-10 * time.Hour),
	// 	}
	// 	if err := s.DB.Create(&promoUsage1).Error; err != nil {
	// 		return err
	// 	}

	// 	// Customer 2 menggunakan Member Exclusive
	// 	if len(customers) > 1 {
	// 		promoUsage2 := entity.PromoUsage{
	// 			PromotionID: promotions[6].ID, // Member Exclusive
	// 			CustomerID:  customers[1].ID,
	// 			OrderID:     2,
	// 			UsedAt:      now.Add(-5 * time.Hour),
	// 		}
	// 		if err := s.DB.Create(&promoUsage2).Error; err != nil {
	// 			return err
	// 		}
	// 	}

	// 	// Customer 3 menggunakan Free Shipping
	// 	if len(customers) > 2 {
	// 		promoUsage3 := entity.PromoUsage{
	// 			PromotionID: promotions[5].ID, // Free Shipping
	// 			CustomerID:  customers[2].ID,
	// 			OrderID:     3,
	// 			UsedAt:      now.Add(-2 * time.Hour),
	// 		}
	// 		if err := s.DB.Create(&promoUsage3).Error; err != nil {
	// 			return err
	// 		}
	// 	}

	// 	// Customer 4 menggunakan Grand Opening Sale
	// 	if len(customers) > 3 {
	// 		promoUsage4 := entity.PromoUsage{
	// 			PromotionID: promotions[0].ID, // Grand Opening
	// 			CustomerID:  customers[3].ID,
	// 			OrderID:     4,
	// 			UsedAt:      now.Add(-1 * time.Hour),
	// 		}
	// 		if err := s.DB.Create(&promoUsage4).Error; err != nil {
	// 			return err
	// 		}
		// }
	// }

	return nil
}

func stringPtr(s string) *string {
	return &s
}
