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

func (controller *UserController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := controller.usecase.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong ...",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (controller *UserController) Create(ctx *gin.Context) {
	input := usecase.CreateUserInput{}
	if err := ctx.Bind(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong ...",
		})
		return
	}

	if err := controller.usecase.Create(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong ...",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}
