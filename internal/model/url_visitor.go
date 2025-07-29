package model

import (
	"time"

	"github.com/google/uuid"
)

type URLVisitor struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	UrlID     uuid.UUID `json:"url_id" gorm:"not null;index"`
	IpAddress string    `json:"ip_address" gorm:"not null"`
	UserAgent string    `json:"user_agent" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (URLVisitor) TableName() string {
	return "url_visitors"
}
