package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/websocket/v2"
)

//var ActiveConnections = make(map[string]*websocket.Conn)	// only has one connection.

/*
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
*/

// var managerConnections = make(map[string]map[*websocket.Conn]bool) //can have multiple connections (e.g., for a manager logged in on multiple devices).
var managerConnections = make(map[string]*websocket.Conn)

func SendNotificationToManager(managerID string, guestID string, Room_number int32, foodItems []string) {
	// Create a message based on the parameters passed.
	message := fmt.Sprintf("🔔 Notification: Guest %s has booked Room %v with food items %v\n", guestID, Room_number, foodItems)

	// Check if the manager is connected
	if connections, exists := managerConnections[managerID]; exists {
		err := connections.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Printf("Failed to Send Notification to manager %s:%v", managerID, err)
		} else {
			log.Printf("Invalid WebSocket connection for manager %s", managerID)
		}
	} else {
		log.Printf("Manager %s is not connected. Please login the manager.", managerID)
	}
}

func WebSocketHandler(c *websocket.Conn) {
	managerID := c.Params("manager_id")

	// Ensure the map exists before adding to it
	if _, exists := managerConnections[managerID]; exists {
		log.Printf("Manager %s is already connected. Only one connection per manager is allowed.", managerID)
		c.Close() // Close the existing connection
		return
	}

	// Store the connection in the map
	managerConnections[managerID] = c
	log.Printf("Manager %s Connected and can receive updates", managerID)

	// Ensure that the connection is cleaned up when the WebSocket connection is closed
	defer func() {
		delete(managerConnections, managerID)
		c.Close()
	}()
	for {
		err := c.WriteMessage(websocket.TextMessage, []byte("Notification test message"))
		if err != nil {
			log.Println("Write error:", err)
			break
		}
		time.Sleep(120 * time.Hour) // Sleep for 120 Hour before sending the next message
	}
}
