package route

import (
	"beep-work-backend/handler"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router) {
	authRouter := router.Group("/auth")

	authRouter.Post("/register", handler.RegisterUser)
	authRouter.Post("/login", handler.LoginUser)
}
