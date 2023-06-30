package mongodb

import (
	"context"
	"testing"

	"github.com/markot99/myinventory-backend/models"
	"github.com/stretchr/testify/assert"
)

func setupUserTableConnection(t *testing.T) (*UserTable, func()) {
	connection, err := NewMongoDBConnection(TestDBHost, TestDBName)
	if err != nil {
		t.Fatalf(`Failed to establish database connection: %v`, err)
	}

	userTable := NewUserTable(connection, "users")

	// drop collection to start with empty collection
	userTable.collection.Drop(context.Background())

	return userTable, func() {
		connection.database.Drop(context.Background())
		connection.Disconnect()
	}
}

func getTestUser() models.User {
	return models.User{
		FirstName: "fistName",
		LastName:  "lastName",
		Email:     "email",
		Password:  "password",
	}
}

func TestUserAddAndRead(t *testing.T) {
	userTable, teardown := setupUserTableConnection(t)
	defer teardown()

	user := getTestUser()

	err := userTable.AddUser(user)
	assert.NoError(t, err)

	userRead, err := userTable.GetUserByEmail(user.Email)
	assert.NoError(t, err)

	assert.Equal(t, user.FirstName, userRead.FirstName)
	assert.Equal(t, user.LastName, userRead.LastName)
	assert.Equal(t, user.Email, userRead.Email)
	assert.Equal(t, user.Password, userRead.Password)
}

func TestDeleteUserCollection(t *testing.T) {
	usersTable, teardown := setupUserTableConnection(t)
	defer teardown()

	user := getTestUser()

	err := usersTable.AddUser(user)
	assert.NoError(t, err)

	err = usersTable.DeleteCollection()
	assert.NoError(t, err)

	_, err = usersTable.GetUserByEmail(user.Email)
	assert.Error(t, err)
}
