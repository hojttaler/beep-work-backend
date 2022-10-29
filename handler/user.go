package handler

import (
	"beep-work-backend/database"
	"beep-work-backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetMe(ctx *fiber.Ctx) error {
	db, err := database.OpenConnection()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}

	credentials := ctx.UserContext().Value("credentials").(models.User)

	user, err := db.GetUserById(credentials.ID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}

	user.Password = ""

	return ctx.JSON(fiber.Map{
		"ok":     true,
		"result": user,
	})
}

func GetUser(ctx *fiber.Ctx) error {
	db, err := database.OpenConnection()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}

	user, err := db.GetUserById(ctx.Params("userId"))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}

	user.Password = ""

	return ctx.JSON(fiber.Map{
		"ok":     true,
		"result": user,
	})
}
