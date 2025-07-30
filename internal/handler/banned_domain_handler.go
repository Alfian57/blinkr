package handler

import (
	"net/http"

	"github.com/Alfian57/belajar-golang/internal/dto"
	"github.com/Alfian57/belajar-golang/internal/response"
	"github.com/Alfian57/belajar-golang/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BannedDomainHandler struct {
	bannedDomainService *service.BannedDomainService
}

func NewBannedDomainHandler(service *service.BannedDomainService) *BannedDomainHandler {
	return &BannedDomainHandler{
		bannedDomainService: service,
	}
}

func (h *BannedDomainHandler) GetAllBannedDomains(ctx *gin.Context) {
	var query dto.GetBannedDomainsFilter
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	query.PaginationRequest.SetDefaults()

	result, err := h.bannedDomainService.GetAllBannedDomains(ctx, query)
	if err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WritePaginatedResponse(ctx, http.StatusOK, result)
}

func (h *BannedDomainHandler) CreateBannedDomain(ctx *gin.Context) {
	var request dto.CreateBannedDomainRequest
	if err := ctx.ShouldBind(&request); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	if err := h.bannedDomainService.CreateBannedDomain(ctx, request); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WriteMessageResponse(ctx, http.StatusCreated, "banned domain successfully created")
}

func (h *BannedDomainHandler) UpdateBannedDomain(ctx *gin.Context) {
	var request dto.UpdateBannedDomainRequest
	if err := ctx.ShouldBind(&request); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}
	request.ID = id

	if err := h.bannedDomainService.UpdateBannedDomain(ctx, request); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WriteMessageResponse(ctx, http.StatusOK, "banned domain successfully updated")
}

func (h *BannedDomainHandler) DeleteBannedDomain(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	if err := h.bannedDomainService.DeleteBannedDomain(ctx, id); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WriteMessageResponse(ctx, http.StatusOK, "banned domain successfully deleted")
}
