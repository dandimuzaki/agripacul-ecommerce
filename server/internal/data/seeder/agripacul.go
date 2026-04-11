package seeder

import (
	"debian-ecommerce/internal/data/entity"
	"log"
	"math/rand"

	"gorm.io/gorm"
)

type AgripaculProductSeeder struct {
	db *gorm.DB
}

func NewAgripaculProductSeeder(db *gorm.DB) *AgripaculProductSeeder {
	return &AgripaculProductSeeder{db: db}
}

func (s *AgripaculProductSeeder) Seed() error {
	// Seed Categories first
	if err := s.seedCategories(); err != nil {
		return err
	}

	// Seed Products
	if err := s.seedProducts(); err != nil {
		return err
	}

	return nil
}

func (s *AgripaculProductSeeder) seedCategories() error {
	var categories []entity.Category
	if err := s.db.Find(&categories).Error; err != nil {
		return err
	}

	if len(categories) > 0 {
		return nil // Sudah ada categories, skip seeding
	}

	categories = []entity.Category{
		{
			Name:       "Fresh Vegetables",
			IconURL:    "https://res.cloudinary.com/agripacul/image/upload/v1/categories/fresh-vegetables",
			IconPublicID: "categories/fresh-vegetables",
		},
		{
			Name:       "Ready to Eat",
			IconURL:    "https://res.cloudinary.com/agripacul/image/upload/v1/categories/makanan-ready to eat",
			IconPublicID: "categories/makanan-ready to eat",
		},
		{
			Name:       "Homemade Foods",
			IconURL:    "https://res.cloudinary.com/agripacul/image/upload/v1/categories/makanan-fermentation",
			IconPublicID: "categories/makanan-fermentation",
		},
		{
			Name:       "Seeds",
			IconURL:    "https://res.cloudinary.com/agripacul/image/upload/v1/categories/benih",
			IconPublicID: "categories/benih",
		},
		{
			Name:       "Gardening Tools",
			IconURL:    "https://res.cloudinary.com/agripacul/image/upload/v1/categories/peralatan-kebun",
			IconPublicID: "categories/peralatan-kebun",
		},
	}

	for _, category := range categories {
		if err := s.db.FirstOrCreate(&category, entity.Category{Name: category.Name}).Error; err != nil {
			return err
		}
	}

	log.Println("Categories seeded successfully")
	return nil
}

