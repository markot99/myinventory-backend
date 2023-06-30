package itemcontroller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddItem_Valid(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	_, token := testHelper.GetUserIDAndToken()
	addItemJson := `{
		"name": "itemName",
		"description": "itemDescription",
		"purchaseInfo": {
		  "date": "2022-02-02",
		  "place": "itemBuyPlace",
		  "quantity": 3,
		  "unitPrice": 4.44
		}
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/items", strings.NewReader(addItemJson))
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)

	var itemResponse AddItemResponseBody
	err := json.Unmarshal(w.Body.Bytes(), &itemResponse)
	assert.NoError(t, err)
	assert.NotEmpty(t, itemResponse.ID)
}

func TestAddItem_WrongDate(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	_, token := testHelper.GetUserIDAndToken()

	addItemJsonWrongDate := `{
		"name": "itemName",
		"description": "itemDescription",
		"purchaseInfo": {
		  "date": "02.04.2022",
		  "place": "itemBuyPlace",
		  "quantity": 3,
		  "unitPrice": 4.44
		}
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/items", strings.NewReader(addItemJsonWrongDate))
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestAddItem_WrongBody(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	_, token := testHelper.GetUserIDAndToken()

	addItemJsonWrongDate := `{
		"namesdada": "itemName",
		"purchaseInfo": {
		  "date": "2022-02-02",
		  "place": "itemBuyPlace",
		  "quantity": 3,
		  "unitPrice": 4.44
		}
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/items", strings.NewReader(addItemJsonWrongDate))
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestAddItem_Unauthorized(t *testing.T) {
	r, _, teardown := setupTest(t)
	defer teardown()

	addItemJson := `{
		"name": "itemName",
		"description": "itemDescription",
		"purchaseInfo": {
		  "date": "2022-02-02",
		  "place": "itemBuyPlace",
		  "quantity": 3,
		  "unitPrice": 4.44
		}
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/items", strings.NewReader(addItemJson))
	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}

func TestAddItem_DateInFuture(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	_, token := testHelper.GetUserIDAndToken()
	addItemJson := `{
		"name": "itemName",
		"description": "itemDescription",
		"purchaseInfo": {
		  "date": "2100-02-02",
		  "place": "itemBuyPlace",
		  "quantity": 3,
		  "unitPrice": 4.44
		}
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/items", strings.NewReader(addItemJson))
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}
