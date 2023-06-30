package mongodb

import (
	"context"
	"reflect"
	"testing"

	"github.com/markot99/myinventory-backend/models"
	"github.com/markot99/myinventory-backend/pkg/database"
	"github.com/stretchr/testify/assert"
)

func setupItemsTableConnection(t *testing.T) (*ItemTable, func()) {
	connection, err := NewMongoDBConnection(TestDBHost, TestDBName)
	assert.NoError(t, err)

	itemTable := NewItemTable(connection, "items")

	// drop collection to start with empty collection
	itemTable.collection.Drop(context.Background())

	return itemTable, func() {
		connection.database.Drop(context.Background())
		connection.Disconnect()
	}
}

func getTestItem() models.Item {
	return models.Item{
		Name:        "Test item",
		Description: "Test item description",
		OwnerID:     "test owner id",
		Images: models.Images{
			PreviewImage: "previewImage",
			Images:       []models.File{},
		},
	}
}

func TestItemsAddAndReadItemSameUser(t *testing.T) {
	itemsTable, teardown := setupItemsTableConnection(t)
	defer teardown()

	item := getTestItem()

	itemID, err := itemsTable.AddItem(item)
	assert.NoError(t, err)

	itemRead, err := itemsTable.GetItem(itemID)
	assert.NoError(t, err)

	assert.Equal(t, item.Name, itemRead.Name)
	assert.Equal(t, item.Description, itemRead.Description)
	assert.Equal(t, item.OwnerID, itemRead.OwnerID)
}

func TestItemsDeleteItem(t *testing.T) {
	itemsTable, teardown := setupItemsTableConnection(t)
	defer teardown()

	item := getTestItem()

	itemID, err := itemsTable.AddItem(item)
	assert.NoError(t, err)

	err = itemsTable.DeleteItem(itemID)
	assert.NoError(t, err)

	_, err = itemsTable.GetItem(itemID)
	assert.ErrorIs(t, err, database.ErrNothingFound)
}

func TestItemsAddImages(t *testing.T) {
	itemsTable, teardown := setupItemsTableConnection(t)
	defer teardown()

	item := getTestItem()

	itemID, err := itemsTable.AddItem(item)
	assert.NoError(t, err)

	images := []models.File{
		{FileName: "test1"},
		{FileName: "test2"},
	}

	err = itemsTable.AddImages(itemID, images)

	if err != nil {
		t.Fatalf(`Failed to add images: %v`, err)
	}

	itemRead, err := itemsTable.GetItem(itemID)
	assert.NoError(t, err)

	assert.Equal(t, len(images), len(itemRead.Images.Images))
}

func TestItemsAddImagesItemNotFound(t *testing.T) {
	itemsTable, teardown := setupItemsTableConnection(t)
	defer teardown()

	images := []models.File{
		{FileName: "test1"},
		{FileName: "test2"},
	}

	err := itemsTable.AddImages("649aec7beb656369627835da", images)
	assert.ErrorIs(t, err, database.ErrNothingFound)
}

func TestItemsSetPreviewImage(t *testing.T) {
	itemsTable, teardown := setupItemsTableConnection(t)
	defer teardown()

	item := getTestItem()

	itemID, err := itemsTable.AddItem(item)
	assert.NoError(t, err)

	newPreviewImageName := "new_name"

	err = itemsTable.SetPreviewImage(itemID, newPreviewImageName)
	assert.NoError(t, err)

	itemRead, err := itemsTable.GetItem(itemID)
	assert.NoError(t, err)

	assert.Equal(t, itemRead.Images.PreviewImage, newPreviewImageName)
}

func TestItemsSetPreviewImageFail(t *testing.T) {
	itemsTable, teardown := setupItemsTableConnection(t)
	defer teardown()

	newPreviewImageName := "new_name"

	err := itemsTable.SetPreviewImage("649aec7beb656369627835da", newPreviewImageName)
	assert.ErrorIs(t, err, database.ErrNothingFound)
}

func TestDeleteItemCollection(t *testing.T) {
	itemsTable, teardown := setupItemsTableConnection(t)
	defer teardown()

	item := getTestItem()

	itemID, err := itemsTable.AddItem(item)
	assert.NoError(t, err)

	err = itemsTable.DeleteCollection()
	assert.NoError(t, err)

	_, err = itemsTable.GetItem(itemID)
	assert.ErrorIs(t, err, database.ErrNothingFound)
}

func TestGetItemsValid(t *testing.T) {
	itemsTable, teardown := setupItemsTableConnection(t)
	defer teardown()

	item := getTestItem()

	itemCount := 4

	for i := 0; i < itemCount; i++ {
		_, err := itemsTable.AddItem(item)
		assert.NoError(t, err)
	}

	items, err := itemsTable.GetItems(item.OwnerID)
	assert.NoError(t, err)

	assert.Equal(t, itemCount, len(items))
}

func TestGetItemsNoItem(t *testing.T) {
	itemsTable, teardown := setupItemsTableConnection(t)
	defer teardown()

	item := getTestItem()

	items, err := itemsTable.GetItems(item.OwnerID)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(items))
}

