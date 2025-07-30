package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Alfian57/belajar-golang/internal/database"
	errs "github.com/Alfian57/belajar-golang/internal/errors"
	"github.com/Alfian57/belajar-golang/internal/logger"
	"github.com/Alfian57/belajar-golang/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UrlRepository struct {
	db *gorm.DB
}

func NewUrlRepository() *UrlRepository {
	return &UrlRepository{db: database.DB}
}

func (r *UrlRepository) GetAllWithFilterPagination(ctx context.Context, search string, orderBy string, orderType string, limit int, offset int) ([]model.Url, error) {
	var urls []model.Url

	query := r.db.WithContext(ctx)

	// Apply search filter
	if search != "" {
		query = query.Where("short_url LIKE ?", "%"+search+"%")
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

	err := query.Find(&urls).Error
	return urls, err
}

func (r *UrlRepository) GetAll(ctx context.Context) ([]model.Url, error) {
	var urls []model.Url
	err := r.db.WithContext(ctx).Find(&urls).Error
	return urls, err
}

func (r *UrlRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Url{}).Count(&count).Error
	return count, err
}

func (r *UrlRepository) CountByShortUrl(ctx context.Context, search string) (int64, error) {
	var count int64

	query := r.db.WithContext(ctx).Model(&model.Url{})

	// Apply search filter
	if search != "" {
		query = query.Where("short_url LIKE ?", "%"+search+"%")
	}

	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *UrlRepository) Create(ctx context.Context, url *model.Url) error {
	url.ID = uuid.New()

	err := r.db.WithContext(ctx).Create(url).Error
	logger.Log.Debug(err)
	return err
}

func (r *UrlRepository) GetByID(ctx context.Context, id string) (model.Url, error) {
	var url model.Url

	err := r.db.WithContext(ctx).First(&url, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return url, errs.ErrUrlNotFound
		}
		return url, err
	}

	return url, nil
}

func (r *UrlRepository) GetByExpiredMoreThan(ctx context.Context, expiredTime time.Time) ([]model.Url, error) {
	var urls []model.Url

	err := r.db.WithContext(ctx).Where("expired_at > ?", expiredTime).Find(&urls).Error
	if err != nil {
		return nil, err
	}
	if len(urls) == 0 {
		return nil, errs.ErrUrlNotFound
	}

	return urls, nil
}

func (r *UrlRepository) Update(ctx context.Context, url *model.Url) error {
	err := r.db.WithContext(ctx).Model(url).Select("short_url", "long_url", "user_id").Updates(url).Error
	return err
}

func (r *UrlRepository) Delete(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&model.Url{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errs.ErrUrlNotFound
	}

	return nil
}
