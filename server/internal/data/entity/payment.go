package entity

import "gorm.io/gorm"

type PaymentType struct {
	gorm.Model
	Name string `json:"name"`

	Methods []PaymentMethod `gorm:"foreignKey:PaymentTypeID" json:"methods"`
}

type PaymentMethod struct {
	gorm.Model
	PaymentTypeID uint `json:"payment_type_id"`
	Name     string `gorm:"uniqueIndex;not null" json:"name"`
	IsActive bool   `gorm:"default:true" json:"is_active"`
	IconURL string `gorm:"type:text" json:"icon_url,omitempty"`
}