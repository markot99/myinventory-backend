package itemcontroller

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditItem_Valid(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	userID, token := testHelper.GetUserIDAndToken()
	item := testHelper.AddExampleItem(userID)
	testHelper.UploadImages(item.ID, 3)
	testHelper.UploadInvoice(item.ID)
	newItemName := "newItemName"
	newItemDescription := "newItemDescription"
	newItemQuantity := 7
	editItemJson := `{
		"name": "` + newItemName + `",
		"description": "` + newItemDescription + `",
		"purchaseInfo": {
		  "date": "2022-02-02",
		  "place": "itemBuyPlace",
		  "quantity": ` + strconv.Itoa(newItemQuantity) + `,
		  "unitPrice": 4.44
		}
	  }`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/items/"+item.ID, strings.NewReader(editItemJson))
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)

	modifiedItem := testHelper.GetItem(item.ID)
	assert.Equal(t, newItemName, modifiedItem.Name)
	assert.Equal(t, newItemDescription, modifiedItem.Description)
	assert.Equal(t, newItemQuantity, modifiedItem.PurchaseInfo.Quantity)
}

func TestEditItem_WrongBody(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	userID, token := testHelper.GetUserIDAndToken()
	item := testHelper.AddExampleItem(userID)
	editItemJson := `{
		"wrongBody": "itemName",
		"description": "itemDescription",
		"purchaseInfo": {
		  "date": "2022-02-02",
		  "place": "itemBuyPlace",
		  "quantity": 4,
		  "unitPrice": 4.44
		}
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/items/"+item.ID, strings.NewReader(editItemJson))
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestEditItem_Unauthorized(t *testing.T) {
	r, _, teardown := setupTest(t)
	defer teardown()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/items/323", strings.NewReader(``))
	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}

func TestEditItem_DateInFuture(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	userID, token := testHelper.GetUserIDAndToken()
	item := testHelper.AddExampleItem(userID)
	testHelper.UploadImages(item.ID, 3)
	testHelper.UploadInvoice(item.ID)
	newItemName := "newItemName"
	newItemDescription := "newItemDescription"
	newItemQuantity := 7
	editItemJson := `{
		"name": "` + newItemName + `",
		"description": "` + newItemDescription + `",
		"purchaseInfo": {
		  "date": "2100-02-02",
		  "place": "itemBuyPlace",
		  "quantity": ` + strconv.Itoa(newItemQuantity) + `,
		  "unitPrice": 4.44
		}
	  }`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/items/"+item.ID, strings.NewReader(editItemJson))
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}
