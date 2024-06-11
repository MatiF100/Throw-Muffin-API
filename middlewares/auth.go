package middlewares

import (
	tokenService "github.com/MatiF100/Throw-Muffin-API/services"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		authorizationHeader := context.GetHeader("Authorization")
		prefix_len := len("Bearer ")
		if authorizationHeader == "" || len(authorizationHeader) <= prefix_len {
			context.JSON(401, gin.H{"error": "Authorization header required"})
			context.Abort()
			return
		}
		tokenString := authorizationHeader[prefix_len:]
		_, err := tokenService.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		context.Next()
	}
}
