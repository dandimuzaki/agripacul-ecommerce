package usecase

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/data/repository"
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/dto/response"
	"debian-ecommerce/pkg/utils"
	"errors"
	"fmt"

	"go.uber.org/zap"
)

type ReviewUsecase interface {
	CreateReview(ctx context.Context, userID uint, req request.CreateReviewRequest) (*response.ReviewResponse, error)
	BatchCreateReview(ctx context.Context, req request.BatchCreateReview) (error)
	GetAllReviews(ctx context.Context, f request.ReviewsFilter) (*response.PaginatedResponse[response.ReviewResponse], error)
	GetReviewsByProduct(ctx context.Context, productID uint, f request.ReviewsFilter) (*response.PaginatedResponse[response.ReviewResponse], error)
	GetReviewDetails(ctx context.Context, id uint) (*response.ReviewResponse, error)
	UpdateReviewComment(ctx context.Context, userID, id uint, req request.UpdateReviewRequest) (*response.ReviewResponse, error)
	DeleteReview(ctx context.Context, userID, id uint) error
	GetReviewStats(ctx context.Context, productID uint) (*response.ReviewStatsResponse, error)
}

type reviewUsecase struct {
	tx TxManager
	repo *repository.Repository
	log          *zap.Logger
}

func NewReviewUsecase(tx TxManager, repo *repository.Repository, log *zap.Logger) ReviewUsecase {
	return &reviewUsecase{
		tx: tx,
		repo: repo,
		log:          log,
	}
}

func (u *reviewUsecase) getCustomerID(ctx context.Context, userID uint) (uint, error) {
	customer, err := u.repo.CustomerRepo.FindCustomerByUserID(ctx, userID)
	if err != nil {
		return 0, utils.WrapError("Customer not found", err)
	}
	return customer.ID, nil
}

func (u *reviewUsecase) CreateReview(ctx context.Context, userID uint, req request.CreateReviewRequest) (*response.ReviewResponse, error) {
	var review entity.Review
	
	err := u.tx.WithinTx(ctx, func(ctx context.Context) error {
	customerID, err := u.getCustomerID(ctx, userID)
	if err != nil {
		return err
	}

	// 1. Verify order ownership
	owned, err := u.repo.ReviewRepo.IsOrderOwnedByCustomer(ctx, req.OrderID, customerID)
	if err != nil {
		return utils.WrapError("Failed to check order ownership", err)
	}
	if !owned {
		return utils.NewAppError(utils.ErrCodeForbidden, "Order does not belong to customer", errors.New("forbidden"))
	}

	// 2. Verify product in order
	inOrder, err := u.repo.ReviewRepo.IsProductInOrder(ctx, req.OrderID, req.ProductID)
	if err != nil {
		return utils.WrapError("Failed to check product in order", err)
	}
	if !inOrder {
		return utils.NewAppError(utils.ErrCodeBadRequest, "Product not found in this order", errors.New("bad request"))
	}

	// 3. Check for existing review
	existing, err := u.repo.ReviewRepo.GetReviewByOrderIDAndProductID(ctx, req.OrderID, req.ProductID)
	if err != nil {
		return utils.WrapError("Failed to check existing review", err)
	}
	if existing != nil {
		return utils.NewAppError(utils.ErrCodeConflict, "Review already exists for this product in this order", errors.New("conflict"))
	}

	review = entity.Review{
		CustomerID: customerID,
		ProductID:  req.ProductID,
		OrderID:    req.OrderID,
		Rating:     req.Rating,
		Comment:    req.Comment,
	}

	if err := u.repo.ReviewRepo.CreateReview(ctx, &review); err != nil {
		return utils.WrapError("Failed to create review", err)
	}

	// Assuming we want to return the response with customer name, we might need to load it.
	// But getting customerID implies we found the customer.
	// We can manually populate Customer field partially if needed, or re-fetch.
	// For simplicity, let's re-fetch details or manually set customer name if we had it.
	// We don't have customer name easily unless we fetched full customer obj.
	// FindCustomerByUserID returns *entity.Customer.

	// Let's optimize: getCustomerID calls FindCustomerByUserID which returns obj.
	// We should reuse it.
	customer, err := u.repo.CustomerRepo.FindCustomerByUserID(ctx, userID)
	if err == nil {
		review.Customer = *customer
	}

	// Update product rating_average
	product, err := u.repo.ProductRepo.GetByID(ctx, req.ProductID)
	if err != nil {
		return err
	}
	product.AverageRating = (product.AverageRating * float64(product.ReviewCount) + float64(req.Rating)) / (float64(product.ReviewCount) + 1)
	product.ReviewCount += 1
	if err = u.repo.ProductRepo.Update(ctx, req.ProductID, product); err != nil {
		return err
	}

	return err
	})

	if err != nil {
		return nil, err
	}

	return response.ConvertReviewToResponse(&review), nil
}

