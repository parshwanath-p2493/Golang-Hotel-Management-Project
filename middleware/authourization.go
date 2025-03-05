package middleware

import (
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/parshwanath-p2493/Project/helpers"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

// func (allowedRoles ...string) fiber.Handler {
func ManagerAuthentication(c *fiber.Ctx) error {
	clientToken := c.Get("X-Auth-Token")
	clientToken = strings.Replace(clientToken, "Bearer ", "", 1)
	if clientToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Token",
		})
	}
	// 	// Parse the token

	claims := &helpers.Info{}
	_, err := jwt.ParseWithClaims(clientToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	log.Printf("SECRETKEY: %s", SECRET_KEY)
	if err != nil {
		log.Println("Error Parsing token", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "INVALID OR EXPIRED TOKEN"})
	}
	if claims.Role != "manager" && claims.Role != "Manager" && claims.Role != "MANAGER" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access forbidden Only Manager Can access "})
	}
	c.Locals("department", claims.Department)
	return c.Next()
}

/*

package middleware

import (
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/parshwanath-p2493/Project/helpers"
)

func ManagerAuth(c *fiber.Ctx) error {
	clientToken := c.Get("X-Auth-Token")
	clientToken = strings.Replace(clientToken, "Bearer ", "", 1)
	if clientToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Token",
		})
	}

	claims := &helpers.Info{}
	_, err := jwt.ParseWithClaims(clientToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		log.Println("Error Parsing token", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "INVALID OR EXPIRED TOKEN"})
	}

	// Ensure the user is a Manager
	if claims.Role != "manager" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access Denied: Only Managers can access",
		})
	}

	// Store department in context
	c.Locals("department", claims.Department)

	return c.Next()
}



*/

func AdminAuthentication(c *fiber.Ctx) error {
	clientToken := c.Get("X-Auth-Token")
	clientToken = strings.Replace(clientToken, "Bearer ", "", 1)
	if clientToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Token",
		})
	}
	// 	// Parse the token

	claims := &helpers.Info{}
	_, err := jwt.ParseWithClaims(clientToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		log.Println("Error Parsing token", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "INVALID OR EXPIRED TOKEN"})
	}
	if claims.Role != "admin" && claims.Role != "Admin" && claims.Role != "ADMIN" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access forbidden Only Admin Can access "})
	}
	c.Locals("role", claims.Role)
	return c.Next()

}
