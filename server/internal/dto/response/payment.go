package response

import "debian-ecommerce/internal/data/entity"

type PaymentMethodResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	IconURL  string `json:"icon_url"`
}

func ToPaymentMethodResponse(pm []entity.PaymentMethod) []PaymentMethodResponse {
	var paymentMethods []PaymentMethodResponse
	for _, p := range pm {
		payment := PaymentMethodResponse{
			ID: p.ID,
			Name: p.Name,
			IsActive: p.IsActive,
			IconURL: p.IconURL,
		}
		paymentMethods = append(paymentMethods, payment)
	}

	return paymentMethods
}

type PaymentMethodListResponse struct {
	PaymentMethods []PaymentMethodResponse `json:"payment_methods"`
	Total          int64                   `json:"total"`
	Page           int                     `json:"page"`
	Limit          int                     `json:"limit"`
	TotalPages     int64                   `json:"total_pages"`
}

type PaymentListResponse struct {
	PaymentTypes []entity.PaymentType `json:"payment_types"`
	Total          int64                   `json:"total"`
	Page           int                     `json:"page"`
	Limit          int                     `json:"limit"`
	TotalPages     int64                   `json:"total_pages"`
}
