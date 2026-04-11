package entity

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	CustomerID uint `gorm:"not null" json:"customer_id"`
	ProductID uint `gorm:"not null" json:"product_id"`
	SKUID uint `gorm:"column:sku_id;not null" json:"sku_id"`
	OrderID uint `gorm:"not null" json:"order_id"`
	Rating int `gorm:"check:rating >=1 AND rating <= 5" json:"rating"`
	Comment string `gorm:"type:text" json:"comment"`

	// Relations
	Customer Customer `gorm:"foreignKey:CustomerID" json:"customer"`
	Product Product `gorm:"foreignKey:ProductID" json:"product"`
	SKU SKU `gorm:"foreignKey:SKUID" json:"sku"`
	Order Order `gorm:"foreignKey:OrderID" json:"order"`
}