package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LogError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		if query != "" {
			path = path + "?" + query
		}
		ctx.Next()
		errMsg := ctx.Errors.String()
		statusCode := ctx.Writer.Status()
		latency := time.Since(start)
		method := ctx.Request.Method
		clientIP := ctx.ClientIP()

		var event *zerolog.Event
		switch {
		case statusCode >= 400 && statusCode < 500:
			event = log.Warn()
		case statusCode >= 500:
			event = log.Error()
		default:
			event = log.Info()
		}
		event.Str("method", method).Str("path", path).Dur("response_time", latency).Int("status", statusCode).Str("client_ip", clientIP).Msg(errMsg)
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
