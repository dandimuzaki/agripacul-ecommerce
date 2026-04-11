package response

import (
	"time"

	"debian-ecommerce/internal/data/entity"
)

type PromotionResponse struct {
	ID                uint                   `json:"id"`
	Name              string                 `json:"name"`
	StartDate         time.Time              `json:"start_date"`
	EndDate           time.Time              `json:"end_date"`
	Type              entity.PromoType       `json:"type"`
	Description       string                 `json:"description"`
	IsPublished       bool                   `json:"is_published"`
	DiscountType      entity.DiscountType    `json:"discount_type"`
	DiscountValue     float64                `json:"discount_value"`
	MinimumOrderValue float64                `json:"minimum_order_value"`
	MaximumDiscount   float64                `json:"maximum_discount"`
	UsageLimit        int                    `json:"usage_limit"`
	VoucherCode       *string                `json:"voucher_code"`
	IsPublic           bool                   `json:"is_shown"`
	CreatedAt         time.Time              `json:"created_at"`
	UpdatedAt         time.Time              `json:"updated_at"`
	PromoProducts     []PromoProductResponse `json:"promo_products,omitempty"`
}

type PromoProductResponse struct {
	ID          uint    `json:"id"`
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	ProductSKU  string  `json:"product_sku"`
	Price       float64 `json:"price"`
}

type PromotionListResponse struct {
	Promotions []PromotionResponse `json:"promotions"`
	Pagination PaginationResponse  `json:"pagination"`
}
