package entity

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name         string `gorm:"type:varchar(255);not null" json:"name"`
	Description  string `gorm:"type:text" json:"description"`
	LogoURL      string `gorm:"type:text" json:"logo_url"`
	PhoneNumber  string `gorm:"type:varchar(50)" json:"phone_number"`
	InstagramURL string `gorm:"type:text" json:"instagram_url"`
	TwitterURL   string `gorm:"type:text" json:"twitter_url"`
	WhatsappURL  string `gorm:"type:text" json:"whatsapp_url"`
	ContactEmail string `gorm:"type:varchar(255)" json:"contact_email"`
}
