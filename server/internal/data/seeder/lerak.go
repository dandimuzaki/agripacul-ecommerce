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

// SeedLerakCollection seeds the complete Buah Lerak (Soap Nut) product line
func SeedLerakCollection(db *gorm.DB) error {
	rand.Seed(time.Now().UnixNano())
	
	// Create specialized categories for lerak products
	category, err := createHomeCleaningCategory(db)
	if err != nil {
		return err
	}
	
	// Seed product lines (6 products instead of 1 mega-product)
	products := []struct {
		name     string
		category uint
		fn       func(*gorm.DB, uint) error
	}{
		{"Whole Dried Lerak", category.ID, seedWholeLerak},
		{"Lerak Powder", category.ID, seedLerakPowder},
		{"Liquid Lerak", category.ID, seedLiquidLerak},
		{"Lerak Paste", category.ID, seedLerakPaste},
		{"Trial Packs", category.ID, seedLerakTrialPacks},
		{"Bulk Economy", category.ID, seedLerakBulk},
	}
	
	for _, p := range products {
		if err := p.fn(db, p.category); err != nil {
			return fmt.Errorf("failed to seed %s: %v", p.name, err)
		}
	}
	
	fmt.Println("✅ Successfully seeded Buah Lerak Collection (6 product lines)")
	return nil
}

// ==================== PRODUCT LINE 1: WHOLE DRIED LERAK ====================
// Target: Traditional users, batik enthusiasts, eco-conscious households
// Based on Tokopedia data: 100g = Rp 12,500, 1kg = Rp 95,000 [citation:1][citation:5]

func seedWholeLerak(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Whole Dried Lerak - Natural Soap Nut for Batik & Laundry",
		Description: `Traditional Indonesian soap nut (*Sapindus rarak DC.*) harvested from Java and Sumatra. Used for centuries to clean batik and delicate fabrics without harsh chemicals.

**🌿 What is Lerak?**
Lerak contains natural saponins—plant-based surfactants that create gentle foam and clean effectively. Unlike chemical detergents containing SLS, lerak is completely biodegradable and safe for aquatic ecosystems [citation:4].

**✨ Why Choose Whole Lerak:**
- ✅ **100% Natural**: No chemicals, preservatives, or synthetic fragrances
- ✅ **Batik Safe**: Cleans without causing colors to fade [citation:4]
- ✅ **Hypoallergenic**: Perfect for baby clothes and sensitive skin [citation:1]
- ✅ **Economical**: Each fruit can be used 6-8 times [citation:1]
- ✅ **Zero Waste**: Used fruits can be composted [citation:1]

**📋 How to Use:**
1. Soak 7-10 fruits in water for 2 nights until soft
2. Remove seeds, mash the flesh to release saponin sap
3. Place in cloth bag and add to washing machine [citation:5]

**🌍 Sustainability:**
Lerak trees take 10-15 years before first harvest—each fruit is precious! Your purchase supports traditional farmers in Java and preserves Indonesian heritage.`,
		IsPublished:      true,
		Tags:             pq.StringArray{"whole lerak", "buah lerak utuh", "soap nut", "natural detergent", "batik care", "traditional", "saponin"},
		Slug:             "whole-dried-lerak-natural-soap-nut-batik-laundry",
		MainImageURL:     "https://down-id.img.susercontent.com/file/sg-11134201-22110-591szrfpjckva1_tn",
		MainImagePublicID: "whole-lerak-main",
		AverageRating:    4.9,
		ReviewCount:      456,
		SoldCount:        2345,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	// Create variant type (just pack size - simple!)
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Pack Size",
	}
	db.Create(&variantType)
	
	// Whole lerak variants - based on real market prices
	variants := []struct {
		Name        string
		WeightGrams int
		Price       float64
		ShipWeight  float64
		Stock       int
		MinStock    int
		Uses        string
	}{
		{
			Name:        "Trial Pack - 100g",
			WeightGrams: 100,
			Price:       12500,
			ShipWeight:  110,
			Stock:       500,
			MinStock:    50,
			Uses:        "~15-20 washes",
		},
		{
			Name:        "Small Pack - 250g",
			WeightGrams: 250,
			Price:       30000,
			ShipWeight:  260,
			Stock:       400,
			MinStock:    40,
			Uses:        "~40-50 washes",
		},
		{
			Name:        "Medium Pack - 500g",
			WeightGrams: 500,
			Price:       55000,
			ShipWeight:  510,
			Stock:       300,
			MinStock:    30,
			Uses:        "~80-100 washes",
		},
		{
			Name:        "Family Pack - 1kg",
			WeightGrams: 1000,
			Price:       95000,
			ShipWeight:  1020,
			Stock:       200,
			MinStock:    20,
			Uses:        "~160-200 washes (6 months)",
		},
		{
			Name:        "Bulk Economy - 5kg",
			WeightGrams: 5000,
			Price:       425000,
			ShipWeight:  5050,
			Stock:       50,
			MinStock:    5,
			Uses:        "~800-1000 washes (2+ years)",
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
		skuCode := fmt.Sprintf("LER-WHL-%dg-%03d-%s", v.WeightGrams, rand.Intn(1000), str)
		
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
			Weight:    v.ShipWeight,
		}
		db.Create(&sku)
		
		// Link SKU to variant
		db.Create(&entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: value.ID,
		})
	}
	
	// Update product price range
	product.MinPrice = 12500
	product.MaxPrice = 425000
	db.Save(&product)
	
	if err := addLerakImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Whole Dried Lerak: %d sizes (Rp 12.5rb - 425rb)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 2: LERAK POWDER ====================
