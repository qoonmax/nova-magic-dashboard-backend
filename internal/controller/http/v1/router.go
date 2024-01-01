package v1

import "github.com/gin-gonic/gin"

type HandlerIterface interface {
	Generate(*gin.Context)
}

type Router struct {
	Handler HandlerIterface
}

func NewRouter(handler HandlerIterface) *Router {
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
