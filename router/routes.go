package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gsxhnd/BangumiServer/controller"
)

// Load loads the middleware, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(mw...)
	g.GET("/code", controller.GetCode)
	g.GET("/wait_login", controller.WaitLogin)

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.JSON(404, struct {
			Message string `json:"message"`
		}{
			Message: "The incorrect API route.",
		})
	})

	if gin.Mode() == "debug" {
	}

	return g
}
