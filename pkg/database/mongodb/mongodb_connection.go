// The package mongodb is used to integrate the mongodb library
package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/markot99/myinventory-backend/pkg/database"
	"github.com/markot99/myinventory-backend/pkg/logger"
)

// MongoDBConnection is the struct that contains the endpoint of the mongodb and the client for the mongodb database
type MongoDBConnection struct {
	endpoint string          // endpoint url of the mongodb
	client   *mongo.Client   // mongodb client
	database *mongo.Database // client for the mongodb database
}

// NewMongoDBConnection is the constructor for creating a MongoDBConnection object.
func NewMongoDBConnection(endpoint string, host string) (MongoDBConnection, error) {
	clientOptions := options.Client().ApplyURI(endpoint)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logger.Error("Failed to connect to mongodb: ", err.Error())
		return MongoDBConnection{}, database.ErrDatabaseConnection
	}

	database := client.Database(host)
	p := MongoDBConnection{endpoint: endpoint, client: client, database: database}

	return p, nil
}

// DropDatabase is used to delete the database.
func (connection_p *MongoDBConnection) DropDatabase() error {
	err := connection_p.database.Drop(context.Background())
	if err != nil {
		logger.Error("Failed to drop database: ", err.Error())
		return database.ErrDatabaseDrop
	}
	return nil
}

// Disconnect is used to disconnect from the mongodb database.
func (connection_p *MongoDBConnection) Disconnect() error {
	err := connection_p.client.Disconnect(context.Background())
	if err != nil {
		logger.Error("Failed to disconnect from mongodb: ", err.Error())
		return database.ErrDatabaseConnection
	}
	return nil
}

// stringToObjectID is used to generate the mongodb specific object id from a string.
func stringToObjectID(id string) (primitive.ObjectID, error) {
	objectid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Debug("Failed to generate object id from string: ", err.Error())
		return primitive.ObjectID{}, database.ErrGeneratingObjectIDFailed
	}
	return objectid, nil
}

// objectIDToString is used to generate a string from a mongodb specific object id.
func objectIDToString(objectID primitive.ObjectID) string {
	return objectID.Hex()
}