// Target: Modern users who want convenience, no soaking required

func seedLerakPowder(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Lerak Powder - Bubuk Lerak Halus Siap Pakai",
		Description: `Ground lerak powder for instant use. No soaking required! Simply add powder directly to your washing machine or mix with water for immediate cleaning.

**✨ Benefits of Powder Form:**
- ✅ Instant use - no overnight soaking
- ✅ Easy to measure exact amounts
- ✅ Dissolves quickly in water
- ✅ Same natural saponin power
- ✅ Perfect for travel

**🌿 Processing:**
Our lerak is sun-dried traditionally then finely ground to preserve all natural saponins. No additives, no preservatives—just pure ground lerak.

**📋 How to Use:**
- **Top Load Machine**: Add 2-3 tablespoons directly to drum
- **Front Load**: Mix with water first, then add to detergent drawer
- **Hand Wash**: Dissolve 1 tablespoon in basin of water

**💚 Eco Packaging:**
Comes in kraft paper bags or jute sacks—no plastic!`,
		IsPublished:      true,
		Tags:             pq.StringArray{"lerak powder", "bubuk lerak", "ground lerak", "instant", "easy use", "natural detergent"},
		Slug:             "lerak-powder-bubuk-halus-siap-pakai",
		MainImageURL:     "https://benihbunbun.com/wp-content/uploads/2023/08/wp-1692842234781.jpg",
		MainImagePublicID: "lerak-powder-main",
		AverageRating:    4.7,
		ReviewCount:      234,
		SoldCount:        1234,
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
		Name        string
		WeightGrams int
		Price       float64
		ShipWeight  float64
		Stock       int
		MinStock    int
	}{
		{"Trial Pack - 100g", 100, 18000, 120, 400, 40},
		{"Small Pack - 250g", 250, 42000, 270, 300, 30},
		{"Medium Pack - 500g", 500, 75000, 520, 200, 25},
		{"Family Pack - 1kg", 1000, 135000, 1030, 150, 15},
		{"Bulk Economy - 5kg", 5000, 600000, 5100, 40, 4},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Name,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("LER-PWD-%dg-%03d-%s", v.WeightGrams, rand.Intn(1000), str)
		
		sku := entity.SKU{
			ProductID: product.ID,
			SKUCode:   skuCode,
			Price:     v.Price,
			Stock:     v.Stock,
			MinStock:  v.MinStock,
			Status:    entity.SKUStatusActive,
			Weight:    v.ShipWeight,
		}
		db.Create(&sku)
		
		db.Create(&entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: value.ID,
		})
	}
	
	product.MinPrice = 18000
	product.MaxPrice = 600000
	db.Save(&product)
	
	if err := addLerakImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Lerak Powder: %d sizes (Rp 18rb - 600rb)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 3: LIQUID LERAK ====================
