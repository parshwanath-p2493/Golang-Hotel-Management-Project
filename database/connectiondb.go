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

func ConnectDB() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	uri := os.Getenv("DB_URI")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo Connected Successfully")
	return client
}

func OpenCollection(collectionName string) *mongo.Collection {
	return client.Database("Hotel-Management").Collection(collectionName)
}
