package entity

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	UserID            uint      `gorm:"uniqueIndex;not null" json:"user_id"`
	FullName          string    `gorm:"not null" json:"full_name"`
	PhoneNumber             string    `json:"phone_number"`
	DateOfBirth       *time.Time `json:"date_of_birth"`
	ProfileImageURL   string    `gorm:"type:text" json:"profile_image_url"`
	ProfileImagePublicID string `gorm:"type:text" json:"profile_image_public_id"`

	// Relations
	User   User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	Addresses []Address `gorm:"foreignKey:CustomerID" json:"addresses,omitempty"`
	Cart *Cart `gorm:"foreignKey:CustomerID" json:"cart,omitempty"`
	Orders []Order `gorm:"foreignKey:CustomerID" json:"orders,omitempty"`
}
