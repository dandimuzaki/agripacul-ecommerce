package repository

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/dto/request"
	infra "debian-ecommerce/internal/infra/transaction"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ReviewRepository interface {
	CreateReview(ctx context.Context, review *entity.Review) error
	BatchCreateReview(ctx context.Context, reviews []entity.Review) error
	GetAllReviews(ctx context.Context, f request.ReviewsFilter) ([]entity.Review, int64, error)
	GetReviewByProductID(ctx context.Context, productID uint, f request.ReviewsFilter) ([]entity.Review, int64, error)
	GetReviewByID(ctx context.Context, id uint) (*entity.Review, error)
	GetReviewByOrderIDAndProductID(ctx context.Context, orderID, productID uint) (*entity.Review, error)
	GetReviewedSKUs(ctx context.Context, customerID uint, orderID uint, skuIDs []uint) (map[uint]bool, error)
	UpdateReview(ctx context.Context, review *entity.Review) error
	DeleteReview(ctx context.Context, id uint) error
	GetReviewStats(ctx context.Context, productID uint) (float64, int64, error)
	IsOrderOwnedByCustomer(ctx context.Context, orderID, customerID uint) (bool, error)
	IsProductInOrder(ctx context.Context, orderID, productID uint) (bool, error)
	IsProductsInOrder(ctx context.Context, orderID uint, productIDs []uint) (bool, error)
}

type reviewRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewReviewRepository(db *gorm.DB, log *zap.Logger) ReviewRepository {
	return &reviewRepository{
		DB:  db,
		Log: log,
	}
}

func (r *reviewRepository) CreateReview(ctx context.Context, review *entity.Review) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Create(review).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}

func (r *reviewRepository) BatchCreateReview(ctx context.Context, reviews []entity.Review) error {
	db := infra.GetDB(ctx, r.DB)
	err := db.Create(&reviews).Error
	if err != nil {
		r.Log.Error("Error query batch create reviews", zap.Error(err))
		return err
	}
	return nil
}

func (r *reviewRepository) GetAllReviews(ctx context.Context, f request.ReviewsFilter) ([]entity.Review, int64, error) {
	db := infra.GetDB(ctx, r.DB)

	// Base query without preloads for counting
	countQuery := db.Model(&entity.Review{})
	
	// Apply filters to count query
	if f.ProductID != nil {
		countQuery = countQuery.Where("product_id = ?", f.ProductID)
	}

	if f.Search != "" {
		searchPattern := "%" + strings.ToLower(f.Search) + "%"
		countQuery = countQuery.Where("LOWER(comment) LIKE ?", searchPattern)
	}

	// Get total count (now without unnecessary JOINs)
	var total int64
	if err := countQuery.Count(&total).Error; err != nil {
		r.Log.Error("Error counting reviews", zap.Error(err))
		return nil, 0, err
	}

	query := db.Model(&entity.Review{}).Preload("Customer").Preload("Product")

	// Sorting
	sortDesc := f.SortOrder == request.SortDesc
	switch f.SortBy {
	case request.SortReviewByRating:
		query = query.Order(clause.OrderByColumn{
			Column: clause.Column{Name: "rating"},
			Desc:   sortDesc,
		})
	case request.SortReviewByCreatedAt:
		query = query.Order(clause.OrderByColumn{
			Column: clause.Column{Name: "created_at"},
			Desc:   sortDesc,
		})
	case request.SortReviewByUpdatedAt:
		query = query.Order(clause.OrderByColumn{
			Column: clause.Column{Name: "updated_at"},
			Desc:   sortDesc,
		})
	default:
		query = query.Order("created_at DESC")
	}

	// Execute main query with pagination
	var reviews []entity.Review
	err := query.Limit(f.Limit).Offset(f.Offset).Find(&reviews).Error
	if err != nil {
		r.Log.Error("Error query get rreview list", zap.Error(err))
		return nil, 0, err
	}

	return reviews, total, nil
}

func (r *reviewRepository) GetReviewByProductID(ctx context.Context, productID uint, f request.ReviewsFilter) ([]entity.Review, int64, error) {
	db := infra.GetDB(ctx, r.DB)

	// Get total count (now without unnecessary JOINs)
	var total int64
	if err := db.Model(&entity.Review{}).Where("product_id = ?", productID).Count(&total).Error; err != nil {
		r.Log.Error("Error counting reviews", zap.Error(err))
		return nil, 0, err
	}

	query := db.Model(&entity.Review{}).Preload("Customer").Where("product_id = ?", productID)

	// Execute main query with pagination
	var reviews []entity.Review
	err := query.Order("created_at DESC").Limit(f.Limit).Offset(f.Offset).Find(&reviews).Error
	if err != nil {
		r.Log.Error("Error query get rreview list", zap.Error(err))
		return nil, 0, err
	}

	return reviews, total, nil
}

