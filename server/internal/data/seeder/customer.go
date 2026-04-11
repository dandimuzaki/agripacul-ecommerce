package seeder

import (
	"time"

	"debian-ecommerce/internal/data/entity"

	"gorm.io/gorm"
)

type CustomerSeeder struct {
	DB *gorm.DB
}

func NewCustomerSeeder(db *gorm.DB) *CustomerSeeder {
	return &CustomerSeeder{DB: db}
}

func (s *CustomerSeeder) Run() error {
	var customers []entity.Customer
	if err := s.DB.Find(&customers).Error; err != nil {
		return err
	}

	if len(customers) > 0 {
		return nil // Sudah ada customers, skip seeding
	}

	// Cari user dengan role customer
	var customerUsers []entity.User
	if err := s.DB.Where("role = ?", entity.RoleCustomer).Find(&customerUsers).Error; err != nil {
		return err
	}

	// Jika tidak ada user customer, buat dulu atau return
	if len(customerUsers) == 0 {
		return nil // Atau buat user customer terlebih dahulu
	}

	// Data customer dengan referensi ke user yang sudah ada
	dob1 := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
	dob2 := time.Date(1985, 8, 22, 0, 0, 0, 0, time.UTC)
	dob3 := time.Date(1995, 3, 10, 0, 0, 0, 0, time.UTC)
	dob4 := time.Date(1978, 11, 30, 0, 0, 0, 0, time.UTC)

	customers = []entity.Customer{
		{
			UserID:          customerUsers[0].ID, // customer1@example.com
			FullName:        "John Doe",
			PhoneNumber:           "+6281234567890",
			DateOfBirth:     &dob1,
			ProfileImageURL: "https://example.com/profiles/john_doe.jpg",
		},
		{
			UserID:          customerUsers[1].ID, // customer2@example.com
			FullName:        "Jane Smith",
			PhoneNumber:           "+6289876543210",
			DateOfBirth:     &dob2,
			ProfileImageURL: "https://example.com/profiles/jane_smith.jpg",
		},
		{
			UserID:          customerUsers[2].ID, // customer3@example.com
			FullName:        "Robert Johnson",
			PhoneNumber:           "+6281122334455",
			DateOfBirth:     &dob3,
			ProfileImageURL: "",
		},
		{
			UserID:          customerUsers[3].ID, // customer.unverified@example.com
			FullName:        "Michael Brown",
			PhoneNumber:           "+6285566778899",
			DateOfBirth:     &dob4,
			ProfileImageURL: "https://example.com/profiles/michael_brown.png",
		},
		{
			UserID:          customerUsers[4].ID, // customer.inactive@example.com
			FullName:        "Sarah Wilson",
			PhoneNumber:           "+6289988776655",
			DateOfBirth:     nil, // Tidak mengisi tanggal lahir
			ProfileImageURL: "",
		},
	}

	// Hapus data lama jika ada (optional)
	s.DB.Exec("DELETE FROM customers")

	// Create customers
	for i := range customers {
		if err := s.DB.Create(&customers[i]).Error; err != nil {
			return err
		}
	}

	return nil
}
