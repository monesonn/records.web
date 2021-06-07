// Package routes ...
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/monesonn/records.web/app/controllers"
)

func PublicRoutes(a *fiber.App) {
	api := a.Group("api")

	a.Static("/", "./assets")

	a.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./www/index.html")
	})

	// Routes for GET method:
	api.Get("/genres", controllers.GetGenres)
	api.Get("/genre/:id", controllers.GetGenre)
	api.Get("/artists", controllers.GetArtists)
	api.Get("/artist/:id", controllers.GetArtist)
	api.Get("/record", controllers.GetRecordURL)
	api.Get("/records", controllers.GetRecords)
	api.Get("/record/:id", controllers.GetRecord)
	api.Get("/products", controllers.GetProducts)
	api.Get("/product", controllers.GetProductURL)
	api.Get("/review/:id", controllers.GetReview)
	// Routes for POST method:
	api.Post("/sign/up", controllers.UserSignUp) // register a new user
	api.Post("/sign/in", controllers.UserSignIn) // auth

	a.Get("/explore", controllers.RenderProducts)
	// a.Post("/explore", controllers.RenderProducts)

	api.Post("/client", controllers.CreateProfile)

	a.Get("/admin", controllers.RenderAdminPanel)
}
