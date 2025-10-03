package models

type Customer struct {
	ID          int    `gorm:"column:customerid;primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"column:customername" json:"name"`
	ContactName string `gorm:"column:contactname" json:"contactName"`
	Address     string `gorm:"column:address" json:"address"`
	City        string `gorm:"column:city" json:"city"`
	PostalCode  string `gorm:"column:postalcode" json:"postalCode"`
	Country     string `gorm:"column:country" json:"country"`
}

func (Customer) TableName() string {
	return "customers"
}
