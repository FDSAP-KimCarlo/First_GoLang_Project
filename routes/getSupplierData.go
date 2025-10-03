package routes

import (
	"first_golang_project/handlers"

	"github.com/gofiber/fiber/v2"
)

func SupplierRoutes(app *fiber.App) {
	suppliers := app.Group("/suppliers")
	suppliers.Get("/", handlers.GetSuppliers)
}
