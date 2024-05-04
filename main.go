package main

import (
	"context"
	"ginboy/database"
	"ginboy/upload"
	"ginboy/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = utils.MaxMultipartMem
	router.Static("/photos", utils.FileDir) // This code makes any file in the photos directory accessible via a URL like http://localhost:8080/photos/filename.jpg/png

	// MongoDB
	mongoClient := database.InitClient()
	defer mongoClient.MongoClient.Disconnect(context.TODO()) // Ensure disconnection of MongoDB when main() exits

	router.POST("/upload", func(c *gin.Context) {
		upload.FileUpload(c, mongoClient)
	})
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
