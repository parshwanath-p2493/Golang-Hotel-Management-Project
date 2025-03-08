package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/controllers"
	"github.com/parshwanath-p2493/Project/middleware"
)

func FoodRoutes(c *fiber.App) {
	food := c.Group("/adminfood", middleware.AdminAuthentication)
	{
		//accessable for admins and managers
		food.Post("/addfood", controllers.AddFood)
		food.Get("/getall", controllers.GetFood)
		food.Patch("/filter/:category_name/:food_name", controllers.ChangeFood)
		food.Delete("/delete/:category_name/:food_name", controllers.DeleteFood)
	}
}
func FoodRoutes2(c *fiber.App) {
	food := c.Group("/managerfood", middleware.ManagerAuthentication)
	{
		//accessable for admins and managers
		food.Post("/addfood", controllers.AddFood)
		food.Get("/getall", controllers.GetFood)
		food.Patch("/filter/:category_name/:food_name", controllers.ChangeFood)
		food.Delete("/delete/:category_name/:food_name", controllers.DeleteFood)
	}
}
