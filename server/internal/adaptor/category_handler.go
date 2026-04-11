package adaptor

import (
	"net/http"
	"strconv"

	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/usecase"
	"debian-ecommerce/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CategoryHandler struct {
	categoryUC usecase.CategoryUsecase
	log        *zap.Logger
}

func NewCategoryHandler(categoryUC usecase.CategoryUsecase, log *zap.Logger) *CategoryHandler {
	return &CategoryHandler{
		categoryUC: categoryUC,
		log:        log,
	}
}

// GetAllCategories godoc
// @Summary Get all categories
// @Description Get list of categories with pagination, sorting, and search
// @Tags Categories
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param search query string false "Search by name"
// @Param sort_by query string false "Sort field (id, name, created_at, updated_at)" default(created_at)
// @Param sort_order query string false "Sort order (asc, desc)" default(desc)
// @Param with_product_count query bool false "Include product count" default(false)
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /categories [get]
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	var req request.CategoryListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err)
		return
	}
	params := request.ToQuery(req)
	withProductCount, _ := strconv.ParseBool(c.DefaultQuery("with_product_count", "false"))

	if withProductCount {
		response, err := h.categoryUC.GetAllCategoriesWithProductCount(c, params)
		if err != nil {
			if utils.IsValidationError(err) {
				utils.BadRequest(c, "Invalid parameters", err)
			} else {
				utils.InternalServerError(c, "Failed to get categories", err)
			}
			h.log.Error("Failed to get categories with product count", zap.Error(err))
			return
		}

		utils.OK(c, "Categories retrieved successfully", response)
	} else {
		response, err := h.categoryUC.GetAllCategories(c, params)
		if err != nil {
			if utils.IsValidationError(err) {
				utils.BadRequest(c, "Invalid parameters", err)
			} else {
				utils.InternalServerError(c, "Failed to get categories", err)
			}
			h.log.Error("Failed to get categories", zap.Error(err))
			return
		}

		utils.ResponsePagination(c, http.StatusOK, "Categories retrieved successfully", response.Data, response.Pagination)
	}
}

// GetCategoryByID godoc
// @Summary Get category by ID
// @Description Get category details by ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param with_product_count query bool false "Include product count" default(false)
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid category ID", utils.ErrInvalidCategoryID)
		return
	}

	withProductCount, _ := strconv.ParseBool(c.DefaultQuery("with_product_count", "false"))

	if withProductCount {
		response, err := h.categoryUC.GetCategoryWithProductCount(uint(id))
		if err != nil {
			if err == utils.ErrCategoryNotFound {
				utils.NotFound(c, "Category not found", err)
			} else {
				utils.InternalServerError(c, "Failed to get category", err)
			}
			h.log.Error("Failed to get category with product count",
				zap.Uint("category_id", uint(id)),
				zap.Error(err))
			return
		}

		utils.OK(c, "Category retrieved successfully", response)
	} else {
		response, err := h.categoryUC.GetCategoryByID(uint(id))
		if err != nil {
			if err == utils.ErrCategoryNotFound {
				utils.NotFound(c, "Category not found", err)
			} else {
				utils.InternalServerError(c, "Failed to get category", err)
			}
			h.log.Error("Failed to get category",
				zap.Uint("category_id", uint(id)),
				zap.Error(err))
			return
		}

		utils.OK(c, "Category retrieved successfully", response)
	}
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new product category
// @Tags Categories
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param name formData string true "Category Name"
// @Param icon formData file false "Category Icon"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 409 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /admin/categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req request.CreateCategoryRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.BadRequest(c, "Invalid request body", utils.ErrInvalidJSON)
		return
	}

	response, err := h.categoryUC.CreateCategory(c, req)
	if err != nil {
		if err == utils.ErrCategoryExists {
			utils.Conflict(c, "Category already exists", err)
		} else if utils.IsValidationError(err) {
			utils.BadRequest(c, "Validation failed", err)
		} else if utils.IsDuplicateError(err) {
			utils.Conflict(c, "Category with this name already exists", err)
		} else {
			utils.InternalServerError(c, "Failed to create category", err)
		}
		h.log.Error("Failed to create category", zap.Error(err), zap.Any("request", req))
		return
	}

	h.log.Info("Category created successfully", zap.Uint("category_id", response.ID))
	utils.Created(c, "Category created successfully", response)
}

// UpdateCategory godoc
// @Summary Update category
// @Description Update existing category
// @Tags Categories
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Param name formData string true "Category Name"
// @Param icon formData file false "Category Icon"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 409 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /admin/categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid category ID", utils.ErrInvalidCategoryID)
		return
	}

	var req request.UpdateCategoryRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.BadRequest(c, "Invalid request body", utils.ErrInvalidJSON)
		return
	}

	response, err := h.categoryUC.UpdateCategory(c, uint(id), req)
	if err != nil {
		if err == utils.ErrCategoryNotFound {
			utils.NotFound(c, "Category not found", err)
		} else if err == utils.ErrCategoryExists {
			utils.Conflict(c, "Category with this name already exists", err)
		} else if utils.IsValidationError(err) {
			utils.BadRequest(c, "Validation failed", err)
		} else if utils.IsDuplicateError(err) {
			utils.Conflict(c, "Category with this name already exists", err)
		} else {
			utils.InternalServerError(c, "Failed to update category", err)
		}
		h.log.Error("Failed to update category",
			zap.Uint("category_id", uint(id)),
			zap.Error(err))
		return
	}

	h.log.Info("Category updated successfully", zap.Uint("category_id", response.ID))
	utils.OK(c, "Category updated successfully", response)
}

// DeleteCategory godoc
// @Summary Delete category
// @Description Delete category by ID
// @Tags Categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 409 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /admin/categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid category ID", utils.ErrInvalidCategoryID)
		return
	}

	if err := h.categoryUC.DeleteCategory(uint(id)); err != nil {
		if err == utils.ErrCategoryNotFound {
			utils.NotFound(c, "Category not found", err)
		} else if err == utils.ErrCategoryNotEmpty {
			utils.Conflict(c, "Cannot delete category with associated products", err)
		} else {
			utils.InternalServerError(c, "Failed to delete category", err)
		}
		h.log.Error("Failed to delete category",
			zap.Uint("category_id", uint(id)),
			zap.Error(err))
		return
	}

	h.log.Info("Category deleted successfully", zap.Uint("category_id", uint(id)))
	utils.OK(c, "Category deleted successfully", nil)
}
