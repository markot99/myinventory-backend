package imagecontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/api"
	"github.com/markot99/myinventory-backend/controller/middleware"
	"github.com/markot99/myinventory-backend/models"
	"github.com/markot99/myinventory-backend/pkg/logger"
)

const SET_PREVIEW_IMAGE_PARAM_ITEM_ID = "id"

type SetPreviewImageRequestBody struct {
	Name string `json:"name" binding:"required"`
}

// SetPreviewImage, uses authentication and itemaccess middleware
//
//	@Summary      Set preview image
//	@Description  Set preview image for an item
//	@Security     JWT
//	@Tags         images
//	@Accept       json
//	@Param        itemID path string true "item id"
//	@Param        item body SetPreviewImageRequestBody true "Set Preview Image Name"
//	@Success      204
//	@Failure      400 {object} api.APIErrorResponse
//	@Failure      401 {object} api.APIErrorResponse
//	@Failure      500 {object} api.APIErrorResponse
//	@Router       /v1/items/{itemID}/images/preview [put]
func (controller ImageController) SetPreviewImage(c *gin.Context) {
	item := middleware.GetItemIDMiddlewareItem(c)

	body := SetPreviewImageRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		logger.Debug("Failed to bind json: ", err.Error())
		api.SetBadRequestWrongBody(c)
		return
	}

	// image does not exist on item
	image, err := models.GetImageFileById(item, body.Name)
	if err != nil {
		logger.Debug("Failed to get image file by id: ", err.Error())
		api.SetBadRequestImageDoesNotExist(c)
		return
	}

	err = controller.ItemTable.SetPreviewImage(item.ID, image.ID)
	if err != nil {
		logger.Error("Failed to set preview image: ", err.Error())
		api.SetInternalServerDatabaseError(c)
	}

	c.JSON(http.StatusNoContent, nil)
}
