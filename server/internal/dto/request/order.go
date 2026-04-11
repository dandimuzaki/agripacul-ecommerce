package request

import "time"

type CreateOrderRequest struct {
	PromotionID       *uint            `json:"selected_promotion_id"`
	VoucherCode       *string          `json:"selected_voucher_code"`
	PaymentMethodID   uint             `json:"selected_payment_method_id"`
	ShippingAddressID uint             `json:"shipping_address_id"`
	Notes             string           `json:"notes"`
	Shipping          ShippingSnapshot `json:"selected_shipping_option"`
}

type ShippingSnapshot struct {
	Name    string  `json:"name"`
	Code    string  `json:"code"`
	Service string  `json:"service"`
	Cost    float64 `json:"cost"`
	ETD     string  `json:"etd"`
}

type OrderSortOption string

const (
	OrderSortByCreatedAt OrderSortOption = "created_at"
	OrderSortByUpdatedAt OrderSortOption = "updated_at"
	OrderSortByTotal OrderSortOption = "total"
)

type OrderListRequest struct {
	PaginationRequest

	// Filtering
	Status         string `form:"status"`
	Search         string `form:"search"`
	CustomerID     *uint `form:"customer"`
	StartDate      string `form:"start"`
	EndDate        string `form:"end"`
	Period         string `form:"period"`
	ShippingMethod string `form:"shipping"`

	// Sorting
	SortBy  OrderSortOption `form:"sort_by" binding:"omitempty,oneof=created_at updated_at total"`
	SortOrder SortOrder         `form:"sort_order" binding:"omitempty,oneof=asc desc"`
}

type PeriodOption string
const (
	Today PeriodOption = "today"
	Week PeriodOption = "last_7_days"
	Month PeriodOption = "this_month"
)

type OrderQueryParams struct {
	Status         string `form:"status"`
	Search         string `form:"search"`
	CustomerID     *uint `form:"customer"`
	StartDate      time.Time `form:"start"`
	EndDate        *time.Time `form:"end"`
	Period         PeriodOption `form:"period"`
	ShippingMethod string `form:"shipping"`

	SortOrder SortOrder
	SortBy  OrderSortOption
	Page int
	Limit   int
	Offset  int
}

func (r OrderListRequest) ToQuery() (*OrderQueryParams, error) {
	var start time.Time
	var end *time.Time
	if r.StartDate != "" {
		from, err := time.Parse("02-01-2006", r.StartDate)
		if err != nil {
			return nil, err
		}
		start = from
	} else {
		from, err := time.Parse("02-01-2006", "01-01-2000")
		if err != nil {
			return nil, err
		}
		start = from
	}
	
	if r.EndDate != "" {
		to, err := time.Parse("02-01-2006", r.EndDate)
		if err != nil {
			return nil, err
		}
		end = &to
	} else {
		now := time.Now()
		end = &now
	}

	return &OrderQueryParams{
		Status: r.Status,
		Search: r.Search,
		CustomerID: r.CustomerID,
		StartDate: start,
		EndDate: end,
		Period: PeriodOption(r.Period),
		ShippingMethod: r.ShippingMethod,
		SortOrder: r.SortOrder,
		SortBy: r.SortBy,
		Limit: r.GetPerPage(),
		Offset: r.GetOffset(),
		Page: r.Page,
	}, nil
}

type CancelOrderRequest struct {
	CancelReason string `json:"cancel_reason"`
}