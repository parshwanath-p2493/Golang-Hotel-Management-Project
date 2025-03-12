package utils

import "fmt"

func SendNotificationToManager(guestID string, Room_number int32, foodItems []string) {
	fmt.Printf("🔔 Notification: Guest %s has booked Room %v with food items %v\n", guestID, Room_number, foodItems)
}
