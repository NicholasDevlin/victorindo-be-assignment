package product

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type ProductDto struct {
	Id uint
	UUID uuid.UUID
	Name        string     
	Type        ProductType
	Code string     
	Price       float64    
	Description string     
	Stock       int        
	CreatedAt time.Time
	UpdatedAt time.Time 
}