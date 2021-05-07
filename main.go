package main

import (
	"log"

	"github.com/monesonn/records.web/pkg/configs"
	"github.com/monesonn/records.web/pkg/middleware"
	"github.com/monesonn/records.web/pkg/routes"
	_ "github.com/monesonn/records.web/platform/database"

	"github.com/gofiber/fiber/v2"

	_ "github.com/monesonn/records.web/docs" // load API Docs files (Swagger)

	_ "github.com/joho/godotenv/autoload"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @BasePath /api
func main() {
	config := configs.FiberConfig()
	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)
	routes.NotFoundRoute(app)
	routes.SwaggerRoute(app)

	log.Fatal(app.Listen(":3000"))
}
