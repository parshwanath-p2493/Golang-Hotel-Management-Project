// package controllers

// import (
// 	"context"
// 	"net/http"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/parshwanath-p2493/Project/database"
// 	"github.com/parshwanath-p2493/Project/models"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func SignUpAdmin(c *fiber.Ctx) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	var admin models.Admin
// 	defer cancel()
// 	collection := database.OpenCollection("Admin")
// 	if err := c.BodyParser(&admin); err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
// 	}

// 	admin.ID = primitive.NewObjectID()
// 	admin.Admin_id = admin.ID.Hex()
// 	result, err := collection.InsertOne(ctx, admin)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponses{
// 			Status:  http.StatusInternalServerError,
// 			Message: "Error",
// 			Data:    &fiber.Map{"data": err.Error()}})
// 	}
// 	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
// }

package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/database"
	"github.com/parshwanath-p2493/Project/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignUpAdmin(c *fiber.Ctx) error {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the admin struct
	var admin models.Admin

	// Open MongoDB collection
	collection := database.OpenCollection("Admin")

	// Parse the request body into the admin struct
	if err := c.BodyParser(&admin); err != nil {
		// Return a BadRequest response if parsing fails
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "error",
			"data":    fiber.Map{"data": err.Error()},
		})
	}

	// Generate ObjectID for the new admin
	admin.ID = primitive.NewObjectID()
	admin.Admin_id = admin.ID.Hex()

	// Insert the new admin into MongoDB
	result, err := collection.InsertOne(ctx, admin)
	if err != nil {
		// Return an InternalServerError response if insertion fails
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "error",
			"data":    fiber.Map{"data": err.Error()},
		})
	}

	// Return a success response after successfully creating the admin
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "success",
		"data":    fiber.Map{"data": result},
	})
}
