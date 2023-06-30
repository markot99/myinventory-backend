package invoicecontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/api"
	"github.com/markot99/myinventory-backend/controller/middleware"
	"github.com/markot99/myinventory-backend/pkg/logger"
)

// DeleteInvoice, uses authentication and itemaccess middleware
//
//	@Summary      Delete invoice
//	@Description  Delete the invoice from the item and from the storage
//	@Security     JWT
//	@Tags         invoice
//	@Accept       json
//	@Param        itemID path string true "item id"
//	@Success      204
//	@Failure      400 {object} api.APIErrorResponse
//	@Failure      401 {object} api.APIErrorResponse
//	@Failure      500 {object} api.APIErrorResponse
//	@Router       /v1/items/{itemID}/invoice [delete]
func (controller InvoiceController) DeleteInvoice(c *gin.Context) {
	item := middleware.GetItemIDMiddlewareItem(c)

	err := controller.ItemTable.DeleteItemInvoice(item.ID)
	if err != nil {
		logger.Error("Failed to delete invoice from database:", err.Error())
		api.SetInternalServerDatabaseError(c)
		return
	}

	err = controller.InvoiceStorage.DeleteFile(item.PurchaseInfo.Invoice.ID)
	if err != nil {
		logger.Error("Failed to delete invoice file:", err.Error())
		api.SetInternalServerStorageError(c)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
