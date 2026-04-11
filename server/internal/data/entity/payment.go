package entity

import "gorm.io/gorm"

type PaymentMethod struct {
	gorm.Model
	Name     string `gorm:"uniqueIndex;not null" json:"name"`
	IsActive bool   `gorm:"default:true" json:"is_active"`
	IconURL string `gorm:"type:text" json:"icon_url,omitempty"`
}