package businessPartner

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type BusinessPartnerDto struct {
	Id uint
	UUID uuid.UUID
	Name          string             
	Type          BusinessPartnerType
	Address       string             
	PhoneNumber   string             
	Email         string     
	CreatedAt time.Time
	UpdatedAt time.Time        
}