func (r *reviewRepository) GetReviewByID(ctx context.Context, id uint) (*entity.Review, error) {
	db := infra.GetDB(ctx, r.DB)
	var review entity.Review
	if err := db.WithContext(ctx).Preload("Customer").Preload("Product").First(&review, id).Error; err != nil {
		r.Log.Error(err.Error())
		return nil, err
	}
	return &review, nil
}

func (r *reviewRepository) GetReviewByOrderIDAndProductID(ctx context.Context, orderID, productID uint) (*entity.Review, error) {
	db := infra.GetDB(ctx, r.DB)
	var review entity.Review
	if err := db.
		Model(&review).
		Joins("JOIN order_items oi ON oi.order_id = reviews.order_id").
		Joins("JOIN skus s ON oi.sku_id = s.id").
		Where("reviews.order_id = ? AND s.product_id = ?", orderID, productID).First(&review).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		r.Log.Error(err.Error())
		return nil, err
	}
	return &review, nil
}

func (r *reviewRepository) GetReviewedSKUs(
    ctx context.Context,
    customerID uint,
    orderID uint,
    skuIDs []uint,
) (map[uint]bool, error) {

	db := infra.GetDB(ctx, r.DB)

	var skuList []uint

	err := db.
		Model(&entity.Review{}).
		Where("customer_id = ? AND order_id = ? AND sku_id IN ?", customerID, orderID, skuIDs).
		Pluck("sku_id", &skuList).Error

	if err != nil {
		r.Log.Error(err.Error())
		return nil, err
	}

	// Convert slice → map for O(1) lookup
	reviewedMap := make(map[uint]bool)
	for _, skuID := range skuList {
		reviewedMap[skuID] = true
	}

	return reviewedMap, nil
}

func (r *reviewRepository) UpdateReview(ctx context.Context, review *entity.Review) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Save(review).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}

func (r *reviewRepository) DeleteReview(ctx context.Context, id uint) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Delete(&entity.Review{}, id).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}

type ReviewStats struct {
	AvgRating float64
	Count     int64
}

func (r *reviewRepository) GetReviewStats(ctx context.Context, productID uint) (float64, int64, error) {
	db := infra.GetDB(ctx, r.DB)
	var stats ReviewStats
	// COALESCE handles null case when there are no reviews
	result := db.WithContext(ctx).Model(&entity.Review{}).
		Select("COALESCE(AVG(rating), 0) as avg_rating, COUNT(*) as count").
		Where("product_id = ?", productID).
		Scan(&stats)

	if result.Error != nil {
		r.Log.Error(result.Error.Error())
		return 0, 0, result.Error
	}
	return stats.AvgRating, stats.Count, nil
}

func (r *reviewRepository) IsOrderOwnedByCustomer(ctx context.Context, orderID, customerID uint) (bool, error) {
	db := infra.GetDB(ctx, r.DB)
	var count int64
	// Check if order exists and belongs to customer
	if err := db.WithContext(ctx).Model(&entity.Order{}).Where("id = ? AND customer_id = ?", orderID, customerID).Count(&count).Error; err != nil {
		r.Log.Error(err.Error())
		return false, err
	}
	return count > 0, nil
}

func (r *reviewRepository) IsProductInOrder(ctx context.Context, orderID, productID uint) (bool, error) {
	db := infra.GetDB(ctx, r.DB)
	var count int64
	// Check if product is in order items
	if err := db.WithContext(ctx).Model(&entity.OrderItem{}).Joins("JOIN skus s ON s.id = order_items.sku_id").Where("order_items.order_id = ? AND s.product_id = ?", orderID, productID).Count(&count).Error; err != nil {
		r.Log.Error(err.Error())
		return false, err
	}
	return count > 0, nil
}

func (r *reviewRepository) IsProductsInOrder(ctx context.Context, orderID uint, productIDs []uint) (bool, error) {
	db := infra.GetDB(ctx, r.DB)
	var count int64
	// Check if product is in order items
	if err := db.WithContext(ctx).Model(&entity.OrderItem{}).Joins("JOIN skus s ON s.id = order_items.sku_id").Where("order_items.order_id = ? AND s.product_id IN ?", orderID, productIDs).Count(&count).Error; err != nil {
		r.Log.Error(err.Error())
		return false, err
	}
	return count == int64(len(productIDs)), nil
}