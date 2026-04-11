package adaptor

import (
	"debian-ecommerce/internal/usecase"
	"debian-ecommerce/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WishlistHandler struct {
	Service usecase.WishlistUsecase
	Log     *zap.Logger
}

func NewWishlistHandler(service usecase.WishlistUsecase, log *zap.Logger) *WishlistHandler {
	return &WishlistHandler{
		Service: service,
		Log:     log,
	}
}

type AddWishlistRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
}

func (h *WishlistHandler) AddWishlist(c *gin.Context) {
	val, exists := c.Get("user_id")
	if !exists {
		utils.ResponseFailed(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	userID := val.(uint)

	var req AddWishlistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	if err := h.Service.AddWishlist(c.Request.Context(), userID, req.ProductID); err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, "Failed to add wishlist", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Product added to wishlist", nil)
}

func (h *WishlistHandler) GetWishlist(c *gin.Context) {
	val, exists := c.Get("user_id")
	if !exists {
		utils.ResponseFailed(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	userID := val.(uint)

	wishlists, err := h.Service.GetWishlist(c.Request.Context(), userID)
	if err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, "Failed to get wishlist", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Success get wishlist", wishlists)
}

func (h *WishlistHandler) RemoveWishlist(c *gin.Context) {
	val, exists := c.Get("user_id")
	if !exists {
		utils.ResponseFailed(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	userID := val.(uint)

	productIDstr := c.Param("id")
	productID, err := strconv.ParseUint(productIDstr, 10, 64)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	if err := h.Service.RemoveWishlist(c.Request.Context(), userID, uint(productID)); err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, "Failed to remove wishlist", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Product removed from wishlist", nil)
}
