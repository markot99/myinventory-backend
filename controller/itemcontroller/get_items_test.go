package itemcontroller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetItems_Valid(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	userID, token := testHelper.GetUserIDAndToken()
	testHelper.AddExampleItem(userID)
	testHelper.AddExampleItem(userID)
	testHelper.AddExampleItem(userID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/items", nil)
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var itemsResponse GetItemsResponseBody
	err := json.Unmarshal(w.Body.Bytes(), &itemsResponse)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(itemsResponse.ResponseItems))
}

func TestGetItems_ValidNoItems(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	_, token := testHelper.GetUserIDAndToken()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/items", nil)
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var itemsResponse GetItemsResponseBody
	err := json.Unmarshal(w.Body.Bytes(), &itemsResponse)
	assert.NoError(t, err)
	assert.Equal(t, "{\"items\":[]}", w.Body.String())
}

func TestGetItems_Unauthorized(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	userID, _ := testHelper.GetUserIDAndToken()
	testHelper.AddExampleItem(userID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/items", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}
