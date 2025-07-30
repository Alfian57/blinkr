package seeder

import (
	"context"

	"github.com/Alfian57/belajar-golang/internal/logger"
)

type Seeder interface {
	Seed(ctx context.Context) error
}

type SeederConfig struct {
	UseFactory      bool
	UserCount       int
	UrlCount        int
	UrlVisitorCount int
}

type DatabaseSeeder struct {
	config  SeederConfig
	seeders []Seeder
}

func NewDatabaseSeeder(config SeederConfig) *DatabaseSeeder {
	seeders := []Seeder{
		NewUserSeeder(config.UseFactory, config.UserCount),
		NewUrlSeeder(config.UseFactory, config.UrlCount),
		NewUrlVisitorSeeder(config.UseFactory, config.UrlVisitorCount),
		// Add other seeders here as needed
	}

	return &DatabaseSeeder{
		config:  config,
		seeders: seeders,
	}
}

func (ds *DatabaseSeeder) SeedAll(ctx context.Context) error {
	if ds.config.UseFactory {
		logger.Log.Info("Starting database seeding with factory...")
	} else {
		logger.Log.Info("Starting database seeding manually...")
	}

	for _, seeder := range ds.seeders {
		if err := seeder.Seed(ctx); err != nil {
			logger.Log.Errorw("Seeding failed", "error", err)
			return err
		}
	}

	logger.Log.Info("All database seeding completed successfully!")
	return nil
}
