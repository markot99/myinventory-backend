// The middleware package implements functions that can be executed before and after a request.
package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/api"
	"github.com/markot99/myinventory-backend/models"
	"github.com/markot99/myinventory-backend/pkg/database"
	"github.com/markot99/myinventory-backend/pkg/logger"
	"github.com/markot99/myinventory-backend/utils/ginutils"
)

// ItemIDMiddlewareParam is the name of the parameter that is read and set by the ItemIDMiddleware
const ItemIDMiddlewareParam = "itemID"

// ItemIDMiddleware checks if the user has submitted a valid item ID.
// The item is retrieved from the database. If the user is not the owner of the item,
// the request is canceled with the status forbidden (403). If the item ID is not valid,
// the request is canceled with the status bad request (400). If the database is not available,
// the request is canceled with the status internal server error (500).
func ItemIDMiddleware(itemTable database.ItemTable) gin.HandlerFunc {
	return func(c *gin.Context) {
		// check parameter
		itemID, err := ginutils.GetParameter(c, ItemIDMiddlewareParam)
		if err != nil {
			logger.Debug("Failed to get item id parameter: ", err.Error())
			api.SetBadRequestParameterItemIDFaulty(c)
			return
		}

		// check if item exists
		item, err := itemTable.GetItem(itemID)
		if errors.Is(err, database.ErrNothingFound) || errors.Is(err, database.ErrGeneratingObjectIDFailed) {
			logger.Debug("Failed to find item in the database: ", err.Error())
			api.SetBadRequestItemNotFound(c)
			return
		}
		if err != nil {
			logger.Error("Failed to get item from database: ", err.Error())
			api.SetInternalServerDatabaseError(c)
			return
		}

		if item.OwnerID != GetAuthMiddlewareClaims(c).ID {
			logger.Warn("User tried to access item that does not belong to him")
			api.SetForbiddenItemAccess(c)
			return
		}

		c.Set(ItemIDMiddlewareParam, item)
		c.Next()
	}
}

// GetItemIDMiddlewareItem can be used to retrieve the item extracted by the ItemIDMiddleware
func GetItemIDMiddlewareItem(c *gin.Context) models.Item {
	return c.MustGet(ItemIDMiddlewareParam).(models.Item)
}
