package response

import "debian-ecommerce/internal/data/entity"

type CartItemResponse struct {
	ID          uint                        `json:"id"`
	SKUID       uint                        `json:"sku_id"`
	SKUCode     string                      `json:"sku_code"`
	Product     ProductSnapshot             `json:"product"`
	Variants    []entity.VariantCombination `json:"variants"`
	Price       PriceSnapshot               `json:"price"`
	Quantity    int                         `json:"quantity"`
	Stock       int                         `json:"stock"`
	SubTotal    float64                     `json:"subtotal"`
	IsSelected  bool                        `json:"is_selected"`
	IsAvailable bool                        `json:"is_available"`
}

type CartResponse struct {
	ID         uint               `json:"id"`
	CustomerID uint               `json:"customer_id"`
	Items      []CartItemResponse `json:"items"`
	Summary TotalSnapshot `json:"summary"`
}

type ProductSnapshot struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	MainImageURL string `json:"main_image_url"`
}

type PriceSnapshot struct {
	UnitPrice          float64 `json:"unit_price"`
	SalePrice          *float64 `json:"sale_price,omitempty"`
	DiscountPercentage float64 `json:"discount_percentage,omitempty"`
}

type TotalSnapshot struct {
	TotalItems int                `json:"total_items"`
	TotalPrice float64            `json:"total_price"`
	TotalSelectedItems int                `json:"total_selected_items"`
	TotalSelectedPrice float64            `json:"total_selected_price"`
}