// The models package implements all the necessary structures for handling data between the database and the controllers
package models

// Structure for storing user data
type User struct {
	ID        string `json:"_id" bson:"_id,omitempty"`   // user identification number
	FirstName string `json:"firstName" bson:"firstName"` // first name of the user
	LastName  string `json:"lastName" bson:"lastName"`   // last name of the user
	Email     string `json:"email" bson:"email"`         // email of the user
	Password  string `json:"password" bson:"password"`   // hashed password of the user
}
