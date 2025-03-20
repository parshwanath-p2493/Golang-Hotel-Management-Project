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
	filter := bson.M{"admin_id": admin.Admin_id}
	update := bson.M{
		"$set": bson.M{
			"email": admin.Email,
			"token": token,
		},
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Token not updated  "))
	}
	log.Println("Refreshed token added to the Mongosuccessfuly", result)

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

func LogOutAdmin(c *fiber.Ctx) error {
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
		deleted, delErr := utils.DeleteAuth(claims.Email, "Admin")
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

/**

updateObj := []bson.E{
	// Start with the required fields for update (email, token)
	{Key: "email", Value: admin.Email},
	{Key: "token", Value: token},
}

// Conditionally add the "updated_at" field if necessary
updateObj = append(updateObj, bson.E{Key: "updated_at", Value: branch.Updated_at})

// Create the update document
update := bson.M{
	"$set": updateObj,  // Use the dynamically built update object
}

filter := bson.M{"admin_id": admin.Admin_id}

result, err := collection.UpdateOne(ctx, filter, update)
if err != nil {
	return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Token not updated"))
}

log.Println("Refreshed token added to the MongoDB successfully", result)

return c.Status(http.StatusOK).JSON(utils.Response(c, token, "Successfully Logged IN"))

**/
