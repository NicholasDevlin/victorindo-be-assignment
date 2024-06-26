package controller

import (
	"assignment/feature/service"
	businesspartner "assignment/models/businessPartner"
	baseresponse "assignment/utils/baseResponse"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type businessPartnerController struct {
	businessPartnerService service.IBusinessPartnerService
}

func NewBusinessPartnerController(businessPartnerService service.IBusinessPartnerService) *businessPartnerController {
	return &businessPartnerController{businessPartnerService}
}

func (u *businessPartnerController) SaveBusinessPartner(e echo.Context) error {
	var input businesspartner.BusinessPartnerReq
	e.Bind(&input)

	res, err := u.businessPartnerService.SaveBusinessPartner(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessResponse(e, res)
}

func (u *businessPartnerController) UpdateBusinessPartner(e echo.Context) error {
	var input businesspartner.BusinessPartnerReq
	e.Bind(&input)
	uuid, err := uuid.FromString(e.Param("id"))
	input.UUID = uuid

	res, err := u.businessPartnerService.SaveBusinessPartner(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessResponse(e, res)
}

func (u *businessPartnerController) GetAllBusinessPartner(e echo.Context) error {
	var filter businesspartner.BusinessPartnerFilter
	if err := e.Bind(&filter); err != nil {
		return err
	}
	var pagination baseresponse.Pagination
	if err := e.Bind(&pagination); err != nil {
		return err
	}

	res, err := u.businessPartnerService.GetAllBusinessPartner(filter, &pagination)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessPaginationResponse(e, res, pagination)
}

func (u *businessPartnerController) DeleteBusinessPartner(e echo.Context) error {
	uuid, err := uuid.FromString(e.Param("id"))
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	err = u.businessPartnerService.DeleteBusinessPartner(uuid)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, businesspartner.BusinessPartnerRes{})
}