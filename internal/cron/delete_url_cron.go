package cron

import (
	"context"
	"fmt"
	"time"

	"github.com/Alfian57/belajar-golang/internal/logger"
	"github.com/Alfian57/belajar-golang/internal/repository"
	"github.com/go-co-op/gocron/v2"
)

type DeleteUrlCron struct {
	scheduler     gocron.Scheduler
	urlRepository *repository.UrlRepository
}

func NewDeleteUrlCron(scheduler gocron.Scheduler) *DeleteUrlCron {
	return &DeleteUrlCron{
		scheduler:     scheduler,
		urlRepository: repository.NewUrlRepository(),
	}
}

func (c *DeleteUrlCron) Start(ctx context.Context) error {
	_, err := c.scheduler.NewJob(
		gocron.DailyJob(
			1, // Every 1 day
			gocron.NewAtTimes(
				gocron.NewAtTime(0, 0, 0), // At midnight
			),
		),
		gocron.NewTask(
			func() {
				expiredUrls, err := c.urlRepository.GetByExpiredMoreThan(ctx, time.Now())
				if err != nil {
					logger.Log.Infoln("URL Cleanup: Failed to retrieve expired URLs: %v", err)
					return
				}

				if len(expiredUrls) == 0 {
					logger.Log.Infoln("URL Cleanup: No expired URLs found")
					return
				}

				for _, url := range expiredUrls {
					if err := c.urlRepository.Delete(ctx, url.ID.String()); err != nil {
						logger.Log.Infoln("URL Cleanup: Failed to delete URL %s: %v", url.ShortUrl, err)
					} else {
						logger.Log.Infoln("URL Cleanup: Successfully deleted URL %s", url.ShortUrl)
					}
				}

				logger.Log.Infoln("URL Cleanup: URL cleanup process completed successfully")
			},
		),
	)

	if err != nil {
		return fmt.Errorf("failed to create delete URL cron job: %w", err)
	}

	return nil
}
