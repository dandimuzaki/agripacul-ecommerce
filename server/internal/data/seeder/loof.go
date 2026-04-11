package seeder

import (
	"debian-ecommerce/internal/data/entity"
	"fmt"
	"math/rand"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func SeedNaturalLoofahSponge(db *gorm.DB) error {
    rand.Seed(time.Now().UnixNano())
    
    // First, ensure category exists (category ID 2 for Personal Care or 3 for Home Cleaning)
    // This product can serve both bath and kitchen purposes
    var category entity.Category
    result := db.First(&category, 2) // Personal Care category
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

    // Create the main product - Natural Loofah Sponge
    product := entity.Product{
        CategoryID:       2,
        Name:             "Natural Loofah Sponge (Gambas Kering) - Plastic-Free Bath & Kitchen Scrub",
        Description:      `Meet the **Natural Loofah Sponge**—a traditional Indonesian zero-waste essential experiencing a modern revival. Made from the dried fruit of the *Luffa cylindrica* plant (known locally as *gambas* or *oyong*), this 100% natural sponge is the perfect eco-friendly alternative to plastic shower puffs and synthetic kitchen scrubbers.

**🌱 WHAT IS LOOFAH?**

Loofah is not a sea creature—it's a plant! It comes from the mature fruit of the *Luffa* gourd, a climbing vine related to cucumbers and pumpkins. When the fruit ripens and dries, it develops a dense network of fibrous tissue that forms a natural, biodegradable sponge .

For generations, Indonesian farmers have grown gambas primarily for vegetable consumption, with the dried fibrous interior often considered waste or used casually for cleaning. Today, innovative local brands and social enterprises are transforming this traditional knowledge into sustainable products that replace plastic alternatives [citation:2][citation:3].

**✨ MULTIPLE USES:**

**For Bath & Body:**
- **Natural Exfoliator**: Gently removes dead skin cells and dirt, revitalizing skin without harsh chemicals [citation:1]
- **Shower Puff Alternative**: Replaces plastic shower poufs that shed microplastics into waterways [citation:3]
- **Circulation Booster**: The mild scrubbing action stimulates blood flow during bathing

**For Kitchen & Home:**
- **Dish Scrubber**: Effectively cleans dishes, pots, and pans without scratching surfaces [citation:6]
- **Produce Washer**: Gently scrubs fruits and vegetables without chemical residues
- **Surface Cleaner**: Ideal for countertops and sinks

**🌿 SUSTAINABILITY CREDENTIALS:**

- ✅ **100% Biodegradable**: At end of life (3-6 months), compost it—it will return to soil naturally [citation:2]
- ✅ **Plastic-Free**: No synthetic polymers—prevents microplastic pollution in oceans [citation:3]
- ✅ **Renewable Resource**: Grown by farmers in Java, Bali, and Sumatra—supports local agriculture [citation:2][citation:4]
- ✅ **Chemical-Free Processing**: Sun-dried naturally, no bleaching or chemical treatments [citation:2]
- ✅ **Carbon Negative**: Growing plants absorb CO2; processing requires minimal energy

**🇮🇩 SUPPORTING INDONESIAN FARMERS & THE CIRCULAR ECONOMY**

The natural loofah movement in Indonesia is more than just a product—it's a growing ecosystem of sustainable livelihoods:

- **Arrafah Shop** (Bandung): Transforms oyong fibers into beautiful, skin-friendly bath sponges while supporting local farmers [citation:3]
- **Kini Bumi** (Jember): Offers multiple loofah formats, from whole elongated sponges to ready-to-use oval pads [citation:4]
- **Kedai Ramah Bumi**: Sources from farmers in East Java and Bali, emphasizing sun-drying without chemical bleaching [citation:2]
- **Mahadewi Bali**: Bali-based brand integrating loofah into their eco-friendly body care line [citation:4]

When you choose natural loofah, you're not just avoiding plastic—you're creating economic opportunities for Indonesian farming communities and supporting the transition away from agricultural waste burning.

**📋 HOW TO USE:**

**First Time Use:**
1. Soak the loofah in warm water for a few minutes until it softens [citation:2]
2. For bath use, some recommend an initial hot water soak to ensure sterility [citation:2]
3. Squeeze out excess water—it becomes wonderfully pliable when wet

**Daily Bath Use:**
1. Wet the loofah and apply liquid soap or rub against a soap bar
2. Gently massage in circular motions over skin—the texture exfoliates without irritation
3. Rinse thoroughly after use

**Daily Kitchen Use:**
1. Cut the loofah into smaller pieces (one whole loofah can become 2-3 kitchen scrubbers) [citation:6]
2. Use with dish soap to clean dishes, pots, and pans
3. Rinse thoroughly after each use

**🧼 CARE & MAINTENANCE:**

To extend the life of your loofah (typically 1-6 months depending on use):

- **Rinse thoroughly** after each use to remove soap residue and trapped debris
- **Squeeze out excess water**—loofahs should not remain soaking wet
- **Hang to dry** in a well-ventilated area, preferably in sunlight which has natural sterilizing properties [citation:2]
- **Replace when** it becomes dark, develops an odor, or starts breaking down (signs it's time to compost)

**⚠️ IMPORTANT NOTE:**

Unlike plastic poufs, natural loofah is a plant product and will eventually biodegrade. This is a feature, not a bug! When it shows signs of wear, simply compost it and purchase a fresh one—each loofah supports farmers and keeps plastic out of landfills.

**🧴 AVAILABLE IN THESE VARIATIONS:**

- **Whole Loofah**: Traditional elongated form, can be cut to desired size
- **Loofah Pads**: Pre-cut oval or round shapes, ready to use
- **Loofah with Handle**: Some variations include cotton or bamboo handles for easier grip
- **Loofah Scrub Mitts**: Fabric mitts with integrated loofah panel

Join the growing movement of Indonesians rediscovering traditional wisdom for modern sustainable living. Each loofah you choose is one less plastic puff in our oceans! [citation:3]`,
        IsPublished:      true,
        Tags:             pq.StringArray{"loofah", "gambas kering", "natural sponge", "oyong", "body scrub", "zero waste", "plastic-free", "biodegradable", "exfoliator", "bath sponge", "kitchen scrubber", "eco-friendly", "traditional", "local farmers", "sustainable", "shower puff", "dish scrub", "sel kulit mati", "alami", "ramah lingkungan"},
        Slug:             "natural-loofah-sponge-gambas-kering-plastic-free-bath-kitchen-scrub",
        MainImageURL:     "https://down-id.img.susercontent.com/file/id-11134207-7r98v-lurepfihsxwxf0@resize_w900_nl.webp",
        MainImagePublicID: "loofah-sponge-main",
        AverageRating:    4.8,
        ReviewCount:      567,
        SoldCount:        2345,
        MinPrice:         0, // Will be updated after SKUs
        MaxPrice:         0, // Will be updated after SKUs
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }

    // Create Variant Types
    // Variant 1: Product Form
    formType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Form / Shape",
    }
    if err := db.Create(&formType).Error; err != nil {
        return err
    }

    // Variant 2: Size
    sizeType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Size",
    }
    if err := db.Create(&sizeType).Error; err != nil {
        return err
    }

    // Variant 3: Purpose (Bath or Kitchen)
    purposeType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Primary Use",
    }
    if err := db.Create(&purposeType).Error; err != nil {
        return err
    }

    // Variant 4: Origin/Processing
    originType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Origin / Processing",
    }
    if err := db.Create(&originType).Error; err != nil {
        return err
    }

    // Create Variant Values
    // Form / Shape options
    formValues := []entity.VariantValue{
        {
            VariantTypeID: formType.ID,
            Value:         "Whole - Traditional Elongated",
        },
        {
            VariantTypeID: formType.ID,
            Value:         "Pad - Oval / Round (Ready to Use)",
        },
        {
            VariantTypeID: formType.ID,
            Value:         "With Cotton Handle - Bath Style",
        },
        {
            VariantTypeID: formType.ID,
            Value:         "Scrub Mitt - With Loofah Panel",
        },
        {
            VariantTypeID: formType.ID,
            Value:         "Cut Pieces - Multipack for Kitchen",
        },
    }
    
    for i := range formValues {
        if err := db.Create(&formValues[i]).Error; err != nil {
            return err
        }
    }

    // Size options (based on actual market sizes)
    sizeValues := []entity.VariantValue{
        {
            VariantTypeID: sizeType.ID,
            Value:         "Extra Small (S) - 12-18cm / 2-4cm diameter",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "Small (S) - 12-18cm",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "Medium (M) - 19-25cm",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "Large (L) - 26-30cm",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "Extra Large (XL) - 31-45cm",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "Pad - Small (8-10cm diameter)",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "Pad - Medium (10-12cm diameter)",
        },
        {
            VariantTypeID: sizeType.ID,
            Value:         "Pad - Large (12-15cm diameter)",
        },
    }
    
    for i := range sizeValues {
        if err := db.Create(&sizeValues[i]).Error; err != nil {
            return err
        }
    }

    // Purpose options
    purposeValues := []entity.VariantValue{
        {
            VariantTypeID: purposeType.ID,
            Value:         "Multi-Purpose (Bath + Kitchen)",
        },
        {
            VariantTypeID: purposeType.ID,
            Value:         "Bath & Body - Premium",
        },
        {
            VariantTypeID: purposeType.ID,
            Value:         "Kitchen & Dish - Heavy Duty",
        },
        {
            VariantTypeID: purposeType.ID,
            Value:         "Facial Scrub - Ultra Soft",
        },
    }
    
    for i := range purposeValues {
        if err := db.Create(&purposeValues[i]).Error; err != nil {
            return err
        }
    }

    // Origin / Processing options (important for consumer preference)
    originValues := []entity.VariantValue{
        {
            VariantTypeID: originType.ID,
            Value:         "Local Java - Sun-Dried, No Bleach",
        },
        {
            VariantTypeID: originType.ID,
            Value:         "Local Bali - Sun-Dried, No Bleach",
        },
        {
            VariantTypeID: originType.ID,
            Value:         "Local Sumatra - Sun-Dried, No Bleach",
        },
        {
            VariantTypeID: originType.ID,
            Value:         "Premium Processed - Gentle Washed",
        },
        {
            VariantTypeID: originType.ID,
            Value:         "Organic Certified - Chemical-Free",
        },
    }
    
    for i := range originValues {
        if err := db.Create(&originValues[i]).Error; err != nil {
            return err
        }
    }

    // Base price mapping based on market research
    // Source: Tokopedia listings and brand data [citation:1][citation:2][citation:4]
    // Kini Bumi: Rp 5,200 - 15,000
    // Mahadewi Bali: Rp 15,000
    // Segara Naturals: Rp 18,000 - 35,000
    // Demi Bumi: Rp 26,000
    // Ecotools: Rp 69,000 (premium)
    // Generic: Rp 12,000 - 25,000

    basePriceMap := map[string]map[string]float64{
        "Whole - Traditional Elongated": {
            "Extra Small (S) - 12-18cm / 2-4cm diameter": 12500,
            "Small (S) - 12-18cm":                         15000,
            "Medium (M) - 19-25cm":                        20000,
            "Large (L) - 26-30cm":                         25000,
            "Extra Large (XL) - 31-45cm":                  30000,
        },
        "Pad - Oval / Round (Ready to Use)": {
            "Pad - Small (8-10cm diameter)":               8500,
            "Pad - Medium (10-12cm diameter)":             12000,
            "Pad - Large (12-15cm diameter)":              16000,
        },
        "With Cotton Handle - Bath Style": {
            "Medium (M) - 19-25cm":                         28000,
            "Large (L) - 26-30cm":                          32000,
        },
        "Scrub Mitt - With Loofah Panel": {
            "Medium (M) - 19-25cm":                         35000,
            "Large (L) - 26-30cm":                          40000,
        },
        "Cut Pieces - Multipack for Kitchen": {
            "Pack of 3 pieces":                             20000,
            "Pack of 5 pieces":                             30000,
            "Pack of 10 pieces":                            50000,
        },
    }

    // Purpose premium adjustments
    purposePremium := map[string]float64{
        "Multi-Purpose (Bath + Kitchen)":       0,
        "Bath & Body - Premium":                5000,   // Premium for finer texture
        "Kitchen & Dish - Heavy Duty":          -2000,  // Discount for tougher fibers
        "Facial Scrub - Ultra Soft":            8000,   // Premium for gentle exfoliation
    }

    // Origin premium (local, ethical sourcing commands premium)
    originPremium := map[string]float64{
        "Local Java - Sun-Dried, No Bleach":     0,      // Standard
        "Local Bali - Sun-Dried, No Bleach":     2000,   // Bali premium (brand perception)
        "Local Sumatra - Sun-Dried, No Bleach":  0,      // Standard
        "Premium Processed - Gentle Washed":     5000,   // Extra processing
        "Organic Certified - Chemical-Free":     8000,   // Certification premium
    }

    var allSKUs []entity.SKU
    var minPrice float64 = 999999
    var maxPrice float64 = 0

    // Generate SKUs for combinations
    for _, form := range formValues {
        formPrices, exists := basePriceMap[form.Value]
        if !exists {
            continue
        }
        
        for _, size := range sizeValues {
            // Check if this size exists for this form
            basePrice, sizeExists := formPrices[size.Value]
            if !sizeExists {
                continue
            }
            
            for _, purpose := range purposeValues {
                // Skip incompatible purpose-form combinations
                if purpose.Value == "Facial Scrub - Ultra Soft" && 
                   form.Value == "Cut Pieces - Multipack for Kitchen" {
                    continue // Kitchen multipack not for facial use
                }
                
                if purpose.Value == "Kitchen & Dish - Heavy Duty" && 
                   (form.Value == "Scrub Mitt - With Loofah Panel" || 
                    form.Value == "With Cotton Handle - Bath Style") {
                    continue // Bath-specific forms not for kitchen
                }
                
                for _, origin := range originValues {
                    // Skip incompatible origin-form combinations
                    // Most origins work with all forms, but organic certification more common for bath
                    
                    // Calculate final price
                    purposeAdj := purposePremium[purpose.Value]
                    originAdj := originPremium[origin.Value]
                    
                    finalPrice := basePrice + purposeAdj + originAdj
                    
                    // Ensure price doesn't go negative
                    if finalPrice < 5000 {
                        finalPrice = 5000
                    }
                    
                    // Round to nearest 500 Rupiah
                    finalPrice = float64(int(finalPrice/500)) * 500
                    
                    // Generate SKU code
                    skuCode := fmt.Sprintf("LOOF-%s-%s-%s-%s-%d", 
                        abbreviateLoofahForm(form.Value),
                        abbreviateLoofahSize(size.Value),
                        abbreviateLoofahPurpose(purpose.Value),
                        abbreviateLoofahOrigin(origin.Value),
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
                    if form.Value == "Whole - Traditional Elongated" && 
                       size.Value == "Medium (M) - 19-25cm" {
                        stock = 200 + rand.Intn(300) // Most popular: 200-500
                    } else if form.Value == "Pad - Oval / Round (Ready to Use)" {
                        stock = 150 + rand.Intn(250) // 150-400
                    } else if form.Value == "Cut Pieces - Multipack for Kitchen" {
                        stock = 100 + rand.Intn(200) // 100-300
                    } else if purpose.Value == "Facial Scrub - Ultra Soft" {
                        stock = 50 + rand.Intn(100) // 50-150 (niche)
                    } else {
                        stock = 75 + rand.Intn(150) // 75-225
                    }
                    
                    // Adjust stock for premium variants
                    if purposeAdj > 5000 || originAdj > 5000 {
                        stock = stock / 2
                    }
                    
                    // Calculate weight
                    weight := 0.0
                    switch {
                    case form.Value == "Whole - Traditional Elongated":
                        if size.Value == "Extra Small (S) - 12-18cm / 2-4cm diameter" {
                            weight = 20
                        } else if size.Value == "Small (S) - 12-18cm" {
                            weight = 30
                        } else if size.Value == "Medium (M) - 19-25cm" {
                            weight = 45
                        } else if size.Value == "Large (L) - 26-30cm" {
                            weight = 60
                        } else if size.Value == "Extra Large (XL) - 31-45cm" {
                            weight = 80
                        }
                    case form.Value == "Pad - Oval / Round (Ready to Use)":
                        if size.Value == "Pad - Small (8-10cm diameter)" {
                            weight = 15
                        } else if size.Value == "Pad - Medium (10-12cm diameter)" {
                            weight = 22
                        } else if size.Value == "Pad - Large (12-15cm diameter)" {
                            weight = 30
                        }
                    case form.Value == "With Cotton Handle - Bath Style":
                        weight = 50
                    case form.Value == "Scrub Mitt - With Loofah Panel":
                        weight = 80
                    case form.Value == "Cut Pieces - Multipack for Kitchen":
                        if size.Value == "Pack of 3 pieces" {
                            weight = 60
                        } else if size.Value == "Pack of 5 pieces" {
                            weight = 100
                        } else if size.Value == "Pack of 10 pieces" {
                            weight = 200
                        }
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
                        VariantValueID: purpose.ID,
                    }
                    db.Create(&skuVariant3)
                    
                    skuVariant4 := entity.SKUVariantValue{
                        SKUID:          sku.ID,
                        VariantValueID: origin.ID,
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
			"https://down-id.img.susercontent.com/file/id-11134207-7r98r-lurepfihuchd07@resize_w900_nl.webp",
			"https://down-id.img.susercontent.com/file/id-11134207-7r98r-luroby2iizzw5b@resize_w900_nl.webp",
			"https://down-id.img.susercontent.com/file/id-11134207-7r990-luroby2ihlfg65.webp",
		}

		for _, url := range imageURLs {
			img := entity.Image{
				ProductID: product.ID,
				ImageURL: url,
			}

			db.Create(&img)
		}

    fmt.Printf("✅ Successfully seeded Natural Loofah Sponge product with %d SKU variants\n", len(allSKUs))
    fmt.Printf("   Price range: Rp %.0f - Rp %.0f\n", minPrice, maxPrice)
    
    return nil
}

// Helper functions for SKU code generation
func abbreviateLoofahForm(form string) string {
    switch {
    case form == "Whole - Traditional Elongated":
        return "WHL"
    case form == "Pad - Oval / Round (Ready to Use)":
        return "PAD"
    case form == "With Cotton Handle - Bath Style":
        return "HDL"
    case form == "Scrub Mitt - With Loofah Panel":
        return "MIT"
    case form == "Cut Pieces - Multipack for Kitchen":
        return "KIT"
    default:
        return "FRM"
    }
}

func abbreviateLoofahSize(size string) string {
    switch {
    case size == "Extra Small (S) - 12-18cm / 2-4cm diameter":
        return "XS"
    case size == "Small (S) - 12-18cm":
        return "S"
    case size == "Medium (M) - 19-25cm":
        return "M"
    case size == "Large (L) - 26-30cm":
        return "L"
    case size == "Extra Large (XL) - 31-45cm":
        return "XL"
    case size == "Pad - Small (8-10cm diameter)":
        return "PDS"
    case size == "Pad - Medium (10-12cm diameter)":
        return "PDM"
    case size == "Pad - Large (12-15cm diameter)":
        return "PDL"
    case size == "Pack of 3 pieces":
        return "3PC"
    case size == "Pack of 5 pieces":
        return "5PC"
    case size == "Pack of 10 pieces":
        return "10PC"
    default:
        return "SIZ"
    }
}

func abbreviateLoofahPurpose(purpose string) string {
    switch {
    case purpose == "Multi-Purpose (Bath + Kitchen)":
        return "ALL"
    case purpose == "Bath & Body - Premium":
        return "BTH"
    case purpose == "Kitchen & Dish - Heavy Duty":
        return "KIT"
    case purpose == "Facial Scrub - Ultra Soft":
        return "FAC"
    default:
        return "USE"
    }
}

func abbreviateLoofahOrigin(origin string) string {
    switch {
    case origin == "Local Java - Sun-Dried, No Bleach":
        return "JV"
    case origin == "Local Bali - Sun-Dried, No Bleach":
        return "BL"
    case origin == "Local Sumatra - Sun-Dried, No Bleach":
        return "SM"
    case origin == "Premium Processed - Gentle Washed":
        return "PRM"
    case origin == "Organic Certified - Chemical-Free":
        return "ORG"
    default:
        return "LOC"
    }
}