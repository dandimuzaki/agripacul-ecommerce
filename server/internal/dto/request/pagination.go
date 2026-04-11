package request

type PaginationRequest struct {
	Page  int `form:"page,default=1" binding:"min=1"`
	Limit int `form:"limit,default=20" binding:"min=1,max=100"`
}

// Getter dengan default values
func (p PaginationRequest) GetPage() int {
	if p.Page < 1 {
		return 1
	}
	return p.Page
}

func (p PaginationRequest) GetPerPage() int {
	if p.Limit < 1 {
		return 10 // DEFAULT 10
	}
	if p.Limit > 100 {
		return 100 // MAX 100
	}
	return p.Limit
}

func (p PaginationRequest) GetOffset() int {
	return (p.GetPage() - 1) * p.GetPerPage()
}

type Pagination struct {
	Page      int    `json:"page" form:"page" validate:"min=1"`
	Limit     int    `json:"limit" form:"limit" validate:"min=1,max=100"`
	SortBy    string `json:"sort_by" form:"sort_by"`
	SortOrder string `json:"sort_order" form:"sort_order" validate:"omitempty,oneof=asc desc"`
}