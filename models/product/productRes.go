package product

import uuid "github.com/satori/go.uuid"

type ProductRes struct {
	UUID        uuid.UUID   `json:"uuid" form:"uuid"`
	Name        string      `json:"name" form:"name"`
	Type        ProductType `json:"type" form:"type"`
	Code        string      `json:"code" form:"code"`
	Price       float64     `json:"price" form:"price"`
	Description string      `json:"description" form:"description"`
	Stock       int         `json:"stock" form:"stock"`
}