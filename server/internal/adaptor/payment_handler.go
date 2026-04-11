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

type PaymentMethodHandler struct {
	paymentMethodUsecase *usecase.PaymentMethodUsecase
	log                  *zap.Logger
}

func NewPaymentMethodHandler(
	paymentMethodUsecase *usecase.PaymentMethodUsecase,
	log *zap.Logger,
) *PaymentMethodHandler {
	return &PaymentMethodHandler{
		paymentMethodUsecase: paymentMethodUsecase,
		log:                  log,
	}
}

// CreatePaymentMethod godoc
// @Summary Create payment method
// @Description Create a new payment method
// @Tags Payment Method
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body request.CreatePaymentMethodRequest true "Payment method data"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 403 {object} utils.Response
// @Failure 409 {object} utils.Response
// @Failure 422 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /payment-methods [post]
func (h *PaymentMethodHandler) CreatePaymentMethod(c *gin.Context) {
	var req request.CreatePaymentMethodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("failed to bind request")
		utils.ResponseFailed(c, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	// Validate request with custom validator
	if fieldErrors, err := utils.ValidateErrors(req); err != nil {
		h.log.Error("validation failed")
		c.JSON(http.StatusUnprocessableEntity, utils.Response{
			Success:  false,
			Message: "Validation failed",
			Fields:  fieldErrors,
		})
		return
	}

	// Panggil CreatePaymentMethod, bukan Create
	result, err := h.paymentMethodUsecase.CreatePaymentMethod(c.Request.Context(), &req)
	if err != nil {
		h.log.Error("failed to create payment method")

		switch {
		case err == utils.ErrNameRequired:
			utils.ResponseFailed(c, http.StatusBadRequest, err.Error(), nil)
		case err == utils.ErrDuplicatePaymentMethod:
			utils.ResponseFailed(c, http.StatusConflict, err.Error(), nil)
		default:
			utils.ResponseFailed(c, http.StatusInternalServerError, "Failed to create payment method", nil)
		}
		return
	}

	utils.ResponseSuccess(c, http.StatusCreated, "Payment method created successfully", result)
}

// GetPaymentMethods godoc
// @Summary Get all payment methods
// @Description Get list of payment methods with pagination and filters
// @Tags Payment Method
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param search query string false "Search by name"
// @Param is_active query bool false "Filter by active status"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 403 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /payment-methods [get]
func (h *PaymentMethodHandler) GetPaymentMethods(c *gin.Context) {
	var req request.GetPaymentMethodsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error("failed to bind query parameters")
		utils.ResponseFailed(c, http.StatusBadRequest, "Invalid query parameters", err.Error())
		return
	}

	// Validate pagination
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}

	// Panggil GetPaymentMethods, bukan GetAll
	result, err := h.paymentMethodUsecase.GetPaymentMethods(c.Request.Context(), &req)
	if err != nil {
		h.log.Error("failed to get payment methods")

		switch {
		case utils.IsNotFoundError(err):
			utils.ResponseFailed(c, http.StatusNotFound, "No payment methods found", nil)
		default:
			utils.ResponseFailed(c, http.StatusInternalServerError, "Failed to get payment methods", nil)
		}
		return
	}

	// Create pagination response
	pagination := map[string]interface{}{
		"page":        result.Page,
		"limit":       result.Limit,
		"total":       result.Total,
		"total_pages": result.TotalPages,
	}

	utils.ResponsePagination(c, http.StatusOK, "Payment methods retrieved successfully", result.PaymentMethods, pagination)
}

// GetPaymentMethodByID godoc
// @Summary Get payment method by ID
// @Description Get payment method details by ID
// @Tags Payment Method
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Payment Method ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 403 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /payment-methods/{id} [get]
func (h *PaymentMethodHandler) GetPaymentMethodByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		h.log.Error("invalid payment method ID")
		utils.ResponseFailed(c, http.StatusBadRequest, "Invalid payment method ID", utils.ErrInvalidUserID.Error())
		return
	}

	// Panggil GetPaymentMethodByID, bukan GetByID
	result, err := h.paymentMethodUsecase.GetPaymentMethodByID(c.Request.Context(), uint(id))
	if err != nil {
		h.log.Error("failed to get payment method")

		switch {
		case err == utils.ErrPaymentMethodNotFound, utils.IsNotFoundError(err):
			utils.ResponseFailed(c, http.StatusNotFound, "Payment method not found", nil)
		default:
			utils.ResponseFailed(c, http.StatusInternalServerError, "Failed to get payment method", nil)
		}
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Payment method retrieved successfully", result)
}

