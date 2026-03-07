package config

import (
	"fmt"

	"github.com/PauloPHAL/microservices/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectToDB() {
	cfg := GetConfig()
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.GetHost(), cfg.GetUser(), cfg.GetPassword(), cfg.GetDBName(), cfg.GetPort())

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
}

func SyncDB() {
	DB.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
	)
}

func GetDB() *gorm.DB {
	if DB == nil {
		connectToDB()
	}
	return DB
}
