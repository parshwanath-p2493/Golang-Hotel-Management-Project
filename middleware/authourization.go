package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/parshwanath-p2493/Project/helpers"
)

// func (allowedRoles ...string) fiber.Handler {
func ManagerAuthentication(c *fiber.Ctx) error {

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	//var SECRET_KEY string = os.Getenv("SECRET_KEY")

	clientToken := c.Get("X-Auth-ManagerToken")
	clientToken = strings.Replace(clientToken, "Bearer ", "", 1)
	if clientToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Token Manager",
		})
	}

	// 	// Parse the token
	log.Printf("SECRETKEY: %s", os.Getenv("SECRET_KEY"))

	claims := &helpers.Info{}
	_, err := jwt.ParseWithClaims(clientToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil //this was the error
	})
	//	fmt.Println("Role of from token:", claims.Role)
	// log.Println("Role from token:", claims.Role)
	// log.Println("Department from token:", claims.Department)

	if err != nil {
		log.Println("Error Parsing token", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "INVALID OR EXPIRED TOKEN"})
	}
	// if err.Error() == "signature is invalid" {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": "Invalid Token Signature",
	// 	})
	// }
	log.Println("Name of the manager: ", claims.Name)
	log.Println("Email of the manager:", claims.Email)
	log.Println("Email of the manager:", claims.Role)
	log.Println("Email of the manager:", claims.Department)

	if claims.Role != "manager" && claims.Role != "Manager" && claims.Role != "MANAGER" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access forbidden Only Manager Can access "})
	}

	c.Locals("department", claims.Department)
	c.Locals("role", claims.Role)
	return c.Next()
}

func AdminAuthentication(c *fiber.Ctx) error {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Println("Unable load the seceret key")
	// }
	// var SECRET_KEY = os.Getenv("SECRET_KEY")

	//  Noo need because we used Getenv in inline itself
	clientToken := c.Get("X-Auth-Token")
	clientToken = strings.Replace(clientToken, "Bearer ", "", 1)
	if clientToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Token of ADMIN",
		})
	}
	// 	// Parse the token

	claims := &helpers.Info{}
	_, err2 := jwt.ParseWithClaims(clientToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err2 != nil {
		log.Println("Error Parsing token", err2)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "INVALID OR EXPIRED TOKEN"})
	}

	log.Println("Name of the Admin: ", claims.Name)
	log.Println("Email of the Admin:", claims.Email)
	log.Println("Role from token:", claims.Role)
	log.Println("Department from token:", claims.Department)

	if claims.Role != "admin" && claims.Role != "Admin" && claims.Role != "ADMIN" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access forbidden Only Admin Can access "})
	}
	//fmt.Print("\n \n")
	//fmt.Println(claims.Role)
	c.Locals("role", claims.Role)
	return c.Next()

}

/*
func GuestAuth(c *fiber.Ctx) error {
	clientToken := c.Get("X-Auth-GuestToken")
	clientToken = strings.Replace(clientToken, "Bearer", "", 1)
	if clientToken == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Token of Guest",
		})
	}
	claims := &helpers.Info{}
	_, err := jwt.ParseWithClaims(clientToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		log.Println("Error Parsing token", err)
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Token is mismatched or Token is expired login again "})
	}
	log.Println("Name of the guest", claims.Name)
	log.Println("Email of the guest:", claims.Email)
	if claims.Role != "Guest" && claims.Role != "guest" && claims.Role != "GUEST" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access forbidden Only Guest Can access "})
	}
	c.Locals("role", claims.Role)
	return c.Next()

}
*/

func GuestAuth(c *fiber.Ctx) error {
	clientToken := c.Get("X-Auth-GuestToken")
	// Clean the token by removing the Bearer prefix if present
	clientToken = strings.TrimPrefix(clientToken, "Bearer ")

	// Check if the token is empty
	if clientToken == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Token of Guest",
		})
	}

	// Debug: Print the token to verify its format
	log.Printf("Received Token: %s", clientToken)

	// Parse the token
	claims := &helpers.Info{}
	_, err := jwt.ParseWithClaims(clientToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	// Check if there was an error parsing the token
	if err != nil {
		log.Printf("Error Parsing Token: %s", err.Error())
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token is mismatched or Token is expired. Please login again.",
		})
	}

	log.Printf("Name of the guest: %s", claims.Name)
	log.Printf("Email of the guest: %s", claims.Email)

	// Check if the role is correct
	if claims.Role != "Guest" && claims.Role != "guest" && claims.Role != "GUEST" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access forbidden. Only Guest can access.",
		})
	}

	c.Locals("role", claims.Role)
	return c.Next()
}
