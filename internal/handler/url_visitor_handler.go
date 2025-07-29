package handler

import (
	"net/http"

	"github.com/Alfian57/belajar-golang/internal/response"
	"github.com/Alfian57/belajar-golang/internal/service"
	"github.com/gin-gonic/gin"
)

type UrlVisitorHandler struct {
	service *service.UrlVisitorService
}

func NewUrlVisitorHandler(s *service.UrlVisitorService) *UrlVisitorHandler {
	return &UrlVisitorHandler{
		service: s,
	}
}

func (h *UrlVisitorHandler) Count(ctx *gin.Context) {
	count, err := h.service.Count(ctx)
	if err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WriteDataResponse(ctx, http.StatusOK, count)
}

func (h *UrlVisitorHandler) CountByID(ctx *gin.Context) {
	urlIDParam := ctx.Param("urlID")

	count, err := h.service.CountByUrlID(ctx, urlIDParam)
	if err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WriteDataResponse(ctx, http.StatusOK, count)
}
