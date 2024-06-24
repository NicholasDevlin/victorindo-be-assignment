package controller

import (
	"assignment/feature/service"
	"assignment/models/user"
	baseresponse "assignment/utils/baseResponse"
	"assignment/utils/errors"
	"assignment/utils/middleware"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type userController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *userController {
	return &userController{userService}
}

func (u *userController) RegisterUsers(e echo.Context) error {
	var input user.UserReq
	e.Bind(&input)

	res, err := u.userService.RegisterUser(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	token, err := middleware.CreateToken(res.UUID, res.Name)
	if err != nil {
		return baseresponse.NewErrorResponse(e, errors.ERR_TOKEN)
	}
	res.Token = token

	return baseresponse.NewSuccessResponse(e, res)
}

func (u *userController) LoginUser(e echo.Context) error {
	var input user.UserReq
	e.Bind(&input)

	res, err := u.userService.LoginUser(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	token, err := middleware.CreateToken(res.UUID, res.Name)
	if err != nil {
		return baseresponse.NewErrorResponse(e, errors.ERR_TOKEN)
	}
	res.Token = token

	return baseresponse.NewSuccessResponse(e, res)
}

func (u *userController) GetAllUser(e echo.Context) error {
	var filter user.UserFilter
	if err := e.Bind(&filter); err != nil {
		return err
	}
	var pagination baseresponse.Pagination
	if err := e.Bind(&pagination); err != nil {
		return err
	}

	res, err := u.userService.GetAllUser(filter, &pagination)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessPaginationResponse(e, res, pagination)
}

func (u *userController) SaveUser(e echo.Context) error {
	userId, err := middleware.ExtractToken(e)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	var input user.UserReq
	e.Bind(&input)
	uuid, err := uuid.FromString(e.Param("id"))
	if userId != uuid {
		return baseresponse.NewErrorResponseUnauthorize(e)
	}
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	input.UUID = uuid
	res, err := u.userService.SaveUser(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, res)
}

func (u *userController) DeleteUser(e echo.Context) error {
	userId, err := middleware.ExtractToken(e)

	uuid, err := uuid.FromString(e.Param("id"))
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	if userId != uuid {
		return baseresponse.NewErrorResponseUnauthorize(e)
	}

	err = u.userService.DeleteUser(uuid)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, user.UserRes{})
}