// Target: Modern households, ready-to-use detergent users

func seedLiquidLerak(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Liquid Lerak Concentrate - Ready to Use Natural Detergent",
		Description: `Pre-made liquid lerak concentrate. Just pour and wash! Perfect for modern households who want the benefits of lerak without any preparation.

**✨ Features:**
- ✅ Ready to use - no soaking, no mixing
- ✅ Concentrated formula - a little goes a long way
- ✅ Preserved naturally with salt and citrus [citation:5]
- ✅ Works in all washing machines
- ✅ Same natural cleaning power

**📋 How to Use:**
- **Top Load**: 50-100ml per large load
- **Front Load**: 30-50ml per load
- **Hand Wash**: 20ml per basin

**🧪 Natural Preservation:**
We preserve our liquid lerak with salt and citrus peels (traditional method) - no synthetic preservatives! [citation:5]

**🌱 Zero Waste Options:**
Choose "Bring Your Own Container" for discount and less packaging.`,
		IsPublished:      true,
		Tags:             pq.StringArray{"liquid lerak", "lerak cair", "ready to use", "natural detergent", "concentrate"},
		Slug:             "liquid-lerak-concentrate-ready-to-use",
		MainImageURL:     "https://down-id.img.susercontent.com/file/id-11134207-7ras8-m1zivyovk1kt54.webp",
		MainImagePublicID: "liquid-lerak-main",
		AverageRating:    4.6,
		ReviewCount:      189,
		SoldCount:        876,
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
		Name       string
		VolumeMl   int
		Price      float64
		ShipWeight float64
		Stock      int
		MinStock   int
	}{
		{"Standard - 500ml", 500, 45000, 550, 300, 30},
		{"Economy - 1L", 1000, 75000, 1050, 200, 20},
		{"Family Refill - 5L", 5000, 325000, 5200, 50, 5},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Name,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("LER-LIQ-%dml-%03d-%s", v.VolumeMl, rand.Intn(1000), str)
		
		sku := entity.SKU{
			ProductID: product.ID,
			SKUCode:   skuCode,
			Price:     v.Price,
			Stock:     v.Stock,
			MinStock:  v.MinStock,
			Status:    entity.SKUStatusActive,
			Weight:    v.ShipWeight,
		}
		db.Create(&sku)
		
		db.Create(&entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: value.ID,
		})
	}
	
	product.MinPrice = 45000
	product.MaxPrice = 325000
	db.Save(&product)
	
	if err := addLerakImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Liquid Lerak: %d sizes (Rp 45rb - 325rb)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 4: LERAK PASTE ====================
// Target: Users who want concentrated form, easy to store

