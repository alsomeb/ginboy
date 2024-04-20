package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

// https://gin-gonic.com/docs/examples/upload-file/single-file/

const MaxMultipartMem = 8 * 1024 * 1024 // 8 MiB
// const MaxMultipartMem = 1 * 1024 // testing a low size to give err
const FileDir = "photos"

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = MaxMultipartMem
	router.POST("/upload", fileUpload)
	_ = router.Run(":8080")
}

func fileUpload(c *gin.Context) {
	file, err := processFileUpload(c)

	// dynamic err msg depending on file processing err
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// If successful, send back a positive response
	c.JSON(http.StatusOK, gin.H{
		"message":  "File uploaded successfully",
		"filename": file.Filename,
	})
}

// processes file and either returns it or a specific error with its error message that we can send in the c.JSON()
func processFileUpload(c *gin.Context) (*multipart.FileHeader, error) {
	// FormFile returns the first file for the provided form key.
	file, err := c.FormFile("file")
	if err != nil {
		// %w is a format specifier used within the fmt.Errorf function to indicate that you want to format an error
		return nil, fmt.Errorf("no file provided")
	}

	if !isAllowedExt(file.Filename) {
		return nil, fmt.Errorf("only Jpg, Jpeg and png allowed")
	}

	// Check if the file size exceeds the maximum allowed size
	if file.Size > MaxMultipartMem {
		return nil, fmt.Errorf("the file size exceeds the maximum limit of 8 MB")
	}

	// Upload the file to specific dst, if same filename exists it will replace it
	err = c.SaveUploadedFile(file, FileDir+"/"+file.Filename)
	if err != nil {
		return nil, fmt.Errorf("processing file err: %w", err)
	}

	return file, nil
}

func isAllowedExt(filename string) bool {
	fileExt := filepath.Ext(strings.ToLower(filename))
	log.Println("file ext:", fileExt)
	switch fileExt {
	case ".jpg", ".jpeg", ".png":
		return true
	default:
		return false
	}
}
