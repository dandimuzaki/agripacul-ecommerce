package entity

import (
	"time"

	"gorm.io/gorm"
)

type PromoType string

const (
	DirectDiscount PromoType = "direct discount"
	VoucherCode    PromoType = "voucher code"
)

type DiscountType string

const (
	DiscountAmount     DiscountType = "amount"
	DiscountPercentage DiscountType = "percentage"
)

type Promotion struct {
	gorm.Model
	Name string `gorm:"type:text;not null" json:"name"`
	StartDate time.Time `gorm:"not null" json:"start_date"`
	EndDate time.Time `gorm:"not null" json:"end_date"`
	Type PromoType `gorm:"type:varchar(50);not null" json:"type"`
	Description string `gorm:"type:text" json:"description"`
	IsPublished bool `gorm:"default:true" json:"is_published"`
	DiscountType DiscountType `gorm:"type:varchar(50);not null" json:"discount_type"`
	DiscountValue float64 `gorm:"type:numeric(12,2);default:0" json:"discount_value,omitempty"`
	MinimumOrderValue float64 `gorm:"type:numeric(12,2);default:0" json:"minimum_order_value,omitempty"`
	MaximumDiscount float64 `gorm:"type:numeric(12,2)" json:"maximum_discount,omitempty"`
	UsageLimit int `gorm:"default:0" json:"usage_limit"`
	VoucherCode *string `gorm:"type:varchar(50)" json:"voucher_code"`
	IsPublic bool `gorm:"default:true" json:"is_public"`
	UsedCount int `gorm:"default:0" json:"used_count"`

	// Relations
	PromoProducts []PromoProduct `gorm:"foreignKey:PromotionID" json:"promo_products,omitempty"`
	PromoUsages []PromoUsage `gorm:"foreignKey:PromotionID" json:"promo_usages,omitempty"`
}

type PromoProduct struct {
	gorm.Model
	PromotionID uint `gorm:"not null;index" json:"promotion_id"`
	ProductID uint `gorm:"not null;index" json:"product_id"`

	// Relations
	Promotion Promotion `gorm:"foreignKey:PromotionID;constraint:OnDelete:CASCADE" json:"promotion,omitempty"`
	Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"product,omitempty"`
}

type PromoUsage struct {
	gorm.Model
	PromotionID uint `gorm:"not null;index" json:"promotion_id"`
	CustomerID uint `gorm:"not null;index" json:"customer_id"`
	OrderID uint `gorm:"not null;uniqueIndex" json:"order_id"`
	UsedAt time.Time `gorm:"not null;default:now()" json:"used_at,omitempty"`

	// Relations
	Promotion Promotion `gorm:"foreignKey:PromotionID;constraint:OnDelete:CASCADE" json:"promotion,omitempty"`
	Customer Customer `gorm:"foreignKey:CustomerID;constraint:OnDelete:CASCADE" json:"customer,omitempty"`
	Order Order `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"order,omitempty"`
}

type PromotionSnapshot struct {
	ID uint `json:"id"`
	Name string `gorm:"type:text;not null" json:"name"`
	Type PromoType `gorm:"type:varchar(50);not null" json:"type"`
	DiscountType DiscountType `gorm:"type:varchar(50);not null" json:"discount_type"`
	DiscountValue float64 `gorm:"type:numeric(12,2);default:0" json:"discount_value,omitempty"`
	MinimumOrderValue float64 `gorm:"type:numeric(12,2);default:0" json:"minimum_order_value,omitempty"`
	MaximumDiscount float64 `gorm:"type:numeric(12,2)" json:"maximum_discount,omitempty"`
	VoucherCode *string `gorm:"type:varchar(50)" json:"voucher_code,omitempty"`
}

func NewPromotionSnapshot(p Promotion) PromotionSnapshot {
	return PromotionSnapshot{
		ID: p.ID,
		Name: p.Name,
		Type: p.Type,
		DiscountType: p.DiscountType,
		DiscountValue: p.DiscountValue,
		MinimumOrderValue: p.MinimumOrderValue,
		MaximumDiscount: p.MaximumDiscount,
		VoucherCode: p.VoucherCode,
	}
}
