package seeder

import (
	"debian-ecommerce/internal/data/entity"
	"fmt"
	"math/rand"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func SeedCompostableCuttingResistantTrays(db *gorm.DB) error {
    rand.Seed(time.Now().UnixNano())
    
    // First, ensure category exists (assuming category ID 6 is Straw & Cutlery, or create new)
    // For this product, category ID 6 (Straw & Cutlery) or 3 (Home & Cleaning) would fit
    var category entity.Category
    result := db.First(&category, 6)
    if result.Error != nil {
        // Create category if it doesn't exist
        category = entity.Category{
            Model:       gorm.Model{ID: 6},
            Name:        "Straw & Cutlery",
        }
        if err := db.Create(&category).Error; err != nil {
            return err
        }
    }

    // Create the main product - Compostable Cutting-Resistant Food Trays
    product := entity.Product{
        CategoryID:       6,
        Name:             "Heavy-Duty Compostable Food Trays - Cutting-Resistant, Microwave-Safe | 3-Compartment",
        Description:      `Introducing our premium **Cutting-Resistant Compostable Food Trays**—the perfect eco-friendly solution for meal delivery services, catering, school lunch programs, and the Indonesian government's "Makan Bergizi Gratis" (Free Nutritious Meal) initiative. These heavy-duty trays are engineered to withstand knife contact while being fully compostable, offering the durability of plastic without the environmental harm.

**🌱 THE INNOVATION: CUTTING-RESISTANT COMPOSTABLE MATERIAL**

Unlike standard paper or bagasse trays that tear or leak when cut, our trays feature advanced **cutting-resistant technology** that allows users to cut food directly on the tray without penetration. This is achieved through:

- **High-density bagasse fiber construction** with proprietary compression technology
- **Reinforced fiber matrix** that resists knife pressure while remaining fully compostable
- **Natural plant-based binders** that enhance strength without synthetic additives

**🇮🇩 MEETING INDONESIA'S GROWING DEMAND**

The Indonesian food tray market is projected to grow from **$428.52 million in 2025 to $446.13 million in 2026**, with a CAGR of 4.11% through 2031 [citation:1]. Key drivers include:

- **"Free Nutritious Meal Program" (Makan Bergizi Gratis)** : This government initiative operates at **30,000 supply points** across Indonesia, creating massive institutional demand for durable, food-safe trays [citation:1]
- **Online food delivery growth**: App-based meal delivery continues to expand in metropolitan areas, with urban households spending **35.35% of monthly food budget on prepared meals** (vs. 26.32% in rural areas) [citation:2]
- **Modern convenience store expansion**: Thousands of branded minimarts expanding to regional cities require barrier packaging for ready-to-eat meals [citation:2]

**✨ KEY FEATURES:**

- ✅ **Cutting-Resistant Technology**: Specially engineered to withstand knife contact during eating—no tearing or penetration
- ✅ **3-Compartment Design**: Perfect for balanced meals with separate sections for rice, protein, and vegetables
- ✅ **Heavy-Duty Construction**: Holds hot, saucy foods without leaking or softening
- ✅ **Microwave Safe**: Reheat food directly in the tray (up to 3 minutes)
- ✅ **Freezer Safe**: Suitable for frozen meal prep
- ✅ **100% Compostable**: Biodegrades in commercial composting facilities within 60-90 days
- ✅ **Chemical-Free**: No plastic lining, wax coating, or PFAS—just natural plant fibers

**🌿 MATERIAL OPTIONS:**

Our trays are available in three premium compostable materials:

1. **Bagasse (Sugarcane Fiber)** : Made from sugarcane pulp—a byproduct of sugar production. Most popular choice, excellent balance of strength and cost [citation:3]
2. **Wheat Straw Fiber**: Made from agricultural waste; naturally lighter color, good strength
3. **Bamboo Fiber**: Premium option; naturally antibacterial, superior strength, fastest renewability

**📋 PRODUCT SPECIFICATIONS:**

| Specification | Details |
|---------------|---------|
| Dimensions | 9.5" x 7.5" x 1.2" (24cm x 19cm x 3cm) |
| Compartments | 3 sections (rice section: 500ml, protein section: 250ml, vegetable section: 200ml) |
| Material Thickness | 3.5mm (heavy-duty) |
| Temperature Range | -20°C to 120°C |
| Microwave Safe | Yes (up to 3 minutes) |
| Composting Time | 60-90 days in commercial facility |
| Certification | OK Compost, BPOM compliant |

**🎯 PERFECT APPLICATIONS:**

- **School Lunch Programs**: Ideal for "Makan Bergizi Gratis" government feeding program at 30,000+ locations [citation:1]
- **Restaurant Delivery**: Durable enough for saucy Indonesian dishes like rendang, gulai, and sambal-based meals
- **Catering Events**: Pre-plated meals for corporate events and weddings
- **Meal Prep Services**: Ready-to-eat refrigerated meals for convenience stores
- **Hospital & Institutional Food Service**: Safe for patient meals requiring utensil use

**🌍 SUSTAINABILITY IMPACT:**

The Indonesian plastics industry faces increasing regulatory pressure, with proposed taxes on virgin plastics and stricter BPOM regulations on food contact chemicals [citation:1]. By switching to compostable trays:

- Each tray diverts plastic waste from oceans and landfills
- Supports Indonesia's commitment to reducing marine plastic debris by 70% by 2025
- Creates demand for agricultural waste, providing additional income for farmers
- Bioplastics segment growing at **6.55% CAGR**, outpacing conventional plastics [citation:2]

**📋 HOW TO USE & DISPOSE:**

1. **Use**: Tray is ready to use—no preparation needed. Microwave safe for reheating.
2. **Dispose**: After use, place in compost bin or organic waste stream. Do not recycle with paper (food-contaminated).
3. **Composting**: Breaks down in commercial composting facilities within 60-90 days. Home composting may take longer.

**🔬 REGULATORY COMPLIANCE:**

All trays comply with Indonesian BPOM (Food and Drug Supervisory Agency) standards for food contact materials. Our manufacturing process meets the proposed stricter migration limits for additives, ensuring safety for hot food applications [citation:1].

Make the switch to sustainable food service packaging without compromising on durability. Perfect for businesses serving Indonesia's growing appetite for convenient, eco-friendly meal solutions!`,
        IsPublished:      true,
        Tags:             pq.StringArray{"compostable trays", "food trays", "bagasse trays", "sugarcane fiber", "cutting resistant", "microwave safe", "3 compartment", "meal delivery", "takeaway packaging", "eco-friendly", "biodegradable", "makan bergizi gratis", "school lunch", "catering", "zero waste", "plant-based", "heavy-duty", "restaurant supplies", "food packaging", "ramah lingkungan"},
        Slug:             "heavy-duty-compostable-food-trays-cutting-resistant-3-compartment",
        MainImageURL:     "https://ecolipak.com/cdn/shop/files/14-in-disposable-cutting-resistant-compostable-traysbpi-certified-pfas-free-6290227.jpg?v=1762583270&width=1024",
        MainImagePublicID: "compostable-trays-main",
        AverageRating:    4.6,
        ReviewCount:      187,
        SoldCount:        892,
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

    // Variant 2: Compartment Configuration
    compartmentType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Compartment Style",
    }
    if err := db.Create(&compartmentType).Error; err != nil {
        return err
    }

    // Variant 3: Size/Capacity
    sizeType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Size",
    }
    if err := db.Create(&sizeType).Error; err != nil {
        return err
    }

    // Variant 4: Pack Quantity
    quantityType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Pack Quantity",
    }
    if err := db.Create(&quantityType).Error; err != nil {
        return err
    }

    // Create Variant Values
    // Material Types (compostable options)
    materialValues := []entity.VariantValue{
        {
            VariantTypeID: materialType.ID,
            Value:         "Bagasse (Sugarcane Fiber) - Standard",
        },
        {
            VariantTypeID: materialType.ID,
            Value:         "Wheat Straw Fiber - Natural Light",
        },
        {
            VariantTypeID: materialType.ID,
            Value:         "Bamboo Fiber - Premium",
        },
        {
            VariantTypeID: materialType.ID,
            Value:         "Recycled Paper Pulp - Economy",
        },
    }
    
    for i := range materialValues {
        if err := db.Create(&materialValues[i]).Error; err != nil {
            return err
        }
    }

    // Compartment Configurations
    compartmentValues := []entity.VariantValue{
        {
            VariantTypeID: compartmentType.ID,
            Value:         "3-Compartment - Rice + Protein + Vegetable",
        },
        {
            VariantTypeID: compartmentType.ID,
            Value:         "2-Compartment - Main + Side",
        },
        {
            VariantTypeID: compartmentType.ID,
            Value:         "1-Compartment - Single Large (Nasi Box Style)",
        },
        {
            VariantTypeID: compartmentType.ID,
            Value:         "5-Compartment - Tiffin Style (Complete Meal)",
        },
        {
            VariantTypeID: compartmentType.ID,
            Value:         "Divided Tray with Lid Slot - For Delivery",
        },
    }
    
    for i := range compartmentValues {
        if err := db.Create(&compartmentValues[i]).Error; err != nil {
            return err
        }
    }

    // Size Options
    sizeValues := []entity.VariantValue{
        {
            VariantTypeID: sizeType.ID,
            Value:         "Standard - 9.5x7.5 inches (Adult Meal)",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "Junior - 7.5x5.5 inches (Child/Snack)",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "Large - 11x8.5 inches (Hearty Meal)",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "Mini - 6x4.5 inches (Appetizer/Sample)",
        },
    }
    
    for i := range sizeValues {
        if err := db.Create(&sizeValues[i]).Error; err != nil {
            return err
        }
    }

    // Pack Quantity (based on common Indonesian e-commerce pack sizes)
    quantityValues := []entity.VariantValue{
        {
            VariantTypeID: quantityType.ID,
            Value:         "Sample Pack - 10 trays",
        },
        {
            VariantTypeID: quantityType.ID,
            Value:         "Home Pack - 25 trays",
        },
        {
            VariantTypeID: quantityType.ID,
            Value:         "Small Business Pack - 50 trays",
        },
        {
            VariantTypeID: quantityType.ID,
            Value:         "Restaurant Starter - 100 trays",
        },
        {
            VariantTypeID: quantityType.ID,
            Value:         "Catering Pack - 250 trays",
        },
        {
            VariantTypeID: quantityType.ID,
            Value:         "Bulk Wholesale - 500 trays",
        },
        {
            VariantTypeID: quantityType.ID,
            Value:         "Institutional (School Program) - 1000 trays",
        },
        {
            VariantTypeID: quantityType.ID,
            Value:         "Pallet - 5000 trays (Container Load)",
        },
    }
    
    for i := range quantityValues {
        if err := db.Create(&quantityValues[i]).Error; err != nil {
            return err
        }
    }

    // Base price mapping per tray (converted to IDR)
    // Based on market pricing for compostable trays [citation:3]
    
    // Prices per tray (before quantity discount)
    basePricePerTrayMap := map[string]map[string]float64{
        "Bagasse (Sugarcane Fiber) - Standard": {
            "Standard - 9.5x7.5 inches (Adult Meal)": 3500,
            "Junior - 7.5x5.5 inches (Child/Snack)":  2500,
            "Large - 11x8.5 inches (Hearty Meal)":    4500,
            "Mini - 6x4.5 inches (Appetizer/Sample)": 1800,
        },
        "Wheat Straw Fiber - Natural Light": {
            "Standard - 9.5x7.5 inches (Adult Meal)": 4000,
            "Junior - 7.5x5.5 inches (Child/Snack)":  3000,
            "Large - 11x8.5 inches (Hearty Meal)":    5000,
            "Mini - 6x4.5 inches (Appetizer/Sample)": 2200,
        },
        "Bamboo Fiber - Premium": {
            "Standard - 9.5x7.5 inches (Adult Meal)": 5500,
            "Junior - 7.5x5.5 inches (Child/Snack)":  4200,
            "Large - 11x8.5 inches (Hearty Meal)":    6800,
            "Mini - 6x4.5 inches (Appetizer/Sample)": 3200,
        },
        "Recycled Paper Pulp - Economy": {
            "Standard - 9.5x7.5 inches (Adult Meal)": 2800,
            "Junior - 7.5x5.5 inches (Child/Snack)":  2000,
            "Large - 11x8.5 inches (Hearty Meal)":    3500,
            "Mini - 6x4.5 inches (Appetizer/Sample)": 1500,
        },
    }

    // Compartment premium adjustments
    compartmentPremium := map[string]float64{
        "3-Compartment - Rice + Protein + Vegetable":   0,      // Standard
        "2-Compartment - Main + Side":                  -200,   // Slightly less material
        "1-Compartment - Single Large (Nasi Box Style)": -500,   // Less complex molding
        "5-Compartment - Tiffin Style (Complete Meal)": 1500,   // Premium for multiple sections
        "Divided Tray with Lid Slot - For Delivery":    2000,   // Specialized design
    }

    // Quantity discount tiers
    quantityDiscount := map[string]float64{
        "Sample Pack - 10 trays":                    1.0,   // No discount
        "Home Pack - 25 trays":                      0.95,  // 5% off
        "Small Business Pack - 50 trays":            0.90,  // 10% off
        "Restaurant Starter - 100 trays":            0.85,  // 15% off
        "Catering Pack - 250 trays":                 0.80,  // 20% off
        "Bulk Wholesale - 500 trays":                0.75,  // 25% off
        "Institutional (School Program) - 1000 trays": 0.70, // 30% off (government program)
        "Pallet - 5000 trays (Container Load)":       0.65,  // 35% off (wholesale)
    }

    // Minimum stock levels by pack quantity
    minStockByPack := map[string]int{
        "Sample Pack - 10 trays":                    20,
        "Home Pack - 25 trays":                      15,
        "Small Business Pack - 50 trays":            10,
        "Restaurant Starter - 100 trays":            8,
        "Catering Pack - 250 trays":                 5,
        "Bulk Wholesale - 500 trays":                3,
        "Institutional (School Program) - 1000 trays": 2,
        "Pallet - 5000 trays (Container Load)":      1,
    }

    var allSKUs []entity.SKU
    var minPrice float64 = 999999
    var maxPrice float64 = 0

    // Generate SKUs for combinations
    for _, material := range materialValues {
        materialPrices, exists := basePricePerTrayMap[material.Value]
        if !exists {
            continue
        }
        
        for _, size := range sizeValues {
            pricePerTray, sizeExists := materialPrices[size.Value]
            if !sizeExists {
                continue
            }
            
            for _, compartment := range compartmentValues {
                compartmentAdj := compartmentPremium[compartment.Value]
                
                for _, quantity := range quantityValues {
                    discountFactor := quantityDiscount[quantity.Value]
                    
                    // Parse pack quantity to get number of trays
                    packSize := 0
                    switch quantity.Value {
                    case "Sample Pack - 10 trays":
                        packSize = 10
                    case "Home Pack - 25 trays":
                        packSize = 25
                    case "Small Business Pack - 50 trays":
                        packSize = 50
                    case "Restaurant Starter - 100 trays":
                        packSize = 100
                    case "Catering Pack - 250 trays":
                        packSize = 250
                    case "Bulk Wholesale - 500 trays":
                        packSize = 500
                    case "Institutional (School Program) - 1000 trays":
                        packSize = 1000
                    case "Pallet - 5000 trays (Container Load)":
                        packSize = 5000
                    }
                    
                    // Calculate price per tray with compartment adjustment
                    adjustedPricePerTray := pricePerTray + compartmentAdj
                    if adjustedPricePerTray < 1000 {
                        adjustedPricePerTray = 1000 // Floor price
                    }
                    
                    // Calculate pack price with quantity discount
                    packPrice := adjustedPricePerTray * float64(packSize) * discountFactor
                    
                    // Round to nearest 1000 Rupiah for packs > 100, otherwise nearest 500
                    var finalPrice float64
                    if packSize >= 100 {
                        finalPrice = float64(int(packPrice/1000)) * 1000
                    } else {
                        finalPrice = float64(int(packPrice/500)) * 500
                    }
                    
                    // Generate SKU code
                    skuCode := fmt.Sprintf("TRAY-%s-%s-%s-%s-%d", 
                        abbreviateMaterial(material.Value),
                        abbreviateCompartment(compartment.Value),
                        abbreviateSize(size.Value),
                        abbreviateQuantity(quantity.Value),
                        rand.Intn(1000))
                    
                    // Calculate sale price (25% chance of being on sale)
                    var salePrice *float64
                    if rand.Float64() < 0.25 {
                        discount := 0.85 + (rand.Float64() * 0.10) // 5-15% off
                        sp := finalPrice * discount
                        if packSize >= 100 {
                            sp = float64(int(sp/1000)) * 1000
                        } else {
                            sp = float64(int(sp/500)) * 500
                        }
                        salePrice = &sp
                    }
                    
                    // Determine stock levels based on pack quantity
                    stock := 0
                    minStock := minStockByPack[quantity.Value]
                    
                    // Higher stock for popular combinations
                    if material.Value == "Bagasse (Sugarcane Fiber) - Standard" && 
                       compartment.Value == "3-Compartment - Rice + Protein + Vegetable" && 
                       size.Value == "Standard - 9.5x7.5 inches (Adult Meal)" {
                        // Most popular configuration
                        stock = minStock * (15 + rand.Intn(20))
                    } else if quantity.Value == "Institutional (School Program) - 1000 trays" {
                        // Government program volume
                        stock = 5 + rand.Intn(15) // 5-20 pallets worth
                    } else if packSize <= 50 {
                        stock = minStock * (10 + rand.Intn(30))
                    } else {
                        stock = minStock * (5 + rand.Intn(15))
                    }
                    
                    // Calculate weight (varies by pack size)
                    // Each tray approx 40-60g depending on size/material
                    weightPerTray := 50.0 // average grams
                    if size.Value == "Large - 11x8.5 inches (Hearty Meal)" {
                        weightPerTray = 70.0
                    } else if size.Value == "Mini - 6x4.5 inches (Appetizer/Sample)" {
                        weightPerTray = 30.0
                    }
                    
                    // Bamboo is slightly heavier
                    if material.Value == "Bamboo Fiber - Premium" {
                        weightPerTray *= 1.2
                    }
                    
                    totalWeight := weightPerTray * float64(packSize)
                    
                    sku := entity.SKU{
                        ProductID: product.ID,
                        SKUCode:   skuCode,
                        Price:     finalPrice,
                        SalePrice: salePrice,
                        Stock:     stock,
                        MinStock:  minStock,
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
                        VariantValueID: material.ID,
                    }
                    db.Create(&skuVariant1)
                    
                    skuVariant2 := entity.SKUVariantValue{
                        SKUID:          sku.ID,
                        VariantValueID: compartment.ID,
                    }
                    db.Create(&skuVariant2)
                    
                    skuVariant3 := entity.SKUVariantValue{
                        SKUID:          sku.ID,
                        VariantValueID: size.ID,
                    }
                    db.Create(&skuVariant3)
                    
                    skuVariant4 := entity.SKUVariantValue{
                        SKUID:          sku.ID,
                        VariantValueID: quantity.ID,
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
			"https://ecolipak.com/cdn/shop/files/14-in-disposable-cutting-resistant-compostable-traysbpi-certified-pfas-free-4713263.jpg?v=1762583270&width=1024",
			"https://ecolipak.com/cdn/shop/files/14-in-disposable-cutting-resistant-compostable-traysbpi-certified-pfas-free-6819939.jpg?v=1762583269&width=1024",
			"https://ecolipak.com/cdn/shop/files/14-in-disposable-cutting-resistant-compostable-traysbpi-certified-pfas-free-7092684.jpg?v=1762583270&width=1024",
		}

		for _, url := range imageURLs {
			img := entity.Image{
				ProductID: product.ID,
				ImageURL: url,
			}

			db.Create(&img)
		}

    fmt.Printf("✅ Successfully seeded Compostable Cutting-Resistant Trays product with %d SKU variants\n", len(allSKUs))
    fmt.Printf("   Price range: Rp %.0f - Rp %.0f\n", minPrice, maxPrice)
    
    return nil
}

// Helper functions for SKU code generation
func abbreviateMaterial(material string) string {
    switch {
    case material == "Bagasse (Sugarcane Fiber) - Standard":
        return "BGS"
    case material == "Wheat Straw Fiber - Natural Light":
        return "WHT"
    case material == "Bamboo Fiber - Premium":
        return "BMB"
    case material == "Recycled Paper Pulp - Economy":
        return "RPP"
    default:
        return "MAT"
    }
}

func abbreviateCompartment(comp string) string {
    switch {
    case comp == "3-Compartment - Rice + Protein + Vegetable":
        return "C3"
    case comp == "2-Compartment - Main + Side":
        return "C2"
    case comp == "1-Compartment - Single Large (Nasi Box Style)":
        return "C1"
    case comp == "5-Compartment - Tiffin Style (Complete Meal)":
        return "C5"
    case comp == "Divided Tray with Lid Slot - For Delivery":
        return "LID"
    default:
        return "COM"
    }
}

func abbreviateSize(size string) string {
    switch {
    case size == "Standard - 9.5x7.5 inches (Adult Meal)":
        return "STD"
    case size == "Junior - 7.5x5.5 inches (Child/Snack)":
        return "JNR"
    case size == "Large - 11x8.5 inches (Hearty Meal)":
        return "LRG"
    case size == "Mini - 6x4.5 inches (Appetizer/Sample)":
        return "MIN"
    default:
        return "SIZ"
    }
}

func abbreviateQuantity(qty string) string {
    switch {
    case qty == "Sample Pack - 10 trays":
        return "P10"
    case qty == "Home Pack - 25 trays":
        return "P25"
    case qty == "Small Business Pack - 50 trays":
        return "P50"
    case qty == "Restaurant Starter - 100 trays":
        return "P100"
    case qty == "Catering Pack - 250 trays":
        return "P250"
    case qty == "Bulk Wholesale - 500 trays":
        return "P500"
    case qty == "Institutional (School Program) - 1000 trays":
        return "P1K"
    case qty == "Pallet - 5000 trays (Container Load)":
        return "P5K"
    default:
        return "PKG"
    }
}