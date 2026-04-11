package entity

import "gorm.io/gorm"

type CompanyAddress struct {
	gorm.Model
	CompanyID        uint   `gorm:"not null" json:"company_id"`
	Label            string `gorm:"type:varchar(100)" json:"label"`
	ProvinceID       uint   `gorm:"not null" json:"province_id"`
	RegencyID        uint   `gorm:"not null" json:"regency_id"`
	DistrictID       uint   `gorm:"not null" json:"district_id"`
	SubdistrictID    uint   `gorm:"not null" json:"subdistrict_id"`
	PostalCode       string `gorm:"type:varchar(10)" json:"postal_code"`
	DetailAddress    string `gorm:"type:text" json:"detail_address"`
	IsShippingOrigin bool   `gorm:"type:boolean;default:false" json:"is_shipping_origin"`

	// Relations
	Company     Company     `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
	Province    Province    `gorm:"foreignKey:ProvinceID" json:"province,omitempty"`
	Regency     Regency     `gorm:"foreignKey:RegencyID" json:"regency,omitempty"`
	District    District    `gorm:"foreignKey:DistrictID" json:"district,omitempty"`
	Subdistrict Subdistrict `gorm:"foreignKey:SubdistrictID" json:"subdistrict,omitempty"`
}
