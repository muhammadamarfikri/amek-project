package user

import (
	"hotel/user/internal"

	"go.uber.org/zap"
)

type UserUsecase interface {
	Registration()
}

func NewUserUsecase(ur internal.UserRepository, log *zap.Logger) UserUsecase {
	return internal.NewUserUsecase(ur, log)
}
