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

	// Right after the MongoDB client is initialized, the defer statement is set up. This line doesn't disconnect from MongoDB immediately. Instead, it schedules the Disconnect function to be called later.
	// When the main function is about to complete execution (which would typically happen if the server shuts down or there's a fatal error causing the program to exit), the defer statement triggers
	// This calls the Disconnect() function ensuring that MONGODB Connection is cleanly closed
	defer mongoClient.MongoClient.Disconnect(context.TODO())

	router.POST("/upload", func(c *gin.Context) {
		upload.FileUpload(c, mongoClient)
	})
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
