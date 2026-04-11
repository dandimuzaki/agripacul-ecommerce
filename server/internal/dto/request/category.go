package request

import (
	"mime/multipart"

	"github.com/go-playground/validator/v10"
)

type CreateCategoryRequest struct {
	Name    string `form:"name" binding:"required,min=3,max=100"`
	Icon *multipart.FileHeader `form:"icon"`
}

type UpdateCategoryRequest struct {
	Name    string `form:"name" binding:"required,min=3,max=100"`
	Icon *multipart.FileHeader `form:"icon"`
}

// Validate melakukan validasi tambahan jika diperlukan
func (req *CreateCategoryRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(req)
}

func (req *UpdateCategoryRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(req)
}

type SortOption string
const (
	SortByID SortOption = "id"
	SortByName SortOption = "name"
	SortByCreatedAt SortOption = "created_at"
	SortByUpdatedAt SortOption = "updated_at"
)

type CategoryListRequest struct {
	PaginationRequest
	
	// Filtering
	Search     string `form:"search"`

	// Sorting
	SortBy  SortOption `form:"sort_by" binding:"omitempty,oneof=id name created_at updated_at"`
	SortOrder SortOrder `form:"sort_order" binding:"omitempty,oneof=asc desc"`

	WithProductCount bool `form:"with_product_count"`
}

type CategoryQueryParams struct {
	// filtering
	Search     string

	// sorting
	SortBy  SortOption
	SortOrder SortOrder

	// pagination
	Page int
	Offset  int
	Limit int
}

func ToQuery(r CategoryListRequest) CategoryQueryParams {
	return CategoryQueryParams{
		Page: r.GetPage(),
		Offset:        r.GetOffset(),
		Limit:       r.GetPerPage(),
		Search:      r.Search,
		SortBy:      r.SortBy,
		SortOrder:    r.SortOrder,
	}
}