func (u *reviewUsecase) BatchCreateReview(ctx context.Context, req request.BatchCreateReview) (error) {
	err := u.tx.WithinTx(ctx, func(ctx context.Context) error {
	customer, err := GetCustomerFromContext(ctx, u.repo)
	if err != nil {
		return err
	}

	// 1. Verify order ownership
	owned, err := u.repo.ReviewRepo.IsOrderOwnedByCustomer(ctx, req.OrderID, customer.ID)
	if err != nil {
		return utils.WrapError("Failed to check order ownership", err)
	}
	if !owned {
		return utils.NewAppError(utils.ErrCodeForbidden, "Order does not belong to customer", errors.New("forbidden"))
	}

	// 2. Verify SKUs in order
	skuMap, err := u.repo.OrderRepo.GetOrderSKUsMap(ctx, req.OrderID)
	if err != nil {
		return err
	}
	var skuIDs []uint
	for _, rev := range req.Reviews {
		skuIDs = append(skuIDs, rev.SKUID)
		if _, exists := skuMap[rev.SKUID]; !exists {
			return utils.NewAppError(utils.ErrCodeBadRequest, "Product not found in this order", errors.New("bad request"))
		}
	}

	// 3. Check for existing review
	reviewedMap, err := u.repo.ReviewRepo.GetReviewedSKUs(ctx, customer.ID, req.OrderID, skuIDs)
	if err != nil {
    return err
	}

	var alreadyReviewed []uint

	for _, rev := range req.Reviews {
		if reviewedMap[rev.SKUID] {
			alreadyReviewed = append(alreadyReviewed, rev.SKUID)
		}
	}

	if len(alreadyReviewed) > 0 {
		return fmt.Errorf("some SKUs already reviewed: %v", alreadyReviewed)
	}

	// Initialize review data
	var reviews []entity.Review
	for _, rev := range req.Reviews {
		review := entity.Review{
			CustomerID: customer.ID,
			ProductID: skuMap[rev.SKUID],
			SKUID: rev.SKUID,
			OrderID: req.OrderID,
			Rating: rev.Rating,
			Comment: rev.Comment,
		}

		reviews = append(reviews, review)
	}

	if err := u.repo.ReviewRepo.BatchCreateReview(ctx, reviews); err != nil {
		return utils.WrapError("Failed to create review", err)
	}

	type Agg struct {
    TotalRating int
    Count       int
	}

	aggMap := make(map[uint]*Agg) // product_id → aggregation

	for _, rev := range req.Reviews {
		productID, exists := skuMap[rev.SKUID]
		if !exists {
			return fmt.Errorf("Product id missing for sku_id %d", rev.SKUID)
		}

		if aggMap[productID] == nil {
				aggMap[productID] = &Agg{}
		}

		aggMap[productID].TotalRating += rev.Rating
		aggMap[productID].Count++
	}

	var productIDs []uint
	for pid := range aggMap {
			productIDs = append(productIDs, pid)
	}

	products, err := u.repo.ProductRepo.BatchGetByIDs(ctx, productIDs)
	if err != nil {
		return err
	}

	if len(products) != len(aggMap) {
		return errors.New("product aggregation mismatch")
	}

	for _, product := range products {
    agg, exists := aggMap[product.ID]
		if !exists {
			return fmt.Errorf("aggregation missing for product_id %d", product.ID)		
		}

    totalRating := product.AverageRating * float64(product.ReviewCount)
    totalRating += float64(agg.TotalRating)

    product.ReviewCount += agg.Count
    product.AverageRating = totalRating / float64(product.ReviewCount)
	}

	err = u.repo.ProductRepo.BatchUpdate(ctx, products)
	if err != nil {
		return err
	}

	return nil
	})

	if err != nil {
		u.log.Error("Failed to create reviews", zap.Error(err))
		return err
	}

	return nil
}

