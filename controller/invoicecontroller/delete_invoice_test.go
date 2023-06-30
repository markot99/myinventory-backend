package invoicecontroller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteInvoice_Valid(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	userID, token := testHelper.GetUserIDAndToken()
	item := testHelper.AddExampleItem(userID)
	testHelper.UploadInvoice(item.ID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/items/"+item.ID+"/invoice", nil)
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)

	modifiedItem := testHelper.GetItem(item.ID)
	assert.Empty(t, modifiedItem.PurchaseInfo.Invoice.ID)
}

func TestDeleteInvoice_NoInvoice(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	userID, token := testHelper.GetUserIDAndToken()
	item := testHelper.AddExampleItem(userID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/items/"+item.ID+"/invoice", nil)
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)

	modifiedItem := testHelper.GetItem(item.ID)
	assert.Empty(t, modifiedItem.PurchaseInfo.Invoice.ID)
}

func TestDeleteInvoice_ItemDoesNotExist(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	_, token := testHelper.GetUserIDAndToken()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/items/45543/invoice", nil)
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestDeleteInvoice_Unauthorized(t *testing.T) {
	r, _, teardown := setupTest(t)
	defer teardown()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/items/45543/invoice", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}
