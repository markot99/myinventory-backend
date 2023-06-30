package localstorage

import (
	"os"

	"github.com/google/uuid"
	"github.com/markot99/myinventory-backend/pkg/storage"
)

// LocalStorage is the struct that contains the base path for the local storage
type LocalStorage struct {
	basePath string // base path for the local storage
}

// NewLocalStorage is the constructor for creating a LocalStorage object
func NewLocalStorage(basePath string) *LocalStorage {
	localStorage := LocalStorage{basePath: basePath}
	localStorage.createDirectory(basePath)
	return &localStorage
}

// FileExists is used to check if a file exists in the local storage
func (localStorage *LocalStorage) FileExists(fileName string) bool {
	_, err := os.Stat(localStorage.basePath + "/" + fileName)
	return err == nil
}

// generateFileIdentifier is used to generate a unique file identifier that is used to identify the file in the local storage
func (localStorage *LocalStorage) generateFileIdentifier() (string, error) {
	for i := 0; i < 10; i++ {
		uuidString := uuid.New().String()

		if !localStorage.FileExists(localStorage.basePath + "/" + uuidString) {
			return uuidString, nil
		}
	}
	return "", storage.ErrFailedToGenerateFileID
}

// createDirectory is used to create a directory in the local storage.
func (localStorage *LocalStorage) createDirectory(dirPath string) error {
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return storage.ErrFailedToCreateDirectory
	}
	return nil
}

// SaveFile is used to save a file in the local storage.
func (localStorage *LocalStorage) SaveFile(file []byte) (string, error) {
	id, err := localStorage.generateFileIdentifier()
	if err != nil {
		return "", storage.ErrFailedToSaveFile
	}

	err = os.WriteFile(localStorage.basePath+"/"+id, file, os.ModePerm)
	if err != nil {
		return "", storage.ErrFailedToSaveFile
	}
	return id, nil
}

// GetFile is used to retrieve a file from the local storage by its id.
func (localStorage *LocalStorage) GetFile(fileID string) ([]byte, error) {
	file, err := os.ReadFile(localStorage.basePath + "/" + fileID)
	if err != nil {
		return nil, storage.ErrFailedToGetFile
	}
	return file, nil
}

// DeleteFile is used to delete a file from the local storage by its id.
func (localStorage *LocalStorage) DeleteFile(fileID string) error {
	// if file does not exist no need to delete it
	if !localStorage.FileExists(localStorage.basePath + "/" + fileID) {
		return nil
	}

	err := os.Remove(localStorage.basePath + "/" + fileID)
	if err != nil {
		return storage.ErrFailedToDeleteFile
	}
	return nil
}

// ClearStorage is used to delete all files from the local storage.
func (localStorage *LocalStorage) ClearStorage() error {
	err := os.RemoveAll(localStorage.basePath)
	if err != nil {
		return storage.ErrFailedToClearStorage
	}
	return nil
}
