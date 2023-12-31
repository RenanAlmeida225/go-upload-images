package config

import (
	"fmt"
	"os"

	"github.com/RenanAlmeida225/go-upload-images/internal/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostegres() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&schemas.User{}); err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&schemas.Confirmation{}); err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&schemas.Images{}); err != nil {
		return nil, err
	}
	return db, nil
}
