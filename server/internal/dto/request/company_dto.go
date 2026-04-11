package request

type CreateCompanyRequest struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description"`
	InstagramURL string `json:"instagram_url"`
	TwitterURL   string `json:"twitter_url"`
	WhatsappURL  string `json:"whatsapp_url"`
	ContactEmail string `json:"contact_email" binding:"omitempty,email"`
}

type UpdateCompanyRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	InstagramURL string `json:"instagram_url"`
	TwitterURL   string `json:"twitter_url"`
	WhatsappURL  string `json:"whatsapp_url"`
	ContactEmail string `json:"contact_email" binding:"omitempty,email"`
}

type CreateCompanyAddressRequest struct {
	CompanyID        uint   `json:"company_id" binding:"required"`
	ProvinceID       uint   `json:"province_id" binding:"required"`
	RegencyID        uint   `json:"regency_id" binding:"required"`
	DistrictID       uint   `json:"district_id" binding:"required"`
	SubdistrictID    uint   `json:"subdistrict_id" binding:"required"`
	PostalCode       string `json:"postal_code" binding:"required"`
	DetailAddress    string `json:"detail_address" binding:"required"`
	IsShippingOrigin bool   `json:"is_shipping_origin"`
}

type UpdateCompanyAddressRequest struct {
	ProvinceID       uint   `json:"province_id"`
	RegencyID        uint   `json:"regency_id"`
	DistrictID       uint   `json:"district_id"`
	SubdistrictID    uint   `json:"subdistrict_id"`
	PostalCode       string `json:"postal_code"`
	DetailAddress    string `json:"detail_address"`
	IsShippingOrigin *bool  `json:"is_shipping_origin"`
}

type SendMessageRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Subject     string `json:"subject"`
	Body        string `json:"body"`
}