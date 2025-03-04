package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/database"
	"github.com/parshwanath-p2493/Project/models"
	"github.com/parshwanath-p2493/Project/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignUpAdmin(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var admin models.Admin
	collection := database.OpenCollection("Admin")

	if err := c.BodyParser(&admin); err != nil {
		// Handle error during adding
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Invalid request body"))
	}

	// Generate a new ObjectID for the admin
	admin.ID = primitive.NewObjectID()
	admin.Admin_id = admin.ID.Hex()

	// Insert the admin document into the database
	result, err := collection.InsertOne(ctx, admin)
	if err != nil {
		// Handle error during insertion
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Error inserting admin into database"))
	}

	// Return success response with the result of insertion
	return c.Status(http.StatusCreated).JSON(utils.Response(c, result, "Operation completed successfully"))
}
