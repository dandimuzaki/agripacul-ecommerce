package seeder

import (
	"debian-ecommerce/internal/data/entity"

	"gorm.io/gorm"
)

type BankPaymentMethodSeeder struct {
	DB *gorm.DB
}

func NewBankPaymentMethodSeeder(db *gorm.DB) *BankPaymentMethodSeeder {
	return &BankPaymentMethodSeeder{DB: db}
}

func (s *BankPaymentMethodSeeder) Run() error {
	var paymentMethods []entity.PaymentMethod
	if err := s.DB.Find(&paymentMethods).Error; err != nil {
		return err
	}

	if len(paymentMethods) > 0 {
		return nil // Sudah ada paymentMethods, skip seeding
	}

	// Data bank-bank di Indonesia
	paymentMethods = []entity.PaymentMethod{
		// Bank Transfer Methods
		{Name: "Bank Transfer - BCA", IsActive: true},
		{Name: "Bank Transfer - Mandiri", IsActive: true},
		{Name: "Bank Transfer - BNI", IsActive: true},
		{Name: "Bank Transfer - BRI", IsActive: true},
		{Name: "Bank Transfer - CIMB Niaga", IsActive: true},
		{Name: "Bank Transfer - Bank Permata", IsActive: true},
		{Name: "Bank Transfer - Bank Danamon", IsActive: true},
		{Name: "Bank Transfer - Bank Mega", IsActive: true},
		{Name: "Bank Transfer - BTPN/Jenius", IsActive: true},
		{Name: "Bank Transfer - Bank Syariah Indonesia", IsActive: true},

		// Virtual Account Methods
		{Name: "BCA Virtual Account", IsActive: true},
		{Name: "Mandiri Virtual Account", IsActive: true},
		{Name: "BNI Virtual Account", IsActive: true},
		{Name: "BRI Virtual Account", IsActive: true},
		{Name: "Permata Virtual Account", IsActive: true},
		{Name: "CIMB Virtual Account", IsActive: true},

		// E-Wallets
		{Name: "OVO", IsActive: true},
		{Name: "GoPay", IsActive: true},
		{Name: "DANA", IsActive: true},
		{Name: "LinkAja", IsActive: true},
		{Name: "ShopeePay", IsActive: true},

		// Credit/Debit Cards
		{Name: "Visa", IsActive: true},
		{Name: "MasterCard", IsActive: true},
		{Name: "JCB", IsActive: true},
		{Name: "American Express", IsActive: true},

		// Retail Payments
		{Name: "Alfamart", IsActive: true},
		{Name: "Indomaret", IsActive: true},
		{Name: "Alfamidi", IsActive: true},
		{Name: "Lawson", IsActive: true},

		// Internet Banking
		{Name: "BCA KlikPay", IsActive: true},
		{Name: "Mandiri e-Cash", IsActive: true},
		{Name: "BNI e-Banking", IsActive: true},
		{Name: "BRI e-Pay", IsActive: true},

		// Payment Gateways
		// {Name: "Midtrans", IsActive: true},
		// {Name: "Xendit", IsActive: true},
		// {Name: "Doku", IsActive: true},
		// {Name: "iPay88", IsActive: true},

		// Others
		{Name: "COD (Bayar di Tempat)", IsActive: true},
		// {Name: "QRIS", IsActive: true},
		// {Name: "PayLater - Akulaku", IsActive: true},
		// {Name: "PayLater - Kredivo", IsActive: true},
		// {Name: "PayLater - Atome", IsActive: true},
		// {Name: "PayPal", IsActive: false},        // Biasanya tidak aktif di Indonesia
		// {Name: "Western Union", IsActive: false}, // Untuk international
	}

	// Batch create dengan OnConflict
	for i := range paymentMethods {
		// Use Upsert (Create or Update)
		if err := s.DB.Where(entity.PaymentMethod{Name: paymentMethods[i].Name}).
			Assign(entity.PaymentMethod{IsActive: paymentMethods[i].IsActive}).
			FirstOrCreate(&paymentMethods[i]).Error; err != nil {
			return err
		}
	}

	return nil
}
