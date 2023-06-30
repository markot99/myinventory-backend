package usercontroller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginUser_Valid(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	user := testHelper.GetExampleUser()
	testHelper.RegisterUser(user)
	loginJson := `{
		"email": "` + user.Email + `",
		"password": "` + user.Password + `"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/login", strings.NewReader(loginJson))
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var token LoginUserResponseBody
	err := json.Unmarshal(w.Body.Bytes(), &token)
	assert.NoError(t, err)
	assert.NotEmpty(t, token.Token)
}

func TestLoginUser_InvalidBody(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	user := testHelper.GetExampleUser()
	testHelper.RegisterUser(user)

	loginJson := `{
		"password": "test"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/login", strings.NewReader(loginJson))
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestLoginUser_UserDoesNotExist(t *testing.T) {
	r, _, teardown := setupTest(t)
	defer teardown()

	loginJson := `{
		"email": "test@mail.de",
		"password": "test"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/login", strings.NewReader(loginJson))
	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}

func TestLoginUser_InvalidPassword(t *testing.T) {
	r, testHelper, teardown := setupTest(t)
	defer teardown()

	user := testHelper.GetExampleUser()
	testHelper.RegisterUser(user)

	loginJson := `{
		"email": "test@mail.de",
		"password": "4324242"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/login", strings.NewReader(loginJson))
	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}
