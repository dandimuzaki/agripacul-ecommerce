package seeder

import (
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/pkg/utils"
	"fmt"
	"math/rand"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

// SeedStrawCollection seeds the complete stainless steel straw product line
func SeedStrawCollection(db *gorm.DB) error {
	rand.Seed(time.Now().UnixNano())
	
	// Create specialized categories for straw products
	category, err := createStrawCutleryCategory(db)
	if err != nil {
		return err
	}
	
	// Seed product lines (7 products instead of 1 mega-product)
	products := []struct {
		name     string
		category uint
		fn       func(*gorm.DB, uint) error
	}{
		{"Individual Straws", category.ID, seedIndividualStraws},
		{"Couple Straw Sets", category.ID, seedCoupleSets},
		{"Family Straw Sets", category.ID, seedFamilySets},
		{"Premium Colored Straws", category.ID, seedPremiumColoredStraws},
		{"Bulk Packs for Business", category.ID, seedBulkPacks},
		{"Travel Straw Sets", category.ID, seedTravelSets},
		{"Starter Kits", category.ID, seedStarterKits},
	}
	
	for _, p := range products {
		if err := p.fn(db, p.category); err != nil {
			return fmt.Errorf("failed to seed %s: %v", p.name, err)
		}
	}
	
	fmt.Println("✅ Successfully seeded Stainless Steel Straw Collection (7 product lines)")
	return nil
}

// ==================== PRODUCT LINE 1: INDIVIDUAL STRAWS ====================
// Target: Solo drinkers, students, minimalists
// Price based on Tokopedia: Rp 25,000 - 45,000

func seedIndividualStraws(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Individual Stainless Steel Straw - Single with Cleaning Brush",
		Description: `A single premium stainless steel straw with its own cleaning brush. Perfect for personal use, travel, or trying reusable straws for the first time.

**✨ What's Included:**
- 1 Stainless steel straw (your choice of style)
- 1 Cleaning brush
- Optional silicone tip

**🌱 Why Choose Reusable Straws?**
Each reusable straw replaces hundreds of single-use plastic straws annually. Indonesia is one of the largest contributors to ocean plastic, and switching to reusable straws is a simple but impactful change.

**📏 Size Options:**
- **Regular (6mm)**: For water, juice, soft drinks
- **Wide (8mm)**: For smoothies, boba tea
- **Jumbo (10mm)**: For thick smoothies, milkshakes

**🧹 Cleaning Brush Included:**
The included brush has a stainless steel handle and durable nylon bristles to keep your straw hygienic.`,
		IsPublished:      true,
		Tags:             pq.StringArray{"individual", "single straw", "personal", "student", "travel", "sedotan satuan"},
		Slug:             "individual-stainless-steel-straw-cleaning-brush",
		MainImageURL:     "https://down-id.img.susercontent.com/file/ac1cd9b9c0750b191865430c1cf4cd54@resize_w900_nl.webp",
		MainImagePublicID: "individual-straw-main",
		AverageRating:    4.7,
		ReviewCount:      345,
		SoldCount:        1234,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	// Create variant type (style + diameter combination)
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Straw Type",
	}
	db.Create(&variantType)
	
	// Individual straw variants
	variants := []struct {
		Name        string
		Style       string // straight or bent
		Diameter    string // regular, wide, jumbo
		Price       float64
		Weight      float64
		Stock       int
		MinStock    int
	}{
		{
			Name:     "Straight - Regular (6mm) for Standard Drinks",
			Style:    "straight",
			Diameter: "regular",
			Price:    25000,
			Weight:   25,
			Stock:    500,
			MinStock: 50,
		},
		{
			Name:     "Straight - Wide (8mm) for Boba/Smoothies",
			Style:    "straight",
			Diameter: "wide",
			Price:    28000,
			Weight:   28,
			Stock:    400,
			MinStock: 40,
		},
		{
			Name:     "Straight - Jumbo (10mm) for Thick Drinks",
			Style:    "straight",
			Diameter: "jumbo",
			Price:    32000,
			Weight:   32,
			Stock:    300,
			MinStock: 30,
		},
		{
			Name:     "Bent - Regular (6mm) for Standard Drinks",
			Style:    "bent",
			Diameter: "regular",
			Price:    27000,
			Weight:   26,
			Stock:    400,
			MinStock: 40,
		},
		{
			Name:     "Bent - Wide (8mm) for Boba/Smoothies",
			Style:    "bent",
			Diameter: "wide",
			Price:    30000,
			Weight:   29,
			Stock:    350,
			MinStock: 35,
		},
		{
			Name:     "Bent - Jumbo (10mm) for Thick Drinks",
			Style:    "bent",
			Diameter: "jumbo",
			Price:    35000,
			Weight:   33,
			Stock:    250,
			MinStock: 25,
		},
	}
	
	for _, v := range variants {
		// Create variant value
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Name,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		// Generate SKU code
		skuCode := fmt.Sprintf("STR-IND-%s-%s-%03d-%s", v.Style[:1], v.Diameter[:1], rand.Intn(1000), str)
		
		// 25% chance of sale price
		var salePrice *float64
		if rand.Float64() < 0.25 {
			sp := v.Price * 0.9 // 10% off
			sp = float64(int(sp/1000)) * 1000
			salePrice = &sp
		}
		
		// Create SKU
		sku := entity.SKU{
			ProductID: product.ID,
			SKUCode:   skuCode,
			Price:     v.Price,
			SalePrice: salePrice,
			Stock:     v.Stock,
			MinStock:  v.MinStock,
			Status:    entity.SKUStatusActive,
			Weight:    v.Weight,
		}
		db.Create(&sku)
		
		// Link SKU to variant
		db.Create(&entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: value.ID,
		})
	}
	
	// Update product price range
	product.MinPrice = 25000
	product.MaxPrice = 35000
	db.Save(&product)
	
	if err := addStainlessStrawImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Individual Straws: %d options (Rp 25rb - 35rb)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 2: COUPLE SETS ====================
