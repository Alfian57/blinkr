package model

import (
	"time"

	"github.com/google/uuid"
)

type Url struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	ShortUrl  string    `json:"short_url" gorm:"not null"`
	LongUrl   string    `json:"long_url" gorm:"not null"`
	UserID    uuid.UUID `json:"user_id" gorm:"not null"`
	ExpiredAt time.Time `json:"expired_at" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Url) TableName() string {
	return "urls"
}
