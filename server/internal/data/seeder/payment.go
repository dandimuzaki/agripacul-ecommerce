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
	var paymentTypes []entity.PaymentType
	if err := s.DB.Find(&paymentTypes).Error; err != nil {
		return err
	}

	if len(paymentTypes) > 0 {
		return nil // Sudah ada paymentTypes, skip seeding
	}

	paymentTypes = []entity.PaymentType{
		{Name: "Bank Transfer"},
		{Name: "Virtual Account"},
		{Name: "E-Wallet"},
		{Name: "Debit/Credit Card"},
		{Name: "Retail Payment"},
	}

	var bankTransfer entity.PaymentType
	var virtualAccount entity.PaymentType
	var eWallet entity.PaymentType
	var debitCredit entity.PaymentType
	var retail entity.PaymentType
	if err := s.DB.Where("name = ?", "Bank Transfer").First(&bankTransfer).Error; err != nil {
		return err
	}
	if err := s.DB.Where("name = ?", "Virtual Account").First(&virtualAccount).Error; err != nil {
		return err
	}
	if err := s.DB.Where("name = ?", "E-Wallet").First(&eWallet).Error; err != nil {
		return err
	}
	if err := s.DB.Where("name = ?", "Debit/Credit Card").First(&debitCredit).Error; err != nil {
		return err
	}
	if err := s.DB.Where("name = ?", "Retail Payment").First(&retail).Error; err != nil {
		return err
	}

	for i := range paymentTypes {
		// Use Upsert (Create or Update)
		if err := s.DB.Where(entity.PaymentType{Name: paymentTypes[i].Name}).
			FirstOrCreate(&paymentTypes[i]).Error; err != nil {
			return err
		}
	}

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
		{PaymentTypeID: bankTransfer.ID, Name: "BCA", IsActive: true, IconURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQyyDL48YXf0J3DjfcEKeyvCxbT9uJVND1kEQ&s"},
		{PaymentTypeID: bankTransfer.ID, Name: "Mandiri", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/a/ad/Bank_Mandiri_logo_2016.svg/1280px-Bank_Mandiri_logo_2016.svg.png"},
		{PaymentTypeID: bankTransfer.ID, Name: "BNI", IsActive: true, IconURL: `https://upload.wikimedia.org/wikipedia/commons/thumb/f/f0/Bank_Negara_Indonesia_logo_%282004%29.svg/3840px-Bank_Negara_Indonesia_logo_%282004%29.svg.png`},
		{PaymentTypeID: bankTransfer.ID, Name: "BRI", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/6/68/BANK_BRI_logo.svg/1280px-BANK_BRI_logo.svg.png"},
		{PaymentTypeID: bankTransfer.ID, Name: "CIMB Niaga", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/3/38/CIMB_Niaga_logo.svg"},
		{PaymentTypeID: bankTransfer.ID, Name: "Permata Bank", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/id/thumb/4/48/PermataBank_logo.svg/3840px-PermataBank_logo.svg.png"},
		{PaymentTypeID: bankTransfer.ID, Name: "Bank Danamon", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/a/a1/Danamon_%282024%29.svg"},
		{PaymentTypeID: bankTransfer.ID, Name: "Bank Mega", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/a/af/Bank_Mega_2013.svg/1280px-Bank_Mega_2013.svg.png"},
		{PaymentTypeID: bankTransfer.ID, Name: "Bank BTPN", IsActive: true, IconURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT7b7mFFLkz0G210Vdh4gwW-rZyjCxmwQDzmA&s"},
		{PaymentTypeID: bankTransfer.ID, Name: "Bank Syariah Indonesia", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a0/Bank_Syariah_Indonesia.svg/960px-Bank_Syariah_Indonesia.svg.png"},

		// Virtual Account Methods
		{PaymentTypeID: virtualAccount.ID, Name: "BCA Virtual Account", IsActive: true, IconURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQyyDL48YXf0J3DjfcEKeyvCxbT9uJVND1kEQ&s"},
		{PaymentTypeID: virtualAccount.ID, Name: "Mandiri Virtual Account", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/a/ad/Bank_Mandiri_logo_2016.svg/1280px-Bank_Mandiri_logo_2016.svg.png"},
		{PaymentTypeID: virtualAccount.ID, Name: "BNI Virtual Account", IsActive: true, IconURL: `https://upload.wikimedia.org/wikipedia/commons/thumb/f/f0/Bank_Negara_Indonesia_logo_%282004%29.svg/3840px-Bank_Negara_Indonesia_logo_%282004%29.svg.png`},
		{PaymentTypeID: virtualAccount.ID, Name: "BRI Virtual Account", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/6/68/BANK_BRI_logo.svg/1280px-BANK_BRI_logo.svg.png"},
		{PaymentTypeID: virtualAccount.ID, Name: "Permata Virtual Account", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/id/thumb/4/48/PermataBank_logo.svg/3840px-PermataBank_logo.svg.png"},
		{PaymentTypeID: virtualAccount.ID, Name: "CIMB Virtual Account", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/3/38/CIMB_Niaga_logo.svg"},
		
		// E-Wallets
		{PaymentTypeID: eWallet.ID, Name: "OVO", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/e/eb/Logo_ovo_purple.svg/3840px-Logo_ovo_purple.svg.png"},
		{PaymentTypeID: eWallet.ID, Name: "GoPay", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/8/86/Gopay_logo.svg/1280px-Gopay_logo.svg.png"},
		{PaymentTypeID: eWallet.ID, Name: "DANA", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/7/72/Logo_dana_blue.svg/1280px-Logo_dana_blue.svg.png"},
		{PaymentTypeID: eWallet.ID, Name: "LinkAja", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/8/85/LinkAja.svg/1280px-LinkAja.svg.png"},
		{PaymentTypeID: eWallet.ID, Name: "ShopeePay", IsActive: true, IconURL: "https://www.koleksilogo.com/2023/02/logo-shopeepay.html"},

		// Credit/Debit Cards
		{PaymentTypeID: debitCredit.ID, Name: "Visa", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/5/5c/Visa_Inc._logo_%282021%E2%80%93present%29.svg/3840px-Visa_Inc._logo_%282021%E2%80%93present%29.svg.png"},
		{PaymentTypeID: debitCredit.ID, Name: "MasterCard", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/b/b7/MasterCard_Logo.svg/1280px-MasterCard_Logo.svg.png"},
		{PaymentTypeID: debitCredit.ID, Name: "JCB", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/4/40/JCB_logo.svg/960px-JCB_logo.svg.png"},
		{PaymentTypeID: debitCredit.ID, Name: "American Express", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/f/fa/American_Express_logo_%282018%29.svg/1280px-American_Express_logo_%282018%29.svg.png"},

		// Retail Payments
		{PaymentTypeID: retail.ID, Name: "Alfamart", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/8/86/Alfamart_logo.svg"},
		{PaymentTypeID: retail.ID, Name: "Indomaret", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/9/9d/Logo_Indomaret.png"},
		{PaymentTypeID: retail.ID, Name: "Alfamidi", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/id/7/7f/Alfamidi.svg"},
		{PaymentTypeID: retail.ID, Name: "Lawson", IsActive: true, IconURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/b/b0/Lawson.svg/3840px-Lawson.svg.png"},

		// Internet Banking
		// {Name: "BCA KlikPay", IsActive: true, IconURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTuzI2bqGNUhwQlvV9Uledvvr7cHTADkHXVxQ&s"},
		// {Name: "Mandiri e-Cash", IsActive: true, IconURL: "https://rs-setohasbadi.com/wp-content/uploads/2018/07/mandiri-e-cach.png"},
		// {Name: "BNI e-Banking", IsActive: true, IconURL: ""},
		// {Name: "BRI e-Pay", IsActive: true, IconURL: ""},

		// Payment Gateways
		// {Name: "Midtrans", IsActive: true},
		// {Name: "Xendit", IsActive: true},
		// {Name: "Doku", IsActive: true},
		// {Name: "iPay88", IsActive: true},

		// Others
		// {Name: "COD (Bayar di Tempat)", IsActive: true},
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
