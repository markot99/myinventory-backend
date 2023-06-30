package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIErrorResponse represents the response for an error
type APIErrorResponse struct {
	DiagnosisCode    uint32 `json:"diagnosisCode"`   // diagnosis code of the error
	DiagnosisMessage string `json:"diagnosisMessage"` // diagnosis message of the error
}

// bad request errors

func SetBadRequestParameterItemIDFaulty(c *gin.Context) {
	response := APIErrorResponse{
		DiagnosisCode:    1001,
		DiagnosisMessage: "The parameter itemID is missing or faulty.",
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, &response)
}

func SetBadRequestItemNotFound(c *gin.Context) {
	response := APIErrorResponse{
		DiagnosisCode:    1002,
		DiagnosisMessage: "No item exists with the given itemID.",
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, &response)
}

func SetBadRequestParameterImageIDFaulty(c *gin.Context) {
	response := APIErrorResponse{
		DiagnosisCode:    1004,
		DiagnosisMessage: "The parameter imageID is missing or faulty",
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, &response)
}

func SetBadRequestImageDoesNotExist(c *gin.Context) {
	response := APIErrorResponse{
		DiagnosisCode:    1003,
		DiagnosisMessage: "The provided imageID does not exist on the item.",
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, &response)
}

func SetBadRequestWrongBody(c *gin.Context) {
	response := APIErrorResponse{
		DiagnosisCode:    1005,
		DiagnosisMessage: "The request body is faulty.",
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, &response)
}

func SetBadRequestItemHasInvoice(c *gin.Context) {
	response := APIErrorResponse{
		DiagnosisCode:    1006,
		DiagnosisMessage: "The item already has an invoice. Please delete the old invoice first.",
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, &response)
}

func SetBadRequestItemHasNoInvoice(c *gin.Context) {
	response := APIErrorResponse{
		DiagnosisCode:    1007,
		DiagnosisMessage: "The item has no invoice.",
	}
	c.AbortWithStatusJSON(http.StatusNotFound, &response)
}

func SetBadRequestWrongFileType(c *gin.Context, allowedFiles string) {
	response := APIErrorResponse{
		DiagnosisCode:    1008,
		DiagnosisMessage: "Only the following file types are allowed: '" + allowedFiles + "'.",
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, &response)
}

func SetBadRequestWrongDateFormat(c *gin.Context) {
	response := APIErrorResponse{
		DiagnosisCode:    1009,
		DiagnosisMessage: "The purchase date is formatted incorrectly. Format: yyyy-mm-dd.",
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, &response)
}

func SetBadRequestPurchaseDateInFuture(c *gin.Context) {
	response := APIErrorResponse{
		DiagnosisCode:    1010,
		DiagnosisMessage: "The purchase date is in the future.",
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, &response)
}

func SetBadRequestWrongEmailFormat(c *gin.Context) {
	response := APIErrorResponse{
		DiagnosisCode:    1011,
		DiagnosisMessage: "The email is not valid.",
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, &response)
}

// unauthorized errors

func SetUnauthorizedErrorTokenMissing(c *gin.Context) {
	response := APIErrorResponse{
		DiagnosisCode:    2001,
		DiagnosisMessage: "The authorization token is missing.",
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, &response)
}

func SetUnauthorizedErrorTokenFaulty(c *gin.Context) {
	response := APIErrorResponse{
		DiagnosisCode:    2002,
		DiagnosisMessage: "The authorization token is faulty or has expired.",
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, &response)
}

func SetUnauthorizedLoginFailed(c *gin.Context) {
	response := APIErrorResponse{
		DiagnosisCode:    2003,
		DiagnosisMessage: "The email or password is incorrect.",
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, &response)
}

// forbidden errors

func SetForbiddenItemAccess(c *gin.Context) {
	response := APIErrorResponse{
		DiagnosisCode:    3001,
		DiagnosisMessage: "The access to the item is forbidden.",
	}
	c.AbortWithStatusJSON(http.StatusForbidden, &response)
}

// confict errors

func SetConflictUserAlreadyExists(c *gin.Context) {
	response := APIErrorResponse{
		DiagnosisCode:    5001,
		DiagnosisMessage: "Registration failed because a user with the given email already exists.",
	}
	c.AbortWithStatusJSON(http.StatusConflict, &response)
}

// internal server errors

func SetInternalServerGeneralError(c *gin.Context) {
	response := APIErrorResponse{
		DiagnosisCode:    6001,
		DiagnosisMessage: "An server error has occurred while processing the request.",
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, &response)
}

func SetInternalServerDatabaseError(c *gin.Context) {
	response := APIErrorResponse{
		DiagnosisCode:    6001,
		DiagnosisMessage: "A database error occurred while processing the request.",
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, &response)
}

func SetInternalServerStorageError(c *gin.Context) {
	response := APIErrorResponse{
		DiagnosisCode:    6002,
		DiagnosisMessage: "A storage error occurred while processing the request.",
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, &response)
}
