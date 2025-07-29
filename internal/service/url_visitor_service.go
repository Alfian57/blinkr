package service

import (
	"context"
	"errors"
	"time"

	errs "github.com/Alfian57/belajar-golang/internal/errors"
	"github.com/Alfian57/belajar-golang/internal/logger"
	"github.com/Alfian57/belajar-golang/internal/repository"
)

type UrlVisitorService struct {
	urlVisitorRepository *repository.UrlVisitorRepository
	urlRepository        *repository.UrlRepository
}

func NewUrlVisitorService(urlVisitorRepository *repository.UrlVisitorRepository, urlRepository *repository.UrlRepository) *UrlVisitorService {
	return &UrlVisitorService{
		urlVisitorRepository: urlVisitorRepository,
		urlRepository:        urlRepository,
	}
}

// CountUrls retrieves the total number of URLs in the repository.
// It returns the count of URLs or an error if the operation fails.
func (s *UrlVisitorService) Count(ctx context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Count the total number of url visitors
	count, err := s.urlVisitorRepository.Count(ctx)
	if err != nil {
		logger.Log.Errorw("failed to count url visitors", "error", err)
		return 0, errs.NewAppError(500, "failed to count url visitors", err)
	}

	return count, nil
}

// CountByUrlID retrieves the total number of URL visitors for a specific URL ID.
// It returns the count of visitors or an error if the operation fails.
func (s *UrlVisitorService) CountByUrlID(ctx context.Context, urlID string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Validate the URL ID
	_, err := s.urlRepository.GetByID(ctx, urlID)
	if err != nil {
		if errors.Is(err, errs.ErrUrlNotFound) {
			logger.Log.Errorw("invalid url id", "url_id", urlID)
			return 0, errs.NewAppError(400, "invalid url id", nil)
		} else {
			logger.Log.Errorw("failed to validate url id", "error", err)
			return 0, errs.NewAppError(500, "failed to validate url id", err)
		}
	}

	// Count the total number of url visitors
	count, err := s.urlVisitorRepository.CountByUrlID(ctx, urlID)
	if err != nil {
		logger.Log.Errorw("failed to count url visitors", "error", err)
		return 0, errs.NewAppError(500, "failed to count url visitors", err)
	}

	return count, nil
}
