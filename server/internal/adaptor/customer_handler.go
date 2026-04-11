package adaptor

import (
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/usecase"
	"debian-ecommerce/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CustomerHandler struct {
	svc usecase.CustomerUsecase
	log *zap.Logger
}

func NewCustomerHandler(svc usecase.CustomerUsecase, log *zap.Logger) *CustomerHandler {
	return &CustomerHandler{
		svc: svc,
		log: log,
	}
}

func (h *CustomerHandler) UpdateProfile(c *gin.Context) {
	userId, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.Unauthorized(c, "Unauthorized", err)
		return
	}

	var req request.UpdateProfileRequest
	if err := c.ShouldBind(&req); err != nil {
		h.log.Error("Failed to retrieve request body", zap.Error(err))
		utils.BadRequest(c, err.Error(), err)
		return
	}

	if err := h.svc.UpdateProfile(c.Request.Context(), userId, req); err != nil {
		h.log.Error("Failed to update profile", zap.Error(err))
		if err.Error() == "customer profile not found" {
			utils.NotFound(c, err.Error(), err)
			return
		}
		utils.InternalServerError(c, err.Error(), err)
		return
	}

	utils.OK(c, "Profile updated successfully", nil)
}

func (h *CustomerHandler) GetProfile(c *gin.Context) {
	userId, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.Unauthorized(c, "Unauthorized", err)
		return
	}

	customer, err := h.svc.GetCustomerByUserID(c, userId)
	if err != nil {
		if err.Error() == "customer profile not found" {
			utils.NotFound(c, err.Error(), err)
			return
		}
		utils.InternalServerError(c, err.Error(), err)
		return
	}

	utils.OK(c, "Get profile successfully", customer)
}
