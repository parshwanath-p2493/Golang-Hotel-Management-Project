package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/controllers"
	"github.com/parshwanath-p2493/Project/middleware"
)

func StaffRoutes(c *fiber.App) {
	staff := c.Group("/staff", middleware.AdminAuthentication)
	{
		staff.Get("/getall", controllers.GetAllStaff)
		staff.Post("/addstaff", controllers.AddStaff)
		//staff.Get("/get/:staff_id", controllers.GetStaff)
		staff.Put("/change/:staff_id", controllers.ChangeStaff)
		staff.Delete("/delete/:staff_id", controllers.DeleteStaff)
	}
}

func StaffRoutes2(c *fiber.App) {
	staff := c.Group("/staff", middleware.ManagerAuthentication)
	{
		staff.Get("/getall", controllers.GetAllStaff)
		staff.Post("/addstaff", controllers.AddStaff)
		//staff.Get("/get/:staff_id", controllers.GetStaff)
		staff.Put("/change/:staff_id", controllers.ChangeStaff)
		staff.Delete("/delete/:staff_id", controllers.DeleteStaff)
	}
}
