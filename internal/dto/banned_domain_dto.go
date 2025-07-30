package dto

import "github.com/google/uuid"

type CreateBannedDomainRequest struct {
	Url string `json:"url" form:"url" binding:"required,min=3,max=255,url"`
}

type UpdateBannedDomainRequest struct {
	ID  uuid.UUID `json:"id" form:"id"`
	Url string    `json:"url" form:"url" binding:"required,min=3,max=255,url"`
}

type GetBannedDomainsFilter struct {
	PaginationRequest
	Search    string `json:"search" form:"search" binding:"omitempty,max=255"`
	OrderBy   string `json:"order_by" form:"order_by" binding:"omitempty,oneof=url created_at"`
	OrderType string `json:"order_type" form:"order_type" binding:"omitempty,oneof=ASC DESC asc desc"`
}
