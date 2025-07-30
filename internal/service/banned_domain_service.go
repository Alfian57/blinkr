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

type BannedDomainService struct {
	bannedDomainRepository *repository.BannedDomainRepository
}

func NewBannedDomainService(repository *repository.BannedDomainRepository) *BannedDomainService {
	return &BannedDomainService{
		bannedDomainRepository: repository,
	}
}

// GetAllBannedDomains retrieves all banned domains with optional filtering and pagination.
// It returns a paginated result containing banned domain data.
func (s *BannedDomainService) GetAllBannedDomains(ctx context.Context, query dto.GetBannedDomainsFilter) (dto.PaginatedResult[model.BannedDomain], error) {
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

	bannedDomains, err := s.bannedDomainRepository.GetAllWithFilterPagination(ctx, query.Search, orderBy, orderType, limit, offset)
	if err != nil {
		logger.Log.Errorw("failed to retrieve banned domains", "error", err)
		return dto.PaginatedResult[model.BannedDomain]{}, errs.NewAppError(500, "failed to retrieve banned domains", err)
	}

	// Count total banned domains for pagination
	count, err := s.bannedDomainRepository.CountByUrl(ctx, query.Search)
	if err != nil {
		logger.Log.Errorw("failed to count banned domains", "error", err)
		return dto.PaginatedResult[model.BannedDomain]{}, errs.NewAppError(500, "failed to retrieve banned domains", err)
	}

	// Create pagination response
	pagination := dto.NewPaginationResponse(query.Page, query.Limit, count)
	result := dto.PaginatedResult[model.BannedDomain]{
		Data:       bannedDomains,
		Pagination: pagination,
	}

	return result, nil
}

// CreateBannedDomain creates a new banned domain with the provided request data.
// It returns an error if the creation fails.
func (s *BannedDomainService) CreateBannedDomain(ctx context.Context, request dto.CreateBannedDomainRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	bannedDomain := model.BannedDomain{
		URL: request.Url,
	}

	// Create the banned domain
	if err := s.bannedDomainRepository.Create(ctx, &bannedDomain); err != nil {
		logger.Log.Errorw("failed to create banned domain", "url", request.Url, "error", err)
		return errs.NewAppError(500, "failed to create banned domain", err)
	}

	logger.Log.Infow("banned domain created successfully", "url", request.Url)
	return nil
}

// UpdateBannedDomain updates an existing banned domain with the provided request data.
// It returns an error if the banned domain does not exist or if the update fails.
func (s *BannedDomainService) UpdateBannedDomain(ctx context.Context, request dto.UpdateBannedDomainRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Check if the banned domain exists
	_, err := s.bannedDomainRepository.GetByID(ctx, request.ID.String())
	if err != nil {
		if err == errs.ErrBannedDomainNotFound {
			return err
		}

		logger.Log.Errorw("failed to check banned domain existence for update", "id", request.ID, "error", err)
		return errs.NewAppError(500, "failed to validate banned domain", err)
	}

	// Prepare banned domain data for update
	bannedDomain := model.BannedDomain{
		ID:  request.ID,
		URL: request.Url,
	}

	// Update the banned domain
	if err := s.bannedDomainRepository.Update(ctx, &bannedDomain); err != nil {
		logger.Log.Errorw("failed to update banned domain", "id", request.ID, "error", err)
		return errs.NewAppError(500, "failed to update banned domain", err)
	}

	logger.Log.Infow("banned domain updated successfully", "id", request.ID)
	return nil
}

// DeleteBannedDomain deletes a banned domain by their ID.
// It returns an error if the banned domain does not exist or if the deletion fails.
func (s *BannedDomainService) DeleteBannedDomain(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Check if the banned domain exists
	if err := s.bannedDomainRepository.Delete(ctx, id.String()); err != nil {
		if err == errs.ErrBannedDomainNotFound {
			return err
		}
		logger.Log.Errorw("failed to delete banned domain", "id", id, "error", err)
		return errs.NewAppError(500, "failed to delete banned domain", err)
	}

	logger.Log.Infow("banned domain deleted successfully", "id", id)
	return nil
}
