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

type AddItemPurchaseInfo struct {
	Date      string  `json:"date"`
	Place     string  `json:"place"`
	UnitPrice float64 `json:"unitPrice"`
	Quantity  int     `json:"quantity"`
}

type AddItemRequestBody struct {
	Name         string              `json:"name" binding:"required"`
	Description  string              `json:"description"`
	PurchaseInfo AddItemPurchaseInfo `json:"purchaseInfo" binding:"required"`
}

type AddItemResponseBody struct {
	ID string `json:"id"`
}

// AddItem, uses authentication middleware
//
//	@Summary      Add an item
//	@Description  Add an item to the database
//	@Security     JWT
//	@Tags         items
//	@Accept	      json
//	@Param        item body AddItemRequestBody true "Item Body"
//	@Success      201 {object} AddItemResponseBody
//	@Failure      400 {object} api.APIErrorResponse
//	@Failure      401 {object} api.APIErrorResponse
//	@Failure      500 {object} api.APIErrorResponse
//	@Router       /v1/items [post]
func (controller ItemController) AddItem(c *gin.Context) {
	claims := middleware.GetAuthMiddlewareClaims(c)

	body := AddItemRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		logger.Debug("Failed to bind add item json: ", err.Error())
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

	purchaseInfo := models.PurchaseInfo{
		Date:      puchaseDate,
		Place:     body.PurchaseInfo.Place,
		UnitPrice: body.PurchaseInfo.UnitPrice,
		Quantity:  body.PurchaseInfo.Quantity,
		Invoice:   models.File{},
	}

	images := models.Images{PreviewImage: "", Images: []models.File{}}
	item := models.Item{Name: body.Name, Description: body.Description, PurchaseInfo: purchaseInfo, OwnerID: claims.ID, Images: images}

	itemID, err := controller.ItemTable.AddItem(item)
	if err != nil {
		logger.Error("Failed to add item to database: ", err.Error())
		api.SetInternalServerDatabaseError(c)
		return
	}

	response := AddItemResponseBody{ID: itemID}

	c.JSON(http.StatusCreated, response)
}
