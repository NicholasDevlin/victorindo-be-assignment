package migrations

import (
	"assignment/models/businessPartner"
	"assignment/models/product"
	"assignment/models/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&businessPartner.BusinessPartner{})
	db.AutoMigrate(&product.Product{})
}
