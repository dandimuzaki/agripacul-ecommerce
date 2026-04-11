package entity

import (
	"time"

	"gorm.io/gorm"
)

type Campaign struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	Type string `gorm:"type:varchar(100)"` // collection/discount
	StartDate time.Time
	EndDate time.Time
	IsActive bool `gorm:"default:true"`

	CampaignProducts []CampaignProduct `gorm:"foreignKey:CampaignID"`
}

type CampaignProduct struct {
	gorm.Model
	CampaignID uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`

	Campaign Campaign `gorm:"foreignKey:CampaignID" json:"campaign,omitempty"`
	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}