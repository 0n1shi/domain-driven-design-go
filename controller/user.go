package controller

import (
	"net/http"

	"github.com/0n1shi/domain-driven-design/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecase *usecase.UserUsecase
}

func NewUserController(usecase *usecase.UserUsecase) *UserController {
	return &UserController{usecase: usecase}
}

func (controller *UserController) FindAll(ctx *gin.Context) {
	users, err := controller.usecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong ...",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
