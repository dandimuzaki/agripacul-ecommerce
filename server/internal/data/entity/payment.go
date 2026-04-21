package entity

import (
	"time"

	"gorm.io/gorm"
)

type PaymentType struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name string `json:"name"`

	Methods []PaymentMethod `gorm:"foreignKey:PaymentTypeID" json:"methods"`
}

type PaymentMethod struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	PaymentTypeID uint `json:"payment_type_id"`
	Name     string `gorm:"uniqueIndex;not null" json:"name"`
	IsActive bool   `gorm:"default:true" json:"is_active"`
	IconURL string `gorm:"type:text" json:"icon_url,omitempty"`
}