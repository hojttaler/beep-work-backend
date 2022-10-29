package middleware

import (
	"beep-work-backend/config"
	"beep-work-backend/database"
	"beep-work-backend/utils"
	"context"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"time"
)

func AuthRequired() fiber.Handler {
	jwtConfig := jwtware.Config{
		SigningKey: []byte(config.GetEnv("JWT_ACCESS_SECRET")),
		ContextKey: "auth_required",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":      false,
					"message": err.Error(),
				})
			}

			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"ok":      false,
				"message": err.Error(),
			})
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			tokenMetadata, err := utils.ExtractJWT(ctx)
			if err != nil {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"ok":      false,
					"message": err.Error(),
				})
			}

			if time.Now().Unix() > tokenMetadata.Expires {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"ok":      false,
					"message": "token expired",
				})
			}

			db, err := database.OpenConnection()
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":      false,
					"message": err.Error(),
				})
			}

			sessionWithUser, err := db.GetSessionWithUser(tokenMetadata.SessionID, tokenMetadata.UserID)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":      false,
					"message": err.Error(),
				})
			}

			sessionWithUser.User.ID = sessionWithUser.UserID
			userContext := context.WithValue(ctx.UserContext(), "credentials", sessionWithUser.User)
			ctx.SetUserContext(userContext)

			return ctx.Next()
		},
	}

	return jwtware.New(jwtConfig)
}
