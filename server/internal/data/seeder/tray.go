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

func SeedTrayCollection(db *gorm.DB) error {
	rand.Seed(time.Now().UnixNano())
	
	// First, create specialized categories
	category, err := createFoodPackagingCategory(db)
	if err != nil {
		return err
	}
    
	// Seed product lines (6 products instead of 1 mega-product)
	products := []struct {
		name     string
		category uint
		fn       func(*gorm.DB, uint) error
	}{
		{"Bagasse Sugarcane Trays", category.ID, seedBagasseTrayLine},
		{"Wheat Straw Trays", category.ID, seedWheatStrawTrayLine},
		{"Bamboo Premium Trays", category.ID, seedBambooTrayLine},
		{"Home & Small Business Packs", category.ID, seedHomeTrayPacks},
		{"Institutional & Government", category.ID, seedInstitutionalTrays},
		{"Sample & Trial Packs", category.ID, seedSamplePacks},
	}
	
	for _, p := range products {
		if err := p.fn(db, p.category); err != nil {
			return fmt.Errorf("failed to seed %s: %v", p.name, err)
		}
	}
	
	fmt.Println("✅ Successfully seeded Compostable Tray Collection (6 product lines)")
	return nil
}

// ==================== PRODUCT LINE 1: BAGASSE TRAYS ====================
// Target: Restaurants, food delivery, cafes
// Based on market data: Most popular material, best price-performance [citation:3]

func seedBagasseTrayLine(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Bagasse Food Trays - 3 Compartment, Cutting Resistant",
		Description:  `Made from sugarcane bagasse, a byproduct of sugar production. Our best-selling trays for restaurants and food delivery in Indonesia.

**🌱 Why Bagasse?**
- Made from agricultural waste (no trees cut)
- Composts in 60-90 days commercially
- Grease-resistant without plastic lining
- BPOM certified for food contact

**✨ Key Features:**
- ✅ Cutting-resistant technology
- ✅ 3-compartment design (nasi, lauk, sayur)
- ✅ Microwave & freezer safe
- ✅ No PFAS or plastic lining
- ✅ Heavy-duty 3.5mm thickness

**🎯 Perfect For:**
- Restaurants & cafes
- GoFood/Grab merchants
- Catering events
- Office lunches

**📦 Pack Sizes:**
- 25 pcs (Home trial)
- 100 pcs (Restaurant starter)
- 500 pcs (Bulk wholesale)
- 1000 pcs (Institutional)`,
		IsPublished:      true,
		Tags:             pq.StringArray{"bagasse", "sugarcane", "ampas tebu", "food tray", "3 compartment", "cutting resistant", "restaurant", "delivery", "gofood", "grab"},
		Slug:             "bagasse-food-trays-3-compartment-cutting-resistant",
		MainImageURL:     "https://ecolipak.com/cdn/shop/files/14-in-disposable-cutting-resistant-compostable-traysbpi-certified-pfas-free-6290227.jpg?v=1762583270&width=1024",
		MainImagePublicID: "bagasse-trays-main",
		AverageRating:    4.7,
		ReviewCount:      245,
		SoldCount:        1250,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	// Create variant type (just pack size - much simpler!)
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Pack Size",
	}
	db.Create(&variantType)
	
	// Bagasse tray variants - real Indonesian e-commerce pricing
	variants := []struct {
		Name        string
		PackSize    int
		Price       float64
		Weight      float64
		Stock       int
		MinStock    int
		Description string
	}{
		{
			Name:        "Home Pack - 25 trays",
			PackSize:    25,
			Price:       85000,
			Weight:      1250,
			Stock:       200,
			MinStock:    20,
			Description: "Cocok untuk coba-coba di rumah",
		},
		{
			Name:        "Restaurant Starter - 100 trays",
			PackSize:    100,
			Price:       299000,
			Weight:      5000,
			Stock:       100,
			MinStock:    15,
			Description: "Untuk warung atau restoran kecil",
		},
		{
			Name:        "Business Pack - 250 trays",
			PackSize:    250,
			Price:       699000,
			Weight:      12500,
			Stock:       50,
			MinStock:    10,
			Description: "Untuk katering dan restoran besar",
		},
		{
			Name:        "Bulk Wholesale - 500 trays",
			PackSize:    500,
			Price:       1_299_000,
			Weight:      25000,
			Stock:       25,
			MinStock:    5,
			Description: "Harga grosir untuk bisnis mapan",
		},
		{
			Name:        "Institutional - 1000 trays",
			PackSize:    1000,
			Price:       2_450_000,
			Weight:      50000,
			Stock:       15,
			MinStock:    2,
			Description: "Untuk program sekolah atau kantin",
		},
	}
	
	for _, v := range variants {
		// Create variant value
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Name,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		// Generate SKU code
		skuCode := fmt.Sprintf("BGS-%d-%03d-%s", v.PackSize, rand.Intn(1000), str)
		
		// 20% chance of sale price
		var salePrice *float64
		if rand.Float64() < 0.2 {
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
			MinStock:  v.MinStock,
			Status:    entity.SKUStatusActive,
			Weight:    v.Weight,
		}
		db.Create(&sku)
		
		// Link SKU to variant
		db.Create(&entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: value.ID,
		})
	}
	
	// Update product price range
	product.MinPrice = 85000
	product.MaxPrice = 2450000
	db.Save(&product)
	
	if err := addTrayImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Bagasse Trays: %d variants (Rp 85rb - 2.45jt)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 2: WHEAT STRAW TRAYS ====================
