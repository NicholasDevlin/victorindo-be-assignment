package route

import (
	"assignment/feature/controller"
	"assignment/feature/repository"
	"assignment/feature/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoute(e *echo.Echo, db *gorm.DB, eJwt *echo.Group) {
	repository := repository.NewUsersRepository(db)
	service := service.NewUserService(repository)
	controller := controller.NewUserController(service)

	e.POST("/user/register", controller.RegisterUsers)
	e.POST("/user/login", controller.LoginUser)

	eJwt.GET("/user", controller.GetAllUser)
	// eJwt.GET("/user/:id", controller.GetUser)
	eJwt.PUT("/user/:id", controller.SaveUser)
	eJwt.DELETE("/user/:id", controller.DeleteUser)
}
