package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/controllers"
)

func StaffRoutes(c *fiber.App) {
	staff := c.Group("/staff")
	{
		staff.Get("/getall", controllers.GetAllStaff())
		staff.Post("/add/:staff_id", controllers.AddStaff())
		staff.Get("/get/:staff_id", controllers.GetStaff())
		staff.Patch("/change/:staff_id", controllers.ChangeStaff())
		staff.Delete("/delete/:staff_id", controllers.DeleteStaff())
	}
}
