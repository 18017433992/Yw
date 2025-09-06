package middleware

import "github.com/gin-gonic/gin"

func Somehandler(ctx *gin.Context) (uint, bool) {
	userID, exists := ctx.Get("userID")
	if !exists {
		return 0, false
	}

	id, ok := userID.(uint)
	return id, ok

}
