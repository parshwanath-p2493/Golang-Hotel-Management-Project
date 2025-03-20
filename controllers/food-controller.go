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

func AddFood(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var food models.Food
	collection := database.OpenCollection("Food")
	if err := c.BodyParser(&food); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "err"))
	}
	food.ID = primitive.NewObjectID()
	food.Food_id = food.ID.Hex()
	food.Created_time = time.Now()
	result, err := collection.InsertOne(ctx, food)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "err"))
	}
	return c.Status(http.StatusCreated).JSON(utils.Response(c, result, "Operation completed successfully"))
}

func GetFood(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var food []models.Food
	category := c.Query("category")
	sortByPrice := c.Query("sortByPrice")
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 5)
	filter := bson.M{}
	if category != "" {
		filter["category_name"] = category
	}
	var SortOrder int
	if sortByPrice == "asc" {
		SortOrder = 1
	} else if sortByPrice == "desc" {
		SortOrder = -1
	} else {
		SortOrder = 1
	}
	sort := bson.M{"price": SortOrder}
	collection := database.OpenCollection("Food")
	cursor, err := collection.Find(ctx, filter, options.Find().
		SetSort(sort).
		SetSkip(int64((page-1)*limit)).
		SetLimit(int64(limit)))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Unable get data"))
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &food); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, err.Error()))
	}
	return c.Status(http.StatusOK).JSON(utils.Response(c, food, "Operation completed successfully"))
}
func ChangeFood(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	CategoryName := c.Params("category_name")
	FoodName := c.Params("food_name")
	var UpdateFood models.Food
	if err := c.BodyParser(&UpdateFood); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Invalid input"))
	}
	filter := bson.M{
		"category_name": CategoryName,
		"item_name":     FoodName,
	}
	// update object with the fields that can be changed
	update := bson.M{
		"$set": bson.M{
			"food_name":     UpdateFood.Item_name,
			"category_name": UpdateFood.Category_name,
			"vegornonveg":   UpdateFood.VegorNonveg,
			"price":         UpdateFood.Price,
			"update_time":   time.Now(),
		},
	}
	collection := database.OpenCollection("Food")
	UpdateFood.Updated_time = time.Now()
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Failed to UPDATE"))
	}
	if result.MatchedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(utils.Error(c, utils.NotFound, "Food Item Not Found"))
	}
	return c.Status(http.StatusOK).JSON(utils.Response(c, result, "Operation completed successfully"))
}

func DeleteFood(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	CategoryName := c.Params("category_name")
	FoodName := c.Params("food_name")
	//var DeleteFood models.Food
	filter := bson.M{
		"category_name": CategoryName,
		"item_name":     FoodName,
	}
	collection := database.OpenCollection("Food")
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Failed to Delete"))
	}
	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(utils.Error(c, utils.NotFound, "Data not Found"))
	}
	return c.Status(http.StatusOK).JSON(utils.Response(c, result, "Item Deleted successfully"))

}
