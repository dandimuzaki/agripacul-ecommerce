package repository

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	infra "debian-ecommerce/internal/infra/transaction"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CompanyRepository interface {
	CreateCompany(ctx context.Context, company *entity.Company) error
	GetCompany(ctx context.Context) (*entity.Company, error)
	GetCompanyByID(ctx context.Context, id uint) (*entity.Company, error)
	UpdateCompany(ctx context.Context, company *entity.Company) error

	CreateAddress(ctx context.Context, address *entity.CompanyAddress) error
	GetAddressByID(ctx context.Context, id uint) (*entity.CompanyAddress, error)
	GetShippingOriginAddress(ctx context.Context) (*entity.CompanyAddress, error)
	UpdateAddress(ctx context.Context, address *entity.CompanyAddress) error
	UnsetShippingOrigin(ctx context.Context, companyID uint) error
}

type companyRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewCompanyRepository(db *gorm.DB, log *zap.Logger) CompanyRepository {
	return &companyRepository{
		DB:  db,
		Log: log,
	}
}

func (r *companyRepository) CreateCompany(ctx context.Context, company *entity.Company) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Create(company).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}

func (r *companyRepository) GetCompany(ctx context.Context) (*entity.Company, error) {
	db := infra.GetDB(ctx, r.DB)
	var company entity.Company
	if err := db.WithContext(ctx).First(&company).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *companyRepository) GetCompanyByID(ctx context.Context, id uint) (*entity.Company, error) {
	db := infra.GetDB(ctx, r.DB)
	var company entity.Company
	if err := db.WithContext(ctx).First(&company, id).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *companyRepository) UpdateCompany(ctx context.Context, company *entity.Company) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Save(company).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}

func (r *companyRepository) CreateAddress(ctx context.Context, address *entity.CompanyAddress) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Create(address).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}

func (r *companyRepository) GetAddressByID(ctx context.Context, id uint) (*entity.CompanyAddress, error) {
	db := infra.GetDB(ctx, r.DB)
	var address entity.CompanyAddress
	if err := db.WithContext(ctx).Preload("Province").Preload("Regency").Preload("District").Preload("Subdistrict").Preload("Company").First(&address, id).Error; err != nil {
		return nil, err
	}
	return &address, nil
}

func (r *companyRepository) GetShippingOriginAddress(ctx context.Context) (*entity.CompanyAddress, error) {
	db := infra.GetDB(ctx, r.DB)
	var address entity.CompanyAddress
	if err := db.WithContext(ctx).Where("is_shipping_origin = ?", true).Preload("Province").Preload("Regency").Preload("District").Preload("Subdistrict").First(&address).Error; err != nil {
		return nil, err
	}
	return &address, nil
}

func (r *companyRepository) UpdateAddress(ctx context.Context, address *entity.CompanyAddress) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Save(address).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}

func (r *companyRepository) UnsetShippingOrigin(ctx context.Context, companyID uint) error {
	db := infra.GetDB(ctx, r.DB)
	return db.WithContext(ctx).Model(&entity.CompanyAddress{}).Where("company_id = ?", companyID).Update("is_shipping_origin", false).Error
}
