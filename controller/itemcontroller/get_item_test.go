package itemcontroller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/markot99/myinventory-backend/models"
	"github.com/stretchr/testify/assert"
)

func TestGetItem_Valid(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	userID, token := testHelper.GetUserIDAndToken()
	item := testHelper.AddExampleItem(userID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/items/"+item.ID, nil)
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var itemResponse models.Item
	err := json.Unmarshal(w.Body.Bytes(), &itemResponse)
	assert.NoError(t, err)
	assert.NotEmpty(t, itemResponse.ID)
}

func TestGetItem_Unauthorized(t *testing.T) {
	r, _, teardown := setupTest(t)
	defer teardown()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/items/3232331", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}
