package config

import (
	"first_golang_project/models"
)

// Users/ Customers
func GetUsers(limit int) ([]models.Customer, error) {
	var users []models.Customer
	result := DB.Limit(limit).Find(&users)
	return users, result.Error
}

// Suppliers
func GetSupplier(limit int) ([]models.GetSupplier, error) {
	var supplier []models.GetSupplier
	result := DB.Limit(limit).Find(&supplier)
	return supplier, result.Error
}
