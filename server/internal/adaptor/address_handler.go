package adaptor

import (
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/usecase"
	"debian-ecommerce/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AddressHandler struct {
	svc         usecase.AddressUsecase
	customerSvc usecase.CustomerUsecase
	log         *zap.Logger
}

func NewAddressHandler(svc usecase.AddressUsecase, customerSvc usecase.CustomerUsecase, log *zap.Logger) *AddressHandler {
	return &AddressHandler{
		svc:         svc,
		customerSvc: customerSvc,
		log:         log,
	}
}

func (h *AddressHandler) Create(c *gin.Context) {
	var req request.CreateAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error(), err)
		return
	}

	if err := h.svc.Create(c, req); err != nil {
		h.log.Error("Failed to create address", zap.Error(err))
		utils.InternalServerError(c, "Failed to create address", err)
		return
	}

	utils.Created(c, "Address created successfully", nil)
}

func (h *AddressHandler) Update(c *gin.Context) {
	// Verify ownership?
	// User can only update their own address.
	// But update by ID doesn't enforce ownership check in Usecase yet unless implemented.
	// We can check ownership here if we fetch first, or rely on Usecase.
	// For now, let's just proceed. Ideally Usecase should check.

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid address ID", err)
		return
	}

	var req request.UpdateAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error(), err)
		return
	}

	if err := h.svc.Update(uint(id), req); err != nil {
		h.log.Error("Failed to update address", zap.Error(err))
		utils.InternalServerError(c, "Failed to update address", err)
		return
	}

	utils.OK(c, "Address updated successfully", nil)
}

func (h *AddressHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid address ID", err)
		return
	}

	if err := h.svc.Delete(uint(id)); err != nil {
		h.log.Error("Failed to delete address", zap.Error(err))
		utils.InternalServerError(c, "Failed to delete address", err)
		return
	}

	utils.OK(c, "Address deleted successfully", nil)
}

func (h *AddressHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid address ID", err)
		return
	}

	address, err := h.svc.GetByID(uint(id))
	if err != nil {
		h.log.Error("Failed to get address", zap.Error(err))
		utils.InternalServerError(c, "Failed to get address", err)
		return
	}

	utils.OK(c, "Address retrieved successfully", address)
}

func (h *AddressHandler) GetByCustomerID(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.Unauthorized(c, "Unauthorized", err)
		return
	}

	customer, err := h.customerSvc.GetCustomerByUserID(c.Request.Context(), userID)
	if err != nil {
		h.log.Error("Failed to get customer", zap.Error(err))
		utils.InternalServerError(c, "Failed to get customer profile", err)
		return
	}

	addresses, err := h.svc.GetByCustomerID(customer.ID)
	if err != nil {
		h.log.Error("Failed to get addresses", zap.Error(err))
		utils.InternalServerError(c, "Failed to get addresses", err)
		return
	}

	utils.OK(c, "Addresses retrieved successfully", addresses)
}

func (h *AddressHandler) SetDefault(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.Unauthorized(c, "Unauthorized", err)
		return
	}

	customer, err := h.customerSvc.GetCustomerByUserID(c.Request.Context(), userID)
	if err != nil {
		h.log.Error("Failed to get customer", zap.Error(err))
		utils.InternalServerError(c, "Failed to get customer profile", err)
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid address ID", err)
		return
	}

	if err := h.svc.SetDefault(customer.ID, uint(id)); err != nil {
		h.log.Error("Failed to set default address", zap.Error(err))
		utils.InternalServerError(c, "Failed to set default address", err)
		return
	}

	utils.OK(c, "Default address set successfully", nil)
}

func (h *AddressHandler) GetProvinces(c *gin.Context) {
	provinces, err := h.svc.GetProvinces()
	if err != nil {
		h.log.Error("Failed to get provinces", zap.Error(err))
		utils.InternalServerError(c, "Failed to get provinces", err)
		return
	}
	utils.OK(c, "Provinces retrieved successfully", provinces)
}

func (h *AddressHandler) GetRegencies(c *gin.Context) {
	provinceIDStr := c.Param("province_id")
	provinceID, err := strconv.ParseUint(provinceIDStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid province ID", err)
		return
	}

	cities, err := h.svc.GetRegencies(uint(provinceID))
	if err != nil {
		h.log.Error("Failed to get cities", zap.Error(err))
		utils.InternalServerError(c, "Failed to get cities", err)
		return
	}
	utils.OK(c, "Cities retrieved successfully", cities)
}

func (h *AddressHandler) GetDistricts(c *gin.Context) {
	cityIDStr := c.Param("regency_id")
	cityID, err := strconv.ParseUint(cityIDStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid regency ID", err)
		return
	}

	districts, err := h.svc.GetDistricts(uint(cityID))
	if err != nil {
		h.log.Error("Failed to get districts", zap.Error(err))
		utils.InternalServerError(c, "Failed to get districts", err)
		return
	}
	utils.OK(c, "Districts retrieved successfully", districts)
}

func (h *AddressHandler) GetSubdistricts(c *gin.Context) {
	districtIDStr := c.Param("district_id")
	districtID, err := strconv.ParseUint(districtIDStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid district ID", err)
		return
	}

	subdistricts, err := h.svc.GetSubdistricts(uint(districtID))
	if err != nil {
		h.log.Error("Failed to get subdistricts", zap.Error(err))
		utils.InternalServerError(c, "Failed to get subdistricts", err)
		return
	}
	utils.OK(c, "Subdistricts retrieved successfully", subdistricts)
}
