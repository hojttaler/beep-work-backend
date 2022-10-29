package route

import (
	"beep-work-backend/handler"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	userRouter := router.Group("/user")

	userRouter.Get("/me", handler.GetMe)
	userRouter.Get("/:userId", handler.GetUser)
}
