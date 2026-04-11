package request

type PreviewCheckout struct {
	SelectedShippingOption *ShippingSnapshot `json:"selected_shipping_option"`
	SelectedPromotionID    *uint             `json:"selected_promotion_id,omitempty"`
	SelectedVoucher        *string           `json:"selected_voucher,omitempty"`
}

type ShippingOptionsRequest struct {
	ShippingAddressID *uint `json:"shipping_address_id"`
}