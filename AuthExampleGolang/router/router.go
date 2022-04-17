package router

import (
	"AuthExampleGolang/handler"
	"AuthExampleGolang/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	//Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	//user
	user := api.Group("user")
	user.Post("/", handler.CreateUser)
	user.Get("/:id", handler.GetUser)

	//product
	product := api.Group("/product")
	product.Get("/", handler.GetAllProducts)
	product.Post("/", middleware.Protected(), handler.CreateProduct)
	product.Delete("/:id", middleware.Protected(), handler.DeleteProduct)
	product.Get("/:id", handler.GetProduct)

}
