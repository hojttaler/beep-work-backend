package route

import (
	"beep-work-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	router := app.Group("/api/v1")

	// Unauthenticated routes
	AuthRoutes(router)

	router.Use(middleware.AuthRequired())

	// Routes with auth
	UserRoutes(router)
}
