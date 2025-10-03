package routes

import (
	"first_golang_project/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	users := app.Group("/users")
	users.Get("/", handlers.GetUsers)
	users.Post("/create", handlers.CreateCustomer)
	users.Patch("/UpdateCustomer", handlers.UpdateCustomer)

}
