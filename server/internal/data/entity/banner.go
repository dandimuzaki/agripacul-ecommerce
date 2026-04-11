package entity

import (
	"time"

	"gorm.io/gorm"
)

type BannerType string

const (
	BannerMain      BannerType = "main"
	BannerSecondary BannerType = "secondary"
)

type Banner struct {
	gorm.Model
	Name        string     `gorm:"not null" json:"name"`
	ImageURL    string     `gorm:"type:text" json:"image_url,omitempty"`
	ImagePublicID string `gorm:"column:image_public_id;type:text" json:"image_public_id"`
	TargetURL   string     `gorm:"type:text" json:"target_url,omitempty"`
	StartDate   time.Time  `gorm:"not null" json:"start_date"`
	EndDate     time.Time  `gorm:"not null" json:"end_date"`
	Type        BannerType `gorm:"type:varchar(50)" json:"type"`
	IsPublished bool       `gorm:"default:false" json:"is_published"`
}