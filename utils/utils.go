package utils

import (
	"github.com/google/uuid"
	"path/filepath"
	"strings"
)

const MaxMultipartMem = 8 * 1024 * 1024 // 8 MiB
const FileDir = "photos"

func IsAllowedExt(filename string) bool {
	fileExt := filepath.Ext(strings.ToLower(filename))
	switch fileExt {
	case ".jpg", ".jpeg", ".png":
		return true
	default:
		return false
	}
}

func GenerateFileName(filename string) string {
	fileExt := filepath.Ext(strings.ToLower(filename))
	return uuid.NewString() + fileExt
}
