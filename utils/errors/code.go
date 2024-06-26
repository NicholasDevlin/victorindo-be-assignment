package errors

import "net/http"


func GetCodeError(err error) int {
	switch err {
	case ERR_ADDRESS_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_TOKEN:
		return http.StatusInternalServerError
	case ERR_PHONE_NUMBER_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_EMAIL_IS_EMPTY: 
		return http.StatusBadRequest
	case ERR_BCRYPT_PASSWORD:
		return http.StatusInternalServerError
	case ERR_NAME_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_PASSWORD_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_USER_NOT_FOUND:
		return http.StatusNotFound
	case ERR_WRONG_PASSWORD:
		return http.StatusConflict
	case ERR_EMAIL_IS_TAKEN:
		return http.StatusConflict
	case ERR_SAVE_DATA:
		return http.StatusInternalServerError
	case ERR_UNAUTHORIZE:
		return http.StatusUnauthorized
	case ERR_DELETE_DATA:
		return http.StatusInternalServerError
	case ERR_BUSINESS_PARTNER_NOT_FOUND:
		return http.StatusNotFound
	case ERR_BUSINESS_PARTNER_TYPE_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_INVALID_BUSINESS_PARTNER_TYPE:
		return http.StatusBadRequest
	case ERR_CODE_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_PRICE_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_PRODUCT_TYPE_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_INVALID_PRODUCT_TYPE:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
