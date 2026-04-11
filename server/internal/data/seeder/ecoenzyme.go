package seeder

import (
	"debian-ecommerce/internal/data/entity"
	"fmt"
	"math/rand"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func SeedEcoEnzyme(db *gorm.DB) error {
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

    // Create the main product - Eco-Enzyme
    product := entity.Product{
        CategoryID:       3,
        Name:             "Eco-Enzyme Multipurpose Cleaner - 100% Natural from Fermented Fruit Waste",
        Description:      `Eco-enzyme is a versatile, multipurpose liquid produced through the natural fermentation of organic waste—typically fresh fruit peels, brown sugar (molasses), and water. This remarkable amber-colored liquid is the result of a 3-month fermentation process that harnesses the power of beneficial microorganisms to create a potent, eco-friendly solution for modern households [citation:2][citation:6].

**🌱 What Makes Eco-Enzyme Special?**

The fermentation process follows a precise 1:3:10 ratio—1 part brown sugar, 3 parts fruit/vegetable peels, and 10 parts water [citation:2][citation:3][citation:10]. During 90 days of fermentation, the mixture produces enzymes, organic acids (particularly acetic acid), and bioactive compounds that give eco-enzyme its cleaning, antibacterial, and deodorizing properties [citation:2][citation:10]. The final product has a characteristic sour-sweet fermented smell and a pH below 4.0, indicating successful fermentation [citation:2].

Unlike chemical cleaners, eco-enzyme is completely biodegradable and actually benefits the environment—during fermentation, it produces ozone (O₃) which helps reduce the greenhouse effect [citation:10].

**🧺 MULTIPLE USES:**

**Home Cleaning Applications:**
- **Floor Cleaner**: Add 30-50ml to mop water for streak-free, naturally disinfected floors [citation:1][citation:6][citation:7]
- **Kitchen Cleaner**: Effectively cuts grease on countertops, stoves, and sinks [citation:3][citation:4]
- **Drain Cleaner**: Pour down drains to break down organic buildup and eliminate odors [citation:8]
- **Glass Cleaner**: Dilute with water for streak-free windows and mirrors
- **Odor Eliminator**: Neutralizes unpleasant smells in trash cans, bathrooms, and refrigerators [citation:3][citation:7]
- **Laundry Booster**: Add 50ml to washing machine for brighter, cleaner clothes [citation:4][citation:5]

**Gardening & Outdoor Uses:**
- **Natural Fertilizer**: Dilute 1:100 with water and apply to plants for essential nutrients [citation:3][citation:6]
- **Pesticide Alternative**: Helps repel pests naturally without harmful chemicals [citation:3][citation:8]
- **Soil Conditioner**: Improves soil health and microbial activity

**Personal Care Applications:**
- **Hand Sanitizer**: Diluted solution with natural antibacterial properties (alcohol-free, gentle on skin) [citation:1][citation:4]
- **Natural Disinfectant**: Safe for surfaces that contact food [citation:8]
- **Face Wash**: Some users dilute heavily for gentle facial cleansing

**✨ KEY BENEFITS:**

- ✅ **100% Natural & Biodegradable**: No harsh chemicals, phosphates, or synthetic fragrances
- ✅ **Antibacterial Properties**: Contains acetic acid and natural enzymes that inhibit bacterial growth [citation:2][citation:5][citation:10]
- ✅ **Zero Waste Production**: Made from discarded fruit peels—each bottle diverts organic waste from landfills [citation:2]
- ✅ **Non-Toxic**: Safe for homes with children and pets
- ✅ **Gentle on Skin**: Unlike bleach or strong chemical cleaners
- ✅ **Economical**: Highly concentrated—dilute before use, one bottle lasts months
- ✅ **No Synthetic Fragrances**: Naturally scented from fermented citrus or fruit peels

**🌍 SUSTAINABILITY IMPACT:**

- **Waste Reduction**: Indonesia generated 38.7 million tonnes of waste in 2024, with 37.87% unmanaged—eco-enzyme directly addresses this crisis [citation:10]
- **Circular Economy**: Transforms waste into valuable household products [citation:3][citation:9]
- **Greenhouse Gas Reduction**: Fermentation produces ozone that helps counteract the greenhouse effect [citation:10]
- **Water Protection**: No phosphates or surfactants that cause eutrophication
- **Community Empowerment**: Increasingly used in community programs and women's groups across Indonesia [citation:2][citation:6][citation:7]

**📋 HOW ECO-ENZYME IS MADE (Traditional Method):**

The traditional 1:3:10 formula [citation:2][citation:3][citation:10]:
- **1 part** Brown sugar (molasses/gula merah)
- **3 parts** Fresh fruit peels (citrus, pineapple, papaya, banana, etc.)
- **10 parts** Water

Fermentation takes **3 months (90 days)** in an airtight container. During the first month, the container must be opened daily to release gases, then sealed for the remaining two months. The final product is strained and ready for use [citation:6][citation:8].

**🧴 AVAILABLE IN THESE VARIATIONS:**

Our eco-enzyme comes in multiple forms to suit your needs:

- **Pure Eco-Enzyme Liquid**: The standard fermented liquid, ready for dilution
- **Concentrated Eco-Enzyme**: Higher potency, more economical
- **Scented Eco-Enzyme**: Infused with specific fruit essences (citrus, lemongrass, etc.)
- **Derivative Products**: Hand soap, dishwashing liquid, and laundry detergent made with eco-enzyme base [citation:4][citation:5][citation:9]`,
        IsPublished:      true,
        Tags:             pq.StringArray{"ecoenzyme", "eco enzyme", "natural cleaner", "fermented cleaner", "multipurpose", "zero waste", "organic waste", "fruit fermentation", "biodegradable", "non-toxic", "floor cleaner", "natural disinfectant", "hand sanitizer", "plant fertilizer", "eco-friendly", "rumah tangga", "pembersih alami"},
        Slug:             "eco-enzyme-multipurpose-cleaner-natural-fermented-fruit-waste",
        MainImageURL:     "https://down-id.img.susercontent.com/file/id-11134207-7rash-m5l74q91i8zqd4.webp",
        MainImagePublicID: "eco-enzyme-main",
        AverageRating:    4.7,
        ReviewCount:      654,
        SoldCount:        2789,
        MinPrice:         0, // Will be updated after SKUs
        MaxPrice:         0, // Will be updated after SKUs
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }

    // Create Variant Types
    // Variant 1: Product Type/Formulation
    productType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Product Type",
    }
    if err := db.Create(&productType).Error; err != nil {
        return err
    }

    // Variant 2: Fruit Base / Scent
    fruitBaseType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Fruit Base / Scent",
    }
    if err := db.Create(&fruitBaseType).Error; err != nil {
        return err
    }

    // Variant 3: Size/Volume
    sizeType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Volume",
    }
    if err := db.Create(&sizeType).Error; err != nil {
        return err
    }

    // Variant 4: Fermentation Age
    ageType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Fermentation Age",
    }
    if err := db.Create(&ageType).Error; err != nil {
        return err
    }

    // Variant 5: Packaging Type
    packagingType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Packaging",
    }
    if err := db.Create(&packagingType).Error; err != nil {
        return err
    }

    // Create Variant Values
    // Product Type / Formulation
    productTypeValues := []entity.VariantValue{
        {
            VariantTypeID: productType.ID,
            Value:         "Pure Eco-Enzyme Liquid - Standard",
        },
        {
            VariantTypeID: productType.ID,
            Value:         "Concentrated Eco-Enzyme - Double Strength",
        },
        {
            VariantTypeID: productType.ID,
            Value:         "Eco-Enzyme Based Hand Soap",
        },
        {
            VariantTypeID: productType.ID,
            Value:         "Eco-Enzyme Dishwashing Liquid",
        },
        {
            VariantTypeID: productType.ID,
            Value:         "Eco-Enzyme Laundry Detergent",
        },
        {
            VariantTypeID: productType.ID,
            Value:         "Eco-Enzyme Floor Cleaner Concentrate",
        },
        {
            VariantTypeID: productType.ID,
            Value:         "Eco-Enzyme Starter Kit (DIY Materials)",
        },
    }
    
    for i := range productTypeValues {
        if err := db.Create(&productTypeValues[i]).Error; err != nil {
            return err
        }
    }

    // Fruit Base / Scent - based on common Indonesian fruit waste [citation:2][citation:4][citation:10]
    fruitBaseValues := []entity.VariantValue{
        {
            VariantTypeID: fruitBaseType.ID,
            Value:         "Mixed Citrus (Jeruk) - Orange, Lemon, Lime",
        },
        {
            VariantTypeID: fruitBaseType.ID,
            Value:         "Pineapple (Nanas) - Sweet Tropical Scent",
        },
        {
            VariantTypeID: fruitBaseType.ID,
            Value:         "Papaya (Pepaya) - Mild, Neutral Scent",
        },
        {
            VariantTypeID: fruitBaseType.ID,
            Value:         "Banana (Pisang) - Sweet, Earthy",
        },
        {
            VariantTypeID: fruitBaseType.ID,
            Value:         "Watermelon (Semangka) - Light, Fresh",
        },
        {
            VariantTypeID: fruitBaseType.ID,
            Value:         "Mixed Tropical - Pineapple, Papaya, Banana",
        },
        {
            VariantTypeID: fruitBaseType.ID,
            Value:         "Lemongrass Infused (Sereh) - Fresh, Citronella",
        },
        {
            VariantTypeID: fruitBaseType.ID,
            Value:         "Unscented / Neutral",
        },
    }
    
    for i := range fruitBaseValues {
        if err := db.Create(&fruitBaseValues[i]).Error; err != nil {
            return err
        }
    }

    // Size/Volume options
    sizeValues := []entity.VariantValue{
        {
            VariantTypeID: sizeType.ID,
            Value:         "100ml - Trial Size",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "250ml - Small Bottle",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "500ml - Standard",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "1L - Economy",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "5L - Jerigen (Refill)",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "20L - Bulk / Gallon",
        },
    }
    
    for i := range sizeValues {
        if err := db.Create(&sizeValues[i]).Error; err != nil {
            return err
        }
    }

    // Fermentation Age - older = more potent [citation:10]
    ageValues := []entity.VariantValue{
        {
            VariantTypeID: ageType.ID,
            Value:         "3 Months - Standard",
        },
        {
            VariantTypeID: ageType.ID,
            Value:         "6 Months - Premium Aged",
        },
        {
            VariantTypeID: ageType.ID,
            Value:         "1 Year - Extra Aged (Highest Potency)",
        },
    }
    
    for i := range ageValues {
        if err := db.Create(&ageValues[i]).Error; err != nil {
            return err
        }
    }

    // Packaging Type (eco-friendly options)
    packagingValues := []entity.VariantValue{
        {
            VariantTypeID: packagingType.ID,
            Value:         "Recycled Plastic Bottle",
        },
        {
            VariantTypeID: packagingType.ID,
            Value:         "Glass Bottle - Reusable",
        },
        {
            VariantTypeID: packagingType.ID,
            Value:         "Jerigen Plastic - Refillable",
        },
        {
            VariantTypeID: packagingType.ID,
            Value:         "Bring Your Own Container (BYOC) - Refill Discount",
        },
        {
            VariantTypeID: packagingType.ID,
            Value:         "Compostable PLA Pouch",
        },
        {
            VariantTypeID: packagingType.ID,
            Value:         "Aluminum Bottle - Premium",
        },
    }
    
    for i := range packagingValues {
        if err := db.Create(&packagingValues[i]).Error; err != nil {
            return err
        }
    }

    // Base price mapping based on market research [citation:1][citation:3]
    // Eco-enzyme typically sells for Rp 20,000 - 50,000 per liter [citation:3]
    
    // Base prices by product type and volume
    basePriceMap := map[string]map[string]float64{
        "Pure Eco-Enzyme Liquid - Standard": {
            "100ml - Trial Size":      12000,
            "250ml - Small Bottle":    20000,
            "500ml - Standard":        35000,
            "1L - Economy":            55000,
            "5L - Jerigen (Refill)":   225000,
            "20L - Bulk / Gallon":     750000,
        },
        "Concentrated Eco-Enzyme - Double Strength": {
            "250ml - Small Bottle":    35000,
            "500ml - Standard":        60000,
            "1L - Economy":            100000,
            "5L - Jerigen (Refill)":    425000,
        },
        "Eco-Enzyme Based Hand Soap": {
            "250ml - Small Bottle":    25000,
            "500ml - Standard":        45000,
            "1L - Economy":            75000,
        },
        "Eco-Enzyme Dishwashing Liquid": {
            "250ml - Small Bottle":    22000,
            "500ml - Standard":        40000,
            "1L - Economy":            70000,
            "5L - Jerigen (Refill)":    300000,
        },
        "Eco-Enzyme Laundry Detergent": {
            "500ml - Standard":        45000,
            "1L - Economy":            80000,
            "5L - Jerigen (Refill)":    350000,
        },
        "Eco-Enzyme Floor Cleaner Concentrate": {
            "500ml - Standard":        40000,
            "1L - Economy":            70000,
            "5L - Jerigen (Refill)":    300000,
        },
        "Eco-Enzyme Starter Kit (DIY Materials)": {
            "1 Kit - 1L Water + Molasses + Guide": 75000,
        },
    }

    // Fruit base premium (some scents more desirable)
    fruitPremium := map[string]float64{
        "Mixed Citrus (Jeruk) - Orange, Lemon, Lime":      0,      // Standard, most popular
        "Pineapple (Nanas) - Sweet Tropical Scent":        5000,   // Premium tropical
        "Papaya (Pepaya) - Mild, Neutral Scent":           0,      // Standard
        "Banana (Pisang) - Sweet, Earthy":                 -2000,  // Less popular, discount
        "Watermelon (Semangka) - Light, Fresh":            3000,   // Premium
        "Mixed Tropical - Pineapple, Papaya, Banana":      8000,   // Premium blend
        "Lemongrass Infused (Sereh) - Fresh, Citronella":  10000,  // High premium (disinfectant scent)
        "Unscented / Neutral":                              -5000,  // Discount
    }

    // Age premium - older = more potent
    agePremium := map[string]float64{
        "3 Months - Standard":             0,
        "6 Months - Premium Aged":          20000,
        "1 Year - Extra Aged (Highest Potency)": 45000,
    }

    // Packaging premium/discount
    packagingPremium := map[string]float64{
        "Recycled Plastic Bottle":              0,
        "Glass Bottle - Reusable":               10000,
        "Jerigen Plastic - Refillable":          -2000,   // Discount for bulk refill
        "Bring Your Own Container (BYOC) - Refill Discount": -10000, // Biggest discount for eco-conscious
        "Compostable PLA Pouch":                 8000,    // Premium eco material
        "Aluminum Bottle - Premium":              20000,   // Premium packaging
    }

    var allSKUs []entity.SKU
    var minPrice float64 = 999999
    var maxPrice float64 = 0

    // Generate SKUs for combinations
    for _, prodType := range productTypeValues {
        typePrices, exists := basePriceMap[prodType.Value]
        if !exists {
            continue // Skip product types without price mapping
        }
        
        for _, size := range sizeValues {
            // Check if this size exists for this product type
            basePrice, sizeExists := typePrices[size.Value]
            if !sizeExists {
                continue
            }
            
            // Skip size combinations that don't make sense
            if prodType.Value == "Eco-Enzyme Starter Kit (DIY Materials)" && size.Value != "1 Kit - 1L Water + Molasses + Guide" {
                continue
            }
            
            for _, fruitBase := range fruitBaseValues {
                // Skip incompatible fruit bases
                // Starter kits only come with mixed fruit or unscented
                if prodType.Value == "Eco-Enzyme Starter Kit (DIY Materials)" && 
                   fruitBase.Value != "Mixed Citrus (Jeruk) - Orange, Lemon, Lime" && 
                   fruitBase.Value != "Unscented / Neutral" {
                    continue
                }
                
                // Derivative products (soap, detergent) typically use citrus or neutral
                if (prodType.Value == "Eco-Enzyme Based Hand Soap" || 
                    prodType.Value == "Eco-Enzyme Dishwashing Liquid" || 
                    prodType.Value == "Eco-Enzyme Laundry Detergent") && 
                   (fruitBase.Value == "Banana (Pisang) - Sweet, Earthy" || 
                    fruitBase.Value == "Watermelon (Semangka) - Light, Fresh") {
                    continue
                }
                
                for _, age := range ageValues {
                    // Age options only for pure liquid, not derivatives
                    if (prodType.Value != "Pure Eco-Enzyme Liquid - Standard" && 
                        prodType.Value != "Concentrated Eco-Enzyme - Double Strength") && 
                       age.Value != "3 Months - Standard" {
                        continue
                    }
                    
                    for _, packaging := range packagingValues {
                        // Skip incompatible packaging-size combinations
                        // Glass bottles only for smaller sizes
                        if packaging.Value == "Glass Bottle - Reusable" && 
                           (size.Value == "5L - Jerigen (Refill)" || size.Value == "20L - Bulk / Gallon") {
                            continue
                        }
                        
                        // Aluminum bottle only for premium sizes
                        if packaging.Value == "Aluminum Bottle - Premium" && 
                           (size.Value == "100ml - Trial Size" || size.Value == "20L - Bulk / Gallon") {
                            continue
                        }
                        
                        // BYOC only for refill sizes
                        if packaging.Value == "Bring Your Own Container (BYOC) - Refill Discount" && 
                           !(size.Value == "5L - Jerigen (Refill)" || size.Value == "1L - Economy") {
                            continue
                        }
                        
                        // Compostable pouch only for smaller sizes
                        if packaging.Value == "Compostable PLA Pouch" && 
                           (size.Value == "5L - Jerigen (Refill)" || size.Value == "20L - Bulk / Gallon") {
                            continue
                        }
                        
                        // Calculate final price
                        fruitAdj := fruitPremium[fruitBase.Value]
                        ageAdj := agePremium[age.Value]
                        packagingAdj := packagingPremium[packaging.Value]
                        
                        finalPrice := basePrice + fruitAdj + ageAdj + packagingAdj
                        
                        // Ensure price doesn't go negative with discounts
                        if finalPrice < 5000 {
                            finalPrice = 5000
                        }
                        
                        // Round to nearest 500 Rupiah (common in Indonesia)
                        finalPrice = float64(int(finalPrice/500)) * 500
                        
                        // Generate SKU code
                        skuCode := fmt.Sprintf("ECO-%s-%s-%s-%s-%s-%d", 
                            abbreviateProductType(prodType.Value),
                            abbreviateFruitBase(fruitBase.Value),
                            abbreviateEcoenzymeSize(size.Value),
                            abbreviateAge(age.Value),
                            abbreviateEcoenzymePackaging(packaging.Value),
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
                        if prodType.Value == "Pure Eco-Enzyme Liquid - Standard" && size.Value == "500ml - Standard" {
                            stock = 200 + rand.Intn(300) // Most popular: 200-500
                        } else if prodType.Value == "Pure Eco-Enzyme Liquid - Standard" && size.Value == "1L - Economy" {
                            stock = 150 + rand.Intn(250) // 150-400
                        } else if prodType.Value == "Eco-Enzyme Based Hand Soap" && size.Value == "250ml - Small Bottle" {
                            stock = 100 + rand.Intn(200) // 100-300
                        } else if size.Value == "5L - Jerigen (Refill)" {
                            stock = 30 + rand.Intn(70) // 30-100 (bulk)
                        } else if size.Value == "20L - Bulk / Gallon" {
                            stock = 10 + rand.Intn(30) // 10-40 (very bulk)
                        } else {
                            stock = 50 + rand.Intn(150) // 50-200
                        }
                        
                        // Adjust stock for premium variants
                        if fruitAdj > 5000 || ageAdj > 10000 || packagingAdj > 5000 {
                            stock = stock / 2
                        }
                        
                        // Calculate weight (varies by volume)
                        weight := 0.0
                        switch size.Value {
                        case "100ml - Trial Size":
                            weight = 120
                        case "250ml - Small Bottle":
                            weight = 280
                        case "500ml - Standard":
                            weight = 550
                        case "1L - Economy":
                            weight = 1050
                        case "5L - Jerigen (Refill)":
                            weight = 5200
                        case "20L - Bulk / Gallon":
                            weight = 20500
                        case "1 Kit - 1L Water + Molasses + Guide":
                            weight = 1200
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
                            VariantValueID: prodType.ID,
                        }
                        db.Create(&skuVariant1)
                        
                        skuVariant2 := entity.SKUVariantValue{
                            SKUID:          sku.ID,
                            VariantValueID: fruitBase.ID,
                        }
                        db.Create(&skuVariant2)
                        
                        skuVariant3 := entity.SKUVariantValue{
                            SKUID:          sku.ID,
                            VariantValueID: size.ID,
                        }
                        db.Create(&skuVariant3)
                        
                        skuVariant4 := entity.SKUVariantValue{
                            SKUID:          sku.ID,
                            VariantValueID: age.ID,
                        }
                        db.Create(&skuVariant4)
                        
                        skuVariant5 := entity.SKUVariantValue{
                            SKUID:          sku.ID,
                            VariantValueID: packaging.ID,
                        }
                        db.Create(&skuVariant5)
                        
                        allSKUs = append(allSKUs, sku)
                    }
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
			"https://down-id.img.susercontent.com/file/id-11134207-7rasa-m5bl9p7q4q2ud4@resize_w900_nl.webp",
			"https://down-id.img.susercontent.com/file/id-11134207-7rasm-m5bl9p7qaccm68.webp",
			"https://down-id.img.susercontent.com/file/id-11134207-7ras9-m5l74q91ffuufd.webp",
		}

		for _, url := range imageURLs {
			img := entity.Image{
				ProductID: product.ID,
				ImageURL: url,
			}

			db.Create(&img)
		}

    fmt.Printf("✅ Successfully seeded Eco-Enzyme product with %d SKU variants\n", len(allSKUs))
    fmt.Printf("   Price range: Rp %.0f - Rp %.0f\n", minPrice, maxPrice)
    
    return nil
}

// Helper functions for SKU code generation
func abbreviateProductType(prodType string) string {
    switch {
    case prodType == "Pure Eco-Enzyme Liquid - Standard":
        return "PURE"
    case prodType == "Concentrated Eco-Enzyme - Double Strength":
        return "CONC"
    case prodType == "Eco-Enzyme Based Hand Soap":
        return "SOAP"
    case prodType == "Eco-Enzyme Dishwashing Liquid":
        return "DISH"
    case prodType == "Eco-Enzyme Laundry Detergent":
        return "LAUN"
    case prodType == "Eco-Enzyme Floor Cleaner Concentrate":
        return "FLR"
    case prodType == "Eco-Enzyme Starter Kit (DIY Materials)":
        return "DIY"
    default:
        return "TYPE"
    }
}

func abbreviateFruitBase(fruit string) string {
    switch {
    case fruit == "Mixed Citrus (Jeruk) - Orange, Lemon, Lime":
        return "CIT"
    case fruit == "Pineapple (Nanas) - Sweet Tropical Scent":
        return "PIN"
    case fruit == "Papaya (Pepaya) - Mild, Neutral Scent":
        return "PAP"
    case fruit == "Banana (Pisang) - Sweet, Earthy":
        return "BAN"
    case fruit == "Watermelon (Semangka) - Light, Fresh":
        return "WAT"
    case fruit == "Mixed Tropical - Pineapple, Papaya, Banana":
        return "TROP"
    case fruit == "Lemongrass Infused (Sereh) - Fresh, Citronella":
        return "SER"
    case fruit == "Unscented / Neutral":
        return "NEU"
    default:
        return "FRT"
    }
}

func abbreviateEcoenzymeSize(size string) string {
    switch {
    case size == "100ml - Trial Size":
        return "100M"
    case size == "250ml - Small Bottle":
        return "250M"
    case size == "500ml - Standard":
        return "500M"
    case size == "1L - Economy":
        return "1L"
    case size == "5L - Jerigen (Refill)":
        return "5L"
    case size == "20L - Bulk / Gallon":
        return "20L"
    case size == "1 Kit - 1L Water + Molasses + Guide":
        return "KIT"
    default:
        return "SIZ"
    }
}

func abbreviateAge(age string) string {
    switch {
    case age == "3 Months - Standard":
        return "3M"
    case age == "6 Months - Premium Aged":
        return "6M"
    case age == "1 Year - Extra Aged (Highest Potency)":
        return "1Y"
    default:
        return "AGE"
    }
}

func abbreviateEcoenzymePackaging(pkg string) string {
    switch {
    case pkg == "Recycled Plastic Bottle":
        return "PLA"
    case pkg == "Glass Bottle - Reusable":
        return "GLS"
    case pkg == "Jerigen Plastic - Refillable":
        return "JER"
    case pkg == "Bring Your Own Container (BYOC) - Refill Discount":
        return "BYOC"
    case pkg == "Compostable PLA Pouch":
        return "CPLA"
    case pkg == "Aluminum Bottle - Premium":
        return "ALM"
    default:
        return "PKG"
    }
}