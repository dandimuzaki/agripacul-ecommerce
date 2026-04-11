package response

import (
	"debian-ecommerce/internal/data/entity"
	"time"
)

type PreviewCheckout struct {
	Items []PreviewItem `json:"selected_items"`
	Totals Totals `json:"totals"`
}

type PreviewItem struct {
	ID          uint                        `json:"id"`
	SKUID       uint                        `json:"sku_id"`
	SKUCode     string                      `json:"sku_code"`
	Product     ProductSnapshot             `json:"product"`
	Variants    []entity.VariantCombination `json:"variants"`
	Price       PriceSnapshot               `json:"price"`
	Quantity    int                         `json:"quantity"`
	Stock       int                         `json:"stock"`
	SubTotal    float64                     `json:"subtotal"`
}

type PromotionSummary struct {
	ID      uint      `json:"id"`
	Name    string    `json:"name"`
	EndDate time.Time `json:"end_date"`
	DiscountType      entity.DiscountType    `json:"discount_type"`
	DiscountValue     float64                `json:"discount_value"`
	MinimumOrderValue float64                `json:"minimum_order_value"`
	MaximumDiscount   float64                `json:"maximum_discount"`
	UsageLeft        int                    `json:"usage_left"`
	VoucherCode       *string                `json:"voucher_code"`
}

func ToPromoSummary(p entity.Promotion) PromotionSummary {
	return PromotionSummary{
		ID: p.ID,
		Name: p.Name,
		EndDate: p.EndDate,
		DiscountType: p.DiscountType,
		DiscountValue: p.DiscountValue,
		MinimumOrderValue: p.MinimumOrderValue,
		MaximumDiscount: p.MaximumDiscount,
		UsageLeft: p.UsageLimit - p.UsedCount,
		VoucherCode: p.VoucherCode,
	}
}