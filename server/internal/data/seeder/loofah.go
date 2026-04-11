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

// SeedLoofahCollection seeds the complete natural loofah product line
func SeedLoofahCollection(db *gorm.DB) error {
	rand.Seed(time.Now().UnixNano())
	
	// Create specialized categories for loofah products
	category, err := createPersonalCareCategory(db)
	if err != nil {
		return err
	}
	
	// Seed product lines (7 products instead of 1 mega-product)
	products := []struct {
		name     string
		category uint
		fn       func(*gorm.DB, uint) error
	}{
		{"Whole Traditional Loofah", category.ID, seedWholeLoofah},
		{"Loofah Pads - Siap Pakai", category.ID, seedLoofahPads},
		{"Loofah with Cotton Handle", category.ID, seedLoofahWithHandle},
		{"Kitchen Multipack Loofah", category.ID, seedKitchenLoofah},
		{"Premium Facial Loofah", category.ID, seedFacialLoofah},
		{"Loofah Scrub Mitts", category.ID, seedScrubMitts},
		{"Variety Pack - Mixed Loofah", category.ID, seedVarietyPack},
	}
	
	for _, p := range products {
		if err := p.fn(db, p.category); err != nil {
			return fmt.Errorf("failed to seed %s: %v", p.name, err)
		}
	}
	
	fmt.Println("✅ Successfully seeded Natural Loofah Collection (7 product lines)")
	return nil
}

// ==================== PRODUCT LINE 1: WHOLE TRADITIONAL LOOFAH ====================
// Target: Traditional bath users, people who want to cut their own size
// Based on Kini Bumi and Mahadewi Bali pricing

