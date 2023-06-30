package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/pkg/logger"
	"github.com/markot99/myinventory-backend/utils/envutils"
)

func main() {
	logger.Info("Starting server...")
	logger.InitializeLogger()

	port := envutils.GetEnvVariableOrDefault("PORT", "8080")

	router := gin.Default()

	router.StaticFS("/frontend", http.Dir("frontend"))

	router.Run(":" + port)

	logger.Info("Bye!")
}
