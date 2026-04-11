package usecase

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/dto/response"
	"debian-ecommerce/pkg/utils"
	"errors"
	"fmt"

	"debian-ecommerce/internal/data/repository"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type EmployeeUsecase interface {
	CreateEmployeeByAdmin(ctx context.Context, req request.CreateEmployeeByAdminRequest) error
	GetEmployeeList(ctx context.Context, page, limit int, search, sortBy, sortOrder string) (*response.PaginatedResponse[response.EmployeeResponse], error)
	GetEmployeeByID(ctx context.Context, id uint) (*response.EmployeeResponse, error)
	UpdateEmployee(ctx context.Context, id uint, req request.UpdateEmployeeRequest) error
	DeleteEmployee(ctx context.Context, id uint) error
}

type employeeUsecase struct {
	UserRepo     repository.UserRepository
	EmployeeRepo repository.EmployeeRepository
	Tx           TxManager
	EmailService utils.EmailService
	Log          *zap.Logger
}

func NewEmployeeUsecase(
	tx TxManager,
	userRepo repository.UserRepository,
	employeeRepo repository.EmployeeRepository,
	emailService utils.EmailService,
	log *zap.Logger,
) EmployeeUsecase {
	return &employeeUsecase{
		UserRepo:     userRepo,
		EmployeeRepo: employeeRepo,
		Tx:           tx,
		EmailService: emailService,
		Log:          log,
	}
}

func (u *employeeUsecase) CreateEmployeeByAdmin(ctx context.Context, req request.CreateEmployeeByAdminRequest) error {
	// Check if email exists
	_, err := u.UserRepo.FindUserByEmail(ctx, req.Email)
	if err == nil {
		return errors.New("email already registered")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	randomPassword, err := utils.GenerateRandomString(10)
	if err != nil {
		return fmt.Errorf("failed to generate password: %w", err)
	}
	hashedPassword := utils.HashPassword(randomPassword)

	// Transaction
	err = u.Tx.WithinTx(ctx, func(ctx context.Context) error {
		user := &entity.User{
			Email:        req.Email,
			PasswordHash: hashedPassword,
			Role:         entity.UserRole(req.Role),
			IsActive:     true,
		}
		if _, err := u.UserRepo.CreateUser(ctx, user); err != nil {
			return err
		}

		employee := &entity.Employee{
			UserID:   user.ID,
			FullName: req.Name,
		}
		if err := u.EmployeeRepo.CreateEmployee(ctx, employee); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	// Send Email (Best effort, outside transaction)
	subject := "Welcome to Lumoshive"
	body := fmt.Sprintf("Hello %s,\n\nYour account has been created successfully.\n\nEmail: %s\nPassword: %s\n\nPlease login and change your password immediately.", req.Name, req.Email, randomPassword)

	if err := u.EmailService.SendEmail(req.Email, subject, body); err != nil {
		u.Log.Error("failed to send welcome email", zap.Error(err), zap.String("email", req.Email))
		return nil // Return nil as user creation was successful
	}

	return nil
}

func (u *employeeUsecase) GetEmployeeList(ctx context.Context, page, limit int, search, sortBy, sortOrder string) (*response.PaginatedResponse[response.EmployeeResponse], error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	employees, total, err := u.EmployeeRepo.ListEmployees(ctx, limit, offset, search, sortBy, sortOrder)
	if err != nil {
		return nil, err
	}

	var list []response.EmployeeResponse
	for _, e := range employees {
		role := ""
		if e.User.ID != 0 {
			role = string(e.User.Role)
		}
		list = append(list, response.EmployeeResponse{
			ID:        e.ID,
			UserID:    e.UserID,
			Name:      e.FullName,
			Email:     e.User.Email,
			Role:      role,
			Phone:     e.Phone,
			Salary:    e.Salary,
			IsActive:  e.User.IsActive,
			CreatedAt: e.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	// Perbaikan: Panggil NewPaginatedResponse sebagai FUNCTION, bukan sebagai struct literal
	return response.NewPaginatedResponse(
		list,  // data []EmployeeResponse
		page,  // page int
		limit, // perPage int
		total, // total int64
	), nil
}

func (u *employeeUsecase) GetEmployeeByID(ctx context.Context, id uint) (*response.EmployeeResponse, error) {
	employee, err := u.EmployeeRepo.FindEmployeeByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("employee not found")
		}
		return nil, err
	}

	return &response.EmployeeResponse{
		ID:        employee.ID,
		UserID:    employee.UserID,
		Name:      employee.FullName,
		Email:     employee.User.Email,
		Role:      string(employee.User.Role),
		Phone:     employee.Phone,
		Salary:    employee.Salary,
		IsActive:  employee.User.IsActive,
		CreatedAt: employee.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (u *employeeUsecase) UpdateEmployee(ctx context.Context, id uint, req request.UpdateEmployeeRequest) error {
	employee, err := u.EmployeeRepo.FindEmployeeByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("employee not found")
		}
		return err
	}

	err = u.Tx.WithinTx(ctx, func(ctx context.Context) error {
		if req.Name != "" {
			employee.FullName = req.Name
		}
		if req.Phone != "" {
			employee.Phone = req.Phone
		}
		if req.Salary > 0 {
			employee.Salary = req.Salary
		}

		if err := u.EmployeeRepo.UpdateEmployee(ctx, employee); err != nil {
			return err
		}

		// Update User fields if necessary
		if req.Email != "" || req.Role != "" {
			user, err := u.UserRepo.FindUserByID(ctx, employee.UserID)
			if err != nil {
				return err
			}

			if req.Email != "" {
				user.Email = req.Email
			}
			if req.Role != "" {
				user.Role = entity.UserRole(req.Role)
			}

			if err := u.UserRepo.UpdateUser(ctx, user); err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func (u *employeeUsecase) DeleteEmployee(ctx context.Context, id uint) error {
	// First find the employee to get user ID
	employee, err := u.EmployeeRepo.FindEmployeeByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("employee not found")
		}
		return err
	}

	// Delete user (cascade should handle relation logic if configured, but here we explicitly delete User)
	// Usually deleting user is enough.
	// But let's check constraints.
	// entity/customer.go: `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	// entity/employee.go: `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	// So deleting User should delete Employee.

	// Transaction
	err = u.Tx.WithinTx(ctx, func(ctx context.Context) error {
		// Delete User
		if err := u.UserRepo.DeleteUser(ctx, employee.UserID); err != nil {
			return err
		}
		return nil
	})

	return err
}
