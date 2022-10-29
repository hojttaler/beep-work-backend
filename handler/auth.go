package handler

import (
	"beep-work-backend/database"
	"beep-work-backend/models"
	"beep-work-backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

func RegisterUser(ctx *fiber.Ctx) error {
	register := &models.Register{}

	if err := ctx.BodyParser(register); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}

	if register.UserRole != "employer" && register.UserRole != "employee" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ok":      false,
			"message": "invalid role",
		})
	}

	db, err := database.OpenConnection()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}

	user := &models.User{
		ID:        uuid.NewString(),
		Email:     register.Email,
		Username:  register.Username,
		Role:      register.UserRole,
		Status:    0,
		Password:  register.Password, // TODO: Hash Password
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.CreateUser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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

func LoginUser(ctx *fiber.Ctx) error {
	login := &models.Login{}

	if err := ctx.BodyParser(login); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}

	db, err := database.OpenConnection()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}

	user, err := db.GetUserByEmail(login.Email)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"ok":      false,
			"message": "user not found",
		})
	}

	if user.Password != login.Password { // TODO: Hash Password
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"ok":      false,
			"message": "user not found",
		})
	}

	session := models.Session{
		ID:        uuid.NewString(),
		UserID:    user.ID,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	err = db.CreateSession(&session)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}

	tokens, err := utils.GenerateTokens(user.ID, session.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"ok": true,
		"result": fiber.Map{
			"access":  tokens.Access,
			"refresh": tokens.Refresh,
		},
	})
}
