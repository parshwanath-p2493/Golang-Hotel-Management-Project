package utils

import (
	"context"
	"log"
	"time"

	"github.com/parshwanath-p2493/Project/database"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteAuth(uuid string, role string, id string) (int64, error) {
	collection := database.OpenCollection(role)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{id: uuid}
	update := bson.M{
		"$unset": bson.M{
			//"email": uuid,
			"token": "",
		},
	}
	log.Println(uuid)
	log.Println(id)
	log.Println(update)

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	log.Println("The token is removed successfuly")

	return result.ModifiedCount, nil
}