func (s *AgripaculProductSeeder) seedProducts() error {
	var products []entity.Product
	if err := s.db.Find(&products).Error; err != nil {
		return err
	}

	if len(products) > 0 {
		return nil // Sudah ada products, skip seeding
	}

	// Get category IDs
	var freshVegetables entity.Category
	var readyToEat entity.Category
	var homemadeFoods entity.Category
	var seeds entity.Category
	var gardeningTools entity.Category

	if err := s.db.Where("name = ?", "Fresh Vegetables").First(&freshVegetables).Error; err != nil {
		return err
	}

	if err := s.db.Where("name = ?", "Ready to Eat").First(&readyToEat).Error; err != nil {
		return err
	}

	if err := s.db.Where("name = ?", "Homemade Foods").First(&homemadeFoods).Error; err != nil {
		return err
	}

	if err := s.db.Where("name = ?", "Seeds").First(&seeds).Error; err != nil {
		return err
	}

	if err := s.db.Where("name = ?", "Gardening Tools").First(&gardeningTools).Error; err != nil {
		return err
	}

	// Define products
	products = []entity.Product{
		// Pakcoy
		{
			CategoryID:      freshVegetables.ID,
			Name:            "Organic Pakcoy",
			Description:     "Fresh, crisp, and naturally vibrant, our pakcoy is harvested at peak maturity to ensure optimal flavor and nutrition. With its tender stems and rich green leaves, pakcoy is perfect for stir-fries, soups, or as a healthy side dish. Grown with care, it delivers a mild, slightly sweet taste and a satisfying crunch in every bite. A staple for balanced, wholesome meals.",
			IsPublished:     true,
			Tags:           []string{"organic", "pakcoy", "fresh vegetables"},
			Slug:           "organic-pakcoy",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/pakcoy/main",
			MainImagePublicID: "products/pakcoy/main",
		},
		// Hidroponic Pakcoy
		{
			CategoryID:      freshVegetables.ID,
			Name:            "Hidroponic Pakcoy",
			Description:     `Grown using modern hydroponic systems, this pakcoy is cultivated without soil in a controlled, clean environment. The result is a fresher, more hygienic vegetable with consistent quality and a longer shelf life. Its crisp texture and pure taste make it ideal for health-conscious consumers who value both flavor and sustainability.`,
			IsPublished:     true,
			Tags:           []string{"organic", "pakcoy", "hidroponic", "fresh vegetables"},
			Slug:           "pakcoy-hidroponic-fresh",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/pakcoy/main",
			MainImagePublicID: "products/pakcoy/main",
		},
		// Cherry Tomato
		{
			CategoryID:      freshVegetables.ID,
			Name:            "Cherry Tomatoes",
			Description:     `Our cherry tomatoes are small bursts of sweetness, offering a juicy, refreshing taste in every bite. Perfect for readyToEat, snacks, or garnishes, they bring color and flavor to any dish. Carefully grown and handpicked, these tomatoes are rich in vitamins and antioxidants, making them as nutritious as they are delicious.`,
			IsPublished:     true,
			Tags:           []string{"cherry-tomato", "fresh vegetables", "premium", "antioxidant"},
			Slug:           "cherry-tomato-premium",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/cherry-tomato/main",
			MainImagePublicID: "products/cherry-tomato/main",
		},
		// Red Tomato
		{
			CategoryID:      freshVegetables.ID,
			Name:            "Local Red Tomatoes",
			Description:     `Plump, juicy, and full of natural goodness, our red tomatoes are a kitchen essential. Their balanced flavor—slightly sweet with a hint of acidity—makes them perfect for sauces, soups, and fresh dishes. Harvested at peak ripeness, they provide both taste and nutrition for everyday cooking.`,
			IsPublished:     true,
			Tags:           []string{"red tomatoes", "fruit vegetables", "local", "fresh vegetables"},
			Slug:           "local-red-tomatoes",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/red-tomatoes/main",
			MainImagePublicID: "products/red-tomatoes/main",
		},
		// Sweet Corn
		{
			CategoryID:      freshVegetables.ID,
			Name:            "Sweet Corn",
			Description:     `Naturally sweet and juicy, our sweet corn delivers a delightful burst of flavor in every kernel. Whether grilled, boiled, or added to dishes, it provides both taste and texture. Freshly harvested to retain its sweetness, it’s a favorite for all ages.`,
			IsPublished:     true,
			Tags:           []string{"sweet corn", "fresh vegetables"},
			Slug:           "sweet-corn-hibrida",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/sweet-corn/main",
			MainImagePublicID: "products/sweet-corn/main",
		},
		// Green Beans
		{
			CategoryID:      freshVegetables.ID,
			Name:            "Green Beans (Buncis)",
			Description:     `Crisp and tender, our green beans are freshly harvested to preserve their natural flavor and nutrients. They are perfect for stir-fries, steaming, or sautéing, adding both texture and color to your meals. A simple yet essential vegetable for healthy, home-cooked dishes.`,
			IsPublished:     true,
			Tags:           []string{"green beans", "fresh vegetables"},
			Slug:           "fresh-green-beans",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/green-beans/main",
			MainImagePublicID: "products/green-beans/main",
		},
		// Cabbage
		{
			CategoryID:      freshVegetables.ID,
			Name:            "Fresh Cabbage",
			Description:     `Fresh, firm, and full of natural goodness, our cabbage is carefully cultivated to deliver both quality and flavor in every layer. With its tightly packed leaves and satisfying crunch, this versatile vegetable is a staple in kitchens around the world. Whether sliced for fresh salads, stir-fried for a quick meal, or slow-cooked into hearty dishes, cabbage adapts beautifully to a wide range of recipes.`,
			IsPublished:     true,
			Tags:           []string{"cabbage", "fresh vegetables"},
			Slug:           "fresh-cabbage",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/cabbage/main",
			MainImagePublicID: "products/cabbage/main",
		},
		// Broccoli
		{
			CategoryID:      freshVegetables.ID,
			Name:            "Fresh Green Broccoli",
			Description:     `Our broccoli is carefully grown to produce dense, vibrant green florets packed with nutrients. With its slightly earthy taste and satisfying crunch, it is perfect for steaming, roasting, or adding to stir-fries. A powerhouse vegetable for those who prioritize healthy eating.`,
			IsPublished:     true,
			Tags:           []string{"broccoli", "fresh vegetables", "superfood"},
			Slug:           "fresh-green-broccoli",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/broccoli/main",
			MainImagePublicID: "products/broccoli/main",
		},
		// Long beans
		{
			CategoryID:      freshVegetables.ID,
			Name:            "Long Beans",
			Description:     `Known for their length and distinctive taste, long beans are a versatile ingredient in many traditional dishes. They offer a slightly earthy flavor and a firm texture that holds up well in cooking. Freshly picked, they bring authenticity and nutrition to your kitchen.`,
			IsPublished:     true,
			Tags:           []string{"long beans", "fresh vegetables", "superfood"},
			Slug:           "long-beans",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/long-beans/main",
			MainImagePublicID: "products/long-beans/main",
		},
		// Bird's eye chilli
		{
			CategoryID:      freshVegetables.ID,
			Name:            "Bird's Eye Chili (Cabai Rawit)",
			Description:     `Small but powerful, these chilies bring intense heat and bold flavor to your cooking. Perfect for sambals, sauces, and spicy dishes, they are a must-have for those who love a fiery kick. Fresh and aromatic, they elevate any dish instantly.`,
			IsPublished:     true,
			Tags:           []string{"birds eye chili", "chili", "seasoning", "spicy", "fresh vegetables"},
			Slug:           "birds-eye-chili-cabai-rawit-merah",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/birds-eye-chili-cabai-rawit/main",
			MainImagePublicID: "products/birds-eye-chili-cabai-rawit/main",
		},
		// Curly red chillies
		{
			CategoryID:      freshVegetables.ID,
			Name:            "Curly Red Chilies",
			Description:     `With their vibrant color and moderate heat, curly red chilies add both spice and visual appeal to your meals. Ideal for cooking, blending into sauces, or garnishing dishes, they strike the perfect balance between flavor and heat.`,
			IsPublished:     true,
			Tags:           []string{"curly red chilies", "seasoning", "spicy", "fresh vegetables"},
			Slug:           "curly-red-chilies-keriting",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/curly-red-chilies/main",
			MainImagePublicID: "products/curly-red-chilies/main",
		},
		// Iceberg Lettuce
		{
			CategoryID:      freshVegetables.ID,
			Name:            "Iceberg Lettuce",
			Description:     `Crisp, refreshing, and irresistibly light, our iceberg lettuce is the perfect foundation for fresh and vibrant meals. Known for its tightly layered, pale green leaves and signature crunch, this lettuce delivers a clean, mildly sweet flavor that complements a wide variety of dishes without overpowering them.`,
			IsPublished:     true,
			Tags:           []string{"iceberg lettuce", "crispy", "salad", "fresh vegetables"},
			Slug:           "iceberg-lettuce",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/iceberg-lettuce/main",
			MainImagePublicID: "products/iceberg-lettuce/main",
		},
		// Curly Green Lettuce
		{
			CategoryID:      freshVegetables.ID,
			Name:            "Curly Green Lettuce",
			Description:     `With its vibrant color and frilly texture, green curly lettuce adds both crunch and visual appeal to any dish. Fresh and refreshing, it’s perfect for salads, sandwiches, and garnishes.`,
			IsPublished:     true,
			Tags:           []string{"curly lettuce", "green lettuce", "salad", "fresh vegetables"},
			Slug:           "curly-green-lettuce",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/curly-green-lettuce/main",
			MainImagePublicID: "products/curly-green-lettuce/main",
		},
		// Curly Red Lettuce
		{
			CategoryID:      freshVegetables.ID,
			Name:            "Curly Red Lettuce",
			Description:     `A colorful twist on classic lettuce, red curly lettuce offers a mild flavor with a slightly crisp texture. Its striking appearance makes it perfect for elevating presentation while maintaining freshness and taste.`,
			IsPublished:     true,
			Tags:           []string{"curly lettuce", "red lettuce", "salad", "fresh vegetables"},
			Slug:           "curly-red-lettuce",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/curly-red-lettuce/main",
			MainImagePublicID: "products/curly-red-lettuce/main",
		},
		// Romaine Lettuce
		{
			CategoryID:      freshVegetables.ID,
			Name:            "Romaine Lettuce",
			Description:     `Crisp, refreshing, and slightly sweet, romaine lettuce is perfect for salads and wraps. Its sturdy leaves hold dressings well, making it a favorite for classic dishes like Caesar salad. Grown fresh to ensure maximum crunch and flavor.`,
			IsPublished:     true,
			Tags:           []string{"romaine lettuce", "caesar salad", "fresh vegetables"},
			Slug:           "romaine-lettuce",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/romaine-lettuce/main",
			MainImagePublicID: "products/romaine-lettuce/main",
		},
		// Butterhead Lettuce
		{
			CategoryID:      freshVegetables.ID,
			Name:            "Butterhead Lettuce",
			Description:     `Soft, tender, and delicately sweet, butterhead lettuce offers a melt-in-your-mouth texture. Its gentle flavor makes it ideal for fresh salads and light dishes. A premium choice for those who appreciate subtlety and freshness.`,
			IsPublished:     true,
			Tags:           []string{"butterhead lettuce", "salad", "fresh vegetables"},
			Slug:           "butterhead-lettuce",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/butterhead-lettuce/main",
			MainImagePublicID: "products/butterhead-lettuce/main",
		},
		// Caisim
		{
			CategoryID:      freshVegetables.ID,
			Name:            "Caisim (Green Mustard)",
			Description:     `Caisim features tender stems and slightly peppery leaves, making it a flavorful addition to soups and stir-fries. Freshly harvested, it delivers both nutrition and a distinctive taste that enhances traditional dishes.`,
			IsPublished:     true,
			Tags:           []string{"caisim", "fresh vegetables", "salad"},
			Slug:           "green-mustard-caisim",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/green-mustard-caisim/main",
			MainImagePublicID: "products/green-mustard-caisim/main",
		},
		// Salad
		{
			CategoryID:      readyToEat.ID, // Anda mungkin perlu membuat kategori baru "Makanan Siap Saji"
			Name:            "Salad Shake Japanese Style Cup",
			Description:     `A convenient, ready-to-enjoy salad inspired by Japanese flavors. Carefully curated ingredients combine freshness with a light, savory dressing. Just shake, open, and enjoy a balanced, refreshing meal anytime, anywhere.`,
			IsPublished:     true,
			Tags:           []string{"salad", "japanese salad", "healty foods", "ready to eat", "diet food"},
			Slug:           "salad-shake-japanese-style",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/salad-japanese/main",
			MainImagePublicID: "products/salad-japanese/main",
		},
		{
			CategoryID:      readyToEat.ID,
			Name:            "Salad Shake Western Style Cup",
			Description:     `A modern, on-the-go salad packed with fresh vegetables and complemented by a rich, flavorful dressing. Designed for convenience without compromising nutrition, it’s perfect for busy lifestyles seeking healthy options.`,
			IsPublished:     true,
			Tags:           []string{"salad", "western salad", "healty foods", "ready to eat", "caesar salad"},
			Slug:           "salad-shake-western-style",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/salad-western/main",
			MainImagePublicID: "products/salad-western/main",
		},
		// Kimchi
		{
			CategoryID:      homemadeFoods.ID,
			Name:            "Kimchi",
			Description:     `Crafted through traditional fermentation, our kimchi offers a bold, tangy, and slightly spicy flavor. Rich in probiotics and nutrients, it supports digestive health while adding depth to any meal. A timeless staple with a modern twist.`,
			IsPublished:     true,
			Tags:           []string{"kimchi", "korean food", "fermentation", "probiotic", "healty foods"},
			Slug:           "kimchi",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/kimchi/main",
			MainImagePublicID: "products/kimchi/main",
		},
		// Pakcoy Seeds
		{
			CategoryID:      seeds.ID,
			Name:            "Pakcoy Seeds",
			Description:     `High-quality pakcoy seeds selected for strong growth and high yield. Ideal for both beginners and experienced growers, these seeds ensure healthy plants and consistent harvests. Start your own fresh supply at home.`,
			IsPublished:     true,
			Tags:           []string{"seeds", "pakcoy-seeds", "gardening", "hidroponic", "vegetables"},
			Slug:           "pakcoy-seeds",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/pakcoy-seeds/main",
			MainImagePublicID: "products/pakcoy-seeds/main",
		},
		// Caisim Seeds
		{
			CategoryID:      seeds.ID,
			Name:            "Caisim Seeds",
			Description:     `Carefully selected caisim seeds designed for optimal germination and growth. Suitable for home gardening or small-scale farming, they produce fresh, flavorful greens you can enjoy straight from your garden.`,
			IsPublished:     true,
			Tags:           []string{"seeds", "caisim-seeds", "gardening", "hidroponic"},
			Slug:           "caisim-seeds-green-mustard",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/caisim-seeds/main",
			MainImagePublicID: "products/caisim-seeds/main",
		},
		// Shovel
		{
			CategoryID:      gardeningTools.ID,
			Name:            "Gardening Hand Shovel (Wooden Handle)",
			Description:     `Durable and comfortable to use, this hand shovel features a sturdy metal head and a smooth wooden handle for a natural grip. Perfect for planting, digging, and transplanting, it’s an essential tool for every gardener.`,
			IsPublished:     true,
			Tags:           []string{"shovel", "gardening tools", "gardening"},
			Slug:           "gardening-hand-shovel-wooden-handle)",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/shovel/main",
			MainImagePublicID: "products/shovel/main",
		},
		// Watering Can
		{
			CategoryID:      gardeningTools.ID,
			Name:            "Plastic Watering Can",
			Description:     `Lightweight yet durable, this watering can is designed for easy handling and efficient watering. Its ergonomic design ensures smooth pouring, making it ideal for both indoor and outdoor plants.`,
			IsPublished:     true,
			Tags:           []string{"watering can", "gardening tools", "gardening"},
			Slug:           "plastic-watering-can",
			MainImageURL:   "https://res.cloudinary.com/agripacul/image/upload/v1/products/watering-can/main",
			MainImagePublicID: "products/watering-can/main",
		},
	}

	// Seed products and their variants
	for i, product := range products {
		if err := s.db.FirstOrCreate(&product, entity.Product{Slug: product.Slug}).Error; err != nil {
			return err
		}
		products[i].ID = product.ID // Update ID for later use

		// Seed variant types and SKUs for each product
		if err := s.seedVariantTypesAndSKUs(&products[i]); err != nil {
			return err
		}

		// Update MinPrice and MaxPrice after SKUs are created
		if err := s.updateProductPriceRange(&products[i]); err != nil {
			return err
		}
	}

	log.Println("Products seeded successfully")
	return nil
}

