package usecase

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/data/repository"
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/dto/response"
	"debian-ecommerce/pkg/utils"

	"go.uber.org/zap"
)

type PaymentMethodUsecase struct {
	paymentMethodRepo repository.PaymentMethodRepository // Gunakan interface, bukan pointer
	tx                TxManager
	log               *zap.Logger
}

func NewPaymentMethodUsecase(
	paymentMethodRepo repository.PaymentMethodRepository, // Interface
	tx TxManager,
	log *zap.Logger,
) *PaymentMethodUsecase {
	return &PaymentMethodUsecase{
		paymentMethodRepo: paymentMethodRepo,
		tx:                tx,
		log:               log,
	}
}

// toResponse converts entity to response DTO
func (u *PaymentMethodUsecase) toResponse(pm *entity.PaymentMethod) *response.PaymentMethodResponse {
	if pm == nil {
		return nil
	}
	return &response.PaymentMethodResponse{
		ID:        pm.ID,
		Name:      pm.Name,
		IsActive:  pm.IsActive,
		IconURL:   pm.IconURL,
	}
}

// CreatePaymentMethod creates a new payment method
func (u *PaymentMethodUsecase) CreatePaymentMethod(ctx context.Context, req *request.CreatePaymentMethodRequest) (*response.PaymentMethodResponse, error) {
	// Validate required fields
	if req.Name == "" {
		u.log.Warn("payment method name is required")
		return nil, utils.ErrNameRequired
	}

	var result *response.PaymentMethodResponse
	err := u.tx.WithinTx(ctx, func(ctx context.Context) error {
		// Check if payment method with same name exists
		exists, err := u.paymentMethodRepo.CheckNameExists(ctx, req.Name, 0)
		if err != nil {
			u.log.Error("failed to check payment method name")
			return err
		}
		if exists {
			return utils.ErrDuplicatePaymentMethod
		}

		// Create payment method entity
		paymentMethod := &entity.PaymentMethod{
			Name: req.Name,
		}

		// Set is_active if provided
		if req.IsActive != nil {
			paymentMethod.IsActive = *req.IsActive
		}

		// Set icon_url if provided
		if req.IconURL != "" {
			paymentMethod.IconURL = req.IconURL
		}

		// Save to database
		if err := u.paymentMethodRepo.Create(ctx, paymentMethod); err != nil {
			u.log.Error("failed to create payment method")
			return err
		}

		result = u.toResponse(paymentMethod)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdatePaymentMethod updates an existing payment method
func (u *PaymentMethodUsecase) UpdatePaymentMethod(ctx context.Context, id uint, req *request.UpdatePaymentMethodRequest) (*response.PaymentMethodResponse, error) {
	if id == 0 {
		return nil, utils.ErrInvalidUserID
	}

	var result *response.PaymentMethodResponse
	err := u.tx.WithinTx(ctx, func(ctx context.Context) error {
		// Find payment method by ID
		paymentMethod, err := u.paymentMethodRepo.FindByID(ctx, id)
		if err != nil {
			u.log.Error("failed to find payment method")
			return err
		}
		if paymentMethod == nil {
			return utils.ErrPaymentMethodNotFound
		}

		// Check if name already exists (if updating name)
		if req.Name != "" && req.Name != paymentMethod.Name {
			exists, err := u.paymentMethodRepo.CheckNameExists(ctx, req.Name, id)
			if err != nil {
				u.log.Error("failed to check payment method name")
				return err
			}
			if exists {
				u.log.Warn("payment method name already exists")
				return utils.ErrDuplicatePaymentMethod
			}
			paymentMethod.Name = req.Name
		}

		// Update fields if provided
		if req.IsActive != nil {
			paymentMethod.IsActive = *req.IsActive
		}
		if req.IconURL != "" {
			paymentMethod.IconURL = req.IconURL
		}

		// Save updates
		if err := u.paymentMethodRepo.Update(ctx, paymentMethod); err != nil {
			u.log.Error("failed to update payment method")
			return err
		}

		result = u.toResponse(paymentMethod)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetPaymentMethodByID retrieves payment method by ID
func (u *PaymentMethodUsecase) GetPaymentMethodByID(ctx context.Context, id uint) (*response.PaymentMethodResponse, error) {
	if id == 0 {
		return nil, utils.ErrInvalidUserID
	}

	paymentMethod, err := u.paymentMethodRepo.FindByID(ctx, id)
	if err != nil {
		u.log.Error("failed to find payment method")
		return nil, err
	}
	if paymentMethod == nil {
		return nil, utils.ErrPaymentMethodNotFound
	}

	return u.toResponse(paymentMethod), nil
}

// GetPaymentMethods retrieves all payment methods with pagination
func (u *PaymentMethodUsecase) GetPaymentMethods(ctx context.Context, req *request.GetPaymentMethodsRequest) (*response.PaymentListResponse, error) {
	// Set default values
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}

	paymentTypes, total, err := u.paymentMethodRepo.FindAll(ctx, req.Page, req.Limit, req.Search, req.IsActive)
	if err != nil {
		u.log.Error("failed to get payment methods")
		return nil, err
	}

	totalPages := int64(0)
	if total > 0 {
		totalPages = total / int64(req.Limit)
		if total%int64(req.Limit) > 0 {
			totalPages++
		}
	}

	return &response.PaymentListResponse{
		PaymentTypes: paymentTypes,
		Total: total,
		Page: req.Page,
		Limit: req.Limit,
		TotalPages: totalPages,
	}, nil
}

func (u *PaymentMethodUsecase) GetAllPaymentMethods(ctx context.Context) ([]entity.PaymentType, error) {
	paymentTypes, err := u.paymentMethodRepo.GetAll(ctx)
	if err != nil {
		u.log.Error("failed to get payment methods")
		return nil, err
	}

	return paymentTypes, nil
}

// DeletePaymentMethod soft deletes a payment method
func (u *PaymentMethodUsecase) DeletePaymentMethod(ctx context.Context, id uint) error {
	if id == 0 {
		return utils.ErrInvalidUserID
	}

	err := u.tx.WithinTx(ctx, func(ctx context.Context) error {
		// Check if payment method exists
		paymentMethod, err := u.paymentMethodRepo.FindByID(ctx, id)
		if err != nil {
			u.log.Error("failed to find payment method")
			return err
		}
		if paymentMethod == nil {
			return utils.ErrPaymentMethodNotFound
		}

		// Soft delete
		if err := u.paymentMethodRepo.Delete(ctx, paymentMethod.ID); err != nil {
			u.log.Error("failed to delete payment method")
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	u.log.Info("payment method deleted successfully")
	return nil
}

// TogglePaymentMethodStatus toggles the active status of a payment method
func (u *PaymentMethodUsecase) TogglePaymentMethodStatus(ctx context.Context, id uint) (*response.PaymentMethodResponse, error) {
	if id == 0 {
		return nil, utils.ErrInvalidUserID
	}

	var result *response.PaymentMethodResponse
	err := u.tx.WithinTx(ctx, func(ctx context.Context) error {
		// Find payment method by ID
		paymentMethod, err := u.paymentMethodRepo.FindByID(ctx, id)
		if err != nil {
			u.log.Error("failed to find payment method")
			return err
		}
		if paymentMethod == nil {
			return utils.ErrPaymentMethodNotFound
		}

		// Toggle status
		paymentMethod.IsActive = !paymentMethod.IsActive

		// Save updates
		if err := u.paymentMethodRepo.Update(ctx, paymentMethod); err != nil {
			u.log.Error("failed to update payment method status")
			return err
		}

		result = u.toResponse(paymentMethod)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