// Target: Couples, small households, sharing

func seedCoupleSets(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Couple Stainless Steel Straw Set - 2 Straws + 1 Brush",
		Description: `Perfect for couples or small households. This set includes two straws and one shared cleaning brush, with matching or complementary styles.

**✨ What's Included:**
- 2 Stainless steel straws (choose your combination)
- 1 Cleaning brush
- 1 Cotton storage pouch
- 2 Silicone tips

**💑 Perfect For:**
- Couples who share drinks
- Small apartments
- Roommates
- Guest supplies

**📦 Set Options:**
- Both straight
- Both bent
- One of each (mixed)

Available in regular, wide, or jumbo diameters.`,
		IsPublished:      true,
		Tags:             pq.StringArray{"couple", "pasangan", "dua", "set for two", "sharing", "household"},
		Slug:             "couple-stainless-steel-straw-set-2-straws",
		MainImageURL:     "https://down-id.img.susercontent.com/file/d6bbcb2b9502fd90254fe7f35d788634.webp",
		MainImagePublicID: "couple-straw-main",
		AverageRating:    4.8,
		ReviewCount:      234,
		SoldCount:        890,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Set Type",
	}
	db.Create(&variantType)
	
	variants := []struct {
		Name     string
		Style    string // both straight, both bent, mixed
		Diameter string
		Price    float64
		Weight   float64
		Stock    int
		MinStock int
	}{
		{
			Name:     "Both Straight - Regular (6mm)",
			Style:    "both straight",
			Diameter: "regular",
			Price:    45000,
			Weight:   55,
			Stock:    300,
			MinStock: 40,
		},
		{
			Name:     "Both Straight - Wide (8mm)",
			Style:    "both straight",
			Diameter: "wide",
			Price:    49000,
			Weight:   60,
			Stock:    250,
			MinStock: 35,
		},
		{
			Name:     "Both Bent - Regular (6mm)",
			Style:    "both bent",
			Diameter: "regular",
			Price:    49000,
			Weight:   56,
			Stock:    250,
			MinStock: 35,
		},
		{
			Name:     "Mixed (1 Straight + 1 Bent) - Regular (6mm)",
			Style:    "mixed",
			Diameter: "regular",
			Price:    52000,
			Weight:   58,
			Stock:    300,
			MinStock: 40,
		},
		{
			Name:     "Mixed (1 Straight + 1 Bent) - Wide (8mm)",
			Style:    "mixed",
			Diameter: "wide",
			Price:    56000,
			Weight:   63,
			Stock:    250,
			MinStock: 35,
		},
		{
			Name:     "Mixed Diameters (6mm + 8mm)",
			Style:    "mixed",
			Diameter: "mixed",
			Price:    65000,
			Weight:   65,
			Stock:    200,
			MinStock: 30,
		},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Name,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("STR-CPL-%03d-%s", rand.Intn(1000), str)
		
		sku := entity.SKU{
			ProductID: product.ID,
			SKUCode:   skuCode,
			Price:     v.Price,
			Stock:     v.Stock,
			MinStock:  v.MinStock,
			Status:    entity.SKUStatusActive,
			Weight:    v.Weight,
		}
		db.Create(&sku)
		
		db.Create(&entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: value.ID,
		})
	}
	
	product.MinPrice = 45000
	product.MaxPrice = 65000
	db.Save(&product)
	
	if err := addStainlessStrawImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Couple Sets: %d options (Rp 45rb - 65rb)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 3: FAMILY SETS ====================
