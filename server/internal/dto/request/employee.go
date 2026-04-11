package request

type CreateEmployeeByAdminRequest struct {
	Name  string `json:"full_name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Role  string `json:"role" binding:"required,oneof=admin staff superadmin"`
}

type UpdateEmployeeRequest struct {
	Name   string  `json:"full_name"`
	Email  string  `json:"email" binding:"omitempty,email"`
	Role   string  `json:"role" binding:"omitempty,oneof=admin staff superadmin"`
	Phone  string  `json:"phone"`
	Salary float64 `json:"salary" binding:"omitempty,min=0"`
}

type PaginationResponse struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
}
