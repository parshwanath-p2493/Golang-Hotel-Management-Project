package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/controllers"
	"github.com/parshwanath-p2493/Project/middleware"
)

func ManagerRoutes(c *fiber.App) {
	manager := c.Group("/manager", middleware.AdminAuthentication)
	{
		manager.Post("/signup", controllers.ManagerSignup)
		manager.Post("/login", controllers.ManagerLogin)
		manager.Delete("/delete/:id", controllers.DeleteManager)
		// guest.Get("/verify-email/confirm", controllers.VerifyGuest())
		// guest.Get("/get/:id", middleware.Authentication(models.G_Acc), controllers.GetGuest())
		// guest.Put("/update/:id", middleware.Authentication(models.G_Acc), controllers.UpdateGuestDetails())
		// guest.Patch("/update-password", middleware.Authentication(models.G_Acc), controllers.ResetGuestPassword())
		// // * Admin
		// guest.Get("/getall", middleware.Authentication(models.A_Acc), controllers.GetAllGuest())

	}
}
func ManagerRoutesAUTH(app *fiber.App) {
	manager := app.Group("/manager", middleware.ManagerAuthentication)

	//	manager := app.Group("/manager", middleware.ManagerAuthentication("admin", "manager"))

	manager.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Manager Dashboard"})
	})

	manager.Get("/getall", controllers.GetAllStaff)
	manager.Post("/addstaff", controllers.AddStaff)
	//staff.Get("/get/:staff_id", controllers.GetStaff)
	manager.Put("/change/:staff_id", controllers.ChangeStaff)
	manager.Delete("/delete/:staff_id", controllers.DeleteStaff)
}
