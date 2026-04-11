package adaptor

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/usecase"
	"debian-ecommerce/pkg/utils"
)

type PromotionHandler struct {
	promotionUsecase usecase.PromotionUsecase
}

func NewPromotionHandler(promotionUsecase usecase.PromotionUsecase) *PromotionHandler {
	return &PromotionHandler{
		promotionUsecase: promotionUsecase,
	}
}

// CreatePromotion godoc
// @Summary Create a new promotion
// @Description Create a new promotion with products
// @Tags Promotions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param promotion body request.PromotionRequest true "Promotion data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /promotions [post]
func (h *PromotionHandler) CreatePromotion(c *gin.Context) {
	var req request.PromotionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err.Error())
		return
	}

	// Validasi sederhana
	if req.Name == "" && req.Type == "" && req.VoucherCode == nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid request", "invalid request")
		return
	}

	promotion, err := h.promotionUsecase.CreatePromotion(c, req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errorMsg := err.Error()

		// Tentukan status code berdasarkan error
		if errorMsg == "promotion not found" {
			statusCode = http.StatusNotFound
		} else if errorMsg == "voucher code already exists" {
			statusCode = http.StatusConflict
		}

		utils.ResponseFailed(c, statusCode, "create promotion failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "create promotion success", promotion)
}

// GetPromotionByID godoc
// @Summary Get promotion by ID
// @Description Get promotion details by ID
// @Tags Promotions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Promotion ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /promotions/{id} [get]
func (h *PromotionHandler) GetPromotionByID(c *gin.Context) {
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid product id", err)
		return
	}

	promotion, err := h.promotionUsecase.GetPromotionByID(c, uint(id))
	if err != nil {
		statusCode := http.StatusInternalServerError
		errorMsg := err.Error()

		// Jika promotion tidak ditemukan
		if errorMsg == "promotion not found" {
			statusCode = http.StatusNotFound
		}

		utils.ResponseFailed(c, statusCode, "get promotion failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "get promotion success", promotion)
}

// GetPromotionList godoc
// @Summary Get promotion list with filters
// @Description Get paginated list of promotions with optional filters
// @Tags Promotions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param name query string false "Promotion name"
// @Param type query string false "Promotion type" Enums(direct discount, voucher code)
// @Param is_published query boolean false "Is published"
// @Param start_date_from query string false "Start date from (ISO 8601)"
// @Param start_date_to query string false "Start date to (ISO 8601)"
// @Param end_date_from query string false "End date from (ISO 8601)"
// @Param end_date_to query string false "End date to (ISO 8601)"
// @Param voucher_code query string false "Voucher code"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param sort_by query string false "Sort by field" default(created_at)
// @Param sort_order query string false "Sort order" default(desc)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /promotions [get]
func (h *PromotionHandler) GetPromotionList(c *gin.Context) {
	// Bind filter
	var filter request.PromotionFilterRequest
	if err := c.ShouldBindQuery(&filter); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err.Error())
		return
	}

	query, err := filter.ToQuery()
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err.Error())
		return
	}

	// Panggil usecase dengan filter dan pagination request
	// GetPromotionList mengembalikan 2 nilai: (*response.PromotionListResponse, error)
	result, err := h.promotionUsecase.GetPromotionList(c, *query)
	if err := c.ShouldBindQuery(&filter); err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, "failed to get promotion", err.Error())
		return
	}

	// Return response dengan data promotions dan pagination meta
	// result berisi Promotions dan Pagination dari response.PromotionListResponse
	utils.ResponseSuccess(c, http.StatusOK, "get promotion success", result)
}

// UpdatePromotion godoc
// @Summary Update promotion
// @Description Update promotion details
// @Tags Promotions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Promotion ID"
// @Param promotion body request.PromotionUpdateRequest true "Promotion update data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /promotions/{id} [put]
func (h *PromotionHandler) UpdatePromotion(c *gin.Context) {
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid promotion id", err)
		return
	}

	var req request.PromotionUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "get promotion failed", err.Error())
		return
	}

	// Validasi sederhana
	if req.Name == "" && req.Type == "" && req.VoucherCode == nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "get promotion failed", "invalid request")
		return
	}

	promotion, err := h.promotionUsecase.UpdatePromotion(c, uint(id), req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errorMsg := err.Error()

		// Tentukan status code berdasarkan error
		if errorMsg == "promotion not found" {
			statusCode = http.StatusNotFound
		} else if errorMsg == "voucher code already exists" {
			statusCode = http.StatusConflict
		}

		utils.ResponseFailed(c, statusCode, "update promotion failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "update promotion success", promotion)
}

// DeletePromotion godoc
// @Summary Delete promotion
// @Description Delete promotion by ID
// @Tags Promotions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Promotion ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /promotions/{id} [delete]
func (h *PromotionHandler) DeletePromotion(c *gin.Context) {
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid promotion id", err)
		return
	}

	err = h.promotionUsecase.DeletePromotion(c, uint(id))
	if err != nil {
		statusCode := http.StatusInternalServerError
		errorMsg := err.Error()

		// Tentukan status code berdasarkan error
		if errorMsg == "promotion not found" {
			statusCode = http.StatusNotFound
		} else if errorMsg == "voucher code already exists" {
			statusCode = http.StatusConflict
		}

		utils.ResponseFailed(c, statusCode, "delete promotion failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "delete promotion success", nil)
}
