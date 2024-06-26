package businessPartner

import (
    "database/sql/driver"
)

type BusinessPartnerType string
const (
    Customer  BusinessPartnerType = "Customer"
    Supplier  BusinessPartnerType = "Supplier"
    Affiliate BusinessPartnerType = "Affiliate"
)

func (b *BusinessPartnerType) Scan(value interface{}) error {
    *b = BusinessPartnerType(value.([]byte))
    return nil
}

func (b BusinessPartnerType) Value() (driver.Value, error) {
    return string(b), nil
}

func (b BusinessPartnerType) String() string {
    return string(b)
}