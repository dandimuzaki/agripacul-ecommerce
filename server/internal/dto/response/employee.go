package response

type EmployeeResponse struct {
	ID        uint    `json:"id"`
	UserID    uint    `json:"user_id"`
	Name      string  `json:"full_name"`
	Email     string  `json:"email"`
	Role      string  `json:"role"`
	Phone     string  `json:"phone"`
	Salary    float64 `json:"salary"`
	IsActive  bool    `json:"is_active"`
	CreatedAt string  `json:"created_at"`
}
