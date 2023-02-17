package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		err, exists := ctx.Get("error")
		if exists {
			eVal, ok := err.(error)
			if ok {
				ctx.JSON(http.StatusOK, gin.H{
					"error": eVal.Error(),
				})
			}
		}
	}
}
