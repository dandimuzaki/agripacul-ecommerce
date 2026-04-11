package repository

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/dto/response"
	infra "debian-ecommerce/internal/infra/transaction"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InventoryRepository interface {
	GetInventoryLogs(ctx context.Context, f request.InventoryQueryParams) ([]entity.InventoryLog, int64, error)
	CreateInventoryLog(ctx context.Context, inventory *entity.InventoryLog) (*entity.InventoryLog, error)
	GetInventory(ctx context.Context, f request.InventoryQueryParams) ([]response.InventoryResponse, int64, error)
	BatchCreateInventoryLogs(ctx context.Context, inventoryLogs []entity.InventoryLog) error
	GetInventoryLogsBySKUID(ctx context.Context, skuID uint, f request.InventoryQueryParams) ([]entity.InventoryLog, int64, error)
	GetInventoryBySKUID(ctx context.Context, skuID uint) (*response.InventoryResponse, error)
}

type inventoryRepository struct {
	db *gorm.DB
	log *zap.Logger
}

func NewInventoryRepo(db *gorm.DB, log *zap.Logger) InventoryRepository {
	return &inventoryRepository{
		db: db,
		log: log,
	}
}

func (r *inventoryRepository) GetInventoryLogs(ctx context.Context, f request.InventoryQueryParams) ([]entity.InventoryLog, int64, error) {
	db := infra.GetDB(ctx, r.db)
	var logs []entity.InventoryLog
	var total int64
	query := db.Model(&entity.InventoryLog{}).Order("created_at desc")
	
	// Get total inventory logs
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Limit(f.Limit).Offset(f.Offset).Find(&logs).Error
	if err != nil {
		r.log.Error("Error query get inventory logs", zap.Error(err))
		return nil, 0, err
	}

	return logs, total, nil
}

func (r *inventoryRepository) GetInventoryLogsBySKUID(ctx context.Context, skuID uint, f request.InventoryQueryParams) ([]entity.InventoryLog, int64, error) {
	db := infra.GetDB(ctx, r.db)
	var logs []entity.InventoryLog
	var total int64
	query := db.Model(&entity.InventoryLog{}).Order("created_at desc").Where("sku_id", skuID)
	
	// Get total inventory logs
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Limit(f.Limit).Offset(f.Offset).Find(&logs).Error
	if err != nil {
		r.log.Error("Error query get inventory logs", zap.Error(err))
		return nil, 0, err
	}

	return logs, total, nil
}

func (r *inventoryRepository) CreateInventoryLog(ctx context.Context, inventory *entity.InventoryLog) (*entity.InventoryLog, error) {
	db := infra.GetDB(ctx, r.db)
	err := db.Create(&inventory).Error
	if err != nil {
		r.log.Error("Error query create inventory", zap.Error(err))
		return nil, err
	}
	return inventory, err
}

