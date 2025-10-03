package handlers

import (
	"first_golang_project/config"

	"github.com/gofiber/fiber/v2"
)

func GetSuppliers(c *fiber.Ctx) error {
	suppliers, err := config.GetSupplier(10)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(suppliers)
}
