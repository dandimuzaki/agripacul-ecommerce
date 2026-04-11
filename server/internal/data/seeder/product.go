package seeder

import (
	"fmt"
	"strings"

	"debian-ecommerce/internal/data/entity"

	"gorm.io/gorm"
)

type ProductSeeder struct {
	DB *gorm.DB
}

func NewProductSeeder(db *gorm.DB) *ProductSeeder {
	return &ProductSeeder{DB: db}
}

func (s *ProductSeeder) Run() error {
	var products []entity.Product
	if err := s.DB.Find(&products).Error; err != nil {
		return err
	}

	if len(products) >= 6 {
		return nil // Sudah ada product, skip seeding
	}

	if err := SeedReusableBags(s.DB); err != nil {
		return err
	}

	if err := SeedReusableBags(s.DB); err != nil {
		return err
	}

	if err := SeedComposterCollection(s.DB); err != nil {
		return err
	}

	if err := SeedLerakCollection(s.DB); err != nil {
		return err
	}

	if err := SeedLoofahCollection(s.DB); err != nil {
		return err
	}

	if err := SeedTrayCollection(s.DB); err != nil {
		return err
	}

	if err := SeedStrawCollection(s.DB); err != nil {
		return err
	}

		// if err := SeedReusableProduceBags(s.DB); err != nil {
		// 	return err
		// }
		// if err := SeedHomeComposter(s.DB); err != nil {
		// 	return err
		// }
		// if err := SeedBambooCutlerySet(s.DB); err != nil {
		// 	return err
		// }
		// if err := SeedEcoEnzyme(s.DB); err != nil {
		// 	return err
		// }
		// if err := SeedCompostableCuttingResistantTrays(s.DB); err != nil {
		// 	return err
		// }
		// if err := SeedZeroWasteStarterKit(s.DB); err != nil {
		// 	return err
		// }
		// if err := SeedBuahLerak(s.DB); err != nil {
		// 	return err
		// }
		// if err := SeedNaturalLoofahSponge(s.DB); err != nil {
		// 	return err
		// }
		// if err := SeedPalmLeafDinnerwareSet(s.DB); err != nil {
		// 	return err
		// }
		// if err := SeedStainlessSteelStrawSet(s.DB); err != nil {
		// 	return err
		// }
		// if err := SeedBambooToothbrush(s.DB); err != nil {
		// 	return err
		// }

	// // Ambil semua kategori untuk referensi
	// var categories []entity.Category
	// if err := s.DB.Find(&categories).Error; err != nil {
	// 	return err
	// }

	// if len(categories) == 0 {
	// 	return nil // Tidak ada kategori, skip seeding
	// }

	// // Data produk dengan kategori acak
	// products = []entity.Product{
	// 	{
	// 		CategoryID:  categories[0].ID, // Food Packaging
	// 		Name:        "Compostable Molded Bamboo pulp Lunch Box Bottom With T-buckle",
	// 		Description: "An eco-friendly lunch box made from molded bamboo pulp, designed to be fully compostable and plastic-free. The sturdy structure provides excellent resistance to heat and moisture, making it suitable for hot and cold foods. Equipped with a secure T-buckle system, the box helps prevent spills during transport. Ideal for takeaways, catering services, and sustainable food businesses. A responsible packaging choice that supports a greener environment.",
	// 		IsPublished: true,
	// 		Tags:        pq.StringArray{"lunch", "box", "bamboo"},
	// 		Slug: utils.GenerateSlug("Compostable Molded Bamboo pulp Lunch Box Bottom With T-buckle"),
	// 	},
	// 	{
	// 		CategoryID:  categories[0].ID, // Food Packaging
	// 		Name:        "Compostable Molded Pulp 5-Compartment Tray Lid",
	// 		Description: "This compostable tray lid is made from molded pulp and designed to fit 5-compartment food trays perfectly. It helps keep food portions separated, fresh, and hygienic during storage and delivery. The material is biodegradable and environmentally friendly, offering a sustainable alternative to plastic lids. Strong yet lightweight, it is suitable for restaurants, meal prep services, and eco-conscious catering. A practical solution for organized and sustainable food packaging.",
	// 		IsPublished: true,
	// 		Tags:        pq.StringArray{"tray", "lid", "packaging"},
	// 		Slug: utils.GenerateSlug("Compostable Molded Pulp 5-Compartment Tray Lid"),
	// 	},
	// 	{
	// 		CategoryID:  categories[1].ID, // Women Cares
	// 		Name:        "Baby Oz Reusable Cloth Menstrual Pad – Anti-Itch & Leak-Proof",
	// 		Description: "The Baby Oz reusable cloth menstrual pad is designed for comfort, safety, and sustainability. Made from soft, breathable fabric, it helps prevent itching and irritation even during long wear. The leak-proof inner layer provides reliable protection throughout the day. Washable and reusable, this menstrual pad reduces waste and is gentle on sensitive skin. A healthy and eco-friendly alternative to disposable pads.",
	// 		IsPublished: true,
	// 		Tags:        pq.StringArray{"women", "menstrual", "pad", "reusable"},
	// 		Slug: utils.GenerateSlug("Baby Oz Reusable Cloth Menstrual Pad – Anti-Itch & Leak-Proof"),
	// 	},
	// 	{
	// 		CategoryID:  categories[1].ID, // Women Cares
	// 		Name:        "Sustaination Reusable Cloth Pantyliner 19 cm (Pack of 3) – Eco-Friendly & Leak-Resistant",
	// 		Description: "This reusable cloth pantyliner set from Sustaination offers daily protection while caring for the environment. Measuring 19 cm, each pantyliner is designed to absorb moisture effectively and prevent leaks. The breathable fabric helps keep the skin dry and comfortable all day long. Easy to wash and reuse, this pack of 3 is perfect for daily wear. An eco-friendly choice for a healthier and more sustainable lifestyle.",
	// 		IsPublished: true,
	// 		Tags:        pq.StringArray{"women", "pantyliner", "cloth"},
	// 		Slug: utils.GenerateSlug("Sustaination Reusable Cloth Pantyliner 19 cm (Pack of 3) – Eco-Friendly & Leak-Resistant"),
	// 	},
	// 	{
	// 		CategoryID:  categories[2].ID, // Gardening Tools
	// 		Name:        "3-in-1 Gardening Tools Set (Shovel, Hoe, Fork) – Large Size (22 × 6 cm)",
	// 		Description: "This complete 3-in-1 gardening tool set includes a shovel, hoe, and fork for versatile garden use. Designed in a large size (22 × 6 cm), the tools are strong and comfortable to handle. Suitable for digging, loosening soil, planting, and general garden maintenance. Ideal for home gardening, urban farming, and outdoor planting activities. A practical and durable set for both beginners and experienced gardeners.",
	// 		IsPublished: false, // Belum dipublish
	// 		Tags:        pq.StringArray{"gardening", "shovel", "hoe", "fork"},
	// 		Slug: utils.GenerateSlug("3-in-1 Gardening Tools Set (Shovel, Hoe, Fork) – Large Size (22 × 6 cm)"),
	// 	},
	// 	{
	// 		CategoryID:  categories[2].ID, // Olahraga & Outdoor
	// 		Name:        "Sustaination Composting Bucket – Household Organic Waste Solution",
	// 		Description: "The Sustaination composting bucket is a practical solution for managing household organic waste. Designed for easy composting at home, it helps turn kitchen waste into nutrient-rich compost. The bucket reduces unpleasant odors and supports cleaner waste management. Suitable for apartments and houses, it encourages an eco-friendly lifestyle. A simple yet effective way to reduce waste and support sustainable living.",
	// 		IsPublished: true,
	// 		Tags:        pq.StringArray{"composting", "bucket", "organic"},
	// 		Slug: utils.GenerateSlug("Sustaination Composting Bucket – Household Organic Waste Solution"),
	// 	},
	// }

	// // Hapus data lama (dengan cascade delete)
	// s.DB.Exec("DELETE FROM products CASCADE")

	// // Create products
	// for i := range products {
	// 	if err := s.DB.Create(&products[i]).Error; err != nil {
	// 		return err
	// 	}
	// }

	// // Populate variants
	// s.seedFoodPackaging(&products[0])
	// s.seedFoodPackaging(&products[1])
	// s.seedWomenCares(&products[2])
	// s.seedWomenCares(&products[3])
	// s.seedToolsVariants(&products[4])
	// s.seedToolsVariants(&products[5])

	return nil
}

