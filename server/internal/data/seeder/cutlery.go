package seeder

import (
	"debian-ecommerce/internal/data/entity"
	"fmt"
	"math/rand"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func SeedBambooCutlerySet(db *gorm.DB) error {
    rand.Seed(time.Now().UnixNano())
    
    // First, ensure category exists (assuming category ID 6 is Straw & Cutlery)
    // If not, create it first
    var category entity.Category
    result := db.First(&category, 6)
    if result.Error != nil {
        // Create Straw & Cutlery category if it doesn't exist
        category = entity.Category{
            Model:       gorm.Model{ID: 6},
            Name:        "Straw & Cutlery",
        }
        if err := db.Create(&category).Error; err != nil {
            return err
        }
    }

    // Create the main product - Bamboo Travel Cutlery Set
    product := entity.Product{
        CategoryID:       6,
        Name:             "EcoBamboo Travel Cutlery Set - 7 Piece Complete Set",
        Description:       `Say goodbye to single-use plastic with our premium 7-piece bamboo cutlery set. Perfect for travel, office lunches, camping, and picnics. Made from 100% organic Moso bamboo, this complete set includes everything you need for sustainable dining on-the-go. The bamboo is naturally antibacterial, lightweight, and durable. Each piece is finely sanded for a smooth finish and coated with food-grade mineral oil. Comes in a stylish, washable cotton pouch with carabiner for easy attachment to your bag. Join the zero-waste movement and make every meal a sustainable one! 🌱

**Set Includes:**
- 1x Bamboo Fork (7.9 inches / 20cm)
- 1x Bamboo Knife (7.9 inches / 20cm) - Serrated edge for cutting
- 1x Bamboo Spoon (7.9 inches / 20cm)
- 1x Bamboo Chopsticks (8.3 inches / 21cm)
- 1x Bamboo Straw (7.9 inches / 20cm)
- 1x Stainless Steel Straw Cleaner Brush
- 1x Organic Cotton Travel Pouch with Carabiner

**Why Choose Bamboo?**
- 🌿 Biodegradable and compostable at end of life
- 🌿 Grows 30x faster than trees, regenerates after harvesting
- 🌿 Naturally antibacterial and odor-resistant
- 🌿 Renewable resource - FSC certified
- 🌿 Saves 1000+ plastic utensils from landfills over its lifetime

**Care Instructions:**
- Hand wash with mild soap and dry thoroughly
- Not dishwasher safe (to preserve natural oils)
- Periodically oil with food-grade mineral oil to maintain luster
- Avoid prolonged soaking`,
        IsPublished:      true,
        Tags:             pq.StringArray{"bamboo", "cutlery", "travel set", "reusable", "eco-friendly", "zero waste", "camping", "picnic", "sustainable", "plastic-free", "fork spoon knife", "chopsticks", "bamboo straw"},
        Slug:             "ecobamboo-travel-cutlery-set-7-piece-complete",
        MainImageURL:     "https://down-id.img.susercontent.com/file/id-11134207-7r990-lzldce687dwoe9_tn",
        MainImagePublicID: "bamboo-cutlery-main",
        AverageRating:    4.7,
        ReviewCount:      892,
        SoldCount:        3456,
        MinPrice:         0, // Will be updated after SKUs
        MaxPrice:         0, // Will be updated after SKUs
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }

    // Create Variant Types
    // Variant 1: Case Color / Pattern
    caseType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Case Color",
    }
    if err := db.Create(&caseType).Error; err != nil {
        return err
    }

    // Variant 2: Set Composition (some people want different combinations)
    compositionType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Set Type",
    }
    if err := db.Create(&compositionType).Error; err != nil {
        return err
    }

    // Variant 3: Finish Type
    finishType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Finish",
    }
    if err := db.Create(&finishType).Error; err != nil {
        return err
    }

    // Create Variant Values
    // Case Colors
    caseValues := []entity.VariantValue{
        {
            VariantTypeID: caseType.ID,
            Value:         "Natural Cotton - Beige",
        },
        {
            VariantTypeID: caseType.ID,
            Value:         "Sage Green",
        },
        {
            VariantTypeID: caseType.ID,
            Value:         "Terracotta Orange",
        },
        {
            VariantTypeID: caseType.ID,
            Value:         "Navy Blue",
        },
        {
            VariantTypeID: caseType.ID,
            Value:         "Black",
        },
        {
            VariantTypeID: caseType.ID,
            Value:         "Batik Pattern - Cendrawasih",
        },
        {
            VariantTypeID: caseType.ID,
            Value:         "Batik Pattern - Mega Mendung",
        },
        {
            VariantTypeID: caseType.ID,
            Value:         "Rainbow Stripe",
        },
    }
    
    for i := range caseValues {
        if err := db.Create(&caseValues[i]).Error; err != nil {
            return err
        }
    }

    // Set Composition Types
    compositionValues := []entity.VariantValue{
        {
            VariantTypeID: compositionType.ID,
            Value:         "Complete Set (7 pieces) - Fork, Knife, Spoon, Chopsticks, Straw, Brush, Pouch",
        },
        {
            VariantTypeID: compositionType.ID,
            Value:         "Essential Set (5 pieces) - Fork, Knife, Spoon, Pouch",
        },
        {
            VariantTypeID: compositionType.ID,
            Value:         "Asian Set (6 pieces) - Chopsticks, Spoon, Fork, Straw, Brush, Pouch",
        },
        {
            VariantTypeID: compositionType.ID,
            Value:         "Straw-Free Set (6 pieces) - Fork, Knife, Spoon, Chopsticks, Brush, Pouch",
        },
        {
            VariantTypeID: compositionType.ID,
            Value:         "Family Pack (2 Complete Sets) - 2x of everything",
        },
    }
    
    for i := range compositionValues {
        if err := db.Create(&compositionValues[i]).Error; err != nil {
            return err
        }
    }

    // Finish Types
    finishValues := []entity.VariantValue{
        {
            VariantTypeID: finishType.ID,
            Value:         "Natural - Unfinished (Eco Choice)",
        },
        {
            VariantTypeID: finishType.ID,
            Value:         "Polished - Food Grade Mineral Oil",
        },
        {
            VariantTypeID: finishType.ID,
            Value:         "Carbonized - Dark Walnut Shade",
        },
        {
            VariantTypeID: finishType.ID,
            Value:         "Engraved - Custom Name Available",
        },
    }
    
    for i := range finishValues {
        if err := db.Create(&finishValues[i]).Error; err != nil {
            return err
        }
    }

    // Base price mapping for different set types (in IDR)
    basePriceMap := map[string]float64{
        "Complete Set (7 pieces) - Fork, Knife, Spoon, Chopsticks, Straw, Brush, Pouch":      129000,
        "Essential Set (5 pieces) - Fork, Knife, Spoon, Pouch":                                89000,
        "Asian Set (6 pieces) - Chopsticks, Spoon, Fork, Straw, Brush, Pouch":                 115000,
        "Straw-Free Set (6 pieces) - Fork, Knife, Spoon, Chopsticks, Brush, Pouch":            119000,
        "Family Pack (2 Complete Sets) - 2x of everything":                                    219000,
    }

    // Finish premium adjustments
    finishPremium := map[string]float64{
        "Natural - Unfinished (Eco Choice)":                0,
        "Polished - Food Grade Mineral Oil":                15000,
        "Carbonized - Dark Walnut Shade":                   25000,
        "Engraved - Custom Name Available":                 35000,
    }

    // Case color premium (batik patterns are premium, others standard)
    casePremium := map[string]float64{
        "Natural Cotton - Beige":       0,
        "Sage Green":                    0,
        "Terracotta Orange":              0,
        "Navy Blue":                      0,
        "Black":                          0,
        "Batik Pattern - Cendrawasih":    25000,
        "Batik Pattern - Mega Mendung":   25000,
        "Rainbow Stripe":                  5000,
    }

    var allSKUs []entity.SKU
    var minPrice float64 = 999999
    var maxPrice float64 = 0

    // Generate SKUs for combinations
    for _, composition := range compositionValues {
        basePrice := basePriceMap[composition.Value]
        
        for _, finish := range finishValues {
            // Skip some combinations (e.g., Family Pack doesn't come with engraving due to complexity)
            if composition.Value == "Family Pack (2 Complete Sets) - 2x of everything" && 
               finish.Value == "Engraved - Custom Name Available" {
                continue
            }
            
            // Skip unfinished for premium patterns (they need protection)
            for _, caseColor := range caseValues {
                // Skip natural finish for batik cases (they need sealing)
                if finish.Value == "Natural - Unfinished (Eco Choice)" && 
                   (caseColor.Value == "Batik Pattern - Cendrawasih" || 
                    caseColor.Value == "Batik Pattern - Mega Mendung") {
                    continue
                }
                
                // Calculate final price
                caseAdjustment := casePremium[caseColor.Value]
                finishAdjustment := finishPremium[finish.Value]
                
                finalPrice := basePrice + caseAdjustment + finishAdjustment
                
                // Volume discount for family pack already factored in base price
                
                // Round to nearest 500 Rupiah (common in Indonesia)
                finalPrice = float64(int(finalPrice/500)) * 500
                
                // Generate SKU code
                skuCode := fmt.Sprintf("BCT-%s-%s-%s-%d", 
                    abbreviateComposition(composition.Value),
                    abbreviateCase(caseColor.Value),
                    abbreviateFinish(finish.Value),
                    rand.Intn(1000))
                
                // Calculate sale price (30% chance of being on sale)
                var salePrice *float64
                if rand.Float64() < 0.3 {
                    discount := 0.80 + (rand.Float64() * 0.15) // 5-20% off
                    sp := finalPrice * discount
                    sp = float64(int(sp/500)) * 500
                    salePrice = &sp
                }
                
                // Determine stock levels
                stock := 0
                if composition.Value == "Complete Set (7 pieces) - Fork, Knife, Spoon, Chopsticks, Straw, Brush, Pouch" {
                    stock = 200 + rand.Intn(300) // Most popular: 200-500
                } else if composition.Value == "Essential Set (5 pieces) - Fork, Knife, Spoon, Pouch" {
                    stock = 150 + rand.Intn(200) // 150-350
                } else if composition.Value == "Family Pack (2 Complete Sets) - 2x of everything" {
                    stock = 50 + rand.Intn(100)  // 50-150 (bulk)
                } else {
                    stock = 100 + rand.Intn(150)  // 100-250
                }
                
                // Adjust stock for premium variants
                if caseAdjustment > 0 || finishAdjustment > 15000 {
                    stock = stock / 2 // Half stock for premium variants
                }
                
                // Calculate weight (varies by set type)
                weight := 0.0
                switch composition.Value {
                case "Complete Set (7 pieces) - Fork, Knife, Spoon, Chopsticks, Straw, Brush, Pouch":
                    weight = 250
                case "Essential Set (5 pieces) - Fork, Knife, Spoon, Pouch":
                    weight = 180
                case "Asian Set (6 pieces) - Chopsticks, Spoon, Fork, Straw, Brush, Pouch":
                    weight = 220
                case "Straw-Free Set (6 pieces) - Fork, Knife, Spoon, Chopsticks, Brush, Pouch":
                    weight = 230
                case "Family Pack (2 Complete Sets) - 2x of everything":
                    weight = 500
                }
                
                sku := entity.SKU{
                    ProductID: product.ID,
                    SKUCode:   skuCode,
                    Price:     finalPrice,
                    SalePrice: salePrice,
                    Stock:     stock,
                    MinStock:  15,
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
                    VariantValueID: composition.ID,
                }
                db.Create(&skuVariant1)
                
                skuVariant2 := entity.SKUVariantValue{
                    SKUID:          sku.ID,
                    VariantValueID: finish.ID,
                }
                db.Create(&skuVariant2)
                
                skuVariant3 := entity.SKUVariantValue{
                    SKUID:          sku.ID,
                    VariantValueID: caseColor.ID,
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
			"https://down-id.img.susercontent.com/file/id-11134207-7r98s-lzldce68a6o515@resize_w900_nl.webp",
			"https://m.media-amazon.com/images/I/81JPUvZ3AhS._AC_UF894,1000_QL80_.jpg",
			"https://m.media-amazon.com/images/I/81uw8HQownL._SX569_.jpghttps://biopac.id/wp-content/uploads/2023/10/2.png",
		}

		for _, url := range imageURLs {
			img := entity.Image{
				ProductID: product.ID,
				ImageURL: url,
			}

			db.Create(&img)
		}

    fmt.Printf("✅ Successfully seeded Bamboo Travel Cutlery Set product with %d SKU variants\n", len(allSKUs))
    fmt.Printf("   Price range: Rp %.0f - Rp %.0f\n", minPrice, maxPrice)
    
    return nil
}

// Helper functions for SKU code generation
func abbreviateComposition(comp string) string {
    switch {
    case comp == "Complete Set (7 pieces) - Fork, Knife, Spoon, Chopsticks, Straw, Brush, Pouch":
        return "CMP"
    case comp == "Essential Set (5 pieces) - Fork, Knife, Spoon, Pouch":
        return "ESS"
    case comp == "Asian Set (6 pieces) - Chopsticks, Spoon, Fork, Straw, Brush, Pouch":
        return "ASN"
    case comp == "Straw-Free Set (6 pieces) - Fork, Knife, Spoon, Chopsticks, Brush, Pouch":
        return "NOB"
    case comp == "Family Pack (2 Complete Sets) - 2x of everything":
        return "FAM"
    default:
        return "STD"
    }
}

func abbreviateCase(caseColor string) string {
    switch {
    case caseColor == "Natural Cotton - Beige":
        return "NAT"
    case caseColor == "Sage Green":
        return "SGE"
    case caseColor == "Terracotta Orange":
        return "TER"
    case caseColor == "Navy Blue":
        return "NVY"
    case caseColor == "Black":
        return "BLK"
    case caseColor == "Batik Pattern - Cendrawasih":
        return "BAT1"
    case caseColor == "Batik Pattern - Mega Mendung":
        return "BAT2"
    case caseColor == "Rainbow Stripe":
        return "RNB"
    default:
        return "CLR"
    }
}

func abbreviateFinish(finish string) string {
    switch {
    case finish == "Natural - Unfinished (Eco Choice)":
        return "RAW"
    case finish == "Polished - Food Grade Mineral Oil":
        return "POL"
    case finish == "Carbonized - Dark Walnut Shade":
        return "CAR"
    case finish == "Engraved - Custom Name Available":
        return "ENG"
    default:
        return "FIN"
    }
}