package handlers

import (
	"first_golang_project/config"
	"first_golang_project/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	users, err := config.GetUsers(50)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func CreateCustomer(c *fiber.Ctx) error {
	customer := new(models.Customer)
	c.BodyParser(customer)      // parse JSON (without ID)
	config.DB.Create(&customer) // database auto-generates customerid

	if err := c.BodyParser(customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := config.DB.Create(&customer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(customer)
}

func UpdateCustomer(c *fiber.Ctx) error {
	db := config.DB
	var customerUpdate models.Customer
	if err := c.BodyParser(&customerUpdate); err != nil {
		fmt.Print("error in parsing")
	}
	result := db.Model(&models.Customer{}).Where("customerid = ?", customerUpdate.ID).Updates(&customerUpdate)

	if result.Error != nil {
		fmt.Print("There is some error")
	}
	return c.JSON(customerUpdate)
}

/*func UpdateCustomer(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var customer models.Customer

		// Find existing customer
		if err := db.First(&customer, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "Customer not found",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// Parse request body
		var input models.Customer
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		// Update fields
		customer.Name = input.Name
		customer.City = input.City

		if err := db.Save(&customer).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(customer)
	}
}*/