func (s *ProductSeeder) seedFoodPackaging(product *entity.Product) error {
	// Variant types untuk food packaging
	colorType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Color",
	}
	if err := s.DB.Create(&colorType).Error; err != nil {
		return err
	}

	// Variant values untuk warna
	colorValues := []entity.VariantValue{
		{VariantTypeID: colorType.ID, Value: "Yellow"},
		{VariantTypeID: colorType.ID, Value: "Pink"},
		{VariantTypeID: colorType.ID, Value: "Blue"},
	}
	for i := range colorValues {
		if err := s.DB.Create(&colorValues[i]).Error; err != nil {
			return err
		}
	}

	// Variant types untuk food packaging
	sizeType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Size",
	}
	if err := s.DB.Create(&sizeType).Error; err != nil {
		return err
	}

	// Variant values untuk warna
	sizeValues := []entity.VariantValue{
		{VariantTypeID: sizeType.ID, Value: "700 ml"},
		{VariantTypeID: sizeType.ID, Value: "800 ml"},
	}
	for i := range sizeValues {
		if err := s.DB.Create(&sizeValues[i]).Error; err != nil {
			return err
		}
	}

	// Buat SKUs untuk semua kombinasi
	skuCounter := 1
	for _, color := range colorValues {
		for _, storage := range sizeValues {
			sku := entity.SKU{
				ProductID: product.ID,
				SKUCode:   fmt.Sprintf("MI-RN12-%03d", skuCounter),
				Price:     2999000.00,
				Stock:     50,
			}
			if err := s.DB.Create(&sku).Error; err != nil {
				return err
			}

			// Link variant values ke SKU
			skuVariant1 := entity.SKUVariantValue{
				SKUID:          sku.ID,
				VariantValueID: color.ID,
			}
			if err := s.DB.Create(&skuVariant1).Error; err != nil {
				return err
			}

			skuVariant2 := entity.SKUVariantValue{
				SKUID:          sku.ID,
				VariantValueID: storage.ID,
			}
			if err := s.DB.Create(&skuVariant2).Error; err != nil {
				return err
			}

			// Tambahkan gambar untuk SKU
			skuImages := []entity.Image{
				{ProductID: sku.ProductID, SKUID: &sku.ID, ImageURL: fmt.Sprintf("https://example.com/products/mi-rn12-%s-1.jpg", strings.ToLower(color.Value))},
				{ProductID: sku.ProductID, SKUID: &sku.ID, ImageURL: fmt.Sprintf("https://example.com/products/mi-rn12-%s-2.jpg", strings.ToLower(color.Value))},
			}
			for i := range skuImages {
				if err := s.DB.Create(&skuImages[i]).Error; err != nil {
					return err
				}
			}

			skuCounter++
		}
	}

	return nil
}

