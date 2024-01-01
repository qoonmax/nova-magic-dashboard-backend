package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MainServiceInterface interface {
	Run() (string, error)
}

type Handler struct {
	MainService MainServiceInterface
}

func NewHandler(service MainServiceInterface) *Handler {
	return &Handler{
		MainService: service,
	}
}

func (h *Handler) Generate(ctx *gin.Context) {
	if 1 > 2 {
		newErrorResponse(ctx, http.StatusInternalServerError, "Error in generate")
	}

	var req Request

	if err := ctx.BindJSON(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.MainService.Run()

	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "MainService error")
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":         "Hello",
		"response":        res,
		"laravel_version": req.LaravelVersion,
		"nova_version":    req.NovaVersion,
		"key":             req.Key,
	})
}
