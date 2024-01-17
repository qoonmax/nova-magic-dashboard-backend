package handlers

import (
	"github.com/gin-gonic/gin"
	"magic-dashboard-backend/internal/http/requests"
	"magic-dashboard-backend/internal/http/responses"
	"net/http"
)

func (h *Handler) Generate(ctx *gin.Context) {
	var req requests.GenerateRequest

	if err := ctx.BindJSON(&req); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	instructions, err := h.service.GenerationService.Generate(&req)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"for_creating": instructions})
}
