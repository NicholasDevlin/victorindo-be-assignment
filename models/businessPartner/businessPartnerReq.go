package businessPartner

import uuid "github.com/satori/go.uuid"

type BusinessPartnerReq struct {
	UUID        uuid.UUID           `json:"uuid" form:"uuid"`
	Name        string              `json:"name" form:"name"`
	Type        BusinessPartnerType `json:"type" form:"type"`
	Address     string              `json:"address" form:"address"`
	PhoneNumber string              `json:"phoneNumber" form:"phoneNumber"`
	Email       string              `json:"email" form:"email"`
}
