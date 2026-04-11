package adaptor

import (
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/usecase"
	"debian-ecommerce/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CheckoutHandler struct {
	service usecase.CheckoutService
	log  *zap.Logger
}

func NewCheckoutHandler(service usecase.CheckoutService, log *zap.Logger) *CheckoutHandler {
	return &CheckoutHandler{
		service: service,
		log:  log,
	}
}

func (h *CheckoutHandler) PreviewCheckout(c *gin.Context) {
	var req request.PreviewCheckout
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err)
		return
	}

	result, err := h.service.GetPreviewCheckout(c, req)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadGateway, "get checkout failed", nil)
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "get checkout success", result)
}

func (h *CheckoutHandler) GetShippingOptions(c *gin.Context) {
	var req request.ShippingOptionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err)
		return
	}

	result, err := h.service.FetchShippingOptions(c, req)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadGateway, "get shipping options failed", nil)
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "get shipping options success", result)
}

func (h *CheckoutHandler) GetValidPromotions(c *gin.Context) {
	var req request.PromotionFilterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err)
		return
	}

	result, err := h.service.GetValidPromotions(c, req)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadGateway, "get valid promotions failed", nil)
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "get valid promotions success", result)
}