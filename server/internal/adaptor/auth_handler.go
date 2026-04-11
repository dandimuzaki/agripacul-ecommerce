package adaptor

import (
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/dto/response"
	"debian-ecommerce/internal/usecase"
	"debian-ecommerce/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler struct {
	service usecase.AuthUsecase
	log     *zap.Logger
}

func NewAuthHandler(service usecase.AuthUsecase, log *zap.Logger) *AuthHandler {
	return &AuthHandler{
		service: service,
		log:     log,
	}
}

func (h *AuthHandler) RegisterCustomer(ctx *gin.Context) {
	var req request.RegisterCustomerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(ctx, http.StatusBadRequest, "invalid request", err.Error())
		return
	}

	res, err := h.service.RegisterCustomer(ctx.Request.Context(), req)
	if err != nil {
		utils.ResponseFailed(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "success", res)
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(ctx, http.StatusBadRequest, "invalid request", err.Error())
		return
	}

	res, err := h.service.Login(ctx, req)
	if err != nil {
		utils.ResponseFailed(ctx, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "success", res)
}

func (h *AuthHandler) RegisterEmployee(ctx *gin.Context) {
	var req request.RegisterEmployeeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(ctx, http.StatusBadRequest, "invalid request", err.Error())
		return
	}

	res, err := h.service.RegisterEmployee(ctx.Request.Context(), req)
	if err != nil {
		utils.ResponseFailed(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "success", res)
}

func (h *AuthHandler) CheckEmail(ctx *gin.Context) {
	var req request.CheckEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(ctx, http.StatusBadRequest, "invalid request", err.Error())
		return
	}

	available, err := h.service.IsEmailAvailable(ctx.Request.Context(), req.Email)
	if err != nil {
		utils.ResponseFailed(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "success", response.CheckEmailResponse{IsAvailable: available})
}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	// Get user ID from context (set by middleware)
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		utils.ResponseFailed(ctx, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	if err := h.service.Logout(ctx.Request.Context(), userID); err != nil {
		utils.ResponseFailed(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "logout success", nil)
}

func (h *AuthHandler) ForgotPassword(ctx *gin.Context) {
	var req request.ForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(ctx, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	err := h.service.RequestResetPassword(ctx, req)
	if err != nil {
		utils.ResponseFailed(ctx, http.StatusBadRequest, "failed to request reset password", err.Error())
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "success", nil)
}

func (h *AuthHandler) ResetPassword(ctx *gin.Context) {
	var req request.ResetPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(ctx, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	err := h.service.ResetPassword(ctx, req)
	if err != nil {
		utils.ResponseFailed(ctx, http.StatusBadRequest, "failed to reset password", err.Error())
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "success", nil)
}
