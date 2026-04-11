package seeder

import (
	"debian-ecommerce/internal/data/entity"

	"gorm.io/gorm"
)

type CartDetailSeeder struct {
	DB *gorm.DB
}

func NewCartDetailSeeder(db *gorm.DB) *CartDetailSeeder {
	return &CartDetailSeeder{DB: db}
}

func (s *CartDetailSeeder) Run() error {
	var carts []entity.Cart
	if err := s.DB.Find(&carts).Error; err != nil {
		return err
	}

	if len(carts) > 0 {
		return nil // Sudah ada cart, skip seeding
	}

	// Cari customer aktif
	var customers []entity.Customer
	if err := s.DB.Preload("User", "is_active = ?", true).Limit(5).Find(&customers).Error; err != nil {
		return err
	}

	// Cari beberapa SKU yang berbeda-beda untuk demo
	var (
		cheapSKUs      []entity.SKU
		midSKUs        []entity.SKU
		expensiveSKUs  []entity.SKU
		fashionSKUs    []entity.SKU
		electronicSKUs []entity.SKU
	)

	// Query SKU berdasarkan kategori/kriteria
	s.DB.Joins("JOIN products ON products.id = skus.product_id").
		Joins("JOIN categories ON categories.id = products.category_id").
		Where("skus.status = ? AND skus.stock > ? AND categories.name IN ?",
			entity.SKUStatusActive, 0, []string{"Personal Care", "Home & Cleaning", "Straw & Cutlery"}).
		Preload("Product").
		Preload("Product.Category").
		Find(&cheapSKUs).
		Where("skus.price < ?", 500000)

	s.DB.Where("status = ? AND stock > ? AND price BETWEEN ? AND ?",
		entity.SKUStatusActive, 0, 500000, 2000000).
		Preload("Product").
		Limit(10).
		Find(&midSKUs)

	s.DB.Where("status = ? AND stock > ? AND price > ?",
		entity.SKUStatusActive, 0, 3000000).
		Preload("Product").
		Limit(5).
		Find(&expensiveSKUs)

	// Buat cart untuk setiap customer dengan item yang bermakna
	for i, customer := range customers {
		cart := entity.Cart{
			CustomerID: customer.ID,
		}

		if err := s.DB.Create(&cart).Error; err != nil {
			return err
		}

		// Tambahkan item berdasarkan tipe customer
		switch i {
		case 0: // Customer 1: Mixed cart (elektronik + fashion)
			s.addCartItem(cart.ID, electronicSKUs, 1, 2)
			s.addCartItem(cart.ID, fashionSKUs, 0, 1)

		case 1: // Customer 2: Electronics enthusiast
			s.addCartItem(cart.ID, electronicSKUs, 0, 3)
			if len(expensiveSKUs) > 0 {
				s.addCartItem(cart.ID, expensiveSKUs, 0, 1)
			}

		case 2: // Customer 3: Budget shopper
			s.addCartItem(cart.ID, cheapSKUs, 0, 4)

		case 3: // Customer 4: Single expensive item
			if len(expensiveSKUs) > 0 {
				s.addCartItem(cart.ID, expensiveSKUs, 0, 1)
			}

		case 4: // Customer 5: Multiple mid-range items
			s.addCartItem(cart.ID, midSKUs, 0, 3)
		}
	}

	return nil
}

func (s *CartDetailSeeder) addCartItem(cartID uint, skus []entity.SKU, startIndex, maxItems int) {
	count := 0
	for i := startIndex; i < len(skus) && count < maxItems; i++ {
		// Hitung quantity yang wajar berdasarkan harga
		quantity := 1
		if skus[i].Price < 200000 {
			quantity = 2
		}

		// Pastikan tidak melebihi stock
		if quantity > skus[i].Stock {
			quantity = skus[i].Stock
		}

		if quantity > 0 {
			cartItem := entity.CartItem{
				CartID:   cartID,
				SKUID:    skus[i].ID,
				Quantity: quantity,
			}
			s.DB.Create(&cartItem)
			count++
		}
	}
}
