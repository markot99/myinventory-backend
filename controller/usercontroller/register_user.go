package usercontroller

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/api"
	"github.com/markot99/myinventory-backend/models"
	"github.com/markot99/myinventory-backend/pkg/logger"
	"github.com/markot99/myinventory-backend/utils/passwordutils"
)

// RegisterUserRequestBody is the struct that contains the necessary data for registering a user
type RegisterUserRequestBody struct {
	FirstName string `json:"firstName" binding:"required"` // first name of the user
	LastName  string `json:"lastName" binding:"required"`  // last name of the user
	Email     string `json:"email" binding:"required"`     // email of the user
	Password  string `json:"password" binding:"required"`  // unhashed password of the user
}

// validateEmail validates the email format
func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// RegisterUser
//
//  @Summary      Register user
//  @Description  Register a new user
//  @Tags         users
//  @Accept       json
//  @Param        item body RegisterUserRequestBody true "Registration Body"
//  @Success      204
//  @Failure      400 {object} api.APIErrorResponse
//  @Failure      409 {object} api.APIErrorResponse
//  @Failure      500 {object} api.APIErrorResponse
//  @Router       /v1/users/register [post]
func (controller UserController) RegisterUser(c *gin.Context) {
	body := RegisterUserRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		logger.Debug("Failed to bind register user json: ", err.Error())
		api.SetBadRequestWrongBody(c)
		return
	}

	if !validateEmail(body.Email) {
		logger.Debug("Failed to validate email: ", body.Email)
		api.SetBadRequestWrongEmailFormat(c)
		return
	}

	_, err := controller.UserTable.GetUserByEmail(body.Email)
	if err == nil {
		logger.Debug("User already exists: ", body.Email)
		api.SetConflictUserAlreadyExists(c)
		return
	}

	hashedPassword, err := passwordutils.HashPassword(&body.Password)
	if err != nil {
		logger.Error("Failed to hash password: ", err.Error())
		api.SetInternalServerGeneralError(c)
		return
	}

	user := models.User{FirstName: body.FirstName, LastName: body.LastName, Email: body.Email, Password: hashedPassword}
	err = controller.UserTable.AddUser(user)
	if err != nil {
		logger.Error("Failed to add user: ", err.Error())
		api.SetInternalServerDatabaseError(c)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
