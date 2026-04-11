package usecase

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/data/repository"
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/dto/response"
	"debian-ecommerce/pkg/utils"
	"errors"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
)

type ProductService interface{
	GetAll(ctx context.Context, req request.ProductQueryParams) (*response.PaginatedResponse[response.ProductSummary], error)
	GetBySlug(ctx context.Context, slug string) (*response.ProductDetails, error)
	GetByID(ctx context.Context, id uint) (*response.ProductDetails, error)
	GetProductDetails(ctx context.Context, id uint) (*response.ProductDetails, error)
	Create(ctx context.Context, req request.CreateProductRequest) error
	UpdateProduct(ctx context.Context, id uint, req request.UpdateProductRequest) error
	GetSKUsByProductID(ctx context.Context, id uint) ([]response.SKUDetails, error)
	BatchUpdateSKU(ctx context.Context, productID uint, req []request.UpdateSKURequest) error
	UpdatePublish(ctx context.Context, id uint, req request.UpdatePublishRequest) error
	UploadMainImage(ctx context.Context, req request.UploadMainImageRequest) error
	UploadProductGallery(ctx context.Context, req request.UploadProductImagesRequest) error
	DeleteImage(ctx context.Context, imageID uint) error
	DeleteProduct(ctx context.Context, productID uint) error
}

type productService struct {
	tx   TxManager
	cloudinary ImageUploader
	repo *repository.Repository
	log  *zap.Logger
}

func NewProductService(tx TxManager, cloudinary ImageUploader, repo *repository.Repository, log *zap.Logger) ProductService {
	return &productService{
		tx: tx,
		cloudinary: cloudinary,
		repo: repo,
		log: log,
	}
}

func (s *productService) GetAll(ctx context.Context, req request.ProductQueryParams) (*response.PaginatedResponse[response.ProductSummary], error) {
	products, total, err := s.repo.ProductRepo.GetAll(ctx, req)
	if err != nil {
		s.log.Error("Failed to browse product", zap.Error(err))
		return nil, err
	}

	var res []response.ProductSummary
	for _, p := range products {
		summary := response.ToProductSummary(&p)
		for _, sku := range p.SKUs {
			if sku.SalePrice != nil {
				summary.SalePrice = sku.SalePrice
				summary.SalePercentage = math.Floor((sku.Price - *sku.SalePrice) / 100)
			}
		}
		res = append(res, *summary)
	}
	return response.NewPaginatedResponse(
		res,
		req.Page,
		req.Limit,
		total,
	), nil
}

func (s *productService) Create(ctx context.Context, req request.CreateProductRequest) error {
	product := req.CreateProduct()
	err := s.tx.WithinTx(ctx, func(ctx context.Context) error {
		createdProduct, err := s.repo.ProductRepo.Create(ctx, product)
		if err != nil {
			s.log.Error("Failed to create product", zap.Error(err))
			return err
		}

		for _, t := range req.VariantTypes {
			vType := t.CreateVariantType(createdProduct.ID)
			createdType, err := s.repo.VariantTypeRepo.Create(ctx, vType)
			if err != nil {
				s.log.Error("Failed to create variant type", zap.Error(err))
				return err
			}

			for _, v := range t.Values {
				vValue := t.CreateVariantValue(createdType.ID, v.Value)
				_, err := s.repo.VariantValueRepo.Create(ctx, vValue)
				if err != nil {
					s.log.Error("Failed to create variant type", zap.Error(err))
					return err
				}
			}
		}

		err = s.ensureSKUCombinations(ctx, createdProduct.ID)
		if err != nil {
			s.log.Error("Failed to create SKUs", zap.Error(err))
			return err
		}

		return nil
	})
	
	if err != nil {
		s.log.Error("Transaction create product failed", zap.Error(err))
		return err
	}

	return nil
}

