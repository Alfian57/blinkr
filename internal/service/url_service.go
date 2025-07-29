package service

import (
	"context"
	"time"

	errs "github.com/Alfian57/belajar-golang/internal/errors"
	"github.com/Alfian57/belajar-golang/internal/logger"
	"github.com/Alfian57/belajar-golang/internal/repository"
)

type UrlService struct {
	urlRepository *repository.UrlRepository
}

func NewUrlService(urlRepository *repository.UrlRepository) *UrlService {
	return &UrlService{
		urlRepository: urlRepository,
	}
}

// CountUrls retrieves the total number of URLs in the repository.
// It returns the count of URLs or an error if the operation fails.
func (s *UrlService) Count(ctx context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Count the total number of urls
	count, err := s.urlRepository.Count(ctx)
	if err != nil {
		logger.Log.Errorw("failed to count urls", "error", err)
		return 0, errs.NewAppError(500, "failed to count urls", err)
	}

	return count, nil
}
