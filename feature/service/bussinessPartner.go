package service

import (
	"assignment/feature/repository"
	businesspartner "assignment/models/businessPartner"

	baseresponse "assignment/utils/baseResponse"
	"assignment/utils/errors"

	uuid "github.com/satori/go.uuid"
)

type IBusinessPartnerService interface {
	GetAllBusinessPartner(filter businesspartner.BusinessPartnerFilter, pagination *baseresponse.Pagination) ([]businesspartner.BusinessPartnerRes, error)
	SaveBusinessPartner(input businesspartner.BusinessPartnerReq) (businesspartner.BusinessPartnerRes, error)
	DeleteBusinessPartner(id uuid.UUID) error
}

type businessPartnerService struct {
	businessPartnerRepository repository.IBusinessPartnerRepository
}

// DeleteBusinessPartner implements IBusinessPartnerService.
func (u *businessPartnerService) DeleteBusinessPartner(id uuid.UUID) error {
	res, err := u.businessPartnerRepository.GetBusinessPartner(businesspartner.BusinessPartnerFilter{UUID: id})
	if err != nil && res.UUID == id {
		return errors.ERR_BUSINESS_PARTNER_NOT_FOUND
	}

	err = u.businessPartnerRepository.DeleteBusinessPartner(id.String())
	if err != nil {
		return errors.ERR_DELETE_DATA
	}
	return nil
}

// GetAllBusinessPartner implements IBusinessPartnerService.
func (u *businessPartnerService) GetAllBusinessPartner(filter businesspartner.BusinessPartnerFilter, pagination *baseresponse.Pagination) ([]businesspartner.BusinessPartnerRes, error) {
	if pagination.PageSize == 0 {
		pagination.PageSize = 10
	}
	if pagination.CurrentPage == 0 {
		pagination.CurrentPage = 1
	}
	res, err := u.businessPartnerRepository.GetAllBusinessPartner(filter, pagination)
	if err != nil {
		return nil, err
	}

	var resBusinessPartner []businesspartner.BusinessPartnerRes
	for i := 0; i < len(res); i++ {
		resBusinessPartner = append(resBusinessPartner, *businesspartner.ConvertDtoToRes(res[i]))
	}

	return resBusinessPartner, nil
}

// SaveBusinessPartner implements IBusinessPartnerService.
func (u *businessPartnerService) SaveBusinessPartner(input businesspartner.BusinessPartnerReq) (businesspartner.BusinessPartnerRes, error) {
	data := new(businesspartner.BusinessPartnerDto)
	var err error
	if input.UUID != uuid.Nil {
		*data, err = u.businessPartnerRepository.GetBusinessPartner(businesspartner.BusinessPartnerFilter{UUID: input.UUID})
		if err != nil {
			return businesspartner.BusinessPartnerRes{}, errors.ERR_BUSINESS_PARTNER_NOT_FOUND
		}
	} else {
		data.UUID = uuid.NewV4()
	}
	if input.Address != "" {
		data.Address = input.Address
	}
	if input.Type != "" {
		data.Type = input.Type
	}
	if input.PhoneNumber != "" {
		data.PhoneNumber = input.PhoneNumber
	}
	if input.Name != "" {
		data.Name = input.Name
	}
	if input.Email != "" {
		data.Email = input.Email
	}

	// validasi
	if data.Address == "" {
		return businesspartner.BusinessPartnerRes{}, errors.ERR_ADDRESS_IS_EMPTY
	}
	if data.Email == "" {
		return businesspartner.BusinessPartnerRes{}, errors.ERR_EMAIL_IS_EMPTY
	}
	if data.Name == "" {
		return businesspartner.BusinessPartnerRes{}, errors.ERR_NAME_IS_EMPTY
	}
	if data.PhoneNumber == "" {
		return businesspartner.BusinessPartnerRes{}, errors.ERR_PHONE_NUMBER_IS_EMPTY
	}
	if data.Type == "" {
		return businesspartner.BusinessPartnerRes{}, errors.ERR_BUSINESS_PARTNER_TYPE_IS_EMPTY
	}
	if !businesspartner.ValidType(data.Type.String()) {
		return businesspartner.BusinessPartnerRes{}, errors.ERR_INVALID_BUSINESS_PARTNER_TYPE
	}

	err = u.businessPartnerRepository.SaveBusinessPartner(data)
	if err != nil {
		return businesspartner.BusinessPartnerRes{}, errors.ERR_SAVE_DATA
	}
	return *businesspartner.ConvertDtoToRes(*data), nil
}

func NewBusinessPartnerService(repo repository.IBusinessPartnerRepository) IBusinessPartnerService {
	return &businessPartnerService{businessPartnerRepository: repo}
}
