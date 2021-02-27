package routers

import (
	c "document-service/services"
	"github.com/gin-gonic/gin"
)

func Folder(route *gin.RouterGroup) {
	folder := route.Group("/folder")
	{
		folder.POST("", c.SetFolder)
		folder.GET("/:folder_id", c.DocumentByFolderID)
		folder.DELETE("", c.DeleteFolder)
	}
}
