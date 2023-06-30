// The package itemcontroller is used to provide the items rest api
package invoicecontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/controller/middleware"
	"github.com/markot99/myinventory-backend/pkg/authenticator"
	"github.com/markot99/myinventory-backend/pkg/database"
	"github.com/markot99/myinventory-backend/pkg/storage"
)

type InvoiceController struct {
	ItemTable      database.ItemTable
	InvoiceStorage storage.Storage
}

func RegisterRoutes(routes *gin.RouterGroup, authenticator authenticator.Authenticator, itemTable database.ItemTable, invoiceStorage storage.Storage) {
	h := &InvoiceController{
		ItemTable:      itemTable,
		InvoiceStorage: invoiceStorage,
	}

	itemIDParamName := middleware.ItemIDMiddlewareParam
	itemIDMiddleware := middleware.ItemIDMiddleware(itemTable)

	authMiddleware := middleware.AuthMiddleware(authenticator)

	invoiceRoutes := routes.Group("")
	invoiceRoutes.Use(authMiddleware)

	invoiceRoutes.POST("/v1/items/:"+itemIDParamName+"/invoice", itemIDMiddleware, h.AddInvoice)
	invoiceRoutes.DELETE("/v1/items/:"+itemIDParamName+"/invoice", itemIDMiddleware, h.DeleteInvoice)
	invoiceRoutes.GET("/v1/items/:"+itemIDParamName+"/invoice", itemIDMiddleware, h.GetInvoice)
}
