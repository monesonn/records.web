package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/monesonn/records.web/app/controllers"
	"github.com/monesonn/records.web/pkg/middleware"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("api")

	a.Static("/", "./assets")

	// Routes for GET method:
	route.Get("/clients", middleware.JWTProtected(), controllers.GetClients)

	// Routes for POST method:
	route.Post("/genre", middleware.JWTProtected(), controllers.CreateGenre)
	route.Post("/artist", middleware.JWTProtected(), controllers.CreateArtist)
	route.Post("/record", middleware.JWTProtected(), controllers.CreateRecord)

	route.Get("/me", middleware.JWTProtected(), controllers.GetMyProfile)
	route.Post("/user/sign/out", middleware.JWTProtected(), controllers.UserSignOut)
	route.Post("/token/renew", middleware.JWTProtected(), controllers.RenewTokens)

	a.Get("/admin", middleware.JWTProtected(), controllers.RenderAdminPanel)

	// Routes for PUT method:
	// route.Put("/genre", controllers.UpdateGenre)

	// Routes for DELETE method:
	// route.Delete("/genre", controllers.DeleteGenre)
}
