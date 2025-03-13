package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	//	"github.com/parshwanath-p2493/Project/utils"
)

//var validato = validator.Validate()

func Validation(c *fiber.Ctx, model interface{}) (error, int16) {
	var validate = validator.New()
	var count int16
	var errormessage string
	err := validate.Struct(model)
	if err != nil {
		var ErrorMSg validator.ValidationErrors
		errors.As(err, &ErrorMSg)
		for _, validationError := range ErrorMSg {
			errormessage += fmt.Sprintf("Field %s is must required %s \n ", validationError.Field(), validationError.Tag())
			log.Println("Validation Error:", ErrorMSg)
			count++
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": map[string]string{
					"field":   validationError.Field(),
					"message": validationError.Tag(),
				},
			}), count
		}
	}
	log.Println("\n", count)
	return nil, count
}
