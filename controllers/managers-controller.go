package controllers

import (
	"context"
	"fmt"
	"log"
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
	hashedPassword, err := helpers.HashPassword(manager.Password)
	if err != nil {
		utils.Error(c, utils.InternalServerError, "Cannot generate the password")
	}
	manager.Password = hashedPassword
	manager.Created_time = time.Now()
	manager.Role = "Manager"

	token, err := helpers.GenerateToken(manager.First_name, manager.Email, manager.Role, manager.Department)
	if err != nil {
		// Handle error during token generation
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Failed to generate token"))
	}
	manager.Token = token
	result, err := collection.InsertOne(ctx, manager)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, err.Error()))
	}
	response := fiber.Map{
		"message": token,                                                   // From utils.Message (you can adjust this as needed)
		"data":    utils.Response(c, result, "Added successfully")["data"], // Get "data" from utils.Response
	}

	return c.Status(http.StatusCreated).JSON(response)
	// return c.Status(http.StatusCreated).JSON(utils.Response(c, result, "Manager added successfully"))
}
func ManagerLogin(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var ExistedManager models.Managers
	var LoginManager models.Managers
	collection := database.OpenCollection("Managers")
	if err := c.BodyParser(&LoginManager); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Invalid Syntax"))
	}
	err := collection.FindOne(ctx, bson.M{"email": LoginManager.Email}).Decode(&ExistedManager)
	if err == mongo.ErrNoDocuments {
		return c.Status(http.StatusUnauthorized).JSON(utils.Error(c, utils.Unauthorized, "Invalid Password or Email"))
	} else if err != nil {
		// Database error
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Error fetching manager from database"))
	}

	_, err = helpers.VerifyPassword(LoginManager.Password, ExistedManager.Password)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(utils.Error(c, utils.Unauthorized, "Wrong Password"))
	}
	log.Println("\n \n Password Matched Login Continue.... ")
	LoginManager.Role = "Manager"
	fmt.Println("\n \n ", LoginManager.Role, LoginManager.Department)

	token, err := helpers.GenerateToken(ExistedManager.First_name, ExistedManager.Email, ExistedManager.Role, ExistedManager.Department)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Failed to generate token"))
	}
	ExistedManager.Updated_time = time.Now().Local()
	filter := bson.M{"manager_id": ExistedManager.Manager_id}
	update := bson.M{
		"$set": bson.M{
			"email":        ExistedManager.Email,
			"token":        token,
			"updated_time": ExistedManager.Updated_time,
		},
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Token not updated  "))
	}
	log.Println("Refreshed token added to the Mongosuccessfuly", result)
	response := fiber.Map{
		"message": token,                                                             // From utils.Message (you can adjust this as needed)
		"data":    utils.Response(c, ExistedManager, "Log in  successfully")["data"], // Get "data" from utils.Response
	}
	return c.Status(http.StatusOK).JSON(response)

}

func DeleteManager(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	managerID := c.Params("id")
	filter := bson.M{"manager_id": managerID}
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
func GetManager(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := database.OpenCollection("Managers")
	var manager []models.Managers
	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Error in fetching data"))
	}
	defer result.Close(ctx)
	if err := result.All(ctx, &manager); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Unable to fettch the data "))
	}
	return c.Status(http.StatusOK).JSON(utils.Response(c, manager, "The Manager details are:"))
}

func LogOutManager(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var manager models.Managers
	id := c.Params("id")
	collection := database.OpenCollection("Managers")
	err := collection.FindOne(ctx, bson.M{"manager_id": id}).Decode(&manager)
	if err == mongo.ErrNilDocument {
		return c.Status(http.StatusNotFound).JSON(utils.Error(c, utils.NotFound, "Not data found "))
	} else if err != nil {
		// Database error
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Error fetching manager from database"))
	}
	log.Println(manager.Email)
	log.Print(manager.First_name)

	// clientToken := c.Get("X-Auth-Token", "")
	// clientToken = strings.Replace(clientToken, "Bearer ", "", 1)
	// claims := &helpers.Info{}
	// _, err2 := jwt.ParseWithClaims(clientToken, claims, func(t *jwt.Token) (interface{}, error) {
	// 	return []byte(os.Getenv("SECRET_KEY")), nil
	// })
	// if err2 != nil {
	// 	log.Println("Error Parsing token in LOGOUT SESSION ", err2)
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "INVALID OR EXPIRED TOKEN"})
	// } else {
	deleted, delErr := utils.DeleteAuth(manager.Manager_id, "Managers", "manager_id")
	if delErr != nil {
		log.Println("Error invalidating the token Metadata")
	}
	if deleted == 0 {
		log.Println("No active Session Found ")
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, delErr.Error()))
	}

	cookie := fiber.Cookie{
		Name:    "jwt", // Name of the cookie
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	}
	c.Cookie(&cookie)
	c.Set("X-Auth-Token", "") //erase the token from request header
	role := c.Locals("role")
	if role != nil {
		c.Locals("role", nil)
	}
	return c.Status(fiber.StatusOK).JSON(utils.Response(c, role, "Logged Out Successfully"))
}
