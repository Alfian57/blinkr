package handler

import (
	"net/http"

	"github.com/Alfian57/belajar-golang/internal/response"
	"github.com/Alfian57/belajar-golang/internal/service"
	"github.com/gin-gonic/gin"
)

type UrlHandler struct {
	service *service.UrlService
}

func NewUrlHandler(s *service.UrlService) *UrlHandler {
	return &UrlHandler{
		service: s,
	}
}

func (h *UrlHandler) CountAllUrl(ctx *gin.Context) {
	count, err := h.service.Count(ctx)
	if err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WriteDataResponse(ctx, http.StatusOK, count)
}
