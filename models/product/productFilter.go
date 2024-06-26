package product

import uuid "github.com/satori/go.uuid"

type ProductFilter struct {
	UUID        uuid.UUID   `query:"uuid"`
	Name        string      `query:"name"`
	Type        ProductType `query:"type"`
	Code        string      `query:"code"`
	Price       float64     `query:"price"`
	Description string      `query:"description"`
	Stock       int         `query:"stock"`
}
