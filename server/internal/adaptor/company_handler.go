package adaptor

import (
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/usecase"
	"debian-ecommerce/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CompanyHandler struct {
	svc usecase.CompanyUsecase
	log *zap.Logger
}

func NewCompanyHandler(svc usecase.CompanyUsecase, log *zap.Logger) *CompanyHandler {
	return &CompanyHandler{
		svc: svc,
		log: log,
	}
}

func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	var req request.CreateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error(), err)
		return
	}

	if err := h.svc.CreateCompany(c.Request.Context(), req); err != nil {
		h.log.Error("Failed to create company", zap.Error(err))
		utils.InternalServerError(c, "Failed to create company", err)
		return
	}

	utils.Created(c, "Company created successfully", nil)
}

func (h *CompanyHandler) GetCompany(c *gin.Context) {
	company, err := h.svc.GetCompany(c.Request.Context())
	if err != nil {
		h.log.Error("Failed to get company", zap.Error(err))
		utils.InternalServerError(c, "Failed to get company", err)
		return
	}

	utils.OK(c, "Company retrieved successfully", company)
}

func (h *CompanyHandler) UpdateCompany(c *gin.Context) {
	var req request.UpdateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error(), err)
		return
	}

	if err := h.svc.UpdateCompany(c.Request.Context(), req); err != nil {
		h.log.Error("Failed to update company", zap.Error(err))
		utils.InternalServerError(c, "Failed to update company", err)
		return
	}

	utils.OK(c, "Company updated successfully", nil)
}

func (h *CompanyHandler) CreateAddress(c *gin.Context) {
	var req request.CreateCompanyAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error(), err)
		return
	}

	if err := h.svc.CreateAddress(c.Request.Context(), req); err != nil {
		h.log.Error("Failed to create company address", zap.Error(err))
		utils.InternalServerError(c, "Failed to create company address", err)
		return
	}

	utils.Created(c, "Company address created successfully", nil)
}

func (h *CompanyHandler) UpdateAddress(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid address ID", err)
		return
	}

	var req request.UpdateCompanyAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error(), err)
		return
	}

	if err := h.svc.UpdateAddress(c.Request.Context(), uint(id), req); err != nil {
		h.log.Error("Failed to update company address", zap.Error(err))
		utils.InternalServerError(c, "Failed to update company address", err)
		return
	}

	utils.OK(c, "Company address updated successfully", nil)
}

func (h *CompanyHandler) GetAddressByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid address ID", err)
		return
	}

	address, err := h.svc.GetAddressByID(c.Request.Context(), uint(id))
	if err != nil {
		h.log.Error("Failed to get company address", zap.Error(err))
		utils.InternalServerError(c, "Failed to get company address", err)
		return
	}

	utils.OK(c, "Company address retrieved successfully", address)
}

func (h *CompanyHandler) GetShippingOrigin(c *gin.Context) {
	// Assuming there is only one company or we pass company ID.
	// For now, let's assume company ID is 1 or we fetch the first one.
	// But `GetShippingOriginAddress` takes companyID.
	// In the request body for CreateAddress we require CompanyID.
	// Here maybe query param? Or assume 1?
	// Let's assume the user passes company_id as query param or we create a dedicated endpoint for specific company?
	// Or maybe the system only supports 1 company.
	// Let's try to get company ID from query, default to 1 if not provided? Or fetch the first company ID first.
	// Or better: pass company_id.
	companyIDStr := c.Query("company_id")
	var companyID uint64 = 1
	var err error
	if companyIDStr != "" {
		companyID, err = strconv.ParseUint(companyIDStr, 10, 32)
		if err != nil {
			utils.BadRequest(c, "Invalid company ID", err)
			return
		}
	} else {
		// Try to fetch first company? Or just fail?
		// User request "Get company address that is shipping origin".
		// I'll assume ID 1 for simplicity if not provided, or error.
		// Let's error if not provided? Or default to 1 (common for single tenant).
		// I'll default to 1.
	}

	address, err := h.svc.GetShippingOriginAddress(c.Request.Context(), uint(companyID))
	if err != nil {
		h.log.Error("Failed to get shipping origin", zap.Error(err))
		utils.InternalServerError(c, "Failed to get shipping origin", err)
		return
	}

	utils.OK(c, "Shipping origin retrieved successfully", address)
}

func (h *CompanyHandler) SendMessage(c *gin.Context) {
	var req request.SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error(), err)
		return
	}

	if err := h.svc.SendMessage(c, req); err != nil {
		h.log.Error("Failed to send message", zap.Error(err))
		utils.InternalServerError(c, "Failed to send message", err)
		return
	}

	utils.Created(c, "Message sent successfully", nil)
}
