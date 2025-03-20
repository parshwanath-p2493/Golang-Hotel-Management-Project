package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/database"
	"github.com/parshwanath-p2493/Project/models"
	"github.com/parshwanath-p2493/Project/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Approve Booking API
func ApproveBooking(c *fiber.Ctx) error {
	managerID := c.Query("managerID")
	bookingID := c.Query("bookingID")
	var guest models.Guest
	if managerID == "" || bookingID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Missing bookingID or managerID"})
	}

	// Find and update booking status
	collection := database.OpenCollection("Bookings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bookingObjectID, err := primitive.ObjectIDFromHex(bookingID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid booking ID"})
	}

	update := bson.M{"$set": bson.M{"status": "Confirmed"}}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": bookingObjectID}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to approve booking"})
	}

	// Notify Guest
	//guestEmail := "guest@example.com" // Fetch from DB
	utils.SendEmail(guest.Email, "Booking Confirmed", "Your booking has been confirmed.", "Your booking is now confirmed.")

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Booking approved successfully"})
}

// Reject Booking API
func RejectBooking(c *fiber.Ctx) error {
	managerID := c.Query("managerID")
	bookingID := c.Query("bookingID")
	var guest models.Guest
	if managerID == "" || bookingID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Missing bookingID or managerID"})
	}

	// Find and update booking status
	collection := database.OpenCollection("Bookings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bookingObjectID, err := primitive.ObjectIDFromHex(bookingID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid booking ID"})
	}

	update := bson.M{"$set": bson.M{"status": "Rejected"}}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": bookingObjectID}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to reject booking"})
	}

	// Notify Guest
	//guestEmail := "guest@example.com" // Fetch from DB
	utils.SendEmail(guest.Email, "Booking Rejected", "Your booking request was rejected.", "Sorry, your booking was rejected due to unavailability.")
	//Better to use Capital Letter while declaring the function name
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Booking rejected successfully"})
}
