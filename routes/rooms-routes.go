package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/controllers"
)

func RoomsRoutes(c *fiber.App) {
	rooms := c.Group("/rooms")
	{
		rooms.Get("/all", controllers.GetAllRooms) //accessable for all
		rooms.Get("/types", controllers.FilterRooms)
		//acceseble for admin and managers
		rooms.Post("/addroom", controllers.AddRooms)
	}
}
