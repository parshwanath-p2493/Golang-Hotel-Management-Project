package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/controllers"
)

func RoomsRoutes(c *fiber.App) {
	rooms := c.Group("/rooms")
	{
		rooms.Get("/all", controllers.GetAllRooms()) //accessable for all

		//acceseble for admin and managers

	}
}
