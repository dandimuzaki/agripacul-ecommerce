package repository

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/dto/response"
	infra "debian-ecommerce/internal/infra/transaction"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ReportRepository interface{
	GetTotalSales(ctx context.Context, f request.ReportQuery) (int64, error)
	GetAverageDailySales(ctx context.Context, f request.ReportQuery) (float64, error)
	GetHourlySales(ctx context.Context, f request.ReportQuery) ([]response.PerHourSales, error)
	GetDailySales(ctx context.Context, f request.ReportQuery) ([]response.DailySales, error)
	GetMonthlySales(ctx context.Context, f request.ReportQuery) ([]response.MonthlySales, error)
	GetTotalRevenue(ctx context.Context, f request.ReportQuery) (float64, error)
	GetAverageDailyRevenue(ctx context.Context, f request.ReportQuery) (float64, error)
	GetHourlyRevenue(ctx context.Context, f request.ReportQuery) ([]response.PerHourRevenue, error)
	GetDailyRevenue(ctx context.Context, f request.ReportQuery) ([]response.DailyRevenue, error)
	GetMonthlyRevenue(ctx context.Context, f request.ReportQuery) ([]response.MonthlyRevenue, error)
	GetProductPerformance(ctx context.Context, f request.ReportQuery) ([]response.ProductPerformance, error)
	GetLoyalCustomer(ctx context.Context) ([]response.CustomerReport, error)
	GetTotalRegisteredCustomer(ctx context.Context) (int64, error)
	GetTotalCustomerWithOrder(ctx context.Context) (int64, error)
	GetTotalNewCustomer(ctx context.Context) (int64, error)
	GetTotalActiveCustomer(ctx context.Context) (int64, error)
}

type reportRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewReportRepo(db *gorm.DB, log *zap.Logger) ReportRepository {
	return &reportRepository{
		db:  db,
		log: log,
	}
}

func (r *reportRepository) GetTotalSales(ctx context.Context, f request.ReportQuery) (int64, error) {
	db := infra.GetDB(ctx, r.db)
	var total int64
	err := db.Model(&entity.Order{}).
		Select("sum(oi.quantity) as total").
		Joins("JOIN order_items oi ON oi.order_id = orders.id").
		Where("orders.created_at > ? AND orders.created_at < ? AND orders.status = ?", f.StartDate, f.EndDate, entity.OrderStatusCompleted).
		Scan(&total).Error

	if err != nil {
		r.log.Error("Error get total sales", zap.Error(err))
		return 0, err
	}

	return total, nil
}

func (r *reportRepository) GetAverageDailySales(ctx context.Context, f request.ReportQuery) (float64, error) {
	db := infra.GetDB(ctx, r.db)
	var avg float64
	err := db.Model(&entity.Order{}).
		Select("avg(oi.quantity) as average").
		Joins("JOIN order_items oi ON oi.order_id = orders.id").
		Where("orders.created_at > ? AND orders.created_at < ? AND orders.status = ?", f.StartDate, f.EndDate, entity.OrderStatusCompleted).
		Group("TO_CHAR(orders.created_at, 'DD-MM-YYYY')").
		Scan(&avg).Error

	if err != nil {
		r.log.Error("Error get average daily sales", zap.Error(err))
		return 0, err
	}

	return avg, nil
}

func (r *reportRepository) GetHourlySales(ctx context.Context, f request.ReportQuery) ([]response.PerHourSales, error) {
	db := infra.GetDB(ctx, r.db)
	var perHour []response.PerHourSales
	err := db.Model(&entity.Order{}).
		Select("TO_CHAR(DATE_TRUNC('hour', orders.created_at), 'HH:00') AS hour, avg(oi.quantity) as average_sales").
		Joins("JOIN order_items oi ON oi.order_id = orders.id").
		Where("orders.created_at > ? AND orders.created_at < ? AND orders.status = ?", f.StartDate, f.EndDate, entity.OrderStatusCompleted).
		Group("hour").
		Order("hour desc").
		Scan(&perHour).Error

	if err != nil {
		r.log.Error("Error get hourly sales", zap.Error(err))
		return nil, err
	}

	return perHour, nil
}

func (r *reportRepository) GetDailySales(ctx context.Context, f request.ReportQuery) ([]response.DailySales, error) {
	db := infra.GetDB(ctx, r.db)
	var daily []response.DailySales
	err := db.Model(&entity.Order{}).
		Select("TO_CHAR(orders.created_at, 'DD-MM-YYYY') as date, sum(oi.quantity) as sales").
		Joins("JOIN order_items oi ON oi.order_id = orders.id").
		Where("orders.created_at > ? AND orders.created_at < ? AND orders.status = ?", f.StartDate, f.EndDate, entity.OrderStatusCompleted).
		Group("date").
		Order("date desc").
		Scan(&daily).Error

	if err != nil {
		r.log.Error("Error get daily sales", zap.Error(err))
		return nil, err
	}

	return daily, nil
}

