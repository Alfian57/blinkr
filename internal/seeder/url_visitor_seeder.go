package seeder

import (
	"context"

	"github.com/Alfian57/belajar-golang/internal/factory"
	"github.com/Alfian57/belajar-golang/internal/logger"
	"github.com/Alfian57/belajar-golang/internal/model"
	"github.com/Alfian57/belajar-golang/internal/repository"
	"github.com/google/uuid"
)

type UrlVisitorSeeder struct {
	urlRepository        *repository.UrlRepository
	urlVisitorRepository *repository.UrlVisitorRepository
	useFactory           bool
	count                int
}

func NewUrlVisitorSeeder(useFactory bool, count int) *UrlVisitorSeeder {
	return &UrlVisitorSeeder{
		urlRepository:        repository.NewUrlRepository(),
		urlVisitorRepository: repository.NewUrlVisitorRepository(),
		useFactory:           useFactory,
		count:                count,
	}
}

func (s *UrlVisitorSeeder) Seed(ctx context.Context) error {
	logger.Log.Info("Starting url visitor seeding...")

	if s.useFactory {
		return s.seedWithFactory(ctx)
	}
	return s.seedManual(ctx)
}

func (s *UrlVisitorSeeder) seedManual(ctx context.Context) error {
	logger.Log.Info("Seeding url visitors manually...")

	urls := []model.URLVisitor{
		// Add your manual URL entries here
	}

	return s.createUrlVisitors(ctx, urls)
}

func (s *UrlVisitorSeeder) seedWithFactory(ctx context.Context) error {
	logger.Log.Infof("Seeding %d url visitors with factory...", s.count)

	var urlVisitors []model.URLVisitor
	var urlIDs uuid.UUIDs

	// Get all urls to assign UserID for URLs
	urls, err := s.urlRepository.GetAll(ctx)
	if err != nil {
		logger.Log.Errorw("Failed to get urls for url visitor seeding", "error", err)
		return err
	}
	for _, url := range urls {
		urlIDs = append(urlIDs, url.ID)
	}

	// Create URL visitors using the factory
	urlVisitorFactory := factory.NewUrlVisitorFactory(urlIDs)
	for i := 0; i < s.count; i++ {
		urlVisitorsDummy := urlVisitorFactory.MustCreate().(*model.URLVisitor)
		urlVisitors = append(urlVisitors, *urlVisitorsDummy)
	}

	return s.createUrlVisitors(ctx, urlVisitors)
}

func (s *UrlVisitorSeeder) createUrlVisitors(ctx context.Context, urls []model.URLVisitor) error {
	for _, url := range urls {
		// Create url visitor
		if err := s.urlVisitorRepository.Create(ctx, &url); err != nil {
			logger.Log.Errorw("Failed to create url visitors", "url visitors", url.IpAddress, "error", err)
			return err
		}
	}

	return nil
}
