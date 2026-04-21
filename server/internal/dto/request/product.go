package request

import (
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/pkg/utils"
	"mime/multipart"
)

type ProductListRequest struct {
	PaginationRequest
	
	// Filtering
	CategoryID *uint   `form:"category_id"`
	Search     string `form:"search"`
	MinPrice  float64 `form:"min_price" binding:"min=0"`
	MaxPrice  float64 `form:"max_price" binding:"min=0"`
	Rating float64 `form:"rating" binding:"min=0,max=5"`

	// Sorting
	SortBy  SortProductOption `form:"sort_by" binding:"omitempty,oneof=id name price rating sold created_at"`
	SortOrder SortOrder `form:"sort_order" binding:"omitempty,oneof=asc desc"`
}

type SortProductOption string
const (
	SortProductByPrice SortProductOption = "price"
	SortProductByRating SortProductOption = "rating"
	SortProductByCreatedAt SortProductOption = "created_at" // terbaru
	SortProductBySold SortProductOption = "sold" // terlaris
	SortProductByName SortProductOption = "name"
	SortProductByID SortProductOption = "id"
)

type SortOrder string
const (
	SortAsc SortOrder = "asc"
	SortDesc SortOrder = "desc"
)

type ProductQueryParams struct {
	// filtering
	CategoryID *uint
	Search     string
	MinPrice   float64
	MaxPrice   float64
	Rating  float64

	// availability
	InStockOnly bool
	IsPublishedOnly bool

	// sorting
	SortBy  SortProductOption
	SortOrder SortOrder

	// pagination
	Page int
	Offset  int
	Limit int
}

func (r ProductListRequest) ToQuery() ProductQueryParams {
	return ProductQueryParams{
		Page: r.GetPage(),
		Offset:        r.GetOffset(),
		Limit:       r.GetPerPage(),
		CategoryID:  r.CategoryID,
		Search:      r.Search,
		MinPrice:    r.MinPrice,
		MaxPrice:    r.MaxPrice,
		Rating:   r.Rating,
		SortBy:      r.SortBy,
		SortOrder:    r.SortOrder,
	}
}

type CreateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	CategoryID  uint   `json:"category_id" validate:"required"`
	Tags []string `json:"tags"`
	Images       []string `json:"images"`
	VariantTypes []CreateVariantTypeRequest `json:"variants" validate:"min=1"`
}

type CreateVariantTypeRequest struct {
	Name   string `json:"name" validate:"required"`
	Values []CreateVariantValueRequest `json:"values" validate:"min=1"`
}

type CreateVariantValueRequest struct {
	Value string `json:"value" validate:"required"`
}

type UpdateProductRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	CategoryID  *uint   `json:"category_id"`

	Images       []string `json:"images"`

	VariantTypes []UpdateVariantTypeRequest `json:"variants"`
}

type UpdateVariantTypeRequest struct {
	ID     *uint   `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status" validate:"oneof=clean created updated deleted"`

	Values []UpdateVariantValueRequest `json:"values"`
}

type UpdateVariantValueRequest struct {
	ID     *uint   `json:"id"`
	Value  string `json:"value"`
	Status string `json:"status" validate:"oneof=clean created updated deleted"`
}

type UpdateSKURequest struct {
	ID            uint    `json:"id"`
	SKUCode       string `json:"sku_code"`
	Price         float64 `json:"price" validate:"gt=0"`
	SalePrice *float64 `json:"sale_price,omitempty"`
	Stock         int     `json:"stock" validate:"gt=0"`
	MinStock int `json:"min_stock"`
	Weight float64 `json:"weight"`
	Status string `json:"status" validate:"oneof=active inactive archived"`
	Images []string `json:"images"`
}

func (r CreateProductRequest) CreateProduct() *entity.Product {
	return &entity.Product{
		Name: r.Name,
		CategoryID: r.CategoryID,
		Description: r.Description,
		Tags: r.Tags,
		Slug: utils.GenerateSlug(r.Name),
	}
}

func (r CreateVariantTypeRequest) CreateVariantType(productID uint) *entity.VariantType {
	return &entity.VariantType{
		ProductID: productID,
		Name: r.Name,
	}
}

func (r CreateVariantTypeRequest) CreateVariantValue(variantTypeID uint, value string) *entity.VariantValue {
	return &entity.VariantValue{
		VariantTypeID: variantTypeID,
		Value: value,
	}
}

func (r CreateProductRequest) CreateProductImages(productID uint) []entity.Image {
	var images []entity.Image
	for _, url := range r.Images {
		img := entity.Image{
			ProductID: productID,
			ImageURL: url,
		}
		images = append(images, img)
	}
	return images
}

func (r UpdateSKURequest) SKUImages(productID uint, SKUID uint) []entity.Image {
	var images []entity.Image
	for _, url := range r.Images {
		img := entity.Image{
			ProductID: productID,
			SKUID: &SKUID,
			ImageURL: url,
		}
		images = append(images, img)
	}
	return images
}

func (r UpdateProductRequest) UpdateProduct() *entity.Product {
	var product entity.Product
	if r.Name != nil {
		product.Name = *r.Name
	}
	if r.Description != nil {
		product.Description = *r.Description
	}
	if r.CategoryID != nil {
		product.CategoryID = *r.CategoryID
	}
	return &product
}

func (r UpdateSKURequest) UpdateSKU(productID uint) *entity.SKU {
	return &entity.SKU{
		ProductID: productID,
		SKUCode: r.SKUCode,
		Price: r.Price,
		SalePrice: r.SalePrice,
		Stock: r.Stock,
		MinStock: r.MinStock,
		Weight: r.Weight,
	}
}

type UpdatePublishRequest struct {
	IsPublished bool `json:"is_published"`
}

type UploadMainImageRequest struct {
	ProductID uint                  `uri:"id" binding:"required"`
	Image     *multipart.FileHeader `form:"image" binding:"required"`
}

type UploadProductImagesRequest struct {
	ProductID uint                     `uri:"id" binding:"required"`
	Images    []*multipart.FileHeader  `form:"images" binding:"required"`
}
