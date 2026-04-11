package adaptor

import (
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/usecase"
	"debian-ecommerce/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductHandler struct {
	service usecase.ProductService
	log *zap.Logger
}

func NewProductHandler(service usecase.ProductService, log *zap.Logger) *ProductHandler {
	return &ProductHandler{
		service: service,
		log: log,
	}
}

// BrowseProducts godoc
// @Summary Get product list
// @Description Get list of products with pagination, sorting, and search
// @Tags Products
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param search query string false "Search by name"
// @Param category_id query int false "Filter by category"
// @Param min_price query int false "Filter with minimum price"
// @Param max_price query int false "Filter with maximum price"
// @Param rating query int false "Filter by rating"
// @Param sort_by query string false "Sort field (id, name, created_at, price, average_rating, sold_count)" default(created_at)
// @Param sort_order query string false "Sort order (asc, desc)" default(desc)
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /products [get]
func (h *ProductHandler) BrowseProducts(c *gin.Context) {
	var req request.ProductListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err.Error())
		return
	}

	params := req.ToQuery()
	params.InStockOnly = true
	params.IsPublishedOnly = true

	result, err := h.service.GetAll(c, params)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "get product failed", err.Error())
		return
	}

	utils.ResponsePagination(c, http.StatusOK, "get product success", result.Data, result.Pagination)
}

// GetAllProducts godoc
// @Summary Get all products by admin
// @Description Get list of products with pagination, sorting, and search
// @Tags Products
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param search query string false "Search by name"
// @Param category_id query int false "Filter by category"
// @Param min_price query int false "Filter with minimum price"
// @Param max_price query int false "Filter with maximum price"
// @Param rating query int false "Filter by rating"
// @Param sort_by query string false "Sort field (id, name, created_at, price, average_rating, sold_count)" default(created_at)
// @Param sort_order query string false "Sort order (asc, desc)" default(desc)
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /admin/products [get]
func (h *ProductHandler) ProductListDashboard(c *gin.Context) {
	var req request.ProductListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err.Error())
		return
	}

	params := req.ToQuery()

	result, err := h.service.GetAll(c, params)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "get product failed", err.Error())
		return
	}

	utils.ResponsePagination(c, http.StatusOK, "get product success", result.Data, result.Pagination)
}

// GetProductDetails godoc
// @Summary Get product details
// @Description Get product details by customer
// @Tags Products
// @Accept json
// @Produce json
// @Params id path int true "Product ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /products/{id} [get]
func (h *ProductHandler) GetProductDetails(c *gin.Context) {
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid product id", err)
		return
	}
	result, err := h.service.GetProductDetails(c, uint(id))
	if err == utils.ErrProductNotFound {
		utils.ResponseFailed(c, http.StatusNotFound, "product not found", err)
		return
	}
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "get product failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "get product success", result)
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Get product details by admin
// @Tags Products
// @Accept json
// @Produce json
// @Params id path int true "Product ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /admin/products/{id} [get]
func (h *ProductHandler) GetByID(c *gin.Context) {
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid product id", err)
		return
	}
	result, err := h.service.GetByID(c, uint(id))
	if err == utils.ErrProductNotFound {
		utils.ResponseFailed(c, http.StatusNotFound, "product not found", err)
		return
	}
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "get product failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "get product success", result)
}

// GetProductBySlug godoc
// @Summary Get product by slug
// @Description Get product details by slug
// @Tags Products
// @Accept json
// @Produce json
// @Params slug path string true "Product Slug"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /products/details/{slug} [get]
func (h *ProductHandler) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")
	result, err := h.service.GetBySlug(c, slug)
	if err == utils.ErrProductNotFound {
		utils.ResponseFailed(c, http.StatusNotFound, "product not found", err)
		return
	}
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "get product failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "get product success", result)
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body request.CreateProductRequest true "Product data"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 409 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /admin/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {	
	var req request.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err.Error())
		return
	}

	err := h.service.Create(c, req)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "create product failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusCreated, "create product success", nil)
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update existing product
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param request body request.UpdateProductRequest true "Product data"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 409 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid product id", err)
		return
	}
	
	var req request.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err.Error())
		return
	}

	err = h.service.UpdateProduct(c, uint(id), req)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "update product failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "update product success", nil)
}

