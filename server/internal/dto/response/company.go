package response

type CompanyResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	InstagramURL string `json:"instagram_url"`
	TwitterURL   string `json:"twitter_url"`
	WhatsappURL  string `json:"whatsapp_url"`
	ContactEmail string `json:"contact_email"`
}

type CompanyAddressResponse struct {
	ID               uint   `json:"id"`
	CompanyID        uint   `json:"company_id"`
	Province         string `json:"province"`
	Regency          string `json:"regency"`
	District         string `json:"district"`
	Subdistrict      string `json:"subdistrict"`
	PostalCode       string `json:"postal_code"`
	DetailAddress    string `json:"detail_address"`
	IsShippingOrigin bool   `json:"is_shipping_origin"`
}