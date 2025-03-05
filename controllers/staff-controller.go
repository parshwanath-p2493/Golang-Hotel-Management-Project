package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/database"
	"github.com/parshwanath-p2493/Project/models"
	"github.com/parshwanath-p2493/Project/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllStaff(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := database.OpenCollection("Staff")

	managerDepartment := c.Locals("department").(string)
	//admin := c.Locals("role").(string)
	filter := bson.M{"department": managerDepartment}

	var staff []models.Staff
	result, err := collection.Find(ctx, filter)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Error in fetching data"))
	}
	defer result.Close(ctx)
	if err := result.All(ctx, &staff); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "Unable to fettch the data "))
	}
	return c.Status(http.StatusOK).JSON(utils.Response(c, staff, "The Staff details are:"))
}
func AddStaff(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var staff models.Staff
	collection := database.OpenCollection("Staff")

	managerDepartment := c.Locals("department").(string)

	if err := c.BodyParser(&staff); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, err.Error()))
	}
	if staff.Department != managerDepartment {
		log.Fatal("MANAGER OF DIFFERENT DEPARTMENT NOT ALLOWED")
		//return exit(0)
	}

	staff.ID = primitive.NewObjectID()
	staff.Staff_id = staff.ID.Hex()
	staff.Created_time = time.Now()
	result, err := collection.InsertOne(ctx, staff)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Syntax error"))
	}
	return c.Status(http.StatusOK).JSON(utils.Response(c, result, "Added Succssfully"))
}
func ChangeStaff(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	staffID := c.Params("staff_id")
	var updateStaff models.Staff
	managerDepartment := c.Locals("department").(string)

	//Staff_ID, _ := primitive.ObjectIDFromHex(staffID)

	if err := c.BodyParser(&updateStaff); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Error(c, utils.BadRequest, "Invalid ID"))
	}
	filter := bson.M{"staff_id": staffID, "department": managerDepartment}
	update := bson.M{
		"$set": bson.M{
			"first_name":   updateStaff.First_name,
			"last_name":    updateStaff.Last_name,
			"department":   updateStaff.Department,
			"age":          updateStaff.Age,
			"gender":       updateStaff.Gender,
			"salary":       updateStaff.Salary,
			"role":         updateStaff.Role,
			"updated_time": time.Now(),
		},
	}
	collection := database.OpenCollection("Staff")
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "ERROR"))
	}
	if result.MatchedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(utils.Error(c, utils.NotFound, "Staff id in incorrect or Manager is of defferent Department"))
	}
	return c.Status(http.StatusOK).JSON(utils.Response(c, result, "Updated Succssfully"))
}
func DeleteStaff(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//var staff_id models.Staff
	staffID := c.BodyParser("staff_id")
	managerDepartment := c.Locals("department").(string)

	collection := database.OpenCollection("Staff")
	filter := bson.M{"staff_id": staffID, "department": managerDepartment}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Error(c, utils.InternalServerError, "ERROR"))
	}
	if result.DeletedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(utils.Error(c, utils.NotFound, "No Data found "))

	}
	return c.Status(http.StatusOK).JSON(utils.Response(c, result, "Deleted  Succssfully"))
}
