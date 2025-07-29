package model

import (
	"time"

	"github.com/google/uuid"
)

type BannedDomain struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	URL       string    `gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (BannedDomain) TableName() string {
	return "banned_domains"
}
