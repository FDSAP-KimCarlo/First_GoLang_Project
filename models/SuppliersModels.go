package models

type GetSupplier struct {
	SupplierID      int    `gorm:"column:supplierid;primaryKey"`
	SupplierName    string `gorm:"column:suppliername"`
	SupplierAddress string `gorm:"column:address"`
}

func (GetSupplier) TableSupplier() string {
	return "suppliers"
}
