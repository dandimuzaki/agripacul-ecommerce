package adaptor

import (
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/usecase"
	"debian-ecommerce/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BannerHandler struct {
	Service usecase.BannerUsecase
	Log     *zap.Logger
}

func NewBannerHandler(service usecase.BannerUsecase, log *zap.Logger) *BannerHandler {
	return &BannerHandler{
		Service: service,
		Log:     log,
	}
}

func (h *BannerHandler) CreateBanner(c *gin.Context) {
	var req request.BannerRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid request", err.Error())
		return
	}

	if err := h.Service.CreateBanner(c, &req); err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(c, http.StatusCreated, "banner created successfully", nil)
}

func (h *BannerHandler) GetBannerList(c *gin.Context) {
	banners, err := h.Service.GetBannerList(c.Request.Context())
	if err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, "get banners failed", err.Error())
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "get banners success", banners)
}

func (h *BannerHandler) GetBannerByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid id", nil)
		return
	}

	banner, err := h.Service.GetBannerByID(c.Request.Context(), uint(id))
	if err != nil {
		utils.ResponseFailed(c, http.StatusNotFound, err.Error(), nil)
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "success", banner)
}

func (h *BannerHandler) UpdateBanner(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid id", nil)
		return
	}

	var req request.BannerRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid request", err.Error())
		return
	}

	if err := h.Service.UpdateBanner(c.Request.Context(), uint(id), &req); err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "banner updated successfully", nil)
}

func (h *BannerHandler) DeleteBanner(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid id", nil)
		return
	}

	if err := h.Service.DeleteBanner(c.Request.Context(), uint(id)); err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "banner deleted successfully", nil)
}