func (s *productService) GetProductDetails(ctx context.Context, id uint) (*response.ProductDetails, error) {
	product, err := s.repo.ProductRepo.GetProductDetails(ctx, id)
	if err != nil {
		s.log.Error("Failed to get product",
			zap.Uint("id", id),
			zap.Error(err))
		return nil, err
	}

	return response.ToProductDetails(product), err
}

func (s *productService) GetByID(ctx context.Context, id uint) (*response.ProductDetails, error) {
	product, err := s.repo.ProductRepo.GetByID(ctx, id)
	if err != nil {
		s.log.Error("Failed to get product",
			zap.Uint("id", id),
			zap.Error(err))
		return nil, err
	}

	return response.ToProductDetails(product), err
}

func (s *productService) GetBySlug(ctx context.Context, slug string) (*response.ProductDetails, error) {
	product, err := s.repo.ProductRepo.GetBySlug(ctx, slug)
	if err != nil {
		s.log.Error("Failed to get product",
			zap.String("slug", slug),
			zap.Error(err))
		return nil, err
	}

	return response.ToProductDetails(product), err
}

func (s *productService) toSKUDetails(ctx context.Context, sku entity.SKU) (*response.SKUDetails, error) {
	var images []string
	for _, img := range sku.Images {
		images = append(images, img.ImageURL)
	}

	comb, err := s.repo.VariantValueRepo.GetVariantCombination(ctx, sku.ID)
	if err != nil {
		return nil, err
	}

	return &response.SKUDetails{
		ID: sku.ID,
		ProductID: sku.ProductID,
		SKUCode: sku.SKUCode,
		Price: sku.Price,
		SalePrice: sku.SalePrice,
		Stock: sku.Stock,
		MinStock: sku.MinStock,
		Status: sku.Status,
		Weight: sku.Weight,
		Images: images,
		Variants: comb,
	}, nil
}

func (s *productService) GetSKUsByProductID(ctx context.Context, id uint) ([]response.SKUDetails, error) {
	var SKUs []entity.SKU
	err := s.tx.WithinTx(ctx, func(ctx context.Context) error {
		err := s.ensureSKUCombinations(ctx, id)
			if err != nil {
				s.log.Error("Failed to create SKUs", zap.Error(err))
				return err
			}
		
		skus, err := s.repo.SKURepo.GetByProductID(ctx, id)
		if err != nil {
			s.log.Error("Failed to get SKUs",
				zap.Uint("id", id),
				zap.Error(err))
			return err
		}
		SKUs = skus

		return nil
	})

	var res []response.SKUDetails 
	for _, sku := range SKUs {
		skuDetails, err := s.toSKUDetails(ctx, sku)
		if err != nil {
			s.log.Error("Failed to get SKU",
				zap.Uint("id", sku.ID),
				zap.Error(err))
			return nil, err
		}
		res = append(res, *skuDetails)
	}

	if err != nil {
		s.log.Error("Get SKU by product ID failed", zap.Error(err))
		return nil, err
	}

	return res, err
}

func (s *productService) UpdateProduct(ctx context.Context, productID uint, req request.UpdateProductRequest) error {
	err := s.tx.WithinTx(ctx, func(ctx context.Context) error {
		product := req.UpdateProduct()
		err := s.repo.ProductRepo.Update(ctx, productID, product)
		if err != nil {
			s.log.Error("Failed to update product", zap.Error(err))
			return err
		}

		deletedValueIDs, err := s.handleVariantTypes(ctx, productID, req.VariantTypes)
		if err != nil {
			s.log.Error("Failed to update variant types", zap.Error(err))
			return err
		}

		err = s.archiveSKUsUsing(ctx, deletedValueIDs)
		if err != nil {
			s.log.Error("Failed to update SKUs", zap.Error(err))
			return err
		}

		err = s.ensureSKUCombinations(ctx, productID)
		if err != nil {
			s.log.Error("Failed to create SKUs", zap.Error(err))
			return err
		}

		return nil
	})  
  if err != nil {
		s.log.Error("Transaction update product failed", zap.Error(err))
		return err
	}

	return nil
}

