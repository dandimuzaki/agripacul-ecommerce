package seeder

import (
	"debian-ecommerce/internal/data/entity"
	"fmt"
	"math/rand"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func SeedBuahLerak(db *gorm.DB) error {
    rand.Seed(time.Now().UnixNano())
    
    // First, ensure category exists (assuming category ID 3 is Home & Cleaning)
    // If not, create it first
    var category entity.Category
    result := db.First(&category, 3)
    if result.Error != nil {
        // Create Home & Cleaning category if it doesn't exist
        category = entity.Category{
            Model:       gorm.Model{ID: 3},
            Name:        "Home & Cleaning",
        }
        if err := db.Create(&category).Error; err != nil {
            return err
        }
    }

    // Create the main product - Buah Lerak (Soap Nut)
    product := entity.Product{
        CategoryID:       3,
        Name:             "Buah Lerak Organic - Natural Soap Nut for Laundry & Batik Care",
        Description:      `Buah Lerak (Sapindus rarak DC.) is a traditional Indonesian natural detergent that has been used for centuries, long before chemical detergents entered the market. This remarkable fruit contains **saponins**—natural surfactants that produce gentle foam and effectively clean clothes without harsh chemicals [citation:4].

**🌿 Why Choose Buah Lerak?**

The main active compound in lerak is saponin, which functions as a natural surfactant with foaming and emulsifying properties. Unlike chemical detergents containing Sodium Lauryl Sulfate (SLS), lerak is completely biodegradable and safe for aquatic ecosystems [citation:4]. For batik enthusiasts, lerak is particularly prized—it cleans batik fabric thoroughly without causing colors to fade, preserving the beauty of your cherished textiles [citation:4].

**🧺 Multiple Uses:**
- **Laundry Detergent**: Perfect for all fabrics, especially delicate batik
- **Dishwashing Liquid**: Cuts grease naturally
- **Floor Cleaner**: Add to mop water for streak-free cleaning
- **Hand Soap**: Gentle on skin, no chemical irritation
- **Shampoo**: Traditional hair cleanser
- **Biopesticide**: Natural plant protection [citation:1]

**✨ Key Benefits:**
- ✅ **100% Natural**: No chemicals, preservatives, or synthetic fragrances
- ✅ **Hypoallergenic**: Safe for sensitive skin and babies—no allergies or irritation [citation:1]
- ✅ **Fabric Safe**: Does not damage clothes or cause fading, ideal for batik [citation:1]
- ✅ **Odorless**: No perfume smell—clothes smell clean, not artificially scented
- ✅ **Easy Rinse**: Requires less water compared to chemical detergents [citation:1]
- ✅ **Compostable**: Used fruits can be composted—zero waste [citation:1]
- ✅ **Economical**: Each fruit can be used 6-8 times [citation:1]

**🌱 Sustainability Impact:**
- **Biodegradable**: Breaks down naturally without polluting waterways
- **Plastic-free packaging**: Packed in jute bags or recycled paper
- **Traditional wisdom**: Supports local farmers and preserves Indonesian heritage
- **Water protection**: No phosphates or surfactants that cause eutrophication

**📋 How to Use:**

*Method 1 - Without Boiling (Gentle):*
1. Soak 7-10 lerak fruits in water for 2 nights until soft
2. Remove seeds, mash the flesh to release saponin sap
3. Place in cloth bag and add to washing machine [citation:5]

*Method 2 - With Boiling (Concentrated):*
1. Soak 20 lerak fruits in 1-2L water for 2 nights
2. Remove seeds, boil flesh with citrus peels/lemongrass (optional) for 60 minutes
3. Add 2 tbsp salt per liter as natural preservative
4. Strain and use liquid as detergent [citation:5]

*Each fruit can be reused 1-6 times until saponin is depleted* [citation:5]

**🌳 About the Plant:**
Lerak comes from the Sapindus rarak tree, native to Indonesia (especially Java and Sumatra), India, Sri Lanka, and South China. These deciduous trees grow 20-50 feet tall and take 10-15 years before first harvest—making each fruit precious! [citation:1][citation:6]`,
        IsPublished:      true,
        Tags:             pq.StringArray{"buah lerak", "soap nut", "natural detergent", "lerak", "batik care", "eco-friendly", "zero waste", "traditional", "saponin", "chemical-free", "hypoallergenic", "baby clothes", "home cleaning", "sustainable", "biodegradable"},
        Slug:             "buah-lerak-organic-natural-soap-nut-laundry-batik-care",
        MainImageURL:     "https://down-id.img.susercontent.com/file/sg-11134201-22110-591szrfpjckva1_tn",
        MainImagePublicID: "buah-lerak-main",
        AverageRating:    4.9,
        ReviewCount:      876,
        SoldCount:        3456,
        MinPrice:         0, // Will be updated after SKUs
        MaxPrice:         0, // Will be updated after SKUs
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }

    // Create Variant Types
    // Variant 1: Form (whole dried fruits or liquid extract)
    formType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Product Form",
    }
    if err := db.Create(&formType).Error; err != nil {
        return err
    }

    // Variant 2: Size/Packaging
    sizeType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Pack Size",
    }
    if err := db.Create(&sizeType).Error; err != nil {
        return err
    }

    // Variant 3: Packaging Type
    packagingType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Packaging Material",
    }
    if err := db.Create(&packagingType).Error; err != nil {
        return err
    }

    // Variant 4: Processing Method
    processingType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Processing",
    }
    if err := db.Create(&processingType).Error; err != nil {
        return err
    }

    // Create Variant Values
    // Product Form
    formValues := []entity.VariantValue{
        {
            VariantTypeID: formType.ID,
            Value:         "Whole Dried Fruits - Traditional",
        },
        {
            VariantTypeID: formType.ID,
            Value:         "Powder - Ground Lerak",
        },
        {
            VariantTypeID: formType.ID,
            Value:         "Liquid Concentrate - Ready to Use",
        },
        {
            VariantTypeID: formType.ID,
            Value:         "Paste - Semi-Solid Concentrate",
        },
    }
    
    for i := range formValues {
        if err := db.Create(&formValues[i]).Error; err != nil {
            return err
        }
    }

    // Pack Size
    sizeValues := []entity.VariantValue{
        {
            VariantTypeID: sizeType.ID,
            Value:         "100g - Trial Size",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "250g - Small Pack",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "500g - Medium Pack",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "1kg - Family Pack",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "5kg - Bulk Economy",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "500ml Liquid - Standard",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "1L Liquid - Economy",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "5L Liquid - Refill/Jerry Can",
        },
    }
    
    for i := range sizeValues {
        if err := db.Create(&sizeValues[i]).Error; err != nil {
            return err
        }
    }

    // Packaging Material (eco-friendly options)
    packagingValues := []entity.VariantValue{
        {
            VariantTypeID: packagingType.ID,
            Value:         "Jute Bag (Karang Goni) - Zero Plastic",
        },
        {
            VariantTypeID: packagingType.ID,
            Value:         "Kraft Paper Bag - Recyclable",
        },
        {
            VariantTypeID: packagingType.ID,
            Value:         "Kantong Blacu (Cotton Sack) - Traditional",
        },
        {
            VariantTypeID: packagingType.ID,
            Value:         "Glass Jar - Reusable",
        },
        {
            VariantTypeID: packagingType.ID,
            Value:         "Recycled Plastic Pouch - Minimalist",
        },
        {
            VariantTypeID: packagingType.ID,
            Value:         "Jerry Can Refill - Bring Your Own",
        },
        {
            VariantTypeID: packagingType.ID,
            Value:         "Compostable PLA Pouch",
        },
    }
    
    for i := range packagingValues {
        if err := db.Create(&packagingValues[i]).Error; err != nil {
            return err
        }
    }

    // Processing Method
    processingValues := []entity.VariantValue{
        {
            VariantTypeID: processingType.ID,
            Value:         "Sun-Dried Traditional",
        },
        {
            VariantTypeID: processingType.ID,
            Value:         "Oven-Dried Low Temperature",
        },
        {
            VariantTypeID: processingType.ID,
            Value:         "Cold-Processed Liquid",
        },
        {
            VariantTypeID: processingType.ID,
            Value:         "Boiled Concentrate",
        },
        {
            VariantTypeID: processingType.ID,
            Value:         "Fermented - Enhanced Saponin",
        },
    }
    
    for i := range processingValues {
        if err := db.Create(&processingValues[i]).Error; err != nil {
            return err
        }
    }

    // Base price mapping (in IDR) - based on Tokopedia data [citation:1][citation:5]
    // Whole Dried Fruits: ~Rp 9,500-12,000 per 100g [citation:5]
    // 250g pack: ~Rp 30,000 [citation:1]
    // 1kg pack: ~Rp 90,000-110,000
    
    basePriceMap := map[string]map[string]float64{
        "Whole Dried Fruits - Traditional": {
            "100g - Trial Size":      12500,
            "250g - Small Pack":      30000,
            "500g - Medium Pack":     55000,
            "1kg - Family Pack":      95000,
            "5kg - Bulk Economy":     425000,
        },
        "Powder - Ground Lerak": {
            "100g - Trial Size":      18000,
            "250g - Small Pack":      42000,
            "500g - Medium Pack":     75000,
            "1kg - Family Pack":      135000,
            "5kg - Bulk Economy":     600000,
        },
        "Liquid Concentrate - Ready to Use": {
            "500ml Liquid - Standard": 45000,
            "1L Liquid - Economy":     75000,
            "5L Liquid - Refill/Jerry Can": 325000,
        },
        "Paste - Semi-Solid Concentrate": {
            "250g - Small Pack":      38000,
            "500g - Medium Pack":     68000,
            "1kg - Family Pack":      125000,
        },
    }

    // Packaging premium adjustments
    packagingPremium := map[string]float64{
        "Jute Bag (Karang Goni) - Zero Plastic":      0,
        "Kraft Paper Bag - Recyclable":               0,
        "Kantong Blacu (Cotton Sack) - Traditional":  2000, // Additional cost for fabric sack [citation:5]
        "Glass Jar - Reusable":                        15000, // Premium glass packaging
        "Recycled Plastic Pouch - Minimalist":        -2000, // Discount for minimal packaging
        "Jerry Can Refill - Bring Your Own":          -5000, // Discount for BYO container
        "Compostable PLA Pouch":                       5000, // Premium eco material
    }

    // Processing premium adjustments
    processingPremium := map[string]float64{
        "Sun-Dried Traditional":          0,
        "Oven-Dried Low Temperature":     5000,   // Controlled process
        "Cold-Processed Liquid":          8000,   // Preserves nutrients
        "Boiled Concentrate":              0,      // Standard method
        "Fermented - Enhanced Saponin":   15000,  // Premium processing
    }

    var allSKUs []entity.SKU
    var minPrice float64 = 999999
    var maxPrice float64 = 0

    // Generate SKUs for combinations
    for _, form := range formValues {
        formPrices, exists := basePriceMap[form.Value]
        if !exists {
            continue // Skip forms without price mapping
        }
        
        for _, size := range sizeValues {
            // Check if this size exists for this form
            basePrice, sizeExists := formPrices[size.Value]
            if !sizeExists {
                continue
            }
            
            for _, packaging := range packagingValues {
                // Skip incompatible packaging-size combinations
                // Glass jars only for smaller sizes
                if packaging.Value == "Glass Jar - Reusable" && 
                   (size.Value == "5kg - Bulk Economy" || size.Value == "5L Liquid - Refill/Jerry Can") {
                    continue
                }
                
                // Jerry can only for liquid sizes
                if packaging.Value == "Jerry Can Refill - Bring Your Own" && 
                   !(size.Value == "1L Liquid - Economy" || size.Value == "5L Liquid - Refill/Jerry Can") {
                    continue
                }
                
                // Blacu sack traditionally for dried fruits only
                if packaging.Value == "Kantong Blacu (Cotton Sack) - Traditional" && 
                   form.Value != "Whole Dried Fruits - Traditional" {
                    continue
                }
                
                for _, processing := range processingValues {
                    // Skip incompatible processing-form combinations
                    // Liquid processing only for liquid forms
                    if (processing.Value == "Cold-Processed Liquid" || processing.Value == "Boiled Concentrate") && 
                       !(form.Value == "Liquid Concentrate - Ready to Use" || form.Value == "Paste - Semi-Solid Concentrate") {
                        continue
                    }
                    
                    // Fermented only for liquid
                    if processing.Value == "Fermented - Enhanced Saponin" && 
                       form.Value != "Liquid Concentrate - Ready to Use" {
                        continue
                    }
                    
                    // Calculate final price
                    packagingAdj := packagingPremium[packaging.Value]
                    processingAdj := processingPremium[processing.Value]
                    
                    finalPrice := basePrice + packagingAdj + processingAdj
                    
                    // Ensure price doesn't go negative with discounts
                    if finalPrice < 5000 {
                        finalPrice = 5000
                    }
                    
                    // Round to nearest 500 Rupiah (common in Indonesia)
                    finalPrice = float64(int(finalPrice/500)) * 500
                    
                    // Generate SKU code
                    skuCode := fmt.Sprintf("LER-%s-%s-%s-%s-%d", 
                        abbreviateForm(form.Value),
                        abbreviateLerakSize(size.Value),
                        abbreviateLerakPackaging(packaging.Value),
                        abbreviateProcessing(processing.Value),
                        rand.Intn(1000))
                    
                    // Calculate sale price (30% chance of being on sale)
                    var salePrice *float64
                    if rand.Float64() < 0.3 {
                        discount := 0.80 + (rand.Float64() * 0.15) // 5-20% off
                        sp := finalPrice * discount
                        sp = float64(int(sp/500)) * 500
                        salePrice = &sp
                    }
                    
                    // Determine stock levels based on popularity
                    stock := 0
                    if form.Value == "Whole Dried Fruits - Traditional" && size.Value == "250g - Small Pack" {
                        stock = 200 + rand.Intn(300) // Most popular: 200-500
                    } else if form.Value == "Whole Dried Fruits - Traditional" && size.Value == "1kg - Family Pack" {
                        stock = 100 + rand.Intn(200) // 100-300
                    } else if form.Value == "Liquid Concentrate - Ready to Use" && size.Value == "1L Liquid - Economy" {
                        stock = 80 + rand.Intn(150) // 80-230
                    } else if size.Value == "5kg - Bulk Economy" || size.Value == "5L Liquid - Refill/Jerry Can" {
                        stock = 20 + rand.Intn(50) // 20-70 (bulk)
                    } else {
                        stock = 50 + rand.Intn(150) // 50-200
                    }
                    
                    // Adjust stock for premium variants
                    if processingAdj > 10000 || packagingAdj > 5000 {
                        stock = stock / 2
                    }
                    
                    // Calculate weight (varies by form and size)
                    weight := 0.0
                    switch {
                    case form.Value == "Whole Dried Fruits - Traditional":
                        if size.Value == "100g - Trial Size" {
                            weight = 110
                        } else if size.Value == "250g - Small Pack" {
                            weight = 260
                        } else if size.Value == "500g - Medium Pack" {
                            weight = 510
                        } else if size.Value == "1kg - Family Pack" {
                            weight = 1020
                        } else if size.Value == "5kg - Bulk Economy" {
                            weight = 5050
                        }
                    case form.Value == "Powder - Ground Lerak":
                        if size.Value == "100g - Trial Size" {
                            weight = 120
                        } else if size.Value == "250g - Small Pack" {
                            weight = 270
                        } else if size.Value == "500g - Medium Pack" {
                            weight = 520
                        } else if size.Value == "1kg - Family Pack" {
                            weight = 1030
                        } else if size.Value == "5kg - Bulk Economy" {
                            weight = 5100
                        }
                    case form.Value == "Liquid Concentrate - Ready to Use":
                        if size.Value == "500ml Liquid - Standard" {
                            weight = 550
                        } else if size.Value == "1L Liquid - Economy" {
                            weight = 1050
                        } else if size.Value == "5L Liquid - Refill/Jerry Can" {
                            weight = 5200
                        }
                    case form.Value == "Paste - Semi-Solid Concentrate":
                        if size.Value == "250g - Small Pack" {
                            weight = 280
                        } else if size.Value == "500g - Medium Pack" {
                            weight = 530
                        } else if size.Value == "1kg - Family Pack" {
                            weight = 1040
                        }
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
                        VariantValueID: form.ID,
                    }
                    db.Create(&skuVariant1)
                    
                    skuVariant2 := entity.SKUVariantValue{
                        SKUID:          sku.ID,
                        VariantValueID: size.ID,
                    }
                    db.Create(&skuVariant2)
                    
                    skuVariant3 := entity.SKUVariantValue{
                        SKUID:          sku.ID,
                        VariantValueID: packaging.ID,
                    }
                    db.Create(&skuVariant3)
                    
                    skuVariant4 := entity.SKUVariantValue{
                        SKUID:          sku.ID,
                        VariantValueID: processing.ID,
                    }
                    db.Create(&skuVariant4)
                    
                    allSKUs = append(allSKUs, sku)
                }
            }
        }
    }

    // Update product with min and max prices
    product.MinPrice = minPrice
    product.MaxPrice = maxPrice
    db.Save(&product)

		// Upload images
		imageURLs := []string{
			"https://down-id.img.susercontent.com/file/id-11134207-7ras8-m1zivyovk1kt54.webp",
			"https://asset.kompas.com/crops/SBorOswHS4Kan2T_HemNCUTQpGk=/100x67:900x601/1200x800/data/photo/2023/02/24/63f8c83700d68.jpg",
			"https://benihbunbun.com/wp-content/uploads/2023/08/wp-1692842234781.jpg",
		}

		for _, url := range imageURLs {
			img := entity.Image{
				ProductID: product.ID,
				ImageURL: url,
			}

			db.Create(&img)
		}

    fmt.Printf("✅ Successfully seeded Buah Lerak (Soap Nut) product with %d SKU variants\n", len(allSKUs))
    fmt.Printf("   Price range: Rp %.0f - Rp %.0f\n", minPrice, maxPrice)
    
    return nil
}

