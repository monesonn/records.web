// Package routes ...
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/monesonn/records.web/app/controllers"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("api")

	a.Static("/", "./public")
	// Routes for GET method:
	// route.Get("/", func(c *fiber.Ctx) error {
	// a.Static("/", "./public")
	// })
	route.Get("/genre", controllers.GetGenres)
	route.Get("/genre/:id", controllers.GetGenre)
	route.Get("/artist", controllers.GetArtists)
	route.Get("/artist/:id", controllers.GetArtist)
	route.Get("/record", controllers.GetRecords)
	route.Get("/record/:id", controllers.GetRecord)
	// route.Get("/record/artist/:artist", controllers.GetRecordArgs)
	route.Get("/record/artist/:artist", controllers.GetRecordArgs)
	route.Get("/record/artist/:artist/genre/:genre", controllers.GetRecordArgs)
	// Routes for POST method:
	route.Post("/user/sign/up", controllers.UserSignUp) // register a new user
	route.Post("/user/sign/in", controllers.UserSignIn) // auth
}
