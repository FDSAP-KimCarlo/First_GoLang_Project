package routes

import (
	"first_golang_project/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	users := app.Group("/users")
	users.Get("/", handlers.GetCustomer)
	users.Post("/create", handlers.CreateCustomer)
	users.Patch("/UpdateCustomer", handlers.UpdateCustomer)
	users.Patch("/delete", handlers.SoftDeleteCustomer)
	users.Get("/selectCustomer/:d", handlers.FindCustomer)

}
