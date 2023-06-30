package itemcontroller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/api"
	"github.com/markot99/myinventory-backend/controller/middleware"
	"github.com/markot99/myinventory-backend/models"
	"github.com/markot99/myinventory-backend/pkg/logger"
)

type ItemEditPurchaseInfo struct {
	Date      string  `json:"date" binding:"required"`
	Place     string  `json:"place" binding:"required"`
	UnitPrice float64 `json:"unitPrice"`
	Quantity  int     `json:"quantity"`
}

type ItemEditRequestBody struct {
	Name         string               `json:"name" binding:"required"`
	Description  string               `json:"description" binding:"required"`
	PurchaseInfo ItemEditPurchaseInfo `json:"purchaseInfo" binding:"required"`
}

// EditItem, uses authentication and itemaccess middleware
//
//	@Summary      Edit an item
//	@Description  Edit an item in the database
//	@Security     JWT
//	@Tags         items
//	@Accept       json
//	@Param        itemID path string true "item id"
//	@Param        item body AddItemRequestBody true "Item Body"
//	@Success      201 {object} AddItemResponseBody
//	@Failure      400 {object} api.APIErrorResponse
//	@Failure      401 {object} api.APIErrorResponse
//	@Failure      500 {object} api.APIErrorResponse
//	@Router       /v1/items/ [put]
func (controller ItemController) EditItem(c *gin.Context) {
	item := middleware.GetItemIDMiddlewareItem(c)

	body := ItemEditRequestBody{}
	if err := c.BindJSON(&body); err != nil {
		logger.Debug("Failed to bind edit item json: ", err.Error())
		api.SetBadRequestWrongBody(c)
		return
	}

	var puchaseDate time.Time

	if body.PurchaseInfo.Date != "" {
		var err error
		puchaseDate, err = models.GetBuyDateFromString(body.PurchaseInfo.Date)
		if err != nil {
			logger.Debug("Failed to parse buy date from string: ", puchaseDate)
			api.SetBadRequestWrongDateFormat(c)
			return
		}

		if puchaseDate.After(time.Now()) {
			logger.Debug("Buy date is after current date: ", puchaseDate)
			api.SetBadRequestPurchaseDateInFuture(c)
			return
		}
	}

	item.Name = body.Name
	item.Description = body.Description
	item.PurchaseInfo.Date = puchaseDate
	item.PurchaseInfo.Place = body.PurchaseInfo.Place
	item.PurchaseInfo.UnitPrice = body.PurchaseInfo.UnitPrice
	item.PurchaseInfo.Quantity = body.PurchaseInfo.Quantity

	err := controller.ItemTable.EditItem(item)
	if err != nil {
		logger.Error("Failed to edit item in database: ", err.Error())
		api.SetInternalServerDatabaseError(c)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