func (s *ProductSeeder) seedWomenCares(product *entity.Product) error {
	// Variant types untuk women cares
	colorType := entity.VariantType{
		ProductID: product.ID,
		Name:      "color",
	}
	if err := s.DB.Create(&colorType).Error; err != nil {
		return err
	}

	// Variant values untuk warna
	colorValues := []entity.VariantValue{
		{VariantTypeID: colorType.ID, Value: "Yellow"},
		{VariantTypeID: colorType.ID, Value: "Pink"},
		{VariantTypeID: colorType.ID, Value: "Blue"},
	}
	for i := range colorValues {
		if err := s.DB.Create(&colorValues[i]).Error; err != nil {
			return err
		}
	}

	// Buat SKUs untuk semua kombinasi
	skuCounter := 1
	for _, col := range colorValues {
		price := 8999000.00
		sku := entity.SKU{
			ProductID: product.ID,
			SKUCode:   fmt.Sprintf("WC-%03d", skuCounter),
			Price:     price,
			Stock:     10,
		}
		if err := s.DB.Create(&sku).Error; err != nil {
			return err
		}

		// Link variant values ke SKU
		skuVariant1 := entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: col.ID,
		}
		if err := s.DB.Create(&skuVariant1).Error; err != nil {
			return err
		}

		// Tambahkan gambar untuk SKU
		skuImages := []entity.Image{
			{ProductID: sku.ProductID, SKUID: &sku.ID, ImageURL: "https://example.com/products/women-cares-1.jpg"},
			{ProductID: sku.ProductID, SKUID: &sku.ID, ImageURL: "https://example.com/products/women-cares-2.jpg"},
			{ProductID: sku.ProductID, SKUID: &sku.ID, ImageURL: "https://example.com/products/women-cares-3.jpg"},
		}
		for i := range skuImages {
			if err := s.DB.Create(&skuImages[i]).Error; err != nil {
				return err
			}
		}

		skuCounter++
	}

	return nil
}

func (s *ProductSeeder) seedToolsVariants(product *entity.Product) error {
	// Variant types untuk gardening tools
	colorType := entity.VariantType{
		ProductID: product.ID,
		Name:      "color",
	}
	if err := s.DB.Create(&colorType).Error; err != nil {
		return err
	}

	// Variant values untuk warna
	colorValues := []entity.VariantValue{
		{VariantTypeID: colorType.ID, Value: "Blue"},
		{VariantTypeID: colorType.ID, Value: "Black"},
	}
	for i := range colorValues {
		if err := s.DB.Create(&colorValues[i]).Error; err != nil {
			return err
		}
	}

	// Buat SKUs untuk semua kombinasi
	skuCounter := 1
	for _, color := range colorValues {
		stock := 100

		sku := entity.SKU{
			ProductID: product.ID,
			SKUCode:   fmt.Sprintf("KEM-OXF-%03d", skuCounter),
			Price:     249000.00,
			Stock:     stock,
		}
		if err := s.DB.Create(&sku).Error; err != nil {
			return err
		}

		// Link variant values ke SKU
		skuVariant1 := entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: color.ID,
		}
		if err := s.DB.Create(&skuVariant1).Error; err != nil {
			return err
		}

		// Tambahkan gambar untuk SKU
		skuImages := []entity.Image{
			{ProductID: sku.ProductID, SKUID: &sku.ID, ImageURL: fmt.Sprintf("https://example.com/products/kemeja-%s-1.jpg", strings.ToLower(color.Value))},
			{ProductID: sku.ProductID, SKUID: &sku.ID, ImageURL: fmt.Sprintf("https://example.com/products/kemeja-%s-2.jpg", strings.ToLower(color.Value))},
		}
		for i := range skuImages {
			if err := s.DB.Create(&skuImages[i]).Error; err != nil {
				return err
			}
		}

		skuCounter++
	}

	return nil
}