// UpdatePublish godoc
// @Summary Publish or unpublish product
// @Description Update product to be published or unpublished
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param request body request.UpdatePublishRequest true "Product is published or unpublished"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 409 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /admin/products/{id}/publish [put]
func (h *ProductHandler) UpdatePublish(c *gin.Context) {
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid product id", err)
		return
	}
	
	var req request.UpdatePublishRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	err = h.service.UpdatePublish(c, uint(id), req)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "update product failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "update product success", nil)
}

// GetSKUsByProductID godoc
// @Summary Get product SKUs
// @Description Get stock keeping units of product by product ID
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 409 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /admin/products/{id}/sku [get]
func (h *ProductHandler) GetSKUsByProductID(c *gin.Context) {
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid product id", err)
		return
	}
	result, err := h.service.GetSKUsByProductID(c, uint(id))
	if err == utils.ErrProductNotFound {
		utils.ResponseFailed(c, http.StatusNotFound, "sku not found", err)
		return
	}
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "get SKUs failed", err.Error())
		h.log.Error("Failed to get SKU", zap.Error(err))
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "get SKUs success", result)
}

// BatchUpdateSKU godoc
// @Summary Update product SKUs
// @Description Update stock keeping units of product by product ID
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param request body request.UpdateSKURequest true "Product SKUs data"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 409 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /admin/products/{id}/sku [put]
func (h *ProductHandler) BatchUpdateSKU(c *gin.Context) {	
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid product id", err)
		return
	}

	var req []request.UpdateSKURequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid params", err.Error())
		return
	}

	err = h.service.BatchUpdateSKU(c, id, req)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "update SKUs failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "update SKUs success", nil)
}

// UploadMainImage godoc
// @Summary Upload product main image
// @Description Upload main image of product
// @Tags Products
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param image formData file false "Product Main Image"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 409 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /admin/products/{id} [post]
func (h *ProductHandler) UploadMainImage(c *gin.Context) {	
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid product id", err)
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		h.log.Error("Failed to retrieve image file", zap.Error(err))
		utils.ResponseFailed(c, http.StatusBadRequest, "image is required", err)
		return
	}

	req := request.UploadMainImageRequest{
		ProductID: id,
		Image:     file,
	}

	err = h.service.UploadMainImage(c, req)
	if err != nil {
		h.log.Error("Failed to upload main image", zap.Error(err))
		utils.ResponseFailed(c, http.StatusBadRequest, "upload main image failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "upload main image success", nil)
}

// UploadProductGallery godoc
// @Summary Upload product gallery
// @Description Upload collection of product images
// @Tags Products
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param images formData file false "Product Images"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 409 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /admin/products/{id}/gallery [post]
func (h *ProductHandler) UploadProductGallery(c *gin.Context) {	
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid product id", err)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		h.log.Error("Failed to retrieve form", zap.Error(err))
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid multipart form", err)
		return
	}

	files := form.File["images"]
	if len(files) == 0 {
		h.log.Error("Failed to retrieve image file", zap.Error(err))
		utils.ResponseFailed(c, http.StatusBadRequest, "images are required", nil)
		return
	}

	req := request.UploadProductImagesRequest{
		ProductID: id,
		Images:    files,
	}

	err = h.service.UploadProductGallery(c, req)
	if err != nil {
		h.log.Error("Failed to upload image gallery", zap.Error(err))
		utils.ResponseFailed(c, http.StatusBadRequest, "upload images failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "upload images success", nil)
}

// DeleteImage godoc
// @Summary Delete product image
// @Description Delete product image by ID
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param imageID path int true "Image ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 409 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /admin/products/{id}/gallery/{imageID} [delete]
func (h *ProductHandler) DeleteImage(c *gin.Context) {	
	id, err := utils.GetUintParam(c, "imageID")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid image id", err)
		return
	}

	err = h.service.DeleteImage(c, id)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "delete image failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "delete image success", nil)
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Delete product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 409 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /admin/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {	
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid product id", err)
		return
	}

	err = h.service.DeleteProduct(c, id)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "delete product failed", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "delete product success", nil)
}