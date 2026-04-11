package usecase

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/data/repository"
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/dto/response"
	"debian-ecommerce/pkg/utils"
	"fmt"

	"go.uber.org/zap"
)

type CategoryUsecase interface {
	GetAllCategories(ctx context.Context, req request.CategoryQueryParams) (*response.PaginatedResponse[response.CategoryResponse], error)
	GetAllCategoriesWithProductCount(ctx context.Context, req request.CategoryQueryParams) (*response.PaginatedResponse[response.CategoryWithProductCountResponse], error)
	GetCategoryByID(id uint) (*response.CategoryResponse, error)
	GetCategoryWithProductCount(id uint) (*response.CategoryWithProductCountResponse, error)
	CreateCategory(ctx context.Context, req request.CreateCategoryRequest) (*response.CategoryResponse, error)
	UpdateCategory(ctx context.Context, id uint, req request.UpdateCategoryRequest) (*response.CategoryResponse, error)
	DeleteCategory(id uint) error
}

type categoryUsecase struct {
	tx   TxManager
	cloudinary ImageUploader
	categoryRepo repository.CategoryRepository
	log          *zap.Logger
}

func NewCategoryUsecase(tx TxManager, cloudinary ImageUploader, categoryRepo repository.CategoryRepository, log *zap.Logger) CategoryUsecase {
	return &categoryUsecase{
		tx: tx,
		cloudinary: cloudinary,
		categoryRepo: categoryRepo,
		log:          log,
	}
}

func (uc *categoryUsecase) GetAllCategories(ctx context.Context, req request.CategoryQueryParams) (*response.PaginatedResponse[response.CategoryResponse], error) {
	if req.Page < 1 {
		return nil, utils.ErrInvalidPagination
	}
	if req.Limit < 1 || req.Limit > 100 {
		return nil, utils.ErrInvalidPagination
	}

	if req.SortBy != "" && !isValidSortField(string(req.SortBy)) {
		return nil, fmt.Errorf("%w: invalid sort field", utils.ErrInvalidSort)
	}
	if req.SortOrder != "" && req.SortOrder != "asc" && req.SortOrder != "desc" {
		return nil, fmt.Errorf("%w: invalid sort order", utils.ErrInvalidSort)
	}

	categories, total, err := uc.categoryRepo.FindAll(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w: FindAll categories failed", utils.ErrInternalServer)
	}

	categoryResponse := uc.mapToCategoryResponseList(categories)
	return response.NewPaginatedResponse(
		categoryResponse,
		req.Page,
		req.Limit,
		total,
	), nil
}

func (uc *categoryUsecase) GetAllCategoriesWithProductCount(ctx context.Context, req request.CategoryQueryParams) (*response.PaginatedResponse[response.CategoryWithProductCountResponse], error) {
	categories, total, err := uc.categoryRepo.FindAllWithProductCount(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w: FindAllWithProductCount failed", utils.ErrInternalServer)
	}

	categoriesResponse := uc.mapToCategoryWithProductCountResponse(categories)

	return response.NewPaginatedResponse(
		categoriesResponse,
		req.Page,
		req.Limit,
		total,
	), nil
}

func (uc *categoryUsecase) GetCategoryByID(id uint) (*response.CategoryResponse, error) {
	if id == 0 {
		return nil, utils.ErrInvalidCategoryID
	}

	category, err := uc.categoryRepo.FindByID(id)
	if err != nil {
		return nil, utils.ErrCategoryNotFound
	}

	return uc.mapToCategoryResponse(category), nil
}

func (uc *categoryUsecase) GetCategoryWithProductCount(id uint) (*response.CategoryWithProductCountResponse, error) {
	if id == 0 {
		return nil, utils.ErrInvalidCategoryID
	}

	category, err := uc.categoryRepo.FindByID(id)
	if err != nil {
		return nil, utils.ErrCategoryNotFound
	}

	productCount, err := uc.categoryRepo.GetProductCount(id)
	if err != nil {
		return nil, fmt.Errorf("%w: GetProductCount failed", utils.ErrInternalServer)
	}

	return &response.CategoryWithProductCountResponse{
		CategoryResponse: response.CategoryResponse{
			ID:        category.ID,
			Name:      category.Name,
			IconURL:   category.IconURL,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		},
		ProductCount: productCount,
	}, nil
}

func (uc *categoryUsecase) CreateCategory(ctx context.Context, req request.CreateCategoryRequest) (*response.CategoryResponse, error) {
	var category entity.Category
	err := uc.tx.WithinTx(ctx, func(ctx context.Context) error {
		// Validate request menggunakan DTO
		if err := req.Validate(); err != nil {
			uc.log.Warn("Category validation failed", zap.Error(err), zap.Any("request", req))
			return fmt.Errorf("%w: validation failed", utils.ErrValidationFailed)
		}

		// Upload image to cloudinary
		url, publicID, err := uc.cloudinary.Upload(ctx, req.Icon, "debian/categories")
		if err != nil {
			uc.log.Error("Failed to upload image to cloudinary", zap.Error(err))
			return err
		}

		category.Name =    req.Name
		category.IconURL = url
		category.IconPublicID = publicID

		if err := uc.categoryRepo.Create(&category); err != nil {
			if utils.IsDuplicateError(err) {
				uc.log.Warn("Category already exists", zap.String("name", req.Name))
				return utils.ErrCategoryExists
			}
			uc.log.Error("Failed to create category", zap.Error(err), zap.Any("category", category))
			return fmt.Errorf("%w: create category failed", utils.ErrDBTransaction)
		}

		uc.log.Info("Category created successfully", zap.Uint("id", category.ID), zap.String("name", category.Name))
	
	return nil
	})

	if err != nil {
		uc.log.Error("Transaction create category failed", zap.Error(err))
		return nil, err
	}
	return uc.mapToCategoryResponse(&category), nil
}

