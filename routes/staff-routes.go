package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/controllers"
	"github.com/parshwanath-p2493/Project/middleware"
)

func StaffRoutes(c *fiber.App) {
	staff := c.Group("/adminstaff", middleware.AdminAuthentication)
	{
		staff.Get("/getall", controllers.GetAllStaff)
		// staff.Post("/addstaff", controllers.AddStaff)
		// //staff.Get("/get/:staff_id", controllers.GetStaff)
		// staff.Put("/change/:staff_id", controllers.ChangeStaff)
		// staff.Delete("/delete/:staff_id", controllers.DeleteStaff)
	}
}

func StaffRoutes2(c *fiber.App) {
	staff := c.Group("/managerstaff", middleware.ManagerAuthentication)
	{
		staff.Get("/getall", controllers.GetAllStaffDept)
		staff.Post("/addstaff", controllers.AddStaff)
		//staff.Get("/get/:staff_id", controllers.GetStaff)
		staff.Put("/change/:staff_id", controllers.ChangeStaff)
		staff.Delete("/delete/:staff_id", controllers.DeleteStaff)
	}
}

/*

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/controllers"
	"github.com/parshwanath-p2493/Project/middleware"
)

func StaffRoutes(c *fiber.App) {
	// Create a new route group for both admin and manager
	staff := c.Group("/staff")

	// Apply manager middleware for manager-specific staff routes
	staff.Use(middleware.ManagerAuthentication) // Middleware applied to all below routes in this group
	{
		staff.Get("/getall", controllers.GetAllStaffDept) // Manager route
		staff.Post("/addstaff", controllers.AddStaff)     // Manager route
		staff.Put("/change/:staff_id", controllers.ChangeStaff)
		staff.Delete("/delete/:staff_id", controllers.DeleteStaff)
	}

	// Apply admin middleware for admin-specific staff routes
	adminStaff := c.Group("/adminstaff", middleware.AdminAuthentication) // Admin route
	{
		adminStaff.Get("/getall", controllers.GetAllStaff)
		adminStaff.Post("/addstaff", controllers.AddStaff)
		adminStaff.Put("/change/:staff_id", controllers.ChangeStaff)
		adminStaff.Delete("/delete/:staff_id", controllers.DeleteStaff)
	}
}


*/
