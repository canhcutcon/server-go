package data

import (
	"fmt"
	"log"
	"server-go/internal/configs"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDb(cfg *configs.Config) error {
	var err error
	cnn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Username, cfg.Postgres.Database, cfg.Postgres.Password)

	dbClient, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)

	log.Println("Database connected")
	return nil
}

func GetDbClient() *gorm.DB {
	return dbClient
}

func CloseDb() {
	sqlDb, _ := dbClient.DB()
	sqlDb.Close()
}
