package invoicecontroller

import (
	"io"
	"mime/multipart"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/api"
	"github.com/markot99/myinventory-backend/controller/middleware"
	"github.com/markot99/myinventory-backend/models"
	"github.com/markot99/myinventory-backend/pkg/logger"
	"github.com/markot99/myinventory-backend/utils/fileutils"
)

type InvoiceForm struct {
	Invoice *multipart.FileHeader `form:"invoice" binding:"required"`
}

// AddInvoice, uses authentication and itemaccess middleware
//
//	@Summary      Upload invoice
//	@Description  Upload invoice to the item
//	@Security     JWT
//	@Tags         invoice
//	@Accept       multipart/form-data
//	@Param        itemID path string true "item id"
//	@Param        invoice formData file true "invoice of the item"
//	@Success      204
//	@Failure      400 {object} api.APIErrorResponse
//	@Failure      401 {object} api.APIErrorResponse
//	@Failure      500 {object} api.APIErrorResponse
//	@Router       /v1/items/{itemID}/invoice [post]
func (controller InvoiceController) AddInvoice(c *gin.Context) {
	item := middleware.GetItemIDMiddlewareItem(c)

	// check if item already has invoice
	if reflect.DeepEqual(item.PurchaseInfo.Invoice, models.PurchaseInfo{}) {
		logger.Debug("Item already has invoice")
		api.SetBadRequestItemHasInvoice(c)
		return
	}

	// validate pdf file
	var invoiceForm InvoiceForm
	if err := c.Bind(&invoiceForm); err != nil {
		logger.Debug("Error binding invoice form: ", err.Error())
		api.SetBadRequestWrongBody(c)
		return
	}
	if !fileutils.IsPdf(invoiceForm.Invoice.Filename) {
		logger.Debug("Wrong file type: ", invoiceForm.Invoice.Filename)
		api.SetBadRequestWrongFileType(c, "pdf")
		return
	}

	// read file
	openedFile, err := invoiceForm.Invoice.Open()
	if err != nil {
		logger.Error("Failed to open file: ", err.Error())
		api.SetInternalServerGeneralError(c)
		return
	}
	defer openedFile.Close()
	fileBytes, err := io.ReadAll(openedFile)
	if err != nil {
		logger.Error("Failed to read file: ", err.Error())
		api.SetInternalServerGeneralError(c)
		return
	}

	// save file
	id, err := controller.InvoiceStorage.SaveFile(fileBytes)
	if err != nil {
		logger.Error("Failed to save file: ", err.Error())
		api.SetInternalServerStorageError(c)
		return
	}

	// add file to item
	invoiceOb := models.File{ID: id, FileName: invoiceForm.Invoice.Filename}
	err = controller.ItemTable.SetInvoice(item.ID, invoiceOb)
	if err != nil {
		logger.Error("Failed to set invoice in the database: ", err.Error())
		api.SetInternalServerDatabaseError(c)
		return
	}

	// return
	c.JSON(http.StatusNoContent, nil)
}
