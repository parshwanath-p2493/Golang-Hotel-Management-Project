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

func CreateBooking(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var booking models.Booking
	collection := database.OpenCollection("Bookings")
	if err := c.BodyParser(&booking); err != nil {
		log.Println("Invalid Syntax.......")
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, err.Error()))
	}
	booking.ID = primitive.NewObjectID()
	booking.BookingId = booking.ID.Hex()
	booking.Status = "Pending"
	booking.Created_time = time.Now()
	result, err := collection.InsertOne(ctx, booking)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Fill The information corectly and book one room at a time "))
	}
	utils.SendNotificationToManager(booking.Guest_id, booking.Room_number, booking.Food_Items)
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
