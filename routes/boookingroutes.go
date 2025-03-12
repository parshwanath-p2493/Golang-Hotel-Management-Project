package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/controllers"
)

func CreateBooking(c *fiber.App) {
	c.Post("/bookhotel", controllers.CreateBooking)
}
