package main

import (
	"ginboy/data"
	"ginboy/utils"
	"github.com/gin-gonic/gin"
	"log"
)

// https://gin-gonic.com/docs/examples/upload-file/single-file/
// https://gin-gonic.com/docs/examples/upload-file/multiple-file/

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = utils.MaxMultipartMem
	router.Static("/photos", utils.FileDir) // This code makes any file in the photos directory accessible via a URL like http://localhost:8080/photos/filename.jpg/png
	router.POST("/upload", data.FileUpload)
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error starting server:", err)
		return
	}
}
