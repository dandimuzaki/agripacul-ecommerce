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

type EmployeeHandler struct {
	service usecase.EmployeeUsecase
	log     *zap.Logger
}

func NewEmployeeHandler(service usecase.EmployeeUsecase, log *zap.Logger) *EmployeeHandler {
	return &EmployeeHandler{
		service: service,
		log:     log,
	}
}

func (h *EmployeeHandler) CreateEmployeeByAdmin(ctx *gin.Context) {
	var req request.CreateEmployeeByAdminRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(ctx, http.StatusBadRequest, "invalid request", err.Error())
		return
	}

	if err := h.service.CreateEmployeeByAdmin(ctx.Request.Context(), req); err != nil {
		utils.ResponseFailed(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusCreated, "employee created successfully", nil)
}

func (h *EmployeeHandler) UpdateEmployee(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ResponseFailed(ctx, http.StatusBadRequest, "invalid id", nil)
		return
	}

	var req request.UpdateEmployeeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(ctx, http.StatusBadRequest, "invalid request", err.Error())
		return
	}

	if err := h.service.UpdateEmployee(ctx.Request.Context(), uint(id), req); err != nil {
		utils.ResponseFailed(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "employee updated successfully", nil)
}

func (h *EmployeeHandler) DeleteEmployee(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ResponseFailed(ctx, http.StatusBadRequest, "invalid id", nil)
		return
	}

	if err := h.service.DeleteEmployee(ctx.Request.Context(), uint(id)); err != nil {
		utils.ResponseFailed(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "employee deleted successfully", nil)
}

func (h *EmployeeHandler) GetEmployeeList(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")
	search := ctx.Query("search")
	sortBy := ctx.Query("sort")
	sortOrder := ctx.Query("order")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	res, err := h.service.GetEmployeeList(ctx.Request.Context(), page, limit, search, sortBy, sortOrder)
	if err != nil {
		utils.ResponseFailed(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "success", res)
}

func (h *EmployeeHandler) GetEmployeeByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ResponseFailed(ctx, http.StatusBadRequest, "invalid id", nil)
		return
	}

	res, err := h.service.GetEmployeeByID(ctx.Request.Context(), uint(id))
	if err != nil {
		utils.ResponseFailed(ctx, http.StatusNotFound, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "success", res)
}
