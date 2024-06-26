package product

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Product struct {
	gorm.Model
	UUID uuid.UUID
	Name        string      `gorm:"type:varchar(255);not null"`
	Type        ProductType `gorm:"type:enum('Item Jual', 'Item Assembly', 'Item Asset');not null"`
	Code string      `gorm:"type:varchar(100);unique;not null"`
	Price       float64     `gorm:"type:decimal(10,2);not null"`
	Description string      `gorm:"type:text"`
	Stock       int         `gorm:"default:0"`
}
