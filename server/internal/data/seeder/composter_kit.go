package seeder

import (
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/pkg/utils"
	"fmt"
	"math/rand"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func SeedComposterCollection(db *gorm.DB) error {
    rand.Seed(time.Now().UnixNano())
    
    // First, ensure categories exist
    category, err := createGardeningToolsCategory(db)
    if err != nil {
        return err
    }
    
    // Seed each composter line as separate products
    if err := seedUrbanComposter(db, category.ID); err != nil {
        return err
    }
    
    if err := seedFamilyComposter(db, category.ID); err != nil {
        return err
    }
    
    if err := seedCommunityComposter(db, category.ID); err != nil {
        return err
    }
    
    if err := seedSmartComposter(db, category.ID); err != nil {
        return err
    }
    
    if err := seedDIYComposterKit(db, category.ID); err != nil {
        return err
    }
    
    if err := seedCompostingStarterKits(db, category.ID); err != nil {
        return err
    }
    
    fmt.Println("✅ Successfully seeded Composter Collection (6 product lines)")
    return nil
}

func createGardeningToolsCategory(db *gorm.DB) (*entity.Category, error) {
  var category entity.Category
	result := db.First(&category, 4)
	if result.Error != nil {
		// Create Home & Cleaning category if it doesn't exist
		category = entity.Category{
			Model:       gorm.Model{ID: 4},
			Name:        "Gardening Tools",
		}
		if err := db.Create(&category).Error; err != nil {
			return nil, err
		}
	}

	return &category, nil
}

// Product 1: Urban Composter - For Apartments & Small Spaces
func seedUrbanComposter(db *gorm.DB, categoryID uint) error {
    product := entity.Product{
        CategoryID: categoryID,
        Name:       "Urban 20L Composter - Solusi Kompos untuk Apartemen & Rumah Kecil",
        Description: `Perfect for apartment dwellers and small households. Our Urban 20L composter fits neatly in your kitchen or balcony, transforming food scraps into liquid fertilizer without odors.

**🇮🇩 Perfect for Urban Indonesian Living:**
Based on successful community initiatives in Serua, Tangerang Selatan, where residents created compact composters from water drums to combat the waste emergency [citation:1]. Now you can do the same with our ready-to-use system.

**✨ What's Included:**
- 20L composter vessel (40×40×50 cm)
- Liquid collection tap
- Aeration lid with carbon filter
- Instruction manual in Bahasa Indonesia
- 1-year warranty

**🌱 How It Works:**
1. Add kitchen scraps daily
2. Stir every 3 days (simple aeration)
3. Harvest liquid fertilizer every week via tap
4. Solid compost ready in 4-6 weeks

**🎯 Best For:**
- 1-2 person households
- Apartment/condo dwellers
- Small balcony gardens
- Kitchen countertop use

**💚 Impact:**
- Diverts 200kg organic waste/year from landfill
- Produces 50L liquid fertilizer annually
- Reduces methane emissions equivalent to planting 2 trees`,
        IsPublished: true,
        Tags:       pq.StringArray{"urban", "apartment", "kecil", "indoor", "tanpa bau", "kitchen composter"},
        Slug:       "urban-20l-composter-apartemen-rumah-kecil",
        MainImageURL: "https://urbankomposter.com/wp-content/uploads/2020/04/urban-komposter-17-liter.png",
        MainImagePublicID: "urban-composter-20l",
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }
    
    // Simple variants: just material options
    variants := []struct {
        Material string
        Price    float64
        Weight   float64
        Stock    int
    }{
        {"Recycled HDPE (Ekonomis)", 350000, 3000, 200},
        {"Stainless Steel (Premium)", 550000, 4500, 50},
        {"Bamboo Composite (Natural)", 475000, 3800, 75},
    }
    
    return createComposterSKUs(db, product, variants, "URB")
}

// Product 2: Family Composter - For Houses with Gardens
func seedFamilyComposter(db *gorm.DB, categoryID uint) error {
    product := entity.Product{
        CategoryID: categoryID,
        Name:       "Family 60L Composter - Untuk Rumah dengan Kebun",
        Description: `Designed for families who garden. The 60L capacity handles all your kitchen waste plus yard trimmings, producing enough compost for a productive vegetable garden.

**🇮🇩 Inspired by Indonesian Communities:**
Following successful programs in Johogunung, Rembang, where UGM students taught residents to make composters from used buckets [citation:2], and Banyuurip where families learned to produce liquid fertilizer within weeks [citation:4][citation:6].

**✨ What's Included:**
- 60L composter vessel (50×50×70 cm)
- Dual collection taps (top/bottom)
- Aeration lid with carbon filter
- Compost thermometer
- Stirring tool
- Detailed guidebook

**🌱 Features:**
- ✅ Handles 3-5kg daily waste
- ✅ Odorless design with carbon filter
- ✅ Dual output: liquid + solid fertilizer
- ✅ Easy harvest system
- ✅ Rat-proof locking lid

**🎯 Best For:**
- 3-4 person households
- Homes with gardens
- Fruit tree owners
- Vegetable gardeners

**💚 Impact:**
- Diverts 400kg organic waste/year
- Produces 150L liquid + 200kg solid compost annually
- Saves Rp 800,000/year on fertilizers`,
        IsPublished: true,
        Tags:       pq.StringArray{"keluarga", "family", "garden", "kebun", "besar", "outdoor"},
        Slug:       "family-60l-composter-rumah-kebun",
        MainImageURL: "https://jubelio-store.s3.ap-southeast-1.amazonaws.com/sustaination/2021/02/31233615/Komposter-Kit-GMI-1-scaled.webp",
        MainImagePublicID: "family-composter-60l",
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }
    
    variants := []struct {
        Material string
        Price    float64
        Weight   float64
        Stock    int
    }{
        {"Recycled HDPE (Standar)", 550000, 5500, 150},
        {"Galvanized Steel (Heavy-Duty)", 850000, 8000, 60},
        {"Stainless Steel (Premium)", 995000, 7500, 40},
    }
    
    return createComposterSKUs(db, product, variants, "FAM")
}

// Product 3: Community Composter - For RT/RW Programs
func seedCommunityComposter(db *gorm.DB, categoryID uint) error {
    product := entity.Product{
        CategoryID: categoryID,
        Name:       "Community 120L Composter - Untuk Program RT/RW & Sekolah",
        Description: `Large-scale composter for community programs, schools, and restaurants. Based on successful initiatives where residents created 40 composter units at Rp 450,000 each to combat the waste crisis [citation:1].

**✨ What's Included:**
- 120L heavy-duty composter (60×60×90 cm)
- Commercial-grade collection system
- Multiple aeration points
- Large harvest door
- Training materials for community programs
- Group buying discount available

**🌱 Perfect For:**
- RT/RW composting programs
- School environmental education
- Small restaurants and cafes
- Community gardens
- Housing complexes

**📊 Specifications:**
- Daily capacity: 6-8kg waste
- Processing time: 3-5 weeks
- Liquid output: 30L/month
- Solid output: 50kg/month

**💚 Community Impact:**
A single unit serving 5-6 families can:
- Divert 1.5 tons of waste from landfill annually
- Create fertilizer for community gardens
- Reduce waste management costs by 40%
- Serve as an educational tool`,
        IsPublished: true,
        Tags:       pq.StringArray{"komunitas", "community", "RT", "RW", "sekolah", "program", "bersama"},
        Slug:       "community-120l-composter-rt-rw-sekolah",
        MainImageURL: "https://www.scdprobiotics.com/cdn/shop/products/all-seasons-indoor-composter.jpg?v=1687467692",
        MainImagePublicID: "community-composter-120l",
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }
    
    variants := []struct {
        Material string
        Price    float64
        Weight   float64
        Stock    int
    }{
        {"Galvanized Steel (Standar)", 995000, 12000, 30},
        {"Stainless Steel (Premium)", 1450000, 14000, 15},
        {"Commercial Grade (Industrial)", 1850000, 18000, 10},
    }
    
    return createComposterSKUs(db, product, variants, "COM")
}

// Product 4: Smart Composter - High-Tech Solution
func seedSmartComposter(db *gorm.DB, categoryID uint) error {
    product := entity.Product{
        CategoryID: categoryID,
        Name:       "Smart Composter WiFi - Otomatis & Terkontrol via Aplikasi",
        Description: `The future of composting is here. Our smart composter uses sensors and automation to optimize the composting process, with real-time monitoring via smartphone.

**✨ Features:**
- ✅ Automatic temperature control
- ✅ Built-in mixing system
- ✅ Moisture sensors
- ✅ WiFi connectivity
- ✅ Mobile app notifications
- ✅ Compost readiness alerts
- ✅ Energy-efficient design

**📱 App Features:**
- Monitor compost temperature
- Receive harvest reminders
- Track waste reduction stats
- Get troubleshooting alerts
- Share progress on social media

**🎯 Perfect For:**
- Tech-savvy homeowners
- Premium restaurant kitchens
- Sustainability enthusiasts
- Smart home integration

**⚙️ Technical Specs:**
- Power: 40W (solar compatible)
- Capacity: 60L
- Processing time: 2-3 weeks
- Noise level: <35dB`,
        IsPublished: true,
        Tags:       pq.StringArray{"smart", "wifi", "otomatis", "high-tech", "aplikasi", "sensor"},
        Slug:       "smart-composter-wi-fi-otomatis",
        MainImageURL: "https://m.media-amazon.com/images/I/71kXnF4b2YL._SX569_.jpg",
        MainImagePublicID: "smart-composter",
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }
    
    variants := []struct {
        Material string
        Price    float64
        Weight   float64
        Stock    int
    }{
        {"Standard Edition", 2850000, 15000, 20},
        {"Solar-Powered Edition", 3450000, 18000, 10},
        {"Commercial Pro", 4250000, 22000, 5},
    }
    
    return createComposterSKUs(db, product, variants, "SMT")
}

// Product 5: DIY Composter Kit - Build Your Own
func seedDIYComposterKit(db *gorm.DB, categoryID uint) error {
    product := entity.Product{
        CategoryID: categoryID,
        Name:       "DIY Composter Kit - Rakit Sendiri dari Ember Bekas",
        Description: `The most affordable way to start composting! Based on successful UGM community programs teaching residents to create composters from used buckets [citation:4]. This kit includes everything you need to convert two 20L buckets into a functioning composter.

**✨ Kit Includes:**
- 2x bucket lids (fits standard 20L ember)
- 1x collection tap with fittings
- 1x aeration grid
- 1x drill bit set
- 1x sealing gasket
- 1x instruction manual
- 1x EM4 starter culture
- 1x molasses sample

**🛠️ You Provide:**
- 2x used 20L buckets (ember bekas)
- Basic tools
- Kitchen waste!

**🌱 Based on Proven Methods:**
This kit follows the successful model used in:
- Banyuurip, Rembang: Residents learned to make composters from two used buckets [citation:4][citation:6]
- Bokoharjo, Sleman: Modified plastic drums with aeration holes and taps [citation:5]
- Tlogotirto, Grobogan: UNDIP's anaerobic RE-DRUM system [citation:9]

**💰 Savings:**
- DIY cost: ~Rp 150,000 + your buckets
- Comparable ready-made: Rp 350,000
- Savings: Rp 200,000 (57% cheaper!)

**💚 Environmental Impact:**
- Reuses existing buckets (zero waste!)
- Same waste reduction as expensive units
- Perfect for community workshops`,
        IsPublished: true,
        Tags:       pq.StringArray{"DIY", "rakit sendiri", "murah", "ekonomis", "ember bekas", "workshop"},
        Slug:       "diy-composter-kit-rakit-sendiri-ember-bekas",
        MainImageURL: "https://urbankomposter.com/wp-content/uploads/2020/04/urban-komposter-17-liter.png",
        MainImagePublicID: "diy-composter-kit",
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }
    
    variants := []struct {
        Material string
        Price    float64
        Weight   float64
        Stock    int
    }{
        {"Basic Kit (Aerobic)", 125000, 800, 300},
        {"Advanced Kit (Anaerobic RE-DRUM)", 175000, 1000, 150},
        {"Community Workshop Pack (10 kits)", 1100000, 9000, 20},
    }
    
    return createComposterSKUs(db, product, variants, "DIY")
}

// Product 6: Composting Starter Kits & Accessories
func seedCompostingStarterKits(db *gorm.DB, categoryID uint) error {
    product := entity.Product{
        CategoryID: categoryID,
        Name:       "Starter Kit Komposter - Aktivator & Perlengkapan",
        Description: `Everything you need to maintain your composter. From EM4 activator to composting tools and educational materials.

**🧪 Activator Options:**
- **EM4 (Effective Microorganisms 4)**: The standard bio-activator used in Indonesian composting programs [citation:4]
- **Molase (Molasses)**: Food source for beneficial bacteria
- **Complete Bio-Starter**: EM4 + molasses + instructions

**🛠️ Accessories:**
- Compost thermometers
- Aeration tools
- Collection containers
- Educational books
- Compost sifters

**📚 Educational Materials:**
- "Panduan Praktis Membuat Kompos" booklet
- Video tutorials (QR code)
- Community workshop guides

**💚 Perfect For:**
- New composters
- Existing composter owners
- Community program supplies
- School environmental projects`,
        IsPublished: true,
        Tags:       pq.StringArray{"starter kit", "aktivator", "EM4", "molase", "perlengkapan", "buku panduan"},
        Slug:       "starter-kit-aktivator-perlengkapan-komposter",
        MainImageURL: "https://jubelio-store.s3.ap-southeast-1.amazonaws.com/sustaination/2021/03/01214020/Komposter-Set-Sekop-3-scaled.webp",
        MainImagePublicID: "composter-accessories",
    }
    
    if err := db.Create(&product).Error; err != nil {
        return err
    }
    
    variants := []struct {
        Material string
        Price    float64
        Weight   float64
        Stock    int
    }{
        {"EM4 Activator (500ml)", 35000, 500, 500},
        {"Molase / Molasses (1L)", 25000, 1000, 400},
        {"Bio-Starter Complete Kit", 75000, 1500, 300},
        {"Compost Thermometer", 125000, 300, 100},
        {"Complete Accessory Set", 225000, 2500, 75},
    }
    
    return createComposterSKUs(db, product, variants, "ACC")
}

// Helper function to create SKUs for composter products
func createComposterSKUs(db *gorm.DB, product entity.Product, variants []struct {
    Material string
    Price    float64
    Weight   float64
    Stock    int
}, prefix string) error {
    
    // Create variant type
    variantType := entity.VariantType{
        ProductID: product.ID,
        Name:      "Model / Material",
    }
    if err := db.Create(&variantType).Error; err != nil {
        return err
    }
    
    var skus []entity.SKU
    var minPrice, maxPrice = variants[0].Price, variants[0].Price
    
    for _, v := range variants {
        // Create variant value
        variantValue := entity.VariantValue{
            VariantTypeID: variantType.ID,
            Value:         v.Material,
        }
        if err := db.Create(&variantValue).Error; err != nil {
            return err
        }

        str, _ := utils.GenerateRandomString(5)
        
        // Generate SKU code
        skuCode := fmt.Sprintf("%s-%s-%03d-%s", prefix, product.Slug[:4], rand.Intn(1000), str)
        
        // 30% chance of sale price
        var salePrice *float64
        if rand.Float64() < 0.3 {
            sp := v.Price * 0.9 // 10% off
            sp = float64(int(sp/1000)) * 1000
            salePrice = &sp
        }
        
        // Create SKU
        sku := entity.SKU{
            ProductID: product.ID,
            SKUCode:   skuCode,
            Price:     v.Price,
            SalePrice: salePrice,
            Stock:     v.Stock,
            MinStock:  int(float64(v.Stock) * 0.2), // 20% of initial stock
            Status:    entity.SKUStatusActive,
            Weight:    v.Weight,
        }
        
        if err := db.Create(&sku).Error; err != nil {
            return err
        }
        
        // Link SKU to variant
        skuVariant := entity.SKUVariantValue{
            SKUID:          sku.ID,
            VariantValueID: variantValue.ID,
        }
        db.Create(&skuVariant)
        
        skus = append(skus, sku)
        
        // Update price range
        if v.Price < minPrice {
            minPrice = v.Price
        }
        if v.Price > maxPrice {
            maxPrice = v.Price
        }
    }
    
    // Update product prices
    product.MinPrice = minPrice
    product.MaxPrice = maxPrice
    db.Save(&product)
    
    // Add images
    addComposterImages(db, product.ID)
    
    fmt.Printf("  ✅ Created %s: %d variants (Rp %.0f - Rp %.0f)\n", 
        product.Name, len(skus), minPrice, maxPrice)
    
    return nil
}

func addComposterImages(db *gorm.DB, productID uint) error {
	// Upload images
	imageURLs := []string{
		"https://jubelio-store.s3.ap-southeast-1.amazonaws.com/sustaination/2021/02/31233615/Komposter-Kit-GMI-1-scaled.webp",
		"https://urbankomposter.com/wp-content/uploads/2020/04/urban-komposter-17-liter.png",
		"https://www.scdprobiotics.com/cdn/shop/products/all-seasons-indoor-composter.jpg?v=1687467692",
	}

	for _, url := range imageURLs {
		img := entity.Image{
			ProductID: productID,
			ImageURL: url,
		}

		if err := db.Create(&img).Error; err != nil {
			return err
		}
	}

	return nil
}