func (r *reportRepository) GetMonthlySales(ctx context.Context, f request.ReportQuery) ([]response.MonthlySales, error) {
	db := infra.GetDB(ctx, r.db)
	var results []response.MonthlySales
	err := db.Model(&entity.Order{}).
		Select("TO_CHAR(orders.created_at, 'YYYY-MM') as month, sum(oi.quantity) as sales").
		Joins("JOIN order_items oi ON oi.order_id = orders.id").
		Where("orders.created_at > ? AND orders.created_at < ? AND orders.status = ?", f.StartDate, f.EndDate, entity.OrderStatusCompleted).
		Group("month").
		Order("month desc").
		Scan(&results).Error
	
	if err != nil {
		r.log.Error("Error get monthly sales", zap.Error(err))
		return nil, err
	}

	return results, err
}

func (r *reportRepository) GetTotalRevenue(ctx context.Context, f request.ReportQuery) (float64, error) {
	db := infra.GetDB(ctx, r.db)
	var total float64
	err := db.Model(&entity.Order{}).
		Select("sum(total) as total").
		Where("created_at > ? AND created_at < ? AND status = ?", f.StartDate, f.EndDate, entity.OrderStatusCompleted).
		Scan(&total).Error

	if err != nil {
		r.log.Error("Error get total revenue", zap.Error(err))
		return 0, err
	}

	return total, nil
}

func (r *reportRepository) GetAverageDailyRevenue(ctx context.Context, f request.ReportQuery) (float64, error) {
	db := infra.GetDB(ctx, r.db)
	var avg float64
	err := db.Model(&entity.Order{}).
		Select("avg(total) as average").
		Where("created_at > ? AND created_at < ? AND status = ?", f.StartDate, f.EndDate, entity.OrderStatusCompleted).
		Group("TO_CHAR(created_at, 'DD-MM-YYYY')").
		Scan(&avg).Error

	if err != nil {
		r.log.Error("Error get average daily revenue", zap.Error(err))
		return 0, err
	}

	return avg, nil
}

func (r *reportRepository) GetHourlyRevenue(ctx context.Context, f request.ReportQuery) ([]response.PerHourRevenue, error) {
	db := infra.GetDB(ctx, r.db)
	var perHour []response.PerHourRevenue
	err := db.Model(&entity.Order{}).
		Select("TO_CHAR(DATE_TRUNC('hour', created_at), 'HH:00') AS hour, avg(total) as average_revenue").
		Where("created_at > ? AND created_at < ? AND status = ?", f.StartDate, f.EndDate, entity.OrderStatusCompleted).
		Group("hour").
		Order("hour desc").
		Scan(&perHour).Error

	if err != nil {
		r.log.Error("Error get hourly revenue", zap.Error(err))
		return nil, err
	}

	return perHour, nil
}

func (r *reportRepository) GetDailyRevenue(ctx context.Context, f request.ReportQuery) ([]response.DailyRevenue, error) {
	db := infra.GetDB(ctx, r.db)
	var daily []response.DailyRevenue
	err := db.Model(&entity.Order{}).
		Select("TO_CHAR(created_at, 'DD-MM-YYYY') as date, sum(total) as revenue").
		Where("created_at > ? AND created_at < ? AND status = ?", f.StartDate, f.EndDate, entity.OrderStatusCompleted).
		Group("date").
		Order("date desc").
		Scan(&daily).Error

	if err != nil {
		r.log.Error("Error get daily revenue", zap.Error(err))
		return nil, err
	}

	return daily, nil
}

func (r *reportRepository) GetMonthlyRevenue(ctx context.Context, f request.ReportQuery) ([]response.MonthlyRevenue, error) {
	db := infra.GetDB(ctx, r.db)
	var results []response.MonthlyRevenue
	err := db.Model(&entity.Order{}).
		Select("TO_CHAR(created_at, 'YYYY-MM') as month, sum(total) as revenue").
		Where("created_at > ? AND created_at < ? AND status = ?", f.StartDate, f.EndDate, entity.OrderStatusCompleted).
		Group("month").
		Order("month desc").
		Scan(&results).Error
	
	if err != nil {
		r.log.Error("Error get monthly revenue", zap.Error(err))
		return nil, err
	}

	return results, err
}

