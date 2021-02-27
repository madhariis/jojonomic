package routers

import (
	c "document-service/services"
	"github.com/gin-gonic/gin"
)

func Document(route *gin.RouterGroup) {
	document := route.Group("/document")
	{
		document.POST("", c.CreateUpdateDocument)
		document.GET("/:document_id", c.DetailDocument)
		document.DELETE("", c.DeleteDocument)
	}
}