func seedLerakPaste(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Lerak Paste - Pasta Lerak Konsentrat",
		Description: `Semi-solid lerak concentrate. More concentrated than liquid, easier to use than whole fruits. Just scoop and dissolve!

**✨ Benefits:**
- ✅ Highly concentrated - small amount goes far
- ✅ Easy to scoop and measure
- ✅ Longer shelf life than liquid
- ✅ No preservatives needed
- ✅ Minimal packaging

**📋 How to Use:**
1. Scoop 1-2 tablespoons of paste
2. Dissolve in warm water
3. Add to washing machine or use for hand wash

**🌿 Traditional Method:**
Made by boiling lerak fruits until reduced to a thick paste, following traditional Javanese recipes passed down for generations.

Perfect for those who want the purity of whole lerak with the convenience of ready-to-use form.`,
		IsPublished:      true,
		Tags:             pq.StringArray{"lerak paste", "pasta lerak", "concentrate", "semi-solid", "traditional"},
		Slug:             "lerak-paste-pasta-konsentrat",
		MainImageURL:     "https://asset.kompas.com/crops/SBorOswHS4Kan2T_HemNCUTQpGk=/100x67:900x601/1200x800/data/photo/2023/02/24/63f8c83700d68.jpg",
		MainImagePublicID: "lerak-paste-main",
		AverageRating:    4.8,
		ReviewCount:      98,
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
		Name        string
		WeightGrams int
		Price       float64
		ShipWeight  float64
		Stock       int
		MinStock    int
	}{
		{"Small Jar - 250g", 250, 38000, 280, 250, 25},
		{"Medium Jar - 500g", 500, 68000, 530, 150, 15},
		{"Family Size - 1kg", 1000, 125000, 1040, 80, 8},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Name,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("LER-PST-%dg-%03d-%s", v.WeightGrams, rand.Intn(1000), str)
		
		sku := entity.SKU{
			ProductID: product.ID,
			SKUCode:   skuCode,
			Price:     v.Price,
			Stock:     v.Stock,
			MinStock:  v.MinStock,
			Status:    entity.SKUStatusActive,
			Weight:    v.ShipWeight,
		}
		db.Create(&sku)
		
		db.Create(&entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: value.ID,
		})
	}
	
	product.MinPrice = 38000
	product.MaxPrice = 125000
	db.Save(&product)
	
	if err := addLerakImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Lerak Paste: %d sizes (Rp 38rb - 125rb)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 5: TRIAL PACKS ====================
// Target: First-time buyers, hesitant customers, students

func seedLerakTrialPacks(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Trial Pack - Coba Lerak untuk Pemula",
		Description: `Not sure which lerak form works for you? Try our affordable trial packs! Perfect for first-time users who want to experience the benefits of natural lerak without committing to large sizes.

**🎁 Trial Options:**

**Whole Lerak Trial (100g):**
- ~15-20 washes worth
- Traditional experience
- Learn to prepare lerak yourself

**Lerak Powder Trial (100g):**
- Instant use, no prep
- Easy to measure
- Perfect for travel

**🌱 Why Try Lerak:**
- Chemical-free cleaning
- Safe for sensitive skin
- Preserves batik colors
- Zero waste alternative
- Supports local farmers

Start your natural detergent journey today!`,
		IsPublished:      true,
		Tags:             pq.StringArray{"trial", "sample", "pemula", "coba", "starter", "small pack"},
		Slug:             "trial-pack-coba-lerak-pemula",
		MainImageURL:     "https://down-id.img.susercontent.com/file/sg-11134201-22110-591szrfpjckva1_tn",
		MainImagePublicID: "lerak-trial-main",
		AverageRating:    4.5,
		ReviewCount:      345,
		SoldCount:        2345,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Trial Type",
	}
	db.Create(&variantType)
	
	variants := []struct {
		Name       string
		Price      float64
		ShipWeight float64
		Stock      int
		MinStock   int
	}{
		{
			Name:       "Whole Lerak Trial - 100g",
			Price:      12500,
			ShipWeight: 110,
			Stock:      500,
			MinStock:   50,
		},
		{
			Name:       "Lerak Powder Trial - 100g",
			Price:      18000,
			ShipWeight: 120,
			Stock:      400,
			MinStock:   40,
		},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Name,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("LER-TRL-%03d-%s", rand.Intn(1000), str)
		
		sku := entity.SKU{
			ProductID: product.ID,
			SKUCode:   skuCode,
			Price:     v.Price,
			Stock:     v.Stock,
			MinStock:  v.MinStock,
			Status:    entity.SKUStatusActive,
			Weight:    v.ShipWeight,
		}
		db.Create(&sku)
		
		db.Create(&entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: value.ID,
		})
	}
	
	product.MinPrice = 12500
	product.MaxPrice = 18000
	db.Save(&product)
	
	if err := addLerakImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Trial Packs: %d options (Rp 12.5rb - 18rb)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 6: BULK ECONOMY ====================
