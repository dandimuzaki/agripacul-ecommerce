package entity

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	UserID            uint      `gorm:"uniqueIndex;not null" json:"user_id"`
	FullName          string    `gorm:"not null" json:"full_name"`
	Phone             string    `json:"phone"`
	DateOfBirth       *time.Time `json:"date_of_birth"`
	Salary            float64   `gorm:"type:numeric(12,2)" json:"salary"`
	ProfileImageURL   string    `json:"profile_image_url"`

	// Relations
	User   User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
}
