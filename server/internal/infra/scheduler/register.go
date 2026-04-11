package scheduler

import (
	"debian-ecommerce/internal/usecase"

	"go.uber.org/zap"
)

func RegisterCronJobs(
	scheduler *Scheduler,
	usecase *usecase.Usecase,
	log *zap.Logger,
) {
	jobs := NewJobs(usecase, log)

	scheduler.AddFunc("0 15 * * *", jobs.AutoPickupToShipped)
	scheduler.AddFunc("0 0 * * *", jobs.AutoDelivered)
	scheduler.AddFunc("5 0 * * *", jobs.AutoCompleted)
}