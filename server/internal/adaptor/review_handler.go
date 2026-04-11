package adaptor

import (
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/usecase"
	"debian-ecommerce/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ReviewHandler struct {
	service usecase.ReviewUsecase
	log     *zap.Logger
}

func NewReviewHandler(service usecase.ReviewUsecase, log *zap.Logger) *ReviewHandler {
	return &ReviewHandler{
		service: service,
		log:     log,
	}
}

func (h *ReviewHandler) CreateReview(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		utils.Forbidden(ctx, "Unauthorized", nil)
		return
	}

	// Assuming userID in context is uint. If it's float64 (from JSON unmarshal sometimes), cast appropriately.
	// But middleware usually sets it as specific type.
	// Based on token.go: Claims.UserID is uint.
	uid, ok := userID.(uint)
	if !ok {
		// Handle float64 case if jwt library decoded numbers as float
		if fId, ok := userID.(float64); ok {
			uid = uint(fId)
		} else {
			utils.InternalServerError(ctx, "invalid user id type", nil)
			return
		}
	}

	var req request.CreateReviewRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, "Invalid request", err)
		return
	}

	res, err := h.service.CreateReview(ctx.Request.Context(), uid, req)
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	utils.Created(ctx, "Review created successfully", res)
}

func (h *ReviewHandler) BatchCreateReview(ctx *gin.Context) {
	var req request.BatchCreateReview
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, "Invalid request", err)
		return
	}

	err := h.service.BatchCreateReview(ctx, req)
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	utils.Created(ctx, "Review created successfully", nil)
}

func (h *ReviewHandler) GetAllReviews(ctx *gin.Context) {
	var req request.ReviewsQuery
	if err := ctx.BindQuery(&req); err != nil {
		utils.ValidationError(ctx, "Invalid request", err)
		return
	}
	f := request.ToReviewsFilter(&req)

	res, err := h.service.GetAllReviews(ctx, *f)
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	utils.OK(ctx, "Reviews retrieved successfully", res)
}

func (h *ReviewHandler) GetReviewsByProduct(ctx *gin.Context) {
	productIDStr := ctx.Query("product_id")
	if productIDStr == "" {
		productIDStr = ctx.Param("product_id")
	}

	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		utils.BadRequest(ctx, "Invalid product ID", err)
		return
	}

	var req request.ReviewsQuery
	if err := ctx.BindQuery(&req); err != nil {
		utils.ValidationError(ctx, "Invalid request", err)
		return
	}
	f := request.ToReviewsFilter(&req)

	res, err := h.service.GetReviewsByProduct(ctx, uint(productID), *f)
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	utils.OK(ctx, "Reviews retrieved successfully", res)
}

func (h *ReviewHandler) GetReviewDetails(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.BadRequest(ctx, "Invalid review ID", err)
		return
	}

	res, err := h.service.GetReviewDetails(ctx.Request.Context(), uint(id))
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	utils.OK(ctx, "Review details retrieved successfully", res)
}

func (h *ReviewHandler) UpdateReview(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		utils.Forbidden(ctx, "Unauthorized", nil)
		return
	}

	uid, ok := userID.(uint)
	if !ok {
		if fId, ok := userID.(float64); ok {
			uid = uint(fId)
		} else {
			utils.InternalServerError(ctx, "invalid user id type", nil)
			return
		}
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.BadRequest(ctx, "Invalid review ID", err)
		return
	}

	var req request.UpdateReviewRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, "Invalid request", err)
		return
	}

	res, err := h.service.UpdateReviewComment(ctx.Request.Context(), uid, uint(id), req)
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	utils.OK(ctx, "Review updated successfully", res)
}

func (h *ReviewHandler) DeleteReview(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		utils.Forbidden(ctx, "Unauthorized", nil)
		return
	}

	uid, ok := userID.(uint)
	if !ok {
		if fId, ok := userID.(float64); ok {
			uid = uint(fId)
		} else {
			utils.InternalServerError(ctx, "invalid user id type", nil)
			return
		}
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.BadRequest(ctx, "Invalid review ID", err)
		return
	}

	err = h.service.DeleteReview(ctx.Request.Context(), uid, uint(id))
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	utils.OK(ctx, "Review deleted successfully", nil)
}

func (h *ReviewHandler) GetReviewStats(ctx *gin.Context) {
	productIDStr := ctx.Query("product_id")
	if productIDStr == "" {
		productIDStr = ctx.Param("id")
	}

	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		utils.BadRequest(ctx, "Invalid product ID", err)
		return
	}

	res, err := h.service.GetReviewStats(ctx.Request.Context(), uint(productID))
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	utils.OK(ctx, "Review stats retrieved successfully", res)
}

func (h *ReviewHandler) handleError(ctx *gin.Context, err error) {
	if utils.IsNotFoundError(err) {
		utils.NotFound(ctx, "Not found", err)
	} else if utils.IsAuthError(err) {
		utils.Forbidden(ctx, "Forbidden", err)
	} else if utils.IsValidationError(err) {
		utils.ValidationError(ctx, "Validation error", err)
	} else if utils.IsDuplicateError(err) {
		utils.Conflict(ctx, "Conflict", err)
	} else {
		// Check for AppError
		if appErr, ok := err.(*utils.AppError); ok {
			switch appErr.Code {
			case utils.ErrCodeForbidden:
				utils.Forbidden(ctx, appErr.Message, appErr)
			case utils.ErrCodeConflict:
				utils.Conflict(ctx, appErr.Message, appErr)
			case utils.ErrCodeBadRequest:
				utils.BadRequest(ctx, appErr.Message, appErr)
			default:
				utils.InternalServerError(ctx, appErr.Message, appErr)
			}
			return
		}
		utils.InternalServerError(ctx, "Internal server error", err)
	}
}
