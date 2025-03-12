package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/controllers"
	"github.com/parshwanath-p2493/Project/middleware"
)

func GuestRoutes(c *fiber.App) {
	guest := c.Group("/guestadmin", middleware.AdminAuthentication)
	{
		guest.Get("/getall", controllers.GetAllGuest)

	}
}
func GuestRoutes2(c *fiber.App) {
	guest := c.Group("/guestmanager", middleware.ManagerAuthentication)
	{
		//guest.Post("/signup", controllers.GuestSignup)
		guest.Get("/getall", controllers.GetAllGuest)

	}
}

//guest.Post("/login", controllers.GuestLogin)
//	guest.Get("/getall", middleware.Authentication(models.A_Acc), controllers.GetAllGuest())
/*
guest.Get("/verify-email/confirm", controllers.VerifyGuest())
guest.Get("/get/:id", middleware.Authentication(models.G_Acc), controllers.GetGuest())
guest.Put("/update/:id", middleware.Authentication(models.G_Acc), controllers.UpdateGuestDetails())
guest.Patch("/update-password", middleware.Authentication(models.G_Acc), controllers.ResetGuestPassword())
// * Admin
*/
