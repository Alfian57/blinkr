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

type UrlRepository struct {
	db *gorm.DB
}

func NewUrlRepository() *UrlRepository {
	return &UrlRepository{db: database.DB}
}

func (r *UrlRepository) GetAll(ctx context.Context) ([]model.Url, error) {
	var urls []model.Url
	err := r.db.WithContext(ctx).Find(&urls).Error
	return urls, err
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

func (r *UrlRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Url{}).Count(&count).Error
	return count, err
}

func (r *UrlRepository) Create(ctx context.Context, url *model.Url) error {
	url.ID = uuid.New()

	err := r.db.WithContext(ctx).Create(url).Error
	logger.Log.Debug(err)
	return err
}
