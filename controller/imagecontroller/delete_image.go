package imagecontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/api"
	"github.com/markot99/myinventory-backend/controller/middleware"
	"github.com/markot99/myinventory-backend/pkg/logger"
)

// DeleteImage, uses authentication, itemaccess and imageaccess middleware
//
//	@Summary      Delete an image
//	@Description  Delete an image from the item
//	@Security     JWT
//	@Tags         images
//	@Accept       json
//	@Param        itemID path string true "item id"
//	@Param        imageID path string true "image id"
//	@Success      200 {file} file "Image"
//	@Failure      400 {object} api.APIErrorResponse
//	@Failure      401 {object} api.APIErrorResponse
//	@Failure      500 {object} api.APIErrorResponse
//	@Router       /v1/items/{itemID}/images/{imageID} [delete]
func (controller ImageController) DeleteImage(c *gin.Context) {
	item := middleware.GetItemIDMiddlewareItem(c)
	image := middleware.GetImageIDMiddlewareImageFile(c)

	err := controller.ImageStorage.DeleteFile(image.ID)
	if err != nil {
		logger.Error("Failed to delete image file:", err.Error())
		api.SetInternalServerStorageError(c)
		return
	}

	err = controller.ItemTable.DeleteItemImage(item.ID, image.ID)
	if err != nil {
		logger.Error("Failed to delete image from database:", err.Error())
		api.SetInternalServerDatabaseError(c)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
