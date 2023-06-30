package itemcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/api"
	"github.com/markot99/myinventory-backend/controller/middleware"
	"github.com/markot99/myinventory-backend/pkg/logger"
)

type ResponseItem struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Quantity     int    `json:"quantity"`
	PreviewImage string `json:"previewImage"`
}

type GetItemsResponseBody struct {
	ResponseItems []ResponseItem `json:"items"`
}

// GetItems, uses authentication middleware
//
//	@Summary      Get a list of items
//	@Description  Get a list of items belongig to the user with reduced information
//	@Security     JWT
//	@Tags         items
//	@Accept       json
//	@Success      200 {object} GetItemsResponseBody
//	@Failure      401 {object} api.APIErrorResponse
//	@Failure      500 {object} api.APIErrorResponse
//	@Router       /v1/items [get]
func (controller ItemController) GetItems(c *gin.Context) {
	claims := middleware.GetAuthMiddlewareClaims(c)

	items, err := controller.ItemTable.GetItems(claims.ID)
	if err != nil {
		logger.Error("Failed to get items from database: ", err.Error())
		api.SetInternalServerDatabaseError(c)
		return
	}

	responseItems := GetItemsResponseBody{ResponseItems: []ResponseItem{}}

	for _, item := range items {
		responseItem := ResponseItem{ID: item.ID, Name: item.Name, Description: item.Description, Quantity: item.PurchaseInfo.Quantity, PreviewImage: item.Images.PreviewImage}
		responseItems.ResponseItems = append(responseItems.ResponseItems, responseItem)
	}

	c.JSON(http.StatusOK, &responseItems)
}
