package businessPartner

import uuid "github.com/satori/go.uuid"

type BusinessPartnerFilter struct {
	Id          uint
	UUID        uuid.UUID           `query:"uuid"`
	Name        string              `query:"name"`
	Type        BusinessPartnerType `query:"type"`
	Address     string              `query:"address"`
	PhoneNumber string              `query:"phoneNumber"`
	Email       string              `query:"email"`
	SortByName  bool                `query:"sortByName"`
}
