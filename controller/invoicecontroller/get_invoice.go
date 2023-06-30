package invoicecontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/api"
	"github.com/markot99/myinventory-backend/controller/middleware"
	"github.com/markot99/myinventory-backend/pkg/logger"
)

// GetInvoice, uses authentication and itemaccess middleware
//
//	@Summary      Get invoice
//	@Description  Get the invoice from the item
//	@Security     JWT
//	@Tags         invoice
//	@Accept       json
//	@Param        itemID path string true "item id"
//	@Success      200 {file} file "Invoice"
//	@Failure      400 {object} api.APIErrorResponse
//	@Failure      401 {object} api.APIErrorResponse
//	@Failure      404 {object} api.APIErrorResponse
//	@Failure      500 {object} api.APIErrorResponse
//	@Router       /v1/items/{itemID}/invoice [get]
func (controller InvoiceController) GetInvoice(c *gin.Context) {
	item := middleware.GetItemIDMiddlewareItem(c)

	if item.PurchaseInfo.Invoice.ID == "" {
		logger.Debug("Item has no invoice")
		api.SetBadRequestItemHasNoInvoice(c)
		return
	}

	// get file from storage
	file, err := controller.InvoiceStorage.GetFile(item.PurchaseInfo.Invoice.ID)
	if err != nil {
		logger.Error("Failed to get invoice file:", err.Error())
		api.SetInternalServerStorageError(c)
		return
	}

	// set necessary headers for pdf file
	mimeType := "application/pdf"
	c.Header("Content-Disposition", "attachment; filename="+item.PurchaseInfo.Invoice.FileName)
	c.Header("Content-Type", mimeType)
	c.Data(http.StatusOK, mimeType, file)
}
