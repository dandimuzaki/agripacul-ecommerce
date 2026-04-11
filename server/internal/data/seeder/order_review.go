package seeder

import (
	"debian-ecommerce/internal/data/entity"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func SeedOrdersAndReviews(db *gorm.DB) error {
    rand.Seed(time.Now().UnixNano())
    
    // First, ensure we have customers (assuming users exist)
    // If not, create 3 test customers
    var customers []entity.Customer
    if err := db.Find(&customers).Error; err != nil {
        return err
    }

    // Get all SKUs for order items
    var skus []entity.SKU
    if err := db.Find(&skus).Error; err != nil {
        return err
    }

    // Check if SKUs exist
    var skuCount int64
    db.Model(&entity.SKU{}).Count(&skuCount)
    if skuCount == 0 {
        fmt.Println("⚠️ No SKUs found in database. Please seed SKUs first")
        return errors.New("No SKU found")
    }
    fmt.Println("skuCount", skuCount)

    // Indonesian provinces, cities, and districts for realistic addresses
    indonesianAddresses := []map[string]string{
        {
            "province": "DKI Jakarta",
            "regency": "Jakarta Selatan",
            "district": "Kebayoran Baru",
            "subdistrict": "Senayan",
            "postal_code": "12190",
        },
        {
            "province": "DKI Jakarta",
            "regency": "Jakarta Timur",
            "district": "Jatinegara",
            "subdistrict": "Cipinang Besar Selatan",
            "postal_code": "13410",
        },
        {
            "province": "Jawa Barat",
            "regency": "Kota Bandung",
            "district": "Coblong",
            "subdistrict": "Dago",
            "postal_code": "40135",
        },
        {
            "province": "Jawa Barat",
            "regency": "Kota Bekasi",
            "district": "Bekasi Selatan",
            "subdistrict": "Pekayon Jaya",
            "postal_code": "17148",
        },
        {
            "province": "Jawa Tengah",
            "regency": "Kota Semarang",
            "district": "Semarang Tengah",
            "subdistrict": "Pekunden",
            "postal_code": "50134",
        },
        {
            "province": "DI Yogyakarta",
            "regency": "Kota Yogyakarta",
            "district": "Gondokusuman",
            "subdistrict": "Baciro",
            "postal_code": "55225",
        },
        {
            "province": "Jawa Timur",
            "regency": "Kota Surabaya",
            "district": "Gubeng",
            "subdistrict": "Kertajaya",
            "postal_code": "60282",
        },
        {
            "province": "Bali",
            "regency": "Kota Denpasar",
            "district": "Denpasar Selatan",
            "subdistrict": "Sanur",
            "postal_code": "80228",
        },
        {
            "province": "Sumatera Utara",
            "regency": "Kota Medan",
            "district": "Medan Petisah",
            "subdistrict": "Petisah Tengah",
            "postal_code": "20112",
        },
        {
            "province": "Sulawesi Selatan",
            "regency": "Kota Makassar",
            "district": "Panakkukang",
            "subdistrict": "Karampuang",
            "postal_code": "90231",
        },
    }

    // Courier options common in Indonesia
    courierOptions := []map[string]interface{}{
        {
            "name": "JNE",
            "code": "jne",
            "service": "REG",
            "cost": 15000.0,
            "etd": 3,
        },
        {
            "name": "JNE",
            "code": "jne",
            "service": "YES",
            "cost": 30000.0,
            "etd": 1,
        },
        {
            "name": "J&T Express",
            "code": "jnt",
            "service": "EZ",
            "cost": 14000.0,
            "etd": 3,
        },
        {
            "name": "SiCepat",
            "code": "sicepat",
            "service": "REG",
            "cost": 16000.0,
            "etd": 2,
        },
        {
            "name": "SiCepat",
            "code": "sicepat",
            "service": "BEST",
            "cost": 25000.0,
            "etd": 1,
        },
        {
            "name": "AnterAja",
            "code": "anteraja",
            "service": "REG",
            "cost": 13000.0,
            "etd": 4,
        },
        {
            "name": "GoSend",
            "code": "gosend",
            "service": "Same Day",
            "cost": 22000.0,
            "etd": 0,
        },
        {
            "name": "GrabExpress",
            "code": "grab",
            "service": "Instant",
            "cost": 24000.0,
            "etd": 0,
        },
    }

    // Payment methods common in Indonesia
    paymentMethods := []map[string]interface{}{
        {
            "id": 1,
            "name": "Bank Transfer BCA",
            "ref_prefix": "BCA",
        },
        {
            "id": 2,
            "name": "Bank Transfer Mandiri",
            "ref_prefix": "MDR",
        },
        {
            "id": 3,
            "name": "Bank Transfer BRI",
            "ref_prefix": "BRI",
        },
        {
            "id": 4,
            "name": "Bank Transfer BNI",
            "ref_prefix": "BNI",
        },
        {
            "id": 5,
            "name": "GoPay",
            "ref_prefix": "GOPAY",
        },
        {
            "id": 6,
            "name": "OVO",
            "ref_prefix": "OVO",
        },
        {
            "id": 7,
            "name": "DANA",
            "ref_prefix": "DANA",
        },
        {
            "id": 8,
            "name": "COD (Cash on Delivery)",
            "ref_prefix": "COD",
        },
    }

    // Generate 5 orders for each customer
    for _, customer := range customers {
        for orderNum := 1; orderNum <= 5; orderNum++ {
            // Random address from list
            addr := indonesianAddresses[rand.Intn(len(indonesianAddresses))]
            
            // Random courier
            courier := courierOptions[rand.Intn(len(courierOptions))]
            
            // Random payment method
            paymentMethod := paymentMethods[rand.Intn(len(paymentMethods))]
            
            // Order date between 1-90 days ago
            daysAgo := rand.Intn(90) + 1
            orderDate := time.Now().AddDate(0, 0, -daysAgo)
            
            // Select 1-4 random SKUs for this order
            numItems := rand.Intn(4) + 1
            selectedSKUs := selectRandomSKUs(skus, numItems)
            
            // Calculate subtotal
            subtotal := 0.0
            orderItems := []entity.OrderItem{}
            
            for _, sku := range selectedSKUs {
                quantity := rand.Intn(3) + 1 // 1-3 items per SKU
                
                // Use sale price if available, otherwise regular price
                unitPrice := sku.Price
                if sku.SalePrice != nil && *sku.SalePrice > 0 {
                    unitPrice = *sku.SalePrice
                }
                
                totalPrice := float64(quantity) * unitPrice
                subtotal += totalPrice
                
                // Create product and SKU snapshots
                var product entity.Product
                db.First(&product, sku.ProductID)
                
                productSnapshot := map[string]interface{}{
                    "id": product.ID,
                    "name": product.Name,
                    "slug": product.Slug,
                    "main_image_url": product.MainImageURL,
                    "category_id": product.CategoryID,
                }
                
                skuSnapshot := map[string]interface{}{
                    "id": sku.ID,
                    "sku_code": sku.SKUCode,
                    "price": sku.Price,
                    "sale_price": sku.SalePrice,
                    "weight_gram": sku.Weight,
                }
                
                orderItems = append(orderItems, entity.OrderItem{
                    SKUID: sku.ID,
                    Quantity: quantity,
                    UnitPrice: unitPrice,
                    TotalPrice: totalPrice,
                    ProductSnapshot: datatypes.JSON(marshalJSON(productSnapshot)),
                    SKUSnapshot: datatypes.JSON(marshalJSON(skuSnapshot)),
                })
            }
            
            // Shipping cost
            shippingCost := courier["cost"].(float64)
            
            // Apply random discount (0%, 5%, 10%, 15%, 20%)
            discountPercent := []float64{0, 5, 10, 15, 20}[rand.Intn(5)]
            discountAmount := subtotal * (discountPercent / 100)
            
            // Calculate total
            total := subtotal - discountAmount + shippingCost
            
            // Determine order status based on order date
            orderStatus := determineOrderStatus(orderDate, daysAgo)
            
            // Create order
            order := entity.Order{
                CustomerID: customer.ID,
                Status: orderStatus,
                Subtotal: subtotal,
                DiscountAmount: discountAmount,
                Total: total,
                Notes: generateOrderNotes(orderNum),
                Items: orderItems,
            }
            
            // Set lifecycle timestamps based on status
            if orderStatus == entity.OrderStatusCompleted || orderStatus == entity.OrderStatusCancelled {
                // For completed orders, set timestamps
                if orderStatus == entity.OrderStatusCompleted {
                    confirmedAt := orderDate.AddDate(0, 0, 1) // Confirmed next day
                    completedAt := orderDate.AddDate(0, 0, rand.Intn(5)+3) // Completed 3-7 days later
                    order.ConfirmedAt = &confirmedAt
                    order.CompletedAt = &completedAt
                } else if orderStatus == entity.OrderStatusCancelled {
                    // Cancelled orders
                    if rand.Float64() < 0.5 { // 50% cancelled by customer, 50% by system
                        cancelledBy := []string{"customer", "system"}[rand.Intn(2)]
                        cancelledAt := orderDate.AddDate(0, 0, rand.Intn(2)+1) // Cancelled 1-2 days later
                        order.CancelledAt = &cancelledAt
                        order.CancelledBy = cancelledBy
                        order.CancelReason = getCancelReason(cancelledBy)
                    }
                }
            }
            
            if err := db.Create(&order).Error; err != nil {
                return err
            }

            // var fixedItems []entity.OrderItem
            // for _, item := range orderItems {
            //     item.OrderID = order.ID
            //     fixedItems = append(fixedItems, item)
            // }
            // if err := db.Create(&fixedItems).Error; err != nil {
            //     return err
            // }
            
            // Create payment record
            paymentStatus := entity.PaymentStatusPending
            paidAt := &orderDate
            expiredAt := orderDate.AddDate(0, 0, 1) // Expires next day
            
            // 80% of orders are paid, 15% pending, 5% expired
            paymentRand := rand.Float64()
            if paymentRand < 0.8 {
                paymentStatus = entity.PaymentStatusPaid
                paidAt = &orderDate
            } else if paymentRand < 0.95 {
                paymentStatus = entity.PaymentStatusPending
                paidAt = nil
            } else {
                paymentStatus = entity.PaymentStatusExpired
                paidAt = nil
            }
            
            transactionRef := fmt.Sprintf("%s-%d-%d", paymentMethod["ref_prefix"], order.ID, rand.Intn(10000))
            
            payment := entity.OrderPayment{
                OrderID: order.ID,
                PaymentMethodID: uint(paymentMethod["id"].(int)),
                Amount: total,
                Status: paymentStatus,
                TransactionRef: transactionRef,
                PaidAt: paidAt,
                ExpiredAt: expiredAt,
            }
            
            if err := db.Create(&payment).Error; err != nil {
                return err
            }
            
            // Create shipping record
            shippingStatus := entity.ShippingStatusPending
            
            // If order is completed, shipping is delivered
            if orderStatus == entity.OrderStatusCompleted {
                shippingStatus = entity.ShippingStatusDelivered
            } else if orderStatus == entity.OrderStatusProcess {
                // Randomly shipped or still processing
                if rand.Float64() < 0.6 {
                    shippingStatus = entity.ShippingStatusShipped
                }
            }
            
            trackingNumber := ""
            shippedAt := (*time.Time)(nil)
            deliveredAt := (*time.Time)(nil)
            
            if shippingStatus == entity.ShippingStatusShipped || shippingStatus == entity.ShippingStatusDelivered {
                trackingNumber = generateTrackingNumber(courier["code"].(string), order.ID)
                shippedAtTime := orderDate.AddDate(0, 0, 1) // Shipped next day
                shippedAt = &shippedAtTime
                
                if shippingStatus == entity.ShippingStatusDelivered {
                    etd := courier["etd"].(int)
                    deliveredAtTime := shippedAtTime.AddDate(0, 0, etd+rand.Intn(2)) // ETD +/- 1 day
                    deliveredAt = &deliveredAtTime
                }
            }
            
            shipping := entity.OrderShipping{
                OrderID: order.ID,
                RecipientName: customer.FullName,
                Label: getAddressLabel(orderNum),
                PhoneNumber: customer.PhoneNumber,
                DetailAddress: fmt.Sprintf("Jl. Contoh No. %d, RT %02d/RW %02d", rand.Intn(100)+1, rand.Intn(20)+1, rand.Intn(10)+1),
                Province: addr["province"],
                Regency: addr["regency"],
                District: addr["district"],
                Subdistrict: addr["subdistrict"],
                PostalCode: addr["postal_code"],
                CourierName: courier["name"].(string),
                CourierCode: courier["code"].(string),
                CourierService: courier["service"].(string),
                Cost: shippingCost,
                ETD: courier["etd"].(int),
                Status: shippingStatus,
                TrackingNumber: trackingNumber,
                ShippedAt: shippedAt,
                DeliveredAt: deliveredAt,
            }
            
            if err := db.Create(&shipping).Error; err != nil {
                return err
            }
            
            // Create reviews for completed orders (only for some items)
            if orderStatus == entity.OrderStatusCompleted && rand.Float64() < 0.7 { // 70% of completed orders have reviews
                // Review 1-3 items from the order
                numReviews := rand.Intn(len(orderItems)) + 1
                if numReviews > 3 {
                    numReviews = 3
                }
                
                // Shuffle items and take first numReviews
                shuffledItems := shuffleOrderItems(orderItems)
                for i := 0; i < numReviews && i < len(shuffledItems); i++ {
                    item := shuffledItems[i]
                    
                    // Get product ID from snapshot
                    var productSnapshot map[string]interface{}
                    unmarshalJSON(item.ProductSnapshot, &productSnapshot)
                    productID := uint(productSnapshot["id"].(float64))
                    
                    // Create review
                    review := entity.Review{
                        CustomerID: customer.ID,
                        ProductID: productID,
                        OrderID: order.ID,
                        Rating: generateRating(),
                        Comment: generateReviewComment(generateRating()),
                    }
                    
                    if err := db.Create(&review).Error; err != nil {
                        return err
                    }
                    
                    // Update product average rating and review count
                    updateProductRatings(db, productID)
                }
            }
        }
    }
    
    fmt.Printf("✅ Successfully seeded orders and reviews for %d customers\n", len(customers))
    return nil
}

// Helper functions

func getCustomerName(i int) string {
    names := []string{
        "Budi Santoso",
        "Siti Rahayu",
        "Dewi Lestari",
        "Ahmad Hidayat",
        "Rina Wijaya",
        "Hendra Gunawan",
    }
    return names[(i-1)%len(names)]
}

func getCustomerPhone(i int) string {
    prefixes := []string{"0812", "0813", "0821", "0856", "0877", "0896"}
    prefix := prefixes[(i-1)%len(prefixes)]
    return fmt.Sprintf("%s%08d", prefix, rand.Intn(10000000))
}

func selectRandomSKUs(skus []entity.SKU, count int) []entity.SKU {
    if len(skus) == 0 {
        return []entity.SKU{}
    }
    
    // Shuffle
    shuffled := make([]entity.SKU, len(skus))
    copy(shuffled, skus)
    rand.Shuffle(len(shuffled), func(i, j int) {
        shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
    })
    
    // Take first count
    if count > len(shuffled) {
        count = len(shuffled)
    }
    return shuffled[:count]
}

func determineOrderStatus(orderDate time.Time, daysAgo int) entity.OrderStatus {
    // Random distribution: 60% completed, 20% processing, 15% created, 5% cancelled
    r := rand.Float64()
    
    if daysAgo > 30 { // Older orders more likely completed
        if r < 0.85 {
            return entity.OrderStatusCompleted
        } else if r < 0.95 {
            return entity.OrderStatusCancelled
        } else {
            return entity.OrderStatusProcess
        }
    } else if daysAgo > 7 { // 7-30 days old
        if r < 0.7 {
            return entity.OrderStatusCompleted
        } else if r < 0.85 {
            return entity.OrderStatusProcess
        } else if r < 0.95 {
            return entity.OrderStatusCreated
        } else {
            return entity.OrderStatusCancelled
        }
    } else { // Recent orders
        if r < 0.3 {
            return entity.OrderStatusCompleted
        } else if r < 0.6 {
            return entity.OrderStatusProcess
        } else if r < 0.85 {
            return entity.OrderStatusCreated
        } else {
            return entity.OrderStatusCancelled
        }
    }
}

func generateOrderNotes(orderNum int) string {
    notes := []string{
        "Tolong dibungkus rapi, untuk hadiah",
        "Jangan pakai plastik, saya zero waste",
        "Titipkan ke satpam jika tidak ada orang",
        "Saya ambil sendiri di toko",
        "Pisahkan produk yang mudah pecah",
        "Tambahkan bubble wrap",
        "Tolong diberi note kecil untuk hadiah",
        "",
        "",
    }
    return notes[rand.Intn(len(notes))]
}

func getAddressLabel(orderNum int) string {
    labels := []string{"Rumah", "Kantor", "Kos", "Apartemen"}
    return labels[orderNum%len(labels)]
}

func generateTrackingNumber(courierCode string, orderID uint) string {
    // Format: COURIER + DATE + ORDERID + RANDOM
    date := time.Now().Format("060102")
    random := rand.Intn(10000)
    return fmt.Sprintf("%s%s%d%04d", strings.ToUpper(courierCode), date, orderID, random)
}

func getCancelReason(cancelledBy string) string {
    if cancelledBy == "customer" {
        reasons := []string{
            "Berubah pikiran, tidak jadi beli",
            "Salah pilih varian produk",
            "Harga terlalu mahal",
            "Ditemukan harga lebih murah di tempat lain",
            "Ingin ganti metode pembayaran",
            "Tidak jadi, butuh uang untuk keperluan lain",
        }
        return reasons[rand.Intn(len(reasons))]
    } else {
        reasons := []string{
            "Pembayaran tidak diterima dalam batas waktu",
            "Stok produk habis",
            "Data pembeli tidak valid",
            "Melebihi batas waktu konfirmasi",
            "Transaksi mencurigakan",
        }
        return reasons[rand.Intn(len(reasons))]
    }
}

func generateRating() int {
    // Distribution: 50% 5-star, 25% 4-star, 15% 3-star, 7% 2-star, 3% 1-star
    r := rand.Float64()
    if r < 0.5 {
        return 5
    } else if r < 0.75 {
        return 4
    } else if r < 0.9 {
        return 3
    } else if r < 0.97 {
        return 2
    } else {
        return 1
    }
}

func generateReviewComment(rating int) string {
    if rating >= 5 {
        comments := []string{
            "Produk sangat bagus, sesuai deskripsi! Recomended!",
            "Pengiriman cepat, barang sesuai pesanan. Makasih",
            "Kualitas bahan premium, puas banget",
            "Sesuai ekspektasi, seller ramah, respon cepat",
            "Barang sampai dengan aman, packaging rapi. Terima kasih",
            "Mantap jiwa! Langsung saya pakai dan hasilnya bagus",
            "Sudah beli kedua kalinya, kualitas konsisten baik",
            "Recommended seller, produk original dan berkualitas",
        }
        return comments[rand.Intn(len(comments))]
    } else if rating >= 4 {
        comments := []string{
            "Produk bagus, tapi pengiriman agak lama",
            "Overall ok, cuma packaging kurang rapi",
            "Barang sesuai, tapi kurang puas sama respon seller",
            "Kualitas ok, tapi harga sedikit mahal",
            "Produk baik, tapi sayang stok warna tidak lengkap",
        }
        return comments[rand.Intn(len(comments))]
    } else if rating >= 3 {
        comments := []string{
            "Biasa saja, standar, tidak ada yang istimewa",
            "Cukup sesuai harga, lumayan",
            "Pengiriman lambat, barang standar",
            "Kualitas pas-pasan, tapi masih ok",
        }
        return comments[rand.Intn(len(comments))]
    } else if rating >= 2 {
        comments := []string{
            "Kurang puas, tidak sesuai foto",
            "Barang rusak saat sampai, pengiriman kurang hati-hati",
            "Kualitas kurang baik, cepat rusak",
            "Seller kurang responsif, pengiriman lama",
        }
        return comments[rand.Intn(len(comments))]
    } else {
        comments := []string{
            "Sangat kecewa, produk berbeda dengan deskripsi",
            "Barang tidak sampai, refund susah",
            "Kualitas jelek, minta refund ditolak",
            "Penipuan, barang tidak sesuai pesanan",
        }
        return comments[rand.Intn(len(comments))]
    }
}

func shuffleOrderItems(items []entity.OrderItem) []entity.OrderItem {
    shuffled := make([]entity.OrderItem, len(items))
    copy(shuffled, items)
    rand.Shuffle(len(shuffled), func(i, j int) {
        shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
    })
    return shuffled
}

func updateProductRatings(db *gorm.DB, productID uint) error {
    // Calculate new average rating
    var result struct {
        AvgRating float64
        Count     int
    }
    
    db.Model(&entity.Review{}).
        Select("COALESCE(AVG(rating), 0) as avg_rating, COUNT(*) as count").
        Where("product_id = ?", productID).
        Scan(&result)
    
    // Update product
    return db.Model(&entity.Product{}).
        Where("id = ?", productID).
        Updates(map[string]interface{}{
            "average_rating": result.AvgRating,
            "review_count":   result.Count,
        }).Error
}

// Helper for JSON marshaling
func marshalJSON(v interface{}) []byte {
    data, _ := json.Marshal(v)
    return data
}

func unmarshalJSON(data []byte, v interface{}) error {
    return json.Unmarshal(data, v)
}