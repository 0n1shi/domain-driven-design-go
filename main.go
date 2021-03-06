package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "github.com/go-redis/redis"
	yaml "gopkg.in/yaml.v3"

	"github.com/0n1shi/domain-driven-design/controller"
	domainUser "github.com/0n1shi/domain-driven-design/domain/user"

	mysqlRepo "github.com/0n1shi/domain-driven-design/infra/repository/mysql"
	// redisRepo "github.com/0n1shi/domain-driven-design/infra/repository/redis"
	"github.com/0n1shi/domain-driven-design/usecase"
)

type Config struct {
	MySQL MySQLConfig `yaml:"mysql"`
	Redis RedisConfig `yaml:"redis"`
}

type MySQLConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DB       string `yaml:"db"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DB       int    `yaml:"db"`
	Password string `yaml:"password"`
}

func main() {
	// config
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "", "config file path")
	flag.Parse()
	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(fmt.Sprintf("failed to read config: %s", err.Error()))
	}
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal config: %s", err.Error()))
	}

	// setup for MySQL
	dbConfig := config.MySQL
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to open connection to db: %s", err.Error()))
	}
	db.AutoMigrate(&mysqlRepo.User{})

	// setup for redis
	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr:     config.Redis.Host + ":" + config.Redis.Port,
	// 	Password: config.Redis.Password, // no password set
	// 	DB:       config.Redis.DB,       // use default DB
	// })

	// setup components
	userRepository := mysqlRepo.NewUserRepository(db)
	// userRepository := redisRepo.NewUserRepository(redisClient)
	userService := domainUser.NewUserService(userRepository)
	userUsecase := usecase.NewUserUsecase(userService)
	userController := controller.NewUserController(userUsecase)

	// setup router
	router := gin.Default()
	router.Use(controller.AbortError())
	users := router.Group("/users")
	{
		users.GET("", userController.FindAll)
		users.GET("/:id", userController.FindByID)
		users.POST("", userController.Create)
	}
	router.Run()
}
