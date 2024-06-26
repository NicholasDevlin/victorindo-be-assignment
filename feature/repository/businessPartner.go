package repository

import (
	businesspartner "assignment/models/businessPartner"
	baseresponse "assignment/utils/baseResponse"
	"math"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IBusinessPartnerRepository interface {
	GetAllBusinessPartner(filter businesspartner.BusinessPartnerFilter, pagination *baseresponse.Pagination) ([]businesspartner.BusinessPartnerDto, error)
	GetBusinessPartner(filter businesspartner.BusinessPartnerFilter) (businesspartner.BusinessPartnerDto, error)
	SaveBusinessPartner(input *businesspartner.BusinessPartnerDto) error
	DeleteBusinessPartner(id string) error
}

type businessPartnerRepository struct {
	db *gorm.DB
}

// DeleteBusinessPartner implements IBusinessPartnerRepository.
func (u *businessPartnerRepository) DeleteBusinessPartner(id string) error {
	BusinessPartnerData := businesspartner.BusinessPartner{}

	err := u.db.Delete(&BusinessPartnerData, "uuid = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

// GetAllBusinessPartner implements IBusinessPartnerRepository.
func (u *businessPartnerRepository) GetAllBusinessPartner(filter businesspartner.BusinessPartnerFilter, pagination *baseresponse.Pagination) ([]businesspartner.BusinessPartnerDto, error) {
	var BusinessPartners []businesspartner.BusinessPartner

	query := u.db.Model(&businesspartner.BusinessPartner{})
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.UUID != uuid.Nil {
		query = query.Where("uuid = ?", filter.UUID)
	}
	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}
	if filter.SortByName {
		query = query.Order("name ASC")
	}

	query.Model(&businesspartner.BusinessPartner{}).Count(&pagination.TotalRecords)
	pagination.AllPages = int64(math.Ceil(float64(pagination.TotalRecords) / float64(pagination.PageSize)))

	err := query.Offset((pagination.CurrentPage - 1) * pagination.PageSize).Limit(pagination.PageSize).Find(&BusinessPartners).Error

	var resBusinessPartner []businesspartner.BusinessPartnerDto
	for i := 0; i < len(BusinessPartners); i++ {
		resBusinessPartner = append(resBusinessPartner, *businesspartner.ConvertModelToDto(BusinessPartners[i]))
	}

	return resBusinessPartner, err
}

// GetBusinessPartner implements IBusinessPartnerRepository.
func (u *businessPartnerRepository) GetBusinessPartner(filter businesspartner.BusinessPartnerFilter) (businesspartner.BusinessPartnerDto, error) {
	var BusinessPartnerData businesspartner.BusinessPartner
	query := u.db
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.UUID != uuid.Nil {
		query = query.Where("uuid = ?", filter.UUID)
	}
	if filter.Email != "" {
		query = query.Where("email = ?", filter.Email)
	}

	err := query.First(&BusinessPartnerData).Error
	if err != nil {
		return businesspartner.BusinessPartnerDto{}, err
	}
	return *businesspartner.ConvertModelToDto(BusinessPartnerData), nil
}

// SaveBusinessPartner implements IBusinessPartnerRepository.
func (u *businessPartnerRepository) SaveBusinessPartner(input *businesspartner.BusinessPartnerDto) error {
	data := businesspartner.ConvertDtoToModel(*input)
	err := u.db.Save(&data).Error
	input = businesspartner.ConvertModelToDto(*data) 
	return err
}

func NewBusinessPartnerRepository(db *gorm.DB) IBusinessPartnerRepository {
	return &businessPartnerRepository{db}
}
