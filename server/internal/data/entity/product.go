package entity

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	CategoryID    uint           `gorm:"index;not null" json:"category_id"`
	Name          string         `gorm:"type:text;not null" json:"name"`
	Description   string         `gorm:"type:text" json:"description,omitempty"`
	IsPublished   bool           `gorm:"default:false" json:"is_published"`
	Tags          pq.StringArray `gorm:"type:text[];default:'{}'" json:"tags,omitempty"`
	Slug          string         `gorm:"type:text" json:"slug"`
	MainImageURL  string         `gorm:"type:text" json:"main_image_url"`
	MainImagePublicID  string         `gorm:"type:text" json:"main_image_public_id"`
	AverageRating float64        `gorm:"type:decimal;default:0" json:"avg_rating,omitempty"`
	ReviewCount   int            `gorm:"type:bigint;default:0" json:"review_count"`
	SoldCount     int            `gorm:"type:bigint;default:0" json:"sold_count"`
	MinPrice      float64        `gorm:"type:numeric(12,2);default:0" json:"min_price"`
	MaxPrice      float64        `gorm:"type:numeric(12,2);default:0" json:"max_price"`

	// Relations
	Category     Category      `gorm:"foreignKey:CategoryID;constraint:OnDelete:RESTRICT" json:"category,omitempty"`
	VariantTypes []VariantType `gorm:"foreignKey:ProductID" json:"variant_types,omitempty"`
	SKUs         []SKU         `gorm:"foreignKey:ProductID" json:"skus,omitempty"`
	Reviews      []Review      `gorm:"foreignKey:ProductID" json:"reviews,omitempty"`
	Images       []Image       `gorm:"foreignKey:ProductID" json:"images,omitempty"`
}

type VariantType struct {
	gorm.Model
	ProductID uint   `gorm:"index;not null" json:"product_id"`
	Name      string `gorm:"type:varchar(50);not null" json:"name"`

	// Relations
	Product Product        `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
	Values  []VariantValue `gorm:"foreignKey:VariantTypeID" json:"values"`
}

type VariantValue struct {
	gorm.Model
	VariantTypeID uint   `gorm:"index;not null" json:"variant_type_id"`
	Value         string `gorm:"not null" json:"value"`

	VariantType VariantType `gorm:"foreignKey:VariantTypeID" json:"variant_type,omitempty"`
}

type SKUStatus string

const (
	SKUStatusActive   SKUStatus = "active"
	SKUStatusInactive SKUStatus = "inactive"
	SKUStatusArchived SKUStatus = "archived"
)

type SKU struct {
	gorm.Model
	ProductID uint   `gorm:"index:idx_product_sku_code,unique"`
	SKUCode   string `gorm:"index:idx_product_sku_code,unique;type:varchar(100)"`
	Price     float64   `gorm:"type:numeric(12,2);not null" json:"price"`
	SalePrice     *float64   `gorm:"type:numeric(12,2)" json:"sale_price"`
	Stock     int       `gorm:"not null;default:0" json:"stock"`
	MinStock     int       `gorm:"not null;default:0" json:"min_stock"`
	Status    SKUStatus `gorm:"type:varchar(50)" json:"status"`
	Weight float64 `gorm:"column:weight_gram;type:decimal" json:"weight_gram"`

	Product          Product           `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"product,omitempty"`
	SKUVariantValues []SKUVariantValue `gorm:"foreignKey:SKUID"`
	Images           []Image           `gorm:"foreignKey:SKUID;constraint:OnDelete:CASCADE" json:"images,omitempty"`
}

type SKUVariantValue struct {
	gorm.Model
	SKUID          uint `gorm:"column:sku_id;not null;index:idx_sku_variant_value,unique" json:"sku_id"`
	VariantValueID uint `gorm:"not null;index:idx_sku_variant_value,unique" json:"variant_value_id"`

	SKU          SKU          `gorm:"foreignKey:SKUID;constraint:OnDelete:CASCADE" json:"sku"`
	VariantValue VariantValue `gorm:"foreignKey:VariantValueID;constraint:OnDelete:CASCADE" json:"variant_value"`
}

type Image struct {
	gorm.Model
	ProductID uint   `gorm:"not null" json:"product_id"`
	SKUID     *uint  `gorm:"column:sku_id" json:"sku_id"`
	ImageURL  string `gorm:"type:text" json:"image_url,omitempty"`
	PublicID string `gorm:"type:text" json:"public_id,omitempty"`
}

func (SKU) TableName() string {
	return "skus"
}

func (SKUVariantValue) TableName() string {
	return "sku_variant_values"
}

type VariantCombination struct {
	SKUID uint   `gorm:"column:sku_id" json:"sku_id,omitempty"`
	Name  string `gorm:"column:name" json:"name"`
	Value string `gorm:"column:value" json:"value"`
}

type ProductSnapshot struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Category string `json:"category"`
	Tags pq.StringArray `json:"tags"`
	MainImageURL string `json:"main_image_url"`
}

type SKUSnapshot struct {
	ID uint `json:"id"`
	SKUCode string `json:"sku_code"`
	Price float64 `json:"price"`
	SalePrice *float64 `json:"sale_price"`
	Variants []VariantCombination `json:"variants"`
}

func NewProductSnapshot(p Product) ProductSnapshot {
	return ProductSnapshot{
		ID: p.ID,
		Name: p.Name,
		Category: p.Category.Name,
		Tags: p.Tags,
		MainImageURL: p.MainImageURL,
	}
}

func NewSKUSnapshot(sku SKU) SKUSnapshot {
	return SKUSnapshot{
		ID: sku.ID,
		SKUCode: sku.SKUCode,
		Price: sku.Price,
		SalePrice: sku.SalePrice,
	}
}