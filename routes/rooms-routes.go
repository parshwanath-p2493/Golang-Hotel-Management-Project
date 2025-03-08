package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/controllers"
	"github.com/parshwanath-p2493/Project/middleware"
)

func RoomsRoutes(c *fiber.App) {
	rooms := c.Group("/rooms")
	{
		rooms.Get("/all", controllers.GetAllRooms) //accessable for all
		rooms.Get("/types/:capacity/:room_type", controllers.FilterRooms)

	}
}

//acceseble for admin and managers

func RoomsRoutesAuth(c *fiber.App) {
	rooms := c.Group("/adminrooms", middleware.AdminAuthentication)
	{
		rooms.Post("/addroom", controllers.AddRooms)
		rooms.Delete("/delete/:room_number", controllers.DeleteRoom)
	}
}

func RoomsRoutesAuthManager(c *fiber.App) {
	rooms := c.Group("/managerrooms", middleware.AdminAuthentication)
	{
		rooms.Post("/addroom", controllers.AddRooms)
		rooms.Delete("/delete/:room_number", controllers.DeleteRoom)
	}
}
