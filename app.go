package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/monesonn/records.web/pkg/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	config := configs.FiberConfig()

	// Create new Fiber instance
	app := fiber.New(config)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("I made a ☕ for you!")
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	// Start server on http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Oopsie!")
}
