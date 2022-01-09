package controller

import (
	"net/http"

	"github.com/0n1shi/domain-driven-design/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	usecase   *usecase.UserUsecase
	validator *validator.Validate
}

func NewUserController(usecase *usecase.UserUsecase) *UserController {
	return &UserController{
		usecase:   usecase,
		validator: validator.New(),
	}
}

func (controller *UserController) FindAll(ctx *gin.Context) {
	users, err := controller.usecase.FindAll()
	if err != nil {
		SetError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (controller *UserController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := controller.usecase.FindByID(id)
	if err != nil {
		SetError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (controller *UserController) Create(ctx *gin.Context) {
	input := usecase.CreateUserInput{}
	if err := ctx.Bind(&input); err != nil {
		SetError(ctx, err)
		return
	}
	if err := controller.validator.Struct(input); err != nil {
		SetErrorPublic(ctx, err)
		return
	}
	user, err := controller.usecase.Create(&input)
	if err != nil {
		SetError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
