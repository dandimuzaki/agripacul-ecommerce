package adaptor

import (
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/usecase"
	"debian-ecommerce/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CartHandler struct {
	svc usecase.CartUsecase
	log *zap.Logger
}

func NewCartHandler(svc usecase.CartUsecase, log *zap.Logger) *CartHandler {
	return &CartHandler{
		svc: svc,
		log: log,
	}
}

func (h *CartHandler) GetCart(c *gin.Context) {
	cart, err := h.svc.GetCart(c)
	if err != nil {
		h.log.Error("Failed to get cart", zap.Error(err))
		utils.InternalServerError(c, "Failed to get cart", err)
		return
	}

	utils.OK(c, "Cart retrieved successfully", cart)
}

func (h *CartHandler) AddItem(c *gin.Context) {
	var req request.AddCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error(), err)
		return
	}

	if err := h.svc.AddItem(c, req); err != nil {
		h.log.Error("Failed to add item to cart", zap.Error(err))
		utils.InternalServerError(c, "Failed to add item to cart", err)
		return
	}

	utils.Created(c, "Item added to cart successfully", nil)
}

func (h *CartHandler) UpdateItem(c *gin.Context) {
	itemIDStr := c.Param("id")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid item ID", err)
		return
	}

	var req request.UpdateCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error(), err)
		return
	}

	if err := h.svc.UpdateItem(c, uint(itemID), req); err != nil {
		h.log.Error("Failed to update cart item", zap.Error(err))
		if utils.IsNotFoundError(err) {
			utils.NotFound(c, "Item not found", err)
			return
		}
		utils.InternalServerError(c, "Failed to update item", err)
		return
	}

	utils.OK(c, "Cart item updated successfully", nil)
}

func (h *CartHandler) BatchSelectItem(c *gin.Context) {
	var req request.BatchSelectCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error(), err)
		return
	}

	if err := h.svc.BatchSelectItem(c, req); err != nil {
		h.log.Error("Failed to batch select cart items", zap.Error(err))
		if utils.IsNotFoundError(err) {
			utils.NotFound(c, "Item not found", err)
			return
		}
		utils.InternalServerError(c, "Failed to batch select cart items", err)
		return
	}

	if req.IsSelected {
		utils.OK(c, "Cart items selected successfully", nil)
		return
	}

	utils.OK(c, "Cart items unselected successfully", nil)
}


func (h *CartHandler) RemoveItem(c *gin.Context) {
	itemIDStr := c.Param("id")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid item ID", err)
		return
	}

	if err := h.svc.RemoveItem(c, uint(itemID)); err != nil {
		h.log.Error("Failed to remove cart item", zap.Error(err))
		if utils.IsNotFoundError(err) {
			utils.NotFound(c, "Item not found", err)
			return
		}
		utils.InternalServerError(c, "Failed to remove item", err)
		return
	}

	utils.OK(c, "Item removed from cart successfully", nil)
}

func (h *CartHandler) ClearCart(c *gin.Context) {
	if err := h.svc.ClearCart(c); err != nil {
		h.log.Error("Failed to clear cart", zap.Error(err))
		utils.InternalServerError(c, "Failed to clear cart", err)
		return
	}

	utils.OK(c, "Cart cleared successfully", nil)
}