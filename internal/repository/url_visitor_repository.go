package repository

import (
	"context"

	"github.com/Alfian57/belajar-golang/internal/database"
	"github.com/Alfian57/belajar-golang/internal/logger"
	"github.com/Alfian57/belajar-golang/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UrlVisitorRepository struct {
	db *gorm.DB
}

func NewUrlVisitorRepository() *UrlVisitorRepository {
	return &UrlVisitorRepository{db: database.DB}
}

func (r *UrlVisitorRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.URLVisitor{}).Count(&count).Error
	return count, err
}

func (r *UrlVisitorRepository) CountByUrlID(ctx context.Context, urlID string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.URLVisitor{}).Where("url_id = ?", urlID).Count(&count).Error
	return count, err
}

func (r *UrlVisitorRepository) Create(ctx context.Context, urlVisitor *model.URLVisitor) error {
	urlVisitor.ID = uuid.New()

	err := r.db.WithContext(ctx).Create(urlVisitor).Error
	logger.Log.Debug(err)
	return err
}
