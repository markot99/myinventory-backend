package passwordutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPasswod(t *testing.T) {
	password := "password"
	hashedPassword, err := HashPassword(&password)
	assert.NoError(t, err)
	assert.NotNil(t, hashedPassword)

	hashedPassword2, err := HashPassword(&password)
	assert.NoError(t, err)
	assert.NotNil(t, hashedPassword2)

	// hashed passwords should not be equal
	assert.NotEqual(t, hashedPassword, hashedPassword2)
}

func TestCheckPassword(t *testing.T) {
	password := "password"
	hashedPassword, err := HashPassword(&password)
	assert.NoError(t, err)
	assert.NotNil(t, hashedPassword)

	assert.True(t, CheckPassword(&hashedPassword, &password))
	password = "password1"
	assert.False(t, CheckPassword(&hashedPassword, &password))
}