func (s *productService) handleVariantTypes(ctx context.Context, productID uint, req []request.UpdateVariantTypeRequest) ([]uint, error) {
  var deletedValueIDs []uint
	for _, t := range req {
		switch t.Status {
		case "clean":
			if t.ID == nil {
        return nil, errors.New("variant type ID is nil for clean")
    	}
			deleted, err := s.handleVariantValues(ctx, *t.ID, t.Values)
			if err != nil {
				s.log.Error("Failed to update variant values", zap.Error(err))
				return nil, err
			}
			deletedValueIDs = append(deletedValueIDs, deleted...)

		case "created":
			vType := entity.VariantType{
				ProductID: productID,
				Name: t.Name,
			}
			createdType, err := s.repo.VariantTypeRepo.Create(ctx, &vType)
			if err != nil {
				s.log.Error("Failed to create variant type", zap.Error(err))
				return nil, err
			}
			_, err = s.handleVariantValues(ctx, createdType.ID, t.Values)
			if err != nil {
				s.log.Error("Failed to create variant value", zap.Error(err))
				return nil, err
			}

		case "updated":
			vType := entity.VariantType{
				ProductID: productID,
				Name: t.Name,
			}
			err := s.repo.VariantTypeRepo.Update(ctx, *t.ID, &vType)
			if err != nil {
				s.log.Error("Failed to update variant types", zap.Error(err))
				return nil, err
			}
			deleted, err := s.handleVariantValues(ctx, *t.ID, t.Values)
			if err != nil {
				s.log.Error("Failed to update variant values", zap.Error(err))
				return nil, err
			}
			deletedValueIDs = append(deletedValueIDs, deleted...)

		case "deleted":
			valueIDs, err := s.repo.VariantValueRepo.GetIDsByTypeID(ctx, *t.ID)
			if err != nil {
				s.log.Error("Failed to get variant value ids", zap.Error(err))
				return nil, err
			}
			for _, valueID := range valueIDs {
				err = s.repo.VariantValueRepo.Delete(ctx, valueID)
				if err != nil {
					s.log.Error("Failed to delete variant value", zap.Error(err))
					return nil, err
				}
			}
			deletedValueIDs = append(deletedValueIDs, valueIDs...)
			err = s.repo.VariantTypeRepo.Delete(ctx, *t.ID)
			if err != nil {
				s.log.Error("Failed to delete variant type", zap.Error(err))
				return nil, err
			}

		default:
			deleted, err := s.handleVariantValues(ctx, *t.ID, t.Values)
			if err != nil {
				s.log.Error("Failed to update variant values", zap.Error(err))
				return nil, err
			}
			deletedValueIDs = append(deletedValueIDs, deleted...)

		}
	}
	return deletedValueIDs, nil
}

func (s *productService) handleVariantValues(ctx context.Context, typeID uint, values []request.UpdateVariantValueRequest) ([]uint, error) {
	var deletedValueIDs []uint
	for _, v := range values {
		switch v.Status {
		case "created":
			vValue := entity.VariantValue{
				VariantTypeID: typeID,
				Value: v.Value,
			}
			_, err := s.repo.VariantValueRepo.Create(ctx, &vValue)
			if err != nil {
				s.log.Error("Failed to create value", zap.Error(err))
				return nil, err
			}

		case "updated":
			vValue := entity.VariantValue{
				VariantTypeID: typeID,
				Value: v.Value,
			}
			err := s.repo.VariantValueRepo.Update(ctx, *v.ID, &vValue)
			if err != nil {
				s.log.Error("Failed to update value", zap.Error(err))
				return nil, err
			}

		case "deleted":
			err := s.repo.VariantValueRepo.Delete(ctx, *v.ID)
			if err != nil {
				s.log.Error("Failed to delete value", zap.Error(err))
				return nil, err
			}
			deletedValueIDs = append(deletedValueIDs, *v.ID)
		}
	}

	return deletedValueIDs, nil
}

