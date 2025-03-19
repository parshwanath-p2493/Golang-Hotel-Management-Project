package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/websocket/v2"
	"gopkg.in/gomail.v2"
)

// var managerConnections = make(map[string]map[*websocket.Conn]bool) //can have multiple connections (e.g., for a manager logged in on multiple devices).
var managerConnections = make(map[string]*websocket.Conn)

func SendNotificationToManager(managerID string, guestID string, Room_number int32, foodItems []string) {
	// Create a message based on the parameters passed.
	message := fmt.Sprintf("🔔 Notification: Guest %s has booked Room %v with food items %v\n", guestID, Room_number, foodItems)

	// Compose the HTML message with approval and rejection buttons
	htmlContent := fmt.Sprintf(`
		<html>
			<body>
				<p>%s</p>
				<p>Do you want to approve or reject the booking?</p>
				<a href="http://yourapp.com/approve?bookingID=1234&managerID=managerID" style="background-color: green; color: white; padding: 10px; text-decoration: none;">Approve</a>
				&nbsp;&nbsp;
				<a href="http://yourapp.com/reject?bookingID=1234&managerID=managerID" style="background-color: red; color: white; padding: 10px; text-decoration: none;">Reject</a>
			</body>
		</html>`, message)

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
	managerEmail := "thekingofmyqueenxyz143@gmail.com"

	subject := "New Booking Request - Action Required"

	er := sendEmail(managerEmail, subject, htmlContent, message)
	if er != nil {
		log.Printf("\n Failed to send msg to manager %s:%v", managerEmail, er)
	} else {
		log.Printf("\n ✅ Email sent  to manager Succesfuly :%s", managerEmail)

	}
}

func sendEmail(toEmail, subject, htmlContent, plainText string) error {
	fromEmail := "parshwanathparamagond1234@gmail.com"
	//fromPassword := "fbfy zhlt csqr djay"
	fromPassword := os.Getenv("FromPassword")

	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	message := gomail.NewMessage()
	message.SetHeader("From", fromEmail)
	message.SetHeader("To", toEmail)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", plainText)
	message.AddAlternative("text/html", htmlContent)

	dialer := gomail.NewDialer(smtpHost, smtpPort, fromEmail, fromPassword)

	// Send Email
	return dialer.DialAndSend(message)

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
