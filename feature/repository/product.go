package repository

import (
	"assignment/models/product"
	baseresponse "assignment/utils/baseResponse"
	"math"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IProductRepository interface {
	GetAllProduct(filter product.ProductFilter, pagination *baseresponse.Pagination) ([]product.ProductDto, error)
	GetProduct(filter product.ProductFilter) (product.ProductDto, error)
	SaveProduct(input *product.ProductDto) error
	DeleteProduct(id string) error
}

type productRepository struct {
	db *gorm.DB
}

// DeleteProduct implements IProductRepository.
func (u *productRepository) DeleteProduct(id string) error {
	ProductData := product.Product{}

	err := u.db.Delete(&ProductData, "uuid = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

// GetAllProduct implements IProductRepository.
func (u *productRepository) GetAllProduct(filter product.ProductFilter, pagination *baseresponse.Pagination) ([]product.ProductDto, error) {
	var products []product.Product

	query := u.db.Model(&product.Product{})
	if filter.UUID != uuid.Nil {
		query = query.Where("uuid = ?", filter.UUID)
	}
	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}
	if filter.SortByName {
		query = query.Order("name ASC")
	}

	query.Model(&product.Product{}).Count(&pagination.TotalRecords)
	pagination.AllPages = int64(math.Ceil(float64(pagination.TotalRecords) / float64(pagination.PageSize)))

	err := query.Offset((pagination.CurrentPage - 1) * pagination.PageSize).Limit(pagination.PageSize).Find(&products).Error

	var resProduct []product.ProductDto
	for i := 0; i < len(products); i++ {
		resProduct = append(resProduct, *product.ConvertModelToDto(products[i]))
	}

	return resProduct, err
}

// GetProduct implements IProductRepository.
func (u *productRepository) GetProduct(filter product.ProductFilter) (product.ProductDto, error) {
	var ProductData product.Product
	query := u.db
	if filter.UUID != uuid.Nil {
		query = query.Where("uuid = ?", filter.UUID)
	}
	if filter.Name != "" {
		query = query.Where("email = ?", filter.Name)
	}

	err := query.First(&ProductData).Error
	if err != nil {
		return product.ProductDto{}, err
	}
	return *product.ConvertModelToDto(ProductData), nil
}

// SaveProduct implements IProductRepository.
func (u *productRepository) SaveProduct(input *product.ProductDto) error {
	data := product.ConvertDtoToModel(*input)
	err := u.db.Save(&data).Error
	input = product.ConvertModelToDto(*data) 
	return err
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &productRepository{db}
}
