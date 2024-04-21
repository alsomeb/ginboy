package data

import (
	"fmt"
	"ginboy/utils"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

func FileUpload(c *gin.Context) {
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

	if !utils.IsAllowedExt(file.Filename) {
		return nil, fmt.Errorf("only Jpg, Jpeg and png allowed")
	}

	// Check if the file size exceeds the maximum allowed size
	if file.Size > utils.MaxMultipartMem {
		return nil, fmt.Errorf("the file size exceeds the maximum limit of 8 MB")
	}

	file.Filename = utils.GenerateFileName(file.Filename)
	err = c.SaveUploadedFile(file, utils.FileDir+"/"+file.Filename)
	if err != nil {
		return nil, fmt.Errorf("processing file err: %w", err)
	}

	return file, nil
}
