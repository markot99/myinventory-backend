// The package ginutils is used to provider helper utils for the gin library
package ginutils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrHeaderNotFound = errors.New("header not found")
var ErrHeaderHasMoreThanOneValue = errors.New("header has more than one value")
var ErrHeaderIsEmpty = errors.New("header is empty")
var ErrParameterNotFound = errors.New("parameter not found")
var ErrParameterIsEmpty = errors.New("parameter is empty")

// GetHeader can be used to retrieve a header from the gin context.
func GetHeader(c *gin.Context, key string) (string, error) {
	item, exists := c.Request.Header[key]

	if !exists {
		return "", ErrHeaderNotFound
	}

	if len(item) != 1 {
		return "", ErrHeaderHasMoreThanOneValue
	}

	if item[0] == "" {
		return "", ErrHeaderIsEmpty
	}

	return item[0], nil
}

// GetParameter can be used to retrieve a url query parameter from the gin context.
func GetParameter(c *gin.Context, key string) (string, error) {
	value, exists := c.Params.Get(key)

	if !exists {
		return "", ErrParameterNotFound
	}

	if value == "" {
		return "", ErrParameterIsEmpty
	}

	return value, nil
}
