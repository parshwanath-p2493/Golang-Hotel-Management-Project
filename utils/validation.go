package utils

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	//	"github.com/parshwanath-p2493/Project/utils"
)

var validate = validator.New()

func Validation(c *fiber.Ctx, model interface{}) error {
	err := validate.Struct(model)
	if err != nil {
		var ErrorMSg string
		for _, e := range err.(validator.ValidationErrors) {
			ErrorMSg += fmt.Sprintf("Field %s is must required ", e.Field())
		}
		log.Println("Validation Error:", ErrorMSg)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": ErrorMSg,
			"status":  "error",
		})
	}
	return nil
}
