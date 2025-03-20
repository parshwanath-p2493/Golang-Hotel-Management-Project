package utils

import (
	"context"
	"log"
	"time"

	"github.com/parshwanath-p2493/Project/database"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteAuth(uuid string, role string) (int64, error) {
	collection := database.OpenCollection(role)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"email": uuid}
	update := bson.M{
		"$set": bson.M{
			"token": "",
		},
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	log.Println("The token is removed successfuly")

	return result.ModifiedCount, nil
}
