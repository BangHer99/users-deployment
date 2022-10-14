package factory

import (
	userData "be12/deploy/features/users/data"
	userDelivery "be12/deploy/features/users/delivery"
	userService "be12/deploy/features/users/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {

	userDataFactory := userData.New(db)
	userUsecaseFactory := userService.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

}