// Target: Families with children, regular household use

func seedFamilySets(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Family Stainless Steel Straw Set - 4 Straws + 2 Brushes",
		Description: `The complete set for a family of four. Includes multiple straw styles and two brushes for convenient cleaning. Perfect for daily household use.

**✨ What's Included:**
- 4 Stainless steel straws (mix of styles)
- 2 Cleaning brushes (keep one upstairs, one downstairs)
- 1 Cotton storage pouch
- 4 Silicone tips

**👨‍👩‍👧‍👦 Perfect For:**
- Families with children
- Small restaurants
- Frequent entertainers
- Daily household use

**📋 Set Includes:**
- 2 Straight straws (regular)
- 2 Bent straws (regular)
- Mixed diameters available

**🌱 Environmental Impact:**
One family set replaces over 1,000 plastic straws per year!`,
		IsPublished:      true,
		Tags:             pq.StringArray{"family", "keluarga", "4 straws", "household", "daily use", "complete set"},
		Slug:             "family-stainless-steel-straw-set-4-straws",
		MainImageURL:     "https://down-id.img.susercontent.com/file/29c865f330b5467a2f4cc5202f5f8331.webp",
		MainImagePublicID: "family-straw-main",
		AverageRating:    4.9,
		ReviewCount:      456,
		SoldCount:        2345,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Diameter Type",
	}
	db.Create(&variantType)
	
	variants := []struct {
		Name     string
		Price    float64
		Weight   float64
		Stock    int
		MinStock int
	}{
		{
			Name:     "All Regular (6mm) - Standard Drinks",
			Price:    75000,
			Weight:   110,
			Stock:    400,
			MinStock: 50,
		},
		{
			Name:     "All Wide (8mm) - Boba/Smoothies",
			Price:    85000,
			Weight:   120,
			Stock:    300,
			MinStock: 40,
		},
		{
			Name:     "Mixed (2 Regular + 2 Wide)",
			Price:    95000,
			Weight:   115,
			Stock:    350,
			MinStock: 45,
		},
		{
			Name:     "Mixed Diameters (6mm + 8mm + 10mm)",
			Price:    115000,
			Weight:   130,
			Stock:    250,
			MinStock: 35,
		},
		{
			Name:     "All Jumbo (10mm) - Thick Smoothies",
			Price:    105000,
			Weight:   140,
			Stock:    200,
			MinStock: 30,
		},
		{
			Name:     "Premium Mixed with Colored Tips",
			Price:    125000,
			Weight:   125,
			Stock:    150,
			MinStock: 25,
		},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Name,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("STR-FAM-%03d-%s", rand.Intn(1000), str)
		
		sku := entity.SKU{
			ProductID: product.ID,
			SKUCode:   skuCode,
			Price:     v.Price,
			Stock:     v.Stock,
			MinStock:  v.MinStock,
			Status:    entity.SKUStatusActive,
			Weight:    v.Weight,
		}
		db.Create(&sku)
		
		db.Create(&entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: value.ID,
		})
	}
	
	product.MinPrice = 75000
	product.MaxPrice = 125000
	db.Save(&product)
	
	if err := addStainlessStrawImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Family Sets: %d options (Rp 75rb - 125rb)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 4: PREMIUM COLORED STRAWS ====================
// Target: Fashion-conscious buyers, gift shoppers

