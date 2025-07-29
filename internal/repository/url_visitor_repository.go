package repository

import (
	"context"

	"github.com/Alfian57/belajar-golang/internal/database"
	"github.com/Alfian57/belajar-golang/internal/model"
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
