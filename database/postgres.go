package database

import (
	"fmt"
	"hotel/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var (
// 	postgresInstance *gorm.DB
// 	postgresOnce     sync.Once
// )

func InitPostgresDB(config *config.PostgresDB, postgresConf *config.PostgresSync) *gorm.DB {

	if postgresConf == nil {
		log.Fatal("postgresConf is nil")
	}
	if postgresConf == nil {
		log.Fatal("postgres config is nil")
	}

	postgresConf.Once.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			config.Host,
			config.User,
			config.Password,
			config.DBName,
			config.Port,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to open connection to database: ", err)
			return
		}
		log.Println("Successfully connected to Postgres")

		postgresConf.DbInstance = db
	})

	return postgresConf.DbInstance
}
