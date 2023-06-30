// The authenticator package is used to include libraries for creating and parsing tokens.
package authenticator

import (
	"errors"

	"github.com/markot99/myinventory-backend/models"
)

var ErrTokenInvalid = errors.New("token invalid")
var ErrGeneratingToken = errors.New("error generating token")

// Authenticator is the interface that wraps the basic methods for creating and parsing tokens
type Authenticator interface {
	// GenerateToken is used to create a token from the given AuthenticationClaims
	GenerateToken(authClaims *models.AuthenticationClaims) (string, error)
	// ValidateToken is used to parse a token and return the AuthenticationClaims
	ValidateToken(token string) (models.AuthenticationClaims, error)
}