func TestSetInvoiceValid(t *testing.T) {
	itemsTable, teardown := setupItemsTableConnection(t)
	defer teardown()

	item := getTestItem()

	itemID, err := itemsTable.AddItem(item)
	assert.NoError(t, err)

	invoice := models.File{
		ID:       "invoice_id",
		FileName: "invoice_name.jpg",
	}

	err = itemsTable.SetInvoice(itemID, invoice)
	assert.NoError(t, err)

	itemRead, err := itemsTable.GetItem(itemID)
	assert.NoError(t, err)

	assert.True(t, reflect.DeepEqual(invoice, itemRead.PurchaseInfo.Invoice))
}

func TestSetInvoiceNoItemFound(t *testing.T) {
	itemsTable, teardown := setupItemsTableConnection(t)
	defer teardown()

	invoice := models.File{}

	err := itemsTable.SetInvoice("649aec7beb656369627835da", invoice)
	assert.Equal(t, database.ErrNothingFound, err)
}

func TestDeleteItemInvoice(t *testing.T) {
	itemsTable, teardown := setupItemsTableConnection(t)
	defer teardown()

	item := getTestItem()

	itemID, err := itemsTable.AddItem(item)
	assert.NoError(t, err)

	invoice := models.File{
		ID:       "invoice_id",
		FileName: "invoice_name.jpg",
	}

	err = itemsTable.SetInvoice(itemID, invoice)
	assert.NoError(t, err)

	err = itemsTable.DeleteItemInvoice(itemID)
	assert.NoError(t, err)

	itemRead, err := itemsTable.GetItem(itemID)
	assert.NoError(t, err)

	assert.True(t, reflect.DeepEqual(models.File{}, itemRead.PurchaseInfo.Invoice))
}

func TestDeleteItemInvoiceItemNotFound(t *testing.T) {
	itemsTable, teardown := setupItemsTableConnection(t)
	defer teardown()

	err := itemsTable.DeleteItemInvoice("649aec7beb656369627835da")
	assert.Equal(t, database.ErrNothingFound, err)
}

func TestDeleteItemImage(t *testing.T) {
	itemsTable, teardown := setupItemsTableConnection(t)
	defer teardown()

	item := getTestItem()

	itemID, err := itemsTable.AddItem(item)
	assert.NoError(t, err)

	images := []models.File{
		{ID: "1", FileName: "test1"},
		{ID: "2", FileName: "test2"},
	}

	err = itemsTable.AddImages(itemID, images)
	assert.NoError(t, err)

	err = itemsTable.DeleteItemImage(itemID, images[0].ID)
	assert.NoError(t, err)

	itemRead, err := itemsTable.GetItem(itemID)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(itemRead.Images.Images))
	assert.Equal(t, images[1].FileName, itemRead.Images.Images[0].FileName)
}

func TestDeleteItemImageItemNotFound(t *testing.T) {
	itemsTable, teardown := setupItemsTableConnection(t)
	defer teardown()

	err := itemsTable.DeleteItemImage("649aec7beb656369627835da", "1")
	assert.Equal(t, database.ErrNothingFound, err)
}

func TestDeleteItemImageImageIDNotFound(t *testing.T) {
	itemsTable, teardown := setupItemsTableConnection(t)
	defer teardown()

	item := getTestItem()

	itemID, err := itemsTable.AddItem(item)
	assert.NoError(t, err)

	images := []models.File{
		{ID: "1", FileName: "test1"},
		{ID: "2", FileName: "test2"},
	}

	err = itemsTable.AddImages(itemID, images)
	assert.NoError(t, err)

	err = itemsTable.DeleteItemImage(itemID, "3")
	assert.Equal(t, database.ErrNothingFound, err)

	itemRead, err := itemsTable.GetItem(itemID)
	assert.NoError(t, err)

	assert.True(t, reflect.DeepEqual(images, itemRead.Images.Images))
}

func TestEditItem(t *testing.T) {
	itemsTable, teardown := setupItemsTableConnection(t)
	defer teardown()

	item := getTestItem()

	itemID, err := itemsTable.AddItem(item)
	assert.NoError(t, err)

	createdItem, err := itemsTable.GetItem(itemID)
	assert.NoError(t, err)

	createdItem.Name = "New Name"
	createdItem.Description = "New Description"

	err = itemsTable.EditItem(createdItem)
	assert.NoError(t, err)

	itemRead, err := itemsTable.GetItem(itemID)
	assert.NoError(t, err)

	assert.True(t, reflect.DeepEqual(itemRead, createdItem))
}

func TestEditItemWrongItemID(t *testing.T) {
	itemsTable, teardown := setupItemsTableConnection(t)
	defer teardown()

	item := getTestItem()
	item.ID = "222"

	err := itemsTable.EditItem(item)
	assert.Equal(t, database.ErrGeneratingObjectIDFailed, err)
}

func TestEditItemNotFound(t *testing.T) {
	itemsTable, teardown := setupItemsTableConnection(t)
	defer teardown()

	item := getTestItem()
	item.ID = "649aec7beb656369627835da"

	err := itemsTable.EditItem(item)
	assert.Equal(t, database.ErrNothingFound, err)
}
