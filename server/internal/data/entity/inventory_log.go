package entity

import (
	"gorm.io/gorm"
)

// InventoryLogType enum
type InventoryLogType string

const (
	InventoryLogTypeIn         InventoryLogType = "in"
	InventoryLogTypeOut        InventoryLogType = "out"
	InventoryLogTypeAdjustment InventoryLogType = "adjustment"
	InventoryLogTypeInitial    InventoryLogType = "initial"
)

type InventoryLog struct {
	gorm.Model
	SKUID         uint             `gorm:"column:sku_id;index;not null" json:"sku_id"`
	Type              InventoryLogType `gorm:"type:varchar(20);not null" json:"type"`
	QuantityChange    int              `gorm:"not null" json:"quantity_change"`
	CurrentStockAfter int              `gorm:"not null" json:"current_stock_after"`
	ReferenceID       *uint            `gorm:"index" json:"reference_id,omitempty"`
	ReferenceType     string           `gorm:"type:varchar(50)" json:"reference_type,omitempty"`
	Notes             string           `json:"notes,omitempty"`

	// Relations
	SKU SKU `gorm:"foreignKey:SKUID" json:"sku"`
}
