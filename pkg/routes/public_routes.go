package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/TheTNB/hitokoto-api/app/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	a.Get("/", controllers.Get)
}
