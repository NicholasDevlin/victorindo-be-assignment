package service

import (
	"assignment/feature/repository"
	"assignment/models/product"

	baseresponse "assignment/utils/baseResponse"
	"assignment/utils/errors"

	uuid "github.com/satori/go.uuid"
)

type IProductService interface {
	GetAllProduct(filter product.ProductFilter, pagination *baseresponse.Pagination) ([]product.ProductRes, error)
	SaveProduct(input product.ProductReq) (product.ProductRes, error)
	DeleteProduct(id uuid.UUID) error
}

type productService struct {
	productRepository repository.IProductRepository
}

// DeleteProduct implements IProductService.
func (u *productService) DeleteProduct(id uuid.UUID) error {
	res, err := u.productRepository.GetProduct(product.ProductFilter{UUID: id})
	if err != nil && res.UUID == id {
		return errors.ERR_BUSINESS_PARTNER_NOT_FOUND
	}

	err = u.productRepository.DeleteProduct(id.String())
	if err != nil {
		return errors.ERR_DELETE_DATA
	}
	return nil
}

// GetAllProduct implements IProductService.
func (u *productService) GetAllProduct(filter product.ProductFilter, pagination *baseresponse.Pagination) ([]product.ProductRes, error) {
	if pagination.PageSize == 0 {
		pagination.PageSize = 10
	}
	if pagination.CurrentPage == 0 {
		pagination.CurrentPage = 1
	}
	res, err := u.productRepository.GetAllProduct(filter, pagination)
	if err != nil {
		return nil, err
	}

	var resBusinessPartner []product.ProductRes
	for i := 0; i < len(res); i++ {
		resBusinessPartner = append(resBusinessPartner, *product.ConvertDtoToRes(res[i]))
	}

	return resBusinessPartner, nil
}

// SaveProduct implements IProductService.
func (u *productService) SaveProduct(input product.ProductReq) (product.ProductRes, error) {
	data := new(product.ProductDto)
	var err error
	if input.UUID != uuid.Nil {
		*data, err = u.productRepository.GetProduct(product.ProductFilter{UUID: input.UUID})
		if err != nil {
			return product.ProductRes{}, errors.ERR_BUSINESS_PARTNER_NOT_FOUND
		}
	} else {
		data.UUID = uuid.NewV4()
	}
	if input.Price != 0 {
		data.Price = input.Price
	}
	if input.Stock != 0 {
		data.Stock = input.Stock
	}
	if input.Type != "" {
		data.Type = input.Type
	}
	if input.Code != "" {
		data.Code = input.Code
	}
	if input.Name != "" {
		data.Name = input.Name
	}

	// validasi
	if data.Code == "" {
		return product.ProductRes{}, errors.ERR_CODE_IS_EMPTY
	}
	if data.Price == 0 {
		return product.ProductRes{}, errors.ERR_PRICE_IS_EMPTY
	}
	if data.Name == "" {
		return product.ProductRes{}, errors.ERR_NAME_IS_EMPTY
	}
	if data.Type == "" {
		return product.ProductRes{}, errors.ERR_PRODUCT_TYPE_IS_EMPTY
	}
	if !product.ValidType(data.Type.String()) {
		return product.ProductRes{}, errors.ERR_INVALID_PRODUCT_TYPE
	}

	err = u.productRepository.SaveProduct(data)
	if err != nil {
		return product.ProductRes{}, errors.ERR_SAVE_DATA
	}
	return *product.ConvertDtoToRes(*data), nil
}

func NewProductService(repo repository.IProductRepository) IProductService {
	return &productService{productRepository: repo}
}
