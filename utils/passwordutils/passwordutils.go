// The passwordutils package is used to provide utilities for hashing and checking passwords.
package passwordutils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword is used to hash a password with a salt
func HashPassword(password *string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// CheckPassword is used to check whether an unhashed password matches a hashed password
func CheckPassword(hashedPassword *string, password *string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*hashedPassword), []byte(*password))
	return err == nil
}
