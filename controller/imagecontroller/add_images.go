package imagecontroller

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/api"
	"github.com/markot99/myinventory-backend/controller/middleware"
	"github.com/markot99/myinventory-backend/models"
	"github.com/markot99/myinventory-backend/pkg/logger"
	"github.com/markot99/myinventory-backend/utils/fileutils"
)

const UPLOAD_IMAGES_PARAM_ITEM_ID = "id"

type ImagesForm struct {
	Images []*multipart.FileHeader `form:"images" binding:"required"`
}

type AddImagesResponseBody struct {
	UploadedImages []models.File `json:"uploadedImages"`
}

// AddImages, uses authentication and itemaccess middleware
//
//	@Summary      Upload images
//	@Description  Upload one ore more images to the server and get their internal ids
//	@Security     JWT
//	@Tags         images
//	@Accept       multipart/form-data
//	@Param        itemID path string true "item id"
//	@Param        images formData []file true "images of the item"
//	@Success      204
//	@Failure      400 {object} api.APIErrorResponse
//	@Failure      401 {object} api.APIErrorResponse
//	@Failure      500 {object} api.APIErrorResponse
//	@Router       /v1/items/{itemID}/images [post]
func (controller ImageController) AddImages(c *gin.Context) {
	item := middleware.GetItemIDMiddlewareItem(c)

	var imagesForm ImagesForm
	if err := c.Bind(&imagesForm); err != nil {
		logger.Debug("Error binding images form: ", err.Error())
		api.SetBadRequestWrongBody(c)
		return
	}

	// check if files were sent
	if len(imagesForm.Images) == 0 {
		logger.Debug("No files were sent")
		api.SetBadRequestWrongFileType(c, "jpg, jpeg, png")
		return
	}

	// check file types
	for _, file := range imagesForm.Images {
		if !fileutils.IsImage(file.Filename) {
			logger.Debug("Wrong file type: ", file.Filename)
			api.SetBadRequestWrongFileType(c, "jpg, jpeg, png")
			return
		}
	}

	// save files and create objects for the database
	imageObjects := []models.File{}
	for _, file := range imagesForm.Images {
		openedFile, err := file.Open()
		if err != nil {
			logger.Error("Error opening file: ", err.Error())
			api.SetInternalServerGeneralError(c)
			return
		}
		defer openedFile.Close()

		fileBytes, _ := io.ReadAll(openedFile)
		id, err := controller.ImageStorage.SaveFile(fileBytes)
		if err != nil {
			logger.Error("Error saving file: ", err.Error())
			api.SetInternalServerStorageError(c)
			return
		}

		imageObjects = append(imageObjects, models.File{ID: id, FileName: file.Filename})
	}

	// save images to the database
	err := controller.ItemTable.AddImages(item.ID, imageObjects)
	if err != nil {
		logger.Error("Error saving images to the database: ", err.Error())
		api.SetInternalServerDatabaseError(c)
		return
	}

	var responseBody AddImagesResponseBody
	responseBody.UploadedImages = append(responseBody.UploadedImages, imageObjects...)

	c.JSON(http.StatusOK, responseBody)
}
