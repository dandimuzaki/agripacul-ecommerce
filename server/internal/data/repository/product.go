package repository

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/dto/request"
	infra "debian-ecommerce/internal/infra/transaction"
	"debian-ecommerce/pkg/utils"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository interface{
	GetAll(ctx context.Context, f request.ProductQueryParams) ([]entity.Product, int64, error)
	Create(ctx context.Context, product *entity.Product) (*entity.Product, error)
	GetByID(ctx context.Context, id uint) (*entity.Product, error)
	GetProductDetails(ctx context.Context, id uint) (*entity.Product, error)
	GetBySlug(ctx context.Context, slug string) (*entity.Product, error)
	Update(ctx context.Context, id uint, product *entity.Product) error
	BatchUpdate(ctx context.Context, products []entity.Product) error
	UpdatePublish(ctx context.Context, id uint, data map[string]interface{}) error
	Delete(ctx context.Context, id uint) error
	RecalculatePrice(ctx context.Context,	productID uint) error
	BatchGetByIDs(ctx context.Context, ids []uint) ([]entity.Product, error)
}

type productRepository struct {
	db *gorm.DB
	log *zap.Logger
}

func NewProductRepo(db *gorm.DB, log *zap.Logger) ProductRepository {
	return &productRepository{
		db: db,
		log: log,
	}
}

func (r *productRepository) Create(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	db := infra.GetDB(ctx, r.db)
	r.log.Info("Creating product",
		zap.String("name", product.Name),
		zap.Uint("category_id", product.CategoryID),
	)

	err := db.Create(product).Error
	if err != nil {
		r.log.Error("Failed to create product",
			zap.String("name", product.Name),
			zap.Error(err))
		return nil, err
	}

	r.log.Info("Product created successfully",
		zap.Uint("id", product.ID),
		zap.String("name", product.Name))

	return product, nil
}

func (r *productRepository) GetAll(ctx context.Context, f request.ProductQueryParams) ([]entity.Product, int64, error) {
	db := infra.GetDB(ctx, r.db)
	
	// Base query without preloads for counting
	countQuery := db.Model(&entity.Product{})
	
	// Apply filters to count query
	if f.CategoryID != nil {
		countQuery = countQuery.Where("category_id = ?", f.CategoryID)
	}

	if f.Search != "" {
		searchPattern := "%" + strings.ToLower(f.Search) + "%"
		countQuery = countQuery.Where("LOWER(name) LIKE ?", searchPattern)
	}

	// Add JOIN only if needed for filtering
	if f.IsPublishedOnly {
		// Use EXISTS instead of JOIN for better performance with COUNT
		countQuery = countQuery.Where("EXISTS (SELECT 1 FROM skus WHERE skus.product_id = products.id AND skus.status = ?)", entity.SKUStatusActive)
	}

	// Price filters
	if f.MinPrice > 0 {
		countQuery = countQuery.Where("min_price >= ?", f.MinPrice)
	}
	if f.MaxPrice > 0 {
		countQuery = countQuery.Where("max_price <= ?", f.MaxPrice)
	}

	// Rating filter
	if f.Rating > 0 {
		countQuery = countQuery.Where("average_rating >= ? AND average_rating < ?", f.Rating-1, f.Rating)
	}

	// Get total count (now without unnecessary JOINs)
	var total int64
	if err := countQuery.Count(&total).Error; err != nil {
		r.log.Error("Error counting products", zap.Error(err))
		return nil, 0, err
	}

	if total == 0 {
		return []entity.Product{}, 0, nil
	}

	// Build main query with preloads
	query := db.Model(&entity.Product{}).
		Preload("Category").
		Preload("SKUs", func(db *gorm.DB) *gorm.DB {
			// Only load active SKUs and limit fields if needed
			return db.Where("status = ?", entity.SKUStatusActive).Select("id", "product_id", "price", "sale_price", "stock", "sku_code")
		})

	// Apply same filters to main query
	if f.CategoryID != nil {
		query = query.Where("category_id = ?", f.CategoryID)
	}

	if f.Search != "" {
		searchPattern := "%" + strings.ToLower(f.Search) + "%"
		query = query.Where("LOWER(name) LIKE ?", searchPattern)
	}

	// Implement InStockOnly filter
	if f.InStockOnly {
		// Only show products with at least one SKU in stock
		query = query.Where("EXISTS (SELECT 1 FROM skus WHERE skus.product_id = products.id AND skus.stock > 0 AND skus.status = ?)", entity.SKUStatusActive)
	}

	if f.IsPublishedOnly {
		query = query.Where("EXISTS (SELECT 1 FROM skus WHERE skus.product_id = products.id AND skus.status = ?)", entity.SKUStatusActive)
	}

	if f.MinPrice > 0 {
		query = query.Where("min_price >= ?", f.MinPrice)
	}
	if f.MaxPrice > 0 {
		query = query.Where("max_price <= ?", f.MaxPrice)
	}
	if f.Rating > 0 {
		query = query.Where("average_rating >= ? AND average_rating < ?", f.Rating-1, f.Rating)
	}

	// Sorting
	sortDesc := f.SortOrder == request.SortDesc
	switch f.SortBy {
	case request.SortProductByPrice:
		query = query.Order(clause.OrderByColumn{
			Column: clause.Column{Name: "min_price"},
			Desc:   sortDesc,
		})
	case request.SortProductByRating:
		query = query.Order(clause.OrderByColumn{
			Column: clause.Column{Name: "average_rating"},
			Desc:   sortDesc,
		})
	case request.SortProductByCreatedAt:
		query = query.Order(clause.OrderByColumn{
			Column: clause.Column{Name: "created_at"},
			Desc:   sortDesc,
		})
	case request.SortProductBySold:
		query = query.Order(clause.OrderByColumn{
			Column: clause.Column{Name: "sold_count"},
			Desc:   sortDesc,
		})
	default:
		query = query.Order("created_at DESC")
	}

	// Execute main query with pagination
	var products []entity.Product
	err := query.Limit(f.Limit).Offset(f.Offset).Find(&products).Error
	if err != nil {
		r.log.Error("Error query get product list", zap.Error(err))
		return nil, 0, err
	}

	return products, total, nil
}

