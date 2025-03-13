package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/parshwanath-p2493/Project/utils"
)

//app.Get("/ws/manager/:manager_id", websocket.New(utils.WebSocketHandler))

func NotificationRoutes(r *fiber.App) {

	// Middleware to check WebSocket upgrade request
	r.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// WebSocket route to handle manager connections
	r.Get("/ws/manager/:manager_id", websocket.New(utils.WebSocketHandler))
}
