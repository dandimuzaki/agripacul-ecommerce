package adaptor

import (
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/usecase"
	"debian-ecommerce/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type InventoryHandler struct {
	service usecase.InventoryService
	log  *zap.Logger
}

func NewInventoryHandler(service usecase.InventoryService, log *zap.Logger) *InventoryHandler {
	return &InventoryHandler{
		service: service,
		log:  log,
	}
}

func (h *InventoryHandler) GetInventoryLogs(c *gin.Context) {
	var req request.InventoryListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err)
		return
	}

	params := req.ToQuery()
	result, err := h.service.GetInventoryLogs(c, params)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadGateway, "get inventory logs failed", nil)
		return
	}

	utils.ResponsePagination(c, http.StatusOK, "get inventory logs success", result.Data, result.Pagination)
}
func (h *InventoryHandler) GetInventoryLogsBySKUID(c *gin.Context) {
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid product id", err)
		return
	}

	var req request.InventoryListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err)
		return
	}

	params := req.ToQuery()
	result, err := h.service.GetInventoryLogsBySKUID(c, id, params)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadGateway, "get inventory logs failed", nil)
		return
	}

	utils.ResponsePagination(c, http.StatusOK, "get inventory logs success", result.Data, result.Pagination)
}

func (h *InventoryHandler) CreateInventoryLog(c *gin.Context) {
	var req request.CreateInventoryLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid request", err)
		return
	}

	// Validation
	messages, err := utils.ValidateErrors(req)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, err.Error(), messages)
		return
	}

	res, err := h.service.CreateInventoryLog(c.Request.Context(), req)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "create inventory log failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusCreated, "create inventory log success", res)
}

func (h *InventoryHandler) GetInventory(c *gin.Context) {
	var req request.InventoryListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err)
		return
	}

	params := req.ToQuery()
	result, err := h.service.GetInventory(c, params)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadGateway, "get inventory failed", nil)
		return
	}

	utils.ResponsePagination(c, http.StatusOK, "get inventory success", result.Data, result.Pagination)
}

func (h *InventoryHandler) GetInventoryBySKUID(c *gin.Context) {
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid sku id", err)
		return
	}

	result, err := h.service.GetInventoryBySKUID(c, id)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadGateway, "get inventory by sku id failed", nil)
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "get inventory by sku id success", result)
}