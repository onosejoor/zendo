package db

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	client   *mongo.Client
	clientMu sync.Mutex
)

// GetClient returns a singleton MongoDB Database
func GetClient() *mongo.Database {
	clientMu.Lock()
	defer clientMu.Unlock()

	if client != nil {
		return client.Database(os.Getenv("DATABASE"))
	}

	MONGODB_URL := os.Getenv("MONGODB_URL")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MONGODB_URL).SetServerAPIOptions(serverAPI)

	var err error
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")

	return client.Database(os.Getenv("DATABASE"))
}

func GetClientWithoutDB() *mongo.Client {
	clientMu.Lock()
	defer clientMu.Unlock()

	if client != nil {
		return client
	}

	MONGODB_URL := os.Getenv("MONGODB_URL")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MONGODB_URL).SetServerAPIOptions(serverAPI)

	var err error
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	return client
}
