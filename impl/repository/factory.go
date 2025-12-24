package repository

import (
	"hotel/impl/repository/internal"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser()
}

func NewRepository(db *gorm.DB, log *zap.Logger) Repository {
	return internal.NewRepository(db, log)
}
