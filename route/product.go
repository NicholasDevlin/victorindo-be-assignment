package route


import (
	"assignment/feature/controller"
	"assignment/feature/repository"
	"assignment/feature/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProductRoute(e *echo.Echo, db *gorm.DB, eJwt *echo.Group) {
	repository := repository.NewProductRepository(db)
	service := service.NewProductService(repository)
	controller := controller.NewProductController(service)

	eJwt.POST("/product", controller.SaveProduct)
	eJwt.GET("/product", controller.GetAllProduct)
	eJwt.PUT("/product/:id", controller.UpdateProduct)
	eJwt.DELETE("/product/:id", controller.DeleteProduct)
}