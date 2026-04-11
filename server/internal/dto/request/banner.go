package request

import (
	"debian-ecommerce/internal/data/entity"
	"mime/multipart"
	"time"
)

type BannerRequest struct {
	Name        string                `form:"name"`
	Image       *multipart.FileHeader `form:"image" binding:"required"`
	TargetURL   string                `form:"target_url,omitempty"`
	StartDate   time.Time             `form:"start_date"`
	EndDate     time.Time             `form:"end_date"`
	Type        entity.BannerType     `form:"type"`
	IsPublished bool                  `form:"is_published"`
}

func (r BannerRequest) ToBanner() entity.Banner {
	return entity.Banner{
		Name: r.Name,
		TargetURL: r.TargetURL,
		StartDate: r.StartDate,
		EndDate: r.EndDate,
		Type: r.Type,
		IsPublished: r.IsPublished,
	}
}