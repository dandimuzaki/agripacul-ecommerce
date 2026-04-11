package response

import (
	"time"

	"debian-ecommerce/internal/data/entity"
	// HAPUS: "debian-ecommerce/internal/data/repository" // Ini penyebab cycle!
)

// CategoryResponse untuk response single category
type CategoryResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	IconURL   string    `json:"icon_url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CategoryWithProductCountResponse untuk response dengan jumlah produk
type CategoryWithProductCountResponse struct {
	CategoryResponse
	ProductCount int64 `json:"product_count"`
}

// CategoriesListResponse untuk response list categories dengan pagination
type CategoriesListResponse struct {
	Categories []CategoryResponse `json:"categories"`
	Pagination PaginationResponse `json:"pagination"`
}

// CategoriesWithProductCountListResponse untuk response list dengan product count
type CategoriesWithProductCountListResponse struct {
	Categories []CategoryWithProductCountResponse `json:"categories"`
	Pagination PaginationResponse                 `json:"pagination"`
}

// PaginationResponse untuk response pagination
type PaginationResponse struct {
	CurrentPage int   `json:"current_page"`
	TotalPages  int   `json:"total_pages"`
	TotalItems  int64 `json:"total_items"`
	PerPage     int   `json:"per_page"`
}

// ConvertEntityToResponse mengkonversi entity.Category ke CategoryResponse
func ConvertEntityToResponse(category *entity.Category) *CategoryResponse {
	return &CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		IconURL:   category.IconURL,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

// ConvertEntityListToResponse mengkonversi list entity.Category ke list CategoryResponse
func ConvertEntityListToResponse(categories []entity.Category) []CategoryResponse {
	responses := make([]CategoryResponse, len(categories))
	for i, category := range categories {
		responses[i] = *ConvertEntityToResponse(&category)
	}
	return responses
}

// ConvertEntityToWithProductCountResponse mengkonversi entity + product count
func ConvertEntityToWithProductCountResponse(category *entity.Category, productCount int64) *CategoryWithProductCountResponse {
	return &CategoryWithProductCountResponse{
		CategoryResponse: CategoryResponse{
			ID:        category.ID,
			Name:      category.Name,
			IconURL:   category.IconURL,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		},
		ProductCount: productCount,
	}
}

// NewCategoriesListResponse membuat response untuk list categories
func NewCategoriesListResponse(categories []entity.Category, pagination PaginationResponse) *CategoriesListResponse {
	return &CategoriesListResponse{
		Categories: ConvertEntityListToResponse(categories),
		Pagination: pagination,
	}
}

// HAPUS fungsi NewCategoriesWithProductCountListResponse atau pindahkan ke usecase
// func NewCategoriesWithProductCountListResponse(categories []repository.CategoryWithProductCount, pagination PaginationResponse) *CategoriesWithProductCountListResponse {
//     responses := make([]CategoryWithProductCountResponse, len(categories))
//     for i, cat := range categories {
//         responses[i] = CategoryWithProductCountResponse{
//             CategoryResponse: CategoryResponse{
//                 ID:        cat.ID,
//                 Name:      cat.Name,
//                 IconURL:   cat.IconURL,
//                 CreatedAt: cat.CreatedAt,
//                 UpdatedAt: cat.UpdatedAt,
//             },
//             ProductCount: cat.ProductCount,
//         }
//     }
//     return &CategoriesWithProductCountListResponse{
//         Categories: responses,
//         Pagination: pagination,
//     }
// }
