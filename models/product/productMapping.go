package product

import "github.com/jinzhu/gorm"

func ConvertReqToDto(input ProductReq) *ProductDto {
	return &ProductDto{
		UUID:        input.UUID,
		Name:        input.Name,
		Type:        input.Type,
		Code:        input.Code,
		Price:       input.Price,
		Description: input.Description,
		Stock:       input.Stock,
	}
}

func ConvertDtoToModel(input ProductDto) *Product {
	return &Product{
		Model: gorm.Model{
			ID:        input.Id,
			CreatedAt: input.CreatedAt,
			UpdatedAt: input.UpdatedAt,
		},
		UUID:        input.UUID,
		Name:        input.Name,
		Type:        input.Type,
		Code:        input.Code,
		Price:       input.Price,
		Description: input.Description,
		Stock:       input.Stock,
	}
}

func ConvertModelToDto(input Product) *ProductDto {
	return &ProductDto{
		Id:          input.ID,
		CreatedAt:   input.CreatedAt,
		UpdatedAt:   input.UpdatedAt,
		UUID:        input.UUID,
		Name:        input.Name,
		Type:        input.Type,
		Code:        input.Code,
		Price:       input.Price,
		Description: input.Description,
		Stock:       input.Stock,
	}
}

func ConvertDtoToRes(input ProductDto) *ProductRes {
	return &ProductRes{
		UUID:        input.UUID,
		Name:        input.Name,
		Type:        input.Type,
		Price:       input.Price,
		Description: input.Description,
		Stock:       input.Stock,
	}
}
