package database

import (
	"github.com/markot99/myinventory-backend/models"
)

// UserTable is the interface that wraps the basic methods for storing and retrieving user data
type UserTable interface {
	// AddUser is used to store a user in the database
	AddUser(user models.User) error
	// GetUserByID is used to retrieve a user from the database by its id
	GetUserByID(userID string) (*models.User, error)
	// GetUserByEmail is used to retrieve a user from the database by its email
	GetUserByEmail(email string) (*models.User, error)
}
