package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	RoleSuperAdmin UserRole = "superadmin"
	RoleAdmin      UserRole = "admin"
	RoleStaff UserRole = "staff"
	RoleCustomer      UserRole = "customer"
)

type User struct {
	gorm.Model
	Email        string         `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string         `gorm:"not null" json:"-"`
	Role         UserRole       `gorm:"type:varchar(50);not null" json:"role"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	IsActive bool `gorm:"default:true;not null" json:"is_active"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`

	// Relations
	Customer *Customer `gorm:"foreignKey:UserID" json:"customer,omitempty"`
	Employee *Employee `gorm:"foreignKey:UserID" json:"employee,omitempty"`
}