// Target: Premium catering, events, hotels

func seedWheatStrawTrayLine(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Wheat Straw Premium Trays - Natural White, Extra Strong",
		Description: `Made from wheat straw fiber, a natural byproduct of wheat harvesting. These trays offer a premium look with natural off-white color and superior strength.

**🌾 Why Wheat Straw?**
- Agricultural waste that would otherwise be burned
- Naturally lighter color needs less processing
- Higher tensile strength than bagasse
- Premium appearance for upscale events

**✨ Features:**
- ✅ Natural off-white color
- ✅ Extra rigid construction
- ✅ Cut-resistant surface
- ✅ 3-compartment design
- ✅ Microwave safe up to 3 minutes

**🎯 Perfect For:**
- Wedding catering
- Hotel room service
- Corporate events
- Premium restaurants`,
		IsPublished:      true,
		Tags:             pq.StringArray{"wheat straw", "jerami gandum", "premium", "catering", "natural white", "event"},
		Slug:             "wheat-straw-premium-trays-natural-white",
		MainImageURL:     "https://ecolipak.com/cdn/shop/files/14-in-disposable-cutting-resistant-compostable-traysbpi-certified-pfas-free-4713263.jpg?v=1762583270&width=1024",
		MainImagePublicID: "wheat-straw-trays-main",
		AverageRating:    4.8,
		ReviewCount:      98,
		SoldCount:        456,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Pack Size",
	}
	db.Create(&variantType)
	
	variants := []struct {
		Name     string
		PackSize int
		Price    float64
		Weight   float64
		Stock    int
		MinStock int
	}{
		{"Sample - 25 trays", 25, 110000, 1300, 100, 15},
		{"Catering Pack - 100 trays", 100, 399000, 5200, 60, 10},
		{"Event Pack - 250 trays", 250, 949000, 13000, 30, 5},
		{"Wedding Special - 500 trays", 500, 1_799_000, 26000, 15, 3},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Name,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("WHT-%d-%03d-%s", v.PackSize, rand.Intn(1000), str)
		
		sku := entity.SKU{
			ProductID: product.ID,
			SKUCode:   skuCode,
			Price:     v.Price,
			Stock:     v.Stock,
			MinStock:  v.MinStock,
			Status:    entity.SKUStatusActive,
			Weight:    v.Weight,
		}
		db.Create(&sku)
		
		db.Create(&entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: value.ID,
		})
	}
	
	product.MinPrice = 110000
	product.MaxPrice = 1799000
	db.Save(&product)
	
	if err := addTrayImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Wheat Straw Trays: %d variants (Rp 110rb - 1.8jt)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 3: BAMBOO PREMIUM TRAYS ====================
// Target: Eco-conscious brands, organic restaurants

