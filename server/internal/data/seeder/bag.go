package seeder

import (
	"debian-ecommerce/internal/data/entity"
	"fmt"
	"math/rand"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func SeedReusableProduceBags(db *gorm.DB) error {
    rand.Seed(time.Now().UnixNano())
    
    // First, ensure category exists (assuming category ID 1 is Zero Waste Kit)
    // If not, create it first
    var category entity.Category
    result := db.First(&category, 1)
    if result.Error != nil {
        // Create Zero Waste Kit category if it doesn't exist
        category = entity.Category{
            Model:       gorm.Model{ID: 1},
            Name:        "Zero Waste Kit",
        }
        if err := db.Create(&category).Error; err != nil {
            return err
        }
    }

    // Create the main product - Reusable Produce Bags
    product := entity.Product{
        CategoryID:       1,
        Name:             "EcoMesh Reusable Produce Bags - Set of 8 | Tare Weight Included",
        Description:       `Say goodbye to single-use plastic produce bags with our complete set of 8 reusable mesh bags. Perfect for fruits, vegetables, and bulk items at supermarkets, traditional markets (pasar), and farmers markets. Each bag features a unique tare weight printed on the tag, making checkout easy and accurate for cashiers.

**🌱 Why Choose Our Produce Bags?**
According to environmental research, each reusable bag that avoids disposal saves approximately 700 single-use plastic bags over its lifetime [citation:3]. With Indonesia's EPR regulations targeting 30% reduction in plastic waste by 2029, switching to reusable bags is both eco-conscious and forward-thinking [citation:1].

**🛍️ Complete Set Includes:**
- 2x Extra Small Bags (20x25cm) - Perfect for garlic, chillies, mushrooms
- 2x Small Bags (25x30cm) - Ideal for apples, oranges, tomatoes
- 2x Medium Bags (30x40cm) - Great for leafy greens, broccoli, carrots
- 2x Large Bags (40x50cm) - Perfect for cabbage, cauliflower, pumpkins
- 1x Cotton Stuff Sack for storage
- 1x Care Guide with tare weight instructions

**✨ Key Features:**
- ✅ **Tare Weight Printed**: Each bag's weight (6-12 grams) printed on durable tag for accurate checkout
- ✅ **High Visibility Colors**: Bright rainbow colors prevent bags from being lost in laundry or forgotten [citation:3]
- ✅ **Breathable Mesh Design**: Allows air circulation to keep produce fresh longer [citation:4]
- ✅ **Reinforced Stitching**: Double-stitched seams for durability up to 2+ years of weekly use
- ✅ **Machine Washable**: Cold wash, air dry (place in delicates bag to prevent loss)
- ✅ **Eco-Friendly Materials**: Made from recycled PET (rPET) - each bag diverts 2-3 plastic bottles from landfills

**💚 Sustainability Impact:**
- Break-even point: Only 10-20 uses needed to offset production footprint [citation:3]
- Lifespan: 2-3 years with proper care (500+ uses)
- Plastic saved: 1,500+ single-use bags over lifetime
- Carbon footprint: 70% lower than conventional plastic bags after 50 uses

**📋 Care Instructions:**
1. Shake out produce scraps after shopping
2. Spot clean when possible; machine wash only when necessary
3. Air dry completely before storing
4. Store in provided cotton sack near shopping list as visual reminder [citation:3]

Join thousands of Indonesian shoppers making the switch to sustainable grocery shopping!`,
        IsPublished:      true,
        Tags:             pq.StringArray{"produce bags", "reusable", "grocery bags", "mesh bags", "zero waste", "plastic-free", "sustainable", "eco-friendly", "tare weight", "bulk shopping", "pasar", "vegetable bags", "fruit bags", "rPET", "recycled material"},
        Slug:             "ecomesh-reusable-produce-bags-set-of-8-tare-weight",
        MainImageURL:     "https://m.media-amazon.com/images/I/91x6TvDCnvS._SX569_.jpg",
        MainImagePublicID: "produce-bags-main",
        AverageRating:    4.8,
        ReviewCount:      1243,
        SoldCount:        5678,
        MinPrice:         0, // Will be updated after SKUs
        MaxPrice:         0, // Will be updated after SKUs
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }

    // Create Variant Types
    // Variant 1: Material Type
    materialType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Material",
    }
    if err := db.Create(&materialType).Error; err != nil {
        return err
    }

    // Variant 2: Color / Pattern
    colorType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Color Pattern",
    }
    if err := db.Create(&colorType).Error; err != nil {
        return err
    }

    // Variant 3: Set Size (Number of bags)
    setSizeType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Set Size",
    }
    if err := db.Create(&setSizeType).Error; err != nil {
        return err
    }

    // Create Variant Values
    // Material Types
    materialValues := []entity.VariantValue{
        {
            VariantTypeID: materialType.ID,
            Value:         "Recycled PET (rPET) - From Plastic Bottles",
        },
        {
            VariantTypeID: materialType.ID,
            Value:         "Organic Cotton - Unbleached Natural",
        },
        {
            VariantTypeID: materialType.ID,
            Value:         "Organic Cotton - Bleached White",
        },
        {
            VariantTypeID: materialType.ID,
            Value:         "Nylon Mesh - Premium Durability",
        },
        {
            VariantTypeID: materialType.ID,
            Value:         "Bamboo Cotton Blend - Eco-Luxe",
        },
    }
    
    for i := range materialValues {
        if err := db.Create(&materialValues[i]).Error; err != nil {
            return err
        }
    }

    // Color Patterns - bright colors recommended to prevent loss in laundry [citation:3]
    colorValues := []entity.VariantValue{
        {
            VariantTypeID: colorType.ID,
            Value:         "Rainbow Mix (Multicolor)",
        },
        {
            VariantTypeID: colorType.ID,
            Value:         "Sunset Gradient (Orange/Pink/Yellow)",
        },
        {
            VariantTypeID: colorType.ID,
            Value:         "Ocean Blues (Navy/Teal/Aqua)",
        },
        {
            VariantTypeID: colorType.ID,
            Value:         "Forest Greens (Sage/Emerald/Mint)",
        },
        {
            VariantTypeID: colorType.ID,
            Value:         "Earth Tones (Terracotta/Ochre/Brown)",
        },
        {
            VariantTypeID: colorType.ID,
            Value:         "Batik Pattern - Kawung",
        },
        {
            VariantTypeID: colorType.ID,
            Value:         "Batik Pattern - Parang",
        },
        {
            VariantTypeID: colorType.ID,
            Value:         "Pastel Rainbow (Soft Pink/Lavender/Mint)",
        },
    }
    
    for i := range colorValues {
        if err := db.Create(&colorValues[i]).Error; err != nil {
            return err
        }
    }

    // Set Size Types
    setSizeValues := []entity.VariantValue{
        {
            VariantTypeID: setSizeType.ID,
            Value:         "Starter Set (4 bags) - 2 Small, 2 Medium",
        },
        {
            VariantTypeID: setSizeType.ID,
            Value:         "Complete Set (8 bags) - 2XS, 2S, 2M, 2L",
        },
        {
            VariantTypeID: setSizeType.ID,
            Value:         "Family Set (12 bags) - 3XS, 3S, 3M, 3L",
        },
        {
            VariantTypeID: setSizeType.ID,
            Value:         "Bulk Shopper Set (6 Large bags)",
        },
        {
            VariantTypeID: setSizeType.ID,
            Value:         "Gift Set (8 bags + Wooden Tags + Recipe Book)",
        },
    }
    
    for i := range setSizeValues {
        if err := db.Create(&setSizeValues[i]).Error; err != nil {
            return err
        }
    }

    // Base price mapping for different set types (in IDR) - based on Indonesian e-commerce pricing
    basePriceMap := map[string]float64{
        "Starter Set (4 bags) - 2 Small, 2 Medium":                    89000,
        "Complete Set (8 bags) - 2XS, 2S, 2M, 2L":                     159000,
        "Family Set (12 bags) - 3XS, 3S, 3M, 3L":                      219000,
        "Bulk Shopper Set (6 Large bags)":                             149000,
        "Gift Set (8 bags + Wooden Tags + Recipe Book)":               199000,
    }

    // Material premium adjustments
    materialPremium := map[string]float64{
        "Recycled PET (rPET) - From Plastic Bottles":   0,      // Standard
        "Organic Cotton - Unbleached Natural":         20000,   // Premium natural material
        "Organic Cotton - Bleached White":             25000,   // Bleached cotton
        "Nylon Mesh - Premium Durability":             15000,   // Durable synthetic
        "Bamboo Cotton Blend - Eco-Luxe":              35000,   // Premium eco material
    }

    // Color premium (batik patterns are premium)
    colorPremium := map[string]float64{
        "Rainbow Mix (Multicolor)":                     0,
        "Sunset Gradient (Orange/Pink/Yellow)":         0,
        "Ocean Blues (Navy/Teal/Aqua)":                 0,
        "Forest Greens (Sage/Emerald/Mint)":            0,
        "Earth Tones (Terracotta/Ochre/Brown)":         0,
        "Batik Pattern - Kawung":                       30000,
        "Batik Pattern - Parang":                        30000,
        "Pastel Rainbow (Soft Pink/Lavender/Mint)":     5000,
    }

    var allSKUs []entity.SKU
    var minPrice float64 = 999999
    var maxPrice float64 = 0

    // Generate SKUs for combinations
    for _, setSize := range setSizeValues {
        basePrice := basePriceMap[setSize.Value]
        
        for _, material := range materialValues {
            materialAdj := materialPremium[material.Value]
            
            for _, color := range colorValues {
                // Skip incompatible combinations
                // Batik patterns only available with cotton materials
                if (color.Value == "Batik Pattern - Kawung" || color.Value == "Batik Pattern - Parang") && 
                   (material.Value != "Organic Cotton - Unbleached Natural" && material.Value != "Organic Cotton - Bleached White") {
                    continue
                }
                
                // Nylon mesh only comes in solid colors (no gradients/patterns)
                if material.Value == "Nylon Mesh - Premium Durability" && 
                   (color.Value == "Sunset Gradient (Orange/Pink/Yellow)" || 
                    color.Value == "Pastel Rainbow (Soft Pink/Lavender/Mint)" ||
                    color.Value == "Batik Pattern - Kawung" || 
                    color.Value == "Batik Pattern - Parang") {
                    continue
                }
                
                // Bamboo blend only in earth tones and natural
                if material.Value == "Bamboo Cotton Blend - Eco-Luxe" && 
                   (color.Value == "Rainbow Mix (Multicolor)" || 
                    color.Value == "Sunset Gradient (Orange/Pink/Yellow)") {
                    continue
                }
                
                // Calculate final price
                colorAdj := colorPremium[color.Value]
                finalPrice := basePrice + materialAdj + colorAdj
                
                // Round to nearest 1000 Rupiah (common in Indonesian e-commerce)
                finalPrice = float64(int(finalPrice/1000)) * 1000
                
                // Generate SKU code
                skuCode := fmt.Sprintf("PRB-%s-%s-%s-%d", 
                    abbreviateBagSetSize(setSize.Value),
                    abbreviateBagMaterial(material.Value),
                    abbreviateBagColor(color.Value),
                    rand.Intn(1000))
                
                // Calculate sale price (30% chance of being on sale - common e-commerce promotion)
                var salePrice *float64
                if rand.Float64() < 0.3 {
                    discount := 0.80 + (rand.Float64() * 0.15) // 5-20% off
                    sp := finalPrice * discount
                    sp = float64(int(sp/1000)) * 1000
                    salePrice = &sp
                }
                
                // Determine stock levels based on popularity
                stock := 0
                if setSize.Value == "Complete Set (8 bags) - 2XS, 2S, 2M, 2L" {
                    stock = 300 + rand.Intn(400) // Most popular: 300-700
                } else if setSize.Value == "Starter Set (4 bags) - 2 Small, 2 Medium" {
                    stock = 200 + rand.Intn(300) // 200-500
                } else if setSize.Value == "Family Set (12 bags) - 3XS, 3S, 3M, 3L" {
                    stock = 100 + rand.Intn(200) // 100-300
                } else if setSize.Value == "Gift Set (8 bags + Wooden Tags + Recipe Book)" {
                    stock = 50 + rand.Intn(100)  // 50-150
                } else {
                    stock = 75 + rand.Intn(150)  // 75-225
                }
                
                // Adjust stock for premium variants
                if colorAdj > 0 || materialAdj > 20000 {
                    stock = stock / 2 // Half stock for premium variants
                }
                
                // Calculate weight (varies by set size)
                weight := 0.0
                switch setSize.Value {
                case "Starter Set (4 bags) - 2 Small, 2 Medium":
                    weight = 120
                case "Complete Set (8 bags) - 2XS, 2S, 2M, 2L":
                    weight = 250
                case "Family Set (12 bags) - 3XS, 3S, 3M, 3L":
                    weight = 380
                case "Bulk Shopper Set (6 Large bags)":
                    weight = 300
                case "Gift Set (8 bags + Wooden Tags + Recipe Book)":
                    weight = 350
                }
                
                sku := entity.SKU{
                    ProductID: product.ID,
                    SKUCode:   skuCode,
                    Price:     finalPrice,
                    SalePrice: salePrice,
                    Stock:     stock,
                    MinStock:  20,
                    Status:    entity.SKUStatusActive,
                    Weight:    weight,
                }
                
                if err := db.Create(&sku).Error; err != nil {
                    return err
                }
                
                // Track min/max prices
                if finalPrice < minPrice {
                    minPrice = finalPrice
                }
                if finalPrice > maxPrice {
                    maxPrice = finalPrice
                }
                
                // Link SKU to variant values
                skuVariant1 := entity.SKUVariantValue{
                    SKUID:          sku.ID,
                    VariantValueID: setSize.ID,
                }
                db.Create(&skuVariant1)
                
                skuVariant2 := entity.SKUVariantValue{
                    SKUID:          sku.ID,
                    VariantValueID: material.ID,
                }
                db.Create(&skuVariant2)
                
                skuVariant3 := entity.SKUVariantValue{
                    SKUID:          sku.ID,
                    VariantValueID: color.ID,
                }
                db.Create(&skuVariant3)
                
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
			"https://m.media-amazon.com/images/I/51NRtt4kMnL._SY300_SX300_QL70_FMwebp_.jpg",
			"https://m.media-amazon.com/images/I/51jGHGjAz0L._SY300_SX300_QL70_FMwebp_.jpg",
			"https://m.media-amazon.com/images/I/81uw8HQownL._SX569_.jpg",
		}

		for _, url := range imageURLs {
			img := entity.Image{
				ProductID: product.ID,
				ImageURL: url,
			}

			db.Create(&img)
		}

    fmt.Printf("✅ Successfully seeded Reusable Produce Bags product with %d SKU variants\n", len(allSKUs))
    fmt.Printf("   Price range: Rp %.0f - Rp %.0f\n", minPrice, maxPrice)
    
    return nil
}

// Helper functions for SKU code generation
func abbreviateBagSetSize(setSize string) string {
    switch {
    case setSize == "Starter Set (4 bags) - 2 Small, 2 Medium":
        return "STR"
    case setSize == "Complete Set (8 bags) - 2XS, 2S, 2M, 2L":
        return "CMP"
    case setSize == "Family Set (12 bags) - 3XS, 3S, 3M, 3L":
        return "FAM"
    case setSize == "Bulk Shopper Set (6 Large bags)":
        return "BLK"
    case setSize == "Gift Set (8 bags + Wooden Tags + Recipe Book)":
        return "GFT"
    default:
        return "SET"
    }
}

func abbreviateBagMaterial(material string) string {
    switch {
    case material == "Recycled PET (rPET) - From Plastic Bottles":
        return "RPT"
    case material == "Organic Cotton - Unbleached Natural":
        return "CTN"
    case material == "Organic Cotton - Bleached White":
        return "CTW"
    case material == "Nylon Mesh - Premium Durability":
        return "NYL"
    case material == "Bamboo Cotton Blend - Eco-Luxe":
        return "BMB"
    default:
        return "MAT"
    }
}

func abbreviateBagColor(color string) string {
    switch {
    case color == "Rainbow Mix (Multicolor)":
        return "RNB"
    case color == "Sunset Gradient (Orange/Pink/Yellow)":
        return "SUN"
    case color == "Ocean Blues (Navy/Teal/Aqua)":
        return "OCN"
    case color == "Forest Greens (Sage/Emerald/Mint)":
        return "FOR"
    case color == "Earth Tones (Terracotta/Ochre/Brown)":
        return "ERT"
    case color == "Batik Pattern - Kawung":
        return "BTK1"
    case color == "Batik Pattern - Parang":
        return "BTK2"
    case color == "Pastel Rainbow (Soft Pink/Lavender/Mint)":
        return "PST"
    default:
        return "CLR"
    }
}