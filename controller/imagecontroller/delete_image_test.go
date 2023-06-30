package imagecontroller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteImage_Valid(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	imageCount := 4
	userID, token := testHelper.GetUserIDAndToken()
	item := testHelper.AddExampleItem(userID)
	testHelper.UploadImages(item.ID, imageCount)
	item = testHelper.GetItem(item.ID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/items/"+item.ID+"/images/"+item.Images.Images[0].ID, nil)
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)

	modifiedItem := testHelper.GetItem(item.ID)
	assert.Equal(t, imageCount-1, len(modifiedItem.Images.Images))
}

func TestDeleteImage_ItemDoesNotExist(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	_, token := testHelper.GetUserIDAndToken()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/items/43242/images/43242", nil)
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestDeleteImage_ImageDoesNotExist(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	imageCount := 4
	userID, token := testHelper.GetUserIDAndToken()
	item := testHelper.AddExampleItem(userID)
	testHelper.UploadImages(item.ID, imageCount)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/items/"+item.ID+"/images/443234", nil)
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestDeleteImage_Unauthorized(t *testing.T) {
	r, _, teardown := setupTest(t)
	defer teardown()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/items/43242/images/43242", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}
