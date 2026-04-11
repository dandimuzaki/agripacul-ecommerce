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

func SeedReusableBags(db *gorm.DB) error {
    rand.Seed(time.Now().UnixNano())
    
    // First, ensure categories exist
    category, err := createZeroWasteKitCategory(db)
    if err != nil {
        return err
    }
    
    // Seed each product line
    if err := seedMeshBagLine(db, category.ID); err != nil {
        return err
    }
    
    if err := seedCottonBagLine(db, category.ID); err != nil {
        return err
    }
    
    if err := seedNylonBagLine(db, category.ID); err != nil {
        return err
    }
    
    if err := seedBambooBagLine(db, category.ID); err != nil {
        return err
    }
    
    if err := seedBatikCollection(db, category.ID); err != nil {
        return err
    }
    
    if err := seedGiftSets(db, category.ID); err != nil {
        return err
    }
    
    fmt.Println("✅ Successfully seeded Reusable Bag Collection (6 product lines)")
    return nil
}

func createZeroWasteKitCategory(db *gorm.DB) (*entity.Category, error) {
    var category entity.Category
    result := db.First(&category, 1)
    if result.Error != nil {
        // Create Zero Waste Kit category if it doesn't exist
        category = entity.Category{
            Model:       gorm.Model{ID: 1},
            Name:        "Zero Waste Kit",
        }
        if err := db.Create(&category).Error; err != nil {
            return nil, err
        }
    }

    return &category, nil
}

// Product 1: Standard Mesh Bags (Most Popular)
func seedMeshBagLine(db *gorm.DB, categoryID uint) error {
    product := entity.Product{
        CategoryID: categoryID,
        Name:       "EcoMesh Reusable Produce Bags - Standard rPET - Set of 8",
        Description: `Our best-selling mesh produce bags, made from recycled plastic bottles. Perfect for everyday grocery shopping at supermarkets and traditional markets (pasar). Each bag features printed tare weight for easy checkout.

**What's Included:**
- 2x Extra Small (20x25cm): Garlic, chillies, mushrooms
- 2x Small (25x30cm): Apples, oranges, tomatoes
- 2x Medium (30x40cm): Leafy greens, broccoli, carrots
- 2x Large (40x50cm): Cabbage, cauliflower, pumpkins

**Features:**
- ✅ Made from recycled PET (rPET) - each set diverts 24 bottles from landfill
- ✅ Tare weight printed on durable tag
- ✅ Bright rainbow colors (prevents loss in laundry)
- ✅ Breathable mesh extends produce freshness
- ✅ Machine washable, cold cycle

**Sustainability Impact:**
- Lifetime: 2-3 years (500+ uses)
- Plastic saved: 1,500+ single-use bags
- Carbon footprint: 70% lower than plastic after 50 uses`,
        IsPublished: true,
        Tags:        pq.StringArray{"mesh bags", "rPET", "recycled", "produce bags", "grocery bags", "rainbow"},
        Slug:        "ecomesh-standard-rpet-produce-bags-set-of-8",
        MainImageURL: "https://m.media-amazon.com/images/I/91x6TvDCnvS._SX569_.jpg",
        MainImagePublicID: "mesh-bags-standard",
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }
    
    // Only color variants for standard line
    colors := []string{
        "Rainbow Mix", "Ocean Blues", "Forest Greens", "Sunset Gradient",
    }
    
    return createSimpleSKUs(db, product, colors, 159000, "MESH")
}

