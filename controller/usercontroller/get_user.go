package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/api"
	"github.com/markot99/myinventory-backend/controller/middleware"
	"github.com/markot99/myinventory-backend/pkg/logger"
)

type MeResponseBody struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// GetUser, uses authentication middleware
//
//	@Summary      Get user
//	@Description  Get user info about the logged in user
//	@Tags         users
//	@Security     JWT
//	@Success      200 {object} MeResponseBody
//	@Failure      401 {object} api.APIErrorResponse
//	@Failure      500 {object} api.APIErrorResponse
//	@Router       /v1/users/me [get]
func (controller UserController) GetUser(c *gin.Context) {
	claims := middleware.GetAuthMiddlewareClaims(c)

	user, err := controller.UserTable.GetUserByID(claims.ID)
	if err != nil {
		logger.Error("Failed to get user from database:", err.Error())
		api.SetInternalServerDatabaseError(c)
		return
	}

	response := MeResponseBody{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName, Email: user.Email}

	c.JSON(http.StatusOK, response)
}
