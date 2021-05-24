package routes

import "github.com/gofiber/fiber/v2"

// NotFoundRoute func for describe 404 Error route.
func NotFoundRoute(a *fiber.App) {
	a.Use(func(c *fiber.Ctx) error {
		// a.Static("404", "./public/routes/404")
		c.SendFile("./public/routes/404/index.html")
		return c.SendStatus(404)
	})
}
