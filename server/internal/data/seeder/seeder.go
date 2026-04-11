package seeder

import (
	"gorm.io/gorm"
)

type Seeder struct {
	UserSeeder *UserSeeder
	CustomerSeeder *CustomerSeeder
	EmployeeSeeder *EmployeeSeeder
	CategorySeeder *CategorySeeder
	ProductSeeder *ProductSeeder
	CartDetailSeeder *CartDetailSeeder
	BankPaymentMethodSeeder *BankPaymentMethodSeeder
	PromotionSeeder *PromotionSeeder
	OrderSeeder *OrderSeeder
	ReviewSeeder *ReviewSeeder
	BannerSeeder *BannerSeeder
	AgripaculProductSeeder *AgripaculProductSeeder
}

func NewSeeder(db *gorm.DB) Seeder {
	return Seeder{
		UserSeeder : NewUserSeeder(db),
		CustomerSeeder : NewCustomerSeeder(db),
		EmployeeSeeder : NewEmployeeSeeder(db),
		CategorySeeder : NewCategorySeeder(db),
		ProductSeeder : NewProductSeeder(db),
		CartDetailSeeder : NewCartDetailSeeder(db),
		BankPaymentMethodSeeder : NewBankPaymentMethodSeeder(db),
		PromotionSeeder : NewPromotionSeeder(db),
		OrderSeeder : NewOrderSeeder(db),
		ReviewSeeder : NewReviewSeeder(db),
		BannerSeeder : NewBannerSeeder(db),
		AgripaculProductSeeder: NewAgripaculProductSeeder(db),
	}
}

func SeedAll(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		s := NewSeeder(tx)

		if err := s.UserSeeder.Run(); err != nil {
				return err
		}
		if err := s.CustomerSeeder.Run(); err != nil {
				return err
		}
		if err := s.EmployeeSeeder.Run(); err != nil {
				return err
		}
		if err := s.AgripaculProductSeeder.Seed(); err != nil {
				return err
		}
		
		if err := s.BankPaymentMethodSeeder.Run(); err != nil {
				return err
		}

		return nil
})
}

