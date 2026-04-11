package entity

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100);not null" json:"name"`
	IconURL string `gorm:"type:text" json:"icon_url,omitempty"`
	IconPublicID string `gorm:"column:icon_public_id;type:text" json:"icon_public_id"`

	Products []Product `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
}

