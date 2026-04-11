package usecase

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/data/repository"
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/dto/response"
	"debian-ecommerce/pkg/utils"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CompanyUsecase interface {
	CreateCompany(ctx context.Context, req request.CreateCompanyRequest) error
	GetCompany(ctx context.Context) (*response.CompanyResponse, error)
	UpdateCompany(ctx context.Context, req request.UpdateCompanyRequest) error
	CreateAddress(ctx context.Context, req request.CreateCompanyAddressRequest) error
	GetAddressByID(ctx context.Context, id uint) (*response.CompanyAddressResponse, error)
	GetShippingOriginAddress(ctx context.Context, companyID uint) (*response.CompanyAddressResponse, error)
	UpdateAddress(ctx context.Context, id uint, req request.UpdateCompanyAddressRequest) error
	SendMessage(ctx context.Context, req request.SendMessageRequest) error
}

type companyUsecase struct {
	repo repository.CompanyRepository
	log  *zap.Logger
	EmailService utils.EmailService
}

func NewCompanyUsecase(repo repository.CompanyRepository, log *zap.Logger, emailService utils.EmailService) CompanyUsecase {
	return &companyUsecase{
		repo: repo,
		log:  log,
		EmailService: emailService,
	}
}

func (u *companyUsecase) CreateCompany(ctx context.Context, req request.CreateCompanyRequest) error {
	// Check if company already exists (assuming single company profile)
	_, err := u.repo.GetCompany(ctx)
	if err == nil {
		return errors.New("company profile already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	company := &entity.Company{
		Name:         req.Name,
		Description:  req.Description,
		InstagramURL: req.InstagramURL,
		TwitterURL:   req.TwitterURL,
		WhatsappURL:  req.WhatsappURL,
		ContactEmail: req.ContactEmail,
	}

	return u.repo.CreateCompany(ctx, company)
}

func (u *companyUsecase) GetCompany(ctx context.Context) (*response.CompanyResponse, error) {
	company, err := u.repo.GetCompany(ctx)
	if err != nil {
		return nil, err
	}

	return &response.CompanyResponse{
		ID:           company.ID,
		Name:         company.Name,
		Description:  company.Description,
		InstagramURL: company.InstagramURL,
		TwitterURL:   company.TwitterURL,
		WhatsappURL:  company.WhatsappURL,
		ContactEmail: company.ContactEmail,
	}, nil
}

func (u *companyUsecase) UpdateCompany(ctx context.Context, req request.UpdateCompanyRequest) error {
	company, err := u.repo.GetCompany(ctx)
	if err != nil {
		return err
	}

	if req.Name != "" {
		company.Name = req.Name
	}
	if req.Description != "" {
		company.Description = req.Description
	}
	if req.InstagramURL != "" {
		company.InstagramURL = req.InstagramURL
	}
	if req.TwitterURL != "" {
		company.TwitterURL = req.TwitterURL
	}
	if req.WhatsappURL != "" {
		company.WhatsappURL = req.WhatsappURL
	}
	if req.ContactEmail != "" {
		company.ContactEmail = req.ContactEmail
	}

	return u.repo.UpdateCompany(ctx, company)
}

func (u *companyUsecase) CreateAddress(ctx context.Context, req request.CreateCompanyAddressRequest) error {
	// If IsShippingOrigin is true, unset others
	if req.IsShippingOrigin {
		if err := u.repo.UnsetShippingOrigin(ctx, req.CompanyID); err != nil {
			return err
		}
	}

	address := &entity.CompanyAddress{
		CompanyID:        req.CompanyID,
		ProvinceID:       req.ProvinceID,
		RegencyID:        req.RegencyID,
		DistrictID:       req.DistrictID,
		SubdistrictID:    req.SubdistrictID,
		PostalCode:       req.PostalCode,
		DetailAddress:    req.DetailAddress,
		IsShippingOrigin: req.IsShippingOrigin,
	}

	return u.repo.CreateAddress(ctx, address)
}

func (u *companyUsecase) GetAddressByID(ctx context.Context, id uint) (*response.CompanyAddressResponse, error) {
	address, err := u.repo.GetAddressByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &response.CompanyAddressResponse{
		ID:               address.ID,
		CompanyID:        address.CompanyID,
		Province:         address.Province.Name,
		Regency:          address.Regency.Name,
		District:         address.District.Name,
		Subdistrict:      address.Subdistrict.Name,
		PostalCode:       address.PostalCode,
		DetailAddress:    address.DetailAddress,
		IsShippingOrigin: address.IsShippingOrigin,
	}, nil
}

func (u *companyUsecase) GetShippingOriginAddress(ctx context.Context, companyID uint) (*response.CompanyAddressResponse, error) {
	address, err := u.repo.GetShippingOriginAddress(ctx)
	if err != nil {
		return nil, err
	}

	return &response.CompanyAddressResponse{
		ID:               address.ID,
		CompanyID:        address.CompanyID,
		Province:         address.Province.Name,
		Regency:          address.Regency.Name,
		District:         address.District.Name,
		Subdistrict:      address.Subdistrict.Name,
		PostalCode:       address.PostalCode,
		DetailAddress:    address.DetailAddress,
		IsShippingOrigin: address.IsShippingOrigin,
	}, nil
}

func (u *companyUsecase) UpdateAddress(ctx context.Context, id uint, req request.UpdateCompanyAddressRequest) error {
	address, err := u.repo.GetAddressByID(ctx, id)
	if err != nil {
		return err
	}

	if req.ProvinceID != 0 {
		address.ProvinceID = req.ProvinceID
	}
	if req.RegencyID != 0 {
		address.RegencyID = req.RegencyID
	}
	if req.DistrictID != 0 {
		address.DistrictID = req.DistrictID
	}
	if req.SubdistrictID != 0 {
		address.SubdistrictID = req.SubdistrictID
	}
	if req.PostalCode != "" {
		address.PostalCode = req.PostalCode
	}
	if req.DetailAddress != "" {
		address.DetailAddress = req.DetailAddress
	}
	if req.IsShippingOrigin != nil {
		if *req.IsShippingOrigin {
			if err := u.repo.UnsetShippingOrigin(ctx, address.CompanyID); err != nil {
				return err
			}
		}
		address.IsShippingOrigin = *req.IsShippingOrigin
	}

	return u.repo.UpdateAddress(ctx, address)
}

func (u *companyUsecase) SendMessage(ctx context.Context, req request.SendMessageRequest) error {
	if err := u.EmailService.SendMessageFromCustomer(req.Subject, req.Body); err != nil {
		u.log.Error("failed to send message", zap.Error(err), zap.String("email", req.Email))
		return err
	}

	return nil
}