func seedPremiumColoredStraws(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Premium Colored Stainless Steel Straws - Gold, Rose Gold, Rainbow",
		Description: `Elevate your drinking experience with our premium colored straws. Available in stunning finishes that add style to any beverage while being completely eco-friendly.

**✨ Available Finishes:**
- **Gold**: Elegant and timeless
- **Rose Gold**: Trendy and romantic
- **Black**: Sleek and modern
- **Rainbow**: Fun and vibrant
- **Mirror Polish**: Classic shine

**🎁 Perfect For:**
- Gift giving
- Special occasions
- Instagram-worthy drinks
- Weddings and parties

**📦 Set Options:**
- Individual straws
- Couple sets
- Family sets
- All with matching colored brushes and pouches

**✨ Features:**
- High-quality color coating (food-safe)
- Won't fade or peel with proper care
- Includes matching cleaning brush
- Comes in coordinating pouch

The same durability as our standard straws, with added style!`,
		IsPublished:      true,
		Tags:             pq.StringArray{"gold", "rose gold", "rainbow", "colored", "premium", "fashion", "gift", "hadiah"},
		Slug:             "premium-colored-stainless-steel-straws-gold-rose-gold-rainbow",
		MainImageURL:     "https://down-id.img.susercontent.com/file/id-11134207-7r98v-lrwwhgxi0eo9a0.webp",
		MainImagePublicID: "colored-straws-main",
		AverageRating:    4.8,
		ReviewCount:      189,
		SoldCount:        678,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Color & Set",
	}
	db.Create(&variantType)
	
	variants := []struct {
		Name     string
		Price    float64
		Weight   float64
		Stock    int
		MinStock int
	}{
		{
			Name:     "Gold - Individual Straw with Brush",
			Price:    45000,
			Weight:   30,
			Stock:    200,
			MinStock: 25,
		},
		{
			Name:     "Rose Gold - Individual Straw with Brush",
			Price:    45000,
			Weight:   30,
			Stock:    200,
			MinStock: 25,
		},
		{
			Name:     "Black - Individual Straw with Brush",
			Price:    42000,
			Weight:   30,
			Stock:    200,
			MinStock: 25,
		},
		{
			Name:     "Rainbow - Individual Straw with Brush",
			Price:    55000,
			Weight:   32,
			Stock:    150,
			MinStock: 20,
		},
		{
			Name:     "Gold - Couple Set (2 straws + 2 brushes)",
			Price:    85000,
			Weight:   65,
			Stock:    150,
			MinStock: 20,
		},
		{
			Name:     "Rose Gold - Couple Set (2 straws + 2 brushes)",
			Price:    85000,
			Weight:   65,
			Stock:    150,
			MinStock: 20,
		},
		{
			Name:     "Rainbow - Family Set (4 straws + 2 brushes)",
			Price:    175000,
			Weight:   140,
			Stock:    80,
			MinStock: 10,
		},
		{
			Name:     "Mixed Colors - Party Pack (8 straws, 3 brushes)",
			Price:    299000,
			Weight:   280,
			Stock:    50,
			MinStock: 5,
		},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Name,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("STR-CLR-%03d-%s", rand.Intn(1000), str)
		
		sku := entity.SKU{
			ProductID: product.ID,
			SKUCode:   skuCode,
			Price:     v.Price,
			Stock:     v.Stock,
			MinStock:  v.MinStock,
			Status:    entity.SKUStatusActive,
			Weight:    v.Weight,
		}
		db.Create(&sku)
		
		db.Create(&entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: value.ID,
		})
	}
	
	product.MinPrice = 42000
	product.MaxPrice = 299000
	db.Save(&product)
	
	if err := addStainlessStrawImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Premium Colored: %d options (Rp 42rb - 299rb)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 5: BULK PACKS FOR BUSINESS ====================
// Target: Cafes, restaurants, small businesses

