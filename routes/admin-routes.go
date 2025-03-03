package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/controllers"
)

func AdminRoutes(c *fiber.App) {
	c.Post("/admin/signup", controllers.SignUpAdmin)
	c.Get("/admin/login/:admin_id", controllers.LoginAdmin)

}
