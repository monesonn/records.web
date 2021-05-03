package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/monesonn/records.web/app/controllers"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("")

	route.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})
	route.Get("/genres", controllers.GetGenres)
	route.Get("/artists", controllers.GetArtists)
}
