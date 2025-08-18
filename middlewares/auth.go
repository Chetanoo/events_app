package middlewares

import (
	"events_app/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		return
	}
	ctx.Set("userId", userId)
	ctx.Next()
}
