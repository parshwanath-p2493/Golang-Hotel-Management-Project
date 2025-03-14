package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/database"
	"github.com/parshwanath-p2493/Project/models"
	"github.com/parshwanath-p2493/Project/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var validate = Validator.New()

func CreateBooking(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var booking models.Booking
	var room models.Rooms
	collection := database.OpenCollection("Bookings")
	if err := c.BodyParser(&booking); err != nil {
		log.Println("Invalid Syntax.......")
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, err.Error()))
	}
	//assigning the details...
	booking.ID = primitive.NewObjectID()
	booking.BookingId = booking.ID.Hex()
	booking.Status = "Pending"
	booking.Created_time = time.Now()
	booking.Updated_time = booking.ID.Timestamp()

	if room.Availability_status == string(models.Room_Occupied) {
		utils.Error(c, utils.Conflict, "Room already occupied by guest.")
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Room already occupied by guest"))

	}
	if err, count := utils.Validation(c, booking); count > 1 {
		log.Fatal("Enter all the required Fields", err)
		return err
	}
	result, err := collection.InsertOne(ctx, booking)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Fill The information corectly and book one room at a time "))
	}
	utils.SendNotificationToManager("67cca92b5532aeb8476e2334", booking.Guest_id, booking.Room_number, booking.Food_Items)
	return c.Status(http.StatusOK).JSON(utils.Response(c, result, "Booking Successfull"))
}
func GetBooking(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var booking []models.Booking
	collection := database.OpenCollection("Bookings")
	SortByTime := c.Query("sortbytime")
	var SetOrder int
	if SortByTime == "asc" {
		SetOrder = 1
	} else if SortByTime == "des" {
		SetOrder = -1
	} else {
		SetOrder = 1
	}
	filter := bson.M{}
	sort := bson.M{"bookingdate": SetOrder}
	result, err := collection.Find(ctx, filter, options.Find().SetSort(sort))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Error in finding bookings"))
	}
	defer result.Close(ctx)
	if err := result.All(ctx, &booking); err != nil {
		fmt.Println(err.Error())
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Error"))
	}
	return c.Status(http.StatusOK).JSON(utils.Response(c, booking, "Booking data fetched Successfull"))

}
func UpdateBookingStatus(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	bookingID := c.Params("id")
	status := c.Query("status")
	if status != "approved" && status != "rejected" {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Invalid Request"))
	}
	collection := database.OpenCollection("Bookings")
	objectID, err := primitive.ObjectIDFromHex(bookingID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Invalid Booking ID"))
	}
	update := bson.M{
		"$set": bson.M{
			"status":       status,
			"updated_time": time.Now(),
		},
	}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Failed to Update booking"))
	}
	return c.Status(http.StatusOK).JSON(utils.Response(c, "Booking STATUS", "Booking"+status+"Successfuly"))
}
