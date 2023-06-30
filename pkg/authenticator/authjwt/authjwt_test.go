package authjwt

import (
	"testing"

	"github.com/markot99/myinventory-backend/models"
	"github.com/markot99/myinventory-backend/pkg/authenticator"
	"github.com/stretchr/testify/assert"
)

var testSecret = []byte("myJWTSecret")

func TestGenerateAndValidateToken_Valid(t *testing.T) {
	jwtauthenticator := CreateJWTAuthenticator(testSecret)

	authClaims := models.AuthenticationClaims{
		ID: "testID",
	}

	token, err := jwtauthenticator.GenerateToken(&authClaims)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	authClaimsRead, err := jwtauthenticator.ValidateToken(token)
	assert.NoError(t, err)

	assert.Equal(t, authClaims.ID, authClaimsRead.ID)
}

func TestValidateToken_InvalidToken(t *testing.T) {
	jwtauthenticator := CreateJWTAuthenticator(testSecret)

	_, err := jwtauthenticator.ValidateToken("token")
	assert.ErrorIs(t, err, authenticator.ErrTokenInvalid)
}

func TestGenerateAndValidateToken_InvalidSecret(t *testing.T) {
	jwtauthenticator := CreateJWTAuthenticator(testSecret)

	authClaims := models.AuthenticationClaims{
		ID: "testID",
	}

	token, err := jwtauthenticator.GenerateToken(&authClaims)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	jwtauthenticator.secret = []byte("invalidSecret")

	_, err = jwtauthenticator.ValidateToken(token)
	assert.ErrorIs(t, err, authenticator.ErrTokenInvalid)
}

func TestValidateToken_ExpiredToken(t *testing.T) {
	jwtauthenticator := CreateJWTAuthenticator(testSecret)

	expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBdXRoZW50aWNhdGlvbkNsYWltcyI6eyJpZCI6InRlc3RJRCJ9LCJleHAiOjE2NTYyNDExNTJ9.1G81LE-JFGXxnoyQeGN1AfxcCku-1AMTTfjIDo8NH3geyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBdXRoZW50aWNhdGlvbkNsYWltcyI6eyJpZCI6InRlc3RJRCJ9LCJleHAiOjE2NTYyNDExNTJ9.1G81LE-JFGXxnoyQeGN1AfxcCku-1AMTTfjIDo8NH3g"

	_, err := jwtauthenticator.ValidateToken(expiredToken)
	assert.Error(t, err)
}
