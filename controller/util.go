package controller

import (
	"github.com/0n1shi/domain-driven-design/usecase"
	"github.com/gin-gonic/gin"
)

func SetError(ctx *gin.Context, err error) {
	errType := gin.ErrorTypePrivate
	if usecase.IsPublicErorr(err) {
		errType = gin.ErrorTypePublic
	}
	ctx.Error(err).SetType(errType)
}
