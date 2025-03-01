package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/controllers"
)

func GuestRoutes(c *fiber.App) {
	guest := c.Group("/guest")
	{
		guest.Post("/signup", controllers.GuestSignup())
		guest.Post("/login", controllers.GuestLogin())
		/*
			guest.Get("/verify-email/confirm", controllers.VerifyGuest())
			guest.Get("/get/:id", middleware.Authentication(models.G_Acc), controllers.GetGuest())
			guest.Put("/update/:id", middleware.Authentication(models.G_Acc), controllers.UpdateGuestDetails())
			guest.Patch("/update-password", middleware.Authentication(models.G_Acc), controllers.ResetGuestPassword())
			// * Admin
			guest.Get("/getall", middleware.Authentication(models.A_Acc), controllers.GetAllGuest())
		*/
	}
}
