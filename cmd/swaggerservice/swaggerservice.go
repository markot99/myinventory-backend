package main

import (
	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/pkg/logger"
	"github.com/markot99/myinventory-backend/utils/envutils"

	_ "github.com/markot99/myinventory-backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	logger.Info("Starting server...")
	logger.InitializeLogger()

	port := envutils.GetEnvVariableOrDefault("PORT", "8080")
	router := gin.Default()
	routes := router.Group("/api")

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":" + port)

	logger.Info("Bye!")
}
