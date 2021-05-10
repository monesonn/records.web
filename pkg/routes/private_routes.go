package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/monesonn/records.web/app/controllers"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("api")

	// Routes for POST method:
	route.Post("/genre", controllers.CreateGenre)
	route.Post("/artist", controllers.CreateArtist)
	route.Post("/record", controllers.CreateRecord)

	// Routes for PUT method:
	// route.Put("/genre", controllers.UpdateGenre)

	// Routes for DELETE method:
	// route.Delete("/genre", controllers.DeleteGenre)
}
