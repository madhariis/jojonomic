package middleware

import (
	"document-service/helper"
	"document-service/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func JwtAuth(c *gin.Context) {
	res, err := helper.TokenValidation(c.GetHeader("Authorization"), os.Getenv("SECRET_KEY"))
	if err != nil {
		res := helper.Response(true, "Unauthorized", map[string]interface{}{})
		c.JSON(http.StatusUnauthorized, res)
		c.Abort()
		return
	}
	models.UserToken = res
	return
}
