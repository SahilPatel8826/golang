package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

// ConnectDB connects to MongoDB Atlas and returns a client instance
func ConnectDB() *mongo.Client {
	// MongoDB Atlas URI (use env vars in production!)
	uri := "link here"

	// Configure Stable API
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client
	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err)
	}

	// Context with timeout to prevent hanging connections
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ping to test connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("Could not ping MongoDB: %v", err)
	}

	fmt.Println("✅ Successfully connected and pinged MongoDB!")
	return client
}
