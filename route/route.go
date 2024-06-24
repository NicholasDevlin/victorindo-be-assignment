package route

import (
	"gorm.io/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	jwt "assignment/utils/middleware"
)

func Route(e *echo.Echo, db *gorm.DB) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	eJwt := e.Group("")
	eJwt.Use(jwt.JWTMiddleware())
	UserRoute(e, db, eJwt)
}
