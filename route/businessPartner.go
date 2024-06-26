package route

import (
	"assignment/feature/controller"
	"assignment/feature/repository"
	"assignment/feature/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BusinessPartnerRoute(e *echo.Echo, db *gorm.DB, eJwt *echo.Group) {
	repository := repository.NewBusinessPartnerRepository(db)
	service := service.NewBusinessPartnerService(repository)
	controller := controller.NewBusinessPartnerController(service)

	eJwt.POST("/business-partner", controller.SaveBusinessPartner)

	eJwt.GET("/business-partner", controller.GetAllBusinessPartner)
	// eJwt.GET("/business-partner/:id", controller.GetBusinessPartner)
	eJwt.PUT("/business-partner/:id", controller.UpdateBusinessPartner)
	eJwt.DELETE("/business-partner/:id", controller.DeleteBusinessPartner)
}