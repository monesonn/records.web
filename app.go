package main

import (
	"log"

	"github.com/monesonn/records.web/pkg/configs"
	"github.com/monesonn/records.web/pkg/routes"
	_ "github.com/monesonn/records.web/platform/database"

	"github.com/gofiber/fiber/v2"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := configs.FiberConfig()
	app := fiber.New(config)

	routes.PublicRoutes(app)
	routes.NotFoundRoute(app)
	routes.SwaggerRoute(app)

	log.Fatal(app.Listen(":3000"))
}
