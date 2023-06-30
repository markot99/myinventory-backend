package main

import (
	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/controller/usercontroller"
	"github.com/markot99/myinventory-backend/pkg/authenticator/authjwt"
	"github.com/markot99/myinventory-backend/pkg/database/mongodb"
	"github.com/markot99/myinventory-backend/pkg/logger"
	"github.com/markot99/myinventory-backend/utils/envutils"
)

func main() {
	logger.Info("Starting server...")
	logger.InitializeLogger()

	mongodb_host := envutils.GetEnvVariableOrDefault("MONGO_DB_HOST", "mongodb://localhost:27017")
	mongodb_database := envutils.GetEnvVariableOrDefault("MONGO_DB_DATABASE", "myinventory")
	port := envutils.GetEnvVariableOrDefault("PORT", "8080")

	usercollection := envutils.GetEnvVariableOrDefault("USER_COLLECTION", "users")

	authSecret := envutils.GetEnvVariableOrDefault("JWT_SECRET", "defaultsecret")

	router := gin.Default()
	routes := router.Group("/api")

	mongodbConnection, err := mongodb.NewMongoDBConnection(mongodb_host, mongodb_database)
	if err != nil {
		panic(err)
	}
	defer mongodbConnection.Disconnect()

	userTable := mongodb.NewUserTable(mongodbConnection, usercollection)
	authenticator := authjwt.CreateJWTAuthenticator([]byte(authSecret))

	usercontroller.RegisterRoutes(routes, authenticator, userTable)

	router.Run(":" + port)
	logger.Info("Bye!")
}
