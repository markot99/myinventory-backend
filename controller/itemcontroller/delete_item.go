package itemcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/api"
	"github.com/markot99/myinventory-backend/controller/middleware"
	"github.com/markot99/myinventory-backend/pkg/logger"
)

const DELETE_ITEM_PARAM_ITEM_ID = "id"

// DeleteItem, uses authentication and itemaccess middleware
//
//	@Summary      Delete an item
//	@Description  Delete an item from the database including invoice and all images
//	@Security     JWT
//	@Tags         items
//	@Accept       json
//	@Param        itemID path string true "item id"
//	@Success      204
//	@Failure      400 {object} api.APIErrorResponse
//	@Failure      401 {object} api.APIErrorResponse
//	@Failure      500 {object} api.APIErrorResponse
//	@Failure      500
//	@Router       /v1/items/{itemID} [delete]
func (controller ItemController) DeleteItem(c *gin.Context) {
	item := middleware.GetItemIDMiddlewareItem(c)

	// delete all images
	for _, image := range item.Images.Images {
		if image.ID == "" {
			err := controller.ImageStorage.DeleteFile(image.ID)
			if err != nil {
				// continue despite error
				logger.Error("Failed to delete image with id", image.ID, ", error: ", err.Error())
			}
		}
	}

	// delete invoice
	if item.PurchaseInfo.Invoice.ID != "" {
		err := controller.InvoiceStorage.DeleteFile(item.PurchaseInfo.Invoice.ID)
		if err != nil {
			// continue despite error
			logger.Error("Failed to delete invoice with id", item.PurchaseInfo.Invoice.ID, ", error: ", err.Error())
		}
	}

	err := controller.ItemTable.DeleteItem(item.ID)
	if err != nil {
		logger.Error("Failed to delete item from database:", err.Error())
		api.SetInternalServerDatabaseError(c)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
