package handler

import (
	"net/http"

	"github.com/Alfian57/belajar-golang/internal/dto"
	"github.com/Alfian57/belajar-golang/internal/logger"
	"github.com/Alfian57/belajar-golang/internal/response"
	"github.com/Alfian57/belajar-golang/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UrlHandler struct {
	service *service.UrlService
}

func NewUrlHandler(s *service.UrlService) *UrlHandler {
	return &UrlHandler{
		service: s,
	}
}

func (h *UrlHandler) GetAllUrls(ctx *gin.Context) {
	var query dto.GetUrlsFilter
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	query.PaginationRequest.SetDefaults()

	result, err := h.service.GetAllUrls(ctx, query)
	if err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WritePaginatedResponse(ctx, http.StatusOK, result)
}

func (h *UrlHandler) CreateUrl(ctx *gin.Context) {
	var request dto.CreateUrlRequest
	if err := ctx.ShouldBind(&request); err != nil {
		logger.Log.Errorw("error", err)
		response.WriteErrorResponse(ctx, err)
		return
	}

	if err := h.service.CreateUrl(ctx, request); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WriteMessageResponse(ctx, http.StatusCreated, "url successfully created")
}

func (h *UrlHandler) GetUrlByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	user, err := h.service.GetUrlByID(ctx, id.String())
	if err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WriteDataResponse(ctx, http.StatusOK, user)
}

func (h *UrlHandler) UpdateUrl(ctx *gin.Context) {
	var request dto.UpdateUrlRequest
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

	if err := h.service.UpdateUrl(ctx, request); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WriteMessageResponse(ctx, http.StatusOK, "url successfully updated")
}

func (h *UrlHandler) DeleteUrl(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	if err := h.service.DeleteUrl(ctx, id); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WriteMessageResponse(ctx, http.StatusOK, "url successfully deleted")
}

func (h *UrlHandler) CountAllUrl(ctx *gin.Context) {
	count, err := h.service.Count(ctx)
	if err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WriteDataResponse(ctx, http.StatusOK, count)
}