func (s *productService) archiveSKUsUsing(ctx context.Context, valueIDs []uint) error {
	if len(valueIDs) == 0 {
			return nil
	}
		skuIDs, err := s.repo.SKURepo.FindSKUIDsByVariantValues(ctx, valueIDs)
		if err != nil {
			s.log.Error("Failed to find sku ids", zap.Error(err))
			return err
		}

		err = s.repo.SKURepo.ArchiveBatch(ctx, skuIDs)
		if err != nil {
			s.log.Error("Failed to batch deactivate SKUs", zap.Error(err))
			return err
		}

		err = s.repo.SKURepo.DeleteSKUVariantValueBySKUIDs(ctx, skuIDs)
		if err != nil {
			s.log.Error("Failed to batch delete SKU - variant value", zap.Error(err))
			return err
		}
	
	return nil
}

func (s *productService) ensureSKUCombinations(ctx context.Context, productID uint) error {
	vTypes, err := s.repo.VariantTypeRepo.GetVariantTypesByProductID(ctx, productID)
	if err != nil {
		s.log.Error("Error get variant types", zap.Error(err))
		return err
	}

	// Case 1: Product with no variants
	if len(vTypes) == 0 {
		exists, err := s.repo.SKURepo.GetByProductID(ctx, productID)
		if err != nil {
			s.log.Error("Error get SKUs", zap.Error(err))
			return err
		}

		if len(exists) > 0 {
			// already has default SKU, do nothing
			return nil
		}

		skuCode, err := utils.GenerateRandomString(10)
		if err != nil {
			s.log.Error("Error generate SKU Code", zap.Error(err))
			return err
		}

		sku := entity.SKU{
			ProductID: productID,
			SKUCode:   skuCode,
			Price:     0,
			Stock:     0,
			Status:    entity.SKUStatusInactive,
		}

		_, err = s.repo.SKURepo.Create(ctx, &sku)
		if err != nil {
			s.log.Error("Error create SKU", zap.Error(err))
			return err
		}
	}

	var valueIDs [][]uint
	for _, typeID := range vTypes {
		vIDs, err := s.repo.VariantValueRepo.GetIDsByTypeID(ctx, typeID)
		if err != nil {
			s.log.Error("Error get variant values", zap.Error(err))
			return err
		}
		if len(vIDs) == 0 {
			return errors.New("variant type has no values")
		}
		valueIDs = append(valueIDs, vIDs)
	}

	existingSKUs, err := s.repo.SKURepo.GetIDsByProductID(ctx, productID)
	if err != nil {
		s.log.Error("Error get SKUs", zap.Error(err))
		return err
	}
	
	skuMap := map[string]uint{}
	if len(existingSKUs) != 0 {
		for _, skuID := range existingSKUs {
			valueIDs, err := s.repo.SKURepo.GetVariantValueIDs(ctx, skuID)
			if err != nil {
				s.log.Error("Error get variant values", zap.Error(err))
				return err
			}
			key := s.normalizeVariantValues(valueIDs)
			skuMap[key] = skuID
		}
	}

	combinations := s.CartesianProduct(valueIDs)
	for _, com := range combinations {
		key := s.normalizeVariantValues(com)
		if _, ok := skuMap[key]; !ok {
			tempCode, err := s.generateTemporarySKUCode(ctx, productID, com)
			if err != nil {
				s.log.Error("Error generate temporary SKU code", zap.Error(err))
				return err
			}
			sku := entity.SKU{
				ProductID: productID,
				SKUCode: *tempCode,
				Price: 0,
				Stock: 0,
				Status: entity.SKUStatusInactive,
			}
			createdSKU, err := s.repo.SKURepo.Create(ctx, &sku)
			if err != nil {
				s.log.Error("Error create SKU", zap.Error(err))
				return err
			}
			for _, valueID := range com {
				skuValue := entity.SKUVariantValue{
					SKUID: createdSKU.ID,
					VariantValueID: valueID,
				}
				_, err := s.repo.SKURepo.CreateSKUVariantValue(ctx, &skuValue)
				if err != nil {
					s.log.Error("Error create SKU - variant value", zap.Error(err))
					return err
				}
			}
		}
	}
	return nil
}

