package seeder

import (
	"debian-ecommerce/internal/data/entity"
	"fmt"
	"math/rand"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func SeedPalmLeafDinnerwareSet(db *gorm.DB) error {
    rand.Seed(time.Now().UnixNano())
    
    // First, ensure category exists (assuming category ID 6 is Straw & Cutlery, or create new)
    // For this product, let's use category ID 6 (Straw & Cutlery) which fits well
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

    // Create the main product - Palm Leaf Plate & Cutlery Set
    product := entity.Product{
        CategoryID:       6,
        Name:             "Heavy-Duty Compostable Palm Leaf Plate & Cutlery Set - Complete Eco-Friendly Dinnerware",
        Description:      `Elevate your dining experience while protecting the planet with our premium **Palm Leaf Dinnerware Set**. Made from naturally fallen areca palm leaves, this complete tableware solution offers the perfect blend of elegance, durability, and environmental responsibility—ideal for parties, weddings, catering events, corporate functions, and daily use.

**🌿 WHAT ARE PALM LEAF PLATES?**

Palm leaf plates are crafted from the sheath of the areca palm tree—the protective covering that falls naturally to the ground each month [citation:2]. In the past, these fallen sheaths were often burned by farmers, causing pollution and health problems in rural Indonesian communities [citation:2]. Today, innovative Indonesian social enterprises like **Greenie** and **Plépah** transform this agricultural waste into beautiful, functional tableware [citation:2][citation:7].

Unlike paper or plastic disposables, palm leaf products are:
- **100% natural** - No chemicals, waxes, dyes, or additives [citation:5]
- **Biodegradable & compostable** - Decomposes within 60-180 days [citation:7][citation:10]
- **Sturdy & durable** - Stronger than paper plates, holds hot and cold foods without leaking [citation:3][citation:8]
- **Microwave & oven safe** - Can withstand heat up to 220°F (104°C) [citation:10]

**🌱 INDONESIAN INNOVATION & SOCIAL IMPACT**

Indonesia is a global leader in palm leaf product innovation. Social enterprises across Java and Sumatra are empowering local farmers, particularly women in rural communities, by providing alternative income through collecting and processing fallen areca palm sheaths [citation:2][citation:7].

**Plépah**, a UNDP-supported Indonesian social enterprise, uses a decentralized "Micro Manufacturing" concept that provides communities with affordable technology to process palm sheaths into quality products [citation:7]. This approach:
- Creates economic security in villages
- Promotes sustainable business schemes
- Reduces plastic waste by replacing single-use packaging [citation:7]

**Greenie**, another Indonesian social enterprise partnered with IKEA, collects up to **600 tons of areca leaf sheath per week** from 25 local communities in Jambi and South Tangerang [citation:2]. They convert this waste into tableware and furniture, giving farmers, waste collectors, and artisans a decent income [citation:2].

By choosing our Palm Leaf Dinnerware Set, you're not just buying eco-friendly products—you're supporting Indonesian communities and the circular economy.

**🧺 WHAT'S INCLUDED IN THIS COMPLETE SET:**

Our premium party pack includes everything needed for **10 people**:

**Plates & Bowls:**
- **10x Dinner Plates (10-inch / 25cm square)** - Main course plates, sturdy enough for heavy meals like nasi goreng, satay, or roasted meats [citation:3][citation:8]
- **10x Salad/Dessert Plates (7-inch / 18cm square)** - Perfect for appetizers, desserts, or side dishes [citation:8]
- **10x Soup/Cereal Bowls (6-inch / 15cm round)** - Deep enough for soups, curries, or oatmeal [citation:3][citation:8]
- **10x Rectangular Serving Trays (9x6-inch / 23x15cm)** - Ideal for appetizers, sushi, or tapas [citation:3]

**Cutlery Set (Biodegradable Bamboo/Wood):**
- **10x Forks** - Smooth finish, comfortable grip
- **10x Knives** - Serrated edge for cutting
- **10x Spoons** - Deep bowl design perfect for Indonesian cuisine
- **10x Chopsticks (optional)** - For Asian-inspired meals

**Total: 70 pieces** — enough for a complete 10-person dining experience [citation:3].

**✨ KEY FEATURES:**

- ✅ **Heavy-Duty Construction**: Unlike flimsy paper plates, palm leaf dinnerware is naturally strong and rigid. Holds heavy foods, gravies, and saucy dishes without bending or leaking [citation:3][citation:8].
- ✅ **Heat & Liquid Resistant**: Withstands hot foods up to 220°F (104°C). Safe for microwave and oven use (up to 350°F for short periods) [citation:5][citation:10].
- ✅ **100% Chemical-Free**: No plastic lining, wax coating, or artificial binders. The natural heat and pressure manufacturing process binds the leaves together [citation:5].
- ✅ **Beautiful Natural Texture**: Each piece has unique grain patterns and slight color variations—no two plates are exactly alike, adding rustic elegance to any table setting [citation:1][citation:9].
- ✅ **Compostable**: After use, simply compost in your garden or dispose of with organic waste. Biodegrades within 60-180 days [citation:7][citation:10].

**🌍 SUSTAINABILITY IMPACT:**

- **Waste Reduction**: Each set diverts agricultural waste (fallen palm leaves) from being burned, reducing air pollution in rural communities [citation:2].
- **Carbon Neutral**: Production uses fallen leaves—no trees are cut down. Palm trees absorb CO2 while growing, making the process carbon neutral [citation:5][citation:7].
- **Water Conservation**: Unlike paper production, palm leaf manufacturing requires minimal water—no pulping or bleaching processes [citation:5].
- **Plastic Replacement**: Replaces up to 70 single-use plastic items per set, preventing them from entering oceans and landfills.
- **Community Empowerment**: Supports Indonesian farmers and artisans with fair wages and sustainable livelihoods [citation:2][citation:7].

**📋 CARE & DISPOSAL:**

- **Before Use**: Wipe clean with a dry cloth. No washing needed—plates are clean and ready to use.
- **After Use**: Compost in backyard compost bin or dispose of with organic waste. Will break down naturally.
- **Avoid**: Prolonged soaking or dishwasher use (these are single-use compostable items, not reusable).

**🎯 PERFECT FOR:**

- Weddings & receptions
- Corporate events & parties
- Catering & food delivery
- BBQs & picnics
- Camping & outdoor gatherings
- Restaurants & cafes seeking eco-friendly takeout packaging
- Daily home use (especially for large families)

Join the growing movement of eco-conscious Indonesians choosing sustainable alternatives. Make your next event unforgettable AND environmentally responsible!`,
        IsPublished:      true,
        Tags:             pq.StringArray{"palm leaf plates", "areca palm", "compostable dinnerware", "biodegradable", "eco-friendly", "disposable plates", "party supplies", "wedding catering", "bamboo cutlery", "zero waste party", "sustainable tableware", "natural", "heavy-duty plates", "Indonesian handmade", "ramah lingkungan", "piring daun", "eco catering", "bento", "takeaway packaging"},
        Slug:             "heavy-duty-compostable-palm-leaf-plate-cutlery-set-complete-10-person",
        MainImageURL:     "https://ecolipak.com/cdn/shop/files/square-leak-proof-palm-leaf-compostable-plates-cutlery-setsfood-grade-bpa-free-8365275.png?v=1770090716&width=1024",
        MainImagePublicID: "palm-leaf-set-main",
        AverageRating:    4.8,
        ReviewCount:      342,
        SoldCount:        1567,
        MinPrice:         0, // Will be updated after SKUs
        MaxPrice:         0, // Will be updated after SKUs
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }

    // Create Variant Types
    // Variant 1: Set Size / Party Pack
    setSizeType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Set Size",
    }
    if err := db.Create(&setSizeType).Error; err != nil {
        return err
    }

    // Variant 2: Plate Shape
    shapeType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Plate Shape",
    }
    if err := db.Create(&shapeType).Error; err != nil {
        return err
    }

    // Variant 3: Cutlery Material
    cutleryType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Cutlery Material",
    }
    if err := db.Create(&cutleryType).Error; err != nil {
        return err
    }

    // Variant 4: Packaging Type
    packagingType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Packaging",
    }
    if err := db.Create(&packagingType).Error; err != nil {
        return err
    }

    // Create Variant Values
    // Set Size options (based on common party sizes) [citation:3][citation:8]
    setSizeValues := []entity.VariantValue{
        {
            VariantTypeID: setSizeType.ID,
            Value:         "Starter Pack - 30 pieces (Serves 4-5 people)",
        },
        {
            VariantTypeID: setSizeType.ID,
            Value:         "Standard Party Pack - 70 pieces (Serves 10 people)",
        },
        {
            VariantTypeID: setSizeType.ID,
            Value:         "Large Party Pack - 125 pieces (Serves 18-20 people)",
        },
        {
            VariantTypeID: setSizeType.ID,
            Value:         "Wedding/Catering Pack - 250 pieces (Serves 35-40 people)",
        },
        {
            VariantTypeID: setSizeType.ID,
            Value:         "Bulk Wholesale - 500 pieces (Restaurant/Catering)",
        },
    }
    
    for i := range setSizeValues {
        if err := db.Create(&setSizeValues[i]).Error; err != nil {
            return err
        }
    }

    // Plate Shape options [citation:1][citation:4]
    shapeValues := []entity.VariantValue{
        {
            VariantTypeID: shapeType.ID,
            Value:         "Round - Classic",
        },
        {
            VariantTypeID: shapeType.ID,
            Value:         "Square - Modern",
        },
        {
            VariantTypeID: shapeType.ID,
            Value:         "Rectangle - Platter Style",
        },
        {
            VariantTypeID: shapeType.ID,
            Value:         "Heart Shape - Wedding Special",
        },
        {
            VariantTypeID: shapeType.ID,
            Value:         "Leaf Shape - Natural Edge",
        },
        {
            VariantTypeID: shapeType.ID,
            Value:         "Mixed Shapes (Assorted)",
        },
    }
    
    for i := range shapeValues {
        if err := db.Create(&shapeValues[i]).Error; err != nil {
            return err
        }
    }

    // Cutlery Material options [citation:10]
    cutleryValues := []entity.VariantValue{
        {
            VariantTypeID: cutleryType.ID,
            Value:         "Bamboo - Premium",
        },
        {
            VariantTypeID: cutleryType.ID,
            Value:         "Birch Wood - Standard",
        },
        {
            VariantTypeID: cutleryType.ID,
            Value:         "Palm Wood (from leaf stems) - Eco",
        },
        {
            VariantTypeID: cutleryType.ID,
            Value:         "Compostable CPLA (Cornstarch) - Plastic-Free Alternative",
        },
        {
            VariantTypeID: cutleryType.ID,
            Value:         "No Cutlery (Plates Only)",
        },
    }
    
    for i := range cutleryValues {
        if err := db.Create(&cutleryValues[i]).Error; err != nil {
            return err
        }
    }

    // Packaging Type (eco-friendly options)
    packagingValues := []entity.VariantValue{
        {
            VariantTypeID: packagingType.ID,
            Value:         "Kraft Paper Box - Recyclable",
        },
        {
            VariantTypeID: packagingType.ID,
            Value:         "Jute Bag - Reusable",
        },
        {
            VariantTypeID: packagingType.ID,
            Value:         "Shrink Wrap + Cardboard Sleeve - Minimal",
        },
        {
            VariantTypeID: packagingType.ID,
            Value:         "Bamboo Tray + Cellophane - Gift Ready",
        },
        {
            VariantTypeID: packagingType.ID,
            Value:         "Bulk Carton - Wholesale",
        },
    }
    
    for i := range packagingValues {
        if err := db.Create(&packagingValues[i]).Error; err != nil {
            return err
        }
    }

    // Base price mapping based on market research [citation:1][citation:4][citation:8]
    // Prices converted to IDR (approx Rp 15,000 = US$1)
    
    basePriceMap := map[string]map[string]float64{
        "Starter Pack - 30 pieces (Serves 4-5 people)": {
            "Round - Classic":         85000,
            "Square - Modern":         95000,
            "Rectangle - Platter Style": 105000,
            "Heart Shape - Wedding Special": 125000,
            "Leaf Shape - Natural Edge": 115000,
            "Mixed Shapes (Assorted)": 90000,
        },
        "Standard Party Pack - 70 pieces (Serves 10 people)": {
            "Round - Classic":         165000,
            "Square - Modern":         179000,
            "Rectangle - Platter Style": 195000,
            "Heart Shape - Wedding Special": 225000,
            "Leaf Shape - Natural Edge": 209000,
            "Mixed Shapes (Assorted)": 175000,
        },
        "Large Party Pack - 125 pieces (Serves 18-20 people)": {
            "Round - Classic":         275000,
            "Square - Modern":         299000,
            "Rectangle - Platter Style": 325000,
            "Heart Shape - Wedding Special": 375000,
            "Leaf Shape - Natural Edge": 349000,
            "Mixed Shapes (Assorted)": 289000,
        },
        "Wedding/Catering Pack - 250 pieces (Serves 35-40 people)": {
            "Round - Classic":         495000,
            "Square - Modern":         535000,
            "Rectangle - Platter Style": 575000,
            "Heart Shape - Wedding Special": 650000,
            "Leaf Shape - Natural Edge": 599000,
            "Mixed Shapes (Assorted)": 515000,
        },
        "Bulk Wholesale - 500 pieces (Restaurant/Catering)": {
            "Round - Classic":         895000,
            "Square - Modern":         965000,
            "Rectangle - Platter Style": 1025000,
            "Heart Shape - Wedding Special": 1150000,
            "Leaf Shape - Natural Edge": 1075000,
            "Mixed Shapes (Assorted)": 925000,
        },
    }

    // Cutlery material premium adjustments
    cutleryPremium := map[string]float64{
        "Bamboo - Premium":                            35000,
        "Birch Wood - Standard":                       20000,
        "Palm Wood (from leaf stems) - Eco":           25000,
        "Compostable CPLA (Cornstarch) - Plastic-Free Alternative": 30000,
        "No Cutlery (Plates Only)":                    0,
    }

    // Packaging premium/discount
    packagingPremium := map[string]float64{
        "Kraft Paper Box - Recyclable":                 0,
        "Jute Bag - Reusable":                          15000,
        "Shrink Wrap + Cardboard Sleeve - Minimal":    -5000,
        "Bamboo Tray + Cellophane - Gift Ready":        25000,
        "Bulk Carton - Wholesale":                      -10000,
    }

    var allSKUs []entity.SKU
    var minPrice float64 = 999999
    var maxPrice float64 = 0

    // Generate SKUs for combinations
    for _, setSize := range setSizeValues {
        sizePrices, exists := basePriceMap[setSize.Value]
        if !exists {
            continue
        }
        
        for _, shape := range shapeValues {
            basePrice, shapeExists := sizePrices[shape.Value]
            if !shapeExists {
                continue
            }
            
            for _, cutlery := range cutleryValues {
                // Skip incompatible combinations
                // Bulk wholesale often comes without cutlery or with basic
                if setSize.Value == "Bulk Wholesale - 500 pieces (Restaurant/Catering)" && 
                   cutlery.Value == "Compostable CPLA (Cornstarch) - Plastic-Free Alternative" {
                    continue // Premium cutlery not typical for bulk
                }
                
                for _, packaging := range packagingValues {
                    // Skip incompatible packaging-size combinations
                    // Gift packaging only for smaller sets
                    if packaging.Value == "Bamboo Tray + Cellophane - Gift Ready" && 
                       (setSize.Value == "Bulk Wholesale - 500 pieces (Restaurant/Catering)" || 
                        setSize.Value == "Wedding/Catering Pack - 250 pieces (Serves 35-40 people)") {
                        continue
                    }
                    
                    // Bulk carton only for large sets
                    if packaging.Value == "Bulk Carton - Wholesale" && 
                       !(setSize.Value == "Bulk Wholesale - 500 pieces (Restaurant/Catering)" || 
                         setSize.Value == "Wedding/Catering Pack - 250 pieces (Serves 35-40 people)") {
                        continue
                    }
                    
                    // Calculate final price
                    cutleryAdj := cutleryPremium[cutlery.Value]
                    packagingAdj := packagingPremium[packaging.Value]
                    
                    finalPrice := basePrice + cutleryAdj + packagingAdj
                    
                    // Ensure price doesn't go negative
                    if finalPrice < 10000 {
                        finalPrice = 10000
                    }
                    
                    // Round to nearest 1000 Rupiah (common in Indonesia)
                    finalPrice = float64(int(finalPrice/1000)) * 1000
                    
                    // Generate SKU code
                    skuCode := fmt.Sprintf("PALM-%s-%s-%s-%s-%d", 
                        abbreviateSetSize(setSize.Value),
                        abbreviateShape(shape.Value),
                        abbreviateCutlery(cutlery.Value),
                        abbreviatePackaging(packaging.Value),
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
                    if setSize.Value == "Standard Party Pack - 70 pieces (Serves 10 people)" {
                        stock = 150 + rand.Intn(200) // Most popular: 150-350
                    } else if setSize.Value == "Large Party Pack - 125 pieces (Serves 18-20 people)" {
                        stock = 80 + rand.Intn(150) // 80-230
                    } else if setSize.Value == "Starter Pack - 30 pieces (Serves 4-5 people)" {
                        stock = 100 + rand.Intn(200) // 100-300
                    } else if setSize.Value == "Wedding/Catering Pack - 250 pieces (Serves 35-40 people)" {
                        stock = 30 + rand.Intn(70) // 30-100
                    } else if setSize.Value == "Bulk Wholesale - 500 pieces (Restaurant/Catering)" {
                        stock = 10 + rand.Intn(30) // 10-40
                    }
                    
                    // Adjust stock for premium variants
                    if cutleryAdj > 30000 || packagingAdj > 20000 {
                        stock = stock / 2
                    }
                    
                    // Calculate weight (varies by set size)
                    weight := 0.0
                    switch setSize.Value {
                    case "Starter Pack - 30 pieces (Serves 4-5 people)":
                        weight = 800
                    case "Standard Party Pack - 70 pieces (Serves 10 people)":
                        weight = 1880 // ~1.88kg [citation:3]
                    case "Large Party Pack - 125 pieces (Serves 18-20 people)":
                        weight = 3200
                    case "Wedding/Catering Pack - 250 pieces (Serves 35-40 people)":
                        weight = 6000
                    case "Bulk Wholesale - 500 pieces (Restaurant/Catering)":
                        weight = 11500
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
                        VariantValueID: setSize.ID,
                    }
                    db.Create(&skuVariant1)
                    
                    skuVariant2 := entity.SKUVariantValue{
                        SKUID:          sku.ID,
                        VariantValueID: shape.ID,
                    }
                    db.Create(&skuVariant2)
                    
                    skuVariant3 := entity.SKUVariantValue{
                        SKUID:          sku.ID,
                        VariantValueID: cutlery.ID,
                    }
                    db.Create(&skuVariant3)
                    
                    skuVariant4 := entity.SKUVariantValue{
                        SKUID:          sku.ID,
                        VariantValueID: packaging.ID,
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
			"https://ecolipak.com/cdn/shop/files/6-in-10-in-square-heavy-duty-compostable-palm-leaf-platetoxin-free-plastic-free-5347982.png?v=1766614395&width=1024",
			"https://ecolipak.com/cdn/shop/files/6-in-10-in-square-heavy-duty-compostable-palm-leaf-platetoxin-free-plastic-free-2170702.png?v=1766614395&width=1024",
			"https://m.media-amazon.com/images/I/81t24aukheL._AC_UF1000,1000_QL80_.jpg",
		}

		for _, url := range imageURLs {
			img := entity.Image{
				ProductID: product.ID,
				ImageURL: url,
			}

			db.Create(&img)
		}

    fmt.Printf("✅ Successfully seeded Palm Leaf Dinnerware Set product with %d SKU variants\n", len(allSKUs))
    fmt.Printf("   Price range: Rp %.0f - Rp %.0f\n", minPrice, maxPrice)
    
    return nil
}

// Helper functions for SKU code generation
func abbreviateSetSize(setSize string) string {
    switch {
    case setSize == "Starter Pack - 30 pieces (Serves 4-5 people)":
        return "STR"
    case setSize == "Standard Party Pack - 70 pieces (Serves 10 people)":
        return "STD"
    case setSize == "Large Party Pack - 125 pieces (Serves 18-20 people)":
        return "LRG"
    case setSize == "Wedding/Catering Pack - 250 pieces (Serves 35-40 people)":
        return "WED"
    case setSize == "Bulk Wholesale - 500 pieces (Restaurant/Catering)":
        return "BLK"
    default:
        return "SIZ"
    }
}

func abbreviateShape(shape string) string {
    switch {
    case shape == "Round - Classic":
        return "RND"
    case shape == "Square - Modern":
        return "SQR"
    case shape == "Rectangle - Platter Style":
        return "REC"
    case shape == "Heart Shape - Wedding Special":
        return "HRT"
    case shape == "Leaf Shape - Natural Edge":
        return "LEF"
    case shape == "Mixed Shapes (Assorted)":
        return "MIX"
    default:
        return "SHP"
    }
}

func abbreviateCutlery(cutlery string) string {
    switch {
    case cutlery == "Bamboo - Premium":
        return "BAM"
    case cutlery == "Birch Wood - Standard":
        return "BRC"
    case cutlery == "Palm Wood (from leaf stems) - Eco":
        return "PLM"
    case cutlery == "Compostable CPLA (Cornstarch) - Plastic-Free Alternative":
        return "CPL"
    case cutlery == "No Cutlery (Plates Only)":
        return "NOC"
    default:
        return "CUT"
    }
}

func abbreviatePackaging(pkg string) string {
    switch {
    case pkg == "Kraft Paper Box - Recyclable":
        return "KRT"
    case pkg == "Jute Bag - Reusable":
        return "JUT"
    case pkg == "Shrink Wrap + Cardboard Sleeve - Minimal":
        return "SHR"
    case pkg == "Bamboo Tray + Cellophane - Gift Ready":
        return "GFT"
    case pkg == "Bulk Carton - Wholesale":
        return "BLK"
    default:
        return "PKG"
    }
}