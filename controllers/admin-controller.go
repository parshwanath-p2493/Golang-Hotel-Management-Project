package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/database"
	"github.com/parshwanath-p2493/Project/helpers"
	"github.com/parshwanath-p2493/Project/models"
	"github.com/parshwanath-p2493/Project/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SignUpAdmin(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var admin models.Admin
	collection := database.OpenCollection("Admin")

	if err := c.BodyParser(&admin); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Invalid request body"))
	}
	hashedPassword, err := helpers.HashPassword(admin.Password)
	if err != nil {
		utils.Error(c, utils.InternalServerError, "Cant generate the password")
	}
	admin.Password = hashedPassword

	admin.ID = primitive.NewObjectID()
	admin.Admin_id = admin.ID.Hex()
	admin.Role = "ADMIN"
	//admin.created_time = time.Now()
	admin.Created_time = time.Now()

	token, err := helpers.GenerateToken(admin.First_name, admin.Email, admin.Role, "")
	if err != nil {
		// Handle error during token generation
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Failed to generate token"))
	}
	admin.Token = token
	result, err := collection.InsertOne(ctx, admin)
	if err != nil {
		// Handle error during insertion
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Error inserting admin into database"))
	}
	// fiber.Map={utils.Message(c,token)}
	// return c.Status(http.StatusCreated).JSON(utils.Response(c, result, "Added successfully"))

	response := fiber.Map{
		"message": token,
		"data":    utils.Response(c, result, "Added successfully")["data"],
	}

	return c.Status(http.StatusCreated).JSON(response)

}
func LoginAdmin(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var admin models.Admin
	var input models.Admin
	collection := database.OpenCollection("Admin")
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Invalid request body"))
	}
	err := collection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&admin)
	if err == mongo.ErrNilDocument {
		return c.Status(http.StatusUnauthorized).JSON(utils.Error(c, utils.Unauthorized, "Invalid Password or email  "))
	} else if err != nil {
		return c.Status(http.StatusNotFound).JSON(utils.Error(c, utils.NotFound, "Error fetching Admin from database  "))
	}
	// if _, err := helpers.VerifyPassword(input.Password, admin.Password); err != nil {
	// 	c.Status(http.StatusUnauthorized).JSON(utils.Error(c, utils.Unauthorized, "Invalid  Password"))
	// }

	_, err = helpers.VerifyPassword(input.Password, admin.Password)
	if err != nil {
		// Handle incorrect password
		return c.Status(http.StatusUnauthorized).JSON(utils.Error(c, utils.Unauthorized, "Invalid credentials"))
	}
	token, err := helpers.GenerateToken(admin.First_name, admin.Email, admin.Role, "")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Failed to generate token"))
	}
	return c.Status(http.StatusOK).JSON(utils.Response(c, token, "Successfuly Logged IN "))
}

/*
func LogOut(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now().Local().Add(-time.Hour),
	}
	c.Cookie(&cookie)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Logged out Successfuly",
	})
}
*/

func LogOut(c *fiber.Ctx) error {
	role := c.Locals("role")
	c.ClearCookie("jwt")
	c.Set("X-Auth-Token", "")
	if role != nil {
		c.Locals("role", nil)
	}
	return c.Status(fiber.StatusOK).JSON(utils.Response(c, role, "Logged Out Successfully"))
}