// UpdatePaymentMethod godoc
// @Summary Update payment method
// @Description Update payment method by ID
// @Tags Payment Method
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Payment Method ID"
// @Param request body request.UpdatePaymentMethodRequest true "Payment method update data"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 403 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 409 {object} utils.Response
// @Failure 422 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/payment-methods/{id} [put]
func (h *PaymentMethodHandler) UpdatePaymentMethod(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		h.log.Error("invalid payment method ID")
		utils.ResponseFailed(c, http.StatusBadRequest, "Invalid payment method ID", utils.ErrInvalidUserID.Error())
		return
	}

	var req request.UpdatePaymentMethodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("failed to bind request")
		utils.ResponseFailed(c, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	// Validate request with custom validator
	if fieldErrors, err := utils.ValidateErrors(req); err != nil {
		h.log.Error("validation failed")
		c.JSON(http.StatusUnprocessableEntity, utils.Response{
			Success:  false,
			Message: "Validation failed",
			Fields:  fieldErrors,
		})
		return
	}

	// Panggil UpdatePaymentMethod, bukan Update
	result, err := h.paymentMethodUsecase.UpdatePaymentMethod(c.Request.Context(), uint(id), &req)
	if err != nil {
		h.log.Error("failed to update payment method")

		switch {
		case err == utils.ErrPaymentMethodNotFound, utils.IsNotFoundError(err):
			utils.ResponseFailed(c, http.StatusNotFound, "Payment method not found", nil)
		case err == utils.ErrPaymentMethodInactive:
			utils.ResponseFailed(c, http.StatusBadRequest, err.Error(), nil)
		case err == utils.ErrDuplicatePaymentMethod:
			utils.ResponseFailed(c, http.StatusConflict, err.Error(), nil)
		default:
			utils.ResponseFailed(c, http.StatusInternalServerError, "Failed to update payment method", nil)
		}
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Payment method updated successfully", result)
}

// DeletePaymentMethod godoc
// @Summary Delete payment method
// @Description Delete payment method by ID (soft delete)
// @Tags Payment Method
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Payment Method ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 403 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /payment-methods/{id} [delete]
func (h *PaymentMethodHandler) DeletePaymentMethod(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		h.log.Error("invalid payment method ID")
		utils.ResponseFailed(c, http.StatusBadRequest, "Invalid payment method ID", utils.ErrInvalidUserID.Error())
		return
	}

	// Panggil DeletePaymentMethod
	err = h.paymentMethodUsecase.DeletePaymentMethod(c.Request.Context(), uint(id))
	if err != nil {
		h.log.Error("failed to delete payment method")

		switch {
		case err == utils.ErrPaymentMethodNotFound, utils.IsNotFoundError(err):
			utils.ResponseFailed(c, http.StatusNotFound, "Payment method not found", nil)
		default:
			utils.ResponseFailed(c, http.StatusInternalServerError, "Failed to delete payment method", nil)
		}
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Payment method deleted successfully", nil)
}

// TogglePaymentMethodStatus godoc
// @Summary Toggle payment method active status
// @Description Activate or deactivate payment method
// @Tags Payment Method
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Payment Method ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 403 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /payment-methods/{id}/toggle [patch]
func (h *PaymentMethodHandler) TogglePaymentMethodStatus(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		h.log.Error("invalid payment method ID")
		utils.ResponseFailed(c, http.StatusBadRequest, "Invalid payment method ID", utils.ErrInvalidUserID.Error())
		return
	}

	// Panggil TogglePaymentMethodStatus
	result, err := h.paymentMethodUsecase.TogglePaymentMethodStatus(c.Request.Context(), uint(id))
	if err != nil {
		h.log.Error("failed to toggle payment method status")

		switch {
		case err == utils.ErrPaymentMethodNotFound, utils.IsNotFoundError(err):
			utils.ResponseFailed(c, http.StatusNotFound, "Payment method not found", nil)
		default:
			utils.ResponseFailed(c, http.StatusInternalServerError, "Failed to toggle payment method status", nil)
		}
		return
	}

	status := "activated"
	if !result.IsActive {
		status = "deactivated"
	}
	utils.ResponseSuccess(c, http.StatusOK, "Payment method "+status+" successfully", result)
}
