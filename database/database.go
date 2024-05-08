package database

import (
	"context"
	"fmt"
	"ginboy/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// https://www.mongodb.com/developer/products/mongodb/build-go-web-application-gin-mongodb-help-ai/

type MongoClient struct {
	MongoClient *mongo.Client
}

func InitClient() *MongoClient {
	username := utils.LoadEnvVariable("DB_USER")
	password := utils.LoadEnvVariable("DB_PASSWORD")
	mongoURI := fmt.Sprintf("mongodb+srv://%v:%v@alsomeb.jcl49rx.mongodb.net/Files?retryWrites=true&w=majority", username, password)
	log.Println("---- Successfully Loaded .env file with DB Details ----")

	client := connectMongo(mongoURI)

	return client
}

func connectMongo(mongoURI string) *MongoClient {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	// reusing the err variable, which was previously used in the mongo.Connect call, and assigning it a new value based on the result of client.Ping.
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("failed to ping MongoDB: %v", err)
	}

	log.Println("---- Mongo Health Check OK ----")
	return &MongoClient{MongoClient: client}

}
