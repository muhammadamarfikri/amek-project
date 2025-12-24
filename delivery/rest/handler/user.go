package handler

import (
	"fmt"
	"hotel/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type UserHandler struct {
	UserUsecase user.UserUsecase
	Log         *zap.Logger
}

func NewUserHandler(uu user.UserUsecase, log *zap.Logger) *UserHandler {
	return &UserHandler{
		UserUsecase: uu,
		Log:         log,
	}
}

func (uh *UserHandler) RegisterUser(c echo.Context) error {

	fmt.Println("test handler")

	uh.UserUsecase.Registration()

	return c.JSON(http.StatusOK, "Registration success")
}
