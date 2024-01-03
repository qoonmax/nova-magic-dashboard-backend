package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	res, err := h.GenerateService.Generate(&req)

	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": res,
		"php":     req.ClientEnvironment.PhpVersion,
		"laravel": req.ClientEnvironment.LaravelVersion,
		"nova":    req.ClientEnvironment.NovaVersion,
		"dbms":    req.ClientEnvironment.DbmsName,
	})
}
