package adaptor

import (
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/usecase"
	"debian-ecommerce/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CampaignHandler struct {
	service usecase.CampaignService
	log *zap.Logger
}

func NewCampaignHandler(service usecase.CampaignService, log *zap.Logger) *CampaignHandler {
	return &CampaignHandler{
		service: service,
		log: log,
	}
}

func (h *CampaignHandler) GetAll(c *gin.Context) {
	var req request.CampaignFilterRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid query params", err.Error())
		return
	}

	query := req.ToQuery()

	result, err := h.service.GetAll(c, query)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "get campaign failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "get campaign success", result)
}

// GetProducts godoc
// @Summary Get product list
// @Description Get all products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /admin/products [get]
func (h *CampaignHandler) GetByID(c *gin.Context) {
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid campaign id", err)
		return
	}
	result, err := h.service.GetWithProduct(c, uint(id))
	if err == utils.ErrProductNotFound {
		utils.ResponseFailed(c, http.StatusNotFound, "campaign not found", err)
		return
	}
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "get campaign failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "get campaign success", result)
}

func (h *CampaignHandler) Create(c *gin.Context) {	
	var req request.CampaignRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	err := h.service.Create(c, req)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "create product failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "create product success", nil)
}

func (h *CampaignHandler) Update(c *gin.Context) {
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid campaign id", err)
		return
	}
	
	var req request.CampaignRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	err = h.service.Update(c, uint(id), req)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "update campaign failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "update campaign success", nil)
}