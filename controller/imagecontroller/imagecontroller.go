// The package itemcontroller is used to provide the items rest api
package imagecontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/controller/middleware"
	"github.com/markot99/myinventory-backend/pkg/authenticator"
	"github.com/markot99/myinventory-backend/pkg/database"
	"github.com/markot99/myinventory-backend/pkg/storage"
)

type ImageController struct {
	ItemTable    database.ItemTable
	ImageStorage storage.Storage
}

const PARAM_ITEM_ID = "itemID"
const PARAM_IMAGE_NAME = "imageID"

func RegisterRoutes(routes *gin.RouterGroup, authenticator authenticator.Authenticator, itemTable database.ItemTable, imageStorage storage.Storage) {
	h := &ImageController{
		ItemTable:    itemTable,
		ImageStorage: imageStorage,
	}

	itemIDParamName := middleware.ItemIDMiddlewareParam
	imageIDParamName := middleware.ImageIDMiddlewareParam
	itemIDMiddleware := middleware.ItemIDMiddleware(itemTable)
	imageIDMiddleware := middleware.ImageIDMiddleware()

	authMiddleware := middleware.AuthMiddleware(authenticator)

	imageRoutes := routes.Group("")
	imageRoutes.Use(authMiddleware)

	imageRoutes.DELETE("/v1/items/:"+itemIDParamName+"/images/:"+imageIDParamName, itemIDMiddleware, imageIDMiddleware, h.DeleteImage)
	imageRoutes.GET("/v1/items/:"+itemIDParamName+"/images/:"+imageIDParamName, itemIDMiddleware, imageIDMiddleware, h.GetImage)
	imageRoutes.POST("/v1/items/:"+itemIDParamName+"/images", itemIDMiddleware, h.AddImages)
	imageRoutes.PUT("/v1/items/:"+itemIDParamName+"/images/preview", itemIDMiddleware, h.SetPreviewImage)

}
