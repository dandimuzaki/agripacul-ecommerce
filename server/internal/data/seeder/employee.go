package seeder

import (
	"time"

	"debian-ecommerce/internal/data/entity"

	"gorm.io/gorm"
)

type EmployeeSeeder struct {
	DB *gorm.DB
}

func NewEmployeeSeeder(db *gorm.DB) *EmployeeSeeder {
	return &EmployeeSeeder{DB: db}
}

func (s *EmployeeSeeder) Run() error {
	var employees []entity.Employee
	if err := s.DB.Find(&employees).Error; err != nil {
		return err
	}

	if len(employees) > 0 {
		return nil // Sudah ada employees, skip seeding
	}

	// Cari user dengan role admin dan staff
	var adminStaffUsers []entity.User
	if err := s.DB.Where("role IN ?", []entity.UserRole{
		entity.RoleAdmin,
		entity.RoleStaff,
		entity.RoleSuperAdmin,
	}).Find(&adminStaffUsers).Error; err != nil {
		return err
	}

	// Jika tidak ada user yang sesuai, return
	if len(adminStaffUsers) == 0 {
		return nil
	}

	// Data employee dengan referensi ke user yang sudah ada
	dob1 := time.Date(1980, 12, 25, 0, 0, 0, 0, time.UTC)
	dob2 := time.Date(1985, 6, 18, 0, 0, 0, 0, time.UTC)
	dob3 := time.Date(1990, 3, 8, 0, 0, 0, 0, time.UTC)
	dob4 := time.Date(1992, 9, 14, 0, 0, 0, 0, time.UTC)
	dob5 := time.Date(1988, 7, 30, 0, 0, 0, 0, time.UTC)
	dob6 := time.Date(1995, 1, 20, 0, 0, 0, 0, time.UTC)

	employees = []entity.Employee{
		// Super Admin Employee
		{
			UserID:          findUserIDByEmail(adminStaffUsers, "superadmin@example.com"),
			FullName:        "Super Admin Manager",
			Phone:           "+6280000000001",
			DateOfBirth:     &dob1,
			Salary:          25000000.00,
			ProfileImageURL: "https://example.com/profiles/superadmin.jpg",
		},
		// Admin Employees
		{
			UserID:          findUserIDByEmail(adminStaffUsers, "admin1@example.com"),
			FullName:        "Admin One Manager",
			Phone:           "+6280000000002",
			DateOfBirth:     &dob2,
			Salary:          15000000.00,
			ProfileImageURL: "https://example.com/profiles/admin1.jpg",
		},
		{
			UserID:          findUserIDByEmail(adminStaffUsers, "admin2@example.com"),
			FullName:        "Admin Two Supervisor",
			Phone:           "+6280000000003",
			DateOfBirth:     &dob3,
			Salary:          12000000.00,
			ProfileImageURL: "https://example.com/profiles/admin2.jpg",
		},
		// Staff Employees (aktif)
		{
			UserID:          findUserIDByEmail(adminStaffUsers, "staff1@example.com"),
			FullName:        "Staff One",
			Phone:           "+6280000000004",
			DateOfBirth:     &dob4,
			Salary:          8000000.00,
			ProfileImageURL: "https://example.com/profiles/staff1.jpg",
		},
		{
			UserID:          findUserIDByEmail(adminStaffUsers, "staff2@example.com"),
			FullName:        "Staff Two",
			Phone:           "+6280000000005",
			DateOfBirth:     &dob5,
			Salary:          7500000.00,
			ProfileImageURL: "https://example.com/profiles/staff2.jpg",
		},
		// Staff Employee (tidak aktif)
		{
			UserID:          findUserIDByEmail(adminStaffUsers, "staff.inactive@example.com"),
			FullName:        "Inactive Staff",
			Phone:           "+6280000000006",
			DateOfBirth:     &dob6,
			Salary:          7000000.00,
			ProfileImageURL: "",
		},
	}

	// Hapus data lama jika ada (optional)
	s.DB.Exec("DELETE FROM employees")

	// Filter hanya employee yang memiliki user (tidak termasuk yang nil)
	filteredEmployees := []entity.Employee{}
	for _, emp := range employees {
		if emp.UserID != 0 {
			filteredEmployees = append(filteredEmployees, emp)
		}
	}

	// Create employees
	for i := range filteredEmployees {
		if err := s.DB.Create(&filteredEmployees[i]).Error; err != nil {
			return err
		}
	}

	return nil
}

// Helper function untuk mencari UserID berdasarkan email
func findUserIDByEmail(users []entity.User, email string) uint {
	for _, user := range users {
		if user.Email == email {
			return user.ID
		}
	}
	return 0
}
