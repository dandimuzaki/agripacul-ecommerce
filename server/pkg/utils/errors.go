package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// =============================================
// ERROR CONSTANTS
// =============================================

// Authentication & Authorization errors
var (
	ErrUnauthorized       = errors.New("unauthorized access")
	ErrInvalidToken       = errors.New("invalid token")
	ErrTokenExpired       = errors.New("token expired")
	ErrForbidden          = errors.New("access forbidden")
	ErrInsufficientRole   = errors.New("insufficient role privileges")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrAccountLocked      = errors.New("account is locked")
	ErrEmailNotVerified   = errors.New("email not verified")
	ErrInvalidOTP         = errors.New("invalid OTP")
	ErrOTPExpired         = errors.New("OTP expired")
)

// User errors
var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserInactive      = errors.New("user account is inactive")
	ErrInvalidUserID     = errors.New("invalid user ID")
	ErrPasswordMismatch  = errors.New("password mismatch")
	ErrWeakPassword      = errors.New("password is too weak")
	ErrEmailAlreadyUsed  = errors.New("email already registered")
)

// Customer errors
var (
	ErrCustomerNotFound    = errors.New("customer not found")
	ErrCustomerInactive    = errors.New("customer account is inactive")
	ErrInvalidCustomerData = errors.New("invalid customer data")
)

// Employee errors
var (
	ErrEmployeeNotFound = errors.New("employee not found")
	ErrEmployeeInactive = errors.New("employee account is inactive")
)

// Category errors
var (
	ErrCategoryNotFound  = errors.New("category not found")
	ErrCategoryExists    = errors.New("category already exists")
	ErrInvalidCategoryID = errors.New("invalid category ID")
	ErrCategoryNotEmpty  = errors.New("category has associated products")
)

// Product errors
var (
	ErrProductNotFound      = errors.New("product not found")
	ErrProductUnpublished   = errors.New("product is not published")
	ErrInvalidProductData   = errors.New("invalid product data")
	ErrSKUNotFound          = errors.New("SKU not found")
	ErrSKUInactive          = errors.New("SKU is inactive")
	ErrInsufficientStock    = errors.New("insufficient stock")
	ErrVariantTypeNotFound  = errors.New("variant type not found")
	ErrVariantValueNotFound = errors.New("variant value not found")
	ErrImageNotFound        = errors.New("image not found")
)

// Cart errors
var (
	ErrCartNotFound     = errors.New("cart not found")
	ErrCartItemNotFound = errors.New("cart item not found")
	ErrCartEmpty        = errors.New("cart is empty")
)

// Order errors
var (
	ErrOrderNotFound                = errors.New("order not found")
	ErrInvalidOrderStatus           = errors.New("invalid order status")
	ErrOrderCancelled               = errors.New("order is cancelled")
	ErrOrderCompleted               = errors.New("order is already completed")
	ErrInvalidOrderData             = errors.New("invalid order data")
	ErrOrderAlreadyPaid             = errors.New("order is already paid")
	ErrOrderUnpaid                  = errors.New("order is unpaid")
	ErrInvalidOrderStatusTransition = errors.New("invalid order status transition")
)

// Payment errors
var (
	ErrPaymentNotFound        = errors.New("payment not found")
	ErrPaymentMethodNotFound  = errors.New("payment method not found")
	ErrPaymentFailed          = errors.New("payment failed")
	ErrPaymentExpired         = errors.New("payment expired")
	ErrInvalidPaymentAmount   = errors.New("invalid payment amount")
	ErrPaymentMethodInactive  = errors.New("payment method is inactive")
	ErrNameRequired           = errors.New("name is required")
	ErrDuplicatePaymentMethod = errors.New("duplicate payment method")
)

// Promotion errors
var (
	ErrPromotionNotFound      = errors.New("promotion not found")
	ErrPromotionExpired       = errors.New("promotion has expired")
	ErrPromotionInactive      = errors.New("promotion is not active")
	ErrPromotionUsageLimit    = errors.New("promotion usage limit reached")
	ErrInvalidPromoCode       = errors.New("invalid promotion code")
	ErrMinimumOrderNotMet     = errors.New("minimum order price not met")
	ErrPromotionNotApplicable = errors.New("promotion not applicable to selected products")
)

