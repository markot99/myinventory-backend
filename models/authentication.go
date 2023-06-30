// The models package implements all the necessary structures for handling data between the database and the controllers
package models

// Structure for storing additional information in the JWT token
type AuthenticationClaims struct {
	ID string `json:"id"` // user identification number
}
