package usecase

import (
	"context"
	"debian-ecommerce/internal/data/repository"
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/dto/response"

	"go.uber.org/zap"
)

type ReportService interface{
	GetSales(ctx context.Context, f request.ReportQuery) (*response.SalesResponse, error)
	GetRevenue(ctx context.Context, f request.ReportQuery) (*response.RevenueResponse, error)
	GetProductPerformance(ctx context.Context, f request.ReportQuery) ([]response.ProductPerformance, error)
	GetCustomerReport(ctx context.Context) ([]response.CustomerReport, error)
	GetCustomerSummary(ctx context.Context) (*response.CustomerSummary, error)
}

type reportService struct {
	tx   TxManager
	repo *repository.Repository
	log  *zap.Logger
}

func NewReportService(tx TxManager, repo *repository.Repository, log *zap.Logger) ReportService {
	return &reportService{
		tx:   tx,
		repo: repo,
		log:  log,
	}
}

func (s *reportService) GetSales(ctx context.Context, f request.ReportQuery) (*response.SalesResponse, error) {
	total, err := s.repo.ReportRepo.GetTotalSales(ctx, f)
	if err != nil {
		s.log.Error("Error get total sales", zap.Error(err))
		return nil, err
	}

	avgDaily, err := s.repo.ReportRepo.GetAverageDailySales(ctx, f)
	if err != nil {
		s.log.Error("Error get average daily sales", zap.Error(err))
		return nil, err
	}

	perHour, err := s.repo.ReportRepo.GetHourlySales(ctx, f)
	if err != nil {
		s.log.Error("Error get hourly sales", zap.Error(err))
		return nil, err
	}

	daily, err := s.repo.ReportRepo.GetDailySales(ctx, f)
	if err != nil {
		s.log.Error("Error get daily sales", zap.Error(err))
		return nil, err
	}

	monthly, err := s.repo.ReportRepo.GetMonthlySales(ctx, f)
	if err != nil {
		s.log.Error("Error get monthly sales", zap.Error(err))
		return nil, err
	}

	res := response.SalesResponse{
		Summary: response.SalesSummary{
			Total: total,
			AverageDaily: avgDaily,
		},
		PerHour: perHour,
		Daily: daily,
		Monthly: monthly,
	}

	return &res, nil
}

func (s *reportService) GetRevenue(ctx context.Context, f request.ReportQuery) (*response.RevenueResponse, error) {
	total, err := s.repo.ReportRepo.GetTotalRevenue(ctx, f)
	if err != nil {
		s.log.Error("Error get total revenue", zap.Error(err))
		return nil, err
	}

	avgDaily, err := s.repo.ReportRepo.GetAverageDailyRevenue(ctx, f)
	if err != nil {
		s.log.Error("Error get average daily revenue", zap.Error(err))
		return nil, err
	}

	perHour, err := s.repo.ReportRepo.GetHourlyRevenue(ctx, f)
	if err != nil {
		s.log.Error("Error get hourly revenue", zap.Error(err))
		return nil, err
	}

	daily, err := s.repo.ReportRepo.GetDailyRevenue(ctx, f)
	if err != nil {
		s.log.Error("Error get daily revenue", zap.Error(err))
		return nil, err
	}

	monthly, err := s.repo.ReportRepo.GetMonthlyRevenue(ctx, f)
	if err != nil {
		s.log.Error("Error get monthly revenue", zap.Error(err))
		return nil, err
	}

	res := response.RevenueResponse{
		Summary: response.RevenueSummary{
			Total: total,
			AverageDaily: avgDaily,
		},
		PerHour: perHour,
		Daily: daily,
		Monthly: monthly,
	}

	return &res, nil
}

func (s *reportService) GetProductPerformance(ctx context.Context, f request.ReportQuery) ([]response.ProductPerformance, error) {
	result, err := s.repo.ReportRepo.GetProductPerformance(ctx, f)
	if err != nil {
		s.log.Error("Error get product performance", zap.Error(err))
		return nil, err
	}
	return result, nil
}

func (s *reportService) GetCustomerReport(ctx context.Context) ([]response.CustomerReport, error) {
	result, err := s.repo.ReportRepo.GetLoyalCustomer(ctx)
	if err != nil {
		s.log.Error("Error get loyal customer", zap.Error(err))
		return nil, err
	}
	return result, nil
}

func (s *reportService) GetCustomerSummary(ctx context.Context) (*response.CustomerSummary, error) {
	registered, err := s.repo.ReportRepo.GetTotalRegisteredCustomer(ctx)
	if err != nil {
		s.log.Error("Error get total registered customer", zap.Error(err))
		return nil, err
	}

	withOrder, err := s.repo.ReportRepo.GetTotalCustomerWithOrder(ctx)
	if err != nil {
		s.log.Error("Error get total customer with order", zap.Error(err))
		return nil, err
	}

	new, err := s.repo.ReportRepo.GetTotalNewCustomer(ctx)
	if err != nil {
		s.log.Error("Error get total new customer", zap.Error(err))
		return nil, err
	}

	active, err := s.repo.ReportRepo.GetTotalActiveCustomer(ctx)
	if err != nil {
		s.log.Error("Error get total active customer", zap.Error(err))
		return nil, err
	}
	
	return &response.CustomerSummary{
		TotalRegisteredCustomer: registered,
		TotalCustomerWithOrder: withOrder,
		TotalNewCustomer: new,
		TotalActiveCustomer: active,
	}, nil
}