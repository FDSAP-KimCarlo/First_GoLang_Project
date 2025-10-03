package handlers

import (
	"first_golang_project/config"
	"first_golang_project/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

//var db = config.DB

func GetCustomer(c *fiber.Ctx) error {
	users, err := config.GetUsers(50)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func FindCustomer(c *fiber.Ctx) error {
	db := config.DB
	var reqBody struct {
		ID           int    `json:"id"`
		CustomerName string `json:"customerName"`
	}
	var customer models.Customer

	// Parse the request body
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Build query based on available fields
	query := db.Model(&models.Customer{})
	if reqBody.ID != 0 {
		query = query.Where("customerid = ?", reqBody.ID)
	} else if reqBody.CustomerName != "" {
		query = query.Where("customername = ?", reqBody.CustomerName)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Provide either id or customerName"})
	}

	// Execute query
	if err := query.First(&customer).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
	}

	return c.Status(fiber.StatusOK).JSON(customer)
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

func SoftDeleteCustomer(c *fiber.Ctx) error {
	db := config.DB
	var deleteCustomer models.Customer

	// Parse the ID from the request body
	if err := c.BodyParser(&deleteCustomer); err != nil {
		fmt.Print("error in parsing")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Find the record first
	var customer models.Customer
	if err := db.First(&customer, "customerid = ?", deleteCustomer.ID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
	}

	// Soft delete: updates deleted_at automatically
	if err := db.Delete(&customer).Error; err != nil {
		fmt.Print("There is some error")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete customer"})
	}

	// Return response with deleted_at timestamp
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Customer soft deleted",
		"id":        customer.ID,
		"deletedAt": customer.DeletedAt.Time,
	})
}