func (s *productService) CartesianProduct(pools [][]uint) [][]uint {
	if len(pools) == 0 {
		return nil
	}

	// Seed with first pool
	result := make([][]uint, 0)
	for _, v := range pools[0] {
		result = append(result, []uint{v})
	}

	// Combine with remaining pools
	for i := 1; i < len(pools); i++ {
		var next [][]uint
		for _, r := range result {
			for _, v := range pools[i] {
				combination := append(append([]uint{}, r...), v)
				next = append(next, combination)
			}
		}
		result = next
	}

	return result
}

func (s *productService) normalizeVariantValues(ids []uint) string {
	slices.Sort(ids)
	var valueIDs []string
	for _, id := range ids {
		idStr := strconv.FormatUint(uint64(id), 10)
		valueIDs = append(valueIDs, idStr)
	}
	return strings.Join(valueIDs, "-")
}

func (s *productService) generateTemporarySKUCode(ctx context.Context, productID uint, valueIDs []uint) (*string, error) {
	product, err := s.repo.ProductRepo.GetByID(ctx, productID)
	if err != nil {
		s.log.Error("Error get product", zap.Error(err))
		return nil, err
	}

	values, err := s.repo.VariantValueRepo.GetValuesByValueIDs(ctx, valueIDs)
	if err != nil {
		s.log.Error("Error get variant values", zap.Error(err))
		return nil, err
	}

	skuCode := fmt.Sprintf(
		"%s-%s-%d",
		product.Slug,
		strings.Join(values, "-"),
		time.Now().UnixNano(),
	)

	return &skuCode, nil
}

