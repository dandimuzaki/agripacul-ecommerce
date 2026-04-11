package response

import "debian-ecommerce/internal/data/entity"

type ReviewResponse struct {
	ID           uint   `json:"id"`
	ProductID    uint   `json:"product_id"`
	ProductName string `json:"product_name,omitempty"`
	OrderID      uint   `json:"order_id"`
	Rating       int    `json:"rating"`
	Comment      string `json:"comment"`
	CustomerName string `json:"customer_name"`
}

type ReviewStatsResponse struct {
	AvgRating float64 `json:"avg_rating"`
	Count     int64   `json:"count"`
}

func ConvertReviewToResponse(r *entity.Review) *ReviewResponse {
	return &ReviewResponse{
		ID:           r.ID,
		ProductID:    r.ProductID,
		ProductName: r.Product.Name,
		OrderID:      r.OrderID,
		Rating:       r.Rating,
		Comment:      r.Comment,
		CustomerName: r.Customer.FullName, // Assuming Customer is preloaded and has Name
	}
}

func ConvertReviewsToResponse(reviews []entity.Review) []ReviewResponse {
	var response []ReviewResponse
	for _, r := range reviews {
		response = append(response, *ConvertReviewToResponse(&r))
	}
	return response
}
