package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/database"
	"github.com/parshwanath-p2493/Project/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// Approve Booking API
func ApproveBooking(c *fiber.Ctx) error {
	managerID := c.Query("managerID")
	bookingID := c.Query("bookingID")
	//var guest models.Guest
	if managerID == "" || bookingID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Missing bookingID or managerID"})
	}
	log.Println("booking id ", bookingID)
	log.Println("manager id ", managerID)
	// Find and update booking status
	collection := database.OpenCollection("Bookings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// BookingId, err := primitive.ObjectIDFromHex(bookingID)
	// if err != nil {
	// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid booking ID"})
	// }

	update := bson.M{"$set": bson.M{"status": "Confirmed"}}
	_, err := collection.UpdateOne(ctx, bson.M{"bookingid": bookingID}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to approve booking"})
	}
	log.Println("Booking done successfuly")
	// Notify Guest
	guestEmail := "rahulroshu2003@gmail.com" // Fetch from DB
	utils.SendEmail(guestEmail, "Booking Confirmed", "Your booking has been confirmed.", "Your booking is now confirmed.")

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Booking approved successfully"})
}

// Reject Booking API
func RejectBooking(c *fiber.Ctx) error {
	managerID := c.Query("managerID")
	bookingID := c.Query("bookingID")
	//var guest models.Guest
	if managerID == "" || bookingID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Missing bookingID or managerID"})
	}

	// Find and update booking status
	collection := database.OpenCollection("Bookings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{"$set": bson.M{"status": "Rejected"}}
	_, err := collection.UpdateOne(ctx, bson.M{"bookingid": bookingID}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to reject booking"})
	}

	//Notify Guest
	guestEmail := "rahulroshu2003@gmail.com" // Fetch from DB
	utils.SendEmail(guestEmail, "Booking Rejected", "Your booking request was rejected.", "Sorry, your booking was rejected due to unavailability.")
	//Always  use Capital Letter for function name while declaring them
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Booking rejected successfully"})
}
