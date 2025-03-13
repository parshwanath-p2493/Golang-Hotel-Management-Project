package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/websocket/v2"
)

var ActiveConnections = make(map[string]*websocket.Conn)

func SendNotificationToManager(managerID string, guestID string, Room_number int32, foodItems []string) {
	message := fmt.Sprintf("🔔 Notification: Guest %s has booked Room %v with food items %v\n", guestID, Room_number, foodItems)
	if connection, exists := ActiveConnections[managerID]; exists {
		err := connection.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Printf("Failed to Send Notification to manager %s:%v", managerID, err)
		} else {
			log.Printf("Manager %s  Not connected to updateds. Please Login Manager", managerID)
		}
	}
}
func WebSocketHandler(c *websocket.Conn) {
	managerID := c.Params("manager_id")
	ActiveConnections[managerID] = c
	log.Printf("Manager %s Connected and can recieve-updated the notification", managerID)
	defer func() {
		delete(ActiveConnections, managerID)
		c.Close()
	}()
}
