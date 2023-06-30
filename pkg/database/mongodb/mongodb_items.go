package mongodb

import (
	"context"
	"errors"

	"github.com/markot99/myinventory-backend/models"
	"github.com/markot99/myinventory-backend/pkg/database"
	"github.com/markot99/myinventory-backend/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ItemTable is the struct that contains the collection for the items
type ItemTable struct {
	collection *mongo.Collection // collection for the items
}

// NewItemTable is the constructor for creating a ItemTable object
func NewItemTable(connection MongoDBConnection, collectionID string) *ItemTable {
	collection := connection.database.Collection(collectionID)
	return &ItemTable{collection: collection}
}

// AddItem is used to add an item to the collection. When inserting, a new item id is automatically created.
// If successful, the id of the inserted item is returned.
func (table_p *ItemTable) AddItem(item models.Item) (string, error) {
	result, err := table_p.collection.InsertOne(context.TODO(), item)
	if err != nil {
		logger.Error("Failed to insert item to database: ", err.Error())
		return "", database.ErrInsertFailed
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		logger.Error("Failed to convert inserted id to object id")
		return "", database.ErrGettingObjectIDFailed
	}

	return objectIDToString(oid), nil
}

// DeleteItem is used to delete an item from the database.
func (table_p *ItemTable) DeleteItem(itemID string) error {
	objectid, err := stringToObjectID(itemID)
	if err != nil {
		return database.ErrDeleteFailed
	}

	filter := bson.M{"_id": objectid}
	deleteResult, err := table_p.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		logger.Error("Failed to delete item from database: ", err.Error())
		return database.ErrDeleteFailed
	}

	if deleteResult.DeletedCount == 0 {
		logger.Debug("No item deleted. Item with id ", itemID, " not found.")
		return database.ErrDeleteFailed
	}

	return nil
}

// GetItem is used to get an item from the collection.
func (table_p *ItemTable) GetItem(itemID string) (models.Item, error) {
	var item models.Item

	objectid, err := stringToObjectID(itemID)
	if err != nil {
		return models.Item{}, database.ErrGeneratingObjectIDFailed
	}

	filter := bson.M{"_id": objectid}
	err = table_p.collection.FindOne(context.TODO(), filter).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Debug("No item found with id ", itemID)
			return models.Item{}, database.ErrNothingFound
		}
		logger.Error("Failed to get item from database: ", err.Error())
		return models.Item{}, database.ErrGetFailed
	}
	return item, nil
}

// GetItems is used to get all items from the collection.
func (table_p *ItemTable) GetItems(userID string) ([]models.Item, error) {
	filter := bson.M{"ownerID": userID}
	cursor, err := table_p.collection.Find(context.TODO(), filter)
	if err != nil {
		logger.Error("Failed to get items from database: ", err.Error())
		return []models.Item{}, database.ErrGetFailed
	}
	defer cursor.Close(context.TODO())

	var items []models.Item
	if err := cursor.All(context.Background(), &items); err != nil {
		logger.Error("Failed to get items from database: ", err.Error())
		return []models.Item{}, database.ErrGetFailed
	}

	return items, nil
}

// AddImages is used to add images to an item from the collection.
func (table_p *ItemTable) AddImages(itemID string, images []models.File) error {
	objectid, err := stringToObjectID(itemID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectid}
	update := bson.M{"$push": bson.M{"images.images": bson.M{"$each": images}}}

	result, err := table_p.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		logger.Error("Failed to add images to item: ", err.Error())
		return database.ErrUpdateFailed
	}

	if result.MatchedCount == 0 {
		logger.Debug("No item updated. Item with id ", itemID, " not found.")
		return database.ErrNothingFound
	}

	return nil
}

// SetPreviewImage is used to set the preview image of an item from the collection.
func (table_p *ItemTable) SetPreviewImage(itemID string, previewImage string) error {
	objectid, err := stringToObjectID(itemID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectid}
	update := bson.M{"$set": bson.M{"images.previewImage": previewImage}}

	result, err := table_p.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Failed to set preview image of item: ", err.Error())
		return database.ErrUpdateFailed
	}

	if result.MatchedCount == 0 {
		logger.Debug("No item updated. Item with id ", itemID, " not found.")
		return database.ErrNothingFound
	}

	return nil
}

// DeleteCollection is used to delete the complete collection of items.
func (table_p *ItemTable) DeleteCollection() error {
	err := table_p.collection.Drop(context.Background())
	if err != nil {
		logger.Error("Failed to delete collection: ", err.Error())
		return database.ErrDeleteCollectionFailed
	}
	return nil
}

// SetInvoice is used to set the invoice of an item from the collection.
func (table_p *ItemTable) SetInvoice(itemID string, invoice models.File) error {
	objectid, err := stringToObjectID(itemID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectid}
	update := bson.M{"$set": bson.M{"purchaseInfo.invoice": invoice}}

	result, err := table_p.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Failed to set invoice of item: ", err.Error())
		return database.ErrUpdateFailed
	}

	if result.MatchedCount == 0 {
		logger.Debug("No item updated. Item with id ", itemID, " not found.")
		return database.ErrNothingFound
	}

	return nil
}

// DeleteItemInvoice is used to delete the invoice of an item from the collection.
func (table_p *ItemTable) DeleteItemInvoice(itemID string) error {
	objectid, err := stringToObjectID(itemID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectid}
	update := bson.M{"$unset": bson.M{"purchaseInfo.invoice": ""}}

	result, err := table_p.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Failed to delete invoice of item: ", err.Error())
		return database.ErrUpdateFailed
	}

	if result.MatchedCount == 0 {
		logger.Debug("No item updated. Item with id ", itemID, " not found.")
		return database.ErrNothingFound
	}

	return nil
}

// DeleteItemImage is used to delete an image of an item from the collection.
func (table_p *ItemTable) DeleteItemImage(itemID string, imageID string) error {
	objectid, err := stringToObjectID(itemID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectid}
	update := bson.M{"$pull": bson.M{"images.images": bson.M{"id": imageID}}}

	result, err := table_p.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Failed to delete image of item: ", err.Error())
		return database.ErrUpdateFailed
	}

	if (result.MatchedCount == 0) || (result.ModifiedCount == 0) {
		logger.Debug("No item updated. Item with id ", itemID, " not found.")
		return database.ErrNothingFound
	}

	return nil
}

// EditItem is used to edit an item from the collection.
func (table_p *ItemTable) EditItem(item models.Item) error {
	objectid, err := stringToObjectID(item.ID)
	if err != nil {
		return err
	}

	// remove item id to prevent write concern error
	replaceItem := item
	replaceItem.ID = ""

	filter := bson.M{"_id": objectid}
	result, err := table_p.collection.ReplaceOne(context.Background(), filter, replaceItem)
	if err != nil {
		logger.Error("Failed to edit item: ", err.Error())
		return database.ErrUpdateFailed
	}

	if (result.MatchedCount == 0) || (result.ModifiedCount == 0) {
		logger.Debug("No item updated. Item with id ", item.ID, " not found.")
		return database.ErrNothingFound
	}

	return nil
}
