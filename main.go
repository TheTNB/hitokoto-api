package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically

	_ "github.com/TheTNB/hitokoto-api/docs" // load API Docs files (Swagger)
	"github.com/TheTNB/hitokoto-api/pkg/configs"
	"github.com/TheTNB/hitokoto-api/pkg/middleware"
	"github.com/TheTNB/hitokoto-api/pkg/routes"
	"github.com/TheTNB/hitokoto-api/pkg/utils"
)

// @title			Hitokoto
// @version		2.0
// @description	Hitokoto API version 2.0
// @contact.name	TreeNewBee
// @contact.email	haozi@loli.email
// @license.name	WTFPL
// @license.url	http://www.wtfpl.net/
// @BasePath		/
func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