func seedBambooTrayLine(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Bamboo Fiber Premium Trays - Antibacterial, Luxury",
		Description: `Our most premium option. Bamboo fiber trays offer natural antibacterial properties, superior strength, and the fastest renewability of any material.

**🎋 Why Bamboo?**
- Fastest growing plant (harvested in 3-5 years)
- Naturally antibacterial without chemicals
- Strongest fiber of all options
- Beautiful natural texture

**✨ Features:**
- ✅ Natural antibacterial properties
- ✅ Highest cut resistance
- ✅ Beautiful natural grain
- ✅ Available in 3 and 5-compartment
- ✅ Premium packaging option

**🎯 Perfect For:**
- Organic restaurants
- Health-conscious consumers
- Luxury catering
- Premium meal delivery`,
		IsPublished:      true,
		Tags:             pq.StringArray{"bamboo", "bambu", "premium", "antibacterial", "luxury", "organic"},
		Slug:             "bamboo-fiber-premium-trays-antibacterial",
		MainImageURL:     "https://ecolipak.com/cdn/shop/files/14-in-disposable-cutting-resistant-compostable-traysbpi-certified-pfas-free-6819939.jpg?v=1762583269&width=1024",
		MainImagePublicID: "bamboo-premium-trays",
		AverageRating:    4.9,
		ReviewCount:      67,
		SoldCount:        234,
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Pack Size",
	}
	db.Create(&variantType)
	
	variants := []struct {
		Name     string
		PackSize int
		Price    float64
		Weight   float64
		Stock    int
		MinStock int
	}{
		{"Trial Pack - 25 trays", 25, 165000, 1500, 50, 10},
		{"Restaurant Pack - 100 trays", 100, 599000, 6000, 25, 5},
		{"Premium Catering - 250 trays", 250, 1_425_000, 15000, 10, 3},
		{"5-Compartment Tiffin - 50 pcs", 50, 425000, 4000, 15, 4},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Name,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("BMB-%d-%03d-%s", v.PackSize, rand.Intn(1000), str)
		
		sku := entity.SKU{
			ProductID: product.ID,
			SKUCode:   skuCode,
			Price:     v.Price,
			Stock:     v.Stock,
			MinStock:  v.MinStock,
			Status:    entity.SKUStatusActive,
			Weight:    v.Weight,
		}
		db.Create(&sku)
		
		db.Create(&entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: value.ID,
		})
	}
	
	product.MinPrice = 165000
	product.MaxPrice = 1425000
	db.Save(&product)
	
	if err := addTrayImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Bamboo Premium Trays: %d variants (Rp 165rb - 1.43jt)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 4: HOME PACKS ====================
// Target: Individual households, small families
// Based on affordable eco-products for students and families [citation:2]

func seedHomeTrayPacks(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Home Pack - 25 Tray Ramah Lingkungan untuk Keluarga",
		Description: `Paket ekonomis untuk keluarga Indonesia yang ingin mulai beralih ke produk ramah lingkungan. Cocok untuk bekal sekolah, arisan, dan acara keluarga.

**🏠 Untuk Keluarga Indonesia:**
- Ukuran pas untuk porsi makan keluarga
- Harga terjangkau mulai Rp 35.000
- Bisa untuk microwave (panasin lauk)
- Aman untuk anak-anak

**📦 Isi Paket:**
- 25 tray 3-compartment
- Bisa pilih bahan: Bagasse atau Bamboo

**💚 Dampak Lingkungan:**
Dengan menggunakan tray ini, keluarga Anda berkontribusi mengurangi sampah plastik sekali pakai. Produk ini terurai dalam 60-90 hari di fasilitas kompos, berbeda dengan plastik yang butuh 500+ tahun.`,
		IsPublished:      true,
		Tags:             pq.StringArray{"home", "keluarga", "rumah tangga", "bekal", "sekolah", "arisan", "murah"},
		Slug:             "home-pack-tray-ramah-lingkungan-keluarga",
		MainImageURL:     "https://ecolipak.com/cdn/shop/files/14-in-disposable-cutting-resistant-compostable-traysbpi-certified-pfas-free-7092684.jpg?v=1762583270&width=1024",
		MainImagePublicID: "home-pack-trays",
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	// For home packs, we use material as the variant
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Material",
	}
	db.Create(&variantType)
	
	variants := []struct {
		Material string
		Price    float64
		Weight   float64
		Stock    int
		MinStock int
	}{
		{"Bagasse (Tebu) - 25 pcs", 85000, 1250, 300, 30},
		{"Wheat Straw - 25 pcs", 110000, 1300, 150, 20},
		{"Bamboo Premium - 25 pcs", 165000, 1500, 75, 10},
	}
	
	for _, v := range variants {
		value := entity.VariantValue{
			VariantTypeID: variantType.ID,
			Value:         v.Material,
		}
		db.Create(&value)

		str, _ := utils.GenerateRandomString(5)
		
		skuCode := fmt.Sprintf("HOME-%s-%03d-%s", v.Material[:3], rand.Intn(1000), str)
		
		sku := entity.SKU{
			ProductID: product.ID,
			SKUCode:   skuCode,
			Price:     v.Price,
			Stock:     v.Stock,
			MinStock:  v.MinStock,
			Status:    entity.SKUStatusActive,
			Weight:    v.Weight,
		}
		db.Create(&sku)
		
		db.Create(&entity.SKUVariantValue{
			SKUID:          sku.ID,
			VariantValueID: value.ID,
		})
	}
	
	product.MinPrice = 85000
	product.MaxPrice = 165000
	db.Save(&product)
	
	if err := addTrayImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Home Packs: %d variants (Rp 85rb - 165rb)\n", len(variants))
	return nil
}

