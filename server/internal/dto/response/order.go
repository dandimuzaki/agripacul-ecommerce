package response

import (
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/pkg/utils"
	"time"
)

type OrderDetails struct {
	ID             uint           `json:"id"`
	Customer Customer `json:"customer"`
	DisplayStatus  string         `json:"display_status"`
	Steps interface{} `json:"steps"`
	Cancellation interface{} `json:"cancellation"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	BillingAddress BillingAddress `json:"billing_address"`
	Items []OrderItem `json:"items"`
	Shipping ShippingSnapshot `json:"shipping"`
	Totals Totals `json:"totals"`
	Notes string `json:"notes"`
	TrackingNumber string `json:"tracking_number,omitempty"`
	CanReview bool `json:"can_review"`
}

type BillingAddress struct {
	RecipientName string `json:"recipient_name"`
	Label string `json:"label"`
	PhoneNumber string `json:"phone_number"`
	Province string `json:"province"`
	Regency string `json:"regency"`
	District string `json:"district"`
	Subdistrict string `json:"subdistrict"`
	PostalCode string `json:"postal_code"`
	DetailAddress string `json:"detail_address"`
}

type OrderItem struct {
	ID uint `json:"id"`
	SKUID uint `json:"sku_id"`
	Name string `json:"name"`
	MainImageURL string `json:"main_image_url"`
	Variants []entity.VariantCombination `json:"variants"`
	Quantity int `json:"quantity"`
	Price float64 `json:"price"`
	TotalPrice float64 `json:"total_price"`
}

type Totals struct {
	PaymentMethod string `json:"payment_method"`
	Subtotal float64 `json:"subtotal"`
	DiscountAmount *float64 `json:"discount_amount,omitempty"`
	ShippingCost float64 `json:"shipping_cost"`
	GrandTotal float64 `json:"grand_total"`
}

type ShippingSnapshot struct {
	Name    string  `json:"name"`
	Code    string  `json:"code"`
	Service string  `json:"service"`
	Cost    float64 `json:"cost"`
	ETD     int  `json:"etd"`
}

func ToDetails(o entity.Order) (*OrderDetails, error) {
	address := BillingAddress{
		RecipientName: o.Shipping.RecipientName,
		Label: o.Shipping.Label,
		PhoneNumber: o.Shipping.PhoneNumber,
		Province: o.Shipping.Province,
		Regency: o.Shipping.Regency,
		District: o.Shipping.District,
		Subdistrict: o.Shipping.Subdistrict,
		PostalCode: o.Shipping.PostalCode,
		DetailAddress: o.Shipping.DetailAddress,
	}

	var items []OrderItem
	for _, i := range o.Items {
		var product entity.ProductSnapshot
		if err := utils.FromJSONB(i.ProductSnapshot, &product); err != nil {
			return nil, err
		}
		var sku entity.SKUSnapshot
		if err := utils.FromJSONB(i.SKUSnapshot, &sku); err != nil {
			return nil, err
		}
		item := OrderItem{
			ID: i.ID,
			SKUID: i.SKUID,
			Name: product.Name,
			MainImageURL: product.MainImageURL,
			Variants: sku.Variants,
			Quantity: i.Quantity,
			Price: i.UnitPrice,
			TotalPrice: i.TotalPrice,
		}
		items = append(items, item)
	}

	shipping := ShippingSnapshot{
		Name: o.Shipping.CourierName,
		Code: o.Shipping.CourierCode,
		Service: o.Shipping.CourierService,
		Cost: o.Shipping.Cost,
		ETD: o.Shipping.ETD,
	}

	totals := Totals{
		PaymentMethod: o.Payment.PaymentMethod.Name,
		Subtotal: o.Subtotal,
		DiscountAmount: &o.DiscountAmount,
		ShippingCost: shipping.Cost,
		GrandTotal: o.Total,
	}

	trackingNumber := o.Shipping.TrackingNumber
	order := OrderDetails{
		ID: o.ID,
		Customer: Customer{
			Name: o.Customer.FullName,
			Email: o.Customer.User.Email,
		},
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
		BillingAddress: address,
		Items: items,
		Shipping: shipping,
		Totals: totals,
		Notes: o.Notes,
		TrackingNumber: trackingNumber,
	}

	return &order, nil
}

type OrderSummary struct {
	ID uint `json:"id"`
	Customer Customer `json:"customer"`
	CreatedAt time.Time `json:"created_at"`
	DisplayStatus  string         `json:"display_status"`
	GrandTotal float64 `json:"grand_total"`
	FirstItem OrderItem `json:"first_item"`
	ItemCount int `json:"item_count"`
	Cancellation interface{} `json:"cancellation"`
	TrackingNumber *string `json:"tracking_number,omitempty"`
	ShippingCourier string `json:"shipping_courier"`
	ETD int `json:"etd"`
	PaymentMethod string `json:"payment_method"`
	CanReview bool `json:"can_review"`
}

type Customer struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

func ToSummary(o entity.Order) (*OrderSummary, error) {
	customer := Customer{
		Name: o.Customer.FullName,
		Email: o.Customer.User.Email,
	}

	var product entity.ProductSnapshot
	var sku entity.SKUSnapshot
	if len(o.Items) > 0 {
		if err := utils.FromJSONB(o.Items[0].ProductSnapshot, &product); err != nil {
			return nil, err
		}
		
		if err := utils.FromJSONB(o.Items[0].SKUSnapshot, &sku); err != nil {
			return nil, err
		}
	}
	
	item := OrderItem{
		Name: product.Name,
		MainImageURL: product.MainImageURL,
		Variants: sku.Variants,
		Quantity: o.Items[0].Quantity,
		Price: o.Items[0].UnitPrice,
		TotalPrice: o.Items[0].TotalPrice,
	}

	trackingNumber := o.Shipping.TrackingNumber

	order := OrderSummary{
		ID: o.ID,
		Customer: customer,
		CreatedAt: o.CreatedAt,
		GrandTotal: o.Total,
		FirstItem: item,
		ItemCount: len(o.Items),
		TrackingNumber: &trackingNumber,
		ShippingCourier: o.Shipping.CourierName,
		ETD: o.Shipping.ETD,
		PaymentMethod: o.Payment.PaymentMethod.Name,
	}

	return &order, nil
}