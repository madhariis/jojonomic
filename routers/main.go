package routers

import (
	"document-service/helper"
	"github.com/gin-gonic/gin"
)

func Router(route *gin.Engine) {
	routes := route.Group("/document-service")
	{
		routes.Use(helper.JwtAuth)
		Root(routes)
		Folder(routes)
		Document(routes)
	}
}
