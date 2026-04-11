package seeder

import (
	"debian-ecommerce/internal/data/entity"

	"gorm.io/gorm"
)

type CategorySeeder struct {
	DB *gorm.DB
}

func NewCategorySeeder(db *gorm.DB) *CategorySeeder {
	return &CategorySeeder{DB: db}
}

func (s *CategorySeeder) Run() error {
	var categories []entity.Category
	if err := s.DB.Find(&categories).Error; err != nil {
		return err
	}

	if len(categories) >= 3 {
		return nil // Sudah ada categories, skip seeding
	}

	categories = []entity.Category{
		{
			Name:    "Food Packaging",
			IconURL: "https://example.com/icons/food-packaging.svg",
		},
		{
			Name:    "Gardening Tools",
			IconURL: "https://example.com/icons/gardening-tools.svg",
		},
		{
			Name:    "Hair Care",
			IconURL: "https://example.com/icons/food-packaging.svg",
		},
		{
			Name:    "Dental Care",
			IconURL: "https://example.com/icons/women-cares.svg",
		},
		{
			Name:    "Gardening Tools",
			IconURL: "https://example.com/icons/gardening-tools.svg",
		},
	}

	// Gunakan Upsert (Create or Update) untuk menghindari duplikasi
	for i := range categories {
		var existingCategory entity.Category

		// Cek apakah kategori sudah ada berdasarkan nama
		err := s.DB.Where("name = ?", categories[i].Name).First(&existingCategory).Error

		if err != nil && err == gorm.ErrRecordNotFound {
			// Jika tidak ada, create baru
			if err := s.DB.Create(&categories[i]).Error; err != nil {
				return err
			}
		} else if err == nil {
			// Jika sudah ada, update icon_url jika diperlukan
			if existingCategory.IconURL != categories[i].IconURL {
				existingCategory.IconURL = categories[i].IconURL
				if err := s.DB.Save(&existingCategory).Error; err != nil {
					return err
				}
			}
		} else {
			return err
		}
	}

	return nil
}
