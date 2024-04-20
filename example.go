package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

// https://gin-gonic.com/docs/examples/upload-file/single-file/

const MaxMultipartMem = 8 * 1024 * 1024 // 8 MiB
//const MaxMultipartMem = 1 * 1024 // testing a low size to give err
const FileDir = "photos"

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = MaxMultipartMem
	router.POST("/upload", fileUpload)
	_ = router.Run(":8080")
}

func fileUpload(c *gin.Context) {
	// FormFile returns the first file for the provided form key.
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file!"})
		// Without the return, the function would continue executing
		return
	}

	if !isAllowedExt(file.Filename) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only JPG and PNG files are allowed"})
		return
	}

	// Check if the file size exceeds the maximum allowed size
	if file.Size > MaxMultipartMem {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "The file size exceeds the maximum limit of 8 MB",
		})
		return
	}

	// Upload the file to specific dst, if same filename exists it will replace it
	err = c.SaveUploadedFile(file, FileDir+"/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If successful, send back a positive response
	c.JSON(http.StatusOK, gin.H{
		"message":  "File uploaded successfully",
		"filename": file.Filename,
	})
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
