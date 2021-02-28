package routers

import (
	"document-service/middleware"
	"github.com/gin-gonic/gin"
)

func Router(route *gin.Engine) {
	routes := route.Group("/document-service")
	{
		routes.Use(middleware.JwtAuth)
		Root(routes)
		Folder(routes)
		Document(routes)
	}
}
