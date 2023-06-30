// The models package implements all the necessary structures for handling data between the database and the controllers
package models

import (
	"errors"
	"time"
)

// GetImageFileById returns the image file with the given id
// if the image is not found, an error is returned
func GetImageFileById(item Item, imageID string) (File, error) {
	for _, image := range item.Images.Images {
		if image.ID == imageID {
			return image, nil
		}
	}
	return File{}, errors.New("image not found")
}

// GetBuyDateFromString converts a string to a time.Time object
// The string must be in the format "YYYY-MM-DD"
// if the string is not in the correct format, an error is returned
func GetBuyDateFromString(buyDate string) (time.Time, error) {
	layout := "2006-01-02" // set date schema
	t, err := time.Parse(layout, buyDate)
	return t, err
}

// Structure for storing information of a single file
type File struct {
	ID       string `json:"id" bson:"id"`             // identification number of the file
	FileName string `json:"fileName" bson:"fileName"` // name of the file
}

// Structure for storing information of the item images
type Images struct {
	PreviewImage string `json:"previewImage" bson:"previewImage"` // preview image name to be displayed first
	Images       []File `json:"images" bson:"images"`             // all images
}

// Structure for storing purchase information
type PurchaseInfo struct {
	UnitPrice float64   `json:"unitPrice" bson:"unitPrice"` // price of a single item
	Date      time.Time `json:"date" bson:"date"`           // date of purchase
	Place     string    `json:"place" bson:"place"`         // place of purchase
	Quantity  int       `json:"quantity" bson:"quantity"`   // number of items purchased
	Invoice   File      `json:"invoice" bson:"invoice"`     // invoice of the purchase
}

// Structure for storing information about the item
type Item struct {
	ID           string       `json:"id" bson:"_id,omitempty"`          // identification number of the item
	Name         string       `json:"name" bson:"name"`                 // name of the item
	Description  string       `json:"description" bson:"description"`   // description of the item
	PurchaseInfo PurchaseInfo `json:"purchaseInfo" bson:"purchaseInfo"` // purchase information of the item
	OwnerID      string       `json:"ownerID" bson:"ownerID"`           // identification number of the owner of the item
	Images       Images       `json:"images" bson:"images"`             // images of the item
}
