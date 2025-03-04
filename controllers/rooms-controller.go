package controllers

import (
	"context"
	"net/http"
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
func GetAllRooms(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var rooms []models.Rooms
	collection := database.OpenCollection("Rooms")
	result, err := collection.Find(ctx, rooms)
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
