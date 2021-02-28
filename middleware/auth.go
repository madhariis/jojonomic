package middleware

import (
	"document-service/lib"
	"document-service/models"
	"document-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func JwtAuth(c *gin.Context) {
	res, err := lib.TokenValidation(c.GetHeader("Authorization"), os.Getenv("SECRET_KEY"))
	if err != nil {
		res := utils.Response(true, "Unauthorized", map[string]interface{}{})
		c.JSON(http.StatusUnauthorized, res)
		c.Abort()
		return
	}
	models.UserToken = res
	return
}
