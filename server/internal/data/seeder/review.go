package seeder

import (
	"math/rand"
	"time"

	"debian-ecommerce/internal/data/entity"

	"gorm.io/gorm"
)

type ReviewSeeder struct {
	DB *gorm.DB
}

func NewReviewSeeder(db *gorm.DB) *ReviewSeeder {
	rand.Seed(time.Now().UnixNano())
	return &ReviewSeeder{DB: db}
}

func (s *ReviewSeeder) Run() error {
	var reviews []entity.Review
	if err := s.DB.Find(&reviews).Error; err != nil {
		return err
	}

	if len(reviews) > 0 {
		return nil // Sudah ada review, skip seeding
	}

	// Ambil data yang diperlukan
	var customers []entity.Customer
	if err := s.DB.Preload("User", "is_active = ?", true).
		Where("users.is_active = ?", true).
		Joins("JOIN users ON users.id = customers.user_id").
		Limit(10).
		Find(&customers).Error; err != nil {
		return err
	}

	// Query orders dengan join ke SKU dan Product
	var orders []entity.Order
	if err := s.DB.
		Where("status = ?", entity.OrderStatusCompleted).
		Preload("OrderItems", func(db *gorm.DB) *gorm.DB {
			return db.Joins("JOIN skus ON skus.id = order_items.sku_id").
				Joins("JOIN products ON products.id = skus.product_id").
				Select("order_items.*, skus.product_id as product_id")
		}).
		Limit(20).
		Find(&orders).Error; err != nil {
		return err
	}

	if len(customers) == 0 || len(orders) == 0 {
		return nil
	}

	// Hapus data lama
	if err := s.DB.Exec("DELETE FROM reviews").Error; err != nil {
		return err
	}

	// Kumpulan komentar review
	positiveComments := []string{
		"Produk sangat berkualitas, sesuai dengan deskripsi.",
		"Pengiriman cepat, barang sampai dengan baik.",
		"Kualitas produk melebihi ekspektasi, sangat puas!",
		"Barang original, packing rapi, seller recommended.",
		"Warna sesuai foto, bahan bagus, cocok untuk daily use.",
		"Fitur lengkap, performa maksimal, worth it untuk harganya.",
		"Customer service responsif, sangat membantu.",
		"Kemasan aman, barang tidak rusak saat sampai.",
		"Bahan nyaman dipakai, model stylish, suka banget!",
		"Fungsionalitas sesuai kebutuhan, mudah digunakan.",
	}

	neutralComments := []string{
		"Produk cukup baik, sesuai dengan harga.",
		"Pengiriman agak lambat tapi barang oke.",
		"Kualitas standar, tidak ada yang istimewa.",
		"Barang sesuai pesanan, warna sedikit berbeda dengan foto.",
		"Ukuran pas, bahan lumayan, harga wajar.",
		"Fitur sesuai, performa cukup untuk kebutuhan dasar.",
		"Kemasan biasa saja, barang aman sampai.",
		"Bahan nyaman, model simple, cocok untuk casual.",
		"Fungsionalitas oke, ada sedikit kendala kecil.",
		"Pengalaman belanja biasa-biasa saja.",
	}

	negativeComments := []string{
		"Barang rusak saat sampai, sangat kecewa.",
		"Kualitas di bawah ekspektasi, bahan tipis.",
		"Warna sangat berbeda dengan foto, tidak sesuai.",
		"Ukuran tidak pas, terlalu kecil/kecil.",
		"Fitur tidak berfungsi dengan baik, sering error.",
		"Pengiriman sangat lambat, tidak sesuai estimasi.",
		"Customer service tidak responsif, sulit dihubungi.",
		"Kemasan jelek, barang berisiko rusak.",
		"Bahan tidak nyaman, mudah rusak.",
		"Tidak sesuai deskripsi, kualitas buruk.",
	}

	// Buat review untuk beberapa order yang sudah completed
	for _, order := range orders {
		// Hanya buat review untuk 50% order yang completed
		if rand.Intn(100) < 50 {
			continue
		}

		// Pastikan order memiliki items
		if len(order.Items) == 0 {
			continue
		}

		// Cari customer yang sesuai
		var customer entity.Customer
		for _, c := range customers {
			if c.ID == order.CustomerID {
				customer = c
				break
			}
		}

		if customer.ID == 0 {
			continue
		}

		// Ambil product_id untuk setiap order item
		orderItemsWithProductID := s.getOrderItemsWithProductID(order.ID)
		if len(orderItemsWithProductID) == 0 {
			continue
		}

		// Buat review untuk beberapa item dalam order
		itemsToReview := 1
		if len(orderItemsWithProductID) > 1 {
			itemsToReview = rand.Intn(len(orderItemsWithProductID)) + 1
		}

		reviewedItems := make(map[uint]bool)

		for i := 0; i < itemsToReview; i++ {
			// Pilih item random yang belum direview
			var selectedItem struct {
				OrderItemID uint
				ProductID   uint
				UnitPrice   float64
			}

			for {
				idx := rand.Intn(len(orderItemsWithProductID))
				selectedItem = orderItemsWithProductID[idx]
				if !reviewedItems[selectedItem.OrderItemID] {
					reviewedItems[selectedItem.OrderItemID] = true
					break
				}
			}

			// Tentukan rating berdasarkan harga produk
			var rating int
			comment := ""

			// Logika rating berdasarkan harga
			if selectedItem.UnitPrice < 100000 {
				rating = rand.Intn(2) + 4 // 4-5
				comment = positiveComments[rand.Intn(len(positiveComments))]
			} else if selectedItem.UnitPrice < 500000 {
				rating = rand.Intn(3) + 3 // 3-5
				switch rating {
				case 5, 4:
					comment = positiveComments[rand.Intn(len(positiveComments))]
				case 3:
					comment = neutralComments[rand.Intn(len(neutralComments))]
				}
			} else {
				rating = rand.Intn(3) + 2 // 2-4
				switch rating {
				case 4:
					comment = positiveComments[rand.Intn(len(positiveComments))]
				case 3:
					comment = neutralComments[rand.Intn(len(neutralComments))]
				case 2:
					comment = negativeComments[rand.Intn(len(negativeComments))]
				}
			}

			// Sesekali berikan rating 1 untuk realism
			if rand.Intn(100) < 5 {
				rating = 1
				comment = negativeComments[rand.Intn(len(negativeComments))]
			}

			review := entity.Review{
				CustomerID: customer.ID,
				ProductID:  selectedItem.ProductID,
				OrderID:    order.ID,
				Rating:     rating,
				Comment:    comment,
			}

			if err := s.DB.Create(&review).Error; err != nil {
				return err
			}
		}
	}

	// Buat review khusus untuk demo
	s.createDemoReviews(customers, orders)

	return nil
}

