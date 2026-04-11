package seeder

import (
	"time"

	"debian-ecommerce/internal/data/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserSeeder struct {
	DB *gorm.DB
}

func NewUserSeeder(db *gorm.DB) *UserSeeder {
	return &UserSeeder{DB: db}
}

func (s *UserSeeder) Run() error {
	var users []entity.User
	if err := s.DB.Find(&users).Error; err != nil {
		return err
	}

	if len(users) > 0 {
		return nil // Sudah ada user, skip seeding
	}

	users = []entity.User{
		// Super Admin
		{
			Email:             "superadmin@example.com",
			PasswordHash:      hashPassword("SuperAdmin123!"),
			Role:              entity.RoleSuperAdmin,
			PasswordChangedAt: time.Now(),
			IsActive:          true,
			EmailVerifiedAt:   time.Now(),
		},
		// Admin Users
		{
			Email:             "admin1@example.com",
			PasswordHash:      hashPassword("Admin123!"),
			Role:              entity.RoleAdmin,
			PasswordChangedAt: time.Now().Add(-24 * time.Hour),
			IsActive:          true,
			EmailVerifiedAt:   time.Now().Add(-48 * time.Hour),
		},
		{
			Email:             "admin2@example.com",
			PasswordHash:      hashPassword("Admin456!"),
			Role:              entity.RoleAdmin,
			PasswordChangedAt: time.Now().Add(-72 * time.Hour),
			IsActive:          true,
			EmailVerifiedAt:   time.Now().Add(-96 * time.Hour),
		},
		// Staff Users
		{
			Email:             "staff1@example.com",
			PasswordHash:      hashPassword("Staff123!"),
			Role:              entity.RoleStaff,
			PasswordChangedAt: time.Now().Add(-7 * 24 * time.Hour),
			IsActive:          true,
			EmailVerifiedAt:   time.Now().Add(-14 * 24 * time.Hour),
		},
		{
			Email:             "staff2@example.com",
			PasswordHash:      hashPassword("Staff456!"),
			Role:              entity.RoleStaff,
			PasswordChangedAt: time.Now().Add(-14 * 24 * time.Hour),
			IsActive:          true,
			EmailVerifiedAt:   time.Now().Add(-21 * 24 * time.Hour),
		},
		{
			Email:             "staff.inactive@example.com",
			PasswordHash:      hashPassword("Staff789!"),
			Role:              entity.RoleStaff,
			PasswordChangedAt: time.Now().Add(-30 * 24 * time.Hour),
			IsActive:          false,
			EmailVerifiedAt:   time.Now().Add(-60 * 24 * time.Hour),
		},
		// Customer Users
		{
			Email:             "customer1@example.com",
			PasswordHash:      hashPassword("Customer123!"),
			Role:              entity.RoleCustomer,
			PasswordChangedAt: time.Now().Add(-1 * 24 * time.Hour),
			IsActive:          true,
			EmailVerifiedAt:   time.Now().Add(-2 * 24 * time.Hour),
		},
		{
			Email:             "customer2@example.com",
			PasswordHash:      hashPassword("Customer456!"),
			Role:              entity.RoleCustomer,
			PasswordChangedAt: time.Now().Add(-15 * 24 * time.Hour),
			IsActive:          true,
			EmailVerifiedAt:   time.Now().Add(-30 * 24 * time.Hour),
		},
		{
			Email:             "customer3@example.com",
			PasswordHash:      hashPassword("Customer789!"),
			Role:              entity.RoleCustomer,
			PasswordChangedAt: time.Now().Add(-90 * 24 * time.Hour),
			IsActive:          true,
			EmailVerifiedAt:   time.Now().Add(-180 * 24 * time.Hour),
		},
		{
			Email:             "customer.unverified@example.com",
			PasswordHash:      hashPassword("Customer000!"),
			Role:              entity.RoleCustomer,
			PasswordChangedAt: time.Now(),
			IsActive:          true,
			// EmailVerifiedAt kosong (zero value) untuk user belum verifikasi
		},
		{
			Email:             "customer.inactive@example.com",
			PasswordHash:      hashPassword("Customer111!"),
			Role:              entity.RoleCustomer,
			PasswordChangedAt: time.Now().Add(-120 * 24 * time.Hour),
			IsActive:          false,
			EmailVerifiedAt:   time.Now().Add(-240 * 24 * time.Hour),
		},
	}

	// Hapus data lama jika ada (optional)
	s.DB.Exec("DELETE FROM users")

	// Create users
	for i := range users {
		if err := s.DB.Create(&users[i]).Error; err != nil {
			return err
		}
	}

	return nil
}

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err) // Atau handle error sesuai kebutuhan
	}
	return string(hash)
}
