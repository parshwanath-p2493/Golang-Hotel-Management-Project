package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/controllers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Booking Routes
	api.Get("/approve", controllers.ApproveBooking)
	api.Get("/reject", controllers.RejectBooking)
}
