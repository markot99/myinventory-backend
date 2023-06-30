// The package usercontroller is used to provide the users rest api
package usercontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/controller/middleware"
	"github.com/markot99/myinventory-backend/pkg/authenticator"
	"github.com/markot99/myinventory-backend/pkg/database"
)

// UserController is the struct that contains the necessary dependencies
type UserController struct {
	Authenticator authenticator.Authenticator
	UserTable     database.UserTable
}

// RegisterRoutes registers the routes for the users api
func RegisterRoutes(routes *gin.RouterGroup, authenticator authenticator.Authenticator, userTable database.UserTable) {
	h := &UserController{
		Authenticator: authenticator,
		UserTable:     userTable,
	}

	authMiddleware := middleware.AuthMiddleware(authenticator)

	routes.POST("/v1/users/login", h.LoginUser)
	routes.POST("/v1/users/register", h.RegisterUser)
	routes.GET("/v1/users/me", authMiddleware, h.GetUser)
}
