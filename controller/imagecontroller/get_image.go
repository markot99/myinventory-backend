package imagecontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/api"
	"github.com/markot99/myinventory-backend/controller/middleware"
	"github.com/markot99/myinventory-backend/pkg/logger"
	"github.com/markot99/myinventory-backend/utils/fileutils"
)

// GetImage, uses authentication, itemaccess and imageaccess middleware
//
//	@Summary      Get an image
//	@Description  Get an image from the item
//	@Security     JWT
//	@Tags         images
//	@Accept       json
//	@Param        itemID path string true "item id"
//	@Param        imageID path string true "image id"
//	@Success      200 {file} file "Image"
//	@Failure      400 {object} api.APIErrorResponse
//	@Failure      401 {object} api.APIErrorResponse
//	@Failure      500 {object} api.APIErrorResponse
//	@Router       /v1/items/{itemID}/images/{imageID} [get]
func (controller ImageController) GetImage(c *gin.Context) {
	image := middleware.GetImageIDMiddlewareImageFile(c)

	file, err := controller.ImageStorage.GetFile(image.ID)
	if err != nil {
		logger.Error("Failed to get image file:", err.Error())
		api.SetInternalServerStorageError(c)
		return
	}

	mimeType := "image/" + fileutils.GetFileType(image.FileName)

	c.Header("Content-Disposition", "attachment; filename="+image.FileName)
	c.Header("Content-Type", mimeType)

	c.Data(http.StatusOK, mimeType, file)
}