// ==================== PRODUCT LINE 5: INSTITUTIONAL TRAYS ====================
// Target: Government programs, schools, hospitals
// Based on "Makan Bergizi Gratis" program at 30,000+ locations [citation:1]

func seedInstitutionalTrays(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Institutional Trays - Program Makan Bergizi Gratis & Sekolah",
		Description: `Dirancang khusus untuk program pemerintah dan institusi. Memenuhi spesifikasi teknis untuk program Makan Bergizi Gratis yang beroperasi di 30.000+ titik di seluruh Indonesia [citation:1].

**🇮🇩 Untuk Program Nasional:**
Mendukung inisiatif pemerintah menyediakan makanan bergizi untuk siswa Indonesia sambil membangun kebiasaan ramah lingkungan sejak dini.

**📋 Spesifikasi Institusional:**
- ✅ 3-compartment (nasi, protein, sayur)
- ✅ Cutting-resistant (aman untuk anak)
- ✅ Microwave-safe untuk pemanasan
- ✅ 100% compostable
- ✅ BPOM certified

**📦 Opsi Pengadaan:**
- School Pack (1000 tray) - per sekolah
- Sub-District (5000 tray) - per kecamatan
- District Container (50,000+) - per kabupaten/kota

**📞 Harga Khusus:**
Hubungi kami untuk tender dan pengadaan pemerintah. Kami mendukung e-katalog LKPP.`,
		IsPublished:      true,
		Tags:             pq.StringArray{"institutional", "government", "sekolah", "program", "makan bergizi gratis", "BPOM", "LKPP"},
		Slug:             "institutional-trays-program-makan-bergizi-gratis",
		MainImageURL:     "https://ecolipak.com/cdn/shop/files/14-in-disposable-cutting-resistant-compostable-traysbpi-certified-pfas-free-6290227.jpg?v=1762583270&width=1024",
		MainImagePublicID: "institutional-trays",
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	// No variants - each is a separate SKU with institutional pricing
	institutionalSKUs := []struct {
		Name     string
		PackSize int
		Price    float64
		Weight   float64
		Stock    int
		MinStock int
	}{
		{
			Name:     "School Pack - 1000 trays (Bagasse)",
			PackSize: 1000,
			Price:    2_450_000,
			Weight:   50000,
			Stock:    50,
			MinStock: 5,
		},
		{
			Name:     "Sub-District Pallet - 5000 trays",
			PackSize: 5000,
			Price:    11_375_000,
			Weight:   250000,
			Stock:    20,
			MinStock: 2,
		},
		{
			Name:     "District Container - 50,000 trays",
			PackSize: 50000,
			Price:    105_000_000,
			Weight:   2500000,
			Stock:    5,
			MinStock: 1,
		},
	}
	
	for _, s := range institutionalSKUs {
		str, _ := utils.GenerateRandomString(5)

		sku := entity.SKU{
			ProductID: product.ID,
			SKUCode:   fmt.Sprintf("GOV-%dK-%03d-%s", s.PackSize/1000, rand.Intn(100), str),
			Price:     s.Price,
			Stock:     s.Stock,
			MinStock:  s.MinStock,
			Status:    entity.SKUStatusActive,
			Weight:    s.Weight,
		}
		db.Create(&sku)
	}
	
	product.MinPrice = 2450000
	product.MaxPrice = 105000000
	product.SoldCount = 15 // Institutional sales are lower volume
	db.Save(&product)
	
	if err := addTrayImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Institutional Trays: %d SKU (Rp 2.45jt - 105jt)\n", len(institutionalSKUs))
	return nil
}

