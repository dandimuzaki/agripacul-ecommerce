package response

import "debian-ecommerce/internal/data/entity"

type AddressResponse struct {
	ID            uint                `json:"id"`
	RecipientName string              `json:"recipient_name"`
	Label         string              `json:"label"`
	PhoneNumber string `json:"phone_number"`
	CustomerID    uint                `json:"customer_id"`
	Province      ProvinceResponse    `json:"province"`
	Regency       RegencyResponse     `json:"regency"`
	District      DistrictResponse    `json:"district"`
	Subdistrict   SubdistrictResponse `json:"subdistrict"`
	PostalCode    string              `json:"postal_code"`
	DetailAddress string              `json:"detail_address"`
	IsDefault     bool                `json:"is_default"`
}

type ProvinceResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type RegencyResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type DistrictResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type SubdistrictResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ToProvince(prov entity.Province) ProvinceResponse {
	return ProvinceResponse{
		ID: prov.ID,
		Name: prov.Name,
	}
}

func ToRegency(reg entity.Regency) RegencyResponse {
	return RegencyResponse{
		ID: reg.ID,
		Name: reg.Name,
		Type: reg.Type,
	}
}

func ToDistrict(dis entity.District) DistrictResponse {
	return DistrictResponse{
		ID: dis.ID,
		Name: dis.Name,
	}
}

func ToSubdistrict(subdis entity.Subdistrict) SubdistrictResponse {
	return SubdistrictResponse{
		ID: subdis.ID,
		Name: subdis.Name,
	}
}