package entity

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	CustomerID uint `gorm:"not null" json:"customer_id"`
	RecipientName string `gorm:"type:varchar(255);not null" json:"recipient_name"`
	Label string `gorm:"type:varchar(100)" json:"label"`
	PhoneNumber string `gorm:"varchar(50)"`
	ProvinceID uint `gorm:"not null" json:"province_id"`
	RegencyID uint `gorm:"not null" json:"regency_id"`
	DistrictID uint `gorm:"not null" json:"district_id"`
	SubdistrictID uint `gorm:"not null" json:"subdistrict_id"`
	PostalCode string `gorm:"type:varchar(10)" json:"postal_code"`
	DetailAddress string `gorm:"type:text;not null" json:"detail_address"`
	IsDefault bool `gorm:"type:boolean;default:false" json:"is_default"`

	// Relations
	Customer    Customer    `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	Province    Province    `gorm:"foreignKey:ProvinceID" json:"province,omitempty"`
	Regency     Regency     `gorm:"foreignKey:RegencyID" json:"regency,omitempty"`
	District    District    `gorm:"foreignKey:DistrictID" json:"district,omitempty"`
	Subdistrict Subdistrict `gorm:"foreignKey:SubdistrictID" json:"subdistrict,omitempty"`
}

type Province struct {
	gorm.Model
	Code         string `gorm:"type:varchar(50);uniqueIndex;not null"`
	Name         string `gorm:"type:varchar(100);not null"`
	RajaOngkirID *uint  `gorm:"column:raja_ongkir_id;index"`
}

type Regency struct {
	gorm.Model
	Code         string `gorm:"type:varchar(50);uniqueIndex;not null"`
	ProvinceID   uint   `gorm:"not null;index"`
	Name         string `gorm:"type:varchar(100);not null"`
	Type         string `gorm:"type:varchar(20);not null"` // kota / kabupaten
	RajaOngkirID *uint  `gorm:"column:raja_ongkir_id;index"`

	Province Province `gorm:"foreignKey:ProvinceID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type District struct {
	gorm.Model
	Code         string `gorm:"type:varchar(50);uniqueIndex;not null"`
	RegencyID    uint   `gorm:"not null;index"`
	Name         string `gorm:"type:varchar(100);not null"`
	RajaOngkirID *uint  `gorm:"column:raja_ongkir_id;index"`

	Regency Regency `gorm:"foreignKey:RegencyID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type Subdistrict struct {
	gorm.Model
	Code       string `gorm:"type:varchar(50);uniqueIndex;not null"`
	DistrictID uint   `gorm:"not null;index"`
	Name       string `gorm:"type:varchar(100);not null"`

	District District `gorm:"foreignKey:DistrictID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}