func (s *AgripaculProductSeeder) seedVariantTypesAndSKUs(product *entity.Product) error {
	// Create Variant Type "Weight"
	variantWeight := entity.VariantType{
		ProductID: product.ID,
		Name:      "Weight",
	}

	variantColor := entity.VariantType{
		ProductID: product.ID,
		Name:      "Color",
	}

	variantPack := entity.VariantType{
		ProductID: product.ID,
		Name:      "Pack",
	}

	variantMenu := entity.VariantType{
		ProductID: product.ID,
		Name:      "Menu",
	}

	if err := s.db.FirstOrCreate(&variantWeight, entity.VariantType{ProductID: product.ID, Name: "Weight"}).Error; err != nil {
		return err
	}
	if err := s.db.FirstOrCreate(&variantColor, entity.VariantType{ProductID: product.ID, Name: "Color"}).Error; err != nil {
		return err
	}
	if err := s.db.FirstOrCreate(&variantPack, entity.VariantType{ProductID: product.ID, Name: "Pack"}).Error; err != nil {
		return err
	}
	if err := s.db.FirstOrCreate(&variantMenu, entity.VariantType{ProductID: product.ID, Name: "Menu"}).Error; err != nil {
		return err
	}

	// Define variant values and SKUs based on product
	var variantValues []entity.VariantValue
	var skus []entity.SKU

	switch product.Name {
	case "Organic Pakcoy":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantWeight.ID, Value: "250 gram"},
			{VariantTypeID: variantWeight.ID, Value: "500 gram"},
			{VariantTypeID: variantWeight.ID, Value: "1 kg"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-PAK-250",
				Price:     5000,
				SalePrice: float64Ptr(4500),
				Stock:     50,
				MinStock:  10,
				Status:    entity.SKUStatusActive,
				Weight:    250,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-PAK-500",
				Price:     9500,
				SalePrice: float64Ptr(8500),
				Stock:     45,
				MinStock:  8,
				Status:    entity.SKUStatusActive,
				Weight:    500,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-PAK-1000",
				Price:     18000,
				SalePrice: float64Ptr(16500),
				Stock:     30,
				MinStock:  5,
				Status:    entity.SKUStatusActive,
				Weight:    1000,
			},
		}

	case "Hidroponic Pakcoy":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantWeight.ID, Value: "250 gram"},
			{VariantTypeID: variantWeight.ID, Value: "500 gram"},
			{VariantTypeID: variantWeight.ID, Value: "1 kg"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-HDR-PAK-250",
				Price:     5000,
				SalePrice: float64Ptr(4500),
				Stock:     50,
				MinStock:  10,
				Status:    entity.SKUStatusActive,
				Weight:    250,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-HDR-PAK-500",
				Price:     9500,
				SalePrice: float64Ptr(8500),
				Stock:     45,
				MinStock:  8,
				Status:    entity.SKUStatusActive,
				Weight:    500,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-HDR-PAK-1000",
				Price:     18000,
				SalePrice: float64Ptr(16500),
				Stock:     30,
				MinStock:  5,
				Status:    entity.SKUStatusActive,
				Weight:    1000,
			},
		}

	case "Cherry Tomatoes":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantWeight.ID, Value: "250 gram"},
			{VariantTypeID: variantWeight.ID, Value: "500 gram"},
			{VariantTypeID: variantWeight.ID, Value: "1 kg"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-TCH-250",
				Price:     12000,
				SalePrice: float64Ptr(10000),
				Stock:     40,
				MinStock:  8,
				Status:    entity.SKUStatusActive,
				Weight:    250,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-TCH-500",
				Price:     22000,
				SalePrice: float64Ptr(20000),
				Stock:     35,
				MinStock:  7,
				Status:    entity.SKUStatusActive,
				Weight:    500,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-TCH-1000",
				Price:     40000,
				SalePrice: float64Ptr(38000),
				Stock:     25,
				MinStock:  5,
				Status:    entity.SKUStatusActive,
				Weight:    1000,
			},
		}

	case "Local Red Tomatoes":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantWeight.ID, Value: "500 gram"},
			{VariantTypeID: variantWeight.ID, Value: "1 kg"},
			{VariantTypeID: variantWeight.ID, Value: "2 kg"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-TMR-500",
				Price:     8000,
				SalePrice: float64Ptr(7000),
				Stock:     60,
				MinStock:  12,
				Status:    entity.SKUStatusActive,
				Weight:    500,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-TMR-1000",
				Price:     15000,
				SalePrice: float64Ptr(13500),
				Stock:     55,
				MinStock:  10,
				Status:    entity.SKUStatusActive,
				Weight:    1000,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-TMR-2000",
				Price:     28000,
				SalePrice: float64Ptr(26000),
				Stock:     40,
				MinStock:  8,
				Status:    entity.SKUStatusActive,
				Weight:    2000,
			},
		}

	case "Sweet Corn":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantWeight.ID, Value: "2 pcs"},
			{VariantTypeID: variantWeight.ID, Value: "4 pcs"},
			{VariantTypeID: variantWeight.ID, Value: "6 pcs"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-JGM-2",
				Price:     12000,
				SalePrice: float64Ptr(10000),
				Stock:     50,
				MinStock:  10,
				Status:    entity.SKUStatusActive,
				Weight:    500, // approx weight for 2 corns
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-JGM-4",
				Price:     22000,
				SalePrice: float64Ptr(20000),
				Stock:     45,
				MinStock:  9,
				Status:    entity.SKUStatusActive,
				Weight:    1000, // approx weight for 4 corns
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-JGM-6",
				Price:     32000,
				SalePrice: float64Ptr(30000),
				Stock:     35,
				MinStock:  7,
				Status:    entity.SKUStatusActive,
				Weight:    1500, // approx weight for 6 corns
			},
		}

	case "Green Beans (Buncis)":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantWeight.ID, Value: "250 gram"},
			{VariantTypeID: variantWeight.ID, Value: "500 gram"},
			{VariantTypeID: variantWeight.ID, Value: "1 kg"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-BNC-250",
				Price:     5000,
				SalePrice: float64Ptr(4500),
				Stock:     60,
				MinStock:  12,
				Status:    entity.SKUStatusActive,
				Weight:    250,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-BNC-500",
				Price:     9500,
				SalePrice: float64Ptr(8500),
				Stock:     55,
				MinStock:  10,
				Status:    entity.SKUStatusActive,
				Weight:    500,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-BNC-1000",
				Price:     18000,
				SalePrice: float64Ptr(16000),
				Stock:     40,
				MinStock:  8,
				Status:    entity.SKUStatusActive,
				Weight:    1000,
			},
		}

	case "Fresh Cabbage":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantWeight.ID, Value: "500 gram"},
			{VariantTypeID: variantWeight.ID, Value: "1 kg"},
			{VariantTypeID: variantWeight.ID, Value: "1 pcs (±1.5kg)"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-KOL-500",
				Price:     6000,
				SalePrice: float64Ptr(5000),
				Stock:     45,
				MinStock:  9,
				Status:    entity.SKUStatusActive,
				Weight:    500,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-KOL-1000",
				Price:     11000,
				SalePrice: float64Ptr(10000),
				Stock:     40,
				MinStock:  8,
				Status:    entity.SKUStatusActive,
				Weight:    1000,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-KOL-UTUH",
				Price:     15000,
				SalePrice: float64Ptr(14000),
				Stock:     30,
				MinStock:  6,
				Status:    entity.SKUStatusActive,
				Weight:    1500,
			},
		}

	case "Fresh Green Broccoli":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantWeight.ID, Value: "250 gram"},
			{VariantTypeID: variantWeight.ID, Value: "500 gram"},
			{VariantTypeID: variantWeight.ID, Value: "1 kg"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-BRK-250",
				Price:     10000,
				SalePrice: float64Ptr(9000),
				Stock:     35,
				MinStock:  7,
				Status:    entity.SKUStatusActive,
				Weight:    250,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-BRK-500",
				Price:     18000,
				SalePrice: float64Ptr(16000),
				Stock:     30,
				MinStock:  6,
				Status:    entity.SKUStatusActive,
				Weight:    500,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-BRK-1000",
				Price:     32000,
				SalePrice: float64Ptr(30000),
				Stock:     20,
				MinStock:  4,
				Status:    entity.SKUStatusActive,
				Weight:    1000,
			},
		}

	case "Long Beans":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantWeight.ID, Value: "250 gram"},
			{VariantTypeID: variantWeight.ID, Value: "500 gram"},
			{VariantTypeID: variantWeight.ID, Value: "1 bunch"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-KCP-250",
				Price:     4000,
				SalePrice: float64Ptr(3500),
				Stock:     65,
				MinStock:  13,
				Status:    entity.SKUStatusActive,
				Weight:    250,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-KCP-500",
				Price:     7500,
				SalePrice: float64Ptr(6500),
				Stock:     60,
				MinStock:  12,
				Status:    entity.SKUStatusActive,
				Weight:    500,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-KCP-IKAT",
				Price:     5000,
				SalePrice: float64Ptr(4500),
				Stock:     70,
				MinStock:  14,
				Status:    entity.SKUStatusActive,
				Weight:    400, // 1 ikat ±400 gram
			},
		}

	case "Bird's Eye Chili (Cabai Rawit)":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantWeight.ID, Value: "100 gram"},
			{VariantTypeID: variantWeight.ID, Value: "250 gram"},
			{VariantTypeID: variantWeight.ID, Value: "500 gram"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-CRW-100",
				Price:     15000,
				SalePrice: float64Ptr(13000),
				Stock:     40,
				MinStock:  8,
				Status:    entity.SKUStatusActive,
				Weight:    100,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-CRW-250",
				Price:     35000,
				SalePrice: float64Ptr(32000),
				Stock:     35,
				MinStock:  7,
				Status:    entity.SKUStatusActive,
				Weight:    250,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-CRW-500",
				Price:     65000,
				SalePrice: float64Ptr(60000),
				Stock:     25,
				MinStock:  5,
				Status:    entity.SKUStatusActive,
				Weight:    500,
			},
		}

	case "Curly Red Chilies":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantWeight.ID, Value: "250 gram"},
			{VariantTypeID: variantWeight.ID, Value: "500 gram"},
			{VariantTypeID: variantWeight.ID, Value: "1 kg"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-CMR-250",
				Price:     20000,
				SalePrice: float64Ptr(18000),
				Stock:     45,
				MinStock:  9,
				Status:    entity.SKUStatusActive,
				Weight:    250,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-CMR-500",
				Price:     38000,
				SalePrice: float64Ptr(35000),
				Stock:     40,
				MinStock:  8,
				Status:    entity.SKUStatusActive,
				Weight:    500,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-CMR-1000",
				Price:     72000,
				SalePrice: float64Ptr(68000),
				Stock:     30,
				MinStock:  6,
				Status:    entity.SKUStatusActive,
				Weight:    1000,
			},
		}

	case "Iceberg Lettuce":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantWeight.ID, Value: "250 gram"},
			{VariantTypeID: variantWeight.ID, Value: "500 gram"},
			{VariantTypeID: variantWeight.ID, Value: "1 pcs (±500gr)"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-ICE-250",
				Price:     10000,
				SalePrice: float64Ptr(9000),
				Stock:     40,
				MinStock:  8,
				Status:    entity.SKUStatusActive,
				Weight:    250,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-ICE-500",
				Price:     18000,
				SalePrice: float64Ptr(16000),
				Stock:     35,
				MinStock:  7,
				Status:    entity.SKUStatusActive,
				Weight:    500,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-ICE-UTUH",
				Price:     15000,
				SalePrice: float64Ptr(14000),
				Stock:     45,
				MinStock:  9,
				Status:    entity.SKUStatusActive,
				Weight:    500,
			},
		}

	case "Curly Green Lettuce":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantWeight.ID, Value: "100 gram"},
			{VariantTypeID: variantWeight.ID, Value: "250 gram"},
			{VariantTypeID: variantWeight.ID, Value: "500 gram"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-SKH-100",
				Price:     8000,
				SalePrice: float64Ptr(7000),
				Stock:     50,
				MinStock:  10,
				Status:    entity.SKUStatusActive,
				Weight:    100,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-SKH-250",
				Price:     18000,
				SalePrice: float64Ptr(16000),
				Stock:     45,
				MinStock:  9,
				Status:    entity.SKUStatusActive,
				Weight:    250,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-SKH-500",
				Price:     32000,
				SalePrice: float64Ptr(30000),
				Stock:     35,
				MinStock:  7,
				Status:    entity.SKUStatusActive,
				Weight:    500,
			},
		}

	case "Curly Red Lettuce":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantWeight.ID, Value: "100 gram"},
			{VariantTypeID: variantWeight.ID, Value: "250 gram"},
			{VariantTypeID: variantWeight.ID, Value: "500 gram"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-SKM-100",
				Price:     9000,
				SalePrice: float64Ptr(8000),
				Stock:     45,
				MinStock:  9,
				Status:    entity.SKUStatusActive,
				Weight:    100,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-SKM-250",
				Price:     20000,
				SalePrice: float64Ptr(18000),
				Stock:     40,
				MinStock:  8,
				Status:    entity.SKUStatusActive,
				Weight:    250,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-SKM-500",
				Price:     35000,
				SalePrice: float64Ptr(33000),
				Stock:     30,
				MinStock:  6,
				Status:    entity.SKUStatusActive,
				Weight:    500,
			},
		}

	case "Romaine Lettuce":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantWeight.ID, Value: "250 gram"},
			{VariantTypeID: variantWeight.ID, Value: "500 gram"},
			{VariantTypeID: variantWeight.ID, Value: "1 bunch"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-ROM-250",
				Price:     12000,
				SalePrice: float64Ptr(10000),
				Stock:     45,
				MinStock:  9,
				Status:    entity.SKUStatusActive,
				Weight:    250,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-ROM-500",
				Price:     22000,
				SalePrice: float64Ptr(20000),
				Stock:     40,
				MinStock:  8,
				Status:    entity.SKUStatusActive,
				Weight:    500,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-ROM-IKAT",
				Price:     18000,
				SalePrice: float64Ptr(16000),
				Stock:     50,
				MinStock:  10,
				Status:    entity.SKUStatusActive,
				Weight:    400,
			},
		}

	case "Butterhead Lettuce":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantWeight.ID, Value: "250 gram"},
			{VariantTypeID: variantWeight.ID, Value: "500 gram"},
			{VariantTypeID: variantWeight.ID, Value: "1 pcs"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-BUT-250",
				Price:     13000,
				SalePrice: float64Ptr(12000),
				Stock:     40,
				MinStock:  8,
				Status:    entity.SKUStatusActive,
				Weight:    250,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-BUT-500",
				Price:     24000,
				SalePrice: float64Ptr(22000),
				Stock:     35,
				MinStock:  7,
				Status:    entity.SKUStatusActive,
				Weight:    500,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-BUT-UTUH",
				Price:     20000,
				SalePrice: float64Ptr(18000),
				Stock:     45,
				MinStock:  9,
				Status:    entity.SKUStatusActive,
				Weight:    450,
			},
		}

	case "Caisim (Green Mustard)":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantWeight.ID, Value: "250 gram"},
			{VariantTypeID: variantWeight.ID, Value: "500 gram"},
			{VariantTypeID: variantWeight.ID, Value: "1 kg"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-CAI-250",
				Price:     5000,
				SalePrice: float64Ptr(4500),
				Stock:     60,
				MinStock:  12,
				Status:    entity.SKUStatusActive,
				Weight:    250,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-CAI-500",
				Price:     9500,
				SalePrice: float64Ptr(8500),
				Stock:     55,
				MinStock:  11,
				Status:    entity.SKUStatusActive,
				Weight:    500,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-CAI-1000",
				Price:     17000,
				SalePrice: float64Ptr(16000),
				Stock:     45,
				MinStock:  9,
				Status:    entity.SKUStatusActive,
				Weight:    1000,
			},
		}

	case "Salad Shake Japanese Style Cup":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantMenu.ID, Value: "Regular Cup"},
			{VariantTypeID: variantMenu.ID, Value: "With Topping (Edamame + Nori)"},
			{VariantTypeID: variantMenu.ID, Value: "With Protein (Chicken Teriyaki)"},
			{VariantTypeID: variantMenu.ID, Value: "Family Size"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-JPN-REG",
				Price:     25000,
				SalePrice: float64Ptr(22000),
				Stock:     30,
				MinStock:  6,
				Status:    entity.SKUStatusActive,
				Weight:    200, // gram
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-JPN-TOP",
				Price:     32000,
				SalePrice: float64Ptr(28000),
				Stock:     25,
				MinStock:  5,
				Status:    entity.SKUStatusActive,
				Weight:    250, // gram
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-JPN-PRO",
				Price:     38000,
				SalePrice: float64Ptr(35000),
				Stock:     20,
				MinStock:  4,
				Status:    entity.SKUStatusActive,
				Weight:    300, // gram
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-JPN-FAM",
				Price:     65000,
				SalePrice: float64Ptr(60000),
				Stock:     15,
				MinStock:  3,
				Status:    entity.SKUStatusActive,
				Weight:    600, // gram
			},
		}

	case "Salad Shake Western Style Cup":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantMenu.ID, Value: "Regular Cup"},
			{VariantTypeID: variantMenu.ID, Value: "With Topping (Crouton + Parmesan)"},
			{VariantTypeID: variantMenu.ID, Value: "With Protein (Grilled Chicken)"},
			{VariantTypeID: variantMenu.ID, Value: "Family Size"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-WST-REG",
				Price:     25000,
				SalePrice: float64Ptr(22000),
				Stock:     30,
				MinStock:  6,
				Status:    entity.SKUStatusActive,
				Weight:    200, // gram
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-WST-TOP",
				Price:     30000,
				SalePrice: float64Ptr(27000),
				Stock:     25,
				MinStock:  5,
				Status:    entity.SKUStatusActive,
				Weight:    230, // gram
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-WST-PRO",
				Price:     38000,
				SalePrice: float64Ptr(35000),
				Stock:     20,
				MinStock:  4,
				Status:    entity.SKUStatusActive,
				Weight:    300, // gram
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-WST-FAM",
				Price:     65000,
				SalePrice: float64Ptr(60000),
				Stock:     15,
				MinStock:  3,
				Status:    entity.SKUStatusActive,
				Weight:    600, // gram
			},
		}

	case "Kimchi":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantWeight.ID, Value: "100 gram"},
			{VariantTypeID: variantWeight.ID, Value: "250 gram"},
			{VariantTypeID: variantWeight.ID, Value: "500 gram"},
			{VariantTypeID: variantWeight.ID, Value: "1 kg"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-KIM-100",
				Price:     15000,
				SalePrice: float64Ptr(13000),
				Stock:     40,
				MinStock:  8,
				Status:    entity.SKUStatusActive,
				Weight:    100,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-KIM-250",
				Price:     32000,
				SalePrice: float64Ptr(28000),
				Stock:     35,
				MinStock:  7,
				Status:    entity.SKUStatusActive,
				Weight:    250,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-KIM-500",
				Price:     58000,
				SalePrice: float64Ptr(55000),
				Stock:     30,
				MinStock:  6,
				Status:    entity.SKUStatusActive,
				Weight:    500,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-KIM-1000",
				Price:     105000,
				SalePrice: float64Ptr(100000),
				Stock:     20,
				MinStock:  4,
				Status:    entity.SKUStatusActive,
				Weight:    1000,
			},
		}

	case "Pakcoy Seeds":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantPack.ID, Value: "East West - Repack 10gr"},
			{VariantTypeID: variantPack.ID, Value: "East West - 1 Pack Original"},
			{VariantTypeID: variantPack.ID, Value: "Bintang Asia - Repack 10gr"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-BPK-EWR10",
				Price:     12000,
				SalePrice: float64Ptr(10000),
				Stock:     50,
				MinStock:  10,
				Status:    entity.SKUStatusActive,
				Weight:    10, // gram
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-BPK-EWO",
				Price:     25000,
				SalePrice: float64Ptr(22000),
				Stock:     30,
				MinStock:  6,
				Status:    entity.SKUStatusActive,
				Weight:    50, // gram (estimasi berat bungkus original)
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-BPK-KTR10",
				Price:     10000,
				SalePrice: float64Ptr(8500),
				Stock:     50,
				MinStock:  10,
				Status:    entity.SKUStatusActive,
				Weight:    10,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-BPK-KTO",
				Price:     20000,
				SalePrice: float64Ptr(18000),
				Stock:     30,
				MinStock:  6,
				Status:    entity.SKUStatusActive,
				Weight:    40,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-BPK-BAR10",
				Price:     9000,
				SalePrice: float64Ptr(8000),
				Stock:     50,
				MinStock:  10,
				Status:    entity.SKUStatusActive,
				Weight:    10,
			},
		}

	case "Caisim Seeds":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantPack.ID, Value: "East West - Repack 10gr"},
			{VariantTypeID: variantPack.ID, Value: "East West - 1 Pack Original"},
			{VariantTypeID: variantPack.ID, Value: "Bintang Asia - Repack 10gr"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-BCA-EWR10",
				Price:     12000,
				SalePrice: float64Ptr(10000),
				Stock:     50,
				MinStock:  10,
				Status:    entity.SKUStatusActive,
				Weight:    10,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-BCA-EWO",
				Price:     25000,
				SalePrice: float64Ptr(22000),
				Stock:     30,
				MinStock:  6,
				Status:    entity.SKUStatusActive,
				Weight:    50,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-BCA-KTR10",
				Price:     10000,
				SalePrice: float64Ptr(8500),
				Stock:     50,
				MinStock:  10,
				Status:    entity.SKUStatusActive,
				Weight:    10,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-BCA-KTO",
				Price:     20000,
				SalePrice: float64Ptr(18000),
				Stock:     30,
				MinStock:  6,
				Status:    entity.SKUStatusActive,
				Weight:    40,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-BCA-BAR10",
				Price:     9000,
				SalePrice: float64Ptr(8000),
				Stock:     50,
				MinStock:  10,
				Status:    entity.SKUStatusActive,
				Weight:    10,
			},
		}

	case "Gardening Hand Shovel (Wooden Handle)":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantColor.ID, Value: "Black"},
			{VariantTypeID: variantColor.ID, Value: "Brown"},
			{VariantTypeID: variantColor.ID, Value: "Green Army"},
			{VariantTypeID: variantColor.ID, Value: "Blue"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-SKP-HTM",
				Price:     35000,
				SalePrice: float64Ptr(32000),
				Stock:     25,
				MinStock:  5,
				Status:    entity.SKUStatusActive,
				Weight:    300, // gram
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-SKP-CKL",
				Price:     35000,
				SalePrice: float64Ptr(32000),
				Stock:     25,
				MinStock:  5,
				Status:    entity.SKUStatusActive,
				Weight:    300,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-SKP-HJU",
				Price:     38000,
				SalePrice: float64Ptr(35000),
				Stock:     20,
				MinStock:  4,
				Status:    entity.SKUStatusActive,
				Weight:    300,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-SKP-BRU",
				Price:     35000,
				SalePrice: float64Ptr(32000),
				Stock:     25,
				MinStock:  5,
				Status:    entity.SKUStatusActive,
				Weight:    300,
			},
		}

	case "Plastic Watering Can":
		variantValues = []entity.VariantValue{
			{VariantTypeID: variantColor.ID, Value: "1 Liter - Green"},
			{VariantTypeID: variantColor.ID, Value: "1 Liter - Blue"},
			{VariantTypeID: variantColor.ID, Value: "1 Liter - Red"},
			{VariantTypeID: variantColor.ID, Value: "2 Liter - Green"},
			{VariantTypeID: variantColor.ID, Value: "2 Liter - Blue"},
			{VariantTypeID: variantColor.ID, Value: "5 Liter - Green"},
		}
		skus = []entity.SKU{
			{
				ProductID: product.ID,
				SKUCode:   "AGR-GMB-1HJU",
				Price:     25000,
				SalePrice: float64Ptr(22000),
				Stock:     30,
				MinStock:  6,
				Status:    entity.SKUStatusActive,
				Weight:    200, // gram (berat kosong)
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-GMB-1BRU",
				Price:     25000,
				SalePrice: float64Ptr(22000),
				Stock:     30,
				MinStock:  6,
				Status:    entity.SKUStatusActive,
				Weight:    200,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-GMB-1MRH",
				Price:     25000,
				SalePrice: float64Ptr(22000),
				Stock:     30,
				MinStock:  6,
				Status:    entity.SKUStatusActive,
				Weight:    200,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-GMB-2HJU",
				Price:     35000,
				SalePrice: float64Ptr(32000),
				Stock:     25,
				MinStock:  5,
				Status:    entity.SKUStatusActive,
				Weight:    300,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-GMB-2BRU",
				Price:     35000,
				SalePrice: float64Ptr(32000),
				Stock:     25,
				MinStock:  5,
				Status:    entity.SKUStatusActive,
				Weight:    300,
			},
			{
				ProductID: product.ID,
				SKUCode:   "AGR-GMB-5HJU",
				Price:     55000,
				SalePrice: float64Ptr(50000),
				Stock:     15,
				MinStock:  3,
				Status:    entity.SKUStatusActive,
				Weight:    500,
			},
		}
	}

	// Create Variant Values and SKUs
	for i, vv := range variantValues {
		// Create Variant Value
		if err := s.db.FirstOrCreate(&variantValues[i], entity.VariantValue{VariantTypeID: vv.VariantTypeID, Value: vv.Value}).Error; err != nil {
			return err
		}

		// Update SKU with Variant Value ID through SKUVariantValue
		sku := skus[i]
		if err := s.db.FirstOrCreate(&sku, entity.SKU{SKUCode: sku.SKUCode}).Error; err != nil {
			return err
		}

		// Create SKU Variant Value relationship
		skuVariantValue := entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: variantValues[i].ID,
		}
		if err := s.db.FirstOrCreate(&skuVariantValue, entity.SKUVariantValue{SKUID: sku.ID, VariantValueID: variantValues[i].ID}).Error; err != nil {
			return err
		}

		// Update SoldCount randomly
		product.SoldCount = rand.Intn(200) + 50
		s.db.Save(&sku)
	}

	return nil
}

func (s *AgripaculProductSeeder) updateProductPriceRange(product *entity.Product) error {
	var skus []entity.SKU
	if err := s.db.Where("product_id = ?", product.ID).Find(&skus).Error; err != nil {
		return err
	}

	if len(skus) > 0 {
		minPrice := skus[0].Price
		maxPrice := skus[0].Price
		totalSold := 0

		for _, sku := range skus {
			if sku.Price < minPrice {
				minPrice = sku.Price
			}
			if sku.Price > maxPrice {
				maxPrice = sku.Price
			}
			totalSold += product.SoldCount
		}

		// Update product with aggregated data
		if err := s.db.Model(&entity.Product{}).Where("id = ?", product.ID).Updates(map[string]interface{}{
			"min_price":     minPrice,
			"max_price":     maxPrice,
			"sold_count":    totalSold,
			"review_count":  rand.Intn(50) + 10,
			"average_rating": float64(rand.Intn(20)+40) / 10, // Random rating between 4.0 - 5.0
		}).Error; err != nil {
			return err
		}
	}

	return nil
}

// Helper function
func float64Ptr(f float64) *float64 {
	return &f
}