package database

import (
	"github.com/joho/godotenv"
	"github.com/szmulinho/common/config"
	"github.com/szmulinho/common/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func LoadConfigFromEnv() config.StorageConfig {
	return config.StorageConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}
}

func Connect() (*gorm.DB, error) {
	conn := LoadConfigFromEnv()
	connectionString := conn.ConnectionString()

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate models
	if err := db.AutoMigrate(&model.Prescription{}, &model.Drug{}, &model.User{}, &model.Opinion{}, &model.Order{}, &model.Doctor{}); err != nil {
		return nil, err
	}

	return db, nil
}
