package driver 

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MongoDBURL           = "mongodb://localhost:27017"
	MongoDBName          = "e-commerce"
	MongoDBCollection    = "user"
	ConnectionTimeoutSec = 10
)

func ConnectToMongoDb() (*mongo.Database) {
	ctx, cancel := context.WithTimeout(context.Background(), ConnectionTimeoutSec*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(MongoDBURL)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client.Database(MongoDBName)
}