package imagecontroller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetPreviewImage_Valid(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	userID, token := testHelper.GetUserIDAndToken()
	item := testHelper.AddExampleItem(userID)
	testHelper.UploadImages(item.ID, 2)
	item = testHelper.GetItem(item.ID)
	setPreviewImageJson := `{
		"name": "` + item.Images.Images[0].ID + `"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/items/"+item.ID+"/images/preview", strings.NewReader(setPreviewImageJson))
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)
	assert.Equal(t, "", w.Body.String())

	modifiedItem := testHelper.GetItem(item.ID)

	assert.Equal(t, item.Images.Images[0].ID, modifiedItem.Images.PreviewImage)
}

func TestSetPreviewImage_ItemDoesNotExist(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	_, token := testHelper.GetUserIDAndToken()
	setPreviewImageJson := `{
		"name": "previewImage.jpg"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/items/4342334/images/preview", strings.NewReader(setPreviewImageJson))
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestSetPrevieImage_Unauthorized(t *testing.T) {
	r, _, teardown := setupTest(t)
	defer teardown()

	setPreviewImageJson := `{
		"name": "previewImage.jpg"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/items/43243/images/preview", strings.NewReader(setPreviewImageJson))
	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}
