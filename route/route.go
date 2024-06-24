package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	jwt "assignment/utils/middleware"
)

func Route(e *echo.Echo, db *gorm.DB) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	eJwt := e.Group("")
	eJwt.Use(jwt.JWTMiddleware())
}