func seedBulkPacks(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Bulk Stainless Steel Straws for Business - Cafe & Restaurant Pack",
		Description: `Business packs for cafes, restaurants, and small businesses that want to offer sustainable alternatives to customers. Includes multiple straws and cleaning brushes.

**☕ Perfect For:**
- Coffee shops
- Boba tea shops
- Juice bars
- Restaurants
- Hotels
- Office pantries

**📦 Bulk Options:**
- **8-Pack**: Small cafe starter
- **12-Pack**: Busy coffee shop
- **25-Pack**: Restaurant standard
- **50-Pack**: High-volume business

**✨ Business Features:**
- Durable construction for commercial use
- Easy to clean (dishwasher safe)
- Bulk pricing discounts
- Custom logo engraving available (minimum order)
- Replaceable parts available

**💰 Cost Savings:**
Switch once, save thousands on disposable straws annually.`,
		IsPublished:      true,
		Tags:             pq.StringArray{"bulk", "grosir", "business", "cafe", "restaurant", "commercial", "wholesale"},
		Slug:             "bulk-stainless-steel-straws-business-cafe-restaurant",
		MainImageURL:     "https://down-id.img.susercontent.com/file/29c865f330b5467a2f4cc5202f5f8331.webp",
		MainImagePublicID: "bulk-straws-main",
		AverageRating:    4.7,
		ReviewCount:      67,
		SoldCount:        234,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Business Pack",
	}
	db.Create(&variantType)
	
	variants := []struct {
		Name     string
		Price    float64
		Weight   float64
		Stock    int
		MinStock int
	}{
		{
			Name:     "Cafe Starter - 8 Straws + 2 Brushes",
			Price:    199000,
			Weight:   280,
			Stock:    50,
			MinStock: 5,
		},
		{
			Name:     "Coffee Shop - 12 Straws + 3 Brushes",
			Price:    279000,
			Weight:   420,
			Stock:    40,
			MinStock: 4,
		},
		{
			Name:     "Restaurant Standard - 25 Straws + 5 Brushes",
			Price:    499000,
			Weight:   875,
			Stock:    30,
			MinStock: 3,
		},
		{
			Name:     "High Volume - 50 Straws + 10 Brushes",
			Price:    899000,
			Weight:   1750,
			Stock:    20,
			MinStock: 2,
		},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Name,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("STR-BLK-%03d-%s", rand.Intn(1000), str)
		
		sku := entity.SKU{
			ProductID: product.ID,
			SKUCode:   skuCode,
			Price:     v.Price,
			Stock:     v.Stock,
			MinStock:  v.MinStock,
			Status:    entity.SKUStatusActive,
			Weight:    v.Weight,
		}
		db.Create(&sku)
		
		db.Create(&entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: value.ID,
		})
	}
	
	product.MinPrice = 199000
	product.MaxPrice = 899000
	db.Save(&product)
	
	if err := addStainlessStrawImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Bulk Packs: %d options (Rp 199rb - 899rb)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 6: TRAVEL SETS ====================
// Target: Frequent travelers, commuters, on-the-go professionals

func seedTravelSets(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Travel Stainless Steel Straw Set - Compact Case Included",
		Description: `Specially designed for people on the go. These travel sets come in a compact, portable case that fits in your bag, purse, or pocket.

**🧳 What's Included:**
- 1 Straight straw (regular or wide)
- 1 Bent straw (regular or wide)
- 1 Mini cleaning brush
- 1 Travel case (aluminum or bamboo)
- 2 Silicone tips

**✨ Travel Features:**
- Compact case (fits anywhere)
- Hygienic storage
- Keychain option available
- Airplane carry-on friendly
- Perfect for office, travel, daily commute

**🎯 Perfect For:**
- Office workers
- Frequent travelers
- Students
- Commuters
- Coffee enthusiasts

Never use a plastic straw again, even when you're out!`,
		IsPublished:      true,
		Tags:             pq.StringArray{"travel", "portable", "compact", "on-the-go", "case", "office", "commuter"},
		Slug:             "travel-stainless-steel-straw-set-compact-case",
		MainImageURL:     "https://down-id.img.susercontent.com/file/d6bbcb2b9502fd90254fe7f35d788634.webp",
		MainImagePublicID: "travel-straw-main",
		AverageRating:    4.9,
		ReviewCount:      123,
		SoldCount:        567,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Travel Set",
	}
	db.Create(&variantType)
	
	variants := []struct {
		Name     string
		Price    float64
		Weight   float64
		Stock    int
		MinStock int
	}{
		{
			Name:     "Regular Set - Aluminum Case (6mm)",
			Price:    65000,
			Weight:   70,
			Stock:    300,
			MinStock: 40,
		},
		{
			Name:     "Wide Set - Aluminum Case (8mm)",
			Price:    70000,
			Weight:   75,
			Stock:    250,
			MinStock: 35,
		},
		{
			Name:     "Regular Set - Bamboo Case (6mm)",
			Price:    75000,
			Weight:   65,
			Stock:    250,
			MinStock: 35,
		},
		{
			Name:     "Deluxe Travel Set - Both Diameters + Keychain",
			Price:    110000,
			Weight:   95,
			Stock:    200,
			MinStock: 30,
		},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Name,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("STR-TRL-%03d-%s", rand.Intn(1000), str)
		
		sku := entity.SKU{
			ProductID: product.ID,
			SKUCode:   skuCode,
			Price:     v.Price,
			Stock:     v.Stock,
			MinStock:  v.MinStock,
			Status:    entity.SKUStatusActive,
			Weight:    v.Weight,
		}
		db.Create(&sku)
		
		db.Create(&entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: value.ID,
		})
	}
	
	product.MinPrice = 65000
	product.MaxPrice = 110000
	db.Save(&product)
	
	if err := addStainlessStrawImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Travel Sets: %d options (Rp 65rb - 110rb)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 7: STARTER KITS ====================
