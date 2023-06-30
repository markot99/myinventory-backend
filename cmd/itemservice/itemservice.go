package main

import (
	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/controller/imagecontroller"
	"github.com/markot99/myinventory-backend/controller/invoicecontroller"
	"github.com/markot99/myinventory-backend/controller/itemcontroller"
	"github.com/markot99/myinventory-backend/pkg/authenticator/authjwt"
	"github.com/markot99/myinventory-backend/pkg/database/mongodb"
	"github.com/markot99/myinventory-backend/pkg/logger"
	"github.com/markot99/myinventory-backend/pkg/storage/localstorage"
	"github.com/markot99/myinventory-backend/utils/envutils"
)

func main() {
	logger.Info("Starting server...")
	logger.InitializeLogger()

	mongodb_host := envutils.GetEnvVariableOrDefault("MONGO_DB_HOST", "mongodb://localhost:27017")
	mongodb_database := envutils.GetEnvVariableOrDefault("MONGO_DB_DATABASE", "myinventory")
	port := envutils.GetEnvVariableOrDefault("PORT", "8080")

	itemcollection := envutils.GetEnvVariableOrDefault("ITEM_COLLECTION", "items")
	imageStoragePath := envutils.GetEnvVariableOrDefault("IMAGE_STORAGE_PATH", "images")
	invoiceStoragePath := envutils.GetEnvVariableOrDefault("INVOICE_STORAGE_PATH", "invoices")

	authSecret := envutils.GetEnvVariableOrDefault("JWT_SECRET", "defaultsecret")

	router := gin.Default()
	routes := router.Group("/api")

	mongodbConnection, err := mongodb.NewMongoDBConnection(mongodb_host, mongodb_database)
	if err != nil {
		logger.Fatal("Failed to connect to mongodb", err.Error())
	}
	defer mongodbConnection.Disconnect()

	itemTable := mongodb.NewItemTable(mongodbConnection, itemcollection)
	imageStorage := localstorage.NewLocalStorage(imageStoragePath)
	invoiceStorage := localstorage.NewLocalStorage(invoiceStoragePath)
	authenticator := authjwt.CreateJWTAuthenticator([]byte(authSecret))

	imagecontroller.RegisterRoutes(routes, authenticator, itemTable, imageStorage)
	itemcontroller.RegisterRoutes(routes, authenticator, itemTable, imageStorage, invoiceStorage)
	invoicecontroller.RegisterRoutes(routes, authenticator, itemTable, invoiceStorage)

	router.Run(":" + port)
	logger.Info("Bye!")
}
