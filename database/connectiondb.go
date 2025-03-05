// package database

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/joho/godotenv"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var db *mongo.Client

// func connectDB() *mongo.Client {

// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	uri := os.Getenv("DB_URI")
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Mongo Conected Successfully")
// 	db = client
// 	return db
// }

// var client *mongo.Client = connectDB()

//	func OpenCollection(collectionName string) *mongo.Collection {
//		return client.Database("Hotel-Management").Collection(collectionName)
//	}
// package database

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/joho/godotenv"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var client *mongo.Client

// func ConnectDB() *mongo.Client {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	uri := os.Getenv("DB_URI")
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Mongo Connected Successfully")
// 	return client
// }

// func OpenCollection(collectionName string) *mongo.Collection {
// 	return client.Database("Hotel-Management").Collection(collectionName)
// }

package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// ConnectDB initializes the connection to the MongoDB server
func ConnectDB() *mongo.Client {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get MongoDB URI from environment variable
	uri := os.Getenv("DB_URI")
	fmt.Println(uri)
	// Set up a timeout context for the MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB and assign the global client variable
	//var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri)) // Use '=' instead of ':=' to assign to global variable
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Confirm connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	fmt.Println("MongoDB connected successfully")

	// Return the MongoDB client
	return client
}

// OpenCollection opens a specific MongoDB collection
func OpenCollection(collectionName string) *mongo.Collection {
	// Check if client is initialized before using it
	if client == nil {
		log.Fatal("MongoDB client is not initialized")
	}
	// Access the "Hotel-Management" database and return the collection
	return client.Database("Hotel-Management").Collection(collectionName)
}