func (uc *categoryUsecase) UpdateCategory(ctx context.Context, id uint, req request.UpdateCategoryRequest) (*response.CategoryResponse, error) {
	var updatedCategory entity.Category
	err := uc.tx.WithinTx(ctx, func(ctx context.Context) error {
		// Validate request menggunakan DTO
		if err := req.Validate(); err != nil {
			uc.log.Warn("Category validation failed", zap.Error(err), zap.Any("request", req))
			return fmt.Errorf("%w: validation failed", utils.ErrValidationFailed)
		}

		category, err := uc.categoryRepo.FindByID(id)
		if err != nil {
			uc.log.Warn("Category not found", zap.Uint("id", id))
			return utils.ErrCategoryNotFound
		}

		if req.Icon != nil {
			// Get old public id to delete
			oldPublicID := category.IconPublicID

			// Upload image to cloudinary
			url, publicID, err := uc.cloudinary.Upload(ctx, req.Icon, "debian/categories")
			if err != nil {
				uc.log.Error("Failed to upload image to cloudinary", zap.Error(err))
				return err
			}

			// Delete old image
			err = uc.cloudinary.Delete(ctx, oldPublicID)
			if err != nil {
				uc.log.Error("Failed to delete old image", zap.Error(err))
				return err
			}

			category.IconURL = url
			category.IconPublicID = publicID
		}

		if req.Name != "" {
			category.Name = req.Name
		}

		if err := uc.categoryRepo.Update(category); err != nil {
			if utils.IsDuplicateError(err) {
				uc.log.Warn("Category with this name already exists", zap.String("name", req.Name))
				return utils.ErrCategoryExists
			}
			uc.log.Error("Failed to update category", zap.Error(err), zap.Uint("id", id))
			return fmt.Errorf("%w: update category failed", utils.ErrDBTransaction)
		}

		uc.log.Info("Category updated successfully", zap.Uint("id", category.ID))
	
		updatedCategory.Name = category.Name
		updatedCategory.IconURL = category.IconURL
		updatedCategory.IconPublicID = category.IconPublicID
		return nil
	})
	if err != nil {
		uc.log.Error("Transaction update category failed", zap.Error(err))
		return nil, err
	}
	return uc.mapToCategoryResponse(&updatedCategory), nil
}

func (uc *categoryUsecase) DeleteCategory(id uint) error {
	if id == 0 {
		return utils.ErrInvalidCategoryID
	}

	category, err := uc.categoryRepo.FindByID(id)
	if err != nil {
		uc.log.Warn("Category not found", zap.Uint("id", id))
		return utils.ErrCategoryNotFound
	}

	productCount, err := uc.categoryRepo.GetProductCount(id)
	if err != nil {
		uc.log.Error("Failed to get product count", zap.Error(err), zap.Uint("id", id))
		return fmt.Errorf("%w: GetProductCount failed", utils.ErrInternalServer)
	}

	if productCount > 0 {
		uc.log.Warn("Cannot delete category with products",
			zap.Uint("id", id),
			zap.Int64("product_count", productCount),
			zap.String("name", category.Name))
		return utils.ErrCategoryNotEmpty
	}

	if err := uc.categoryRepo.Delete(id); err != nil {
		uc.log.Error("Failed to delete category", zap.Error(err), zap.Uint("id", id))
		return fmt.Errorf("%w: delete category failed", utils.ErrDBTransaction)
	}

	uc.log.Info("Category deleted successfully", zap.Uint("id", id), zap.String("name", category.Name))
	return nil
}

func isValidSortField(field string) bool {
	validFields := map[string]bool{
		"id":         true,
		"name":       true,
		"created_at": true,
		"updated_at": true,
	}
	return validFields[field]
}

// Helper functions untuk mapping

func (uc *categoryUsecase) mapToCategoryResponse(category *entity.Category) *response.CategoryResponse {
	return &response.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		IconURL:   category.IconURL,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func (uc *categoryUsecase) mapToCategoryResponseList(categories []entity.Category) []response.CategoryResponse {
	responses := make([]response.CategoryResponse, len(categories))
	for i, category := range categories {
		responses[i] = *uc.mapToCategoryResponse(&category)
	}
	return responses
}

func (uc *categoryUsecase) mapToCategoryWithProductCountResponse(categories []repository.CategoryWithProductCount) []response.CategoryWithProductCountResponse {
	responses := make([]response.CategoryWithProductCountResponse, len(categories))
	for i, cat := range categories {
		responses[i] = response.CategoryWithProductCountResponse{
			CategoryResponse: response.CategoryResponse{
				ID:        cat.ID,
				Name:      cat.Name,
				IconURL:   cat.IconURL,
				CreatedAt: cat.CreatedAt,
				UpdatedAt: cat.UpdatedAt,
			},
			ProductCount: cat.ProductCount,
		}
	}
	return responses
}
