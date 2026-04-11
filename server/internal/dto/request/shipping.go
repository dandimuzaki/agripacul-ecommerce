package request

type ShippingRequest struct {
	ShippingAddressID uint    `json:"shipping_address_id"`
	TotalWeight       float64 `json:"total_weight"`
}