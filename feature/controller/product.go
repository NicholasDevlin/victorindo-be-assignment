package controller

import (
	"assignment/feature/service"
	"assignment/models/product"
	baseresponse "assignment/utils/baseResponse"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type productController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *productController {
	return &productController{productService}
}

func (u *productController) SaveProduct(e echo.Context) error {
	var input product.ProductReq
	e.Bind(&input)

	res, err := u.productService.SaveProduct(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessResponse(e, res)
}

func (u *productController) UpdateProduct(e echo.Context) error {
	var input product.ProductReq
	e.Bind(&input)
	uuid, err := uuid.FromString(e.Param("id"))
	input.UUID = uuid

	res, err := u.productService.SaveProduct(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessResponse(e, res)
}

func (u *productController) GetAllProduct(e echo.Context) error {
	var filter product.ProductFilter
	if err := e.Bind(&filter); err != nil {
		return err
	}
	var pagination baseresponse.Pagination
	if err := e.Bind(&pagination); err != nil {
		return err
	}

	res, err := u.productService.GetAllProduct(filter, &pagination)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessPaginationResponse(e, res, pagination)
}

func (u *productController) DeleteProduct(e echo.Context) error {
	uuid, err := uuid.FromString(e.Param("id"))
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	err = u.productService.DeleteProduct(uuid)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, product.ProductRes{})
}