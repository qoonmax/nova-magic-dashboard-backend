package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GenerateServiceInterface interface {
	Generate() (string, error)
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

	res, err := h.GenerateService.Generate()

	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Hello",
		"response": res,
		"laravel":  req.LaravelVersion,
		"nova":     req.NovaVersion,
		"key":      req.Key,
	})
}