// Address errors
var (
	ErrAddressNotFound      = errors.New("address not found")
	ErrInvalidAddressData   = errors.New("invalid address data")
	ErrPrimaryAddressExists = errors.New("primary address already exists")
)

// Review errors
var (
	ErrReviewNotFound      = errors.New("review not found")
	ErrReviewAlreadyExists = errors.New("review already exists for this order")
	ErrInvalidRating       = errors.New("invalid rating value")
)

// Banner errors
var (
	ErrBannerNotFound    = errors.New("banner not found")
	ErrBannerInactive    = errors.New("banner is not active")
	ErrInvalidBannerData = errors.New("invalid banner data")
)

// Campaign errors
var (
	ErrCampaignNotFound = errors.New("campaign not found")
)

// Validation errors
var (
	ErrValidationFailed = errors.New("validation failed")
	ErrRequiredField    = errors.New("field is required")
	ErrInvalidEmail     = errors.New("invalid email format")
	ErrInvalidPhone     = errors.New("invalid phone number format")
	ErrInvalidURL       = errors.New("invalid URL format")
	ErrInvalidDate      = errors.New("invalid date format")
	ErrInvalidPrice     = errors.New("invalid price format")
	ErrInvalidQuantity  = errors.New("invalid quantity")
	ErrInsufficientQuantity  = errors.New("quantity must be greater than 0")
	ErrFileTooLarge     = errors.New("file size too large")
	ErrInvalidFileType  = errors.New("invalid file type")
)

// Database errors
var (
	ErrDBConnection     = errors.New("database connection failed")
	ErrDBTransaction    = errors.New("database transaction failed")
	ErrDBConstraint     = errors.New("database constraint violation")
	ErrDBDuplicate      = errors.New("duplicate entry")
	ErrDBRecordNotFound = errors.New("record not found in database")
)

// Business logic errors
var (
	ErrOutOfStock          = errors.New("product is out of stock")
	ErrPriceChanged        = errors.New("product price has changed")
	ErrShippingUnavailable = errors.New("shipping unavailable to this address")
	ErrCourierError        = errors.New("courier service error")
	ErrInvalidTracking     = errors.New("invalid tracking number")
	ErrRefundFailed        = errors.New("refund failed")
	ErrExchangeNotAllowed  = errors.New("exchange not allowed for this product")
)

// File & Storage errors
var (
	ErrFileUploadFailed = errors.New("file upload failed")
	ErrFileNotFound     = errors.New("file not found")
	ErrStorageError     = errors.New("storage error")
	ErrImageProcessing  = errors.New("image processing error")
)

// External service errors
var (
	ErrExternalAPI        = errors.New("external API error")
	ErrPaymentGateway     = errors.New("payment gateway error")
	ErrShippingAPI        = errors.New("shipping API error")
	ErrNotificationFailed = errors.New("notification failed to send")
	ErrSMSSendFailed      = errors.New("SMS sending failed")
	ErrEmailSendFailed    = errors.New("email sending failed")
)

// System errors
var (
	ErrInternalServer     = errors.New("internal server error")
	ErrServiceUnavailable = errors.New("service temporarily unavailable")
	ErrTimeout            = errors.New("request timeout")
	ErrRateLimitExceeded  = errors.New("rate limit exceeded")
	ErrMaintenance        = errors.New("system under maintenance")
)

// Input/Request errors
var (
	ErrBadRequest        = errors.New("bad request")
	ErrInvalidJSON       = errors.New("invalid JSON format")
	ErrMissingParameters = errors.New("missing required parameters")
	ErrInvalidPagination = errors.New("invalid pagination parameters")
	ErrInvalidFilter     = errors.New("invalid filter parameters")
	ErrInvalidSort       = errors.New("invalid sort parameters")
)

// =============================================
// RESPONSE STRUCTS & FUNCTIONS
// =============================================