func (r *productRepository) GetProductDetails(ctx context.Context, id uint) (*entity.Product, error) {
	db := infra.GetDB(ctx, r.db)
	r.log.Info("Get product details",
		zap.Uint("id", id),
	)

	var product entity.Product
	query := db.
  Preload("Category").
  Preload("VariantTypes").
  Preload("VariantTypes.Values").
  Preload("SKUs", "status = ?", entity.SKUStatusActive).
  Preload("SKUs.SKUVariantValues").
  Preload("SKUs.SKUVariantValues.VariantValue").
  Preload("SKUs.Images").
  Preload("Images").
	Where("EXISTS (SELECT 1 FROM skus WHERE skus.product_id = products.id AND skus.status = ?)", entity.SKUStatusActive)

	err := query.First(&product, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			r.log.Warn("Product not found", zap.Uint("id", id))
			return nil, utils.ErrProductNotFound
		} else {
			r.log.Error("Failed to get product",
				zap.Uint("id", id),
				zap.Error(err))
			return nil, err
		}
	}

	return &product, nil
}

func (r *productRepository) GetByID(ctx context.Context, id uint) (*entity.Product, error) {
	db := infra.GetDB(ctx, r.db)
	r.log.Info("Get product by id",
		zap.Uint("id", id),
	)

	var product entity.Product
	query := db.
  Preload("Category").
  Preload("VariantTypes").
  Preload("VariantTypes.Values").
  Preload("SKUs.SKUVariantValues").
  Preload("SKUs.SKUVariantValues.VariantValue").
  Preload("SKUs.Images").
  Preload("Images")

	err := query.First(&product, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			r.log.Warn("Product not found", zap.Uint("id", id))
			return nil, utils.ErrProductNotFound
		} else {
			r.log.Error("Failed to get product",
				zap.Uint("id", id),
				zap.Error(err))
			return nil, err
		}
	}

	return &product, nil
}

func (r *productRepository) BatchGetByIDs(ctx context.Context, ids []uint) ([]entity.Product, error) {
	db := infra.GetDB(ctx, r.db)
	r.log.Info("Get product by ids",
		zap.Any("id", ids),
	)

	var products []entity.Product
	query := db.
  Preload("Category").
  Preload("VariantTypes").
  Preload("VariantTypes.Values").
  Preload("SKUs", "status = ?", entity.SKUStatusActive).
  Preload("SKUs.SKUVariantValues").
  Preload("SKUs.SKUVariantValues.VariantValue").
  Preload("SKUs.Images").
  Preload("Images").Where("id IN ?", ids)

	err := query.Find(&products).Error
	if err != nil {
		r.log.Error("Failed to get products",
			zap.Error(err))
		return nil, err
	}

	return products, nil
}

