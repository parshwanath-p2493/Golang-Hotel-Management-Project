package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/database"
	"github.com/parshwanath-p2493/Project/models"
	"github.com/parshwanath-p2493/Project/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddRooms(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var room models.Rooms
	collection := database.OpenCollection("Rooms")
	if err := c.BodyParser(&room); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Invalid synatx"))
	}

	room.ID = primitive.NewObjectID()
	room.Room_id = room.ID.Hex()
	result, err := collection.InsertOne(ctx, room)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Error"))
	}
	return c.Status(http.StatusOK).JSON(utils.Response(c, result, "Operation completed successfully"))
}
func DeleteRoom(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	RoomNumber := c.Params("room_number")
	// filter := bson.M{
	// 	"room_number": RoomNumber}

	roomNumberInt, err := strconv.ParseInt(RoomNumber, 10, 32)
	if err != nil {
		// If conversion fails, return a bad request response
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Invalid room number format"))
	}

	// Build the filter to match the room_number in the database
	filter := bson.M{
		"room_number": int32(roomNumberInt), // Convert to int32 to match the MongoDB schema
	}
	fmt.Println("Filter:", filter)

	collection := database.OpenCollection("Rooms")
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Unable to delete"))
	}
	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(utils.Error(c, utils.NotFound, "Room number in incorrect"))
	}
	return c.Status(http.StatusOK).JSON(utils.Response(c, result, "Deleted Succssfuly"))
}
func GetAllRooms(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var rooms []models.Rooms
	collection := database.OpenCollection("Rooms")
	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Error"))
	}
	defer result.Close(ctx)
	if err := result.All(ctx, &rooms); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Failed to fetch the data"))
	}
	return c.Status(http.StatusOK).JSON(utils.Response(c, rooms, "Operation completed successfully"))
}

// Accesaable to the users or guestes who are booking the rooms
func FilterRooms(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var rooms []models.Rooms
	capacity := c.Query("capacity")
	RoomType := c.Query("room_type")
	sortByPrice := c.Query("sortByPrice")
	minPrice := c.Query("min_price")
	maxPrice := c.Query("max-price")
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 5)
	collection := database.OpenCollection("Rooms")
	filter := bson.M{}

	if RoomType != "" {
		filter["room_type"] = RoomType
	}
	if capacity != "" {
		filter["capacity"] = capacity
	}
	//Implementing Range of price in the Project for good user experience
	if minPrice != "" || maxPrice != "" {
		minPriceVal, maxPriceVal := 0.0, 0.0
		var PriceFilter bson.M
		if minPrice != "" {
			minPriceVal, _ = strconv.ParseFloat(minPrice, 64)
		}
		if maxPrice != "" {
			maxPriceVal, _ = strconv.ParseFloat(maxPrice, 64)
		}
		if minPrice != " " && maxPrice != "" {
			PriceFilter = bson.M{
				"$gte": minPriceVal,
				"$lte": maxPriceVal,
			}
		} else if minPrice != "" {
			PriceFilter = bson.M{
				"$gte": minPriceVal,
			}
		} else if maxPrice != "" {
			PriceFilter = bson.M{
				"$lte": maxPriceVal,
			}
		}
		filter["price"] = PriceFilter
	}

	var SortOrder int
	if sortByPrice == "asc" {
		SortOrder = 1
	} else if sortByPrice == "des" {
		SortOrder = -1
	} else {
		SortOrder = 1
	}

	sort := bson.M{"price": SortOrder}
	result, err := collection.Find(ctx, filter, options.Find().SetSort(sort).SetSkip(int64(page)).SetLimit(int64(limit)))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Unable to get data"))
	}
	defer result.Close(ctx)
	if err := result.All(ctx, &rooms); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, err.Error()))
	}
	return c.Status(http.StatusOK).JSON(utils.Response(c, rooms, "Operation completed successfully"))

}
func UpdateRoomStatus(id string, status models.Availability_status) error {
	// ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	// defer cancel()
	updated_at := time.Now()
	update := bson.M{
		"$set": bson.M{
			"availability_status": status,
			"updated_time":        updated_at,
		},
	}
	collection := database.OpenCollection("Rooms")
	upsert := true
	filter := bson.M{"room_id": id}
	options := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update, &options)
	if err != nil {
		log.Println(utils.Error(&fiber.Ctx{}, utils.InternalServerError, err.Error()))
		return err
	}
	return nil
}
