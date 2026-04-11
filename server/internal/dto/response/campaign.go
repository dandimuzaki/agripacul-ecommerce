package response

import "time"

type CampaignResponse struct {
	ID          uint             `json:"id"`
	Name string           `json:"name"`
	Description string           `json:"description"`
	Type        string           `json:"type"` // collection/discount
	StartDate   time.Time        `json:"start_date"`
	EndDate     time.Time        `json:"end_date"`
	IsActive    bool             `json:"is_active"`
	Products    []ProductSummary `json:"products"`
}