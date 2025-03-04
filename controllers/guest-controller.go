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
)

func GuestSignup(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var guest models.Guest
	collection := database.OpenCollection("Guest")
	if err := c.BodyParser(&guest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Unformatted data"))
	}
	guest.ID = primitive.NewObjectID()
	guest.Guest_id = guest.ID.Hex()

	result, err := collection.InsertOne(ctx, guest)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Failed to add data"))
	}
	return c.Status(http.StatusOK).JSON(utils.Response(c, result, "Operation completed successfully"))
}

//	func GuestLogin(c *fiber.Ctx) error {
//		return c.Status(http.StatusOK)
//	}
func GetAllGuest(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var guests []models.Guest
	defer cancel()
	collection := database.OpenCollection("Guest")
	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, err.Error()))
	}
	return c.Status(http.StatusOK).JSON(utils.Response(c, result, "Operation completed successfully"))
	defer result.Close(ctx)
	if err := result.All(ctx, &guests); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Failed fetch data"))
	}
	return c.Status(http.StatusOK).JSON(utils.Response(c, guests, "Operation completed successfully"))
}
