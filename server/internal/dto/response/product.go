package response

import (
	"debian-ecommerce/internal/data/entity"
	"time"
)

type ProductSummary struct {
	ID uint `json:"id"`
	Category    CategoryResponse `json:"category"`
	Name        string           `json:"name"`
	IsPublished bool             `json:"is_published"`
	Tags        []string         `json:"tags"`
	Slug string `json:"slug"`
	MainImageURL string `json:"main_image_url"`
	AverageRating float64 `json:"average_rating"`
	ReviewCount int `json:"review_count"`
	SoldCount int `json:"sold_count"`
	MinPrice float64 `json:"min_price"`
	MaxPrice float64 `json:"max_price"`
	Badge string `json:"badge"`
	SalePrice *float64 `json:"sale_price,omitempty"`
	SalePercentage float64 `json:"sale_percentage,omitempty"`
}

type ProductDetails struct {
	ID uint `json:"id"`
	Category    CategoryResponse `json:"category"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	IsPublished bool             `json:"is_published"`
	Tags        []string         `json:"tags"`
	Variants    []Variant        `json:"variants"`
	SKUs        []SKU            `json:"skus"`
	DefaultSKUID *uint `json:"default_sku_id"`
	Slug string `json:"slug"`
	MainImageURL string `json:"main_image_url"`
	AverageRating float64 `json:"average_rating"`
	ReviewCount int `json:"review_count"`
	SoldCount int `json:"sold_count"`
	MinPrice float64 `json:"min_price"`
	MaxPrice float64 `json:"max_price"`
	Images []Image `json:"images"`
}

type Variant struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Values []Value `json:"values"`
}

type Value struct {
	ID    uint   `json:"id"`
	Value string `json:"value"`
}

type SKU struct {
	ID              uint             `json:"id"`
	SKUCode         string           `json:"sku_code"`
	Price float64 `json:"price"`
	SalePrice *float64 `json:"sale_price,omitempty"`
	Stock int `json:"stock"`
	MinStock int `json:"min_stock"`
	Weight float64 `json:"weight"`
	Status          entity.SKUStatus `json:"status"`
	VariantValueIDs []uint           `json:"variant_value_ids"`
	Images []Image `json:"images"`
	IsAvailable bool `json:"is_available"`
}

type Image struct {
	ID       uint   `json:"id"`
	ImageURL string `json:"image_url"`
}

type SKUDetails struct {
	ID uint `json:"id"`
	ProductID uint `json:"product_id"`
	SKUCode string `json:"sku_code"`
	Price float64 `json:"price"`
	SalePrice *float64 `json:"sale_price,omitempty"`
	Stock int `json:"stock"`
	MinStock int `json:"min_stock"`
	Status entity.SKUStatus `json:"status"`
	Weight float64 `json:"weight"`
	Variants []entity.VariantCombination `json:"variants"`
	Images []string `json:"images"`
}

func ToProductDetails(p *entity.Product) *ProductDetails {
	var variantTypes []Variant
	for _, t := range p.VariantTypes {
		var values []Value
		for _, v := range t.Values {
			values = append(values, Value{
				ID: v.ID,
				Value: v.Value,
			})
		}
		variantTypes = append(variantTypes, Variant{
			ID: t.ID,
			Name: t.Name,
			Values: values,
		})
	}

	var images []Image
	for _, img := range p.Images {
		images = append(images, Image{
			ID: img.ID,
			ImageURL: img.ImageURL,
		})
	}

	var skus []SKU
	for _, s := range p.SKUs {
		var variantValueIDs []uint
		for _, v := range s.SKUVariantValues {
			variantValueIDs = append(variantValueIDs, v.VariantValueID)
		}
		var images []Image
		for _, img := range s.Images {
			images = append(images, Image{
				ID: img.ID,
				ImageURL: img.ImageURL,
			})
		}
		skus = append(skus, SKU{
			ID: s.ID,
			SKUCode: s.SKUCode,
			Price: s.Price,
			Stock: s.Stock,
			Status: s.Status,
			VariantValueIDs: variantValueIDs,
			Images: images,
			IsAvailable: s.Stock > 0,
		})
	}
	
	res := ProductDetails{
		ID: p.ID,
		Category: CategoryResponse{
			ID: p.CategoryID,
			Name: p.Category.Name,
			IconURL: p.Category.IconURL,
		},
		Name: p.Name,
		Description: p.Description,
		IsPublished: p.IsPublished,
		Tags: p.Tags,
		Slug: p.Slug,
		MainImageURL: p.MainImageURL,
		AverageRating: p.AverageRating,
		ReviewCount: p.ReviewCount,
		SoldCount: p.SoldCount,
		MinPrice: p.MinPrice,
		MaxPrice: p.MaxPrice,
		Variants: variantTypes,
		SKUs: skus,
		Images: images,
	}

	if len(skus) > 0 {
		res.DefaultSKUID = &skus[0].ID
	}

	return &res
}

func Badge(createdAt time.Time) string {
	if time.Now().Before(createdAt.AddDate(0,0,30)) {
		return "new"
	} else {
		return "regular"
	}
}

func ToProductSummary(p *entity.Product) *ProductSummary {
	return &ProductSummary{
		ID: p.ID,
		Category: CategoryResponse{
			ID: p.CategoryID,
			Name: p.Category.Name,
			IconURL: p.Category.IconURL,
		},
		Name: p.Name,
		IsPublished: p.IsPublished,
		Tags: p.Tags,
		Slug: p.Slug,
		MainImageURL: p.MainImageURL,
		AverageRating: p.AverageRating,
		ReviewCount: p.ReviewCount,
		SoldCount: p.SoldCount,
		MinPrice: p.MinPrice,
		MaxPrice: p.MaxPrice,
		Badge: Badge(p.CreatedAt),
	}
}