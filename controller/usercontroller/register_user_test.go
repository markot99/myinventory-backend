package usercontroller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUser_Valid(t *testing.T) {
	r, _, teardown := setupTest(t)
	defer teardown()

	registerJsonMissingAttributes := `{
		"firstName": "test",
		"lastName": "test",
		"password": "test"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/register", strings.NewReader(registerJsonMissingAttributes))
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestRegisterUser_InvalidEmail(t *testing.T) {
	r, _, teardown := setupTest(t)
	defer teardown()

	registerJsonInvalidEmail := `{
		"email": "test",
		"firstName": "test",
		"lastName": "test",
		"password": "test"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/register", strings.NewReader(registerJsonInvalidEmail))
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestRegisterUser_WrongMethod(t *testing.T) {
	r, _, teardown := setupTest(t)
	defer teardown()

	registerJson := `{
		"email": "test@mail.de",
		"firstName": "test",
		"lastName": "test",
		"password": "test"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/register", strings.NewReader(registerJson))
	r.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
}

func TestRegisterUser_UserAlreadyExists(t *testing.T) {
	r, _, teardown := setupTest(t)
	defer teardown()

	registerJson := `{
		"email": "test@mail.de",
		"firstName": "test",
		"lastName": "test",
		"password": "test"
	}`

	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest(http.MethodPost, "/api/v1/users/register", strings.NewReader(registerJson))
	r.ServeHTTP(w1, req1)
	assert.Equal(t, 204, w1.Code)

	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest(http.MethodPost, "/api/v1/users/register", strings.NewReader(registerJson))
	r.ServeHTTP(w2, req2)
	assert.Equal(t, 409, w2.Code)
}
