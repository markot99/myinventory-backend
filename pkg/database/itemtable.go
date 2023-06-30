package database

import (
	"github.com/markot99/myinventory-backend/models"
)

// ItemTable is the interface that wraps the basic methods for storing and retrieving item data
type ItemTable interface {
	// AddItem is used to store an item in the database
	AddItem(item models.Item) (string, error)
	// EditItem is used to delete the complete collection of items
	EditItem(item models.Item) error
	// DeleteItem is used to delete an item from the database
	DeleteItem(id string) error
	// GetItem is used to retrieve an item from the database by its id
	GetItem(id string) (models.Item, error)
	// GetItems is used to retrieve all items from the database that belong to a user
	GetItems(userID string) ([]models.Item, error)
	// AddImages is used to add images to an item
	AddImages(itemID string, images []models.File) error
	// DeleteItemImage is used to delete an image of an item
	DeleteItemImage(itemID string, imageID string) error
	// SetPreviewImage is used to set the preview image of an item
	SetPreviewImage(itemID string, previewImage string) error
	// SetInvoice is used to set the invoice of an item
	SetInvoice(itemID string, invoice models.File) error
	// DeleteItemInvoice is used to delete the invoice of an item
	DeleteItemInvoice(itemID string) error
}
