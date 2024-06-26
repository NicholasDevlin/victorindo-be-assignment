package businessPartner

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type BusinessPartner struct {
	gorm.Model
	UUID uuid.UUID
	Name          string             `gorm:"type:varchar(255);not null"`
	Type          BusinessPartnerType `gorm:"type:enum('Customer', 'Supplier', 'Affiliate');not null"`
	Address       string             `gorm:"type:varchar(255)"`
	PhoneNumber   string             `gorm:"type:varchar(50)"`
	Email         string             `gorm:"type:varchar(255)"`
}