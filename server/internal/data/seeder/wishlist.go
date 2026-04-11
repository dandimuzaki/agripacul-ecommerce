package seeder

import (
	"debian-ecommerce/internal/data/entity"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

func SeedWishlists(db *gorm.DB) error {
    rand.Seed(time.Now().UnixNano())
    
    // Get all customers
    var customers []entity.Customer
    if err := db.Find(&customers).Error; err != nil {
        return fmt.Errorf("failed to get customers: %w", err)
    }
    
    // Get all published products
    var products []entity.Product
    if err := db.Where("is_published = ?", true).Find(&products).Error; err != nil {
        return fmt.Errorf("failed to get products: %w", err)
    }
    
    if len(products) == 0 {
        return fmt.Errorf("no products found to add to wishlist")
    }
    
    var allWishlists []entity.Wishlist
    var wishlistCount int
    
    // For each customer, add 5-15 items to wishlist
    for _, customer := range customers {
        // Random number of wishlist items per customer (5-15)
        numItems := rand.Intn(11) + 5 // 5 to 15 items
        
        // Shuffle products to get random selection
        shuffledProducts := shuffleProducts(products)
        
        // Track added product IDs to avoid duplicates
        addedProductIDs := make(map[uint]bool)
        
        itemsAdded := 0
        for i := 0; i < len(shuffledProducts) && itemsAdded < numItems; i++ {
            product := shuffledProducts[i]
            
            // Skip if already in wishlist
            if addedProductIDs[product.ID] {
                continue
            }
            
            // Check if this customer already has this product in wishlist
            var existingWishlist entity.Wishlist
            err := db.Where("customer_id = ? AND product_id = ?", customer.ID, product.ID).First(&existingWishlist).Error
            if err == nil {
                // Already exists, skip
                continue
            }
            
            // Create wishlist item with random creation date (1-90 days ago)
            daysAgo := rand.Intn(90) + 1
            createdAt := time.Now().AddDate(0, 0, -daysAgo)
            
            wishlist := entity.Wishlist{
                CustomerID: customer.ID,
                ProductID:  product.ID,
                Model: gorm.Model{
                    CreatedAt: createdAt,
                    UpdatedAt: createdAt,
                },
            }
            
            if err := db.Create(&wishlist).Error; err != nil {
                return fmt.Errorf("failed to create wishlist: %w", err)
            }
            
            allWishlists = append(allWishlists, wishlist)
            addedProductIDs[product.ID] = true
            itemsAdded++
            wishlistCount++
        }
        
        // Add some products that are out of stock (to simulate "notify when available")
        outOfStockProducts := getOutOfStockProducts(db)
        for i := 0; i < len(outOfStockProducts) && i < 3; i++ {
            product := outOfStockProducts[i]
            
            if addedProductIDs[product.ID] {
                continue
            }
            
            // Check if already in wishlist
            var existingWishlist entity.Wishlist
            err := db.Where("customer_id = ? AND product_id = ?", customer.ID, product.ID).First(&existingWishlist).Error
            if err == nil {
                continue
            }
            
            createdAt := time.Now().AddDate(0, 0, -rand.Intn(30))
            
            wishlist := entity.Wishlist{
                CustomerID: customer.ID,
                ProductID:  product.ID,
                Model: gorm.Model{
                    CreatedAt: createdAt,
                    UpdatedAt: createdAt,
                },
            }
            
            if err := db.Create(&wishlist).Error; err != nil {
                return fmt.Errorf("failed to create wishlist for out-of-stock product: %w", err)
            }
            
            allWishlists = append(allWishlists, wishlist)
            addedProductIDs[product.ID] = true
            wishlistCount++
        }
    }
    
    fmt.Printf("✅ Successfully seeded %d wishlist items for %d customers\n", wishlistCount, len(customers))
    return nil
}

// Helper to shuffle products for random selection
func shuffleProducts(products []entity.Product) []entity.Product {
    shuffled := make([]entity.Product, len(products))
    copy(shuffled, products)
    rand.Shuffle(len(shuffled), func(i, j int) {
        shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
    })
    return shuffled
}

// Helper to get products that are out of stock
func getOutOfStockProducts(db *gorm.DB) []entity.Product {
    var products []entity.Product
    
    // Find products where all SKUs have zero stock
    db.Joins("JOIN skus ON skus.product_id = products.id").
        Where("products.is_published = ?", true).
        Group("products.id").
        Having("SUM(skus.stock) = 0").
        Find(&products)
    
    return products
}

// Alternative: Seed with specific wishlist items for realistic scenarios
func SeedWishlistsWithPreferences(db *gorm.DB) error {
    // This function creates more targeted wishlists based on customer preferences
    
    // Customer preferences mapping
    customerPreferences := map[uint][]string{
        1: {"Zero Waste Kit", "Personal Care", "Stainless"}, // Budi - practical items
        2: {"Bamboo", "Kitchen", "Eco-friendly"},            // Siti - kitchen & bamboo
        3: {"Batik", "Premium", "Gift"},                     // Dewi - premium & gifts
    }
    
    var allWishlists []entity.Wishlist
    
    for customerID, preferences := range customerPreferences {
        for _, pref := range preferences {
            // Find products matching preference
            var products []entity.Product
            db.Where("is_published = ? AND (name ILIKE ? OR tags @> ARRAY[?]::text[])", 
                true, "%"+pref+"%", pref).
                Limit(5).
                Find(&products)
            
            for _, product := range products {
                // Check if already in wishlist
                var existingWishlist entity.Wishlist
                err := db.Where("customer_id = ? AND product_id = ?", customerID, product.ID).First(&existingWishlist).Error
                if err == nil {
                    continue
                }
                
                // Create wishlist with realistic date (evening/weekend browsing)
                createdAt := generateRealisticWishlistDate()
                
                wishlist := entity.Wishlist{
                    CustomerID: customerID,
                    ProductID:  product.ID,
                    Model: gorm.Model{
                        CreatedAt: createdAt,
                        UpdatedAt: createdAt,
                    },
                }
                
                if err := db.Create(&wishlist).Error; err != nil {
                    return err
                }
                
                allWishlists = append(allWishlists, wishlist)
            }
        }
    }
    
    fmt.Printf("✅ Successfully seeded %d preference-based wishlist items\n", len(allWishlists))
    return nil
}

// Generate realistic wishlist dates (evenings and weekends when people browse more)
func generateRealisticWishlistDate() time.Time {
    now := time.Now()
    
    daysAgo := rand.Intn(90) + 1
    date := now.AddDate(0, 0, -daysAgo)
    
    // Make it more likely to be evening (after work) or weekend
    hour := rand.Intn(24)
    weekday := date.Weekday()
    
    // If it's a weekday, bias toward evening hours (6-11 PM)
    if weekday >= time.Monday && weekday <= time.Friday {
        if rand.Float64() < 0.6 {
            // 60% chance of evening for weekdays
            hour = rand.Intn(6) + 18 // 6-11 PM
        }
    } else {
        // Weekend - more spread out but still bias to daytime/evening
        if rand.Float64() < 0.8 {
            hour = rand.Intn(12) + 10 // 10 AM - 10 PM
        }
    }
    
    minute := rand.Intn(60)
    second := rand.Intn(60)
    
    return time.Date(
        date.Year(), date.Month(), date.Day(),
        hour, minute, second, 0,
        date.Location(),
    )
}