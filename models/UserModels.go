package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID          int            `gorm:"column:customerid;primaryKey;autoIncrement" json:"id"`
	Name        string         `gorm:"column:customername" json:"name"`
	ContactName string         `gorm:"column:contactname" json:"contactName"`
	Address     string         `gorm:"column:address" json:"address"`
	City        string         `gorm:"column:city" json:"city"`
	PostalCode  string         `gorm:"column:postalcode" json:"postalCode"`
	Country     string         `gorm:"column:country" json:"country"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deletedAt,omitempty"`
}

func (Customer) TableCustomer() string {
	return "customers"
}
