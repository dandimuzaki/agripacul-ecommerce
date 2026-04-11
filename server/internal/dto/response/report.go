package response

type SalesResponse struct {
	Summary SalesSummary   `json:"summary"`
	PerHour []PerHourSales `json:"per_hour"`
	Daily   []DailySales   `json:"daily"`
	Monthly []MonthlySales `json:"monthly"`
}

type SalesSummary struct {
	Total        int64        `json:"total"`
	AverageDaily float64      `json:"average_daily"`
	PeakHour     PerHourSales `json:"peak_hour"`
}

type DailySales struct {
	Date  string `json:"date"`
	Sales int    `json:"sales"`
}

type MonthlySales struct {
	Month string `json:"month"`
	Sales int    `json:"sales"`
}

type PerHourSales struct {
	Hour         string  `json:"hour"`
	AverageSales float64 `json:"average_sales"`
}

type RevenueResponse struct {
	Summary RevenueSummary   `json:"summary"`
	PerHour []PerHourRevenue `json:"per_hour"`
	Daily   []DailyRevenue   `json:"daily"`
	Monthly []MonthlyRevenue `json:"monthly"`
}

type RevenueSummary struct {
	Total        float64        `json:"total"`
	AverageDaily float64        `json:"average_daily"`
	PeakHour     PerHourRevenue `json:"peak_hour"`
}

type DailyRevenue struct {
	Date    string  `json:"date"`
	Revenue float64 `json:"revenue"`
}

type MonthlyRevenue struct {
	Month   string  `json:"month"`
	Revenue float64 `json:"revenue"`
}

type PerHourRevenue struct {
	Hour           string  `json:"hour"`
	AverageRevenue float64 `json:"average_revenue"`
}

type ProductPerformance struct {
	Name       string  `json:"name"`
	Category   string  `json:"category"`
	MinPrice   string  `json:"min_price"`
	MaxPrice   string  `json:"max_price"`
	TotalStock int64   `json:"total_stock"`
	Status     string  `json:"status"`
	Badge      string  `json:"badge"`
	Sales      int64   `json:"sales"`
	Revenue    float64 `json:"revenue"`
}

type CustomerReport struct {
	Name              string  `json:"name"`
	Email             string  `json:"email"`
	PhoneNumber       string  `json:"phone_number"`
	TotalOrder        int     `json:"total_order"`
	TotalOrderValue   float64 `json:"total_order_value"`
	AverageOrderValue float64 `json:"average_order_value"`
}

type CustomerSummary struct {
	TotalRegisteredCustomer int64 `json:"total_registered_customer"`
	TotalCustomerWithOrder  int64 `json:"total_customer_with_order"`
	TotalActiveCustomer     int64 `json:"total_active_customer"`
	TotalNewCustomer        int64 `json:"total_new_customer"`
}