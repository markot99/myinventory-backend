package invoicecontroller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddInvoice_Valid(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	userID, token := testHelper.GetUserIDAndToken()
	item := testHelper.AddExampleItem(userID)
	writer, body := testHelper.GenerateMultipartData(".pdf", "invoice", 1)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/items/"+item.ID+"/invoice", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)

	modifiedItem := testHelper.GetItem(item.ID)
	assert.NotEmpty(t, modifiedItem.PurchaseInfo.Invoice.ID)
}

func TestAddInvoice_ItemDoesNotExist(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	_, token := testHelper.GetUserIDAndToken()
	writer, body := testHelper.GenerateMultipartData(".pdf", "invoice", 1)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/items/434234/invoice", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestAddInvoice_WrongFileType(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	userID, token := testHelper.GetUserIDAndToken()
	item := testHelper.AddExampleItem(userID)
	writer, body := testHelper.GenerateMultipartData(".jpg", "invoice", 1)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/items/"+item.ID+"/invoice", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestAddInvoice_WrongForm(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	userID, token := testHelper.GetUserIDAndToken()
	item := testHelper.AddExampleItem(userID)
	writer, body := testHelper.GenerateMultipartData(".pdf", "invoices", 1)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/items/"+item.ID+"/invoice", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestAddInvoice_Unauthorized(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	writer, body := testHelper.GenerateMultipartData(".pdf", "invoice", 1)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/items/334/invoice", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}
