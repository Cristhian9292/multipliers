package main

import (
	"fmt"
	"trafilea/controllers"
	"trafilea/internal"

	_ "trafilea/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	var collection = []int{}
	fmt.Println(collection)
	internal.CreateList()
}

// @title Multipliers
// @version 1.0
// @description A Multipliers Service API in Go using Gin framework

// @host localhost:8080
// @BasePath /
func main() {
	router := gin.Default()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/healthcheck", controllers.HealthCheck)
	router.POST("/multipliers", controllers.Save)
	router.GET("/multipliers", controllers.Get)
	router.GET("/multipliers/collection", controllers.GetAll)

	router.Run("0.0.0.0:8080")
}
