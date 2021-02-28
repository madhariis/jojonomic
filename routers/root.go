package routers

import (
	"document-service/controllers"

	"github.com/gin-gonic/gin"
)

func Root(route *gin.RouterGroup) {
	root := route.Group("/")
	{
		root.GET("", controllers.GetAll)
	}
}
