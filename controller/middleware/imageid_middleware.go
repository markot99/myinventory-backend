// The middleware package implements functions that can be executed before and after a request.
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/api"
	"github.com/markot99/myinventory-backend/models"
	"github.com/markot99/myinventory-backend/pkg/logger"
	"github.com/markot99/myinventory-backend/utils/ginutils"
)

// ImageIDMiddlewareParam is the name of the parameter that is read and set by the ImageIDMiddleware
const ImageIDMiddlewareParam = "imageID"

// ImageIDMiddleware checks if the user has submitted a valid image ID.
// The image file gets extracted from the item. If the image ID is not valid,
// the request is canceled with the status bad request (400).
//
// ItemIDMiddleware must be called beforehand.
func ImageIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		imageID, err := ginutils.GetParameter(c, ImageIDMiddlewareParam)
		if err != nil {
			logger.Debug("Failed to get image id parameter: ", err.Error())
			api.SetBadRequestParameterImageIDFaulty(c)
			return
		}

		item := GetItemIDMiddlewareItem(c)

		image, err := models.GetImageFileById(item, imageID)
		if err != nil {
			logger.Debug("Failed to get image file by id: ", err.Error())
			api.SetBadRequestImageDoesNotExist(c)
			return
		}

		c.Set(ImageIDMiddlewareParam, image)
		c.Next()
	}
}

// GetImageIDMiddlewareImageFile can be used to retrieve the image file extracted by the ImageIDMiddleware
func GetImageIDMiddlewareImageFile(c *gin.Context) models.File {
	return c.MustGet(ImageIDMiddlewareParam).(models.File)
}
