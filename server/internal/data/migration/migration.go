package migration

import (
	"debian-ecommerce/internal/data/entity"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
		&entity.Customer{},
		&entity.Employee{},

		&entity.Category{},
		&entity.Product{},
		&entity.VariantType{},
		&entity.VariantValue{},
		&entity.SKU{},
		&entity.SKUVariantValue{},
		&entity.Image{},

		&entity.Cart{},
		&entity.CartItem{},
		&entity.Wishlist{},
		&entity.Campaign{},
		&entity.CampaignProduct{},

		&entity.PaymentType{},
		&entity.PaymentMethod{},
		&entity.Promotion{},
		&entity.PromoProduct{},
		&entity.Order{},
		&entity.OrderItem{},
		&entity.PromoUsage{},
		&entity.Review{},
		&entity.OrderPayment{},
		&entity.OrderShipping{},
		&entity.InventoryLog{},

		&entity.Banner{},
		&entity.Province{},
		&entity.Regency{},
		&entity.District{},
		&entity.Subdistrict{},
		&entity.Address{},
		&entity.Company{},
		&entity.CompanyAddress{},
	)
}