// Helper functions for SKU code generation
func abbreviateForm(form string) string {
    switch {
    case form == "Whole Dried Fruits - Traditional":
        return "WDF"
    case form == "Powder - Ground Lerak":
        return "PWD"
    case form == "Liquid Concentrate - Ready to Use":
        return "LIQ"
    case form == "Paste - Semi-Solid Concentrate":
        return "PST"
    default:
        return "FRM"
    }
}

func abbreviateLerakSize(size string) string {
    switch {
    case size == "100g - Trial Size":
        return "100G"
    case size == "250g - Small Pack":
        return "250G"
    case size == "500g - Medium Pack":
        return "500G"
    case size == "1kg - Family Pack":
        return "1KG"
    case size == "5kg - Bulk Economy":
        return "5KG"
    case size == "500ml Liquid - Standard":
        return "500M"
    case size == "1L Liquid - Economy":
        return "1L"
    case size == "5L Liquid - Refill/Jerry Can":
        return "5L"
    default:
        return "SIZ"
    }
}

func abbreviateLerakPackaging(pkg string) string {
    switch {
    case pkg == "Jute Bag (Karang Goni) - Zero Plastic":
        return "JUT"
    case pkg == "Kraft Paper Bag - Recyclable":
        return "KRT"
    case pkg == "Kantong Blacu (Cotton Sack) - Traditional":
        return "BLC"
    case pkg == "Glass Jar - Reusable":
        return "GLZ"
    case pkg == "Recycled Plastic Pouch - Minimalist":
        return "RPP"
    case pkg == "Jerry Can Refill - Bring Your Own":
        return "JRY"
    case pkg == "Compostable PLA Pouch":
        return "PLA"
    default:
        return "PKG"
    }
}

func abbreviateProcessing(proc string) string {
    switch {
    case proc == "Sun-Dried Traditional":
        return "SUN"
    case proc == "Oven-Dried Low Temperature":
        return "OVN"
    case proc == "Cold-Processed Liquid":
        return "CLD"
    case proc == "Boiled Concentrate":
        return "BIL"
    case proc == "Fermented - Enhanced Saponin":
        return "FRM"
    default:
        return "PRO"
    }
}