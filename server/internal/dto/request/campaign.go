package request

import "time"

type CampaignFilterRequest struct {
	PaginationRequest
}

type CampaignListQuery struct {
	Page   int
	Limit  int
	Offset int
}

func (r CampaignFilterRequest) ToQuery() CampaignListQuery {
	return CampaignListQuery{
		Page:   r.GetPage(),
		Limit:  r.GetPerPage(),
		Offset: r.GetOffset(),
	}
}

type CampaignRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"` // collection/discount
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	IsActive    bool      `json:"is_active"`
	ProductIDs  []uint    `json:"product_ids"`
}