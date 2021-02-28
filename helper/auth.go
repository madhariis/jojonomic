package helper

import (
	"document-service/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func JwtAuth(c *gin.Context) {
	res, err := TokenValidation(c.GetHeader("Authorization"), os.Getenv("SECRET_KEY"))
	if err != nil {
		res := Response(true, "Unauthorized", map[string]interface{}{})
		c.JSON(http.StatusUnauthorized, res)
		c.Abort()
		return
	}
	models.UserToken = res
	return
}
