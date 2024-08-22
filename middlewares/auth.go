package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/plug-pathomgphong/golang-rest/utils"
)

func Authentication(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authenticate."})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authenticate."})
		return
	}

	context.Set("userId", userId)

	context.Next()
}
