package middleware

import (
	"net/http"

	"example.com/blog/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			ctx.Abort()
			return
		}
		user, err := utils.ParseJWT(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}
		ctx.Set("username", user.Username)
		ctx.Set("userID", user.ID)

		// fmt.Println("username", user.Username)
		// fmt.Println("userID", user.ID)

		ctx.Next()

	}
}
