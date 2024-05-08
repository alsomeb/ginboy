package upload

import (
	"fmt"
	"ginboy/database"
	"ginboy/utils"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

// https://gin-gonic.com/docs/examples/upload-file/single-file/
// https://gin-gonic.com/docs/examples/upload-file/multiple-file/

func FileUpload(c *gin.Context, mongoClient *database.MongoClient) {
	// Parse multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Check for file input in the form, 'files' since it will accept an input with potential multiple files
	files, found := form.File["files"]
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "no files found"})
		return
	}

	// Process each file
	for _, file := range files {
		// Here, the 'processErr' variable is declared and checked inline within the if condition.
		// This approach scopes the 'processErr' variable to the if block only, meaning it is not accessible outside of it
		if processErr := processAndSaveFile(c, file); processErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": processErr.Error()})
			return
		}
	}

	// Respond success if all files are processed
	c.JSON(http.StatusOK, gin.H{
		"message": "File(s) uploaded successfully",
		"amount":  len(files),
	})
}

// processes file & saving it to a folder or returns dynamic error msg that we can send in the c.JSON()
func processAndSaveFile(c *gin.Context, file *multipart.FileHeader) error {
	if !utils.IsAllowedExt(file.Filename) {
		return fmt.Errorf("only Jpg, Jpeg and png allowed")
	}

	// Check if the file size exceeds the maximum allowed size
	if file.Size > utils.MaxMultipartMem {
		return fmt.Errorf("the file size exceeds the maximum limit of 8 MB")
	}

	file.Filename = utils.GenerateFileName(file.Filename)
	if err := c.SaveUploadedFile(file, utils.FileDir+"/"+file.Filename); err != nil {
		return fmt.Errorf("processing file err: %w", err)
	}

	return nil
}
