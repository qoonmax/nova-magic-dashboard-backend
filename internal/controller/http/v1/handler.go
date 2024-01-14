package v1

import (
	"github.com/gin-gonic/gin"
	"magic-dashboard-backend/internal/controller/http/v1/request"
	"magic-dashboard-backend/internal/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) Generate(ctx *gin.Context) {
	var req request.GenerateRequest

	if err := ctx.BindJSON(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	pctx, err := h.service.GenerationService.Generate(&req)

	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"for_creating": pctx.Instructions})
}
