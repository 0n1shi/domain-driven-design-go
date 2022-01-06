package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

func AbortError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.ByType(gin.ErrorTypePublic).Last()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error_message": fmt.Sprintf("error: %s", err.Err),
			})
		}

		err = ctx.Errors.ByType(gin.ErrorTypePrivate).Last()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error_message": "internal server error",
			})
		}
	}
}
