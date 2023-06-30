package storage

import (
	"errors"
)

var ErrFailedToSaveFile = errors.New("failed to save file")
var ErrFailedToGetFile = errors.New("failed to get file")
var ErrFailedToGenerateFileID = errors.New("failed to generate file id")
var ErrFailedToDeleteFile = errors.New("failed to delete file")
var ErrFailedToClearStorage = errors.New("failed to clear storage")
var ErrFailedToCreateDirectory = errors.New("failed to create directory")

// Storage is the interface that wraps the basic methods for storing files
type Storage interface {
	// FileExists is used to check if a file exists. True if the file exists, false otherwise
	FileExists(fileName string) bool
	// SaveFile is used to save a file in the storage.
	SaveFile(file []byte) (string, error)
	// GetFile is used to retrieve a file from the storage by its name
	GetFile(fileName string) ([]byte, error)
	// DeleteFile is used to delete a file from the storage by its name
	DeleteFile(fileName string) error
	// ClearStorage is used to delete all files from the storage
	ClearStorage() error
}