// Target: Businesses, boarding houses, heavy users

func seedLerakBulk(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Bulk Lerak - Ekonomis untuk Usaha & Keluarga Besar",
		Description: `Economy bulk packs for businesses, boarding houses, and large families. Best value pricing with minimal packaging.

**📦 Bulk Options:**

**Whole Lerak - 5kg Sack:**
- Traditional dried fruits
- Jute or cotton sack
- Best for long-term storage
- ~800-1000 washes

**Lerak Powder - 5kg Sack:**
- Ready-to-use powder
- Easy for staff to use
- Consistent measurements
- ~1000+ washes

**Liquid Lerak - 5L Jerry Can:**
- Ready to pour
- Perfect for restaurants
- Refill available (bring your own)

**💚 Perfect For:**
- Laundry businesses
- Batik workshops
- Boarding houses
- Eco-resorts
- Large families

**💰 Bulk Savings:**
Save 15-25% compared to buying smaller packs. Contact us for larger wholesale orders (25kg+).`,
		IsPublished:      true,
		Tags:             pq.StringArray{"bulk", "grosir", "ekonomis", "usaha", "bisnis", "keluarga besar"},
		Slug:             "bulk-lerak-ekonomis-usaha-keluarga-besar",
		MainImageURL:     "https://benihbunbun.com/wp-content/uploads/2023/08/wp-1692842234781.jpg",
		MainImagePublicID: "lerak-bulk-main",
		AverageRating:    4.7,
		ReviewCount:      67,
		SoldCount:        234,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Bulk Type",
	}
	db.Create(&variantType)
	
	variants := []struct {
		Name       string
		Price      float64
		ShipWeight float64
		Stock      int
		MinStock   int
	}{
		{
			Name:       "Whole Lerak - 5kg Sack",
			Price:      425000,
			ShipWeight: 5100,
			Stock:      50,
			MinStock:   5,
		},
		{
			Name:       "Lerak Powder - 5kg Sack",
			Price:      600000,
			ShipWeight: 5200,
			Stock:      30,
			MinStock:   3,
		},
		{
			Name:       "Liquid Lerak - 5L Jerry Can",
			Price:      325000,
			ShipWeight: 5300,
			Stock:      40,
			MinStock:   4,
		},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Name,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("LER-BLK-%03d-%s", rand.Intn(1000), str)
		
		// Bulk items less likely to be on sale
		var salePrice *float64
		if rand.Float64() < 0.1 { // Only 10% chance
			sp := v.Price * 0.95 // 5% off
			sp = float64(int(sp/1000)) * 1000
			salePrice = &sp
		}
		
		sku := entity.SKU{
			ProductID: product.ID,
			SKUCode:   skuCode,
			Price:     v.Price,
			SalePrice: salePrice,
			Stock:     v.Stock,
			MinStock:  v.MinStock,
			Status:    entity.SKUStatusActive,
			Weight:    v.ShipWeight,
		}
		db.Create(&sku)
		
		db.Create(&entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: value.ID,
		})
	}
	
	product.MinPrice = 325000
	product.MaxPrice = 600000
	db.Save(&product)
	
	if err := addLerakImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Bulk Economy: %d options (Rp 325rb - 600rb)\n", len(variants))
	return nil
}

func createHomeCleaningCategory(db *gorm.DB) (*entity.Category, error) {
  var category entity.Category
	result := db.First(&category, 3)
	if result.Error != nil {
		// Create category if it doesn't exist
		category = entity.Category{
			Model:       gorm.Model{ID: 3},
      Name:        "Home & Cleaning",
		}
		if err := db.Create(&category).Error; err != nil {
			return nil, err
		}
	}

	return &category, nil
}

func addLerakImages(db *gorm.DB, productID uint) error {
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