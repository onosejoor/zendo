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
	client   *mongo.Database
	DB       *mongo.Client
	clientMu sync.Mutex
)

// GetClient returns a singleton MongoDB Database
func GetClient() *mongo.Database {
	clientMu.Lock()
	defer clientMu.Unlock()

	if client != nil {
		return client
	}

	MONGODB_URL := os.Getenv("MONGODB_URL")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MONGODB_URL).SetServerAPIOptions(serverAPI)

	var err error
	conn, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	if err := conn.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	client = conn.Database("zendo")

	log.Println("Connected to MongoDB!")

	return client
}

func GetClientWithoutDB() *mongo.Client {
	clientMu.Lock()
	defer clientMu.Unlock()

	if DB != nil {
		return DB
	}

	MONGODB_URL := os.Getenv("MONGODB_URL")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MONGODB_URL).SetServerAPIOptions(serverAPI)

	var err error
	DB, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	if err := DB.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	return DB
}
