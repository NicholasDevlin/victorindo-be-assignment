package errors

import "errors"

var (
	ERR_CREATE_CUSTOMER                = errors.New("Failed to save customer")
	ERR_ADDRESS_IS_EMPTY               = errors.New("Address is empty")
	ERR_CODE_IS_EMPTY               = errors.New("Code is empty")
	ERR_NAME_IS_EMPTY                  = errors.New("Name is empty")
	ERR_TOKEN                          = errors.New("Failed to create new token")
	ERR_PHONE_NUMBER_IS_EMPTY          = errors.New("Phone number is empty")
	ERR_EMAIL_IS_EMPTY                 = errors.New("Email is empty")
	ERR_PRICE_IS_EMPTY                 = errors.New("Price is empty")
	ERR_BUSINESS_PARTNER_TYPE_IS_EMPTY = errors.New("Business Partner Type is empty")
	ERR_PRODUCT_TYPE_IS_EMPTY = errors.New("Product Type is empty")
	ERR_INVALID_BUSINESS_PARTNER_TYPE                   = errors.New("Invalid Business Partner Type")
	ERR_INVALID_PRODUCT_TYPE                   = errors.New("Invalid Product Type")
	ERR_BCRYPT_PASSWORD                = errors.New("Failed to bcrypt password")
	ERR_PASSWORD_IS_EMPTY              = errors.New("Password is empty")
	ERR_USER_NOT_FOUND                 = errors.New("User not found")
	ERR_BUSINESS_PARTNER_NOT_FOUND     = errors.New("Business Partner not found")
	ERR_WRONG_PASSWORD                 = errors.New("Wrong Password")
	ERR_EMAIL_IS_TAKEN                 = errors.New("Email is taken")
	ERR_SAVE_DATA                      = errors.New("Error saving data")
	ERR_UNAUTHORIZE                    = errors.New("Unauthorized user")
	ERR_DELETE_DATA                    = errors.New("Failed to delete data")
)
