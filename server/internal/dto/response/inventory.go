package response

import (
	"debian-ecommerce/internal/data/entity"
	"time"
)

type InventoryResponse struct {
	ID           uint                 `json:"id"`
	Product      string               `json:"product"`
	SKUCode      string               `json:"sku_code"`
	VariantLabel string `json:"variant_label"`
	Stock        int                  `json:"stock"`
	MinStock     int                  `json:"min_stock"`
	Availability string                 `json:"availability"`
	Status       string               `json:"status"`
}

type InventoryWithVariants struct {
	InventoryResponse
	Variants []entity.VariantCombination `json:"variants"`
}

type InventoryLogResponse struct {
	ID                uint                    `json:"id"`
	SKUID         uint                    `json:"sku_id"`
	Type              entity.InventoryLogType `json:"type"`
	QuantityChange    int                     `json:"quantity_change"`
	CurrentStockAfter int                     `json:"current_stock_after"`
	ReferenceID       *uint                   `json:"reference_id,omitempty"`
	ReferenceType     string                  `json:"reference_type,omitempty"`
	Notes             string                  `json:"notes,omitempty"`
	CreatedAt         time.Time               `json:"created_at"`
}