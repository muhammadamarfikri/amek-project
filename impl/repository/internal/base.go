package internal

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repo struct {
	Db  *gorm.DB
	Log *zap.Logger
}

func NewRepository(db *gorm.DB, log *zap.Logger) *Repo {
	return &Repo{
		Db:  db,
		Log: log,
	}
}
