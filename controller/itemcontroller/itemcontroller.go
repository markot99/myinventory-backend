// The package itemcontroller is used to provide the items rest api
package itemcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/controller/middleware"
	"github.com/markot99/myinventory-backend/pkg/authenticator"
	"github.com/markot99/myinventory-backend/pkg/database"
	"github.com/markot99/myinventory-backend/pkg/storage"
)

type ItemController struct {
	ItemTable      database.ItemTable
	ImageStorage   storage.Storage
	InvoiceStorage storage.Storage
}

const PARAM_ITEM_ID = "itemID"
const PARAM_IMAGE_NAME = "imageID"

func RegisterRoutes(routes *gin.RouterGroup, authenticator authenticator.Authenticator, itemTable database.ItemTable,
	imageStorage storage.Storage, invoiceStorage storage.Storage) {
	h := &ItemController{
		ItemTable:      itemTable,
		ImageStorage:   imageStorage,
		InvoiceStorage: invoiceStorage,
	}

	itemIDParamName := middleware.ItemIDMiddlewareParam
	itemIDMiddleware := middleware.ItemIDMiddleware(itemTable)

	authMiddleware := middleware.AuthMiddleware(authenticator)

	itemRoutes := routes.Group("")
	itemRoutes.Use(authMiddleware)

	itemRoutes.POST("/v1/items", h.AddItem)
	itemRoutes.GET("/v1/items", h.GetItems)
	itemRoutes.DELETE("/v1/items/:"+itemIDParamName, itemIDMiddleware, h.DeleteItem)
	itemRoutes.PUT("/v1/items/:"+itemIDParamName, itemIDMiddleware, h.EditItem)
	itemRoutes.GET("/v1/items/:"+itemIDParamName, itemIDMiddleware, h.GetItem)
}
