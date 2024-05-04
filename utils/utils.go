package utils

import (
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"log"
	"os"
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

// LoadEnvVariable Loads current env file and returns requested value by key
func LoadEnvVariable(key string) string {
	err := godotenv.Load(".env") // Load will read your env file(s) and load them into ENV for this process.

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}
