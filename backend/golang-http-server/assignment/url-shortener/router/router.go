package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/handlers"
)

func SetupRouter(urlHandler handlers.URLHandler) *gin.Engine {
	// insiasi router
	r := gin.Default()

	// set routes
	r.GET("/:path", urlHandler.Get)
	r.POST("/", urlHandler.Create)
	r.POST("/custom", urlHandler.CreateCustom)

	// return
	return r
	// return &gin.Engine{} // TODO: replace this
}
