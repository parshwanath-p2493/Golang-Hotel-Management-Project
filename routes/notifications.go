package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/parshwanath-p2493/Project/utils"
)

func NotificationRoutes(app *fiber.App) {
	app.Get("/ws/manager/:manager_id", websocket.New(utils.WebSocketHandler))
}
