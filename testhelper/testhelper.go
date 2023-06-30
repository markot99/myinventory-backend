// The package testhelper is used to provide basic functions for testing
package testhelper

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/models"
	"github.com/markot99/myinventory-backend/pkg/authenticator/authjwt"
	"github.com/markot99/myinventory-backend/pkg/database/mongodb"
	"github.com/markot99/myinventory-backend/pkg/storage/localstorage"
	"github.com/markot99/myinventory-backend/utils/passwordutils"
	"github.com/stretchr/testify/require"
)

// TestHelper is the struct that contains the necessary objects for testing
type TestHelper struct {
	test              *testing.T                 // test object
	MongoDBConnection *mongodb.MongoDBConnection // mongodb connection
	ItemTable         *mongodb.ItemTable         // item table
	UserTable         *mongodb.UserTable         // user table
	ImageStorage      *localstorage.LocalStorage // image storage
	InvoiceStorage    *localstorage.LocalStorage // invoice storage
	Authenticator     *authjwt.JWTAuthenticator  // jwt authenticator
	Router            *gin.Engine                // gin router
}

const TestDir = ".test_tmp"
const TestDBName = "myinventory_test"
const TestDBHost = "mongodb://localhost:27017"

// CreateTestHelper is the constructor for creating a TestHelper object
func CreateTestHelper(t *testing.T) TestHelper {
	mongodbConnection, err := mongodb.NewMongoDBConnection(TestDBHost, TestDBName)
	require.NoError(t, err)

	authenticator := authjwt.CreateJWTAuthenticator([]byte("myJWTSecret"))

	itemTable := mongodb.NewItemTable(mongodbConnection, "items")
	userTable := mongodb.NewUserTable(mongodbConnection, "users")

	err = os.MkdirAll(TestDir, os.ModePerm)
	require.NoError(t, err)

	imageStorage := localstorage.NewLocalStorage(TestDir + "/" + "images")
	invoiceStorage := localstorage.NewLocalStorage(TestDir + "/" + "invoices")

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	router := gin.Default()

	return TestHelper{test: t, MongoDBConnection: &mongodbConnection, Authenticator: authenticator, ItemTable: itemTable, UserTable: userTable, ImageStorage: imageStorage, InvoiceStorage: invoiceStorage, Router: router}
}

// Teardown is used to teardown the test environment (must be called at the end of each test)
func (testHelper *TestHelper) Teardown() {
	err := testHelper.MongoDBConnection.DropDatabase()
	require.NoError(testHelper.test, err)

	err = testHelper.MongoDBConnection.Disconnect()
	require.NoError(testHelper.test, err)

	err = os.RemoveAll(TestDir)
	require.NoError(testHelper.test, err)
}

// GetExampleUser is used to get an example user
func (testHelper *TestHelper) GetExampleUser() models.User {
	return models.User{
		FirstName: "Test",
		LastName:  "Test",
		Email:     "test@mail.de",
		Password:  "password",
	}
}

// GetExampleItem is used to get an example item with the given owner id
func (testHelper *TestHelper) GetExampleItem(ownerID string) models.Item {
	purchaseInfo := models.PurchaseInfo{
		Date:      time.Now(),
		Place:     "Laden",
		UnitPrice: 10.23,
		Quantity:  5,
		Invoice:   models.File{},
	}

	images := models.Images{PreviewImage: "", Images: []models.File{}}
	return models.Item{Name: "item", Description: "description", PurchaseInfo: purchaseInfo, OwnerID: ownerID, Images: images}
}

// RegisterUser is used to register an user in the database
func (testHelper *TestHelper) RegisterUser(user models.User) {
	hash, err := passwordutils.HashPassword(&user.Password)
	require.NoError(testHelper.test, err)
	modifiedUser := user
	modifiedUser.Password = hash
	err = testHelper.UserTable.AddUser(modifiedUser)
	require.NoError(testHelper.test, err)
}

// GetUserIDAndToken is used to register an user and get the user id and token
func (testHelper *TestHelper) GetUserIDAndToken() (string, string) {
	user := testHelper.GetExampleUser()
	testHelper.RegisterUser(user)
	registeredUser, err := testHelper.UserTable.GetUserByEmail(user.Email)
	require.NoError(testHelper.test, err)

	authClaims := models.AuthenticationClaims{ID: registeredUser.ID}
	token, err := testHelper.Authenticator.GenerateToken(&authClaims)
	require.NoError(testHelper.test, err)

	return registeredUser.ID, token
}

// AddItem is used to add an item to the database
func (testHelper *TestHelper) AddItem(item models.Item) string {
	itemID, err := testHelper.ItemTable.AddItem(item)
	require.NoError(testHelper.test, err)
	return itemID
}

// AddExampleItem is used to add an example item to the database with the given owner id
func (testHelper *TestHelper) AddExampleItem(userID string) models.Item {
	item := testHelper.GetExampleItem(userID)
	itemID := testHelper.AddItem(item)
	return testHelper.GetItem(itemID)
}

// GetItem is used to get an item from the database by item id
func (testHelper *TestHelper) GetItem(itemID string) models.Item {
	item, err := testHelper.ItemTable.GetItem(itemID)
	require.NoError(testHelper.test, err)
	return item
}

// GenerateMultipartData is used to generate multipart data for testing, fileType can be e.g. "pdf" or "jpg"
func (testHelper *TestHelper) GenerateMultipartData(fileType string, formName string, filesCount int) (*multipart.Writer, bytes.Buffer) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	for i := 0; i < filesCount; i++ {
		imageName := fmt.Sprintf("file_%d."+fileType, i)

		part, err := w.CreateFormFile(formName, imageName)
		require.NoError(testHelper.test, err)

		_, err = part.Write([]byte("test"))
		require.NoError(testHelper.test, err)
	}

	w.Close()
	return w, b
}

// GenerateFileContent is used to generate simple file content for testing
func (testHelper *TestHelper) GenerateFileContent() []byte {
	return []byte("datadatadatadatadatadatadatadatadatadatadatadatadatadatadata")
}

// UploadInvoice is used to upload an invoice to the database and storage for the given item id
func (testHelper *TestHelper) UploadInvoice(itemID string) {
	data := testHelper.GenerateFileContent()
	id, err := testHelper.InvoiceStorage.SaveFile(data)
	require.NoError(testHelper.test, err)
	invoiceOb := models.File{ID: id, FileName: "invoice_1.pdf"}
	err = testHelper.ItemTable.SetInvoice(itemID, invoiceOb)
	require.NoError(testHelper.test, err)
}

// UploadImages is used to upload images to the database and storage for the given item id
func (testHelper *TestHelper) UploadImages(itemID string, itemCount int) {
	var images []models.File
	for i := 0; i < itemCount; i++ {
		data := testHelper.GenerateFileContent()
		id, err := testHelper.ImageStorage.SaveFile(data)
		require.NoError(testHelper.test, err)
		images = append(images, models.File{ID: id, FileName: fmt.Sprintf("image_%d.jpg", i)})
	}

	err := testHelper.ItemTable.AddImages(itemID, images)
	require.NoError(testHelper.test, err)
}
