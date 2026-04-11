package seeder

import (
	"debian-ecommerce/internal/data/entity"
	"fmt"
	"math/rand"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func SeedStainlessSteelStrawSet(db *gorm.DB) error {
    rand.Seed(time.Now().UnixNano())
    
    // First, ensure category exists (category ID 6 for Straw & Cutlery)
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

    // Create the main product - Stainless Steel Straw Set with Cleaning Brush
    product := entity.Product{
        CategoryID:       6,
        Name:             "Stainless Steel Straw Set with Cleaning Brush - Reusable, Travel Pouch Included",
        Description:      `Make the switch to plastic-free drinking with our premium **Stainless Steel Straw Set**. Each set includes multiple straw sizes and a dedicated cleaning brush, perfect for everything from bubble tea to smoothies and everyday beverages.

**🌱 WHY CHOOSE STAINLESS STEEL STRAWS?**

Unlike single-use plastic straws that pollute oceans and harm marine life, stainless steel straws offer a durable, reusable alternative. According to environmental studies, Indonesia is one of the largest contributors to plastic pollution in oceans, and switching to reusable straws is a simple but impactful change .

**✨ WHAT'S INCLUDED IN YOUR SET:**

- **4 Stainless Steel Straws**: 2 straight (21cm) + 2 bent (21cm) for different drinking preferences
- **1 Cleaning Brush**: Stainless steel handle with durable nylon bristles
- **1 Organic Cotton Pouch**: For hygienic storage and travel
- **1 Set of Silicone Tips**: For temperature comfort (hot/cold drinks)

**🧹 THE CLEANING BRUSH:**

Our dedicated cleaning brush features:
- **Stainless steel handle**: Rust-proof and durable
- **Nylon bristles**: Effectively removes residue without scratching
- **Long reach**: Extended length to clean full straw depth
- **Easy storage**: Fits inside pouch with straws

**📋 STRAW SPECIFICATIONS:**

| Type | Length | Diameter | Best For |
|------|--------|----------|----------|
| Straight - Regular | 21.5 cm | 6mm | Water, juice, soft drinks |
| Straight - Wide | 21.5 cm | 8mm | Smoothies, boba tea |
| Bent - Regular | 21.5 cm | 6mm | Comfort grip, cocktails |
| Bent - Wide | 21.5 cm | 8mm | Thick drinks, milkshakes |

**✨ KEY FEATURES:**

- ✅ **Food-Grade Stainless Steel**: 18/8 (304) stainless steel—no rust, no metallic taste
- ✅ **Durable Construction**: Will last for years with proper care
- ✅ **Easy to Clean**: Dishwasher safe or use included brush
- ✅ **Travel-Friendly**: Comes with cotton pouch for on-the-go use
- ✅ **Silicone Tips**: Optional tips protect teeth and lips from cold/hot drinks
- ✅ **Eco-Friendly Packaging**: Plastic-free, compostable cardboard box

**🌍 SUSTAINABILITY IMPACT:**

By switching to reusable stainless steel straws:
- Each set replaces hundreds of single-use plastic straws annually
- Reduces plastic waste in Indonesian oceans and landfills
- Saves resources used in plastic production
- Supports the global movement to #StopSucking

**📋 CARE INSTRUCTIONS:**

1. **Rinse immediately** after use to prevent residue buildup
2. **Use cleaning brush** with warm soapy water to scrub interior
3. **Dishwasher safe** (top rack recommended)
4. **Air dry** completely before storing to prevent moisture
5. **Store in cotton pouch** when traveling or not in use

**⚠️ SAFETY NOTE:**

Not recommended for children under 3 years. Adult supervision recommended for children learning to use reusable straws.

Join millions of Indonesians choosing reusable alternatives. Make every sip sustainable!`,
        IsPublished:      true,
        Tags:             pq.StringArray{"stainless steel straw", "reusable straw", "sedotan stainless", "boba straw", "cleaning brush", "zero waste", "plastic-free", "eco-friendly", "travel straw", "bubble tea straw", "sedotan", "ramah lingkungan", "sikat pembersih", "straw set", "reusable"},
        Slug:             "stainless-steel-straw-set-with-cleaning-brush-reusable-travel-pouch",
        MainImageURL:     "https://down-id.img.susercontent.com/file/ac1cd9b9c0750b191865430c1cf4cd54@resize_w900_nl.webp",
        MainImagePublicID: "stainless-straw-main",
        AverageRating:    4.8,
        ReviewCount:      1234,
        SoldCount:        5678,
        MinPrice:         0, // Will be updated after SKUs
        MaxPrice:         0, // Will be updated after SKUs
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }

    // Create Variant Types
    // Variant 1: Straw Count
    countType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Straw Count",
    }
    if err := db.Create(&countType).Error; err != nil {
        return err
    }

    // Variant 2: Straw Style
    styleType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Straw Style",
    }
    if err := db.Create(&styleType).Error; err != nil {
        return err
    }

    // Variant 3: Diameter
    diameterType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Diameter",
    }
    if err := db.Create(&diameterType).Error; err != nil {
        return err
    }

    // Variant 4: Finish Type
    finishType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Finish",
    }
    if err := db.Create(&finishType).Error; err != nil {
        return err
    }

    // Variant 5: Pouch Color
    pouchType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Pouch Color",
    }
    if err := db.Create(&pouchType).Error; err != nil {
        return err
    }

    // Create Variant Values
    // Straw Count options
    countValues := []entity.VariantValue{
        {
            VariantTypeID: countType.ID,
            Value:         "Individual - 1 Straw + Brush",
        },
        {
            VariantTypeID: countType.ID,
            Value:         "Couple Set - 2 Straws + Brush",
        },
        {
            VariantTypeID: countType.ID,
            Value:         "Family Set - 4 Straws + Brush",
        },
        {
            VariantTypeID: countType.ID,
            Value:         "Party Pack - 8 Straws + 2 Brushes",
        },
        {
            VariantTypeID: countType.ID,
            Value:         "Bulk - 12 Straws + 3 Brushes",
        },
    }
    
    for i := range countValues {
        if err := db.Create(&countValues[i]).Error; err != nil {
            return err
        }
    }

    // Straw Style options
    styleValues := []entity.VariantValue{
        {
            VariantTypeID: styleType.ID,
            Value:         "All Straight",
        },
        {
            VariantTypeID: styleType.ID,
            Value:         "All Bent",
        },
        {
            VariantTypeID: styleType.ID,
            Value:         "Mixed (Straight + Bent) - Standard",
        },
    }
    
    for i := range styleValues {
        if err := db.Create(&styleValues[i]).Error; err != nil {
            return err
        }
    }

    // Diameter options
    diameterValues := []entity.VariantValue{
        {
            VariantTypeID: diameterType.ID,
            Value:         "Regular - 6mm (Standard Drinks)",
        },
        {
            VariantTypeID: diameterType.ID,
            Value:         "Wide - 8mm (Boba/Smoothies)",
        },
        {
            VariantTypeID: diameterType.ID,
            Value:         "Jumbo - 10mm (Thick Smoothies)",
        },
        {
            VariantTypeID: diameterType.ID,
            Value:         "Mixed Diameters (6mm + 8mm + 10mm)",
        },
    }
    
    for i := range diameterValues {
        if err := db.Create(&diameterValues[i]).Error; err != nil {
            return err
        }
    }

    // Finish options
    finishValues := []entity.VariantValue{
        {
            VariantTypeID: finishType.ID,
            Value:         "Brushed Stainless - Matte",
        },
        {
            VariantTypeID: finishType.ID,
            Value:         "Polished Stainless - Mirror",
        },
        {
            VariantTypeID: finishType.ID,
            Value:         "Colored Coated - Gold",
        },
        {
            VariantTypeID: finishType.ID,
            Value:         "Colored Coated - Rose Gold",
        },
        {
            VariantTypeID: finishType.ID,
            Value:         "Colored Coated - Black",
        },
        {
            VariantTypeID: finishType.ID,
            Value:         "Colored Coated - Rainbow",
        },
    }
    
    for i := range finishValues {
        if err := db.Create(&finishValues[i]).Error; err != nil {
            return err
        }
    }

    // Pouch Color options
    pouchValues := []entity.VariantValue{
        {
            VariantTypeID: pouchType.ID,
            Value:         "Natural Cotton - Beige",
        },
        {
            VariantTypeID: pouchType.ID,
            Value:         "Black",
        },
        {
            VariantTypeID: pouchType.ID,
            Value:         "Navy Blue",
        },
        {
            VariantTypeID: pouchType.ID,
            Value:         "Forest Green",
        },
        {
            VariantTypeID: pouchType.ID,
            Value:         "Terracotta",
        },
        {
            VariantTypeID: pouchType.ID,
            Value:         "Batik Pattern - Kawung",
        },
        {
            VariantTypeID: pouchType.ID,
            Value:         "No Pouch (Eco Option)",
        },
    }
    
    for i := range pouchValues {
        if err := db.Create(&pouchValues[i]).Error; err != nil {
            return err
        }
    }

    // Base price mapping based on market research
    // Source: Tokopedia, Shopee, and IKEA Indonesia pricing [citation:1]
    
    basePriceMap := map[string]map[string]float64{
        "Individual - 1 Straw + Brush": {
            "Regular - 6mm (Standard Drinks)":       35000,
            "Wide - 8mm (Boba/Smoothies)":           38000,
            "Jumbo - 10mm (Thick Smoothies)":        42000,
        },
        "Couple Set - 2 Straws + Brush": {
            "Regular - 6mm (Standard Drinks)":       55000,
            "Wide - 8mm (Boba/Smoothies)":           59000,
            "Jumbo - 10mm (Thick Smoothies)":        65000,
            "Mixed Diameters (6mm + 8mm + 10mm)":    72000,
        },
        "Family Set - 4 Straws + Brush": {
            "Regular - 6mm (Standard Drinks)":       89000,
            "Wide - 8mm (Boba/Smoothies)":           95000,
            "Jumbo - 10mm (Thick Smoothies)":        105000,
            "Mixed Diameters (6mm + 8mm + 10mm)":    115000,
        },
        "Party Pack - 8 Straws + 2 Brushes": {
            "Regular - 6mm (Standard Drinks)":       155000,
            "Wide - 8mm (Boba/Smoothies)":           165000,
            "Jumbo - 10mm (Thick Smoothies)":        185000,
            "Mixed Diameters (6mm + 8mm + 10mm)":    199000,
        },
        "Bulk - 12 Straws + 3 Brushes": {
            "Regular - 6mm (Standard Drinks)":       215000,
            "Wide - 8mm (Boba/Smoothies)":           229000,
            "Jumbo - 10mm (Thick Smoothies)":        255000,
            "Mixed Diameters (6mm + 8mm + 10mm)":    275000,
        },
    }

    // Style premium adjustments
    stylePremium := map[string]float64{
        "All Straight":            0,
        "All Bent":                 2000,
        "Mixed (Straight + Bent) - Standard":  3000,
    }

    // Finish premium adjustments
    finishPremium := map[string]float64{
        "Brushed Stainless - Matte":       0,
        "Polished Stainless - Mirror":     5000,
        "Colored Coated - Gold":           15000,
        "Colored Coated - Rose Gold":      15000,
        "Colored Coated - Black":          12000,
        "Colored Coated - Rainbow":        20000,
    }

    // Pouch premium/discount
    pouchPremium := map[string]float64{
        "Natural Cotton - Beige":           0,
        "Black":                            0,
        "Navy Blue":                         0,
        "Forest Green":                      0,
        "Terracotta":                        0,
        "Batik Pattern - Kawung":            15000,
        "No Pouch (Eco Option)":             -5000,
    }

    var allSKUs []entity.SKU
    var minPrice float64 = 999999
    var maxPrice float64 = 0

    // Generate SKUs for combinations
    for _, count := range countValues {
        countPrices, exists := basePriceMap[count.Value]
        if !exists {
            continue
        }
        
        for _, diameter := range diameterValues {
            basePrice, diameterExists := countPrices[diameter.Value]
            if !diameterExists {
                continue
            }
            
            for _, style := range styleValues {
                // Skip incompatible count-style combinations
                if count.Value == "Individual - 1 Straw + Brush" && 
                   style.Value == "Mixed (Straight + Bent) - Standard" {
                    continue // Can't mix with single straw
                }
                
                for _, finish := range finishValues {
                    for _, pouch := range pouchValues {
                        // Skip incompatible pouch-finish combinations
                        if finish.Value == "Colored Coated - Rainbow" && 
                           pouch.Value == "Natural Cotton - Beige" {
                            // Rainbow straws often paired with fun pouch colors, but still allowed
                        }
                        
                        // Calculate final price
                        styleAdj := stylePremium[style.Value]
                        finishAdj := finishPremium[finish.Value]
                        pouchAdj := pouchPremium[pouch.Value]
                        
                        finalPrice := basePrice + styleAdj + finishAdj + pouchAdj
                        
                        // Ensure price doesn't go negative
                        if finalPrice < 15000 {
                            finalPrice = 15000
                        }
                        
                        // Round to nearest 1000 Rupiah
                        finalPrice = float64(int(finalPrice/1000)) * 1000
                        
                        // Generate SKU code
                        skuCode := fmt.Sprintf("STR-%s-%s-%s-%s-%s-%d", 
                            abbreviateCount(count.Value),
                            abbreviateDiameter(diameter.Value),
                            abbreviateStyle(style.Value),
                            abbreviateSSFinish(finish.Value),
                            abbreviatePouch(pouch.Value),
                            rand.Intn(1000))
                        
                        // Calculate sale price (30% chance of being on sale)
                        var salePrice *float64
                        if rand.Float64() < 0.3 {
                            discount := 0.80 + (rand.Float64() * 0.15) // 5-20% off
                            sp := finalPrice * discount
                            sp = float64(int(sp/1000)) * 1000
                            salePrice = &sp
                        }
                        
                        // Determine stock levels based on popularity
                        stock := 0
                        if count.Value == "Family Set - 4 Straws + Brush" && 
                           finish.Value == "Brushed Stainless - Matte" {
                            stock = 200 + rand.Intn(300) // Most popular: 200-500
                        } else if count.Value == "Individual - 1 Straw + Brush" {
                            stock = 150 + rand.Intn(250) // 150-400
                        } else if finish.Value == "Colored Coated - Rainbow" {
                            stock = 30 + rand.Intn(70) // 30-100 (niche)
                        } else if count.Value == "Bulk - 12 Straws + 3 Brushes" {
                            stock = 20 + rand.Intn(40) // 20-60
                        } else {
                            stock = 50 + rand.Intn(150) // 50-200
                        }
                        
                        // Adjust stock for premium variants
                        if finishAdj > 10000 || pouchAdj > 10000 {
                            stock = stock / 2
                        }
                        
                        // Calculate weight based on count
                        baseWeight := 20.0 // grams per straw
                        brushWeight := 15.0 // grams per brush
                        
                        // Parse count to get number of straws and brushes
                        strawCount := 1
                        brushCount := 1
                        switch count.Value {
                        case "Individual - 1 Straw + Brush":
                            strawCount = 1
                            brushCount = 1
                        case "Couple Set - 2 Straws + Brush":
                            strawCount = 2
                            brushCount = 1
                        case "Family Set - 4 Straws + Brush":
                            strawCount = 4
                            brushCount = 1
                        case "Party Pack - 8 Straws + 2 Brushes":
                            strawCount = 8
                            brushCount = 2
                        case "Bulk - 12 Straws + 3 Brushes":
                            strawCount = 12
                            brushCount = 3
                        }
                        
                        // Add pouch weight (approx 30g)
                        pouchWeight := 30.0
                        if pouch.Value == "No Pouch (Eco Option)" {
                            pouchWeight = 0
                        }
                        
                        totalWeight := (float64(strawCount) * baseWeight) + (float64(brushCount) * brushWeight) + pouchWeight
                        
                        sku := entity.SKU{
                            ProductID: product.ID,
                            SKUCode:   skuCode,
                            Price:     finalPrice,
                            SalePrice: salePrice,
                            Stock:     stock,
                            MinStock:  15,
                            Status:    entity.SKUStatusActive,
                            Weight:    totalWeight,
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
                            VariantValueID: count.ID,
                        }
                        db.Create(&skuVariant1)
                        
                        skuVariant2 := entity.SKUVariantValue{
                            SKUID:          sku.ID,
                            VariantValueID: diameter.ID,
                        }
                        db.Create(&skuVariant2)
                        
                        skuVariant3 := entity.SKUVariantValue{
                            SKUID:          sku.ID,
                            VariantValueID: style.ID,
                        }
                        db.Create(&skuVariant3)
                        
                        skuVariant4 := entity.SKUVariantValue{
                            SKUID:          sku.ID,
                            VariantValueID: finish.ID,
                        }
                        db.Create(&skuVariant4)
                        
                        skuVariant5 := entity.SKUVariantValue{
                            SKUID:          sku.ID,
                            VariantValueID: pouch.ID,
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
			"https://down-id.img.susercontent.com/file/id-11134207-7r98v-lrwwhgxi0eo9a0.webp",
			"https://down-id.img.susercontent.com/file/d6bbcb2b9502fd90254fe7f35d788634.webp",
			"https://down-id.img.susercontent.com/file/29c865f330b5467a2f4cc5202f5f8331.webp",
		}

		for _, url := range imageURLs {
			img := entity.Image{
				ProductID: product.ID,
				ImageURL: url,
			}

			db.Create(&img)
		}

    fmt.Printf("✅ Successfully seeded Stainless Steel Straw Set product with %d SKU variants\n", len(allSKUs))
    fmt.Printf("   Price range: Rp %.0f - Rp %.0f\n", minPrice, maxPrice)
    
    return nil
}

// Helper functions for SKU code generation
func abbreviateCount(count string) string {
    switch {
    case count == "Individual - 1 Straw + Brush":
        return "IND"
    case count == "Couple Set - 2 Straws + Brush":
        return "CPL"
    case count == "Family Set - 4 Straws + Brush":
        return "FAM"
    case count == "Party Pack - 8 Straws + 2 Brushes":
        return "PRT"
    case count == "Bulk - 12 Straws + 3 Brushes":
        return "BLK"
    default:
        return "CNT"
    }
}

func abbreviateDiameter(diameter string) string {
    switch {
    case diameter == "Regular - 6mm (Standard Drinks)":
        return "R6"
    case diameter == "Wide - 8mm (Boba/Smoothies)":
        return "W8"
    case diameter == "Jumbo - 10mm (Thick Smoothies)":
        return "J10"
    case diameter == "Mixed Diameters (6mm + 8mm + 10mm)":
        return "MIX"
    default:
        return "DIA"
    }
}

func abbreviateStyle(style string) string {
    switch {
    case style == "All Straight":
        return "ST"
    case style == "All Bent":
        return "BN"
    case style == "Mixed (Straight + Bent) - Standard":
        return "MX"
    default:
        return "STY"
    }
}

func abbreviateSSFinish(finish string) string {
    switch {
    case finish == "Brushed Stainless - Matte":
        return "BRS"
    case finish == "Polished Stainless - Mirror":
        return "POL"
    case finish == "Colored Coated - Gold":
        return "GLD"
    case finish == "Colored Coated - Rose Gold":
        return "RSG"
    case finish == "Colored Coated - Black":
        return "BLK"
    case finish == "Colored Coated - Rainbow":
        return "RNB"
    default:
        return "FIN"
    }
}

func abbreviatePouch(pouch string) string {
    switch {
    case pouch == "Natural Cotton - Beige":
        return "NC"
    case pouch == "Black":
        return "BLK"
    case pouch == "Navy Blue":
        return "NVY"
    case pouch == "Forest Green":
        return "GRN"
    case pouch == "Terracotta":
        return "TER"
    case pouch == "Batik Pattern - Kawung":
        return "BTK"
    case pouch == "No Pouch (Eco Option)":
        return "NOP"
    default:
        return "PCH"
    }
}