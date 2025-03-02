package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/database"
	"github.com/parshwanath-p2493/Project/routes"
)

func main() {

	database.ConnectDB()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "5050"
	}
	//r := gin.Default()
	//r.Run(":" + PORT)
	//r.Run(":" + PORT)
	r := fiber.New()

	r.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("THE SERVER IS ALL SET GO ")
		return c.JSON(&fiber.Map{"message": "SERVER RUNNING in PORT : " + PORT})
	})
	routes.AdminRoutes(r)
	log.Fatal(r.Listen(":" + PORT))
}
