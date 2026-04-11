package seeder

import (
	"debian-ecommerce/internal/data/entity"
	"fmt"
	"math/rand"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func SeedHomeComposter(db *gorm.DB) error {
    rand.Seed(time.Now().UnixNano())
    
    // First, ensure category exists (category ID 3 for Home & Cleaning)
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

    // Create the main product - Home Composter
    product := entity.Product{
        CategoryID:       3,
        Name:             "Home Composter Kit - Ubah Sampah Dapur Jadi Pupuk Organik",
        Description:      `Transform your kitchen waste into valuable organic fertilizer with our **Home Composter Kit**. As Indonesia faces growing waste management challenges—with Tangerang Selatan recently declaring a waste emergency and residents taking matters into their own hands by creating composters from water drums [citation:1]—home composting has become an essential practice for environmentally conscious households.

**🌱 WHY COMPOST AT HOME?**

Indonesia generates millions of tons of organic waste annually, much of which ends up in landfills or is burned, creating pollution and health problems [citation:1][citation:2]. By composting at home, you can:
- **Reduce waste volume**: Up to 60% of household waste is organic and can be composted
- **Create free fertilizer**: Produce nutrient-rich compost for your plants and garden
- **Save money**: Reduce spending on chemical fertilizers and garbage bags
- **Help the environment**: Prevent methane production in landfills and reduce carbon footprint

**✨ PRODUCT OVERVIEW**

Our Home Composter Kit is designed for Indonesian households, with simple operation and maintenance. Based on successful community initiatives in Serua, Ciputat, where residents created 40 composter units at Rp 450,000 each [citation:1], and UGM student programs in Rembang teaching residents to make composters from used buckets [citation:4], we offer ready-to-use systems for immediate waste reduction.

**🧪 HOW IT WORKS:**

1. **Add organic waste**: Fruit peels, vegetable scraps, coffee grounds, eggshells, leaves
2. **Add bio-activator**: Use EM4, molasses (or brown sugar/gula merah), and water [citation:4]
3. **Wait 2-4 weeks**: Stir every 3 days to aerate
4. **Harvest liquid fertilizer**: Collect from built-in tap (pupuk cair)
5. **Harvest solid compost**: After 1-2 months, solid compost is ready for your garden

**📋 KEY FEATURES:**

- ✅ **Complete System**: Includes composter vessel, tap for liquid fertilizer collection, aeration lid, and instruction manual
- ✅ **Odorless Design**: Proper aeration prevents bad smells—no more complaints from neighbors!
- ✅ **Dual Output**: Produces both liquid fertilizer (POC) and solid compost
- ✅ **Easy Harvest**: Built-in tap for convenient liquid collection [citation:4]
- ✅ **Durable Construction**: Made from recycled HDPE plastic or stainless steel options
- ✅ **Rat-Proof**: Secure lid prevents pests from accessing waste

**🇮🇩 INDONESIA'S COMPOSTING MOVEMENT**

Communities across Indonesia are embracing composting:

- **Serua, Tangerang Selatan**: Residents created 40 composter units from water drums to combat the waste crisis, with each unit costing Rp 450,000 [citation:1]
- **Johogunung, Rembang**: UGM students taught residents to make simple composters, helping farmers reduce dependence on chemical fertilizers [citation:2]
- **Banyuurip, Rembang**: Using two used buckets, residents learned to process organic waste into liquid fertilizer within 2-4 weeks [citation:4][citation:6]
- **Bokoharjo, Sleman**: Modified plastic drums with aeration holes and taps serve as community composting systems [citation:5]
- **Tlogotirto, Grobogan**: UNDIP students developed anaerobic RE-DRUM composters producing organic liquid fertilizer [citation:9]

**🧴 TYPES OF COMPOSTERS AVAILABLE:**

1. **Aerobic Composter**: Traditional design with air circulation—faster decomposition, no methane
2. **Anaerob RE-DRUM**: Fermentation-based, sealed system producing high-quality liquid fertilizer [citation:9]
3. **Takakura Method**: Japanese-style basket system for small spaces
4. **Bokashi Bucket**: Fermentation with effective microorganisms
5. **Smart Composter**: With thermal sensors and automatic ventilation (for tech-savvy users) [citation:7]

**🌿 INCLUDED IN YOUR KIT:**

- Main composter vessel (20L, 60L, 120L, or 240L options)
- Liquid collection tap with hose connector
- Aeration lid with carbon filter
- Bio-activator starter pack (EM4 + molasses)
- Compost thermometer
- Step-by-step guidebook in Bahasa Indonesia
- 1-year warranty

**📊 SPECIFICATIONS:**

| Model | Capacity | Daily Waste | Household Size | Dimensions | Output/Month |
|-------|----------|-------------|----------------|------------|--------------|
| Urban 20L | 20 liters | 1-2 kg | 1-2 people | 40x40x50 cm | 5L liquid + 8kg solid |
| Family 60L | 60 liters | 3-5 kg | 3-4 people | 50x50x70 cm | 15L liquid + 25kg solid |
| Community 120L | 120 liters | 6-8 kg | 5-6 people | 60x60x90 cm | 30L liquid + 50kg solid |
| Master 240L | 240 liters | 10-15 kg | Small community | 80x80x100 cm | 60L liquid + 100kg solid |

**🎯 PERFECT FOR:**

- Urban households with small gardens or potted plants
- Suburban homes with yards and fruit trees
- Community composting programs (RT/RW level)
- School environmental education programs
- Restaurants and cafes wanting to reduce waste

**🌍 SUSTAINABILITY IMPACT:**

According to market research, the Indonesia residential organic compost market is growing significantly, with import momentum increasing 31.92% in 2024 due to rising sustainability awareness [citation:3]. Government initiatives promoting composting and waste reduction at the residential level are driving demand [citation:3]. Each composter:
- Diverts 200-500 kg of organic waste from landfills annually
- Produces enough fertilizer for a 50m² garden
- Reduces methane emissions equivalent to planting 5 trees per year
- Saves families Rp 500,000-1,000,000 annually on fertilizers

**🛠️ CARE & MAINTENANCE:**

1. **Empty liquid fertilizer** every 3-5 days via tap
2. **Stir contents** every 3 days to aerate (for aerobic entity) [citation:4]
3. **Add bio-activator** monthly to maintain microbial activity
4. **Clean tap** periodically to prevent clogging
5. **Harvest solid compost** every 1-2 months

**📝 PRO TIPS:**

- Maintain brown-to-green ratio (carbon:nitrogen = 30:1)
- Chop large pieces for faster decomposition
- Add shredded paper/cardboard if too wet
- Avoid meat, dairy, and oily foods
- Use rice washing water (air cucian beras) as natural activator [citation:6]

Join thousands of Indonesian households turning waste into resources. Make composting a daily habit—for your garden and our planet!`,
        IsPublished:      true,
        Tags:             pq.StringArray{"komposter", "composter", "pupuk organik", "organic fertilizer", "pupuk cair", "liquid fertilizer", "sampah organik", "organic waste", "zero waste", "rumah tangga", "home composting", "eco-friendly", "berkebun", "gardening", "EM4", "molase", "takakura", "bokashi", "RE-DRUM", "lingkungan", "waste management"},
        Slug:             "home-composter-kit-ubah-sampah-dapur-jadi-pupuk-organik",
        MainImageURL:     "https://jubelio-store.s3.ap-southeast-1.amazonaws.com/sustaination/2021/03/01214020/Komposter-Set-Sekop-3-scaled.webp",
        MainImagePublicID: "composter-main",
        AverageRating:    4.7,
        ReviewCount:      345,
        SoldCount:        1234,
        MinPrice:         0, // Will be updated after SKUs
        MaxPrice:         0, // Will be updated after SKUs
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }

    // Create Variant Types
    // Variant 1: Composter Type / Method
    methodType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Composting Method",
    }
    if err := db.Create(&methodType).Error; err != nil {
        return err
    }

    // Variant 2: Size / Capacity
    capacityType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Capacity",
    }
    if err := db.Create(&capacityType).Error; err != nil {
        return err
    }

    // Variant 3: Material
    materialType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Material",
    }
    if err := db.Create(&materialType).Error; err != nil {
        return err
    }

    // Variant 4: Feature Level
    featureType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Features",
    }
    if err := db.Create(&featureType).Error; err != nil {
        return err
    }

    // Variant 5: Starter Kit Option
    starterType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Includes Activator",
    }
    if err := db.Create(&starterType).Error; err != nil {
        return err
    }

    // Create Variant Values
    // Composting Method
    methodValues := []entity.VariantValue{
        {
            VariantTypeID: methodType.ID,
            Value:         "Aerobic - Traditional (Udara)",
        },
        {
            VariantTypeID: methodType.ID,
            Value:         "Anaerob RE-DRUM - Fermentasi Tertutup",
        },
        {
            VariantTypeID: methodType.ID,
            Value:         "Takakura - Keranjang Jepang",
        },
        {
            VariantTypeID: methodType.ID,
            Value:         "Bokashi - Fermentasi Cepat",
        },
        {
            VariantTypeID: methodType.ID,
            Value:         "Smart Composter - Otomatis",
        },
    }
    
    for i := range methodValues {
        if err := db.Create(&methodValues[i]).Error; err != nil {
            return err
        }
    }

    // Capacity options (based on market research)
    capacityValues := []entity.VariantValue{
        {
            VariantTypeID: capacityType.ID,
            Value:         "Urban 20L - 1-2 Orang / Apartemen",
        },
        {
            VariantTypeID: capacityType.ID,
            Value:         "Family 60L - 3-4 Orang / Rumah Kecil",
        },
        {
            VariantTypeID: capacityType.ID,
            Value:         "Community 120L - 5-6 Orang / Rumah Besar",
        },
        {
            VariantTypeID: capacityType.ID,
            Value:         "Master 240L - Komunitas / RT",
        },
    }
    
    for i := range capacityValues {
        if err := db.Create(&capacityValues[i]).Error; err != nil {
            return err
        }
    }

    // Material options
    materialValues := []entity.VariantValue{
        {
            VariantTypeID: materialType.ID,
            Value:         "Recycled HDPE Plastik - Ekonomis",
        },
        {
            VariantTypeID: materialType.ID,
            Value:         "Stainless Steel - Premium",
        },
        {
            VariantTypeID: materialType.ID,
            Value:         "Galvanized Steel - Heavy Duty",
        },
        {
            VariantTypeID: materialType.ID,
            Value:         "Bamboo & Wood - Natural",
        },
        {
            VariantTypeID: materialType.ID,
            Value:         "DIY Kit - Ember Bekas (Rakit Sendiri)",
        },
    }
    
    for i := range materialValues {
        if err := db.Create(&materialValues[i]).Error; err != nil {
            return err
        }
    }

    // Feature Level
    featureValues := []entity.VariantValue{
        {
            VariantTypeID: featureType.ID,
            Value:         "Standard - Kran + Tutup Aerasi",
        },
        {
            VariantTypeID: featureType.ID,
            Value:         "Deluxe - Kran + Filter Karbon + Termometer",
        },
        {
            VariantTypeID: featureType.ID,
            Value:         "Premium - Sensor Suhu + Kipas Otomatis",
        },
        {
            VariantTypeID: featureType.ID,
            Value:         "Smart - WiFi Monitoring + App Control",
        },
    }
    
    for i := range featureValues {
        if err := db.Create(&featureValues[i]).Error; err != nil {
            return err
        }
    }

    // Starter Kit Option
    starterValues := []entity.VariantValue{
        {
            VariantTypeID: starterType.ID,
            Value:         "Komposter Only (Tanpa Aktivator)",
        },
        {
            VariantTypeID: starterType.ID,
            Value:         "Starter Kit - EM4 + Molase + Panduan",
        },
        {
            VariantTypeID: starterType.ID,
            Value:         "Complete Kit - + Sekam + Arang + Cacing",
        },
    }
    
    for i := range starterValues {
        if err := db.Create(&starterValues[i]).Error; err != nil {
            return err
        }
    }

    // Base price mapping based on market research
    // Source: Community initiatives show Rp 450,000 for DIY water drum composter [citation:1]
    // Commercial composters range from Rp 200,000 - 3,000,000 depending on size and features
    
    basePriceMap := map[string]map[string]float64{
        "Urban 20L - 1-2 Orang / Apartemen": {
            "Aerobic - Traditional (Udara)":                 350000,
            "Anaerob RE-DRUM - Fermentasi Tertutup":         425000,
            "Takakura - Keranjang Jepang":                   275000,
            "Bokashi - Fermentasi Cepat":                    325000,
            "Smart Composter - Otomatis":                     1850000,
        },
        "Family 60L - 3-4 Orang / Rumah Kecil": {
            "Aerobic - Traditional (Udara)":                 550000,
            "Anaerob RE-DRUM - Fermentasi Tertutup":         650000,
            "Takakura - Keranjang Jepang":                   425000,
            "Bokashi - Fermentasi Cepat":                     495000,
            "Smart Composter - Otomatis":                     2250000,
        },
        "Community 120L - 5-6 Orang / Rumah Besar": {
            "Aerobic - Traditional (Udara)":                 850000,
            "Anaerob RE-DRUM - Fermentasi Tertutup":         995000,
            "Takakura - Keranjang Jepang":                   650000,
            "Bokashi - Fermentasi Cepat":                     750000,
            "Smart Composter - Otomatis":                     2850000,
        },
        "Master 240L - Komunitas / RT": {
            "Aerobic - Traditional (Udara)":                 1450000,
            "Anaerob RE-DRUM - Fermentasi Tertutup":         1650000,
            "Bokashi - Fermentasi Cepat":                     1250000,
            "Smart Composter - Otomatis":                     3950000,
        },
    }

    // Material premium adjustments
    materialPremium := map[string]float64{
        "Recycled HDPE Plastik - Ekonomis":       0,
        "Stainless Steel - Premium":              450000,
        "Galvanized Steel - Heavy Duty":          350000,
        "Bamboo & Wood - Natural":                250000,
        "DIY Kit - Ember Bekas (Rakit Sendiri)":  -100000, // Discount for DIY
    }

    // Feature level premium
    featurePremium := map[string]float64{
        "Standard - Kran + Tutup Aerasi":          0,
        "Deluxe - Kran + Filter Karbon + Termometer":     150000,
        "Premium - Sensor Suhu + Kipas Otomatis":         450000,
        "Smart - WiFi Monitoring + App Control":          1250000,
    }

    // Starter kit premium
    starterPremium := map[string]float64{
        "Komposter Only (Tanpa Aktivator)":        0,
        "Starter Kit - EM4 + Molase + Panduan":    75000,
        "Complete Kit - + Sekam + Arang + Cacing": 150000,
    }

    var allSKUs []entity.SKU
    var minPrice float64 = 999999
    var maxPrice float64 = 0

    // Generate SKUs for combinations
    for _, capacity := range capacityValues {
        capacityPrices, exists := basePriceMap[capacity.Value]
        if !exists {
            continue
        }
        
        for _, method := range methodValues {
            basePrice, methodExists := capacityPrices[method.Value]
            if !methodExists {
                continue
            }
            
            // Skip incompatible method-capacity combinations
            if method.Value == "Takakura - Keranjang Jepang" && capacity.Value == "Master 240L - Komunitas / RT" {
                continue // Takakura typically for smaller capacities
            }
            
            for _, material := range materialValues {
                // Skip incompatible material-method combinations
                if material.Value == "Bamboo & Wood - Natural" && method.Value == "Smart Composter - Otomatis" {
                    continue // Smart features need durable materials
                }
                
                if material.Value == "DIY Kit - Ember Bekas (Rakit Sendiri)" && method.Value == "Smart Composter - Otomatis" {
                    continue // DIY not compatible with smart features
                }
                
                for _, feature := range featureValues {
                    // Skip incompatible feature-method combinations
                    if feature.Value == "Smart - WiFi Monitoring + App Control" && method.Value != "Smart Composter - Otomatis" {
                        continue // Smart features only for smart composters
                    }
                    
                    if feature.Value == "Premium - Sensor Suhu + Kipas Otomatis" && method.Value != "Smart Composter - Otomatis" {
                        // Premium sensors only for smart entity
                        continue
                    }
                    
                    for _, starter := range starterValues {
                        // Calculate final price
                        materialAdj := materialPremium[material.Value]
                        featureAdj := featurePremium[feature.Value]
                        starterAdj := starterPremium[starter.Value]
                        
                        finalPrice := basePrice + materialAdj + featureAdj + starterAdj
                        
                        // Ensure price doesn't go negative
                        if finalPrice < 100000 {
                            finalPrice = 100000
                        }
                        
                        // Round to nearest 10,000 Rupiah for large items
                        finalPrice = float64(int(finalPrice/10000)) * 10000
                        
                        // Generate SKU code
                        skuCode := fmt.Sprintf("CMP-%s-%s-%s-%s-%s-%d", 
                            abbreviateMethod(method.Value),
                            abbreviateCapacity(capacity.Value),
                            abbreviateComposterMaterial(material.Value),
                            abbreviateFeature(feature.Value),
                            abbreviateStarter(starter.Value),
                            rand.Intn(1000))
                        
                        // Calculate sale price (20% chance of being on sale)
                        var salePrice *float64
                        if rand.Float64() < 0.2 {
                            discount := 0.85 + (rand.Float64() * 0.10) // 5-15% off
                            sp := finalPrice * discount
                            sp = float64(int(sp/10000)) * 10000
                            salePrice = &sp
                        }
                        
                        // Determine stock levels based on popularity
                        stock := 0
                        if capacity.Value == "Family 60L - 3-4 Orang / Rumah Kecil" && 
                           method.Value == "Aerobic - Traditional (Udara)" {
                            stock = 50 + rand.Intn(100) // Most popular: 50-150
                        } else if capacity.Value == "Urban 20L - 1-2 Orang / Apartemen" {
                            stock = 40 + rand.Intn(80) // 40-120
                        } else if method.Value == "Smart Composter - Otomatis" {
                            stock = 5 + rand.Intn(15) // 5-20 (niche)
                        } else if capacity.Value == "Master 240L - Komunitas / RT" {
                            stock = 10 + rand.Intn(20) // 10-30
                        } else {
                            stock = 20 + rand.Intn(50) // 20-70
                        }
                        
                        // Adjust stock for premium variants
                        if materialAdj > 300000 || featureAdj > 500000 {
                            stock = stock / 2
                        }
                        
                        // Calculate weight
                        weight := 0.0
                        switch capacity.Value {
                        case "Urban 20L - 1-2 Orang / Apartemen":
                            weight = 3000 + materialAdj/100
                        case "Family 60L - 3-4 Orang / Rumah Kecil":
                            weight = 5500 + materialAdj/100
                        case "Community 120L - 5-6 Orang / Rumah Besar":
                            weight = 9000 + materialAdj/100
                        case "Master 240L - Komunitas / RT":
                            weight = 15000 + materialAdj/100
                        }
                        
                        sku := entity.SKU{
                            ProductID: product.ID,
                            SKUCode:   skuCode,
                            Price:     finalPrice,
                            SalePrice: salePrice,
                            Stock:     stock,
                            MinStock:  5,
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
                            VariantValueID: method.ID,
                        }
                        db.Create(&skuVariant1)
                        
                        skuVariant2 := entity.SKUVariantValue{
                            SKUID:          sku.ID,
                            VariantValueID: capacity.ID,
                        }
                        db.Create(&skuVariant2)
                        
                        skuVariant3 := entity.SKUVariantValue{
                            SKUID:          sku.ID,
                            VariantValueID: material.ID,
                        }
                        db.Create(&skuVariant3)
                        
                        skuVariant4 := entity.SKUVariantValue{
                            SKUID:          sku.ID,
                            VariantValueID: feature.ID,
                        }
                        db.Create(&skuVariant4)
                        
                        skuVariant5 := entity.SKUVariantValue{
                            SKUID:          sku.ID,
                            VariantValueID: starter.ID,
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
			"https://jubelio-store.s3.ap-southeast-1.amazonaws.com/sustaination/2021/02/31233615/Komposter-Kit-GMI-1-scaled.webp",
			"https://urbankomposter.com/wp-content/uploads/2020/04/urban-komposter-17-liter.png",
			"https://www.scdprobiotics.com/cdn/shop/products/all-seasons-indoor-composter.jpg?v=1687467692",
		}

		for _, url := range imageURLs {
			img := entity.Image{
				ProductID: product.ID,
				ImageURL: url,
			}

			db.Create(&img)
		}

    fmt.Printf("✅ Successfully seeded Home Composter product with %d SKU variants\n", len(allSKUs))
    fmt.Printf("   Price range: Rp %.0f - Rp %.0f\n", minPrice, maxPrice)
    
    return nil
}

// Helper functions for SKU code generation
func abbreviateMethod(method string) string {
    switch {
    case method == "Aerobic - Traditional (Udara)":
        return "AER"
    case method == "Anaerob RE-DRUM - Fermentasi Tertutup":
        return "ANA"
    case method == "Takakura - Keranjang Jepang":
        return "TKR"
    case method == "Bokashi - Fermentasi Cepat":
        return "BOK"
    case method == "Smart Composter - Otomatis":
        return "SMT"
    default:
        return "MET"
    }
}

func abbreviateCapacity(capacity string) string {
    switch {
    case capacity == "Urban 20L - 1-2 Orang / Apartemen":
        return "U20"
    case capacity == "Family 60L - 3-4 Orang / Rumah Kecil":
        return "F60"
    case capacity == "Community 120L - 5-6 Orang / Rumah Besar":
        return "C120"
    case capacity == "Master 240L - Komunitas / RT":
        return "M240"
    default:
        return "CAP"
    }
}

func abbreviateComposterMaterial(material string) string {
    switch {
    case material == "Recycled HDPE Plastik - Ekonomis":
        return "HDPE"
    case material == "Stainless Steel - Premium":
        return "SS"
    case material == "Galvanized Steel - Heavy Duty":
        return "GALV"
    case material == "Bamboo & Wood - Natural":
        return "BAM"
    case material == "DIY Kit - Ember Bekas (Rakit Sendiri)":
        return "DIY"
    default:
        return "MAT"
    }
}

func abbreviateFeature(feature string) string {
    switch {
    case feature == "Standard - Kran + Tutup Aerasi":
        return "STD"
    case feature == "Deluxe - Kran + Filter Karbon + Termometer":
        return "DLX"
    case feature == "Premium - Sensor Suhu + Kipas Otomatis":
        return "PRM"
    case feature == "Smart - WiFi Monitoring + App Control":
        return "SMT"
    default:
        return "FTR"
    }
}

func abbreviateStarter(starter string) string {
    switch {
    case starter == "Komposter Only (Tanpa Aktivator)":
        return "ONLY"
    case starter == "Starter Kit - EM4 + Molase + Panduan":
        return "STR"
    case starter == "Complete Kit - + Sekam + Arang + Cacing":
        return "CMP"
    default:
        return "STK"
    }
}