package localstorage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TestDir = ".test_tmp"

func setupStorage(t *testing.T) (LocalStorage, func()) {
	storage := NewLocalStorage(TestDir + "/" + "test")

	return *storage, func() {
		err := os.RemoveAll(TestDir)
		assert.NoError(t, err)
	}
}

func TestFileExists_Valid(t *testing.T) {
	storage, teardown := setupStorage(t)
	defer teardown()

	id, err := storage.SaveFile([]byte("data"))
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
	assert.True(t, storage.FileExists(id))
}

func TestFileExists_DoesNotExist(t *testing.T) {
	storage, teardown := setupStorage(t)
	defer teardown()

	assert.False(t, storage.FileExists("123"))
}

func TestClearStorage_Valid(t *testing.T) {
	storage, teardown := setupStorage(t)
	defer teardown()

	id1, err := storage.SaveFile([]byte("data"))
	assert.NoError(t, err)
	id2, err := storage.SaveFile([]byte("data"))
	assert.NoError(t, err)

	storage.ClearStorage()

	assert.False(t, storage.FileExists(id1))
	assert.False(t, storage.FileExists(id2))
}
func TestDeleteFile_Valid(t *testing.T) {
	storage, teardown := setupStorage(t)
	defer teardown()
	id, err := storage.SaveFile([]byte("data"))
	assert.NoError(t, err)

	err = storage.DeleteFile(id)
	assert.NoError(t, err)
}

func TestDeleteFile_FileDoesNotExist(t *testing.T) {
	storage, teardown := setupStorage(t)
	defer teardown()

	err := storage.DeleteFile("123")
	assert.NoError(t, err)
}
