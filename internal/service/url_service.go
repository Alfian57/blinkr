package service

import (
	"context"
	"time"

	"github.com/Alfian57/belajar-golang/internal/dto"
	errs "github.com/Alfian57/belajar-golang/internal/errors"
	"github.com/Alfian57/belajar-golang/internal/logger"
	"github.com/Alfian57/belajar-golang/internal/model"
	"github.com/Alfian57/belajar-golang/internal/repository"
	"github.com/google/uuid"
)

type UrlService struct {
	urlRepository  *repository.UrlRepository
	userRepository *repository.UserRepository
}

func NewUrlService(urlRepository *repository.UrlRepository, userRepository *repository.UserRepository) *UrlService {
	return &UrlService{
		urlRepository:  urlRepository,
		userRepository: userRepository,
	}
}

// GetAllUrls retrieves all urls with optional filtering and pagination.
// It returns a paginated result containing url data.
func (s *UrlService) GetAllUrls(ctx context.Context, query dto.GetUrlsFilter) (dto.PaginatedResult[model.Url], error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Set default ordering
	orderBy := query.OrderBy
	if orderBy == "" {
		orderBy = "created_at"
	}
	orderType := query.OrderType
	if orderType != "ASC" && orderType != "DESC" {
		orderType = "ASC"
	}

	// Validate pagination parameters (SetDefaults should be called before this)
	limit := query.PaginationRequest.Limit
	offset := query.PaginationRequest.GetOffset()

	urls, err := s.urlRepository.GetAllWithFilterPagination(ctx, query.Search, orderBy, orderType, limit, offset)
	if err != nil {
		logger.Log.Errorw("failed to retrieve urls", "error", err)
		return dto.PaginatedResult[model.Url]{}, errs.NewAppError(500, "failed to retrieve urls", err)
	}

	// Count total urls for pagination
	count, err := s.urlRepository.CountByShortUrl(ctx, query.Search)
	if err != nil {
		logger.Log.Errorw("failed to count urls", "error", err)
		return dto.PaginatedResult[model.Url]{}, errs.NewAppError(500, "failed to retrieve urls", err)
	}

	// Create pagination response
	pagination := dto.NewPaginationResponse(query.Page, query.Limit, count)
	result := dto.PaginatedResult[model.Url]{
		Data:       urls,
		Pagination: pagination,
	}

	return result, nil
}

// CreateUrl creates a new url with the provided request data.
// It returns an error if the creation fails.
func (s *UrlService) CreateUrl(ctx context.Context, request dto.CreateUrlRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Validate user existence
	_, err := s.userRepository.GetByID(ctx, request.UserID)
	if err != nil {
		if err == errs.ErrUserNotFound {
			return errs.ErrUserNotFound
		}
		logger.Log.Errorw("failed to check user existence", "userID", request.UserID, "error", err)
		return errs.NewAppError(500, "failed to validate user", err)
	}

	userID, err := uuid.Parse(request.UserID)
	if err != nil {
		logger.Log.Errorw("invalid userID format", "userID", request.UserID, "error", err)
		return errs.NewAppError(400, "invalid userID format", err)
	}

	url := model.Url{
		ShortUrl:  request.ShortUrl,
		LongUrl:   request.LongUrl,
		UserID:    userID,
		ExpiredAt: request.ExpiredAt,
	}

	// Create the url
	if err := s.urlRepository.Create(ctx, &url); err != nil {
		logger.Log.Errorw("failed to create url", "short_url", request.ShortUrl, "error", err)
		return errs.NewAppError(500, "failed to create url", err)
	}

	logger.Log.Infow("url created successfully", "short_url", request.ShortUrl)
	return nil
}

// GetUrlByID retrieves a url by its ID.
// It returns the url data or an error if the url does not exist.
func (s *UrlService) GetUrlByID(ctx context.Context, id string) (model.Url, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Check if the url exists
	url, err := s.urlRepository.GetByID(ctx, id)
	if err != nil {
		if err == errs.ErrUrlNotFound {
			return model.Url{}, err
		}
		logger.Log.Errorw("failed to get url by ID", "id", id, "error", err)
		return model.Url{}, errs.NewAppError(500, "failed to retrieve url", err)
	}
	return url, nil
}

// UpdateUrl updates an existing url with the provided request data.
// It returns an error if the url does not exist or if the update fails.
func (s *UrlService) UpdateUrl(ctx context.Context, request dto.UpdateUrlRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Check if the url exists
	currentUrl, err := s.urlRepository.GetByID(ctx, request.ID.String())
	if err != nil {
		if err == errs.ErrUrlNotFound {
			return err
		}

		logger.Log.Errorw("failed to check url existence for update", "id", request.ID, "error", err)
		return errs.NewAppError(500, "failed to validate url", err)
	}

	// Prepare url data for update
	url := model.Url{
		ID:        request.ID,
		ShortUrl:  request.ShortUrl,
		LongUrl:   request.LongUrl,
		UserID:    currentUrl.UserID, // Keep the original user ID
		ExpiredAt: request.ExpiredAt,
	}

	// Update the url
	if err := s.urlRepository.Update(ctx, &url); err != nil {
		logger.Log.Errorw("failed to update url", "id", request.ID, "error", err)
		return errs.NewAppError(500, "failed to update url", err)
	}

	logger.Log.Infow("url updated successfully", "id", request.ID)
	return nil
}

// DeleteUrl deletes a url by their ID.
// It returns an error if the url does not exist or if the deletion fails.
func (s *UrlService) DeleteUrl(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Check if the url exists
	if err := s.urlRepository.Delete(ctx, id.String()); err != nil {
		if err == errs.ErrUrlNotFound {
			return err
		}
		logger.Log.Errorw("failed to delete url", "id", id, "error", err)
		return errs.NewAppError(500, "failed to delete url", err)
	}

	logger.Log.Infow("url deleted successfully", "id", id)
	return nil
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