func (r *productRepository) GetBySlug(ctx context.Context, slug string) (*entity.Product, error) {
	db := infra.GetDB(ctx, r.db)
	r.log.Info("Get product by slug",
		zap.String("slug", slug),
	)

	var product entity.Product
	query := db.Model(&entity.Product{}).
		Preload("Category").
		Preload("VariantTypes").
		Preload("VariantTypes.Values").
		Preload("SKUs.SKUVariantValues").
		Preload("SKUs.SKUVariantValues.VariantValue").
		Preload("SKUs.Images").
		Preload("Images").Where("slug = ?", slug)

	err := query.Find(&product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			r.log.Warn("Product not found", zap.String("slug", slug))
			return nil, utils.ErrProductNotFound
		} else {
			r.log.Error("Failed to get product",
				zap.String("slug", slug),
				zap.Error(err))
			return nil, err
		}
	}

	return &product, nil
}

func (r *productRepository) Update(ctx context.Context, id uint, product *entity.Product) error {
	db := infra.GetDB(ctx, r.db)

	result := db.Model(&entity.Product{}).
		Where("id = ?", id).
		Updates(product)

	if result.Error != nil {
		r.log.Error("Error query update product", zap.Error(result.Error))
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrProductNotFound
	}

	return nil
}

func (r *productRepository) BatchUpdate(ctx context.Context, products []entity.Product) error {
	db := infra.GetDB(ctx, r.db)
	
	if len(products) == 0 {
		return nil
	}

	// Build CASE expression
	ids := make([]uint, len(products))
	avgRatingCase := "CASE"
  reviewCountCase := "CASE"
	soldCountCase := "CASE"
	
	for i, p := range products {
		ids[i] = p.ID

		avgRatingCase += fmt.Sprintf(" WHEN id = %d THEN %f", p.ID, p.AverageRating)
    reviewCountCase += fmt.Sprintf(" WHEN id = %d THEN %d", p.ID, p.ReviewCount)
    soldCountCase += fmt.Sprintf(" WHEN id = %d THEN %d", p.ID, p.SoldCount)
	}
	
	avgRatingCase += " END"
	reviewCountCase  += " END"
	soldCountCase += " END"

	// Update with raw expression
	result := db.Model(&entity.Product{}).
		Where("id IN ?", ids).
		Updates(map[string]interface{}{
			"average_rating": gorm.Expr(avgRatingCase),
			"review_count": gorm.Expr(reviewCountCase),
			"sold_count": gorm.Expr(soldCountCase),
			"updated_at": time.Now(),
		})

	if result.Error != nil {
		r.log.Error("Error batch updating products", zap.Error(result.Error))
		return result.Error
	}

	return nil
}

func (r *productRepository) UpdatePublish(ctx context.Context, id uint, data map[string]interface{}) error {
	db := infra.GetDB(ctx, r.db)

	result := db.Model(&entity.Product{}).
		Where("id = ?", id).
		Updates(data)

	if result.Error != nil {
		r.log.Error("Error query update product", zap.Error(result.Error))
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrProductNotFound
	}

	return nil
}

func (r *productRepository) Delete(ctx context.Context, id uint) error {
	db := infra.GetDB(ctx, r.db)
	err := db.Delete(&entity.Product{}, id).Error
	if err != nil {
		r.log.Error("Error query delete product", zap.Error(err))
		return err
	}
	return nil
}

func (r *productRepository) RecalculatePrice(
	ctx context.Context,
	productID uint,
) error {
	db := infra.GetDB(ctx, r.db)

	var result struct {
		MinPrice float64
		MaxPrice float64
	}

	err := db.
		Model(&entity.SKU{}).
		Select("MIN(price) as min_price, MAX(price) as max_price").
		Where("product_id = ? AND status = ?", productID, entity.SKUStatusActive).
		Scan(&result).Error
	if err != nil {
		r.log.Error("Error query get product", zap.Error(err))
		return err
	}

	err = db.
		Model(&entity.Product{}).
		Where("id = ?", productID).
		Updates(map[string]interface{}{
			"min_price": result.MinPrice,
			"max_price": result.MaxPrice,
		}).Error

	if err != nil {
		r.log.Error("Error query update product", zap.Error(err))
		return err
	}

	return nil
}