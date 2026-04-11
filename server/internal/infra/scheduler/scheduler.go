package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron *cron.Cron
}

func NewScheduler() *Scheduler {
	loc, _ := time.LoadLocation("Asia/Jakarta")

	c := cron.New(
		cron.WithLocation(loc),
	)

	return &Scheduler{cron: c}
}

func (s *Scheduler) Start() {
	s.cron.Start()
	log.Println("Cron scheduler started")
}

func (s *Scheduler) Stop() context.Context {
	log.Println("Cron scheduler stopped")
	return s.cron.Stop()
}

func (s *Scheduler) AddFunc(spec string, cmd func()) error {
    _, err := s.cron.AddFunc(spec, cmd)
    return err
}