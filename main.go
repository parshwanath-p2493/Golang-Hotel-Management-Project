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
	a := os.Getenv("XYZ") //for testing purpose....
	fmt.Println(a)
	r.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("THE SERVER IS ALL SET GO ")
		return c.JSON(&fiber.Map{"message": "SERVER RUNNING in PORT : " + PORT})
	})

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

	//We need to call  all routes before starting the server else it will be error
	log.Fatal(r.Listen(":" + PORT))

}
