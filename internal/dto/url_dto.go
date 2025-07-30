package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateUrlRequest struct {
	ShortUrl  string    `json:"short_url" form:"short_url" binding:"required,min=3,max=255"`
	LongUrl   string    `json:"long_url" form:"long_url" binding:"required,min=3,max=255,url"`
	UserID    string    `json:"user_id" form:"user_id" binding:"required,uuid"`
	ExpiredAt time.Time `json:"expired_at" form:"expired_at" binding:"required"`
}

type UpdateUrlRequest struct {
	ID        uuid.UUID `json:"id" form:"id"`
	ShortUrl  string    `json:"short_url" form:"short_url" binding:"required,min=3,max=255"`
	LongUrl   string    `json:"long_url" form:"long_url" binding:"required,min=3,max=255,url"`
	ExpiredAt time.Time `json:"expired_at" form:"expired_at" binding:"required"`
}

type GetUrlsFilter struct {
	PaginationRequest
	Search    string `json:"search" form:"search" binding:"omitempty,max=255"`
	OrderBy   string `json:"order_by" form:"order_by" binding:"omitempty,oneof=short_url long_url created_at"`
	OrderType string `json:"order_type" form:"order_type" binding:"omitempty,oneof=ASC DESC asc desc"`
}
