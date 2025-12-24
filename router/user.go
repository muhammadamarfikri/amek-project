package router

import (
	"fmt"
	"hotel/delivery/rest/handler"

	"github.com/labstack/echo/v4"
)

func UserRouter(e *echo.Echo, uh *handler.UserHandler) {

	fmt.Println("test router")

	paymentProjectGroup := e.Group("/hotel")
	register := paymentProjectGroup.Group("/user")

	register.POST("/register", uh.RegisterUser)
}
