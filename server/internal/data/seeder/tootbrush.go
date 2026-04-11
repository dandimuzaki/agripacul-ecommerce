package seeder

import (
	"debian-ecommerce/internal/data/entity"
	"fmt"
	"math/rand"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func SeedBambooToothbrush(db *gorm.DB) error {
    rand.Seed(time.Now().UnixNano())
    
    // First, ensure category exists (assuming category ID 2 is Personal Care)
    // If not, create it first
    var category entity.Category
    result := db.First(&category, 2)
    if result.Error != nil {
        // Create Personal Care category if it doesn't exist
        category = entity.Category{
            Model:       gorm.Model{ID: 2},
            Name:        "Personal Care",
        }
        if err := db.Create(&category).Error; err != nil {
            return err
        }
    }

    // Create the main product - Bamboo Toothbrush
    product := entity.Product{
        CategoryID:       2,
        Name:             "EcoBamboo Toothbrush - Sustainable Oral Care",
        Description:       `Experience the perfect blend of sustainability and oral hygiene with our premium bamboo toothbrush. Made from 100% organic Moso bamboo, this eco-friendly alternative helps reduce plastic waste while providing effective cleaning. The handle is naturally antibacterial, biodegradable, and FSC-certified. Soft bristles are gentle on gums yet effective at plaque removal. Packaged in 100% compostable cardboard. Join the zero-waste movement one brush at a time. 🌱`,
        IsPublished:      true,
        Tags:             pq.StringArray{"bamboo", "toothbrush", "eco-friendly", "zero waste", "sustainable", "oral care", "biodegradable", "plastic-free"},
        Slug:             "ecobamboo-toothbrush-sustainable-oral-care",
        MainImageURL:     "https://down-id.img.susercontent.com/file/50391bf1e0c2a0ab3613fbfdff911733.webp",
        MainImagePublicID: "bamboo-toothbrush-main",
        AverageRating:    4.8,
        ReviewCount:      1245,
        SoldCount:        5678,
        MinPrice:         0, // Will be updated after SKUs are created
        MaxPrice:         0, // Will be updated after SKUs are created
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }

    // Create Variant Types
    // Variant 1: Bristle Type
    bristleType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Bristle Type",
    }
    if err := db.Create(&bristleType).Error; err != nil {
        return err
    }

    // Variant 2: Bristle Color
    colorType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Bristle Color",
    }
    if err := db.Create(&colorType).Error; err != nil {
        return err
    }

    // Create Variant Values
    // Bristle Type values
    bristleValues := []entity.VariantValue{
        {
            VariantTypeID: bristleType.ID,
            Value:         "Soft - Premium Plant-Based Nylon",
        },
        {
            VariantTypeID: bristleType.ID,
            Value:         "Medium - Activated Charcoal Infused",
        },
        {
            VariantTypeID: bristleType.ID,
            Value:         "Extra Soft - Sensitive Gum Care",
        },
        {
            VariantTypeID: bristleType.ID,
            Value:         "Kids - Ultra Soft (Ages 3-12)",
        },
    }
    
    for i := range bristleValues {
        if err := db.Create(&bristleValues[i]).Error; err != nil {
            return err
        }
    }

    // Bristle Color values
    colorValues := []entity.VariantValue{
        {
            VariantTypeID: colorType.ID,
            Value:         "Natural White",
        },
        {
            VariantTypeID: colorType.ID,
            Value:         "Charcoal Black",
        },
        {
            VariantTypeID: colorType.ID,
            Value:         "Rainbow Mix (Multicolor)",
        },
        {
            VariantTypeID: colorType.ID,
            Value:         "Pastel Pink",
        },
        {
            VariantTypeID: colorType.ID,
            Value:         "Forest Green",
        },
    }
    
    for i := range colorValues {
        if err := db.Create(&colorValues[i]).Error; err != nil {
            return err
        }
    }

    // Create SKUs with different combinations and realistic pricing
    // Price mapping based on bristle type and color
    priceMap := map[string]float64{
        "Soft - Premium Plant-Based Nylon":         21900,
        "Medium - Activated Charcoal Infused":      26900,
        "Extra Soft - Sensitive Gum Care":          24900,
        "Kids - Ultra Soft (Ages 3-12)":            17900,
    }

    // Color premium adjustments (some colors are premium)
    colorPremium := map[string]float64{
        "Natural White":      0,
        "Charcoal Black":     5000,
        "Rainbow Mix":        8000,
        "Pastel Pink":        2000,
        "Forest Green":       2000,
    }

    // Bundle discounts for multi-packs
    bundleMultiplier := map[int]float64{
        1: 1.0,   // Single - full price
        3: 0.9,   // Pack of 3 - 10% off
        6: 0.85,  // Family pack of 6 - 15% off
        12: 0.8,  // Bulk pack of 12 - 20% off
    }

    var allSKUs []entity.SKU
    var minPrice float64 = 999999
    var maxPrice float64 = 0

    // Generate SKUs for all combinations
    for _, bristle := range bristleValues {
        basePrice := priceMap[bristle.Value]
        
        for _, color := range colorValues {
            // Skip some incompatible combinations (e.g., Kids brushes don't come in all colors)
            if bristle.Value == "Kids - Ultra Soft (Ages 3-12)" && 
               (color.Value == "Charcoal Black" || color.Value == "Rainbow Mix") {
                continue // Kids brushes only in fun colors
            }
            
            // For each combination, create multiple pack sizes
            for packSize, multiplier := range bundleMultiplier {
                // Calculate final price
                colorAdjustment := colorPremium[color.Value]
                basePriceWithColor := basePrice + colorAdjustment
                
                // Pack price with volume discount
                packPrice := basePriceWithColor * float64(packSize) * multiplier
                
                // Round to nearest 100 Rupiah (common in Indonesia)
                packPrice = float64(int(packPrice/100)) * 100
                
                // Generate SKU code
                skuCode := fmt.Sprintf("BBT-%s-%s-%dpcs-%d", 
                    abbreviateBristle(bristle.Value),
                    abbreviateColor(color.Value),
                    packSize,
                    rand.Intn(1000))
                
                // Calculate sale price (sometimes on promotion)
                var salePrice *float64
                if rand.Float64() < 0.3 { // 30% chance of being on sale
                    discount := 0.85 + (rand.Float64() * 0.10) // 15-25% off
                    sp := packPrice * discount
                    sp = float64(int(sp/100)) * 100
                    salePrice = &sp
                }
                
                // Determine stock levels
                stock := 0
                if packSize == 1 {
                    stock = 500 + rand.Intn(500) // Singles: 500-1000 stock
                } else if packSize == 3 {
                    stock = 200 + rand.Intn(300) // Pack of 3: 200-500 stock
                } else if packSize == 6 {
                    stock = 50 + rand.Intn(150)   // Pack of 6: 50-200 stock
                } else {
                    stock = 20 + rand.Intn(80)    // Pack of 12: 20-100 stock
                }
                
                sku := entity.SKU{
                    ProductID: product.ID,
                    SKUCode:   skuCode,
                    Price:     packPrice,
                    SalePrice: salePrice,
                    Stock:     stock,
                    MinStock:  20,
                    Status:    entity.SKUStatusActive,
                    Weight:    20 * float64(packSize), // ~20g per brush
                }
                
                if err := db.Create(&sku).Error; err != nil {
                    return err
                }
                
                // Track min/max prices
                if packPrice < minPrice {
                    minPrice = packPrice
                }
                if packPrice > maxPrice {
                    maxPrice = packPrice
                }
                
                // Link SKU to variant values
                skuVariant1 := entity.SKUVariantValue{
                    SKUID:          sku.ID,
                    VariantValueID: bristle.ID,
                }
                db.Create(&skuVariant1)
                
                skuVariant2 := entity.SKUVariantValue{
                    SKUID:          sku.ID,
                    VariantValueID: color.ID,
                }
                db.Create(&skuVariant2)
                
                allSKUs = append(allSKUs, sku)
            }
        }
    }

    // Update product with min and max prices
    product.MinPrice = minPrice
    product.MaxPrice = maxPrice
    db.Save(&product)

		// Upload images
		imageURLs := []string{
			"https://down-id.img.susercontent.com/file/6b45fbd3d6d43d3c936332d192ac6f3f.webp",
			"https://down-id.img.susercontent.com/file/6fb6f1f75df7e98a3f0aac3b97919402.webp",
			"https://down-id.img.susercontent.com/file/a96423758721fd03c73f14f68fa2cbc6.webp",
		}

		for _, url := range imageURLs {
			img := entity.Image{
				ProductID: product.ID,
				ImageURL: url,
			}

			db.Create(&img)
		}

    fmt.Printf("✅ Successfully seeded bamboo toothbrush product with %d SKU variants\n", len(allSKUs))
    fmt.Printf("   Price range: Rp %.0f - Rp %.0f\n", minPrice, maxPrice)
    
    return nil
}

// Helper functions
func abbreviateBristle(bristle string) string {
    switch {
    case bristle == "Soft - Premium Plant-Based Nylon":
        return "SFT"
    case bristle == "Medium - Activated Charcoal Infused":
        return "CHR"
    case bristle == "Extra Soft - Sensitive Gum Care":
        return "EXS"
    case bristle == "Kids - Ultra Soft (Ages 3-12)":
        return "KID"
    default:
        return "STD"
    }
}

func abbreviateColor(color string) string {
    switch {
    case color == "Natural White":
        return "WHT"
    case color == "Charcoal Black":
        return "BLK"
    case color == "Rainbow Mix":
        return "RNB"
    case color == "Pastel Pink":
        return "PNK"
    case color == "Forest Green":
        return "GRN"
    default:
        return "CLR"
    }
}