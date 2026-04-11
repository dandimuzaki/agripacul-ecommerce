package usecase

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/data/repository"
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/dto/response"

	"go.uber.org/zap"
)

type InventoryService interface {
	GetInventoryLogs(ctx context.Context, req request.InventoryQueryParams) (*response.PaginatedResponse[response.InventoryLogResponse], error)
	CreateInventoryLog(ctx context.Context, req request.CreateInventoryLogRequest) (*response.InventoryLogResponse, error)
	GetInventory(ctx context.Context, req request.InventoryQueryParams) (*response.PaginatedResponse[response.InventoryResponse], error)
	GetInventoryLogsBySKUID(ctx context.Context, skuID uint, req request.InventoryQueryParams) (*response.PaginatedResponse[response.InventoryLogResponse], error)
	GetInventoryBySKUID(ctx context.Context, skuID uint) (*response.InventoryResponse, error)
}

type inventoryService struct {
	tx TxManager
	repo *repository.Repository
	log *zap.Logger
}

func NewInventoryService(tx TxManager, repo *repository.Repository, log *zap.Logger) InventoryService {
	return &inventoryService{
		tx: tx,
		repo: repo,
		log: log,
	}
}

func (s *inventoryService) GetInventoryLogs(ctx context.Context, req request.InventoryQueryParams) (*response.PaginatedResponse[response.InventoryLogResponse], error) {
	logs, total, err := s.repo.InventoryRepo.GetInventoryLogs(ctx, req)
	if err != nil {
		s.log.Error("Error get inventory logs service", zap.Error(err))
		return nil, err
	}

	// Convert to DTO
	var res []response.InventoryLogResponse
	for _, l := range logs {
		log := response.InventoryLogResponse{
			ID: l.ID,
			SKUID: l.SKUID,
			Type: l.Type,
			QuantityChange: l.QuantityChange,
			CurrentStockAfter: l.CurrentStockAfter,
			ReferenceID: l.ReferenceID,
			ReferenceType: l.ReferenceType,
			Notes: l.Notes,
			CreatedAt: l.CreatedAt,
		}
		res = append(res, log)
	}	

	return response.NewPaginatedResponse(
		res,
		req.Page,
		req.Limit,
		total,
	), nil
}

func (s *inventoryService) GetInventoryLogsBySKUID(ctx context.Context, skuID uint, req request.InventoryQueryParams) (*response.PaginatedResponse[response.InventoryLogResponse], error) {
	logs, total, err := s.repo.InventoryRepo.GetInventoryLogsBySKUID(ctx, skuID, req)
	if err != nil {
		s.log.Error("Error get inventory logs service", zap.Error(err))
		return nil, err
	}

	// Convert to DTO
	var res []response.InventoryLogResponse
	for _, l := range logs {
		log := response.InventoryLogResponse{
			ID: l.ID,
			SKUID: l.SKUID,
			Type: l.Type,
			QuantityChange: l.QuantityChange,
			CurrentStockAfter: l.CurrentStockAfter,
			ReferenceID: l.ReferenceID,
			ReferenceType: l.ReferenceType,
			Notes: l.Notes,
			CreatedAt: l.CreatedAt,
		}
		res = append(res, log)
	}	

	return response.NewPaginatedResponse(
		res,
		req.Page,
		req.Limit,
		total,
	), nil
}

func (s *inventoryService) CreateInventoryLog(ctx context.Context, req request.CreateInventoryLogRequest) (*response.InventoryLogResponse, error) {
	inventory := entity.InventoryLog{
		SKUID: req.SKUID,
		QuantityChange: req.QuantityChange,
		Notes: req.Notes,
	}

	err := s.tx.WithinTx(ctx, func(ctx context.Context) error {
		// Get current stock
		sku, err := s.repo.SKURepo.GetByID(ctx, req.SKUID)
		if err != nil {
			s.log.Error("Failed to get SKU")
			return err
		}

		// Construct inventory log
		inventory.CurrentStockAfter = sku.Stock + inventory.QuantityChange

		if req.Action == "restock" {
			inventory.Type = entity.InventoryLogTypeIn
			inventory.ReferenceType = "purchase"
		}

		if req.Action == "adjustment" {
			inventory.Type = entity.InventoryLogTypeAdjustment
			inventory.ReferenceType = "adjustment"
		}

		sku.Stock += inventory.QuantityChange
		err = s.repo.SKURepo.Update(ctx, sku.ID, sku)
		if err != nil {
			s.log.Error("Failed to update stock")
			return err
		}

		log, err := s.repo.InventoryRepo.CreateInventoryLog(ctx, &inventory)
		if err != nil {
			return err
		}
		inventory.ID = log.ID
		return nil
	})

	if err != nil {
		s.log.Error("Error create inventory log transaction", zap.Error(err))
		return nil, err
	}

	// Construct response
	res := response.InventoryLogResponse{
		ID: inventory.ID,
		SKUID: inventory.SKUID,
		Type: inventory.Type,
		QuantityChange: inventory.QuantityChange,
		CurrentStockAfter: inventory.CurrentStockAfter,
		ReferenceType: inventory.ReferenceType,
		Notes: inventory.Notes,
		CreatedAt: inventory.CreatedAt,
	}
		
	return &res, nil
}

func (s *inventoryService) 	GetInventory(ctx context.Context, req request.InventoryQueryParams) (*response.PaginatedResponse[response.InventoryResponse], error) {
	inventories, total, err := s.repo.InventoryRepo.GetInventory(ctx, req)
	if err != nil {
		s.log.Error("Error get inventory", zap.Error(err))
		return nil, err
	}

	return response.NewPaginatedResponse(
		inventories,
		req.Page,
		req.Limit,
		total,
	), nil
}

func (s *inventoryService) GetInventoryBySKUID(ctx context.Context, skuID uint) (*response.InventoryResponse, error) {
	inventory, err := s.repo.InventoryRepo.GetInventoryBySKUID(ctx, skuID)
	if err != nil {
		s.log.Error("Error get inventory", zap.Error(err))
		return nil, err
	}

	return inventory, nil
}
