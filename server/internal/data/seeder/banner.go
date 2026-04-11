package seeder

import (
	"time"

	"debian-ecommerce/internal/data/entity"

	"gorm.io/gorm"
)

type BannerSeeder struct {
	DB *gorm.DB
}

func NewBannerSeeder(db *gorm.DB) *BannerSeeder {
	return &BannerSeeder{DB: db}
}

func (s *BannerSeeder) Run() error {
	var banners []entity.Banner
	if err := s.DB.Find(&banners).Error; err != nil {
		return err
	}

	if len(banners) > 0 {
		return nil // Sudah ada banners, skip seeding
	}

	now := time.Now()

	banners = []entity.Banner{
		// Banner utama - aktif
		{
			Name:        "Grand Opening Sale 50% Off",
			ImageURL:    "https://example.com/banners/main-sale-2024.jpg",
			TargetURL:   "/promo/grand-opening",
			StartDate:   now.Add(-7 * 24 * time.Hour),
			EndDate:     now.Add(7 * 24 * time.Hour),
			Type:        entity.BannerMain,
			IsPublished: true,
		},
		// Banner utama - akan datang
		{
			Name:        "Black Friday Sale Coming Soon",
			ImageURL:    "https://example.com/banners/black-friday-coming.jpg",
			TargetURL:   "/promo/black-friday",
			StartDate:   now.Add(10 * 24 * time.Hour),
			EndDate:     now.Add(25 * 24 * time.Hour),
			Type:        entity.BannerMain,
			IsPublished: true,
		},
		// Banner utama - sudah berakhir
		{
			Name:        "New Year Sale 2023",
			ImageURL:    "https://example.com/banners/new-year-2023.jpg",
			TargetURL:   "/promo/new-year",
			StartDate:   now.Add(-60 * 24 * time.Hour),
			EndDate:     now.Add(-30 * 24 * time.Hour),
			Type:        entity.BannerMain,
			IsPublished: false,
		},
		// Banner sekunder - elektronik
		{
			Name:        "Composter Flash Sale",
			ImageURL:    "https://example.com/banners/composter-flash.jpg",
			TargetURL:   "/categories/composter",
			StartDate:   now,
			EndDate:     now.Add(2 * 24 * time.Hour),
			Type:        entity.BannerSecondary,
			IsPublished: true,
		},
		// Banner sekunder - fashion
		{
			Name:        "Cutlery Collection 2024",
			ImageURL:    "https://example.com/banners/cutlery-2024.jpg",
			TargetURL:   "/categories/cutlery",
			StartDate:   now.Add(-15 * 24 * time.Hour),
			EndDate:     now.Add(45 * 24 * time.Hour),
			Type:        entity.BannerSecondary,
			IsPublished: true,
		},
		// Banner sekunder - makanan
		{
			Name:        "Premium Food Packaging Products",
			ImageURL:    "https://example.com/banners/food-packaging.jpg",
			TargetURL:   "/categories/food-packaging",
			StartDate:   now.Add(-3 * 24 * time.Hour),
			EndDate:     now.Add(27 * 24 * time.Hour),
			Type:        entity.BannerSecondary,
			IsPublished: true,
		},
		// Banner sekunder - kesehatan
		{
			Name:        "Health & Beauty Week",
			ImageURL:    "https://example.com/banners/health-beauty.jpg",
			TargetURL:   "/categories/health",
			StartDate:   now.Add(-1 * 24 * time.Hour),
			EndDate:     now.Add(6 * 24 * time.Hour),
			Type:        entity.BannerSecondary,
			IsPublished: true,
		},
		// Banner tidak aktif
		{
			Name:        "Test Banner (Inactive)",
			ImageURL:    "https://example.com/banners/test-banner.jpg",
			TargetURL:   "/test",
			StartDate:   now.Add(-1 * 24 * time.Hour),
			EndDate:     now.Add(1 * 24 * time.Hour),
			Type:        entity.BannerSecondary,
			IsPublished: false,
		},
		// Banner untuk event khusus
		{
			Name:        "Free Shipping Weekend",
			ImageURL:    "https://example.com/banners/free-shipping.jpg",
			TargetURL:   "/promo/free-shipping",
			StartDate:   now.Add(-1 * 24 * time.Hour),
			EndDate:     now.Add(2 * 24 * time.Hour),
			Type:        entity.BannerMain,
			IsPublished: true,
		},
		// Banner promosi member
		{
			Name:        "Member Exclusive Benefits",
			ImageURL:    "https://example.com/banners/member-exclusive.jpg",
			TargetURL:   "/membership",
			StartDate:   now.Add(-30 * 24 * time.Hour),
			EndDate:     now.Add(60 * 24 * time.Hour),
			Type:        entity.BannerSecondary,
			IsPublished: true,
		},
	}

	// Hapus data lama
	if err := s.DB.Exec("DELETE FROM banners").Error; err != nil {
		return err
	}

	// Create banners
	for i := range banners {
		if err := s.DB.Create(&banners[i]).Error; err != nil {
			return err
		}
	}

	return nil
}
