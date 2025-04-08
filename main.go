package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/parshwanath-p2493/Project/database"
	_ "github.com/parshwanath-p2493/Project/docs"
	"github.com/parshwanath-p2493/Project/routes"
	// Correct Swagger import
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:2493
// @BasePath /
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
	a := os.Getenv("XYZ") //for testing purpose....
	fmt.Println(a)
	r.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("THE SERVER IS ALL SET GO ")
		return c.JSON(&fiber.Map{"message": "SERVER RUNNING in PORT : " + PORT})
	})
	//r.Get("/swagger/*", swagger.HandlerDefault)

	r.Static("/docs", "./docs")
	r.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "/docs/swagger.json", // Match the actual path where it was generated
	}))

	routes.AdminRoutes(r)
	routes.FoodRoutes(r)
	routes.GuestRoutes(r)
	routes.ManagerRoutes(r)
	routes.RoomsRoutes(r)
	routes.StaffRoutes(r)
	routes.AdminRoutesAuth(r)
	routes.FoodRoutes2(r)
	routes.GuestRoutes2(r)
	routes.StaffRoutes2(r)
	routes.RoomsRoutesAuth(r)
	routes.RoomsRoutesAuthManager(r)
	routes.CreateBooking(r)
	routes.NotificationRoutes(r)
	routes.SetupRoutes(r)

	//r.Get("/swagger/*", fiberSwagger.WrapHandler)
	//We need to call  all routes before starting the server else it will be error
	log.Fatal(r.Listen(":" + PORT))

}
