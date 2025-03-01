package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/controllers"
)

func FoodRoutes(c *fiber.App) {

	//accessable for admins and managers
	c.Post("/food/:category-name", controllers.AddFood())
	c.Get("/food/:category_name", controllers.GetFood())
	c.Patch("/food/:category_name/:item_name", controllers.ChangeFood())
	c.Delete("/food/:category_name/:item_name", controllers.DeleteFood())
}
