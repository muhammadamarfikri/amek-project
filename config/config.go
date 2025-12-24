package config

import (
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	Env                              string
	Port                             string
	Version                          string
	Address                          string
	BaseUrl                          string
	Timezone                         string
	EndpointPrefix                   string
	GracefulShutdownTimeoutInSeconds int
}

type PostgresDB struct {
	Host     string
	Port     string
	User     string
	DBName   string
	SSLMode  string
	Password string
}

type PostgresSync struct {
	DbInstance *gorm.DB
	Once       sync.Once
}

type Logger struct {
	Zap *zap.Logger
}
