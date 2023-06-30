package mongodb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

const TestDBName = "myinventory_test"
const TestDBHost = "mongodb://localhost:27017"

func TestMongoDBConnectAndDisconnect_Valid(t *testing.T) {
	connection, err := NewMongoDBConnection(TestDBHost, TestDBName)
	assert.NoError(t, err)

	// Query the server status to test the connection
	var commandResult bson.M
	command := bson.D{{Key: "serverStatus", Value: 1}}
	err = connection.database.RunCommand(context.TODO(), command).Decode(&commandResult)
	assert.NoError(t, err)

	err = connection.Disconnect()
	assert.NoError(t, err)
}

func TestInvalidMongoDBConnectAndDisconnect(t *testing.T) {
	_, err := NewMongoDBConnection("mongodb://localhost:270172", "test")
	assert.Error(t, err)
}

func TestCorrectObjectID(t *testing.T) {
	_, err := stringToObjectID("649819427ee9b564b50615ce")
	assert.NoError(t, err)
}

func TestWrongObjectID(t *testing.T) {
	_, err := stringToObjectID("123")
	assert.Error(t, err)
}
