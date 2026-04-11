package repository

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/dto/request"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll(ctx context.Context, req request.CategoryQueryParams) ([]entity.Category, int64, error)
	FindByID(id uint) (*entity.Category, error)
	FindByIDs(ids []uint) ([]entity.Category, error)
	Create(category *entity.Category) error
	Update(category *entity.Category) error
	Delete(id uint) error
	CountAll() (int64, error)
	GetProductCount(categoryID uint) (int64, error)
	FindAllWithProductCount(ctx context.Context, req request.CategoryQueryParams) ([]CategoryWithProductCount, int64, error)
}

type categoryRepository struct {
	db *gorm.DB
}

type CategoryWithProductCount struct {
	entity.Category
	ProductCount int64 `json:"product_count"`
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) FindAll(ctx context.Context, req request.CategoryQueryParams) ([]entity.Category, int64, error) {
	var categories []entity.Category
	var total int64

	query := r.db.Model(&entity.Category{})

	if req.Search != "" {
		query = query.Where("name ILIKE ?", "%"+req.Search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if req.SortBy == "" {
		req.SortBy = "created_at"
	}
	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}
	orderClause := string(req.SortBy) + " " + string(req.SortOrder)

	offset := (req.Page - 1) * req.Limit
	if err := query.
		Order(orderClause).
		Offset(offset).
		Limit(req.Limit).
		Find(&categories).Error; err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}

func (r *categoryRepository) FindAllWithProductCount(ctx context.Context, req request.CategoryQueryParams) ([]CategoryWithProductCount, int64, error) {
	var results []CategoryWithProductCount
	var total int64

	subQuery := r.db.Model(&entity.Product{}).
		Select("category_id, COUNT(*) as product_count").
		Where("is_published = ?", true).
		Group("category_id")

	query := r.db.Model(&entity.Category{}).
		Select("categories.*, COALESCE(pc.product_count, 0) as product_count").
		Joins("LEFT JOIN (?) as pc ON pc.category_id = categories.id", subQuery)

	if req.Search != "" {
		query = query.Where("categories.name ILIKE ?", "%"+req.Search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if req.SortBy == "" {
		req.SortBy = "created_at"
	}
	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}
	orderClause := "categories." + string(req.SortBy) + " " + string(req.SortOrder)

	if err := query.
		Order(orderClause).
		Offset(req.Offset).
		Limit(req.Limit).
		Scan(&results).Error; err != nil {
		return nil, 0, err
	}

	return results, total, nil
}

func (r *categoryRepository) FindByID(id uint) (*entity.Category, error) {
	var category entity.Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) FindByIDs(ids []uint) ([]entity.Category, error) {
	var categories []entity.Category
	if err := r.db.Where("id IN ?", ids).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) Create(category *entity.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) Update(category *entity.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Category{}, id).Error
}

func (r *categoryRepository) CountAll() (int64, error) {
	var count int64
	if err := r.db.Model(&entity.Category{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *categoryRepository) GetProductCount(categoryID uint) (int64, error) {
	var count int64
	if err := r.db.Model(&entity.Product{}).
		Where("category_id = ? AND is_published = ?", categoryID, true).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
