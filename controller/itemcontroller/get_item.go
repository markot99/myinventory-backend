package itemcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/controller/middleware"
)

const GET_ITEM_PARAM_ITEM_ID = "id"

// GetItem, uses authentication and itemaccess middleware
//
//	@Summary      Get an item
//	@Description  Get detailed information about an item from the database
//	@Security     JWT
//	@Tags         items
//	@Accept       json
//	@Param        id path string true "Object id"
//	@Success      200 {object} models.Item
//	@Failure      400 {object} api.APIErrorResponse
//	@Failure      401 {object} api.APIErrorResponse
//	@Failure      500 {object} api.APIErrorResponse
//	@Router       /v1/items/{id} [get]
func (controller ItemController) GetItem(c *gin.Context) {
	item := middleware.GetItemIDMiddlewareItem(c)
	c.JSON(http.StatusOK, &item)
}
