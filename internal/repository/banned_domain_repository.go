package repository

import (
	"context"
	"errors"

	"github.com/Alfian57/belajar-golang/internal/database"
	errs "github.com/Alfian57/belajar-golang/internal/errors"
	"github.com/Alfian57/belajar-golang/internal/logger"
	"github.com/Alfian57/belajar-golang/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BannedDomainRepository struct {
	db *gorm.DB
}

func NewBannedDomainRepository() *BannedDomainRepository {
	return &BannedDomainRepository{db: database.DB}
}

// GetAllWithFilterPagination retrieves bannedDomains with optional filters, ordering, and pagination
func (r *BannedDomainRepository) GetAllWithFilterPagination(ctx context.Context, search string, orderBy string, orderType string, limit int, offset int) ([]model.BannedDomain, error) {
	var bannedDomains []model.BannedDomain

	query := r.db.WithContext(ctx)

	// Apply search filter
	if search != "" {
		query = query.Where("url LIKE ?", "%"+search+"%")
	}

	// Apply ordering
	if orderBy != "" && orderType != "" {
		query = query.Order(orderBy + " " + orderType)
	}

	// Apply limit and offset
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	err := query.Find(&bannedDomains).Error
	return bannedDomains, err
}

// GetAll retrieves all bannedDomains without any filters
func (r *BannedDomainRepository) GetAll(ctx context.Context) ([]model.BannedDomain, error) {
	var bannedDomains []model.BannedDomain
	err := r.db.WithContext(ctx).Find(&bannedDomains).Error
	return bannedDomains, err
}

// Count returns the total number of all bannedDomains
func (r *BannedDomainRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.BannedDomain{}).Count(&count).Error
	return count, err
}

// CountWithFilter returns the total number of users matching the search criteria
func (r *BannedDomainRepository) CountByUrl(ctx context.Context, search string) (int64, error) {
	var count int64

	query := r.db.WithContext(ctx).Model(&model.BannedDomain{})

	// Apply search filter
	if search != "" {
		query = query.Where("url LIKE ?", "%"+search+"%")
	}

	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *BannedDomainRepository) Create(ctx context.Context, bannedDomain *model.BannedDomain) error {
	bannedDomain.ID = uuid.New()

	err := r.db.WithContext(ctx).Create(bannedDomain).Error
	logger.Log.Debug(err)
	return err
}

func (r *BannedDomainRepository) GetByID(ctx context.Context, id string) (model.BannedDomain, error) {
	var bannedDomain model.BannedDomain

	err := r.db.WithContext(ctx).First(&bannedDomain, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return bannedDomain, errs.ErrBannedDomainNotFound
		}
		return bannedDomain, err
	}

	return bannedDomain, nil
}

func (r *BannedDomainRepository) Update(ctx context.Context, bannedDomain *model.BannedDomain) error {
	err := r.db.WithContext(ctx).Model(bannedDomain).Select("url").Updates(bannedDomain).Error
	return err
}

func (r *BannedDomainRepository) Delete(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&model.BannedDomain{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errs.ErrBannedDomainNotFound
	}

	return nil
}
