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

func ManagerSignup(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var manager models.Managers
	collection := database.OpenCollection("Managers")
	manager.ID = primitive.NewObjectID()
	manager.Manager_id = manager.ID.Hex()
	if err := c.BodyParser(&manager); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Invalid Synatax"))
	}
	result, err := collection.InsertOne(ctx, manager)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, err.Error()))
	}
	return c.Status(http.StatusCreated).JSON(utils.Response(c, result, "Operation completed successfully"))
}
func DeleteManager(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	managerID := c.Params("id")
	filter := bson.M{"id": managerID}
	collection := database.OpenCollection("Managers")
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, err.Error()))
	}
	if result.DeletedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(utils.Error(c, utils.NotFound, "Not data found "))
	}
	return c.Status(http.StatusCreated).JSON(utils.Response(c, result, "Operation completed successfully"))
}
