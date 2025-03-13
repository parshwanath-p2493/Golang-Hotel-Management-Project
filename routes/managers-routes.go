package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/controllers"
	"github.com/parshwanath-p2493/Project/middleware"
)

func ManagerRoutes(c *fiber.App) {
	manager := c.Group("/manager", middleware.ManagerAuthentication)
	{
		manager.Get("/dashboard", func(c *fiber.Ctx) error { return c.JSON(fiber.Map{"message": "Manager Dashboard"}) })
		manager.Post("/signup", controllers.ManagerSignup)
		manager.Delete("/delete/:id", controllers.DeleteManager)
		manager.Post("/logout", controllers.LogOutManager)
		manager.Patch("/booking/:id", controllers.UpdateBookingStatus) // Manager approves/rejects booking

		manager.Get("/getallbookings", controllers.GetBooking)
	}
}

// func ManagerRoutesAUTH(app *fiber.App) {
// 	manager := app.Group("/manager", middleware.ManagerAuthentication)

// 	//	manager := app.Group("/manager", middleware.ManagerAuthentication("admin", "manager"))

// 	manager.Get("/dashboard", func(c *fiber.Ctx) error {
// 		return c.JSON(fiber.Map{"message": "Manager Dashboard"})
// 	})

// 	manager.Get("/getallstaff", controllers.GetAllStaff)
// 	manager.Post("/addstaff", controllers.AddStaff)
// 	//staff.Get("/get/:staff_id", controllers.GetStaff)
// 	manager.Put("/change/:staff_id", controllers.ChangeStaff)
// 	manager.Delete("/delete/:staff_id", controllers.DeleteStaff)
// }
