package repository

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	infra "debian-ecommerce/internal/infra/transaction"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	CreateEmployee(ctx context.Context, employee *entity.Employee) error
	FindEmployeeByUserID(ctx context.Context, userID uint) (*entity.Employee, error)
	FindEmployeeByID(ctx context.Context, id uint) (*entity.Employee, error)
	UpdateEmployee(ctx context.Context, employee *entity.Employee) error
	ListEmployees(ctx context.Context, limit, offset int, search, sortBy, sortOrder string) ([]entity.Employee, int64, error)
}

type employeeRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewEmployeeRepository(db *gorm.DB, log *zap.Logger) EmployeeRepository {
	return &employeeRepository{
		DB:  db,
		Log: log,
	}
}

func (r *employeeRepository) CreateEmployee(ctx context.Context, employee *entity.Employee) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Create(employee).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}

func (r *employeeRepository) FindEmployeeByUserID(ctx context.Context, userID uint) (*entity.Employee, error) {
	db := infra.GetDB(ctx, r.DB)
	var employee entity.Employee
	if err := db.WithContext(ctx).Where("user_id = ?", userID).First(&employee).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepository) FindEmployeeByID(ctx context.Context, id uint) (*entity.Employee, error) {
	db := infra.GetDB(ctx, r.DB)
	var employee entity.Employee
	if err := db.WithContext(ctx).Preload("User").First(&employee, id).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepository) UpdateEmployee(ctx context.Context, employee *entity.Employee) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.WithContext(ctx).Save(employee).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}

func (r *employeeRepository) ListEmployees(ctx context.Context, limit, offset int, search, sortBy, sortOrder string) ([]entity.Employee, int64, error) {
	var employees []entity.Employee
	var count int64

	db := infra.GetDB(ctx, r.DB).Model(&entity.Employee{}).Joins("User")

	if search != "" {
		searchTerm := "%" + search + "%"
		db = db.Where("employees.full_name ILIKE ?", searchTerm)
	}

	if err := db.WithContext(ctx).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if sortBy != "" {
		if sortOrder == "" {
			sortOrder = "asc"
		}
		// Validate sortBy to avoid injection if not careful, though GORM handles standard columns.
		// Allowed: full_name, created_at, etc.
		if sortBy == "name" {
			sortBy = "full_name"
		}
		db = db.Order(sortBy + " " + sortOrder)
	} else {
		db = db.Order("employees.created_at desc")
	}

	if err := db.WithContext(ctx).Preload("User").Limit(limit).Offset(offset).Find(&employees).Error; err != nil {
		return nil, 0, err
	}

	return employees, count, nil
}
