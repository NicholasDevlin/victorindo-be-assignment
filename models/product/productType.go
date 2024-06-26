package product

import (
    "database/sql/driver"
    "errors"
)

type ProductType string

const (
    ItemJual     ProductType = "Item Jual"
    ItemAssembly ProductType = "Item Assembly"
    ItemAsset    ProductType = "Item Asset"
)

func (p *ProductType) Scan(value interface{}) error {
    if value == nil {
        *p = ""
        return nil
    }

    switch v := value.(type) {
    case []byte:
        *p = ProductType(v)
    case string:
        *p = ProductType(v)
    default:
        return errors.New("invalid type for ProductType")
    }
    return nil
}

func (p ProductType) Value() (driver.Value, error) {
    return string(p), nil
}

func (b ProductType) String() string {
    return string(b)
}