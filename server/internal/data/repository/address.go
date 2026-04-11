package repository

import (
	"debian-ecommerce/internal/data/entity"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AddressRepository interface {
	Create(address *entity.Address) error
	Update(address *entity.Address) error
	Delete(id uint) error
	FindByID(id uint) (*entity.Address, error)
	FindByCustomerID(customerID uint) ([]entity.Address, error)
	SetDefault(customerID uint, addressID uint) error

	GetProvinces() ([]entity.Province, error)
	GetRegencies(provinceID uint) ([]entity.Regency, error)
	GetDistricts(regencyID uint) ([]entity.District, error)
	GetSubdistricts(districtID uint) ([]entity.Subdistrict, error)
}

type addressRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewAddressRepository(db *gorm.DB, log *zap.Logger) AddressRepository {
	return &addressRepository{
		db:  db,
		log: log,
	}
}

func (r *addressRepository) Create(address *entity.Address) error {
	if address.IsDefault {
		// Reset all addresses for this customer to not default
		if err := r.db.Model(&entity.Address{}).Where("customer_id = ?", address.CustomerID).Update("is_default", false).Error; err != nil {
			return err
		}
	}
	return r.db.Create(address).Error
}

func (r *addressRepository) Update(address *entity.Address) error {
	return r.db.Save(address).Error
}

func (r *addressRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Address{}, id).Error
}

func (r *addressRepository) FindByID(id uint) (*entity.Address, error) {
	var address entity.Address
	if err := r.db.Preload("Province").Preload("Regency").Preload("District").Preload("Subdistrict").First(&address, id).Error; err != nil {
		return nil, err
	}
	return &address, nil
}

func (r *addressRepository) FindByCustomerID(customerID uint) ([]entity.Address, error) {
	var addresses []entity.Address
	if err := r.db.Where("customer_id = ?", customerID).
		Preload("Province").Preload("Regency").Preload("District").Preload("Subdistrict").
		Find(&addresses).Error; err != nil {
		r.log.Error("Failed to get addressess", zap.Error(err))
		return nil, err
	}
	return addresses, nil
}

func (r *addressRepository) SetDefault(customerID uint, addressID uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Reset all addresses for this customer to not default
		if err := tx.Model(&entity.Address{}).Where("customer_id = ?", customerID).Update("is_default", false).Error; err != nil {
			return err
		}
		// Set the specific address to default
		if err := tx.Model(&entity.Address{}).Where("id = ? AND customer_id = ?", addressID, customerID).Update("is_default", true).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *addressRepository) GetProvinces() ([]entity.Province, error) {
	var provinces []entity.Province
	if err := r.db.Find(&provinces).Error; err != nil {
		return nil, err
	}
	return provinces, nil
}

func (r *addressRepository) GetRegencies(provinceID uint) ([]entity.Regency, error) {
	var regencies []entity.Regency
	if err := r.db.Where("province_id = ?", provinceID).Find(&regencies).Error; err != nil {
		return nil, err
	}
	return regencies, nil
}

func (r *addressRepository) GetDistricts(regencyID uint) ([]entity.District, error) {
	var districts []entity.District
	if err := r.db.Where("regency_id = ?", regencyID).Find(&districts).Error; err != nil {
		return nil, err
	}
	return districts, nil
}

func (r *addressRepository) GetSubdistricts(districtID uint) ([]entity.Subdistrict, error) {
	var subdistricts []entity.Subdistrict
	if err := r.db.Where("district_id = ?", districtID).Find(&subdistricts).Error; err != nil {
		return nil, err
	}
	return subdistricts, nil
}
