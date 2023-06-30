// the package fileutils provides some functions to analyze and validate files
package fileutils

import (
	"strings"
)

// IsPdf checks if the file is a pdf file by checking the file extension
func IsPdf(file string) bool {
	fileType := GetFileType(file)
	fileName := GetFileName(file)
	if fileType == "pdf" && fileName != "" {
		return true
	} else {
		return false
	}
}

// IsImage checks if the file is an image by checking the file extension, only jpg, jpeg and png files are allowed
func IsImage(file string) bool {
	fileType := GetFileType(file)
	fileName := GetFileName(file)
	if (fileType == "jpg" || fileType == "jpeg" || fileType == "png") && (fileName != "") {
		return true
	} else {
		return false
	}
}

// GetFileType is used to get the file type by extracting the file extension
func GetFileType(file string) string {
	parts := strings.Split(file, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	} else {
		return ""
	}
}

// GetFileName is used to get the file name by removing the file extension
func GetFileName(file string) string {
	parts := strings.Split(file, ".")
	if len(parts) > 1 {
		return strings.Join(parts[:len(parts)-1], ".")
	} else {
		return file
	}
}
