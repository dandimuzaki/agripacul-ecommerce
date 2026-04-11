package repository

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepo         UserRepository
	CustomerRepo     CustomerRepository
	EmployeeRepo     EmployeeRepository
	CategoryRepo     CategoryRepository
	ProductRepo      ProductRepository
	BannerRepo       BannerRepository
	ReviewRepo       ReviewRepository
	CartRepo         CartRepository
	VariantTypeRepo  VariantTypeRepository
	VariantValueRepo VariantValueRepository
	ImageRepo        ImageRepository
	SKURepo          SKURepository
	WishlistRepo     WishlistRepository
	CompanyRepo      CompanyRepository
	OrderRepo OrderRepository
	PromotionRepo PromotionRepository
	AddressRepo AddressRepository
	ReportRepo ReportRepository
	InventoryRepo InventoryRepository
	CampaignRepo CampaignRepository
	ResetPasswordRepo ResetPasswordRepository
	PaymentMethodRepo PaymentMethodRepository
}

func NewRepository(redis *redis.Client, db *gorm.DB, log *zap.Logger) *Repository {
	return &Repository{
		UserRepo:         NewUserRepository(db, log),
		CustomerRepo:     NewCustomerRepository(db, log),
		EmployeeRepo:     NewEmployeeRepository(db, log),
		ProductRepo:      NewProductRepo(db, log),
		CategoryRepo:     NewCategoryRepository(db),
		BannerRepo:       NewBannerRepository(db, log),
		ReviewRepo:       NewReviewRepository(db, log),
		CartRepo:         NewCartRepository(db, log),
		VariantTypeRepo:  NewVariantTypeRepo(db, log),
		VariantValueRepo: NewVariantValueRepo(db, log),
		ImageRepo:        NewImageRepo(db, log),
		SKURepo:          NewSKURepo(db, log),
		WishlistRepo:     NewWishlistRepository(db, log),
		CompanyRepo:      NewCompanyRepository(db, log),
		OrderRepo: NewOrderRepo(db, log),
		PromotionRepo: NewPromotionRepository(db, log),
		AddressRepo: NewAddressRepository(db, log),
		ReportRepo: NewReportRepo(db, log),
		InventoryRepo: NewInventoryRepo(db, log),
		CampaignRepo: NewCampaignRepo(db, log),
		ResetPasswordRepo: NewResetPasswordRepository(redis, log),
		PaymentMethodRepo: NewPaymentMethodRepository(db, log),
	}
}
