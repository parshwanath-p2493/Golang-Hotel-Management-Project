package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/controllers"
	"github.com/parshwanath-p2493/Project/middleware"
)

func CreateBooking(c *fiber.App) {
	guest := c.Group("/guest")
	{
		guest.Post("/bookhotel", middleware.GuestAuth, controllers.CreateBooking) //first call the auth function then call the controller function
		guest.Post("/signup", controllers.GuestSignup)
		guest.Get("/getallfood", controllers.GetFood)
		guest.Post("/logout", middleware.GuestAuth, controllers.LogOutGuest)
		guest.Post("/login", middleware.GuestAuth, controllers.GuestLogin)

	}
}
