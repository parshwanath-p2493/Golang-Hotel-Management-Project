package controllers

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/database"
	"github.com/parshwanath-p2493/Project/helpers"
	"github.com/parshwanath-p2493/Project/models"
	"github.com/parshwanath-p2493/Project/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GuestSignup(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var guest models.Guest
	collection := database.OpenCollection("Guest")
	if err := c.BodyParser(&guest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Unformatted data"))
	}
	hashedPassword, err := helpers.HashPassword(guest.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Fill Passowrd please "))
	}
	token, err := helpers.GenerateToken(guest.First_Name, guest.Email, "Guest", "Guest")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Unable to generate token"))
	}
	guest.ID = primitive.NewObjectID()
	guest.Guest_id = guest.ID.Hex()
	guest.Password = hashedPassword
	guest.Token = token
	guest.Created_time = time.Now()

	result, err := collection.InsertOne(ctx, guest)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Failed to add data"))
	}
	response := fiber.Map{
		"message": token,
		"data":    utils.Response(c, result, "Guest Created Successfully ")["data"],
	}
	return c.Status(http.StatusOK).JSON(response)
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
	//return c.Status(http.StatusOK).JSON(utils.Response(c, result, "Operation completed successfully"))
	defer result.Close(ctx)
	if err := result.All(ctx, &guests); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Failed fetch data"))
	}
	return c.Status(http.StatusOK).JSON(utils.Response(c, guests, "Guest Data and details fetched successfully"))
}
func GuestLogin(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := database.OpenCollection("Guest")
	var LoginGuest models.Guest
	var ExistingGuest models.Guest
	if err := c.BodyParser(&LoginGuest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Enter both Username or Passwrod"))
	}
	err := collection.FindOne(ctx, bson.M{"email": LoginGuest.Email}).Decode(&ExistingGuest)
	if err == mongo.ErrNilDocument {
		return c.Status(http.StatusUnauthorized).JSON(utils.Error(c, utils.Unauthorized, "Invalid Password or email "))
	} else if err != nil {
		return c.Status(http.StatusNotFound).JSON(utils.Error(c, utils.NotFound, "Error fetching Guest from database  "))
	}
	_, err = helpers.VerifyPassword(LoginGuest.Password, ExistingGuest.Password)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(utils.Error(c, utils.Unauthorized, "Enter corrrect password ,Password is wrong "))
	}
	log.Println("\n", ExistingGuest.First_Name, ExistingGuest.Last_Name)
	log.Println("\n \n Password Matched Login Continue.... ")
	token, err := helpers.GenerateToken(ExistingGuest.First_Name, ExistingGuest.Email, "Guest", "Guest")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Failed to generate token"))
	}

	filter := bson.M{"guest_id": ExistingGuest.Guest_id}
	update := bson.M{
		"$set": bson.M{
			"email": ExistingGuest.Email,
			"token": token,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Token not updated  "))
	}
	log.Println("\n Refreshed token added to the Mongosuccessfuly", result)

	response := fiber.Map{
		"message": token,
		"data":    utils.Response(c, ExistingGuest, "Login Successfully")["data"],
	}
	return c.Status(http.StatusOK).JSON(response)
}

func LogOutGuest(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:    "jwt", // Name of the cookie
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	}
	c.Cookie(&cookie)
	role := c.Locals("role")
	//c.ClearCookie("jwt")
	c.Set("X-Auth-Token", "") //erase the token from request header
	clientToken := c.Get("X-Auth-Token")
	clientToken = strings.Replace(clientToken, "Bearer ", "", 1)
	claims := &helpers.Info{}
	_, err2 := jwt.ParseWithClaims(clientToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err2 != nil {
		log.Println("Error Parsing token in LOGOUT SESSION ", err2)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "INVALID OR EXPIRED TOKEN"})
	} else {
		deleted, delErr := utils.DeleteAuth(claims.Email, "Guest", "guest_id")
		if delErr != nil {
			log.Println("Error invalidating the token Metadata")
		}
		if deleted == 0 {
			log.Println("No active Session Found ")
			return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, delErr.Error()))
		}
	}

	if role != nil {
		c.Locals("role", nil)
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response(c, role, "Logged Out Successfully"))
}
