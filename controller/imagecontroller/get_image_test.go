package imagecontroller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetImage_Valid(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	userID, token := testHelper.GetUserIDAndToken()
	item := testHelper.AddExampleItem(userID)
	testHelper.UploadImages(item.ID, 4)
	item = testHelper.GetItem(item.ID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/items/"+item.ID+"/images/"+item.Images.Images[0].ID, nil)
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestGetImage_ItemDoesNotExist(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	_, token := testHelper.GetUserIDAndToken()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/items/33123/images/323123", nil)
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestGetImage_ImageDoesNotExist(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	userID, token := testHelper.GetUserIDAndToken()
	item := testHelper.AddExampleItem(userID)
	testHelper.UploadImages(item.ID, 2)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/items/"+item.ID+"/images/5435345", nil)
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestGetImage_Unauthorized(t *testing.T) {
	r, _, teardown := setupTest(t)
	defer teardown()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/items/34423/images/5435345", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}
