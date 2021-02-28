package routers

import (
	"document-service/services"

	"github.com/gin-gonic/gin"
)

func Root(route *gin.RouterGroup) {
	root := route.Group("/")
	{
		root.GET("", services.GetAll)
	}
}
