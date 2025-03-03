package utils

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// 500, 404, 400, 200, 409, 401
var InternalServerError = http.StatusInternalServerError // for error while id convertion
var NotFound = http.StatusNotFound                       // for id not found
var BadRequest = http.StatusBadRequest                   // for invalid json or empty data
var OK = http.StatusOK                                   // for ok status
var Conflict = http.StatusConflict                       // for inserting data already exist like email, password
var Unauthorized = http.StatusUnauthorized               // Email or password incorrect

func Error(c *fiber.Ctx, statuscode int, errormessage string) fiber.Map {
	log.Printf("Error:%s", errormessage)
	c.Set("Content-Type", "application/json")
	c.Status(statuscode)
	return fiber.Map{"error": errormessage}
}
func Message(c *fiber.Ctx, message string) {
	c.Set("Content-Type", "application/json")
	c.JSON(fiber.Map{"message": message})
}
func Response(c *fiber.Ctx, data interface{}) fiber.Map {
	c.Set("Content-Type", "apllication/json")
	return fiber.Map{"data": data}
}