func seedWholeLoofah(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Whole Natural Loofah - Gambas Kering Tradisional",
		Description: `The traditional Indonesian loofah, harvested and sun-dried naturally. No chemicals, no bleaching—just pure plant fiber. You can cut it to your desired size for bath or kitchen use.

**🌱 What You Get:**
A complete dried Luffa fruit, harvested at peak maturity and sun-dried for 7-10 days. The natural fibers create a gentle yet effective exfoliating surface.

**✨ Features:**
- ✅ 100% natural, no chemicals
- ✅ Sun-dried in Java/Bali
- ✅ Can be cut into multiple pieces
- ✅ Lasts 3-6 months with proper care
- ✅ Fully compostable at end of life

**📋 How To Use:**
1. Soak in warm water for 5-10 minutes to soften
2. Cut into desired sizes (one whole loofah can make 2-4 pieces)
3. Use with soap for bath or kitchen
4. Rinse and hang to dry after use

**🌍 Supporting Local Farmers:**
Each loofah is grown by small-scale farmers in Java, providing sustainable income and promoting agricultural diversity.`,
		IsPublished:      true,
		Tags:             pq.StringArray{"whole loofah", "gambas utuh", "traditional", "natural sponge", "unbleached", "sun-dried", "local farmers"},
		Slug:             "whole-natural-loofah-gambas-kering-tradisional",
		MainImageURL:     "https://down-id.img.susercontent.com/file/id-11134207-7r98v-lurepfihsxwxf0@resize_w900_nl.webp",
		MainImagePublicID: "whole-loofah-main",
		AverageRating:    4.8,
		ReviewCount:      234,
		SoldCount:        987,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	// Create variant type (size)
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Size",
	}
	db.Create(&variantType)
	
	// Whole loofah variants - based on actual market sizes
	variants := []struct {
		Size        string
		LengthCm    int
		Price       float64
		Weight      float64
		Stock       int
		MinStock    int
		Description string
	}{
		{
			Size:        "Small (12-18 cm)",
			LengthCm:    15,
			Price:       15000,
			Weight:      30,
			Stock:       300,
			MinStock:    50,
			Description: "Perfect for travel or small hands",
		},
		{
			Size:        "Medium (19-25 cm)",
			LengthCm:    22,
			Price:       20000,
			Weight:      45,
			Stock:       500,
			MinStock:    75,
			Description: "Most popular size for bath use",
		},
		{
			Size:        "Large (26-30 cm)",
			LengthCm:    28,
			Price:       25000,
			Weight:      60,
			Stock:       250,
			MinStock:    40,
			Description: "Great for back exfoliation",
		},
		{
			Size:        "Extra Large (31-35 cm)",
			LengthCm:    33,
			Price:       30000,
			Weight:      75,
			Stock:       150,
			MinStock:    25,
			Description: "Can be cut into multiple pieces",
		},
		{
			Size:        "Jumbo (36-45 cm)",
			LengthCm:    40,
			Price:       35000,
			Weight:      90,
			Stock:       100,
			MinStock:    15,
			Description: "Best value - cut into 3-4 pieces",
		},
	}
	
	for _, v := range variants {
		// Create variant value
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Size,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		// Generate SKU code
		skuCode := fmt.Sprintf("LOOF-WHL-%dCM-%03d-%s", v.LengthCm, rand.Intn(1000), str)
		
		// 25% chance of sale price
		var salePrice *float64
		if rand.Float64() < 0.25 {
			sp := v.Price * 0.9 // 10% off
			sp = float64(int(sp/500)) * 500
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
	product.MinPrice = 15000
	product.MaxPrice = 35000
	db.Save(&product)
	
	if err := addLoofahImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Whole Loofah: %d sizes (Rp 15rb - 35rb)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 2: LOOFAH PADS ====================
// Target: Convenience seekers, ready-to-use products
// Based on Kini Bumi pricing: Rp 5,200 - 15,000

func seedLoofahPads(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Loofah Pads - Siap Pakai untuk Mandi & Dapur",
		Description: `Pre-cut loofah pads, ready to use right out of the package. No cutting required! Perfect for those who want convenience without compromising on sustainability.

**✨ Features:**
- ✅ Pre-cut and ready to use
- ✅ Consistent round/oval shape
- ✅ Gentle exfoliation for daily use
- ✅ Perfect size for hand grip
- ✅ Works with bar soap or liquid

**🎯 Multiple Uses:**
- Body exfoliation in shower
- Face cleansing (gentle pressure)
- Dish washing (for non-stick pans)
- Surface cleaning

**🌱 Zero Waste:**
When it wears out (3-4 months), simply compost it. No plastic packaging used.`,
		IsPublished:      true,
		Tags:             pq.StringArray{"loofah pads", "siap pakai", "ready to use", "circular", "bath pad", "dish pad"},
		Slug:             "loofah-pads-siap-pakai-mandi-dapur",
		MainImageURL:     "https://down-id.img.susercontent.com/file/id-11134207-7r98r-luroby2iizzw5b@resize_w900_nl.webp",
		MainImagePublicID: "loofah-pads-main",
		AverageRating:    4.7,
		ReviewCount:      156,
		SoldCount:        678,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Size",
	}
	db.Create(&variantType)
	
	variants := []struct {
		Size     string
		Diameter int
		Price    float64
		Weight   float64
		Stock    int
		MinStock int
	}{
		{"Small (8 cm)", 8, 8500, 15, 400, 60},
		{"Medium (10 cm)", 10, 12000, 22, 500, 80},
		{"Large (12 cm)", 12, 16000, 30, 300, 50},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Size,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("LOOF-PAD-%dCM-%03d-%s", v.Diameter, rand.Intn(1000), str)
		
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
	
	product.MinPrice = 8500
	product.MaxPrice = 16000
	db.Save(&product)
	
	if err := addLoofahImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Loofah Pads: %d sizes (Rp 8.5rb - 16rb)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 3: LOOFAH WITH COTTON HANDLE ====================
// Target: Elderly, children, people with mobility issues
// Based on Ecotools and Demi Bumi pricing

func seedLoofahWithHandle(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Loofah with Cotton Handle - Untuk Mandi Praktis",
		Description: `Loofah with a soft cotton handle for easier grip and extended reach. Perfect for elderly users, children, or anyone who wants a more comfortable bathing experience.

**✨ Features:**
- ✅ Soft cotton loop handle
- ✅ Easy to hang dry
- ✅ Reaches back easily
- ✅ Natural loofah pad
- ✅ Machine washable handle

**🎯 Perfect For:**
- Elderly with limited mobility
- Children learning to bathe
- People with arthritis
- Anyone who wants better grip

The cotton handle is removable and washable, extending the life of your loofah.`,
		IsPublished:      true,
		Tags:             pq.StringArray{"with handle", "cotton handle", "easy grip", "elderly", "children", "bath helper"},
		Slug:             "loofah-cotton-handle-mandi-praktis",
		MainImageURL:     "https://down-id.img.susercontent.com/file/id-11134207-7r990-luroby2ihlfg65.webp",
		MainImagePublicID: "loofah-handle-main",
		AverageRating:    4.9,
		ReviewCount:      89,
		SoldCount:        345,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Size",
	}
	db.Create(&variantType)
	
	variants := []struct {
		Size     string
		Price    float64
		Weight   float64
		Stock    int
		MinStock int
	}{
		{"Medium with Handle", 28000, 55, 200, 30},
		{"Large with Handle", 32000, 70, 150, 25},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Size,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("LOOF-HDL-%03d-%s", rand.Intn(1000), str)
		
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
	
	product.MinPrice = 28000
	product.MaxPrice = 32000
	db.Save(&product)
	
	if err := addLoofahImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Loofah with Handle: %d sizes (Rp 28rb - 32rb)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 4: KITCHEN MULTIPACK ====================
// Target: Households, eco-conscious kitchens
// Based on Segara Naturals and generic kitchen pricing

func seedKitchenLoofah(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Kitchen Loofah Multipack - Scrubber Alami untuk Dapur",
		Description: `The perfect zero-waste alternative to plastic kitchen scrubbers. These loofah pieces are cut and ready for dishwashing, counter cleaning, and vegetable scrubbing.

**✨ Why Loofah in the Kitchen:**
- ✅ Tough enough for pots and pans
- ✅ Gentle enough for non-stick cookware
- ✅ Naturally antimicrobial
- ✅ No plastic microfibers released
- ✅ Compostable at end of life

**📦 Multipack Options:**
- **3-pack**: Perfect for trial or small households
- **5-pack**: 2-3 month supply for average family
- **10-pack**: Stock up for the year

**🧼 Care Tips:**
Rinse thoroughly after each use and hang to dry. Replace every 2-3 months. When worn out, toss in the compost bin!

Supporting Indonesian farmers with every purchase.`,
		IsPublished:      true,
		Tags:             pq.StringArray{"kitchen", "dapur", "multipack", "dish scrubber", "pan scrubber", "vegetable washer"},
		Slug:             "kitchen-loofah-multipack-scrubber-alami",
		MainImageURL:     "https://down-id.img.susercontent.com/file/id-11134207-7r98r-luroby2iizzw5b@resize_w900_nl.webp",
		MainImagePublicID: "kitchen-loofah-main",
		AverageRating:    4.6,
		ReviewCount:      112,
		SoldCount:        456,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Pack Size",
	}
	db.Create(&variantType)
	
	variants := []struct {
		PackName  string
		Quantity  int
		Price     float64
		Weight    float64
		Stock     int
		MinStock  int
		PerUnit   string
	}{
		{"Starter Pack - 3 pieces", 3, 20000, 60, 300, 40, "Rp 6,700/pc"},
		{"Family Pack - 5 pieces", 5, 30000, 100, 250, 30, "Rp 6,000/pc"},
		{"Stock Up Pack - 10 pieces", 10, 55000, 200, 150, 20, "Rp 5,500/pc"},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.PackName,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("LOOF-KIT-%dPC-%03d-%s", v.Quantity, rand.Intn(1000), str)
		
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
	
	product.MinPrice = 20000
	product.MaxPrice = 55000
	db.Save(&product)
	
	if err := addLoofahImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Kitchen Multipack: %d options (Rp 20rb - 55rb)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 5: PREMIUM FACIAL LOOFAH ====================
// Target: Sensitive skin, facial care enthusiasts
// Based on premium brands like Ecotools and specialty sellers

func seedFacialLoofah(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Premium Facial Loofah - Ultra Soft untuk Wajah",
		Description: `Specially selected and processed for the delicate skin of your face. These loofah pieces are made from the softest inner fibers, providing gentle exfoliation without irritation.

**✨ Why Facial Loofah:**
- ✅ Ultra-fine fibers from the loofah's inner core
- ✅ Gentle enough for daily use
- ✅ Removes dead skin cells naturally
- ✅ Stimulates circulation
- ✅ No microplastics going down your drain

**🌱 How to Use:**
1. Wet with warm water until soft
2. Apply gentle circular motions to damp skin
3. Use with your favorite facial cleanser
4. Rinse thoroughly and hang to dry

**⚠️ Note:**
These are premium-grade loofah, hand-selected for consistency and softness. Perfect for those taking their zero-waste skincare seriously.`,
		IsPublished:      true,
		Tags:             pq.StringArray{"facial", "wajah", "ultra soft", "gentle", "exfoliation", "skincare", "premium"},
		Slug:             "premium-facial-loofah-ultra-soft-wajah",
		MainImageURL:     "https://down-id.img.susercontent.com/file/id-11134207-7r98r-lurepfihuchd07@resize_w900_nl.webp",
		MainImagePublicID: "facial-loofah-main",
		AverageRating:    4.5,
		ReviewCount:      67,
		SoldCount:        234,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Size",
	}
	db.Create(&variantType)
	
	variants := []struct {
		Size     string
		Price    float64
		Weight   float64
		Stock    int
		MinStock int
	}{
		{"Single Facial Pad", 25000, 15, 200, 30},
		{"Facial Set (2 pads + pouch)", 45000, 35, 100, 15},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Size,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("LOOF-FAC-%03d-%s", rand.Intn(1000), str)
		
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
	
	product.MinPrice = 25000
	product.MaxPrice = 45000
	db.Save(&product)
	
	if err := addLoofahImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Facial Loofah: %d options (Rp 25rb - 45rb)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 6: SCRUB MITTS ====================
// Target: Spa enthusiasts, deep exfoliation seekers

func seedScrubMitts(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Loofah Scrub Mitt - Eksfoliasi Seluruh Tubuh",
		Description: `A fabric mitt with an integrated loofah panel for easy full-body exfoliation. Just slip it on and scrub—no holding required!

**✨ Features:**
- ✅ Soft cotton exterior
- ✅ Loofah panel on palm side
- ✅ Elastic wrist for secure fit
- ✅ Reaches every part of your body
- ✅ Doubles as a washcloth

**🎯 Benefits:**
- Exfoliates back easily
- Removes dead skin before shaving
- Stimulates lymphatic drainage
- Reduces ingrown hairs
- Spa-quality treatment at home

**🧼 Care:**
Machine washable (gentle cycle) and line dry. Replace every 4-6 months.`,
		IsPublished:      true,
		Tags:             pq.StringArray{"scrub mitt", "sarung tangan", "exfoliating mitt", "body scrub", "spa", "back exfoliator"},
		Slug:             "loofah-scrub-mitt-eksfoliasi-tubuh",
		MainImageURL:     "https://down-id.img.susercontent.com/file/id-11134207-7r990-luroby2ihlfg65.webp",
		MainImagePublicID: "scrub-mitt-main",
		AverageRating:    4.8,
		ReviewCount:      45,
		SoldCount:        189,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Size",
	}
	db.Create(&variantType)
	
	variants := []struct {
		Size     string
		Price    float64
		Weight   float64
		Stock    int
		MinStock int
	}{
		{"Small/Medium", 35000, 80, 120, 20},
		{"Large/Extra Large", 40000, 90, 100, 15},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Size,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("LOOF-MIT-%03d-%s", rand.Intn(1000), str)
		
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
	
	if err := addLoofahImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Scrub Mitts: %d sizes (Rp 35rb - 40rb)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 7: VARIETY PACKS ====================
// Target: Gift shoppers, first-time buyers, variety seekers

func seedVarietyPack(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Loofah Variety Pack - Hadiah Zero Waste",
		Description: `The perfect gift for someone starting their zero-waste journey! A selection of our most popular loofah products in one beautiful package.

**🎁 Includes:**
- 1 Medium whole loofah
- 2 Loofah pads (medium)
- 1 Kitchen scrubber
- 1 Cotton storage bag
- Care instructions card

**✨ Perfect For:**
- Birthdays
- Housewarming gifts
- Zero waste starter kits
- Corporate sustainable gifts

**🌱 Eco Packaging:**
Comes in a reusable cotton bag with recycled paper tag. No plastic!`,
		IsPublished:      true,
		Tags:             pq.StringArray{"variety pack", "gift set", "hadiah", "starter kit", "zero waste gift", "mixed"},
		Slug:             "loofah-variety-pack-hadiah-zero-waste",
		MainImageURL:     "https://down-id.img.susercontent.com/file/id-11134207-7r98v-lurepfihsxwxf0@resize_w900_nl.webp",
		MainImagePublicID: "variety-pack-main",
		AverageRating:    4.9,
		ReviewCount:      34,
		SoldCount:        98,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Pack Type",
	}
	db.Create(&variantType)
	
	variants := []struct {
		PackName  string
		Price     float64
		Weight    float64
		Stock     int
		MinStock  int
	}{
		{"Starter Variety Pack", 35000, 120, 100, 15},
		{"Deluxe Gift Set (with mitt)", 60000, 200, 50, 10},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.PackName,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("LOOF-GFT-%03d-%s", rand.Intn(1000), str)
		
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
	product.MaxPrice = 60000
	db.Save(&product)
	
	if err := addLoofahImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Variety Packs: %d options (Rp 35rb - 60rb)\n", len(variants))
	return nil
}

func createPersonalCareCategory(db *gorm.DB) (*entity.Category, error) {
  var category entity.Category
	result := db.First(&category, 2)
	if result.Error != nil {
		// Create category if it doesn't exist
		category = entity.Category{
			Model:       gorm.Model{ID: 2},
      Name:        "Personal Care",
		}
		if err := db.Create(&category).Error; err != nil {
			return nil, err
		}
	}

	return &category, nil
}

func addLoofahImages(db *gorm.DB, productID uint) error {
	imageURLs := []string{
		"https://down-id.img.susercontent.com/file/id-11134207-7r98r-lurepfihuchd07@resize_w900_nl.webp",
		"https://down-id.img.susercontent.com/file/id-11134207-7r98r-luroby2iizzw5b@resize_w900_nl.webp",
		"https://down-id.img.susercontent.com/file/id-11134207-7r990-luroby2ihlfg65.webp",
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