// Helper function untuk mendapatkan order items dengan product_id
func (s *ReviewSeeder) getOrderItemsWithProductID(orderID uint) []struct {
	OrderItemID uint
	ProductID   uint
	UnitPrice   float64
} {
	var results []struct {
		OrderItemID uint
		ProductID   uint
		UnitPrice   float64
	}

	s.DB.Table("order_items").
		Select("order_items.id as order_item_id, skus.product_id as product_id, order_items.unit_price").
		Joins("JOIN skus ON skus.id = order_items.sku_id").
		Where("order_items.order_id = ?", orderID).
		Scan(&results)

	return results
}

func (s *ReviewSeeder) createDemoReviews(customers []entity.Customer, orders []entity.Order) error {
	if len(customers) < 3 || len(orders) < 3 {
		return nil
	}

	// Untuk demo, kita butuh product_id
	for i := 0; i < 5 && i < len(orders); i++ {
		order := orders[i]

		// Ambil satu product dari order ini
		orderItems := s.getOrderItemsWithProductID(order.ID)
		if len(orderItems) == 0 {
			continue
		}

		productID := orderItems[0].ProductID
		customerIndex := i
		if customerIndex >= len(customers) {
			customerIndex = len(customers) - 1
		}

		// Demo reviews dengan rating berbeda
		demoReviews := []struct {
			rating  int
			comment string
		}{
			{
				rating:  5,
				comment: "Sangat puas dengan produk ini! Kualitasnya luar biasa, sesuai dengan deskripsi dan harganya worth it. Pengiriman juga cepat dan packing aman. Seller sangat recommended, akan belanja lagi di sini. Terima kasih!",
			},
			{
				rating:  1,
				comment: "Sangat kecewa! Barang rusak saat sampai, kualitas buruk, tidak sesuai deskripsi sama sekali. Pengiriman lambat dan customer service tidak responsif. Uang saya terbuang percuma.",
			},
			{
				rating:  3,
				comment: "Produk cukup baik untuk harganya, tapi ada beberapa hal yang perlu ditingkatkan. Kualitas bahan bisa lebih baik lagi, pengiriman agak lambat. Overall oke untuk pemakaian sehari-hari.",
			},
			{
				rating:  2,
				comment: "Warna tidak sesuai dengan foto di website, lebih gelap dan kurang menarik. Bahan juga tipis dan mudah kusut. Namun pengiriman cukup cepat dan packing rapi. Harusnya bisa lebih baik untuk harga segini.",
			},
			{
				rating:  4,
				comment: "Kualitas produk sangat bagus, bahan premium dan nyaman dipakai. Desain stylish dan sesuai tren. Hanya saja ukuran agak kecil, mungkin perlu ditambah size chart yang lebih akurat. Overall sangat recommended!",
			},
		}

		for j := 0; j < len(demoReviews) && j < len(customers); j++ {
			review := entity.Review{
				CustomerID: customers[j].ID,
				ProductID:  productID,
				OrderID:    order.ID,
				Rating:     demoReviews[j].rating,
				Comment:    demoReviews[j].comment,
			}
			s.DB.Create(&review)
		}
	}

	return nil
}
