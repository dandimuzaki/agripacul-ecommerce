package scheduler

import (
	"context"
	"debian-ecommerce/internal/usecase"

	"go.uber.org/zap"
)

type Jobs struct {
	service *usecase.Usecase
	log *zap.Logger
}

func NewJobs(service *usecase.Usecase, log *zap.Logger) *Jobs {
	return &Jobs{service: service, log: log}
}

func (j *Jobs) AutoPickupToShipped() {
	defer func() {
		if r := recover(); r != nil {
			j.log.Warn("panic in AutoPickupToShipped", zap.Any("panic", r))
		}
	}()

	ctx := context.Background()

	if err := j.service.OrderService.Ship(ctx); err != nil {
		j.log.Error("cron AutoPickupToShipped failed", zap.Error(err))
	}
}

func (j *Jobs) AutoDelivered() {
	defer func() {
		if r := recover(); r != nil {
			j.log.Warn("panic in AutoDelivered", zap.Any("panic", r))
		}
	}()

	ctx := context.Background()

	if err := j.service.OrderService.Deliver(ctx); err != nil {
		j.log.Error("cron AutoDelivered failed", zap.Error(err))
	}
}

func (j *Jobs) AutoCompleted() {
	defer func() {
		if r := recover(); r != nil {
			j.log.Warn("panic in AutoCompleted", zap.Any("panic", r))
		}
	}()

	ctx := context.Background()

	if err := j.service.OrderService.Complete(ctx); err != nil {
		j.log.Error("cron AutoCompleted failed", zap.Error(err))
	}
}