package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/parshwanath-p2493/Project/database"
	"github.com/parshwanath-p2493/Project/models"
)

func SignUpAdmin(c *fiber.Ctx) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.Admin
	defer cancel()
	collection:=database.OpenCollection("Admin")
	_,err:=
}
