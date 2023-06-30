package fileutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPDFValid(t *testing.T) {
	assert.True(t, IsPdf("test.pdf"))
	assert.True(t, IsPdf("file_321.pdf"))
}

func TestIsPDFInvalid(t *testing.T) {
	assert.False(t, IsPdf("test.txt"))
	assert.False(t, IsPdf("test.pdf.txt"))
	assert.False(t, IsPdf("file_321.pdf.txt"))
	assert.False(t, IsPdf(".pdf"))
}

func TestIsImageValid(t *testing.T) {
	assert.True(t, IsImage("test.jpg"))
	assert.True(t, IsImage("test.jpeg"))
	assert.True(t, IsImage("test.png"))
	assert.True(t, IsImage("file_321.jpg"))
	assert.True(t, IsImage("file_321.jpeg"))
	assert.True(t, IsImage("file_321.png"))
}

func TestIsImageInvalid(t *testing.T) {
	assert.False(t, IsImage("test.txt"))
	assert.False(t, IsImage("test.jpg.txt"))
	assert.False(t, IsImage("file_321.png.txt"))
	assert.False(t, IsImage(".jpg"))
	assert.False(t, IsImage(".jpeg"))
	assert.False(t, IsImage(".png"))
}

func TestGetFileType(t *testing.T) {
	assert.Equal(t, "pdf", GetFileType("test.pdf"))
	assert.Equal(t, "pdf", GetFileType("file_321.pdf"))
	assert.Equal(t, "jpg", GetFileType("test.jpg"))
	assert.Equal(t, "jpg", GetFileType("file_321.jpg"))
	assert.Equal(t, "jpeg", GetFileType("test.jpeg"))
	assert.Equal(t, "jpeg", GetFileType("file_321.jpeg"))
	assert.Equal(t, "png", GetFileType("test.png"))
	assert.Equal(t, "png", GetFileType("file_321.png"))
	assert.Equal(t, "", GetFileType("test"))
	assert.Equal(t, "", GetFileType("file_321"))
	assert.Equal(t, "", GetFileType(""))
}

func TestGetFileName(t *testing.T) {
	assert.Equal(t, "test", GetFileName("test.pdf"))
	assert.Equal(t, "file_321", GetFileName("file_321.pdf"))
	assert.Equal(t, "test.file.312", GetFileName("test.file.312.jpg"))
	assert.Equal(t, "", GetFileName(".jpg"))
}
