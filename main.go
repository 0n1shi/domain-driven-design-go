package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/0n1shi/domain-driven-design/controller"
	domainUser "github.com/0n1shi/domain-driven-design/domain/user"
	"github.com/0n1shi/domain-driven-design/infra/db"
	"github.com/0n1shi/domain-driven-design/usecase"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	userRepository := db.NewDBUserRepository()
	userService := domainUser.NewUserService(userRepository)
	userUsecase := usecase.NewUserUsecase(userService)
	userController := controller.NewUserController(userUsecase)

	users := router.Group("/users")
	{
		users.GET("", userController.FindAll)
	}

	router.Run()
}
