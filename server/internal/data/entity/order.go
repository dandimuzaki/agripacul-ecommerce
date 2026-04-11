package entity

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderStatusCreated   OrderStatus = "created"
	OrderStatusProcess OrderStatus = "processing"
	OrderStatusCompleted OrderStatus = "completed"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusPaid PaymentStatus = "paid"
	PaymentStatusExpired PaymentStatus = "expired"
)

type ShippingStatus string

const (
	ShippingStatusPending ShippingStatus = "pending"
	ShippingStatusShipped ShippingStatus = "shipped"
	ShippingStatusDelivered ShippingStatus = "delivered"
)

type Order struct {
	gorm.Model

	CustomerID uint        `gorm:"index;not null" json:"customer_id"`
	Status     OrderStatus `gorm:"type:varchar(50);not null" json:"status"`

	// Promotion
	PromotionID       *uint          `gorm:"index" json:"promotion_id,omitempty"`
	PromotionSnapshot datatypes.JSON `gorm:"type:jsonb" json:"promotion_snapshot,omitempty"`

	// Pricing snapshot
	Subtotal       float64 `gorm:"type:numeric(12,2);not null" json:"subtotal"`
	DiscountAmount float64 `gorm:"type:numeric(12,2);not null;default:0" json:"discount_amount"`
	Total           float64 `gorm:"type:numeric(12,2);not null" json:"total"`

	// Lifecycle timestamps
	CancelledAt *time.Time `json:"cancelled_at,omitempty"`
	CancelReason string `gorm:"type:text" json:"cancel_reason,omitempty"`
	CancelledBy  string `gorm:"type:varchar(50)" json:"cancelled_by,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	ConfirmedAt *time.Time `json:"confirmed_at,omitempty"`

	Notes string `gorm:"type:text" json:"notes,omitempty"`

	// Relations
	Customer   Customer    `gorm:"foreignKey:CustomerID;constraint:OnDelete:RESTRICT"`
	Items []OrderItem `gorm:"foreignKey:OrderID"`
	Promotion  Promotion  `gorm:"foreignKey:PromotionID;constraint:OnDelete:SET NULL"`
	Payment OrderPayment `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	Shipping OrderShipping `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
}

type OrderItem struct {
	gorm.Model

	OrderID uint `gorm:"not null"`
	SKUID   uint `gorm:"column:sku_id;not null"`

	Quantity   int     `gorm:"not null"`
	UnitPrice  float64 `gorm:"type:numeric(12,2);not null"`
	TotalPrice float64 `gorm:"type:numeric(12,2);not null"`

	ProductSnapshot datatypes.JSON `gorm:"type:jsonb;not null"`
	SKUSnapshot     datatypes.JSON `gorm:"type:jsonb;not null"`
	IsRated bool `gorm:"default:false"`

	SKU SKU `gorm:"foreignKey:SKUID;constraint:OnDelete:CASCADE"`
}

type OrderPayment struct {
	gorm.Model

	OrderID uint `gorm:"uniqueIndex;not null" json:"order_id"`
	PaymentMethodID   uint `gorm:"not null" json:"payment_method_id"`   // bank_transfer, cod, gateway
	Amount float64 `gorm:"type:numeric(12,2);not null" json:"amount"`
	Status PaymentStatus `gorm:"type:varchar(50);not null" json:"status"`
	TransactionRef string   `gorm:"type:varchar(100);index" json:"transaction_ref,omitempty"`
	PaidAt         *time.Time `json:"paid_at,omitempty"`
	FailedAt       *time.Time `json:"failed_at,omitempty"`
	ExpiredAt time.Time `json:"expired_at"`

	// Relations
	PaymentMethod PaymentMethod `gorm:"foreignKey:PaymentMethodID"`
}

type OrderShipping struct {
	gorm.Model

	OrderID uint `gorm:"uniqueIndex;not null" json:"order_id"`

	// Address snapshot
	RecipientName string `gorm:"type:varchar(100);not null"`
	Label string `gorm:"type:varchar(50)"`
	PhoneNumber   string `gorm:"type:varchar(30);not null"`
	DetailAddress   string `gorm:"type:text"`
	Province      string `gorm:"type:varchar(100);not null"`
	Regency          string `gorm:"type:varchar(100);not null"`
	District      string `gorm:"type:varchar(100);not null"`
	Subdistrict      string `gorm:"type:varchar(100);not null"`
	PostalCode    string `gorm:"type:varchar(10)"`

	// Courier snapshot
	CourierName string  `gorm:"type:varchar(100);not null"`
	CourierCode string  `gorm:"type:varchar(20);not null"`
	CourierService string  `gorm:"type:varchar(50)"`
	Cost        float64 `gorm:"type:numeric(12,2);not null"`
	ETD         int  `json:"etd"`

	// Shipping lifecycle
	Status         ShippingStatus `gorm:"type:varchar(50);not null"`
	TrackingNumber string        `gorm:"type:varchar(100);index"`
	ShippedAt      *time.Time
	DeliveredAt    *time.Time
}
