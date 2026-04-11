package adaptor

import (
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/usecase"
	"debian-ecommerce/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ReportHandler struct {
	service usecase.ReportService
	log     *zap.Logger
}

func NewReportHandler(service usecase.ReportService, log *zap.Logger) *ReportHandler {
	return &ReportHandler{
		service: service,
		log:     log,
	}
}

func (h *ReportHandler) GetSales(c *gin.Context) {
	var req request.ReportRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err)
		return
	}

	query, err := req.ToQuery()
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err)
		return
	}

	result, err := h.service.GetSales(c, *query)
	if err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, "get revenue report failed", nil)
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "get revenue report success", result)
}

func (h *ReportHandler) GetRevenue(c *gin.Context) {
	var req request.ReportRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err)
		return
	}

	query, err := req.ToQuery()
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err)
		return
	}

	result, err := h.service.GetRevenue(c, *query)
	if err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, "get revenue report failed", nil)
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "get revenue report success", result)
}

func (h *ReportHandler) GetProductPerformance(c *gin.Context) {
	var req request.ReportRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err)
		return
	}

	query, err := req.ToQuery()
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err)
		return
	}

	result, err := h.service.GetProductPerformance(c, *query)
	if err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, "get product performance failed", nil)
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "get product performance success", result)
}

func (h *ReportHandler) GetLoyalCustomer(c *gin.Context) {
	result, err := h.service.GetCustomerReport(c)
	if err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, "get loyal customer failed", nil)
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "get loyal customer success", result)
}

func (h *ReportHandler) GetCustomerSummary(c *gin.Context) {
	result, err := h.service.GetCustomerSummary(c)
	if err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, "get customer summary failed", nil)
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "get customer summary success", result)
}