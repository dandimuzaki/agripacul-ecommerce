package usecase

import (
	"debian-ecommerce/internal/data/repository"
	"debian-ecommerce/pkg/utils"

	"go.uber.org/zap"
)

type Usecase struct {
	ProductService  ProductService
	TokenService    utils.TokenService
	EmailService    utils.EmailService
	CategoryService CategoryUsecase
	AuthService     AuthUsecase
	EmployeeService EmployeeUsecase
	BannerService   BannerUsecase
	ReviewService   ReviewUsecase
	CartService     CartUsecase
	CustomerService CustomerUsecase
	WishlistService WishlistUsecase
	CompanyService  CompanyUsecase
	OrderService OrderService
	InventoryService InventoryService
	CheckoutService CheckoutService
	AddressUsecase AddressUsecase
	PromotionUsecase PromotionUsecase
	ReportService ReportService
	CampaignService CampaignService
	PaymentMethodUsecase PaymentMethodUsecase
}

func NewUsecase(
	tx TxManager, 
	repo *repository.Repository, 
	tokenRepo repository.TokenRepository, 
	log *zap.Logger, 
	tokenService utils.TokenService, 
	emailService utils.EmailService,
	cloudinary ImageUploader, 
	config utils.Configuration,
	) *Usecase {
	return &Usecase{
		ProductService:  NewProductService(tx, cloudinary, repo, log),
		TokenService:    tokenService,
		EmailService:    emailService,
		CategoryService: NewCategoryUsecase(tx, cloudinary, repo.CategoryRepo, log),
		AuthService:     NewAuthUsecase(tx, repo.UserRepo, repo.CustomerRepo, repo.EmployeeRepo, tokenRepo, tokenService, log),
		EmployeeService: NewEmployeeUsecase(tx, repo.UserRepo, repo.EmployeeRepo, emailService, log),
		BannerService:   NewBannerUsecase(tx, cloudinary, repo.BannerRepo, log),
		ReviewService:   NewReviewUsecase(tx, repo, log),
		CartService:     NewCartUsecase(repo, log),
		CustomerService: NewCustomerUsecase(repo.CustomerRepo, cloudinary, log),
		WishlistService: NewWishlistUsecase(repo.WishlistRepo, repo.CustomerRepo, repo.ProductRepo, log),
		CompanyService:  NewCompanyUsecase(repo.CompanyRepo, log, emailService),
		OrderService: NewOrderService(tx, repo, log),
		InventoryService: NewInventoryService(tx, repo, log),
		CheckoutService: NewCheckoutService(tx, repo, log, config),
		AddressUsecase: NewAddressUsecase(*repo, log),
		PromotionUsecase: NewPromotionUsecase(repo.PromotionRepo, tx),
		ReportService: NewReportService(tx, repo, log),
		CampaignService: NewCampaignService(tx, repo, log),
		PaymentMethodUsecase: *NewPaymentMethodUsecase(repo.PaymentMethodRepo, tx, log),
	}
}
