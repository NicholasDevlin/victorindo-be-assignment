package businessPartner

import "github.com/jinzhu/gorm"

func ConvertReqToDto(input BusinessPartnerReq) *BusinessPartnerDto {
	return &BusinessPartnerDto{
		UUID:        input.UUID,
		Name:        input.Name,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Type:        input.Type,
		Address:     input.Address,
	}
}

func ConvertDtoToModel(input BusinessPartnerDto) *BusinessPartner {
	return &BusinessPartner{
		Model: gorm.Model{
			ID:        input.Id,
			CreatedAt: input.CreatedAt,
			UpdatedAt: input.UpdatedAt,
		},
		UUID:        input.UUID,
		Name:        input.Name,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Type:        input.Type,
		Address:     input.Address,
	}
}

func ConvertModelToDto(input BusinessPartner) *BusinessPartnerDto {
	return &BusinessPartnerDto{
		Id:          input.ID,
		CreatedAt:   input.CreatedAt,
		UpdatedAt:   input.UpdatedAt,
		UUID:        input.UUID,
		Name:        input.Name,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Type:        input.Type,
		Address:     input.Address,
	}
}

func ConvertDtoToRes(input BusinessPartnerDto) *BusinessPartnerRes {
	return &BusinessPartnerRes{
		UUID:        input.UUID,
		Name:        input.Name,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Type:        input.Type,
		Address:     input.Address,
	}
}

var AllowedTypes = []string{"Customer", "Supplier", "Affiliate"}
func ValidType(t string) bool {
	for _, v := range AllowedTypes {
			if v == t {
					return true
			}
	}
	return false
}
