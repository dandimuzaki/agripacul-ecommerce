package request

type CreateAddressRequest struct {
	RecipientName string `json:"recipient_name"`
	Label         string `json:"label"`
	PhoneNumber   string `json:"phone_number"`
	ProvinceID    uint   `json:"province_id" validate:"required"`
	RegencyID     uint   `json:"regency_id" validate:"required"`
	DistrictID    uint   `json:"district_id" validate:"required"`
	SubdistrictID uint   `json:"subdistrict_id" validate:"required"`
	PostalCode    string `json:"postal_code" validate:"max=10"`
	DetailAddress string `json:"detail_address"`
	IsDefault     bool   `json:"is_default"`
}

type UpdateAddressRequest struct {
	RecipientName string `json:"recipient_name"`
	Label         string `json:"label"`
	PhoneNumber   string `json:"phone_number"`
	ProvinceID    uint   `json:"province_id"`
	RegencyID     uint   `json:"regency_id"`
	DistrictID    uint   `json:"district_id"`
	SubdistrictID uint   `json:"subdistrict_id"`
	PostalCode    string `json:"postal_code" validate:"max=10"`
	DetailAddress string `json:"detail_address"`
	IsDefault     bool   `json:"is_default"`
}