// ==================== PRODUCT LINE 6: SAMPLE PACKS ====================
// Target: Customers wanting to try before buying bulk

func seedSamplePacks(db *gorm.DB, categoryID uint) error {
	product := entity.Product{
		CategoryID:   categoryID,
		Name:         "Sample Pack - Coba Semua Varian Tray Ramah Lingkungan",
		Description: `Mau coba dulu sebelum beli banyak? Paket sampel ini cocok untuk Anda. Dapatkan 10 tray dengan berbagai material untuk menemukan yang paling cocok.

**🎁 Isi Paket:**
- 4 tray Bagasse (tebu)
- 3 tray Wheat Straw (gandum)
- 3 tray Bamboo (bambu)

**📋 Cocok Untuk:**
- Pemilik rumah yang baru mulai
- Pemilik kafe yang mau ganti kemasan
- Event organiser yang mau coba produk
- Sekolah yang mau program edukasi

Harga spesial untuk trial!`,
		IsPublished:      true,
		Tags:             pq.StringArray{"sample", "trial", "coba", "starter", "mixed", "varian"},
		Slug:             "sample-pack-coba-semua-varian-tray",
		MainImageURL:     "https://ecolipak.com/cdn/shop/files/14-in-disposable-cutting-resistant-compostable-traysbpi-certified-pfas-free-4713263.jpg?v=1762583270&width=1024",
		MainImagePublicID: "sample-pack-trays",
	}
	
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	
	// Just one variant - the sampler pack
	variantType := entity.VariantType{
		ProductID: product.ID,
		Name:      "Sample Type",
	}
	db.Create(&variantType)
	
	value := entity.VariantValue{
		VariantTypeID: variantType.ID,
		Value:         "Mixed Material Sampler - 10 trays",
	}
	db.Create(&value)

	str, _ := utils.GenerateRandomString(5)
	
	sku := entity.SKU{
		ProductID: product.ID,
		SKUCode:   fmt.Sprintf("SMP-MIX-%03d-%s", rand.Intn(1000), str),
		Price:     65000,
		Stock:     150,
		MinStock:  20,
		Status:    entity.SKUStatusActive,
		Weight:    550,
	}
	db.Create(&sku)
	
	db.Create(&entity.SKUVariantValue{
		SKUID:          sku.ID,
		VariantValueID: value.ID,
	})
	
	product.MinPrice = 65000
	product.MaxPrice = 65000
	db.Save(&product)
	
	if err := addTrayImages(db, product.ID); err != nil {
		return err
	}
	
	fmt.Printf("  ✅ Sample Packs: 1 SKU (Rp 65rb)\n")
	return nil
}

func createFoodPackagingCategory(db *gorm.DB) (*entity.Category, error) {
  var category entity.Category
	result := db.First(&category, 5)
	if result.Error != nil {
		// Create category if it doesn't exist
		category = entity.Category{
			Model:       gorm.Model{ID: 5},
			Name:        "Food Packaging",
		}
		if err := db.Create(&category).Error; err != nil {
			return nil, err
		}
	}

	return &category, nil
}

func addTrayImages(db *gorm.DB, productID uint) error {
	imageURLs := []string{
		"https://ecolipak.com/cdn/shop/files/14-in-disposable-cutting-resistant-compostable-traysbpi-certified-pfas-free-6290227.jpg?v=1762583270&width=1024",
		"https://ecolipak.com/cdn/shop/files/14-in-disposable-cutting-resistant-compostable-traysbpi-certified-pfas-free-4713263.jpg?v=1762583270&width=1024",
		"https://ecolipak.com/cdn/shop/files/14-in-disposable-cutting-resistant-compostable-traysbpi-certified-pfas-free-6819939.jpg?v=1762583269&width=1024",
		"https://ecolipak.com/cdn/shop/files/14-in-disposable-cutting-resistant-compostable-traysbpi-certified-pfas-free-7092684.jpg?v=1762583270&width=1024",
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