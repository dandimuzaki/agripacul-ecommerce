package entity

import "gorm.io/gorm"

type Wishlist struct {
	gorm.Model
	CustomerID uint `gorm:"not null;index" json:"customer_id"`
	ProductID  uint `gorm:"not null;index" json:"product_id"`

	Customer *Customer `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"customer,omitempty"`
	Product  *Product  `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
}
