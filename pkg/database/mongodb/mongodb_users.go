package mongodb

import (
	"context"
	"errors"

	"github.com/markot99/myinventory-backend/models"
	"github.com/markot99/myinventory-backend/pkg/database"
	"github.com/markot99/myinventory-backend/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserTable is the struct that contains the collection for the users
type UserTable struct {
	collection *mongo.Collection // collection for the users
}

// NewUserTable is the constructor for creating a UserTable object
func NewUserTable(connection MongoDBConnection, collectionID string) *UserTable {
	collection := connection.database.Collection(collectionID)
	return &UserTable{collection: collection}
}

// AddUser is used to add an user to the collection. When inserting, a new user id is automatically created.
func (table_p *UserTable) AddUser(user models.User) error {
	_, err := table_p.collection.InsertOne(context.TODO(), user)
	if err != nil {
		logger.Error("Failed to insert user to database: ", err.Error())
		return database.ErrInsertFailed
	}

	return nil
}

// GetUser is used to get an user from the collection by user id.
func (table_p *UserTable) GetUserByID(userID string) (*models.User, error) {
	var user models.User

	objectid, err := stringToObjectID(userID)
	if err != nil {
		return nil, database.ErrGeneratingObjectIDFailed
	}

	filter := bson.M{"_id": objectid}
	err = table_p.collection.FindOne(context.TODO(), filter).Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		logger.Debug("No user found with id ", userID)
		return nil, database.ErrNothingFound
	}
	if err != nil {
		logger.Error("Failed to get user by id: ", err.Error())
		return nil, database.ErrGetFailed
	}
	return &user, nil
}

// GetUser is used to get an user from the collection by user email.
func (table_p *UserTable) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	filter := bson.M{"email": email}
	err := table_p.collection.FindOne(context.TODO(), filter).Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		logger.Debug("No user found with email ", email)
		return nil, database.ErrNothingFound
	}
	if err != nil {
		logger.Error("Failed to get user by email: ", err.Error())
		return nil, database.ErrGetFailed
	}

	return &user, nil
}

// DeleteCollection is used to delete the complete collection of users.
func (table_p *UserTable) DeleteCollection() error {
	err := table_p.collection.Drop(context.Background())
	if err != nil {
		logger.Error("Failed to delete user collection: ", err.Error())
		return database.ErrDeleteCollectionFailed
	}
	return nil
}