func (u *reviewUsecase) GetAllReviews(ctx context.Context, f request.ReviewsFilter) (*response.PaginatedResponse[response.ReviewResponse], error) {
	reviews, total, err := u.repo.ReviewRepo.GetAllReviews(ctx, f)
	if err != nil {
		return nil, utils.WrapError("Failed to get reviews", err)
	}
	data := response.ConvertReviewsToResponse(reviews)

	return response.NewPaginatedResponse(
		data,
		f.Page,
		f.Limit,
		total,
	), nil
}

func (u *reviewUsecase) GetReviewsByProduct(ctx context.Context, productID uint, f request.ReviewsFilter) (*response.PaginatedResponse[response.ReviewResponse], error) {
	reviews, total, err := u.repo.ReviewRepo.GetReviewByProductID(ctx, productID, f)
	if err != nil {
		return nil, utils.WrapError("Failed to get reviews", err)
	}
	data := response.ConvertReviewsToResponse(reviews)

	return response.NewPaginatedResponse(
		data,
		f.Page,
		f.Limit,
		total,
	), nil
}

func (u *reviewUsecase) GetReviewDetails(ctx context.Context, id uint) (*response.ReviewResponse, error) {
	review, err := u.repo.ReviewRepo.GetReviewByID(ctx, id)
	if err != nil {
		return nil, utils.WrapError("Failed to get review", err)
	}
	return response.ConvertReviewToResponse(review), nil
}

func (u *reviewUsecase) UpdateReviewComment(ctx context.Context, userID, id uint, req request.UpdateReviewRequest) (*response.ReviewResponse, error) {
	customerID, err := u.getCustomerID(ctx, userID)
	if err != nil {
		return nil, err
	}

	review, err := u.repo.ReviewRepo.GetReviewByID(ctx, id)
	if err != nil {
		return nil, utils.WrapError("Failed to get review", err)
	}

	if review.CustomerID != customerID {
		return nil, utils.NewAppError(utils.ErrCodeForbidden, "You can only update your own reviews", errors.New("forbidden"))
	}

	review.Comment = req.Comment
	if err := u.repo.ReviewRepo.UpdateReview(ctx, review); err != nil {
		return nil, utils.WrapError("Failed to update review", err)
	}

	return response.ConvertReviewToResponse(review), nil
}

func (u *reviewUsecase) DeleteReview(ctx context.Context, userID, id uint) error {
	customerID, err := u.getCustomerID(ctx, userID)
	if err != nil {
		return err
	}

	review, err := u.repo.ReviewRepo.GetReviewByID(ctx, id)
	if err != nil {
		return utils.WrapError("Failed to get review", err)
	}

	if review.CustomerID != customerID {
		return utils.NewAppError(utils.ErrCodeForbidden, "You can only delete your own reviews", errors.New("forbidden"))
	}

	if err := u.repo.ReviewRepo.DeleteReview(ctx, id); err != nil {
		return utils.WrapError("Failed to delete review", err)
	}
	return nil
}

func (u *reviewUsecase) GetReviewStats(ctx context.Context, productID uint) (*response.ReviewStatsResponse, error) {
	avg, count, err := u.repo.ReviewRepo.GetReviewStats(ctx, productID)
	if err != nil {
		return nil, utils.WrapError("Failed to get stats", err)
	}
	return &response.ReviewStatsResponse{
		AvgRating: avg,
		Count:     count,
	}, nil
}