// Product 2: Organic Cotton Collection
func seedCottonBagLine(db *gorm.DB, categoryID uint) error {
    product := entity.Product{
        CategoryID: categoryID,
        Name:       "Organic Cotton Produce Bags - Unbleached Natural - Set of 6",
        Description: `For those seeking natural fibers, our organic cotton bags are the perfect choice. Unbleached and chemical-free, these bags are ideal for dry goods and bulk shopping.

**What's Included:**
- 2x Small Drawstring Bags (25x30cm)
- 2x Medium Drawstring Bags (30x40cm)
- 2x Large Drawstring Bags (40x50cm)
- 1x Cotton storage pouch

**Features:**
- ✅ GOTS certified organic cotton
- ✅ Unbleached, no chemical treatments
- ✅ Drawstring closure for secure shopping
- ✅ Lightweight and foldable
- ✅ Tare weight printed on fabric label

**Perfect For:**
- Dry goods (rice, beans, lentils)
- Nuts and seeds
- Bread and pastries
- Bulk spices`,
        IsPublished: true,
        Tags:        pq.StringArray{"organic cotton", "natural", "unbleached", "drawstring", "bulk bags"},
        Slug:        "organic-cotton-produce-bags-natural-set-of-6",
        MainImageURL: "https://m.media-amazon.com/images/I/51NRtt4kMnL._SY300_SX300_QL70_FMwebp_.jpg",
        MainImagePublicID: "cotton-bags-natural",
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }
    
    variants := []string{
        "Unbleached Natural", "Bleached White", 
    }
    
    return createSimpleSKUs(db, product, variants, 179000, "CTN")
}

// Product 3: Premium Nylon Durability Line
func seedNylonBagLine(db *gorm.DB, categoryID uint) error {
    product := entity.Product{
        CategoryID: categoryID,
        Name:       "Heavy-Duty Nylon Mesh Bags - Premium Durability - Set of 5",
        Description: `For heavy shoppers and bulk buyers. These premium nylon bags offer superior durability while maintaining the breathability needed for produce.

**What's Included:**
- 1x Extra Small (heavy-duty)
- 1x Small (reinforced)
- 1x Medium (reinforced)
- 2x Large (heavy-duty for bulk)

**Features:**
- ✅ Ultra-strong nylon construction
- ✅ Reinforced seams (tested to 10kg)
- ✅ Water-resistant coating
- ✅ Double-stitched handles
- ✅ Reflective trim for visibility

**Best For:**
- Heavy produce (potatoes, onions, melons)
- Bulk grains (25kg rice/beans)
- Wet produce (washed greens)
- Farmer's market heavy loads`,
        IsPublished: true,
        Tags:        pq.StringArray{"nylon", "heavy-duty", "durable", "bulk", "reinforced"},
        Slug:        "heavy-duty-nylon-mesh-bags-premium-set-of-5",
        MainImageURL: "https://m.media-amazon.com/images/I/51jGHGjAz0L._SY300_SX300_QL70_FMwebp_.jpg",
        MainImagePublicID: "nylon-heavy-duty",
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }
    
    colors := []string{
        "Black (Heavy-Duty)", "Navy (Heavy-Duty)", "Charcoal (Heavy-Duty)",
    }
    
    return createSimpleSKUs(db, product, colors, 229000, "NYL")
}

// Product 4: Bamboo Blend Eco-Luxe
func seedBambooBagLine(db *gorm.DB, categoryID uint) error {
    product := entity.Product{
        CategoryID: categoryID,
        Name:       "Bamboo Cotton Blend Bags - Eco-Luxe Collection - Set of 4",
        Description: `The ultimate in sustainable luxury. Our bamboo-cotton blend bags combine the softness of bamboo with the durability of organic cotton.

**What's Included:**
- 4x Assorted sizes in matching earth tones
- 1x Bamboo storage box
- 1x Care guide

**Features:**
- ✅ 60% Bamboo, 40% Organic Cotton
- ✅ Naturally antimicrobial
- ✅ Ultra-soft feel
- ✅ Biodegradable
- ✅ Luxury gift packaging

**Perfect For:**
- Farmers market shopping
- Eco-conscious gift giving
- Specialty organic purchases`,
        IsPublished: true,
        Tags:        pq.StringArray{"bamboo", "luxury", "eco-luxe", "premium", "gift quality"},
        Slug:        "bamboo-cotton-blend-bags-eco-luxe-set-of-4",
        MainImageURL: "https://m.media-amazon.com/images/I/81uw8HQownL._SX569_.jpg",
        MainImagePublicID: "bamboo-eco-luxe",
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }
    
    variants := []string{
        "Earth Tones (Terracotta)", "Earth Tones (Sage)", "Earth Tones (Sand)",
    }
    
    return createSimpleSKUs(db, product, variants, 289000, "BMB")
}

