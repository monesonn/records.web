package main

import (
	"context"
	_ "database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

const (
	hostname      = "localhost"
	host_port     = 5432
	username      = ""
	password      = "" // I should somehow to hide it...
	database_name = ""
)

func main() {
	pg_con_string := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable", host_port, hostname, username, password, database_name)

	conn, err := pgx.Connect(context.Background(), pg_con_string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// Create new Fiber instance
	app := fiber.New()

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
