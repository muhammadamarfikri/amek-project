package internal

import "go.uber.org/zap"

type UserRepository interface {
	CreateUser()
}

type UserUsecase struct {
	UserRepo UserRepository
	Log      *zap.Logger
}

func NewUserUsecase(ur UserRepository, log *zap.Logger) *UserUsecase {
	return &UserUsecase{
		UserRepo: ur,
		Log:      log,
	}
}