// Product 5: Batik Cultural Collection
func seedBatikCollection(db *gorm.DB, categoryID uint) error {
    product := entity.Product{
        CategoryID: categoryID,
        Name:       "Batik Pattern Reusable Bags - Indonesian Heritage Collection - Set of 3",
        Description: `Celebrate Indonesian heritage with our Batik-patterned reusable bags. Made from premium organic cotton with traditional motifs.

**What's Included:**
- 1x Medium bag with Kawung pattern
- 1x Medium bag with Parang pattern
- 1x Large bag with Ceplok pattern
- 1x Informational card about each motif

**Pattern Meanings:**
- **Kawung**: Represents purity and perfection
- **Parang**: Symbolizes strength and continuity
- **Ceplok**: Represents harmony and balance

**Features:**
- ✅ Premium organic cotton base
- ✅ Traditional hand-stamped patterns
- ✅ Colorfast dyes (cold wash)
- ✅ Double-stitched construction
- ✅ Made in collaboration with Indonesian artisans`,
        IsPublished: true,
        Tags:        pq.StringArray{"batik", "indonesian", "cultural", "traditional", "artisan", "heritage"},
        Slug:        "batik-pattern-reusable-bags-heritage-collection",
        MainImageURL: "https://m.media-amazon.com/images/I/71kXnF4b2YL._SX569_.jpg",
        MainImagePublicID: "batik-collection",
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }
    
    variants := []string{
        "Kawung - Indigo", "Kawung - Sogan", "Parang - Indigo", "Parang - Sogan",
    }
    
    return createSimpleSKUs(db, product, variants, 259000, "BTK")
}

// Product 6: Gift Sets
func seedGiftSets(db *gorm.DB, categoryID uint) error {
    product := entity.Product{
        CategoryID: categoryID,
        Name:       "Complete Zero Waste Gift Set - Starter Kit with Accessories",
        Description: `The perfect gift for someone starting their zero waste journey. Includes everything needed to eliminate single-use plastic from grocery shopping.

**Complete Set Includes:**
- 8x Mesh produce bags (assorted sizes)
- 3x Cotton bulk bags (for dry goods)
- 2x Reusable shopping totes
- 1x Wooden tare weight tags
- 1x Recipe book (20 plastic-free recipes)
- 1x Jute gift box
- 1x "My Zero Waste Journey" tracker

**Perfect For:**
- Housewarming gifts
- Eco-conscious friends
- Corporate sustainable gifts
- Birthday presents

**Impact Promise:**
By gifting this set, you're helping eliminate ~1,800 single-use plastic bags from entering Indonesian landfills.`,
        IsPublished: true,
        Tags:        pq.StringArray{"gift set", "starter kit", "complete", "zero waste", "corporate gift"},
        Slug:        "complete-zero-waste-gift-set-starter-kit",
        MainImageURL: "https://m.media-amazon.com/images/I/81ih+bX2KlL._SX569_.jpg",
        MainImagePublicID: "gift-set-complete",
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }
    
    variants := []string{
        "Standard Gift Box", "Premium Wooden Box", "Corporate Gift (Bulk)",
    }
    
    return createGiftSetSKUs(db, product, variants)
}

