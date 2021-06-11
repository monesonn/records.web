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
	// route.Get("/clients", middleware.JWTProtected(), controllers.GetClients)

	route.Get("/user/uuid/:uuid", controllers.GetClientByUUID)
	// route.Get("/user/uuid/:uuid", middleware.JWTProtected(), controllers.GetClientByUUID)
	route.Get("/me/email/:email", middleware.JWTProtected(), controllers.GetClientByEmail)
	route.Get("/user/:username", controllers.GetUser)

	// Routes for POST method:
	route.Post("/genre", middleware.JWTProtected(), controllers.CreateGenre)
	route.Post("/artist", middleware.JWTProtected(), controllers.CreateArtist)
	route.Post("/record", middleware.JWTProtected(), controllers.CreateRecord)

	route.Post("/sign/out", middleware.JWTProtected(), controllers.UserSignOut)
	route.Post("/token/renew", middleware.JWTProtected(), controllers.RenewTokens)

	route.Post("/order", middleware.JWTProtected(), controllers.CreateOrder)
	route.Get("/order/:uuid", middleware.JWTProtected(), controllers.GetOrderByUUID)
	route.Post("/order/list", middleware.JWTProtected(), controllers.CreateOrderList)
	route.Post("/review", middleware.JWTProtected(), controllers.CreateReview)

	// Routes for PUT method:
	// route.Put("/genre", controllers.UpdateGenre)

	// Routes for DELETE method:
	// route.Delete("/genre", controllers.DeleteGenre)
}
