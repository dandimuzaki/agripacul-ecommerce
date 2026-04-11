package adaptor

import (
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/usecase"
	"debian-ecommerce/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OrderHandler struct {
	service usecase.OrderService
	log *zap.Logger
}

func NewOrderHandler(service usecase.OrderService, log *zap.Logger) *OrderHandler {
	return &OrderHandler{
		service: service,
		log: log,
	}
}

func (h *OrderHandler) GetOrderHistory(c *gin.Context) {
	var req request.OrderListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err)
		return
	}

	params, err := req.ToQuery()
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err)
		return
	}

	result, err := h.service.GetOrderHistory(c, *params)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "get order history failed", err.Error())
		return
	}

	utils.ResponsePagination(c, http.StatusOK, "get order history success", result.Data, result.Pagination)
}

// GetOrders godoc
// @Summary Get order list
// @Description Get all orders
// @Tags Orders
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /admin/orders [get]
func (h *OrderHandler) GetAll(c *gin.Context) {
	var req request.OrderListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err)
		return
	}

	params, err := req.ToQuery()
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err)
		return
	}

	result, err := h.service.GetAll(c, *params)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "get orders failed", err.Error())
		return
	}

	utils.ResponsePagination(c, http.StatusOK, "get orders success", result.Data, result.Pagination)
}

func (h *OrderHandler) GetByID(c *gin.Context) {
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid order id", err.Error())
		return
	}
	result, err := h.service.GetDetails(c, uint(id))
	if err == utils.ErrOrderNotFound {
		utils.ResponseFailed(c, http.StatusNotFound, "order not found", err.Error())
		return
	}
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "get order failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "get order success", result)
}

func (h *OrderHandler) Create(c *gin.Context) {	
	var req request.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	order, err := h.service.Create(c, req)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "create order failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "create order success", order)
}

func (h *OrderHandler) Pay(c *gin.Context) {
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid order id", err.Error())
		return
	}

	err = h.service.Pay(c, uint(id))
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "pay order failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "pay order success", nil)
}

func (h *OrderHandler) Confirm(c *gin.Context) {
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid order id", err)
		return
	}

	err = h.service.Confirm(c, uint(id))
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "confirm order failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "confirm order success", nil)
}

func (h *OrderHandler) Complete(c *gin.Context) {
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid order id", err.Error())
		return
	}

	err = h.service.ManualComplete(c, uint(id))
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "complete order failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "complete order success", nil)
}

func (h *OrderHandler) Cancel(c *gin.Context) {
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid order id", err.Error())
		return
	}

	var req request.CancelOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	err = h.service.Cancel(c, uint(id), req)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "cancel order failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "cancel order success", nil)
}