// Helper function for simple product SKUs
func createSimpleSKUs(db *gorm.DB, product entity.Product, variants []string, basePrice float64, prefix string) error {
    var skus []entity.SKU
    var minPrice, maxPrice float64 = basePrice, basePrice
    
    for i, variant := range variants {
        // Variant type
        variantType := entity.VariantType{
            ProductID: product.ID,
            Name:      "Style",
        }
        db.Create(&variantType)
        
        // Variant value
        variantValue := entity.VariantValue{
            VariantTypeID: variantType.ID,
            Value:         variant,
        }
        db.Create(&variantValue)
        
        // Price variation (±10% for different colors)
        price := basePrice
        if i%2 == 0 {
            price = basePrice * 0.95 // 5% off for some colors
        } else if i%3 == 0 {
            price = basePrice * 1.05 // 5% premium for others
        }
        price = float64(int(price/1000)) * 1000

        str, _ := utils.GenerateRandomString(5)
        
        sku := entity.SKU{
            ProductID: product.ID,
            SKUCode:   fmt.Sprintf("%s-%d-%s-%s", prefix, product.ID, variant[:3], str),
            Price:     price,
            Stock:     100 + rand.Intn(300),
            MinStock:  20,
            Status:    entity.SKUStatusActive,
            Weight:    250,
        }
        
        if err := db.Create(&sku).Error; err != nil {
            return err
        }
        
        skuVariant := entity.SKUVariantValue{
            SKUID:          sku.ID,
            VariantValueID: variantValue.ID,
        }
        db.Create(&skuVariant)
        
        skus = append(skus, sku)
        
        if price < minPrice {
            minPrice = price
        }
        if price > maxPrice {
            maxPrice = price
        }
    }
    
    // Update product prices
    product.MinPrice = minPrice
    product.MaxPrice = maxPrice
    db.Save(&product)
    
    // Add images
    if err := addProductImages(db, product.ID); err != nil {
			return err
		}
    
    fmt.Printf("  ✅ Created %s: %d variants (Rp %.0f - Rp %.0f)\n", 
        product.Name, len(skus), minPrice, maxPrice)
    
    return nil
}

// Special handling for gift sets
func createGiftSetSKUs(db *gorm.DB, product entity.Product, variants []string) error {
    priceMap := map[string]float64{
        "Standard Gift Box":     299000,
        "Premium Wooden Box":    399000,
        "Corporate Gift (Bulk)": 259000, // Bulk pricing
    }
    
    var skus []entity.SKU
    var minPrice, maxPrice float64 = 999999, 0
    
    for _, variant := range variants {
        variantType := entity.VariantType{
            ProductID: product.ID,
            Name:      "Gift Package",
        }
        db.Create(&variantType)
        
        variantValue := entity.VariantValue{
            VariantTypeID: variantType.ID,
            Value:         variant,
        }
        db.Create(&variantValue)
        
        price := priceMap[variant]
        
        // Stock levels vary by package type
        stock := 50
        if variant == "Corporate Gift (Bulk)" {
            stock = 200 // Bulk stock
        }

        str, _ := utils.GenerateRandomString(5)
        
        sku := entity.SKU{
            ProductID: product.ID,
            SKUCode:   fmt.Sprintf("GIFT-%s-%d-%s", variant[:4], rand.Intn(1000), str),
            Price:     price,
            Stock:     stock,
            MinStock:  10,
            Status:    entity.SKUStatusActive,
            Weight:    800,
        }
        
        if err := db.Create(&sku).Error; err != nil {
            return err
        }
        
        skuVariant := entity.SKUVariantValue{
            SKUID:          sku.ID,
            VariantValueID: variantValue.ID,
        }
        db.Create(&skuVariant)
        
        skus = append(skus, sku)
        
        if price < minPrice {
            minPrice = price
        }
        if price > maxPrice {
            maxPrice = price
        }
    }
    
    product.MinPrice = minPrice
    product.MaxPrice = maxPrice
    db.Save(&product)
    
    // Add images
    if err := addProductImages(db, product.ID); err != nil {
			return err
		}
    
    fmt.Printf("  ✅ Created %s: %d gift options (Rp %.0f - Rp %.0f)\n", 
        product.Name, len(skus), minPrice, maxPrice)
    
    return nil
}

func addProductImages(db *gorm.DB, productID uint) error {
	// Upload images
	imageURLs := []string{
		"https://m.media-amazon.com/images/I/51NRtt4kMnL._SY300_SX300_QL70_FMwebp_.jpg",
		"https://m.media-amazon.com/images/I/51jGHGjAz0L._SY300_SX300_QL70_FMwebp_.jpg",
		"https://m.media-amazon.com/images/I/81uw8HQownL._SX569_.jpg",
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