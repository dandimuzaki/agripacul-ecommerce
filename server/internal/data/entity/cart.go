package entity

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	CustomerID uint `gorm:"uniqueIndex;not null" json:"customer_id"`

	// Relations
	Customer Customer   `gorm:"foreignKey:CustomerID;constraint:OnDelete:CASCADE" json:"customer,omitempty"`
	Items    []CartItem `gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE" json:"items,omitempty"`
}

type CartItem struct {
	gorm.Model
	CartID   uint `gorm:"index;not null" json:"cart_id"`
	SKUID    uint `gorm:"column:sku_id;index;not null" json:"sku_id"`
	Quantity int  `gorm:"not null" json:"quantity"`
	IsSelected bool `gorm:"default:false" json:"is_selected"`

	// Relations
	Cart Cart `gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE" json:"cart,omitempty"`
	SKU  SKU  `gorm:"foreignKey:SKUID;constraint:OnDelete:RESTRICT" json:"sku,omitempty"`
}
