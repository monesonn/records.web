package main

import (
	"os"

	"github.com/monesonn/records.web/pkg/configs"
	"github.com/monesonn/records.web/pkg/middleware"
	"github.com/monesonn/records.web/pkg/routes"
	"github.com/monesonn/records.web/pkg/utils"
	_ "github.com/monesonn/records.web/platform/database"

	"github.com/gofiber/fiber/v2"

	_ "github.com/monesonn/records.web/docs" // load API Docs files (Swagger)

	_ "github.com/joho/godotenv/autoload"
)

// @title API for Records.web
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @BasePath /api
func main() {
	config := configs.FiberConfig()
	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.SwaggerRoute(app)
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	// log.Fatal(app.Listen(os.Getenv("SERVER_PORT")))

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