// Target: First-time buyers, hesitant consumers

func seedStarterKits(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Starter Kit - Try Reusable Straws for the First Time",
		Description: `Not sure if reusable straws are for you? Try our affordable starter kit! Everything you need to begin your plastic-free journey at an accessible price point.

**🎁 Kit Includes:**
- 1 Straight straw (6mm - most popular)
- 1 Cleaning brush
- 1 Simple cotton pouch
- Care instructions card

**✨ Why Start Here:**
- Lowest entry price
- Most popular size
- Learn if you prefer straight or bent (buy bent separately)
- Perfect for travel

**🌱 Your First Step:**
Join millions of Indonesians reducing plastic waste. One reusable straw replaces hundreds of disposables. Start today!

After you try this, upgrade to a full family set!`,
		IsPublished:      true,
		Tags:             pq.StringArray{"starter", "pemula", "trial", "first time", "beginner", "entry level"},
		Slug:             "starter-kit-try-reusable-straws-first-time",
		MainImageURL:     "https://down-id.img.susercontent.com/file/ac1cd9b9c0750b191865430c1cf4cd54@resize_w900_nl.webp",
		MainImagePublicID: "starter-straw-main",
		AverageRating:    4.6,
		ReviewCount:      345,
		SoldCount:        1234,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Starter Type",
	}
	db.Create(&variantType)
	
	variants := []struct {
		Name     string
		Price    float64
		Weight   float64
		Stock    int
		MinStock int
	}{
		{
			Name:     "Basic Starter - Straight Straw (6mm)",
			Price:    35000,
			Weight:   35,
			Stock:    500,
			MinStock: 75,
		},
		{
			Name:     "Basic Starter - Bent Straw (6mm)",
			Price:    37000,
			Weight:   36,
			Stock:    400,
			MinStock: 60,
		},
		{
			Name:     "Wide Starter - For Boba Lovers (8mm)",
			Price:    40000,
			Weight:   40,
			Stock:    300,
			MinStock: 50,
		},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Name,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("STR-STRT-%03d-%s", rand.Intn(1000), str)
		
		sku := entity.SKU{
			ProductID: product.ID,
			SKUCode:   skuCode,
			Price:     v.Price,
			Stock:     v.Stock,
			MinStock:  v.MinStock,
			Status:    entity.SKUStatusActive,
			Weight:    v.Weight,
		}
		db.Create(&sku)
		
		db.Create(&entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: value.ID,
		})
	}
	
	product.MinPrice = 35000
	product.MaxPrice = 40000
	db.Save(&product)
	
	if err := addStainlessStrawImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Starter Kits: %d options (Rp 35rb - 40rb)\n", len(variants))
	return nil
}

func createStrawCutleryCategory(db *gorm.DB) (*entity.Category, error) {
  var category entity.Category
	result := db.First(&category, 6)
	if result.Error != nil {
		// Create category if it doesn't exist
		category = entity.Category{
			Model:       gorm.Model{ID: 6},
			Name:        "Straw & Cutlery",
		}
		if err := db.Create(&category).Error; err != nil {
			return nil, err
		}
	}

	return &category, nil
}

func addStainlessStrawImages(db *gorm.DB, productID uint) error {
	imageURLs := []string{
		"https://down-id.img.susercontent.com/file/id-11134207-7r98v-lrwwhgxi0eo9a0.webp",
		"https://down-id.img.susercontent.com/file/d6bbcb2b9502fd90254fe7f35d788634.webp",
		"https://down-id.img.susercontent.com/file/29c865f330b5467a2f4cc5202f5f8331.webp",
	}
	
	for _, url := range imageURLs {
		img := entity.Image{
			ProductID: productID,
			ImageURL: url,
		}

		if err := db.Create(&img).Error; err != nil {
			return err
		}
	}

	return nil
}