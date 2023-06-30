// The middleware package implements functions that can be executed before and after a request.
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/api"
	"github.com/markot99/myinventory-backend/models"
	"github.com/markot99/myinventory-backend/pkg/authenticator"
	"github.com/markot99/myinventory-backend/pkg/logger"
	"github.com/markot99/myinventory-backend/utils/ginutils"
)

// AuthMiddlewareParam is the name of the parameter that is set by the AuthMiddleware
const AuthMiddlewareParam = "claims"

// AuthMiddleware checks if the user provied a valid authentication token, and if so, extracts the claims.
// If the token is missing or faulty, the request is canceled with the status unauthorized (401).
func AuthMiddleware(auth authenticator.Authenticator) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := ginutils.GetHeader(c, "Authorization")
		if err != nil {
			logger.Debug("Failed to get authorization header: ", err.Error())
			api.SetUnauthorizedErrorTokenMissing(c)
			return
		}

		claims, err := auth.ValidateToken(token)
		if err != nil {
			logger.Warn("User tried to access with faulty token: ", err.Error())
			api.SetUnauthorizedErrorTokenFaulty(c)
			return
		}

		c.Set(AuthMiddlewareParam, claims)
		c.Next()
	}
}

// GetAuthMiddlewareClaims can be used to retrieve the claims extracted by the AuthMiddleware
func GetAuthMiddlewareClaims(c *gin.Context) models.AuthenticationClaims {
	return c.MustGet(AuthMiddlewareParam).(models.AuthenticationClaims)
}
