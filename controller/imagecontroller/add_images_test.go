package imagecontroller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddImages_Valid(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	imageCount := 4
	userID, token := testHelper.GetUserIDAndToken()
	item := testHelper.AddExampleItem(userID)
	writer, body := testHelper.GenerateMultipartData(".jpg", "images", imageCount)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/items/"+item.ID+"/images", &body)
	req.Header.Add("Authorization", token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	item = testHelper.GetItem(item.ID)
	assert.Equal(t, imageCount, len(item.Images.Images))
}

func TestAddImages_ItemDoesNotExist(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	_, token := testHelper.GetUserIDAndToken()
	writer, body := testHelper.GenerateMultipartData(".jpg", "images", 4)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/items/443/images", &body)
	req.Header.Add("Authorization", token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestAddImages_Unauthorized(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	writer, body := testHelper.GenerateMultipartData(".jpg", "images", 4)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/items/3434/images", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}