func (r *reportRepository) GetProductPerformance(ctx context.Context, f request.ReportQuery) ([]response.ProductPerformance, error) {
	db := infra.GetDB(ctx, r.db)
	var results []response.ProductPerformance
	err := db.Model(&entity.Product{}).
    Select(`
      products.name AS name,
      c.name AS category,
      products.min_price,
      products.max_price,
      COALESCE(SUM(DISTINCT s.stock), 0) AS total_stock,
      CASE 
          WHEN EXISTS (
              SELECT 1 FROM skus s2 
              WHERE s2.product_id = products.id AND s2.stock < s2.min_stock
          ) THEN 'low stock'
          WHEN SUM(s.stock) <= 0 THEN 'out of stock'
          ELSE 'in stock'
      END AS status,
			CASE
				WHEN products.created_at >= NOW() - INTERVAL '30 days' THEN 'new'
				ELSE 'regular'
			END AS badge,
      COALESCE(SUM(oi.quantity), 0) AS sales,
      COALESCE(SUM(oi.total_price), 0) AS revenue
    `).
    Joins("LEFT JOIN skus s ON s.product_id = products.id").
    Joins("LEFT JOIN order_items oi ON oi.sku_id = s.id").
    // Move the filters HERE inside the ON clause
    Joins(`
      LEFT JOIN orders o ON oi.order_id = o.id 
      AND o.created_at > ? 
      AND o.created_at < ? 
      AND o.status = ?
    `, f.StartDate, f.EndDate, entity.OrderStatusCompleted).
    Joins("LEFT JOIN categories c ON products.category_id = c.id").
    Where("products.deleted_at IS NULL").
    Group(`
      products.id,
      products.name,
      c.name,
      products.min_price,
      products.max_price
    `).
		Order("COALESCE(SUM(oi.total_price), 0) DESC").
		Scan(&results).Error

	
	if err != nil {
		r.log.Error("Error get product performance", zap.Error(err))
		return nil, err
	}

	return results, err
}

func (r *reportRepository) GetLoyalCustomer(ctx context.Context) ([]response.CustomerReport, error) {
	db := infra.GetDB(ctx, r.db)
	var results []response.CustomerReport
	err := db.Model(&entity.Order{}).
		Select(`
			customers.full_name as name,
			users.email as email,
			customers.phone_number as phone_number,
			COALESCE(COUNT(orders.id), 0) AS total_order,
			COALESCE(SUM(orders.total), 0) AS total_order_value,
			COALESCE(AVG(orders.total), 0) AS average_order_value
		`).
		Joins("LEFT JOIN customers ON customers.id = orders.customer_id").
		Joins("LEFT JOIN users ON customers.user_id = users.id").
		Where("orders.status = ?", entity.OrderStatusCompleted).
		Group(`
			customers.full_name,
			users.email,
			customers.phone_number
		`).
		Order("total_order_value DESC").
		Scan(&results).Error

	
	if err != nil {
		r.log.Error("Error get loyal customer", zap.Error(err))
		return nil, err
	}

	return results, err
}

func (r *reportRepository) GetTotalRegisteredCustomer(ctx context.Context) (int64, error) {
	db := infra.GetDB(ctx, r.db)
	var total int64
	err := db.Model(&entity.Customer{}).Count(&total).Error

	if err != nil {
		r.log.Error("Error get total customer", zap.Error(err))
		return 0, err
	}

	return total, nil
}

func (r *reportRepository) GetTotalCustomerWithOrder(ctx context.Context) (int64, error) {
	db := infra.GetDB(ctx, r.db)
	var total int64
	err := db.Model(&entity.Order{}).Distinct("customer_id").Count(&total).Error

	if err != nil {
		r.log.Error("Error get total customer with order", zap.Error(err))
		return 0, err
	}

	return total, nil
}

func (r *reportRepository) GetTotalNewCustomer(ctx context.Context) (int64, error) {
	db := infra.GetDB(ctx, r.db)
	var total int64
	err := db.Model(&entity.Customer{}).Where("created_at >= NOW() - INTERVAL '30 days'").Count(&total).Error

	if err != nil {
		r.log.Error("Error get total new customer", zap.Error(err))
		return 0, err
	}

	return total, nil
}

func (r *reportRepository) GetTotalActiveCustomer(ctx context.Context) (int64, error) {
	db := infra.GetDB(ctx, r.db)
	var total int64
	err := db.Model(&entity.Order{}).Distinct("customer_id").Where("created_at >= NOW() - INTERVAL '30 days'").Count(&total).Error

	if err != nil {
		r.log.Error("Error get total active customer", zap.Error(err))
		return 0, err
	}

	return total, nil
}