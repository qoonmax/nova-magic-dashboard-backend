package routers

import (
	"github.com/gin-gonic/gin"
	"magic-dashboard-backend/internal/http/handlers"
)

type Router struct {
	Handler *handlers.Handler
}

func NewRouter(handler *handlers.Handler) *Router {
	return &Router{
		Handler: handler,
	}
}

func (r *Router) InitRoutes() *gin.Engine {
	ginEngine := gin.New()

	api := ginEngine.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/generate", r.Handler.Generate)
		}
	}

	return ginEngine
}
