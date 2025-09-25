package middlewares

import (
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "error parsing token.",
		})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