func (r *inventoryRepository) GetInventory(
	ctx context.Context,
	f request.InventoryQueryParams,
) ([]response.InventoryResponse, int64, error) {

	db := infra.GetDB(ctx, r.db)

	var results []response.InventoryResponse

	// ===============================
	// BASE QUERY (for reuse)
	// ===============================
	baseQuery := db.Model(&entity.SKU{})

	// Filter by SKU status
	if f.Status != "" {
		baseQuery = baseQuery.Where("skus.status = ?", f.Status)
	}

	// Filter by stock availability
	switch f.Stock {
	case "out":
		baseQuery = baseQuery.Where("skus.stock <= 0")
	case "low":
		baseQuery = baseQuery.Where("skus.stock < skus.min_stock")
	case "in":
		baseQuery = baseQuery.Where("skus.stock > 0 AND skus.stock >= skus.min_stock")
	}

	// ===============================
	// COUNT QUERY (with filters)
	// ===============================
	var total int64
	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// ===============================
	// MAIN QUERY
	// ===============================
	query := baseQuery.
		Select(`
			skus.id AS id,
			p.name AS product,
			skus.sku_code,
			skus.stock,
			skus.min_stock,

			STRING_AGG(DISTINCT vv.value, ' | ') AS variant_label,

			CASE
				WHEN skus.stock <= 0 THEN 'out of stock'
				WHEN skus.stock < skus.min_stock THEN 'low stock'
				ELSE 'in stock'
			END AS availability,

			skus.status
		`).
		Joins("LEFT JOIN products p ON skus.product_id = p.id").
		Joins("LEFT JOIN sku_variant_values svv ON svv.sku_id = skus.id").
		Joins("LEFT JOIN variant_values vv ON vv.id = svv.variant_value_id").
		Group(`
			skus.id,
			p.name,
			skus.sku_code,
			skus.stock,
			skus.min_stock,
			skus.status
		`)

	// Filter by keyword
	if f.Search != "" {
		searchPattern := "%" + strings.ToLower(f.Search) + "%"
		query = query.Where("LOWER(p.name) LIKE ? OR LOWER(skus.sku_code) LIKE ?", searchPattern)
	}

	// ===============================
	// SORTING
	// ===============================
	sortDesc := f.SortOrder == request.SortDesc

	switch f.SortBy {
	case string(request.SortInventoryByID):
		query = query.Order(clause.OrderByColumn{
			Column: clause.Column{Name: "skus.id"},
			Desc:   sortDesc,
		})
	case string(request.SortInventoryByCreatedAt):
		query = query.Order(clause.OrderByColumn{
			Column: clause.Column{Name: "skus.created_at"},
			Desc:   sortDesc,
		})
	case string(request.SortInventoryByUpdatedAt):
		query = query.Order(clause.OrderByColumn{
			Column: clause.Column{Name: "skus.updated_at"},
			Desc:   sortDesc,
		})
	case string(request.SortInventoryByName):
		query = query.Order(clause.OrderByColumn{
			Column: clause.Column{Name: "p.name"},
			Desc:   sortDesc,
		})
	case string(request.SortInventoryByStock):
		query = query.Order(clause.OrderByColumn{
			Column: clause.Column{Name: "skus.stock"},
			Desc:   sortDesc,
		})
	default:
		query = query.Order("skus.created_at DESC")
	}

	// ===============================
	// PAGINATION + EXECUTION
	// ===============================
	err := query.
		Limit(f.Limit).
		Offset(f.Offset).
		Scan(&results).Error

	if err != nil {
		r.log.Error("Error query get inventory", zap.Error(err))
		return nil, 0, err
	}

	return results, total, nil
}

func (r *inventoryRepository) BatchCreateInventoryLogs(ctx context.Context, inventoryLogs []entity.InventoryLog) error {
	db := infra.GetDB(ctx, r.db)
	err := db.Create(&inventoryLogs).Error
	if err != nil {
		r.log.Error("Error query create inventory logs", zap.Error(err))
		return err
	}
	return nil
}

func (r *inventoryRepository) GetInventoryBySKUID(
	ctx context.Context,
	skuID uint,
) (*response.InventoryResponse, error) {

	db := infra.GetDB(ctx, r.db)

	var inventory response.InventoryResponse

	query := db.Model(&entity.SKU{}).
		Select(`
			skus.id AS id,
			p.name AS product,
			skus.sku_code,
			skus.stock,
			skus.min_stock,

			STRING_AGG(vv.value, ' | ') AS variant_label,

			CASE
				WHEN skus.stock <= 0 THEN 'out of stock'
				WHEN skus.stock < skus.min_stock THEN 'low stock'
				ELSE 'in stock'
			END AS availability,

			skus.status
		`).
		Joins("LEFT JOIN products p ON skus.product_id = p.id").
		Joins("LEFT JOIN sku_variant_values svv ON svv.sku_id = skus.id").
		Joins("LEFT JOIN variant_values vv ON vv.id = svv.variant_value_id").
		Where("skus.id = ?", skuID).
		Group(`
			skus.id,
			p.name,
			skus.sku_code,
			skus.stock,
			skus.min_stock,
			skus.status
		`)

	err := query.
		Find(&inventory).Error

	if err != nil {
		r.log.Error("Error query get inventory", zap.Error(err))
		return nil, err
	}

	return &inventory, nil
}