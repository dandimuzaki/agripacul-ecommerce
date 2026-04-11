package wire

import (
	_ "debian-ecommerce/docs"
	"debian-ecommerce/internal/adaptor"
	"debian-ecommerce/internal/data/repository"
	"debian-ecommerce/internal/infra/scheduler"
	infra "debian-ecommerce/internal/infra/transaction"
	"debian-ecommerce/internal/infra/uploader"
	"debian-ecommerce/internal/usecase"
	mCustom "debian-ecommerce/pkg/middleware"
	"debian-ecommerce/pkg/utils"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	Route  *gin.Engine
	Stop   chan struct{}
	WG     *sync.WaitGroup
	Config utils.Configuration
	Scheduler *scheduler.Scheduler
}

func Wiring(db *gorm.DB, rdb *redis.Client, log *zap.Logger, config utils.Configuration) *App {
	r := gin.Default()
	
	r.Use(cors.New(cors.Config{
    AllowOrigins: []string{"http://localhost:3000", config.ClientHost},
    AllowMethods: []string{
        "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS",
    },
    AllowHeaders: []string{
        "Origin", "Content-Type", "Authorization", "Accept",
    },
    AllowCredentials: true,
    MaxAge: 12 * time.Hour,
	}))

	r1 := r.Group("/api/v1")
	r1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	stop := make(chan struct{})
	wg := &sync.WaitGroup{}

	tx := infra.NewGormTxManager(db)
	emailService := utils.NewEmailService(config.SMTP, log)
	tokenService := utils.NewTokenService(config.JWTSecret, config.Issuer)
	tokenRepo := repository.NewTokenRepository(rdb)

	repo := repository.NewRepository(rdb, db, log)
	cloudinary, err := uploader.NewUploader(config.CloudinaryConfig)
	if err != nil {
		log.Error("failed to init cloudinary", zap.Error(err))
	}
	usecase := usecase.NewUsecase(tx, repo, tokenRepo, log, tokenService, emailService, cloudinary, config)
	handler := adaptor.NewHandler(usecase, log)
	mw := mCustom.NewMiddlewareCustom(nil, log, tokenService, tokenRepo) // Passing nil for old Usecase if it was used.

	sch := scheduler.NewScheduler()
	scheduler.RegisterCronJobs(sch, usecase, log)

	ApiV1(r1, &handler, mw)

	return &App{
		Route:  r,
		Stop:   stop,
		WG:     wg,
		Config: config,
		Scheduler: sch,
	}
}

func ApiV1(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	AuthRoute(r.Group("/auth"), handler, mw)
	AdminAuthRoute(r.Group("/admin/auth"), handler, mw)
	EmployeeRoute(r.Group("/admin/employees"), handler, mw)
	CategoryRoute(r.Group("/categories"), handler, mw)
	ProductRoute(r.Group("/products"), handler, mw)
	BannerRoute(r.Group("/banners"), handler, mw)
	AddressRoute(r.Group("/address"), handler, mw)
	AdminBannerRoute(r.Group("/admin/banners"), handler, mw)
	ReviewRoute(r.Group("/reviews"), handler, mw)
	WishlistRoute(r.Group("/wishlist"), handler, mw)
	CartRoute(r.Group("/cart"), handler, mw)
	CheckoutRoute(r.Group("/checkout"), handler, mw)
	OrderRoute(r.Group("/orders"), handler, mw)
	CustomerRoute(r.Group("/profile"), handler, mw)
	DashboardRoute(r.Group("/admin"), handler, mw)
	CompanyRoute(r.Group("/company"), handler, mw)
	ContactRoute(r.Group("/contact"), handler, mw)
	AdminCompanyRoute(r.Group("/admin/company"), handler, mw)
	CampaignRoute(r.Group("/campaigns"), handler, mw)
	AdminCampaignRoute(r.Group("/admin/campaigns"), handler, mw)
	LocationRoute(r.Group(""), handler, mw)
	PaymentRoute(r.Group("/payment-methods"), handler, mw)
}

func CustomerRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.Use(mw.BearerTokenAuth())
	r.GET("", handler.CustomerHandler.GetProfile)
	r.PUT("", handler.CustomerHandler.UpdateProfile)
}

func CartRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.Use(mw.BearerTokenAuth())

	r.GET("", handler.CartHandler.GetCart)
	r.POST("/items", handler.CartHandler.AddItem)
	r.PUT("/items/:id", handler.CartHandler.UpdateItem)
	r.PUT("", handler.CartHandler.BatchSelectItem)
	r.DELETE("/items/:id", handler.CartHandler.RemoveItem)
	r.DELETE("", handler.CartHandler.ClearCart)
}

func AuthRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.POST("/register", handler.AuthHandler.RegisterCustomer)
	r.POST("/login", handler.AuthHandler.Login)
	r.POST("/check-email", handler.AuthHandler.CheckEmail)
	r.POST("/logout", mw.BearerTokenAuth(), handler.AuthHandler.Logout)
	r.POST("/forgot", handler.AuthHandler.ForgotPassword)
	r.POST("/reset", handler.AuthHandler.ResetPassword)
}

func AdminAuthRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.POST("/register", handler.AuthHandler.RegisterEmployee)
	r.POST("/logout", mw.BearerTokenAuth(), handler.AuthHandler.Logout)
}

func EmployeeRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.Use(mw.BearerTokenAuth())
	r.Use(mw.RBAC("admin", "superadmin"))

	r.GET("", handler.EmployeeHandler.GetEmployeeList)
	r.POST("", handler.EmployeeHandler.CreateEmployeeByAdmin)
	r.GET("/:id", handler.EmployeeHandler.GetEmployeeByID)
	r.PUT("/:id", handler.EmployeeHandler.UpdateEmployee)
	r.DELETE("/:id", handler.EmployeeHandler.DeleteEmployee)
}

func CategoryRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.GET("", handler.CategoryHandler.GetAllCategories)
	r.GET("/:id", handler.CategoryHandler.GetCategoryByID)
}

func ProductRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.GET("", handler.ProductHandler.BrowseProducts)
	r.GET("/:id", handler.ProductHandler.GetProductDetails)
	r.GET("/details/:slug", handler.ProductHandler.GetBySlug)
	r.GET("/:id/reviews", handler.ReviewHandler.GetReviewsByProduct)
}

func CheckoutRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.Use(mw.BearerTokenAuth())
	r.Use(mw.RBAC("customer"))
	
	r.POST("/preview", handler.CheckoutHandler.PreviewCheckout)
	r.POST("/shippings", handler.CheckoutHandler.GetShippingOptions)
	r.POST("/promotions", handler.CheckoutHandler.GetValidPromotions)
}

func OrderRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.Use(mw.BearerTokenAuth())
	r.Use(mw.RBAC("customer"))
	
	r.GET("", handler.OrderHandler.GetOrderHistory)
	r.GET("/:id", handler.OrderHandler.GetByID)
	r.POST("", handler.OrderHandler.Create)
	r.PUT("/:id/pay", handler.OrderHandler.Pay)
	r.PUT("/:id/complete", handler.OrderHandler.Complete)
	r.PUT("/:id/cancel", handler.OrderHandler.Cancel)
}

func BannerRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.GET("", handler.BannerHandler.GetBannerList)
	r.GET("/:id", handler.BannerHandler.GetBannerByID)
}

func LocationRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.GET("/provinces", handler.AddressHandler.GetProvinces)
	r.GET("/regencies/:province_id", handler.AddressHandler.GetRegencies)
	r.GET("/districts/:regency_id", handler.AddressHandler.GetDistricts)
	r.GET("/subdistricts/:district_id", handler.AddressHandler.GetSubdistricts)
}

func AddressRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.Use(mw.BearerTokenAuth())
	r.Use(mw.RBAC("customer"))

	r.GET("", handler.AddressHandler.GetByCustomerID)
	r.POST("", handler.AddressHandler.Create)
	r.PUT("/:id", handler.AddressHandler.Update)
	r.PUT("/:id/default", handler.AddressHandler.SetDefault)
	r.GET("/:id", handler.AddressHandler.GetByID)
	r.DELETE("/:id", handler.AddressHandler.Delete)
}

func AdminBannerRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.Use(mw.BearerTokenAuth())
	r.Use(mw.RBAC("admin", "superadmin"))

	r.GET("", handler.BannerHandler.GetBannerList)
	r.GET("/:id", handler.BannerHandler.GetBannerByID)
	r.POST("", handler.BannerHandler.CreateBanner)
	r.PUT("/:id", handler.BannerHandler.UpdateBanner)
	r.DELETE("/:id", handler.BannerHandler.DeleteBanner)
}

func ReviewRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	// Public endpoints
	r.GET("/:product_id", handler.ReviewHandler.GetReviewsByProduct)
	r.GET("/stats", handler.ReviewHandler.GetReviewStats)
	r.GET("/details/:id", handler.ReviewHandler.GetReviewDetails)

	// Protected endpoints
	protected := r.Group("")
	protected.Use(mw.BearerTokenAuth())
	protected.POST("", handler.ReviewHandler.BatchCreateReview)
	protected.PUT("/:id", handler.ReviewHandler.UpdateReview)
	protected.DELETE("/:id", handler.ReviewHandler.DeleteReview)

	// Admin route
	admin := r.Group("")
	admin.Use(mw.BearerTokenAuth())
	admin.GET("", handler.ReviewHandler.GetAllReviews)
}

func DashboardRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.Use(mw.BearerTokenAuth())
	r.Use(mw.RBAC("admin", "superadmin"))
	r.GET("/categories", handler.CategoryHandler.GetAllCategories)
	r.GET("/categories/:id", handler.CategoryHandler.GetCategoryByID)
	r.POST("/categories", handler.CategoryHandler.CreateCategory)
	r.PUT("/categories/:id", handler.CategoryHandler.UpdateCategory)
	r.DELETE("/categories/:id", handler.CategoryHandler.DeleteCategory)
	r.GET("/products", handler.ProductHandler.ProductListDashboard)
	r.GET("/products/:id", handler.ProductHandler.GetByID)
	r.POST("/products", handler.ProductHandler.CreateProduct)
	r.PUT("/products/:id", handler.ProductHandler.UpdateProduct)
	r.PUT("/products/:id/publish", handler.ProductHandler.UpdatePublish)
	r.GET("/products/:id/sku", handler.ProductHandler.GetSKUsByProductID)
	r.PUT("/products/:id/sku", handler.ProductHandler.BatchUpdateSKU)
	r.POST("/products/:id", handler.ProductHandler.UploadMainImage)
	r.POST("/products/:id/gallery", handler.ProductHandler.UploadProductGallery)
	r.DELETE("/products/:id/gallery/:imageID", handler.ProductHandler.DeleteImage)
	r.DELETE("/products/:id", handler.ProductHandler.DeleteProduct)
	r.GET("/orders", handler.OrderHandler.GetAll)
	r.GET("/orders/:id", handler.OrderHandler.GetByID)
	r.PUT("/orders/:id/confirm", handler.OrderHandler.Confirm)
	r.PUT("/orders/:id/complete", handler.OrderHandler.Complete)
	r.PUT("/orders/:id/cancel", handler.OrderHandler.Cancel)
	r.GET("/inventories", handler.InventoryHandler.GetInventory)
	r.GET("/inventories/:id", handler.InventoryHandler.GetInventoryBySKUID)
	r.GET("/inventories/logs", handler.InventoryHandler.GetInventoryLogs)
	r.GET("/inventories/logs/:id", handler.InventoryHandler.GetInventoryLogsBySKUID)
	r.POST("/inventories", handler.InventoryHandler.CreateInventoryLog)
	r.GET("/promotions", handler.PromotionHandler.GetPromotionList)
	r.GET("/promotions/:id", handler.PromotionHandler.GetPromotionByID)
	r.POST("/promotions", handler.PromotionHandler.CreatePromotion)
	r.PUT("/promotions/:id", handler.PromotionHandler.UpdatePromotion)
	r.DELETE("/promotions/:id", handler.PromotionHandler.DeletePromotion)
	r.GET("/report/sales", handler.ReportHandler.GetSales)
	r.GET("/report/revenue", handler.ReportHandler.GetRevenue)
	r.GET("/report/product", handler.ReportHandler.GetProductPerformance)
	r.GET("/report/customer/loyal", handler.ReportHandler.GetLoyalCustomer)
	r.GET("/report/customer", handler.ReportHandler.GetCustomerSummary)
}

func WishlistRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.Use(mw.BearerTokenAuth())
	r.GET("", handler.WishlistHandler.GetWishlist)
	r.POST("", handler.WishlistHandler.AddWishlist)
	r.DELETE("/:id", handler.WishlistHandler.RemoveWishlist)
}

func CompanyRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.GET("", handler.CompanyHandler.GetCompany)
	r.GET("/shipping-origin", handler.CompanyHandler.GetShippingOrigin)
}

func ContactRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.POST("", handler.CompanyHandler.SendMessage)
}

func AdminCompanyRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.Use(mw.BearerTokenAuth())
	r.Use(mw.RBAC("admin", "superadmin"))

	r.POST("", handler.CompanyHandler.CreateCompany)
	r.PUT("", handler.CompanyHandler.UpdateCompany)
	r.POST("/addresses", handler.CompanyHandler.CreateAddress)
	r.GET("/addresses/:id", handler.CompanyHandler.GetAddressByID)
	r.PUT("/addresses/:id", handler.CompanyHandler.UpdateAddress)
}

func CampaignRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.GET("/:id", handler.CampaignHandler.GetByID)
}

func PaymentRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.GET("", handler.PaymentMethodHandler.GetPaymentMethods)
	r.GET("/:id", handler.PaymentMethodHandler.GetPaymentMethodByID)
}

func AdminCampaignRoute(r *gin.RouterGroup, handler *adaptor.Handler, mw mCustom.MiddlewareCustom) {
	r.GET("", handler.CampaignHandler.GetAll)
	r.POST("", handler.CampaignHandler.Create)
	r.PUT("/:id", handler.CampaignHandler.Update)
}
