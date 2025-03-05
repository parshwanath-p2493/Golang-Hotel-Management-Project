package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/controllers"
	"github.com/parshwanath-p2493/Project/middleware"
)

func AdminRoutes(c *fiber.App) {
	c.Post("/admin/signup", controllers.SignUpAdmin)
	c.Post("/admin/login", controllers.LoginAdmin)

}
func AdminRoutesAuth(app *fiber.App) {
	admin := app.Group("/admin", middleware.AdminAuthentication)

	//admin := app.Group("/admin", middleware.AdminAuthentication("admin"))

	admin.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Admin Dashboard"})
	})
	admin.Post("/signupmanager", controllers.ManagerSignup)
	admin.Post("/loginmanager", controllers.ManagerLogin)
	admin.Delete("/delete/:id", controllers.DeleteManager)
	admin.Get("/getallmanagers", controllers.GetManager)
	admin.Get("/getall", controllers.GetAllStaff)
	admin.Post("/addstaff", controllers.AddStaff)

}
