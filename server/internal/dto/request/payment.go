package request

type CreatePaymentMethodRequest struct {
	Name     string `json:"name" validate:"required"`
	IsActive *bool  `json:"is_active"`
	IconURL  string `json:"icon_url"`
}

type UpdatePaymentMethodRequest struct {
	Name     string `json:"name"`
	IsActive *bool  `json:"is_active"`
	IconURL  string `json:"icon_url"`
}

type GetPaymentMethodsRequest struct {
	Page     int    `json:"page" form:"page" default:"1"`
	Limit    int    `json:"limit" form:"limit" default:"10"`
	Search   string `json:"search" form:"search"`
	IsActive *bool  `json:"is_active" form:"is_active"`
}
