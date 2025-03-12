package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/controllers"
	"github.com/parshwanath-p2493/Project/middleware"
)

func CreateBooking(c *fiber.App) {
	guest := c.Group("/guest")
	{
		guest.Post("/bookhotel", controllers.CreateBooking, middleware.GuestAuth)
		c.Get("/getallfood", controllers.GetFood)
		c.Post("/signup", controllers.GuestSignup)
	}
}
