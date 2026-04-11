package seeder

import (
	"debian-ecommerce/internal/data/entity"
	"fmt"
	"math/rand"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func SeedZeroWasteStarterKit(db *gorm.DB) error {
    rand.Seed(time.Now().UnixNano())
    
    // First, ensure category exists (category ID 1 for Zero Waste Kit)
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

    // Create the main product - Zero Waste Starter Kit
    product := entity.Product{
        CategoryID:       1,
        Name:             "Zero Waste Starter Kit - Complete Set for Plastic-Free Living",
        Description:      `Start your journey toward a plastic-free lifestyle with our comprehensive **Zero Waste Starter Kit**. Inspired by successful Indonesian initiatives like **Siklus**, which provides cost-effective refillable household products without plastic waste, this kit contains everything you need to replace single-use items in your daily routine .

**🌱 WHY GO ZERO WASTE IN INDONESIA?**

According to environmental data, 70% of Indonesians purchase household goods in single-serving sachets because they cannot afford to buy in larger quantities, meaning low-income populations often end up paying extra for their everyday needs . This kit helps you save money while saving the environment—Siklus offers products that are on average 20% cheaper and help customers save on everyday necessities .

**🎁 WHAT'S INCLUDED IN YOUR STARTER KIT:**

Our complete kit includes 8 essential items carefully selected for the Indonesian household:

| Item | Quantity | Description |
|------|----------|-------------|
| Stainless Steel Straw Set | 4 straws + brush | 2 straight + 2 bent, 6mm & 8mm diameters |
| Bamboo Cutlery Set | 5 pieces | Fork, knife, spoon, chopsticks, carrying case |
| Reusable Produce Bags | 3 bags | Small, medium, large mesh bags |
| Beeswax Food Wraps | 3 wraps | S, M, L sizes for food covering |
| Cotton Tote Bag | 1 bag | Heavy-duty canvas, folds into pouch |
| Stainless Steel Tumbler | 1 tumbler | 500ml insulated, keeps drinks cold 24h |
| Bamboo Toothbrush | 2 brushes | Soft bristle, biodegradable handle |
| Cleaning Brush Set | 2 brushes | For straws and small containers |

**✨ KEY FEATURES:**

- ✅ **Complete Solution**: Everything needed to replace disposables in one package
- ✅ **Premium Materials**: Food-grade stainless steel, organic bamboo, natural cotton
- ✅ **Indonesian Context**: Designed for local shopping habits—pasar, supermarket, and warung
- ✅ **Travel-Friendly**: All items pack into included cotton tote
- ✅ **Plastic-Free Packaging**: Shipped in recycled cardboard with no plastic

**🌍 SUSTAINABILITY IMPACT:**

Based on Siklus' successful model, each kit helps customers save money while reducing plastic waste . Over one year, this kit prevents approximately:

- **500+** single-use plastic straws
- **300+** plastic cutlery pieces
- **200+** plastic produce bags
- **50+** plastic water bottles
- **100+** plastic food wrap sheets

**📋 KIT CONTENTS DETAILS:**

**1. Stainless Steel Straw Set (4 pcs + Brush)**
- 2 straight (21cm) + 2 bent (21cm)
- 6mm and 8mm diameters for various drinks
- Includes cleaning brush and cotton pouch

**2. Bamboo Cutlery Set (5 pcs)**
- Fork, knife, spoon, chopsticks
- Organic bamboo, naturally antibacterial
- Comes in cotton roll-up case

**3. Reusable Produce Bags (3 pcs)**
- Small (25x30cm), Medium (30x40cm), Large (40x50cm)
- Lightweight mesh, tare weight printed
- Perfect for fruits, vegetables, and bulk items

**4. Beeswax Food Wraps (3 pcs)**
- Small (20x20cm), Medium (25x25cm), Large (30x30cm)
- Organic cotton with beeswax, jojoba oil, tree resin
- Reusable for up to 1 year

**5. Cotton Tote Bag**
- Heavy-duty canvas (140gsm)
- Folds into built-in pouch
- Dimensions: 40x45cm, handles 30cm

**6. Stainless Steel Tumbler**
- Double-wall vacuum insulation
- 500ml capacity
- Keeps cold 24h, hot 12h
- Leak-proof lid

**7. Bamboo Toothbrush (2 pcs)**
- Medium-soft bristles
- Biodegradable bamboo handle
- Compostable packaging

**8. Cleaning Brush Set (2 pcs)**
- Small brush for straws
- Large brush for bottles/containers
- Natural bristles, wooden handles

**🇮🇩 INDONESIA'S ZERO WASTE MOVEMENT**

Indonesia is at the forefront of zero-waste innovation. Social enterprises like **Siklus** have pioneered refill delivery services using bicycles, partnering with major FMCG companies including Nestle, P&G, Mars, and Sinarmas . Through community outreach, Siklus has educated over **1,600 women across 15 villages** on waste issues and invited them to use refill services .

The movement has gained significant momentum, with activities ranging from talk shows for businesspeople to outreach about plastic waste in Jakarta, Depok, and Bekasi, as well as online campaign activities for #KemanaSampahmu .

**📦 PACKAGING:**

Your Zero Waste Starter Kit arrives in:
- Recycled cardboard box
- No plastic tape or wrapping
- Compostable packing peanuts
- Includes educational booklet on zero-waste living

**🎯 PERFECT FOR:**

- **Beginners**: Starting your zero-waste journey
- **Gifts**: Eco-friendly presents for friends and family
- **Students**: Moving out and setting up sustainable home
- **Travelers**: Complete set for plastic-free travel
- **Corporate gifts**: Sustainable option for company events

**📖 INCLUDED GUIDEBOOK:**

Each kit comes with a 20-page guidebook (Bahasa Indonesia + English) covering:
- How to use each item
- Care and maintenance instructions
- Tips for zero-waste shopping in Indonesia
- DIY recipes (natural cleaners, toothpaste)
- Composting basics

Start your zero-waste journey today—join thousands of Indonesians reducing plastic waste and saving money with every refill!`,
        IsPublished:      true,
        Tags:             pq.StringArray{"zero waste kit", "plastic-free", "starter kit", "eco-friendly", "reusable", "sustainable living", "bamboo cutlery", "stainless straw", "produce bags", "beeswax wraps", "tote bag", "tumbler", "ramah lingkungan", "tanpa plastik", "zero waste starter", "gift set", "eco gift"},
        Slug:             "zero-waste-starter-kit-complete-set-plastic-free-living",
        MainImageURL:     "https://ecoclubofficial.com/wp-content/uploads/2019/03/ecozie-4.jpg",
        MainImagePublicID: "zero-waste-kit-main",
        AverageRating:    4.9,
        ReviewCount:      892,
        SoldCount:        3456,
        MinPrice:         0, // Will be updated after SKUs
        MaxPrice:         0, // Will be updated after SKUs
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }

    // Create Variant Types
    // Variant 1: Kit Size
    kitSizeType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Kit Size",
    }
    if err := db.Create(&kitSizeType).Error; err != nil {
        return err
    }

    // Variant 2: Material Preference
    materialPrefType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Material Preference",
    }
    if err := db.Create(&materialPrefType).Error; err != nil {
        return err
    }

    // Variant 3: Color Theme
    colorThemeType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Color Theme",
    }
    if err := db.Create(&colorThemeType).Error; err != nil {
        return err
    }

    // Variant 4: Pouch/Tote Design
    designType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Tote Design",
    }
    if err := db.Create(&designType).Error; err != nil {
        return err
    }

    // Variant 5: Packaging Option
    packagingType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Packaging",
    }
    if err := db.Create(&packagingType).Error; err != nil {
        return err
    }

    // Create Variant Values
    // Kit Size options
    kitSizeValues := []entity.VariantValue{
        {
            VariantTypeID: kitSizeType.ID,
            Value:         "Essential Kit - 5 items (Straws + Cutlery + Tote)",
        },
        {
            VariantTypeID: kitSizeType.ID,
            Value:         "Standard Kit - 8 items (Complete Set)",
        },
        {
            VariantTypeID: kitSizeType.ID,
            Value:         "Premium Kit - 12 items ( + Food Containers + Water Bottle)",
        },
        {
            VariantTypeID: kitSizeType.ID,
            Value:         "Family Kit - 16 items (Double quantities for 2 people)",
        },
    }
    
    for i := range kitSizeValues {
        if err := db.Create(&kitSizeValues[i]).Error; err != nil {
            return err
        }
    }

    // Material Preference
    materialPrefValues := []entity.VariantValue{
        {
            VariantTypeID: materialPrefType.ID,
            Value:         "Standard - Mixed Materials",
        },
        {
            VariantTypeID: materialPrefType.ID,
            Value:         "All Bamboo (No Metal)",
        },
        {
            VariantTypeID: materialPrefType.ID,
            Value:         "All Stainless (Premium)",
        },
        {
            VariantTypeID: materialPrefType.ID,
            Value:         "Budget-Friendly (Recycled Plastic)",
        },
    }
    
    for i := range materialPrefValues {
        if err := db.Create(&materialPrefValues[i]).Error; err != nil {
            return err
        }
    }

    // Color Theme
    colorThemeValues := []entity.VariantValue{
        {
            VariantTypeID: colorThemeType.ID,
            Value:         "Natural/Earthy (Beige + Green)",
        },
        {
            VariantTypeID: colorThemeType.ID,
            Value:         "Modern Minimalist (Black + White)",
        },
        {
            VariantTypeID: colorThemeType.ID,
            Value:         "Boho (Terracotta + Mustard)",
        },
        {
            VariantTypeID: colorThemeType.ID,
            Value:         "Pastel (Pink + Mint)",
        },
        {
            VariantTypeID: colorThemeType.ID,
            Value:         "Rainbow Mix",
        },
    }
    
    for i := range colorThemeValues {
        if err := db.Create(&colorThemeValues[i]).Error; err != nil {
            return err
        }
    }

    // Tote Design
    designValues := []entity.VariantValue{
        {
            VariantTypeID: designType.ID,
            Value:         "Plain - No Print",
        },
        {
            VariantTypeID: designType.ID,
            Value:         "#ZeroWasteIndonesia Print",
        },
        {
            VariantTypeID: designType.ID,
            Value:         "Batik Pattern - Kawung",
        },
        {
            VariantTypeID: designType.ID,
            Value:         "Batik Pattern - Parang",
        },
        {
            VariantTypeID: designType.ID,
            Value:         "Custom - Minimalist Leaf Design",
        },
    }
    
    for i := range designValues {
        if err := db.Create(&designValues[i]).Error; err != nil {
            return err
        }
    }

    // Packaging Option
    packagingValues := []entity.VariantValue{
        {
            VariantTypeID: packagingType.ID,
            Value:         "Standard - Recycled Box",
        },
        {
            VariantTypeID: packagingType.ID,
            Value:         "Gift Wrapped - Reusable Fabric Wrap (Furoshiki)",
        },
        {
            VariantTypeID: packagingType.ID,
            Value:         "Minimal - No Box (Eco Option)",
        },
    }
    
    for i := range packagingValues {
        if err := db.Create(&packagingValues[i]).Error; err != nil {
            return err
        }
    }

    // Base price mapping based on market research
    // Zero waste kits typically range from Rp 250,000 - 1,200,000 depending on contents
    
    basePriceMap := map[string]map[string]float64{
        "Essential Kit - 5 items (Straws + Cutlery + Tote)": {
            "Standard - Mixed Materials":          249000,
            "All Bamboo (No Metal)":               229000,
            "All Stainless (Premium)":              299000,
            "Budget-Friendly (Recycled Plastic)":   189000,
        },
        "Standard Kit - 8 items (Complete Set)": {
            "Standard - Mixed Materials":          449000,
            "All Bamboo (No Metal)":               399000,
            "All Stainless (Premium)":              549000,
            "Budget-Friendly (Recycled Plastic)":   349000,
        },
        "Premium Kit - 12 items ( + Food Containers + Water Bottle)": {
            "Standard - Mixed Materials":          749000,
            "All Bamboo (No Metal)":               679000,
            "All Stainless (Premium)":              899000,
        },
        "Family Kit - 16 items (Double quantities for 2 people)": {
            "Standard - Mixed Materials":          849000,
            "All Bamboo (No Metal)":               779000,
            "All Stainless (Premium)":              999000,
        },
    }

    // Color theme premium adjustments
    colorPremium := map[string]float64{
        "Natural/Earthy (Beige + Green)":       0,
        "Modern Minimalist (Black + White)":    0,
        "Boho (Terracotta + Mustard)":          15000,
        "Pastel (Pink + Mint)":                 10000,
        "Rainbow Mix":                           20000,
    }

    // Tote design premium
    designPremium := map[string]float64{
        "Plain - No Print":                      0,
        "#ZeroWasteIndonesia Print":             5000,
        "Batik Pattern - Kawung":                25000,
        "Batik Pattern - Parang":                 25000,
        "Custom - Minimalist Leaf Design":       10000,
    }

    // Packaging premium
    packagingPremium := map[string]float64{
        "Standard - Recycled Box":                0,
        "Gift Wrapped - Reusable Fabric Wrap (Furoshiki)": 35000,
        "Minimal - No Box (Eco Option)":          -10000,
    }

    var allSKUs []entity.SKU
    var minPrice float64 = 999999
    var maxPrice float64 = 0

    // Generate SKUs for combinations
    for _, kitSize := range kitSizeValues {
        sizePrices, exists := basePriceMap[kitSize.Value]
        if !exists {
            continue
        }
        
        for _, material := range materialPrefValues {
            basePrice, materialExists := sizePrices[material.Value]
            if !materialExists {
                continue
            }
            
            // Skip budget option for premium kits
            if material.Value == "Budget-Friendly (Recycled Plastic)" && 
               (kitSize.Value == "Premium Kit - 12 items ( + Food Containers + Water Bottle)" || 
                kitSize.Value == "Family Kit - 16 items (Double quantities for 2 people)") {
                continue
            }
            
            for _, color := range colorThemeValues {
                for _, design := range designValues {
                    for _, packaging := range packagingValues {
                        // Calculate final price
                        colorAdj := colorPremium[color.Value]
                        designAdj := designPremium[design.Value]
                        packagingAdj := packagingPremium[packaging.Value]
                        
                        finalPrice := basePrice + colorAdj + designAdj + packagingAdj
                        
                        // Ensure price doesn't go negative
                        if finalPrice < 150000 {
                            finalPrice = 150000
                        }
                        
                        // Round to nearest 1000 Rupiah
                        finalPrice = float64(int(finalPrice/1000)) * 1000
                        
                        // Generate SKU code
                        skuCode := fmt.Sprintf("ZWK-%s-%s-%s-%s-%s-%d", 
                            abbreviateKitSize(kitSize.Value),
                            abbreviateKitMaterial(material.Value),
                            abbreviateKitColor(color.Value),
                            abbreviateKitDesign(design.Value),
                            abbreviateKitPackaging(packaging.Value),
                            rand.Intn(1000))
                        
                        // Calculate sale price (25% chance of being on sale)
                        var salePrice *float64
                        if rand.Float64() < 0.25 {
                            discount := 0.80 + (rand.Float64() * 0.15) // 5-20% off
                            sp := finalPrice * discount
                            sp = float64(int(sp/1000)) * 1000
                            salePrice = &sp
                        }
                        
                        // Determine stock levels based on popularity
                        stock := 0
                        if kitSize.Value == "Standard Kit - 8 items (Complete Set)" && 
                           material.Value == "Standard - Mixed Materials" {
                            stock = 100 + rand.Intn(150) // Most popular: 100-250
                        } else if kitSize.Value == "Essential Kit - 5 items (Straws + Cutlery + Tote)" {
                            stock = 80 + rand.Intn(120) // 80-200
                        } else if kitSize.Value == "Premium Kit - 12 items ( + Food Containers + Water Bottle)" {
                            stock = 30 + rand.Intn(50) // 30-80
                        } else if kitSize.Value == "Family Kit - 16 items (Double quantities for 2 people)" {
                            stock = 20 + rand.Intn(40) // 20-60
                        } else {
                            stock = 40 + rand.Intn(80) // 40-120
                        }
                        
                        // Adjust stock for premium variants
                        if colorAdj > 15000 || designAdj > 20000 {
                            stock = stock / 2
                        }
                        
                        // Calculate weight based on kit size
                        weight := 0.0
                        switch kitSize.Value {
                        case "Essential Kit - 5 items (Straws + Cutlery + Tote)":
                            weight = 600
                        case "Standard Kit - 8 items (Complete Set)":
                            weight = 1200
                        case "Premium Kit - 12 items ( + Food Containers + Water Bottle)":
                            weight = 2200
                        case "Family Kit - 16 items (Double quantities for 2 people)":
                            weight = 2500
                        }
                        
                        sku := entity.SKU{
                            ProductID: product.ID,
                            SKUCode:   skuCode,
                            Price:     finalPrice,
                            SalePrice: salePrice,
                            Stock:     stock,
                            MinStock:  10,
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
                            VariantValueID: kitSize.ID,
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
                        
                        skuVariant4 := entity.SKUVariantValue{
                            SKUID:          sku.ID,
                            VariantValueID: design.ID,
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
			"https://memotherearthbrand.com/cdn/shop/files/7-Zero_Waste_Kitchen_Kit5_1_1296x.jpg?v=1762457095",
			"https://memotherearthbrand.com/cdn/shop/files/DSC_2400-Copy_900x.jpg?v=1762457890",
			"https://organics.com/cdn/shop/products/LQ3ph7CgGp.jpg?v=1654117954&width=2048",
		}

		for _, url := range imageURLs {
			img := entity.Image{
				ProductID: product.ID,
				ImageURL: url,
			}

			db.Create(&img)
		}

    fmt.Printf("✅ Successfully seeded Zero Waste Starter Kit product with %d SKU variants\n", len(allSKUs))
    fmt.Printf("   Price range: Rp %.0f - Rp %.0f\n", minPrice, maxPrice)
    
    return nil
}

// Helper functions for SKU code generation
func abbreviateKitSize(size string) string {
    switch {
    case size == "Essential Kit - 5 items (Straws + Cutlery + Tote)":
        return "ESS"
    case size == "Standard Kit - 8 items (Complete Set)":
        return "STD"
    case size == "Premium Kit - 12 items ( + Food Containers + Water Bottle)":
        return "PRM"
    case size == "Family Kit - 16 items (Double quantities for 2 people)":
        return "FAM"
    default:
        return "SIZ"
    }
}

func abbreviateKitMaterial(material string) string {
    switch {
    case material == "Standard - Mixed Materials":
        return "MIX"
    case material == "All Bamboo (No Metal)":
        return "BAM"
    case material == "All Stainless (Premium)":
        return "STL"
    case material == "Budget-Friendly (Recycled Plastic)":
        return "BUD"
    default:
        return "MAT"
    }
}

func abbreviateKitColor(color string) string {
    switch {
    case color == "Natural/Earthy (Beige + Green)":
        return "NAT"
    case color == "Modern Minimalist (Black + White)":
        return "MOD"
    case color == "Boho (Terracotta + Mustard)":
        return "BOH"
    case color == "Pastel (Pink + Mint)":
        return "PAS"
    case color == "Rainbow Mix":
        return "RNB"
    default:
        return "CLR"
    }
}

func abbreviateKitDesign(design string) string {
    switch {
    case design == "Plain - No Print":
        return "PLN"
    case design == "#ZeroWasteIndonesia Print":
        return "ZWI"
    case design == "Batik Pattern - Kawung":
        return "BTK1"
    case design == "Batik Pattern - Parang":
        return "BTK2"
    case design == "Custom - Minimalist Leaf Design":
        return "LEAF"
    default:
        return "DSN"
    }
}

func abbreviateKitPackaging(pkg string) string {
    switch {
    case pkg == "Standard - Recycled Box":
        return "BOX"
    case pkg == "Gift Wrapped - Reusable Fabric Wrap (Furoshiki)":
        return "GFT"
    case pkg == "Minimal - No Box (Eco Option)":
        return "MIN"
    default:
        return "PKG"
    }
}