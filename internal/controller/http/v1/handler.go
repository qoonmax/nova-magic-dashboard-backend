package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type GenerateServiceInterface interface {
	Generate(req *Request) (string, error)
}

type Handler struct {
	GenerateService GenerateServiceInterface
}

func NewHandler(service GenerateServiceInterface) *Handler {
	return &Handler{
		GenerateService: service,
	}
}

func (h *Handler) Generate(ctx *gin.Context) {
	var req Request

	if err := ctx.BindJSON(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.GenerateService.Generate(&req)

	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Mock response data
	instructions := map[string]interface{}{
		"for_creating": []map[string]interface{}{
			{
				"path":    "app/Models/Post.php",
				"content": "<?php echo \"" + time.Now().Format("2006-01-02 15:04:05") + "\";",
			},
			{
				"path":    "app/Nova/Post.php",
				"content": "<?php echo \"" + time.Now().Format("2006-01-02 15:04:05") + "\";",
			},
		},
	}

	ctx.JSON(http.StatusOK, instructions)
}
