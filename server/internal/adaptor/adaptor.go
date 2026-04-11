package adaptor

import (
	"debian-ecommerce/internal/usecase"

	"go.uber.org/zap"
)

type Handler struct {
	AuthHandler     *AuthHandler
	EmployeeHandler *EmployeeHandler
	CategoryHandler *CategoryHandler
	ProductHandler  *ProductHandler
	BannerHandler   *BannerHandler
	ReviewHandler   *ReviewHandler
	WishlistHandler *WishlistHandler
	CartHandler     *CartHandler
	CustomerHandler *CustomerHandler
	CompanyHandler  *CompanyHandler
	OrderHandler *OrderHandler
	InventoryHandler *InventoryHandler
	AddressHandler *AddressHandler
	CheckoutHandler *CheckoutHandler
	PromotionHandler *PromotionHandler
	ReportHandler *ReportHandler
	CampaignHandler *CampaignHandler
	PaymentMethodHandler *PaymentMethodHandler
}

func NewHandler(u *usecase.Usecase, log *zap.Logger) Handler {
	return Handler{
		AuthHandler:     NewAuthHandler(u.AuthService, log),
		EmployeeHandler: NewEmployeeHandler(u.EmployeeService, log),
		CategoryHandler: NewCategoryHandler(u.CategoryService, log),
		ProductHandler:  NewProductHandler(u.ProductService, log),
		BannerHandler:   NewBannerHandler(u.BannerService, log),
		ReviewHandler:   NewReviewHandler(u.ReviewService, log),
		WishlistHandler: NewWishlistHandler(u.WishlistService, log),
		CartHandler:     NewCartHandler(u.CartService, log),
		CustomerHandler: NewCustomerHandler(u.CustomerService, log),
		CompanyHandler:  NewCompanyHandler(u.CompanyService, log),
		OrderHandler: NewOrderHandler(u.OrderService, log),
		InventoryHandler: NewInventoryHandler(u.InventoryService, log),
		AddressHandler: NewAddressHandler(u.AddressUsecase, u.CustomerService, log),
		CheckoutHandler: NewCheckoutHandler(u.CheckoutService, log),
		PromotionHandler: NewPromotionHandler(u.PromotionUsecase),
		ReportHandler: NewReportHandler(u.ReportService, log),
		CampaignHandler: NewCampaignHandler(u.CampaignService, log),
		PaymentMethodHandler: NewPaymentMethodHandler(&u.PaymentMethodUsecase, log),
	}
}
