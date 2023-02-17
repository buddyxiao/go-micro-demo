package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		data, exists := ctx.Get("data")
		if exists {
			ctx.JSON(http.StatusOK, gin.H{
				"data": data,
			})
		}
	}
}
