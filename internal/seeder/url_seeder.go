package seeder

import (
	"context"

	"github.com/Alfian57/belajar-golang/internal/factory"
	"github.com/Alfian57/belajar-golang/internal/logger"
	"github.com/Alfian57/belajar-golang/internal/model"
	"github.com/Alfian57/belajar-golang/internal/repository"
	"github.com/google/uuid"
)

type UrlSeeder struct {
	urlRepository  *repository.UrlRepository
	userRepisotory *repository.UserRepository
	useFactory     bool
	count          int
}

func NewUrlSeeder(useFactory bool, count int) *UrlSeeder {
	return &UrlSeeder{
		urlRepository:  repository.NewUrlRepository(),
		userRepisotory: repository.NewUserRepository(),
		useFactory:     useFactory,
		count:          count,
	}
}

func (s *UrlSeeder) Seed(ctx context.Context) error {
	logger.Log.Info("Starting url seeding...")

	if s.useFactory {
		return s.seedWithFactory(ctx)
	}
	return s.seedManual(ctx)
}

func (s *UrlSeeder) seedManual(ctx context.Context) error {
	logger.Log.Info("Seeding urls manually...")

	urls := []model.Url{
		// Add your manual URL entries here
	}

	return s.createUrls(ctx, urls)
}

func (s *UrlSeeder) seedWithFactory(ctx context.Context) error {
	logger.Log.Infof("Seeding %d urls with factory...", s.count)

	var urls []model.Url
	var userIDs uuid.UUIDs

	// Get all users to assign UserID for URLs
	users, err := s.userRepisotory.GetAll(ctx)
	if err != nil {
		logger.Log.Errorw("Failed to get users for URL seeding", "error", err)
		return err
	}
	for _, user := range users {
		userIDs = append(userIDs, user.ID)
	}

	// Create URLs using the factory
	urlFactory := factory.NewUrlFactory(userIDs)
	for i := 0; i < s.count; i++ {
		urlsDummy := urlFactory.MustCreate().(*model.Url)
		urls = append(urls, *urlsDummy)
	}

	return s.createUrls(ctx, urls)
}

func (s *UrlSeeder) createUrls(ctx context.Context, urls []model.Url) error {
	for _, url := range urls {
		// Create url
		if err := s.urlRepository.Create(ctx, &url); err != nil {
			logger.Log.Errorw("Failed to create urls", "urls", url.ShortUrl, "error", err)
			return err
		}
	}

	return nil
}
