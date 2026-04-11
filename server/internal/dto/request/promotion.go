package request

import (
	"time"

	"debian-ecommerce/internal/data/entity"
)

type PromotionRequest struct {
	Name              string              `json:"name" validate:"required"`
	StartDate         time.Time           `json:"start_date" validate:"required"`
	EndDate           time.Time           `json:"end_date" validate:"required,gtfield=StartDate"`
	Type              entity.PromoType    `json:"type" validate:"required,oneof='direct discount' 'voucher code'"`
	Description       string              `json:"description"`
	IsPublished       bool                `json:"is_published"`
	DiscountType      entity.DiscountType `json:"discount_type" validate:"required,oneof=amount percentage"`
	DiscountValue     float64             `json:"discount_value" validate:"required,min=0"`
	MinimumOrderValue float64             `json:"minimum_order_value" validate:"min=0"`
	MaximumDiscount   float64             `json:"maximum_discount" validate:"min=0"`
	UsageLimit        int                 `json:"usage_limit" validate:"min=0"`
	VoucherCode       *string             `json:"voucher_code"`
	IsPublic           bool                `json:"is_shown"`
	ProductIDs        []uint              `json:"product_ids"`
}

type PromotionUpdateRequest struct {
	Name              string              `json:"name"`
	StartDate         *time.Time          `json:"start_date"`
	EndDate           *time.Time          `json:"end_date" validate:"omitempty,gtfield=StartDate"`
	Type              entity.PromoType    `json:"type" validate:"omitempty,oneof='direct discount' 'voucher code'"`
	Description       *string             `json:"description"`
	IsPublished       *bool               `json:"is_published"`
	DiscountType      entity.DiscountType `json:"discount_type" validate:"omitempty,oneof=amount percentage"`
	DiscountValue     *float64            `json:"discount_value" validate:"omitempty,min=0"`
	MinimumOrderValue *float64            `json:"minimum_order_value" validate:"omitempty,min=0"`
	MaximumDiscount   *float64            `json:"maximum_discount" validate:"omitempty,min=0"`
	UsageLimit        *int                `json:"usage_limit" validate:"omitempty,min=0"`
	VoucherCode       *string             `json:"voucher_code"`
	IsPublic           *bool               `json:"is_shown"`
	ProductIDs        []uint              `json:"product_ids"`
}

type PromotionFilterRequest struct {
	PaginationRequest

	Name          string           `form:"name"`
	Type          entity.PromoType `form:"type"`
	IsPublished   bool            `form:"is_published"`
	IsActive   bool            `form:"is_active"`
	StartDateFrom string       `form:"start_date_from"`
	StartDateTo   string        `form:"start_date_to"`
	EndDateFrom   string        `form:"end_date_from"`
	EndDateTo     string       `form:"end_date_to"`
	VoucherCode   string           `form:"voucher_code"`
	SortBy SortOption `form:"sort_by"`
	SortOrder SortOrder `form:"sort_order"`
	MinimumOrderValue float64 `form:"minimum"`
	Available bool `form:"available"`
}

type PromotionFilterQuery struct {
	Name          string           `form:"name"`
	Type          entity.PromoType `form:"type"`
	IsPublished   bool            `form:"is_published"`
	IsActive   bool            `form:"is_active"`
	StartDateFrom time.Time        `form:"start_date_from"`
	StartDateTo   time.Time        `form:"start_date_to"`
	EndDateFrom   time.Time        `form:"end_date_from"`
	EndDateTo     time.Time        `form:"end_date_to"`
	VoucherCode   string           `form:"voucher_code"`
	Page int
	Limit int
	Offset int
	SortBy SortOption
	SortOrder SortOrder
	MinimumOrderValue float64
	Available bool
}

func (r PromotionFilterRequest) ToQuery() (*PromotionFilterQuery, error) {
	var startFrom time.Time
	var startTo time.Time
	var endFrom time.Time
	var endTo time.Time
	if r.StartDateFrom != "" {
		start, err := time.Parse("02-01-2006", r.StartDateFrom)
		if err != nil {
			return nil, err
		}
		startFrom = start
	}

	if r.StartDateTo != "" {
		start, err := time.Parse("02-01-2006", r.StartDateTo)
		if err != nil {
			return nil, err
		}
		startTo = start
	}
	
	if r.EndDateFrom != "" {
		end, err := time.Parse("02-01-2006", r.EndDateFrom)
		if err != nil {
			return nil, err
		}
		endFrom = end
	}

	if r.EndDateTo != "" {
		end, err := time.Parse("02-01-2006", r.EndDateFrom)
		if err != nil {
			return nil, err
		}
		endTo = end
	}

	return &PromotionFilterQuery{
		Name: r.Name,
		Type: r.Type,
		IsPublished: r.IsPublished,
		IsActive: r.IsActive,
		StartDateFrom: startFrom,
		StartDateTo: startTo,
		EndDateFrom: endFrom,
		EndDateTo: endTo,
		VoucherCode: r.VoucherCode,
		Page: r.GetPage(),
		Limit: r.GetPerPage(),
		Offset: r.GetOffset(),
	}, nil
}
