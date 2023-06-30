// The authjwt package is used to integrate the "github.com/golang-jwt/jwt" library
package authjwt

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/markot99/myinventory-backend/models"
	"github.com/markot99/myinventory-backend/pkg/authenticator"
	"github.com/markot99/myinventory-backend/pkg/logger"
)

// Claims is the struct that contains the AuthenticationClaims and the StandardClaims for the jwt token
type Claims struct {
	AuthenticationClaims models.AuthenticationClaims // additional data stored in the token
	jwt.StandardClaims                               // standard claims for the jwt token
}

// JWTAuthenticator is the struct that contains the secret for signing and verifying tokens
type JWTAuthenticator struct {
	secret         []byte        // secret for signing and verifying tokens
	expirationTime time.Duration // duration of a token until it expires
}

// CreateJWTAuthenticator is the constructor for creating a JWTAuthenticator object.
func CreateJWTAuthenticator(secret []byte) *JWTAuthenticator {
	// use expiration time of 24 hours
	return &JWTAuthenticator{secret: secret, expirationTime: 24 * time.Hour}
}

// GenerateToken is used to create a token with the default expiration time from the given AuthenticationClaims.
// If successful, the JWT token is returned as a string.
func (authenticator_p *JWTAuthenticator) GenerateToken(authClaims *models.AuthenticationClaims) (string, error) {
	expirationTime := time.Now().Add(authenticator_p.expirationTime)

	claims := &Claims{
		AuthenticationClaims: *authClaims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(authenticator_p.secret)
	if err != nil {
		logger.Error("Failed to generate jwt token: ", err.Error())
		return "", authenticator.ErrGeneratingToken
	}

	return tokenString, nil
}

// ValidateToken is used to validate a token. If successful, the AuthenticationClaims are returned.
func (authenticator_p *JWTAuthenticator) ValidateToken(token string) (models.AuthenticationClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return authenticator_p.secret, nil
	}

	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, keyFunc)
	if err != nil {
		logger.Debug("Failed to parse jwt token: ", err.Error())
		return models.AuthenticationClaims{}, authenticator.ErrTokenInvalid
	}

	claims, ok := parsedToken.Claims.(*Claims)
	if !ok || !parsedToken.Valid {
		logger.Debug("Failed to parse jwt token claims")
		return models.AuthenticationClaims{}, authenticator.ErrTokenInvalid
	}

	return claims.AuthenticationClaims, nil
}
