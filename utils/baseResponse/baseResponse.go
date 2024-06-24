package baseresponse

import (
	errors "assignment/utils/errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BasePaginationResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	CurrentPage  int   `json:"current" query:"current"`
	AllPages     int64 `json:"allPages" query:"allPages"`
	TotalRecords int64 `json:"total" query:"total"`
	PageSize     int   `json:"pageSize" query:"pageSize"`
}

func NewSuccessPaginationResponse(c echo.Context, data interface{}, pagination Pagination) error {
	return c.JSON(http.StatusOK, BasePaginationResponse{
		Success:    true,
		Message:    "Success",
		Data:       data,
		Pagination: pagination,
	})
}

func NewSuccessResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, BaseResponse{
		Success: true,
		Message: "Success",
		Data:    data,
	})
}

func NewErrorResponseUnauthorize(c echo.Context) error {
	return c.JSON(errors.GetCodeError(errors.ERR_UNAUTHORIZE), BaseResponse{
		Success:  false,
		Message: errors.ERR_UNAUTHORIZE.Error(),
		Data:    nil,
	})
}

func NewErrorResponse(c echo.Context, err error) error {
	return c.JSON(errors.GetCodeError(err), BaseResponse{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	})
}
