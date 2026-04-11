package usecase

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/data/repository"
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/dto/response"
	"debian-ecommerce/pkg/utils"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthUsecase interface {
	RegisterCustomer(ctx context.Context, req request.RegisterCustomerRequest) (*response.AuthResponse, error)
	RegisterEmployee(ctx context.Context, req request.RegisterEmployeeRequest) (*response.AuthResponse, error)
	Login(ctx context.Context, req request.LoginRequest) (*response.AuthResponse, error)
	IsEmailAvailable(ctx context.Context, email string) (bool, error)
	Logout(ctx context.Context, userID uint) error
	RequestResetPassword(ctx context.Context,	req request.ForgotPasswordRequest) error
	ResetPassword(ctx context.Context, req request.ResetPasswordRequest) error
}

type authUsecase struct {
	UserRepo     repository.UserRepository
	CustomerRepo repository.CustomerRepository
	EmployeeRepo repository.EmployeeRepository
	TokenRepo    repository.TokenRepository
	ResetPasswordRepo repository.ResetPasswordRepository
	EmailService utils.EmailService
	Tx           TxManager
	TokenService utils.TokenService
	Log          *zap.Logger
}

func NewAuthUsecase(
	tx TxManager,
	userRepo repository.UserRepository,
	customerRepo repository.CustomerRepository,
	employeeRepo repository.EmployeeRepository,
	tokenRepo repository.TokenRepository,
	tokenService utils.TokenService,
	log *zap.Logger,
) AuthUsecase {
	return &authUsecase{
		UserRepo:     userRepo,
		CustomerRepo: customerRepo,
		EmployeeRepo: employeeRepo,
		TokenRepo:    tokenRepo,
		Tx:           tx,
		TokenService: tokenService,
		Log:          log,
	}
}

func (u *authUsecase) RegisterCustomer(ctx context.Context, req request.RegisterCustomerRequest) (*response.AuthResponse, error) {
	// Check if email exists
	_, err := u.UserRepo.FindUserByEmail(ctx, req.Email)
	if err == nil {
		return nil, errors.New("email already registered")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hashedPassword := utils.HashPassword(req.Password)

	var token string
	var createdUser *entity.User
	err = u.Tx.WithinTx(ctx, func(ctx context.Context) error {
		user := &entity.User{
			Email:        req.Email,
			PasswordHash: hashedPassword,
			Role:         entity.RoleCustomer,
			IsActive:     true,
		}
		user, err := u.UserRepo.CreateUser(ctx, user)
		if err != nil {
			return err
		}
		createdUser = user

		customer := &entity.Customer{
			UserID:   user.ID,
			FullName: req.Name,
		}
		if err := u.CustomerRepo.CreateCustomer(ctx, customer); err != nil {
			return err
		}

		t, err := u.TokenService.GenerateToken(user.ID, string(user.Role))
		if err != nil {
			return err
		}

		if err := u.TokenRepo.SaveToken(ctx, user.ID, t, 24*time.Hour); err != nil {
			return err
		}
		token = t

		return nil
	})

	if err != nil {
		return nil, err
	}

	userResponse := response.UserResponse{
		ID: createdUser.ID,
		Email: createdUser.Email,
		Role: createdUser.Role,
	}

	return &response.AuthResponse{User: userResponse, Token: token}, nil
}

func (u *authUsecase) Login(ctx context.Context, req request.LoginRequest) (*response.AuthResponse, error) {
	user, err := u.UserRepo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, errors.New("invalid email or password")
	}

	token, err := u.TokenService.GenerateToken(user.ID, string(user.Role))
	if err != nil {
		return nil, err
	}

	if err := u.TokenRepo.SaveToken(ctx, user.ID, token, 24*time.Hour); err != nil {
		return nil, err
	}

	userResponse := response.UserResponse{
		ID: user.ID,
		Email: user.Email,
		Role: user.Role,
	}

	return &response.AuthResponse{User: userResponse, Token: token}, nil
}

func (u *authUsecase) RegisterEmployee(ctx context.Context, req request.RegisterEmployeeRequest) (*response.AuthResponse, error) {
	// Check if email exists
	_, err := u.UserRepo.FindUserByEmail(ctx, req.Email)
	if err == nil {
		return nil, errors.New("email already registered")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hashedPassword := utils.HashPassword(req.Password)

	var token string
	var createdUser *entity.User
	err = u.Tx.WithinTx(ctx, func(ctx context.Context) error {
		user := &entity.User{
			Email:        req.Email,
			PasswordHash: hashedPassword,
			Role:         entity.RoleAdmin, // Assuming admin for employee as per user request context usually implies admin/employee
			IsActive:     true,
		}
		user, err := u.UserRepo.CreateUser(ctx, user)
		if err != nil {
			return err
		}
		createdUser = user

		employee := &entity.Employee{
			UserID:   user.ID,
			FullName: req.Name,
		}
		if err := u.EmployeeRepo.CreateEmployee(ctx, employee); err != nil {
			return err
		}

		t, err := u.TokenService.GenerateToken(user.ID, string(user.Role))
		if err != nil {
			return err
		}

		if err := u.TokenRepo.SaveToken(ctx, user.ID, t, 24*time.Hour); err != nil {
			return err
		}
		token = t

		return nil
	})

	if err != nil {
		return nil, err
	}

	userResponse := response.UserResponse{
		ID: createdUser.ID,
		Email: createdUser.Email,
		Role: createdUser.Role,
	}

	return &response.AuthResponse{User: userResponse, Token: token}, nil
}

func (u *authUsecase) Logout(ctx context.Context, userID uint) error {
	return u.TokenRepo.DeleteToken(ctx, userID)
}
