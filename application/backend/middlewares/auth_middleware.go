package middlewares

import (
	"carlife-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get token from header
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}
		// Get user from token
		userID, userType, err := utils.ParseJWT(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		// Set user to context
		ctx.Set("userID", userID)
		ctx.Set("userType", userType)
		ctx.Next()
	}
}
