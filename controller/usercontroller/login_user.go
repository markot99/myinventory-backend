package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/api"
	"github.com/markot99/myinventory-backend/models"
	"github.com/markot99/myinventory-backend/pkg/logger"
	"github.com/markot99/myinventory-backend/utils/passwordutils"
)

type LoginUserRequestBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponseBody struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Token     string `json:"token"`
}

// LoginUser, uses authentication middleware
//
//  @Summary      Login user
//  @Description  Login user and retrieve an authorization token
//  @Tags         users
//  @Accept       json
//  @Param        item body LoginUserRequestBody true "Login Body"
//  @Success      200 {object} LoginUserResponseBody
//  @Failure      400 {object} api.APIErrorResponse
//  @Failure      401 {object} api.APIErrorResponse
//  @Failure      500 {object} api.APIErrorResponse
//  @Router       /v1/users/login [post]
func (controller UserController) LoginUser(c *gin.Context) {
	body := LoginUserRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		logger.Debug("Failed to bind login user json: ", err.Error())
		api.SetBadRequestWrongBody(c)
		return
	}

	user, err := controller.UserTable.GetUserByEmail(body.Email)
	if err != nil {
		logger.Warn("User tried to login with non existing email: " + body.Email)
		api.SetUnauthorizedLoginFailed(c)
		return
	}

	if !passwordutils.CheckPassword(&user.Password, &body.Password) {
		logger.Warn("User tried to login with wrong password: " + body.Email)
		api.SetUnauthorizedLoginFailed(c)
		return
	}

	authClaims := models.AuthenticationClaims{ID: user.ID}
	token, err := controller.Authenticator.GenerateToken(&authClaims)
	if err != nil {
		logger.Error("Failed to generate token: ", err.Error())
		api.SetInternalServerGeneralError(c)
		return
	}

	response := LoginUserResponseBody{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Token: token}

	c.JSON(http.StatusOK, response)
}