func (s *productService) BatchUpdateSKU(ctx context.Context, productID uint, req []request.UpdateSKURequest) error {
	err := s.tx.WithinTx(ctx, func(ctx context.Context) error {
		for _, r := range req {
			sku, err := s.repo.SKURepo.GetByID(ctx, r.ID)
			if err != nil {
				return err
			}

			if sku.Status == "archived" {
				return errors.New("archived SKU cannot be modified")
			}

			if r.Status == "active" && r.Stock == 0 {
				return errors.New("cannot activate SKU with zero stock")
			}

			if r.Status == "active" && r.Price <= 0 {
				return errors.New("active SKU must have price")
			}

			update := entity.SKU{
				SKUCode: r.SKUCode,
				Price:   r.Price,
				Stock:   r.Stock,
				Status:  entity.SKUStatus(r.Status),
			}

			if err := s.repo.SKURepo.Update(ctx, r.ID, &update); err != nil {
				return err
			}
		}

		// 🔥 Recalculate product price ONCE
		if err := s.repo.ProductRepo.RecalculatePrice(ctx, productID); err != nil {
			s.log.Error("Failed to update product", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		s.log.Error("Error transaction update SKUs", zap.Error(err))
		return err
	}

	return nil
}

func (s *productService) UpdatePublish(ctx context.Context, id uint, req request.UpdatePublishRequest) error {
	data := map[string]interface{}{
		"is_published": req.IsPublished,
	}
	err := s.repo.ProductRepo.UpdatePublish(ctx, id, data)
	if err != nil {
		s.log.Error("Failed to update publish")
		return err
	}
	return nil
}

func (s *productService) UploadMainImage(ctx context.Context, req request.UploadMainImageRequest) error {
	err := s.tx.WithinTx(ctx, func(ctx context.Context) error {
		product, err := s.repo.ProductRepo.GetByID(ctx, req.ProductID)
		if err != nil {
			s.log.Error("Failed to get product", zap.Error(err))
			return err
		}

		// Get old public id to delete
		oldPublicID := product.MainImagePublicID

		// Upload image to cloudinary
		url, publicID, err := s.cloudinary.Upload(ctx, req.Image, "debian/products/main")
		if err != nil {
			s.log.Error("Failed to upload image to cloudinary", zap.Error(err))
			return err
		}

		product.MainImageURL = url
		product.MainImagePublicID = publicID

		// Update product main image
		err = s.repo.ProductRepo.Update(ctx, req.ProductID, product)
		if err != nil {
			s.log.Error("Failed to update product", zap.Error(err))
			return err
		}

		err = s.cloudinary.Delete(ctx, oldPublicID)
		if err != nil {
			s.log.Error("Failed to delete old image", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		s.log.Error("Transaction upload main image failed", zap.Error(err))
		return err
	}

	return nil
}

func (s *productService) UploadProductGallery(ctx context.Context, req request.UploadProductImagesRequest) error {
	err := s.tx.WithinTx(ctx, func(ctx context.Context) error {
		// Initialize public IDs
		var uploaded []string
		for _, file := range req.Images {
			url, publicID, err := s.cloudinary.Upload(ctx, file, "debian/products/gallery")
			if err != nil {
				s.log.Error("Failed to upload image", zap.Error(err))
				s.rollbackImages(ctx, uploaded)
				return err
			}

			image := entity.Image{
				ProductID: req.ProductID,
				ImageURL:  url,
				PublicID:  publicID,
			}

			if err := s.repo.ImageRepo.Create(ctx, &image); err != nil {
				s.log.Error("Failed to create image", zap.Error(err))
				uploaded = append(uploaded, publicID)
				s.rollbackImages(ctx, uploaded)
				return err
			}

			uploaded = append(uploaded, publicID)
		}

		return nil
	})

	if err != nil {
		s.log.Error("Transaction upload main image failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *productService) rollbackImages(ctx context.Context, publicIDs []string) {
	for _, id := range publicIDs {
		_ = s.cloudinary.Delete(ctx, id)
	}
}

func (s *productService) DeleteImage(ctx context.Context, imageID uint) error {
	img, err := s.repo.ImageRepo.GetByID(ctx, imageID)
	if err != nil {
		s.log.Error("Failed to get image", zap.Error(err))
		return err
	}
	err = s.cloudinary.Delete(ctx, img.PublicID)
	if err != nil {
		s.log.Error("Failed to delete image public URL", zap.Error(err))
		return err
	}
	err = s.repo.ImageRepo.Delete(ctx, imageID)
	if err != nil {
		s.log.Error("Failed to delete image", zap.Error(err))
		return err
	}
	return nil
}

func (s *productService) 	DeleteProduct(ctx context.Context, productID uint) error {
	p, err := s.repo.ProductRepo.GetByID(ctx, productID)
	if err != nil {
		s.log.Error("Failed to get product", zap.Error(err))
		return err
	}

	var skuIDs []uint
	for _, sku := range p.SKUs {
		skuIDs = append(skuIDs, sku.ID)
	}
	err = s.repo.SKURepo.ArchiveBatch(ctx, skuIDs)
	if err != nil {
		s.log.Error("Failed to archive SKU", zap.Error(err))
		return err
	}

	for _, t := range p.VariantTypes {
		err = s.repo.VariantTypeRepo.Delete(ctx, t.ID)
			if err != nil {
			s.log.Error("Failed to delete variant type", zap.Error(err))
			return err
		}
		for _, val := range t.Values {
			err = s.repo.VariantValueRepo.Delete(ctx, val.ID)
			if err != nil {
				s.log.Error("Failed to delete variant value", zap.Error(err))
				return err
			}
		}
	}

	err = s.repo.ProductRepo.Delete(ctx, productID)
	if err != nil {
		s.log.Error("Failed to delete product", zap.Error(err))
		return err
	}
	return nil
}