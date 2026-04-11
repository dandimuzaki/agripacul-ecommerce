package request

type BatchCreateReview struct {
	OrderID uint           `json:"order_id"`
	Reviews []CreateReview `json:"reviews"`
}

type CreateReview struct {
	SKUID   uint   `json:"sku_id" binding:"required"`
	Rating  int    `json:"rating" binding:"required,min=1,max=5"`
	Comment string `json:"comment"`
}

type CreateReviewRequest struct {
	ProductID uint   `json:"product_id" binding:"required"`
	OrderID   uint   `json:"order_id" binding:"required"`
	Rating    int    `json:"rating" binding:"required,min=1,max=5"`
	Comment   string `json:"comment"`
}

type UpdateReviewRequest struct {
	Comment string `json:"comment" binding:"required"`
}

type ReviewsQuery struct {
	PaginationRequest

	Search    string `form:"search"`
	ProductID *uint  `form:"product_id"`

	SortBy    string    `form:"sort_by"`
	SortOrder SortOrder `form:"sort_order"`
}

type ReviewsFilter struct {
	Page   int
	Limit  int
	Offset int

	Search    string
	ProductID *uint
	SortBy    SortReviewOption
	SortOrder SortOrder
}

type SortReviewOption string

var (
	SortReviewByCreatedAt SortReviewOption = "created_at"
	SortReviewByUpdatedAt SortReviewOption = "updated_at"
	SortReviewByRating    SortReviewOption = "rating"
)

func ToReviewsFilter(q *ReviewsQuery) *ReviewsFilter {
	return &ReviewsFilter{
		Page:      q.GetPage(),
		Limit:     q.GetPerPage(),
		Offset:    q.GetOffset(),
		Search:    q.Search,
		ProductID: q.ProductID,
		SortBy:    SortReviewOption(q.SortBy),
		SortOrder: q.SortOrder,
	}
}
