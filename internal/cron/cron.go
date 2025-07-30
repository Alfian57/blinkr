package cron

import (
	"context"

	"github.com/Alfian57/belajar-golang/internal/logger"
	"github.com/go-co-op/gocron/v2"
)

type Crobjob interface {
	Start(ctx context.Context) error
}

func Init() {
	s, err := gocron.NewScheduler()
	if err != nil {
		panic("Failed to create scheduler: " + err.Error())
	}

	cronjobs := []Crobjob{
		NewDeleteUrlCron(s),
	}

	for _, job := range cronjobs {
		if err := job.Start(context.Background()); err != nil {
			panic("Failed to start cron job: " + err.Error())
		}
	}

	logger.Log.Infoln("Cron jobs initialized successfully")
	s.Start()
}