// SuccessResponse untuk response sukses
type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse untuk response error
type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// SuccessJSON mengirim response success
func SuccessJSON(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorJSON mengirim response error
func ErrorJSON(c *gin.Context, statusCode int, message string, err error) {
	errorMsg := ""
	if err != nil {
		errorMsg = err.Error()
	}

	c.JSON(statusCode, ErrorResponse{
		Success: false,
		Message: message,
		Error:   errorMsg,
	})
}

// =============================================
// HELPER FUNCTIONS FOR COMMON HTTP STATUS CODES
// =============================================

// BadRequest mengirim 400 Bad Request
func BadRequest(c *gin.Context, message string, err error) {
	ErrorJSON(c, http.StatusBadRequest, message, err)
}

// Unauthorized mengirim 401 Unauthorized
func Unauthorized(c *gin.Context, message string, err error) {
	ErrorJSON(c, http.StatusUnauthorized, message, err)
}

// Forbidden mengirim 403 Forbidden
func Forbidden(c *gin.Context, message string, err error) {
	ErrorJSON(c, http.StatusForbidden, message, err)
}

// NotFound mengirim 404 Not Found
func NotFound(c *gin.Context, message string, err error) {
	ErrorJSON(c, http.StatusNotFound, message, err)
}

// Conflict mengirim 409 Conflict
func Conflict(c *gin.Context, message string, err error) {
	ErrorJSON(c, http.StatusConflict, message, err)
}

// InternalServerError mengirim 500 Internal Server Error
func InternalServerError(c *gin.Context, message string, err error) {
	ErrorJSON(c, http.StatusInternalServerError, message, err)
}

// ValidationError mengirim 422 Unprocessable Entity
func ValidationError(c *gin.Context, message string, err error) {
	ErrorJSON(c, http.StatusUnprocessableEntity, message, err)
}

// =============================================
// CONVENIENCE FUNCTIONS FOR SUCCESS RESPONSES
// =============================================

// Created mengirim 201 Created
func Created(c *gin.Context, message string, data interface{}) {
	SuccessJSON(c, http.StatusCreated, message, data)
}

// OK mengirim 200 OK
func OK(c *gin.Context, message string, data interface{}) {
	SuccessJSON(c, http.StatusOK, message, data)
}

// NoContent mengirim 204 No Content
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// =============================================
// CUSTOM ERROR TYPES & ERROR HELPERS
// =============================================

// AppError untuk error dengan code tambahan
type AppError struct {
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

func NewAppError(code, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// Common error codes
const (
	ErrCodeUnauthorized       = "UNAUTHORIZED"
	ErrCodeForbidden          = "FORBIDDEN"
	ErrCodeNotFound           = "NOT_FOUND"
	ErrCodeValidation         = "VALIDATION_ERROR"
	ErrCodeConflict           = "CONFLICT"
	ErrCodeBadRequest         = "BAD_REQUEST"
	ErrCodeInternal           = "INTERNAL_ERROR"
	ErrCodeServiceUnavailable = "SERVICE_UNAVAILABLE"
	ErrCodeRateLimit          = "RATE_LIMIT"
)

// Helper functions untuk mengecek jenis error
func IsNotFoundError(err error) bool {
	switch err {
	case ErrUserNotFound, ErrCustomerNotFound, ErrProductNotFound,
		ErrCategoryNotFound, ErrOrderNotFound, ErrPaymentNotFound,
		ErrPromotionNotFound, ErrAddressNotFound, ErrReviewNotFound,
		ErrBannerNotFound, ErrDBRecordNotFound, ErrCartNotFound, ErrCartItemNotFound:
		return true
	}
	return false
}

func IsValidationError(err error) bool {
	switch err {
	case ErrValidationFailed, ErrRequiredField, ErrInvalidEmail,
		ErrInvalidPhone, ErrInvalidURL, ErrInvalidDate, ErrInvalidPrice,
		ErrInvalidQuantity, ErrFileTooLarge, ErrInvalidFileType:
		return true
	}
	return false
}

func IsDuplicateError(err error) bool {
	return err == ErrDBDuplicate
}

func IsAuthError(err error) bool {
	switch err {
	case ErrUnauthorized, ErrInvalidToken, ErrTokenExpired,
		ErrForbidden, ErrInvalidCredentials, ErrAccountLocked:
		return true
	}
	return false
}

// WrapError wraps an error with additional context
func WrapError(msg string, err error) error {
	if err == nil {
		return errors.New(msg)
	}
	return errors.New(msg + ": " + err.Error())
}
